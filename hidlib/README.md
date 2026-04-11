# hidlib - Embedded hidapi for purego

This directory contains hidapi shared libraries for cgo-free builds on Windows and macOS.

## Building

From the project root:

```bash
make [hidlib-windows-amd64|hidlib-windows-arm64|hidlib-windows-x86|hidlib-darwin-amd64|hidlib-darwin-arm64|hidlib-all]
```

- **Windows**: Can be built from Linux using zig cc cross-compilation.
- **macOS**: Requires a macOS host (or osxcross) for darwin_amd64/darwin_arm64; zig lacks the macOS SDK for IOKit/CoreFoundation on Linux.

## Contents

| File | Platform | Built by |
|------|----------|----------|
| hidapi_windows_amd64.dll | Windows amd64 | `make hidlib-windows-amd64` |
| hidapi_windows_arm64.dll | Windows amd64 | `make hidlib-windows-arm64` |
| hidapi_windows_x86.dll | Windows amd64 | `make hidlib-windows-x86` |
| hidapi_darwin_amd64.dylib | macOS x86_64 | `make hidlib-darwin-amd64` (on Mac) |
| hidapi_darwin_arm64.dylib | macOS arm64 | `make hidlib-darwin-arm64` (on Mac) |

## Embedding

The purego bindings (`hid_purego_windows.go`, `hid_purego_darwin.go`) use `//go:embed` to bundle these libraries. When present and valid, they are extracted to a temp directory and loaded at runtime. Otherwise, the system library is used (hidapi.dll on Windows, libhidapi.dylib on macOS).

Alternatively, you can use Homebrew on MacOS to install hidapi.
```bash
brew install hidapi
```