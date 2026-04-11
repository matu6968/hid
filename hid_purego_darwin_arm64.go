// purego bindings for hidapi on macOS (no cgo).
// Loads hidapi from embedded dylib or system. Uses github.com/ebitengine/purego.
//
// To embed: run ./make_hidlib.sh darwin_amd64 and darwin_arm64 on a Mac,
// then the embedded libraries will be used. Otherwise loads from system
// (libhidapi.dylib, libhidapi.0.dylib, or /usr/local/lib/libhidapi.dylib).

//go:build !cgo && darwin && arm64 && !ios

package hid

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"unicode/utf8"
	"unsafe"

	"github.com/ebitengine/purego"
)

//go:embed hidlib/hidapi_darwin_arm64.dylib
var darwinEmbeddedDylib []byte

var (
	libHidapi   uintptr
	puregoMu    sync.Mutex
	enumerateMu sync.Mutex
	loadErr     error
)

var (
	hid_init                func() int
	hid_exit                func() int
	hid_enumerate           func(vendorID, productID uint16) uintptr
	hid_free_enumeration    func(devs uintptr)
	hid_open_path           func(path uintptr) uintptr
	hid_close               func(dev uintptr)
	hid_write               func(dev uintptr, data uintptr, length uintptr) int
	hid_read_timeout        func(dev uintptr, data uintptr, length uintptr, ms int) int
	hid_get_feature_report  func(dev uintptr, data uintptr, length uintptr) int
	hid_send_feature_report func(dev uintptr, data uintptr, length uintptr) int
	hid_error               func(dev uintptr) uintptr
)

func loadHidapiDarwin() error {
	puregoMu.Lock()
	defer puregoMu.Unlock()

	if libHidapi != 0 {
		return nil
	}
	if loadErr != nil {
		return loadErr
	}

	// Try embedded library first (darwinEmbeddedDylib set by arch-specific embed file)
	if len(darwinEmbeddedDylib) > 0 {
		tmpDir := filepath.Join(os.TempDir(), "hid-hidapi")
		if err := os.MkdirAll(tmpDir, 0755); err == nil {
			name := "libhidapi_darwin.dylib"
			if runtime.GOARCH == "arm64" {
				name = "libhidapi_darwin_arm64.dylib"
			} else {
				name = "libhidapi_darwin_amd64.dylib"
			}
			dylibPath := filepath.Join(tmpDir, name)
			if err := os.WriteFile(dylibPath, darwinEmbeddedDylib, 0755); err == nil {
				libHidapi, loadErr = purego.Dlopen(dylibPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
				if loadErr == nil {
					runtime.SetFinalizer(new(int), func(_ *int) { _ = os.Remove(dylibPath) })
					goto register
				}
			}
		}
	}

	// Try system library
	for _, name := range []string{"libhidapi.dylib", "libhidapi.0.dylib"} {
		libHidapi, loadErr = purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if loadErr == nil {
			break
		}
	}
	if loadErr != nil {
		for _, p := range []string{"/usr/local/lib/libhidapi.dylib", "/opt/homebrew/lib/libhidapi.dylib"} {
			libHidapi, loadErr = purego.Dlopen(p, purego.RTLD_NOW|purego.RTLD_GLOBAL)
			if loadErr == nil {
				break
			}
		}
	}
	if loadErr != nil {
		loadErr = fmt.Errorf("hid: failed to load libhidapi: %w", loadErr)
		return loadErr
	}

register:

	purego.RegisterLibFunc(&hid_init, libHidapi, "hid_init")
	purego.RegisterLibFunc(&hid_exit, libHidapi, "hid_exit")
	purego.RegisterLibFunc(&hid_enumerate, libHidapi, "hid_enumerate")
	purego.RegisterLibFunc(&hid_free_enumeration, libHidapi, "hid_free_enumeration")
	purego.RegisterLibFunc(&hid_open_path, libHidapi, "hid_open_path")
	purego.RegisterLibFunc(&hid_close, libHidapi, "hid_close")
	purego.RegisterLibFunc(&hid_write, libHidapi, "hid_write")
	purego.RegisterLibFunc(&hid_read_timeout, libHidapi, "hid_read_timeout")
	purego.RegisterLibFunc(&hid_get_feature_report, libHidapi, "hid_get_feature_report")
	purego.RegisterLibFunc(&hid_send_feature_report, libHidapi, "hid_send_feature_report")
	purego.RegisterLibFunc(&hid_error, libHidapi, "hid_error")
	return nil
}

// hid_device_info C struct layout (macOS, wchar_t is 4 bytes)
type hidDeviceInfoDarwin struct {
	path                uintptr
	vendor_id           uint16
	product_id          uint16
	_                   uint32
	serial_number       uintptr
	release_number      uint16
	_                   uint16
	manufacturer_string uintptr
	product_string      uintptr
	usage_page          uint16
	usage               uint16
	interface_number    int32
	_                   int32
	next                uintptr
	bus_type            int32
}

func wcharPtrToStringDarwin(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var b []byte
	for {
		ch := *(*uint32)(unsafe.Pointer(ptr))
		if ch == 0 {
			break
		}
		r := rune(ch)
		if !utf8.ValidRune(r) {
			return ""
		}
		var buf [utf8.UTFMax]byte
		n := utf8.EncodeRune(buf[:], r)
		b = append(b, buf[:n]...)
		ptr += 4
	}
	return string(b)
}

func Supported() bool {
	return loadHidapiDarwin() == nil
}

func Enumerate(vendorID uint16, productID uint16) ([]DeviceInfo, error) {
	if err := loadHidapiDarwin(); err != nil {
		return nil, err
	}
	enumerateMu.Lock()
	defer enumerateMu.Unlock()

	head := hid_enumerate(vendorID, productID)
	if head == 0 {
		return nil, nil
	}
	defer hid_free_enumeration(head)

	var infos []DeviceInfo
	for cur := head; cur != 0; {
		info := (*hidDeviceInfoDarwin)(unsafe.Pointer(cur))
		di := DeviceInfo{
			Path:      goStringDarwin(info.path),
			VendorID:  info.vendor_id,
			ProductID: info.product_id,
			Release:   info.release_number,
			UsagePage: info.usage_page,
			Usage:     info.usage,
			Interface: int(info.interface_number),
		}
		if info.serial_number != 0 {
			di.Serial = wcharPtrToStringDarwin(info.serial_number)
		}
		if info.manufacturer_string != 0 {
			di.Manufacturer = wcharPtrToStringDarwin(info.manufacturer_string)
		}
		if info.product_string != 0 {
			di.Product = wcharPtrToStringDarwin(info.product_string)
		}
		infos = append(infos, di)
		cur = info.next
	}
	return infos, nil
}

func (info DeviceInfo) Open() (Device, error) {
	if err := loadHidapiDarwin(); err != nil {
		return nil, err
	}
	enumerateMu.Lock()
	defer enumerateMu.Unlock()

	path, err := cStringDarwin(info.Path)
	if err != nil {
		return nil, err
	}
	defer freeDarwin(path)

	dev := hid_open_path(path)
	if dev == 0 {
		return nil, fmt.Errorf("hidapi: failed to open device")
	}
	return &puregoDeviceDarwin{DeviceInfo: info, dev: dev}, nil
}

type puregoDeviceDarwin struct {
	DeviceInfo
	dev uintptr
	mu  sync.Mutex
}

func (d *puregoDeviceDarwin) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.dev != 0 {
		puregoMu.Lock()
		hid_close(d.dev)
		puregoMu.Unlock()
		d.dev = 0
	}
	return nil
}

