# Unified build: Linux amalgamation (ccgo) + hidapi shared libs for purego (Windows/macOS).
# Requires: zig (amalgamation + Windows hidlib); clang (Darwin hidlib, typically on macOS).

PROJECT_ROOT := $(CURDIR)
AMALG_DIR := $(PROJECT_ROOT)/amalgamation
DRIVER := $(AMALG_DIR)/driver_linux.c
HIDAPI_DIR := $(PROJECT_ROOT)/hidapi
HIDLIB_DIR ?= $(PROJECT_ROOT)/hidlib

AMALG_CFLAGS := -DOS_LINUX -D_GNU_SOURCE -DHAVE_SYS_TIME_H -DHAVE_CLOCK_GETTIME
AMALG_CFLAGS += -DDEFAULT_VISIBILITY=\"\" -DPOLL_NFDS_TYPE=int
AMALG_INCLUDES := -I$(PROJECT_ROOT)/libusb/libusb -I$(PROJECT_ROOT)/hidapi/hidapi -I$(PROJECT_ROOT)

.PHONY: all amalgamation amalgamation-32 amalgamation-64 \
	hidlib hidlib-all hidlib-windows-amd64 hidlib-windows-arm64 hidlib-windows-x86 \
	hidlib-darwin-amd64 hidlib-darwin-arm64 help

all: amalgamation hidlib-all

amalgamation: amalgamation-32 amalgamation-64

amalgamation-32: $(AMALG_DIR)/hid_amalgamation_linux_32.c

amalgamation-64: $(AMALG_DIR)/hid_amalgamation_linux_64.c

$(AMALG_DIR)/hid_amalgamation_linux_32.c: $(DRIVER)
	@test -f "$(DRIVER)" || (echo "Driver not found: $(DRIVER)" >&2; exit 1)
	@echo "Building amalgamation for Linux (32-bit)..."
	@echo "  Driver: $(DRIVER)"
	@echo "  Output: $@"
	zig cc -target x86-linux-musl -E -P \
		$(AMALG_INCLUDES) \
		$(AMALG_CFLAGS) \
		"$(DRIVER)" \
		> "$@" 2>/dev/null
	@echo "  Generated $$(wc -l < "$@") lines"
	@echo "Done: $@"

$(AMALG_DIR)/hid_amalgamation_linux_64.c: $(DRIVER)
	@test -f "$(DRIVER)" || (echo "Driver not found: $(DRIVER)" >&2; exit 1)
	@echo "Building amalgamation for Linux..."
	@echo "  Driver: $(DRIVER)"
	@echo "  Output: $@"
	zig cc -target x86_64-linux-musl -E -P \
		$(AMALG_INCLUDES) \
		$(AMALG_CFLAGS) \
		"$(DRIVER)" \
		> "$@" 2>/dev/null
	@echo "  Generated $$(wc -l < "$@") lines"
	@echo "Done: $@"

$(HIDLIB_DIR):
	mkdir -p "$(HIDLIB_DIR)"

$(HIDLIB_DIR)/hidapi_windows_amd64.dll: $(HIDAPI_DIR)/windows/hid.c | $(HIDLIB_DIR)
	@echo "Building hidapi for Windows amd64..."
	zig cc -shared \
		-target x86_64-windows-gnu \
		-o "$@" \
		-I"$(PROJECT_ROOT)/hidapi/hidapi" \
		-I"$(PROJECT_ROOT)/hidapi/windows" \
		-lsetupapi \
		"$(HIDAPI_DIR)/windows/hid.c"
	@echo "  -> $@"

$(HIDLIB_DIR)/hidapi_windows_arm64.dll: $(HIDAPI_DIR)/windows/hid.c | $(HIDLIB_DIR)
	@echo "Building hidapi for Windows arm64..."
	zig cc -shared \
		-target aarch64-windows-gnu \
		-o "$@" \
		-I"$(PROJECT_ROOT)/hidapi/hidapi" \
		-I"$(PROJECT_ROOT)/hidapi/windows" \
		-lsetupapi \
		"$(HIDAPI_DIR)/windows/hid.c"
	@echo "  -> $@"

