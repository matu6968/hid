# hid + libusb Amalgamation (Linux)

This directory contains build artifacts for combining hidapi and libusb into a single C file (amalgamation), similar to [SQLite's amalgamation](https://www.sqlite.org/amalgamation.html).

## Purpose

The amalgamation enables:

- **ccgo** (C-to-Go transpiler): Process a single C file to generate pure Go code without cgo on Linux
- **Simpler distribution**: One source file instead of many
- **Potential optimizations**: Single translation unit allows better inlining

## Building

From the project root:

```bash
make amalgamation-32 # or amalgamation-64
```

Output: `amalgamation/hid_amalgamation_linux_32.c` or `amalgamation/hid_amalgamation_linux_64.c` (~10k lines)

Uses `zig cc -E -P` for preprocessing. Linux is the only platform using amalgamation; Windows and macOS use **purego** with embedded shared libraries instead (see `hidlib/`).

## ccgo Usage

[ccgo](https://pkg.go.dev/modernc.org/ccgo/v4) transpiles C to Go. To build the ccgo version of the library:

```bash
go install modernc.org/ccgo/v3@latest
go get modernc.org/libc

# Transpile (from project root)
# For 32 bit targets set GO_ARCH to any 32 bit architecture (like x86 or arm)
export GO_ARCH=386
ccgo -o amalgamation/hid_ccgo.go amalgamation/hid_amalgamation_linux_32.c 
# for 64 bit targets set GO_ARCH to any 64 bit architecture (like amd64 or arm64)
#export GO_ARCH=amd64
#ccgo -o amalgamation/hid_ccgo.go amalgamation/hid_amalgamation_linux_64.c

# Don't forget to rename the top level package name to hid
sed -i -e 's/package main/package hid/g' amalgamation/hid_ccgo.go

```

## cgo-free platform strategy

| Platform | Approach |
|----------|----------|
| **Linux** | ccgo + amalgamation (this directory) |
| **Windows** | purego + embedded DLL (`hidlib/`) |
| **macOS** | purego + embedded dylib (`hidlib/`) |

