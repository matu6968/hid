/*
 * Amalgamation driver for hidapi + libusb (Linux)
 *
 * This file is preprocessed to produce a single C file containing both
 * libraries, suitable for ccgo (C-to-Go transpiler) or other tools that
 * expect a single translation unit.
 *
 * Build with: ./make_amalgamation.sh
 * Output: hid_libusb_amalgamation.c
 */

#define OS_LINUX 1
#define DEFAULT_VISIBILITY ""
#define POLL_NFDS_TYPE int
#define _GNU_SOURCE 1
#define HAVE_SYS_TIME_H 1
#define HAVE_CLOCK_GETTIME 1

#include <poll.h>

/* libusb - must include libusbi.h first (defines ASSERT_EQ before threads_posix.h) */
#include "libusbi.h"
#include "os/events_posix.h"
#include "os/threads_posix.h"

#include "os/events_posix.c"
#include "os/threads_posix.c"

#include "os/linux_usbfs.c"
#include "os/linux_netlink.c"

#include "core.c"
#include "descriptor.c"
#include "hotplug.c"

#include "io.c"
#include "strerror.c"
#include "sync.c"

/* hidapi libusb backend */
#include "hidapi/libusb/hid.c"
