// purego bindings for hidapi on Windows (no cgo).
// Loads hidapi from embedded DLL or system. Uses github.com/ebitengine/purego.

//go:build !cgo && windows && arm64

package hid

import (
	_ "embed"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/ebitengine/purego"
)

//go:embed hidlib/hidapi_windows_arm64.dll
var hidapiDll []byte

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

func loadHidapi() error {
	puregoMu.Lock()
	defer puregoMu.Unlock()

	if libHidapi != 0 {
		return nil
	}
	if loadErr != nil {
		return loadErr
	}

	// Try system library first (e.g. hidapi.dll in PATH)
	h, err := syscall.LoadLibrary("hidapi.dll")
	if err == nil {
		libHidapi = uintptr(h)
	} else {
		// Fall back to embedded
		tmpDir := filepath.Join(os.TempDir(), "hid-hidapi")
		if err := os.MkdirAll(tmpDir, 0755); err != nil {
			loadErr = err
			return loadErr
		}
		dllPath := filepath.Join(tmpDir, "hidapi_windows_arm64.dll")
		if err := os.WriteFile(dllPath, hidapiDll, 0755); err != nil {
			loadErr = err
			return loadErr
		}
		runtime.SetFinalizer(new(int), func(_ *int) {
			_ = os.Remove(dllPath)
		})
		h, loadErr = syscall.LoadLibrary(dllPath)
		if loadErr != nil {
			loadErr = errors.New("hid: failed to load hidapi.dll (embedded or system)")
			return loadErr
		}
		libHidapi = uintptr(h)
	}
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

// hid_device_info C struct layout (Windows arm64 ABI)
type hidDeviceInfo struct {
	path                uintptr
	vendor_id           uint16
	product_id          uint16
	serial_number       uintptr
	release_number      uint16
	manufacturer_string uintptr
	product_string      uintptr
	usage_page          uint16
	usage               uint16
	interface_number    int32
	next                uintptr
	_                   uintptr // bus_type
}

func wcharPtrToString(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var u16 []uint16
	for {
		ch := *(*uint16)(unsafe.Pointer(ptr))
		if ch == 0 {
			break
		}
		u16 = append(u16, ch)
		ptr += 2
	}
	return string(utf16.Decode(u16))
}

func Supported() bool {
	return loadHidapi() == nil
}

func Enumerate(vendorID uint16, productID uint16) ([]DeviceInfo, error) {
	if err := loadHidapi(); err != nil {
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
		info := (*hidDeviceInfo)(unsafe.Pointer(cur))
		di := DeviceInfo{
			Path:      goString(info.path),
			VendorID:  info.vendor_id,
			ProductID: info.product_id,
			Release:   info.release_number,
			UsagePage: info.usage_page,
			Usage:     info.usage,
			Interface: int(info.interface_number),
		}
		if info.serial_number != 0 {
			di.Serial = wcharPtrToString(info.serial_number)
		}
		if info.manufacturer_string != 0 {
			di.Manufacturer = wcharPtrToString(info.manufacturer_string)
		}
		if info.product_string != 0 {
			di.Product = wcharPtrToString(info.product_string)
		}
		infos = append(infos, di)
		cur = info.next
	}
	return infos, nil
}

func (info DeviceInfo) Open() (Device, error) {
	if err := loadHidapi(); err != nil {
		return nil, err
	}
	enumerateMu.Lock()
	defer enumerateMu.Unlock()

	path, err := cString(info.Path)
	if err != nil {
		return nil, err
	}
	defer free(path)

	dev := hid_open_path(path)
	if dev == 0 {
		return nil, errors.New("hidapi: failed to open device")
	}
	return &puregoDevice{DeviceInfo: info, dev: dev}, nil
}

type puregoDevice struct {
	DeviceInfo
	dev uintptr
	mu  sync.Mutex
}

func (d *puregoDevice) Close() error {
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

func (d *puregoDevice) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	// Windows expects report ID as first byte
	report := append([]byte{0x00}, b...)
	written := callWithBytes(dev, report, func(dev, data, size uintptr) int {
		return hid_write(dev, data, size)
	})
	if written < 0 {
		return 0, lastError(dev)
	}
	if written > 0 {
		written--
	}
	return written, nil
}

func (d *puregoDevice) Read(b []byte) (int, error) {
	return d.ReadTimeout(b, -1)
}

func (d *puregoDevice) ReadTimeout(b []byte, timeout int) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	read := callWithBytes(dev, b, func(dev, data, size uintptr) int {
		return hid_read_timeout(dev, data, size, timeout)
	})
	if read >= 0 {
		return read, nil
	}
	return 0, lastError(dev)
}

func (d *puregoDevice) GetFeatureReport(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	read := callWithBytes(dev, b, func(dev, data, size uintptr) int {
		return hid_get_feature_report(dev, data, size)
	})
	if read >= 0 {
		return read, nil
	}
	return 0, lastError(dev)
}

func (d *puregoDevice) SendFeatureReport(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	written := callWithBytes(dev, b, func(dev, data, size uintptr) int {
		return hid_send_feature_report(dev, data, size)
	})
	if written < 0 {
		return 0, lastError(dev)
	}
	return written, nil
}

func callWithBytes(dev uintptr, b []byte, fn func(dev, data, size uintptr) int) int {
	puregoMu.Lock()
	defer puregoMu.Unlock()
	data := unsafe.Pointer(unsafe.SliceData(b))
	return fn(dev, uintptr(data), uintptr(len(b)))
}

func lastError(dev uintptr) error {
	puregoMu.Lock()
	msg := hid_error(dev)
	puregoMu.Unlock()
	if msg != 0 {
		s := wcharPtrToString(msg)
		if s != "" {
			return errors.New("hidapi: " + s)
		}
	}
	return errors.New("hidapi: unknown failure")
}

func goString(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}
	var n int
	for *(*byte)(unsafe.Pointer(ptr + uintptr(n))) != 0 {
		n++
	}
	return string(unsafe.Slice((*byte)(unsafe.Pointer(ptr)), n))
}

func cString(s string) (uintptr, error) {
	b := append([]byte(s), 0)
	return uintptr(unsafe.Pointer(unsafe.SliceData(b))), nil
}

func free(uintptr) {}