$(HIDLIB_DIR)/hidapi_windows_x86.dll: $(HIDAPI_DIR)/windows/hid.c | $(HIDLIB_DIR)
	@echo "Building hidapi for Windows x86..."
	zig cc -shared \
		-target x86-windows-gnu \
		-o "$@" \
		-I"$(PROJECT_ROOT)/hidapi/hidapi" \
		-I"$(PROJECT_ROOT)/hidapi/windows" \
		-lsetupapi \
		"$(HIDAPI_DIR)/windows/hid.c"
	@echo "  -> $@"

$(HIDLIB_DIR)/hidapi_darwin_amd64.dylib: $(HIDAPI_DIR)/mac/hid.c | $(HIDLIB_DIR)
ifeq ($(shell uname), Darwin)
	@echo "Building hidapi for macOS x86_64..."
	clang -shared \
		-o "$@" \
		-I"$(PROJECT_ROOT)/hidapi/hidapi" \
		-I"$(PROJECT_ROOT)/hidapi/mac" \
		-framework CoreFoundation \
		-framework IOKit \
		"$(HIDAPI_DIR)/mac/hid.c"
	@echo "  -> $@"
else
	@echo "Error: Building hidapi for macOS x86_64 is not supported on this platform"
	@exit 1
endif

$(HIDLIB_DIR)/hidapi_darwin_arm64.dylib: $(HIDAPI_DIR)/mac/hid.c | $(HIDLIB_DIR)
ifeq ($(shell uname), Darwin)
	@echo "Building hidapi for macOS arm64..."
	clang -shared \
		-o "$@" \
		-I"$(PROJECT_ROOT)/hidapi/hidapi" \
		-I"$(PROJECT_ROOT)/hidapi/mac" \
		-framework CoreFoundation \
		-framework IOKit \
		"$(HIDAPI_DIR)/mac/hid.c"
	@echo "  -> $@"
else
	@echo "Error: Building hidapi for macOS arm64 is not supported on this platform"
	@exit 1
endif

hidlib-windows-amd64: $(HIDLIB_DIR)/hidapi_windows_amd64.dll
hidlib-windows-arm64: $(HIDLIB_DIR)/hidapi_windows_arm64.dll
hidlib-windows-x86: $(HIDLIB_DIR)/hidapi_windows_x86.dll
hidlib-darwin-amd64: $(HIDLIB_DIR)/hidapi_darwin_amd64.dylib
hidlib-darwin-arm64: $(HIDLIB_DIR)/hidapi_darwin_arm64.dylib

hidlib hidlib-all: $(HIDLIB_DIR)/hidapi_windows_amd64.dll $(HIDLIB_DIR)/hidapi_windows_arm64.dll $(HIDLIB_DIR)/hidapi_windows_x86.dll
	@$(MAKE) hidlib-darwin-amd64
	@$(MAKE) hidlib-darwin-arm64
	@echo "Done. Libraries in $(HIDLIB_DIR)"

demo:
	go build -o demo demo.go

help:
	@echo "Targets:"
	@echo "  amalgamation-32 / amalgamation-64  - preprocess driver -> hid_amalgamation_linux_{32,64}.c"
	@echo "  amalgamation                       - both amalgamation outputs"
	@echo "  hidlib / hidlib-all              - Windows DLLs + best-effort Darwin dylibs"
	@echo "  hidlib-windows-amd64|arm64|x86   - single Windows library"
	@echo "  hidlib-darwin-amd64|arm64        - single macOS library (needs macOS toolchain)"
	@echo "  all                                - amalgamation + hidlib-all"
	@echo "  demo                               - build demo binary"
	@echo "Variables: HIDLIB_DIR (default: ./hidlib)"