func (d *puregoDeviceDarwin) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	// macOS does not prepend report ID
	written := callWithBytesDarwin(dev, b, func(dev, data, size uintptr) int {
		return hid_write(dev, data, size)
	})
	if written < 0 {
		return 0, lastErrorDarwin(dev)
	}
	return written, nil
}

func (d *puregoDeviceDarwin) Read(b []byte) (int, error) {
	return d.ReadTimeout(b, -1)
}

func (d *puregoDeviceDarwin) ReadTimeout(b []byte, timeout int) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	read := callWithBytesDarwin(dev, b, func(dev, data, size uintptr) int {
		return hid_read_timeout(dev, data, size, timeout)
	})
	if read >= 0 {
		return read, nil
	}
	return 0, lastErrorDarwin(dev)
}

func (d *puregoDeviceDarwin) GetFeatureReport(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	read := callWithBytesDarwin(dev, b, func(dev, data, size uintptr) int {
		return hid_get_feature_report(dev, data, size)
	})
	if read >= 0 {
		return read, nil
	}
	return 0, lastErrorDarwin(dev)
}

func (d *puregoDeviceDarwin) SendFeatureReport(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	written := callWithBytesDarwin(dev, b, func(dev, data, size uintptr) int {
		return hid_send_feature_report(dev, data, size)
	})
	if written < 0 {
		return 0, lastErrorDarwin(dev)
	}
	return written, nil
}

func callWithBytesDarwin(dev uintptr, b []byte, fn func(dev, data, size uintptr) int) int {
	puregoMu.Lock()
	defer puregoMu.Unlock()
	data := unsafe.Pointer(unsafe.SliceData(b))
	return fn(dev, uintptr(data), uintptr(len(b)))
}

func lastErrorDarwin(dev uintptr) error {
	puregoMu.Lock()
	msg := hid_error(dev)
	puregoMu.Unlock()
	if msg != 0 {
		s := wcharPtrToStringDarwin(msg)
		if s != "" {
			return fmt.Errorf("hidapi: %s", s)
		}
	}
	return fmt.Errorf("hidapi: unknown failure")
}

func goStringDarwin(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var n int
	for *(*byte)(unsafe.Pointer(ptr + uintptr(n))) != 0 {
		n++
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(ptr)), n))
}

func cStringDarwin(s string) (uintptr, error) {
	b := append([]byte(s), 0)
	return uintptr(unsafe.Pointer(unsafe.SliceData(b))), nil
}

func freeDarwin(uintptr) {}
