// Pure Go wchar handling for nocgo Linux build.
// On Linux, wchar_t is 4 bytes (UTF-32).

//go:build !cgo && linux

package hid

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

// wcharTToString converts a wchar_t* (Linux UTF-32) to a Go string.
// wcsptr is a uintptr pointing to a null-terminated array of uint32.
func wcharTToStringNocgo(wcsptr uintptr) (string, error) {
	if wcsptr == 0 {
		return "", nil
	}
	var b []byte
	for {
		ch := *(*uint32)(unsafe.Pointer(wcsptr))
		if ch == 0 {
			break
		}
		r := rune(ch)
		if !utf8.ValidRune(r) {
			return "", fmt.Errorf("invalid rune at wchar_t* offset %d", len(b))
		}
		var buf [utf8.UTFMax]byte
		n := utf8.EncodeRune(buf[:], r)
		b = append(b, buf[:n]...)
		wcsptr += 4 // sizeof(wchar_t) on Linux
	}
	return string(b), nil
}
