// nocgo adapter layer for hidapi on Linux using ccgo-transpiled code.

//go:build !cgo && linux

package hid

import (
	"errors"
	"sync"
	"unsafe"

	"modernc.org/libc"
)

var (
	nocgoTLS     *libc.TLS
	nocgoTLSOnce sync.Once
	nocgoMu      sync.Mutex
)

func getNocgoTLS() *libc.TLS {
	nocgoTLSOnce.Do(func() {
		nocgoTLS = libc.NewTLS()
	})
	return nocgoTLS
}

// Supported returns whether this platform is supported by the HID library or not.
func Supported() bool {
	return true
}

// Enumerate returns a list of all the HID devices attached to the system which
// match the vendor and product id.
func Enumerate(vendorID uint16, productID uint16) ([]DeviceInfo, error) {
	nocgoMu.Lock()
	defer nocgoMu.Unlock()

	tls := getNocgoTLS()
	head := hid_enumerate(tls, vendorID, productID)
	if head == 0 {
		return nil, nil
	}
	defer hid_free_enumeration(tls, head)

	var infos []DeviceInfo
	for cur := head; cur != 0; cur = (*hid_device_info)(unsafe.Pointer(cur)).Fnext {
		info := (*hid_device_info)(unsafe.Pointer(cur))
		di := DeviceInfo{
			Path:      libc.GoString(info.Fpath),
			VendorID:  info.Fvendor_id,
			ProductID: info.Fproduct_id,
			Release:   info.Frelease_number,
			UsagePage: info.Fusage_page,
			Usage:     info.Fusage,
			Interface: int(info.Finterface_number),
		}
		if info.Fserial_number != 0 {
			di.Serial, _ = wcharTToStringNocgo(info.Fserial_number)
		}
		if info.Fproduct_string != 0 {
			di.Product, _ = wcharTToStringNocgo(info.Fproduct_string)
		}
		if info.Fmanufacturer_string != 0 {
			di.Manufacturer, _ = wcharTToStringNocgo(info.Fmanufacturer_string)
		}
		infos = append(infos, di)
	}
	return infos, nil
}

// Open connects to a previously discovered HID device.
func (info DeviceInfo) Open() (Device, error) {
	nocgoMu.Lock()
	defer nocgoMu.Unlock()

	tls := getNocgoTLS()
	pathPtr, err := libc.CString(info.Path)
	if err != nil {
		return nil, err
	}
	defer libc.Xfree(nil, pathPtr)

	dev := hid_open_path(tls, pathPtr)
	if dev == 0 {
		return nil, errors.New("hidapi: failed to open device")
	}
	return &nocgoDevice{
		DeviceInfo: info,
		dev:        dev,
	}, nil
}

// nocgoDevice implements Device using the ccgo hidapi.
type nocgoDevice struct {
	DeviceInfo
	dev uintptr
	mu  sync.Mutex
}

func (d *nocgoDevice) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.dev != 0 {
		nocgoMu.Lock()
		hid_close(getNocgoTLS(), d.dev)
		nocgoMu.Unlock()
		d.dev = 0
	}
	return nil
}

func (d *nocgoDevice) Write(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()

	if dev == 0 {
		return 0, ErrDeviceClosed
	}

	// Linux does not prepend report ID
	nocgoMu.Lock()
	tls := getNocgoTLS()
	dataPtr := tls.Alloc(len(b))
	defer tls.Free(len(b))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dataPtr)), len(b)), b)
	written := hid_write(tls, dev, dataPtr, size_t(len(b)))
	nocgoMu.Unlock()

	if written < 0 {
		d.mu.Lock()
		dev = d.dev
		d.mu.Unlock()
		if dev == 0 {
			return 0, ErrDeviceClosed
		}
		nocgoMu.Lock()
		msg := hid_error(tls, dev)
		nocgoMu.Unlock()
		if msg != 0 {
			failure, _ := wcharTToStringNocgo(msg)
			return 0, errors.New("hidapi: " + failure)
		}
		return 0, errors.New("hidapi: unknown failure")
	}
	return int(written), nil
}

func (d *nocgoDevice) Read(b []byte) (int, error) {
	return d.ReadTimeout(b, -1)
}

func (d *nocgoDevice) ReadTimeout(b []byte, timeout int) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()

	if dev == 0 {
		return 0, ErrDeviceClosed
	}

	nocgoMu.Lock()
	tls := getNocgoTLS()
	dataPtr := tls.Alloc(len(b))
	defer tls.Free(len(b))
	read := hid_read_timeout(tls, dev, dataPtr, size_t(len(b)), int32(timeout))
	nocgoMu.Unlock()

	if read >= 0 {
		copy(b, unsafe.Slice((*byte)(unsafe.Pointer(dataPtr)), read))
		return int(read), nil
	}

	d.mu.Lock()
	dev = d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	nocgoMu.Lock()
	msg := hid_error(tls, dev)
	nocgoMu.Unlock()
	if msg != 0 {
		failure, _ := wcharTToStringNocgo(msg)
		return 0, errors.New("hidapi: " + failure)
	}
	return 0, errors.New("hidapi: unknown failure")
}

func (d *nocgoDevice) GetFeatureReport(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()

	if dev == 0 {
		return 0, ErrDeviceClosed
	}

	nocgoMu.Lock()
	tls := getNocgoTLS()
	dataPtr := tls.Alloc(len(b))
	defer tls.Free(len(b))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dataPtr)), len(b)), b)
	read := hid_get_feature_report(tls, dev, dataPtr, size_t(len(b)))
	nocgoMu.Unlock()

	if read >= 0 {
		copy(b, unsafe.Slice((*byte)(unsafe.Pointer(dataPtr)), read))
		return int(read), nil
	}

	d.mu.Lock()
	dev = d.dev
	d.mu.Unlock()
	if dev == 0 {
		return 0, ErrDeviceClosed
	}
	nocgoMu.Lock()
	msg := hid_error(tls, dev)
	nocgoMu.Unlock()
	if msg != 0 {
		failure, _ := wcharTToStringNocgo(msg)
		return 0, errors.New("hidapi: " + failure)
	}
	return 0, errors.New("hidapi: unknown failure")
}

func (d *nocgoDevice) SendFeatureReport(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d.mu.Lock()
	dev := d.dev
	d.mu.Unlock()

	if dev == 0 {
		return 0, ErrDeviceClosed
	}

	nocgoMu.Lock()
	tls := getNocgoTLS()
	dataPtr := tls.Alloc(len(b))
	defer tls.Free(len(b))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dataPtr)), len(b)), b)
	written := hid_send_feature_report(tls, dev, dataPtr, size_t(len(b)))
	nocgoMu.Unlock()

	if written < 0 {
		d.mu.Lock()
		dev = d.dev
		d.mu.Unlock()
		if dev == 0 {
			return 0, ErrDeviceClosed
		}
		nocgoMu.Lock()
		msg := hid_error(tls, dev)
		nocgoMu.Unlock()
		if msg != 0 {
			failure, _ := wcharTToStringNocgo(msg)
			return 0, errors.New("hidapi: " + failure)
		}
		return 0, errors.New("hidapi: unknown failure")
	}
	return int(written), nil
}
