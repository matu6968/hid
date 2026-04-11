// Code generated for linux/s390x by 'ccgo -o /zram/hid/hid_ccgo_linux_s390x.go /zram/hid/amalgamation/hid_libusb_amalgamation.c', DO NOT EDIT.

//go:build !cgo && linux && s390x

package hid

import (
	"reflect"
	"unsafe"

	"modernc.org/libc"
)

var _ reflect.Type
var _ unsafe.Pointer

const _GNU_SOURCE = 1
const _LP64 = 1
const _STDC_PREDEF_H = 1
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BITINT_MAXWIDTH__ = 65535
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 2
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 14
const __GXX_ABI_VERSION = 1019
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __NO_INLINE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE2__ = 1
const __SSE_MATH__ = 1
const __SSE__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_UTF_16__ = 1
const __STDC_UTF_32__ = 1
const __STDC_VERSION__ = 201710
const __STDC__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __VERSION__ = "14.2.0"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const linux = 1
const unix = 1

type __builtin_va_list = uintptr

type __predefined_size_t = uint64

type __predefined_wchar_t = int32

type __predefined_ptrdiff_t = int64

type nfds_t = uint64

type pollfd = struct {
	Ffd      int32
	Fevents  int16
	Frevents int16
}

type time_t = int64

type timespec = struct {
	Ftv_sec  time_t
	Ftv_nsec int64
}

type sigset_t = struct {
	F__bits [16]uint64
}

type __sigset_t = sigset_t

type uintptr_t = uint64

type intptr_t = int64

type int8_t = int8

type int16_t = int16

type int32_t = int32

type int64_t = int64

type intmax_t = int64

type uint8_t = uint8

type uint16_t = uint16

type uint32_t = uint32

type uint64_t = uint64

type uintmax_t = uint64

type int_fast8_t = int8

type int_fast64_t = int64

type int_least8_t = int8

type int_least16_t = int16

type int_least32_t = int32

type int_least64_t = int64

type uint_fast8_t = uint8

type uint_fast64_t = uint64

type uint_least8_t = uint8

type uint_least16_t = uint16

type uint_least32_t = uint32

type uint_least64_t = uint64

type int_fast16_t = int32

type int_fast32_t = int32

type uint_fast16_t = uint32

type uint_fast32_t = uint32

type wchar_t = int32

type imaxdiv_t = struct {
	Fquot intmax_t
	Frem  intmax_t
}

type va_list = uintptr

type max_align_t = struct {
	F__ll int64
	F__ld float64
}

type size_t = uint64

type ptrdiff_t = int64

type div_t = struct {
	Fquot int32
	Frem  int32
}

type ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type lldiv_t = struct {
	Fquot int64
	Frem  int64
}

type suseconds_t = int64

type timeval = struct {
	Ftv_sec  time_t
	Ftv_usec suseconds_t
}

type fd_mask = uint64

type fd_set = struct {
	Ffds_bits [16]uint64
}

type itimerval = struct {
	Fit_interval timeval
	Fit_value    timeval
}

type timezone1 = struct {
	Ftz_minuteswest int32
	Ftz_dsttime     int32
}

type ssize_t = int64

type register_t = int64

type u_int64_t = uint64

type mode_t = uint32

type nlink_t = uint64

type off_t = int64

type ino_t = uint64

type dev_t = uint64

type blksize_t = int64

type blkcnt_t = int64

type fsblkcnt_t = uint64

type fsfilcnt_t = uint64

type timer_t = uintptr

type clockid_t = int32

type clock_t = int64

type pid_t = int32

type id_t = uint32

type uid_t = uint32

type gid_t = uint32

type key_t = int32

type useconds_t = uint32

type pthread_t = uintptr

type pthread_once_t = int32

type pthread_key_t = uint32

type pthread_spinlock_t = int32

type pthread_mutexattr_t = struct {
	F__attr uint32
}

type pthread_condattr_t = struct {
	F__attr uint32
}

type pthread_barrierattr_t = struct {
	F__attr uint32
}

type pthread_rwlockattr_t = struct {
	F__attr [2]uint32
}

type pthread_attr_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__s  [0][7]uint64
		F__i  [14]int32
	}
}

type pthread_mutex_t = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

type pthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][6]uintptr
		F__i  [12]int32
	}
}

type pthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__p  [0][7]uintptr
		F__i  [14]int32
	}
}

type pthread_barrier_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][4]uintptr
		F__i  [8]int32
	}
}

type u_int8_t = uint8

type u_int16_t = uint16

type u_int32_t = uint32

type caddr_t = uintptr

type u_char = uint8

type u_short = uint16

type ushort = uint16

type u_int = uint32

type uint1 = uint32

type u_long = uint64

type ulong = uint64

type quad_t = int64

type u_quad_t = uint64

func __bswap16(tls *libc.TLS, __x uint16_t) (r uint16_t) {
	return libc.Uint16FromInt32(libc.Int32FromUint16(__x)<<int32(8) | libc.Int32FromUint16(__x)>>int32(8))
}

func __bswap32(tls *libc.TLS, __x uint32_t) (r uint32_t) {
	return __x>>int32(24) | __x>>int32(8)&uint32(0xff00) | __x<<int32(8)&uint32(0xff0000) | __x<<int32(24)
}

func __bswap64(tls *libc.TLS, __x uint64_t) (r uint64_t) {
	return (uint64(__bswap32(tls, uint32(__x)))+0)<<int32(32) | uint64(__bswap32(tls, uint32(__x>>int32(32))))
}

type locale_t = uintptr

type tm = struct {
	Ftm_sec    int32
	Ftm_min    int32
	Ftm_hour   int32
	Ftm_mday   int32
	Ftm_mon    int32
	Ftm_year   int32
	Ftm_wday   int32
	Ftm_yday   int32
	Ftm_isdst  int32
	Ftm_gmtoff int64
	Ftm_zone   uintptr
}

type itimerspec = struct {
	Fit_interval timespec
	Fit_value    timespec
}

func libusb_cpu_to_le16(tls *libc.TLS, x uint16_t) (r uint16_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* _tmp at bp+0 */ struct {
		Fb16 [0]uint16_t
		Fb8  [2]uint8_t
	}
	*(*uint8_t)(unsafe.Pointer(bp + 1)) = libc.Uint8FromInt32(libc.Int32FromUint16(x) >> libc.Int32FromInt32(8))
	*(*uint8_t)(unsafe.Pointer(bp)) = libc.Uint8FromInt32(libc.Int32FromUint16(x) & libc.Int32FromInt32(0xff))
	return *(*uint16_t)(unsafe.Pointer(bp))
}

type libusb_class_code = int32

const LIBUSB_CLASS_PER_INTERFACE = 0
const LIBUSB_CLASS_AUDIO = 1
const LIBUSB_CLASS_COMM = 2
const LIBUSB_CLASS_HID = 3
const LIBUSB_CLASS_PHYSICAL = 5
const LIBUSB_CLASS_IMAGE = 6
const LIBUSB_CLASS_PTP = 6
const LIBUSB_CLASS_PRINTER = 7
const LIBUSB_CLASS_MASS_STORAGE = 8
const LIBUSB_CLASS_HUB = 9
const LIBUSB_CLASS_DATA = 10
const LIBUSB_CLASS_SMART_CARD = 11
const LIBUSB_CLASS_CONTENT_SECURITY = 13
const LIBUSB_CLASS_VIDEO = 14
const LIBUSB_CLASS_PERSONAL_HEALTHCARE = 15
const LIBUSB_CLASS_AUDIO_VIDEO = 16
const LIBUSB_CLASS_BILLBOARD = 17
const LIBUSB_CLASS_TYPE_C_BRIDGE = 18
const LIBUSB_CLASS_BULK_DISPLAY_PROTOCOL = 19
const LIBUSB_CLASS_MCTP = 20
const LIBUSB_CLASS_I3C = 60
const LIBUSB_CLASS_DIAGNOSTIC_DEVICE = 220
const LIBUSB_CLASS_WIRELESS = 224
const LIBUSB_CLASS_MISCELLANEOUS = 239
const LIBUSB_CLASS_APPLICATION = 254
const LIBUSB_CLASS_VENDOR_SPEC = 255

type libusb_descriptor_type = int32

const LIBUSB_DT_DEVICE = 1
const LIBUSB_DT_CONFIG = 2
const LIBUSB_DT_STRING = 3
const LIBUSB_DT_INTERFACE = 4
const LIBUSB_DT_ENDPOINT = 5
const LIBUSB_DT_INTERFACE_ASSOCIATION = 11
const LIBUSB_DT_BOS = 15
const LIBUSB_DT_DEVICE_CAPABILITY = 16
const LIBUSB_DT_HID = 33
const LIBUSB_DT_REPORT = 34
const LIBUSB_DT_PHYSICAL = 35
const LIBUSB_DT_HUB = 41
const LIBUSB_DT_SUPERSPEED_HUB = 42
const LIBUSB_DT_SS_ENDPOINT_COMPANION = 48

type libusb_endpoint_direction = int32

const LIBUSB_ENDPOINT_OUT = 0
const LIBUSB_ENDPOINT_IN = 128

type libusb_endpoint_transfer_type = int32

const LIBUSB_ENDPOINT_TRANSFER_TYPE_CONTROL = 0
const LIBUSB_ENDPOINT_TRANSFER_TYPE_ISOCHRONOUS = 1
const LIBUSB_ENDPOINT_TRANSFER_TYPE_BULK = 2
const LIBUSB_ENDPOINT_TRANSFER_TYPE_INTERRUPT = 3

type libusb_standard_request = int32

const LIBUSB_REQUEST_GET_STATUS = 0
const LIBUSB_REQUEST_CLEAR_FEATURE = 1
const LIBUSB_REQUEST_SET_FEATURE = 3
const LIBUSB_REQUEST_SET_ADDRESS = 5
const LIBUSB_REQUEST_GET_DESCRIPTOR = 6
const LIBUSB_REQUEST_SET_DESCRIPTOR = 7
const LIBUSB_REQUEST_GET_CONFIGURATION = 8
const LIBUSB_REQUEST_SET_CONFIGURATION = 9
const LIBUSB_REQUEST_GET_INTERFACE = 10
const LIBUSB_REQUEST_SET_INTERFACE = 11
const LIBUSB_REQUEST_SYNCH_FRAME = 12
const LIBUSB_REQUEST_SET_SEL = 48
const LIBUSB_SET_ISOCH_DELAY = 49

type libusb_request_type = int32

const LIBUSB_REQUEST_TYPE_STANDARD = 0
const LIBUSB_REQUEST_TYPE_CLASS = 32
const LIBUSB_REQUEST_TYPE_VENDOR = 64
const LIBUSB_REQUEST_TYPE_RESERVED = 96

type libusb_request_recipient = int32

const LIBUSB_RECIPIENT_DEVICE = 0
const LIBUSB_RECIPIENT_INTERFACE = 1
const LIBUSB_RECIPIENT_ENDPOINT = 2
const LIBUSB_RECIPIENT_OTHER = 3

type libusb_iso_sync_type = int32

const LIBUSB_ISO_SYNC_TYPE_NONE = 0
const LIBUSB_ISO_SYNC_TYPE_ASYNC = 1
const LIBUSB_ISO_SYNC_TYPE_ADAPTIVE = 2
const LIBUSB_ISO_SYNC_TYPE_SYNC = 3

type libusb_iso_usage_type = int32

const LIBUSB_ISO_USAGE_TYPE_DATA = 0
const LIBUSB_ISO_USAGE_TYPE_FEEDBACK = 1
const LIBUSB_ISO_USAGE_TYPE_IMPLICIT = 2

type libusb_supported_speed = int32

const LIBUSB_LOW_SPEED_OPERATION = 1
const LIBUSB_FULL_SPEED_OPERATION = 2
const LIBUSB_HIGH_SPEED_OPERATION = 4
const LIBUSB_SUPER_SPEED_OPERATION = 8

type libusb_usb_2_0_extension_attributes = int32

const LIBUSB_BM_LPM_SUPPORT = 2

type libusb_ss_usb_device_capability_attributes = int32

const LIBUSB_BM_LTM_SUPPORT = 2

type libusb_bos_type = int32

const LIBUSB_BT_WIRELESS_USB_DEVICE_CAPABILITY = 1
const LIBUSB_BT_USB_2_0_EXTENSION = 2
const LIBUSB_BT_SS_USB_DEVICE_CAPABILITY = 3
const LIBUSB_BT_CONTAINER_ID = 4
const LIBUSB_BT_PLATFORM_DESCRIPTOR = 5
const LIBUSB_BT_SUPERSPEED_PLUS_CAPABILITY = 10

type libusb_device_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbcdUSB             uint16_t
	FbDeviceClass       uint8_t
	FbDeviceSubClass    uint8_t
	FbDeviceProtocol    uint8_t
	FbMaxPacketSize0    uint8_t
	FidVendor           uint16_t
	FidProduct          uint16_t
	FbcdDevice          uint16_t
	FiManufacturer      uint8_t
	FiProduct           uint8_t
	FiSerialNumber      uint8_t
	FbNumConfigurations uint8_t
}

type libusb_endpoint_descriptor = struct {
	FbLength          uint8_t
	FbDescriptorType  uint8_t
	FbEndpointAddress uint8_t
	FbmAttributes     uint8_t
	FwMaxPacketSize   uint16_t
	FbInterval        uint8_t
	FbRefresh         uint8_t
	FbSynchAddress    uint8_t
	Fextra            uintptr
	Fextra_length     int32
}

type libusb_interface_association_descriptor = struct {
	FbLength           uint8_t
	FbDescriptorType   uint8_t
	FbFirstInterface   uint8_t
	FbInterfaceCount   uint8_t
	FbFunctionClass    uint8_t
	FbFunctionSubClass uint8_t
	FbFunctionProtocol uint8_t
	FiFunction         uint8_t
}

type libusb_interface_association_descriptor_array = struct {
	Fiad    uintptr
	Flength int32
}

type libusb_interface_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbInterfaceNumber   uint8_t
	FbAlternateSetting  uint8_t
	FbNumEndpoints      uint8_t
	FbInterfaceClass    uint8_t
	FbInterfaceSubClass uint8_t
	FbInterfaceProtocol uint8_t
	FiInterface         uint8_t
	Fendpoint           uintptr
	Fextra              uintptr
	Fextra_length       int32
}

type libusb_interface = struct {
	Faltsetting     uintptr
	Fnum_altsetting int32
}

type libusb_config_descriptor = struct {
	FbLength             uint8_t
	FbDescriptorType     uint8_t
	FwTotalLength        uint16_t
	FbNumInterfaces      uint8_t
	FbConfigurationValue uint8_t
	FiConfiguration      uint8_t
	FbmAttributes        uint8_t
	FMaxPower            uint8_t
	Finterface1          uintptr
	Fextra               uintptr
	Fextra_length        int32
}

type libusb_ss_endpoint_companion_descriptor = struct {
	FbLength           uint8_t
	FbDescriptorType   uint8_t
	FbMaxBurst         uint8_t
	FbmAttributes      uint8_t
	FwBytesPerInterval uint16_t
}

type libusb_bos_dev_capability_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbDevCapabilityType uint8_t
}

type libusb_bos_descriptor = struct {
	F__ccgo_align    [0]uint64
	FbLength         uint8_t
	FbDescriptorType uint8_t
	FwTotalLength    uint16_t
	FbNumDeviceCaps  uint8_t
}

type libusb_usb_2_0_extension_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbDevCapabilityType uint8_t
	FbmAttributes       uint32_t
}

type libusb_ss_usb_device_capability_descriptor = struct {
	FbLength               uint8_t
	FbDescriptorType       uint8_t
	FbDevCapabilityType    uint8_t
	FbmAttributes          uint8_t
	FwSpeedSupported       uint16_t
	FbFunctionalitySupport uint8_t
	FbU1DevExitLat         uint8_t
	FbU2DevExitLat         uint16_t
}

type libusb_superspeedplus_sublink_attribute_sublink_type = int32

const LIBUSB_SSPLUS_ATTR_TYPE_SYM = 0
const LIBUSB_SSPLUS_ATTR_TYPE_ASYM = 1

type libusb_superspeedplus_sublink_attribute_sublink_direction = int32

const LIBUSB_SSPLUS_ATTR_DIR_RX = 0
const LIBUSB_SSPLUS_ATTR_DIR_TX = 1

type libusb_superspeedplus_sublink_attribute_exponent = int32

const LIBUSB_SSPLUS_ATTR_EXP_BPS = 0
const LIBUSB_SSPLUS_ATTR_EXP_KBS = 1
const LIBUSB_SSPLUS_ATTR_EXP_MBS = 2
const LIBUSB_SSPLUS_ATTR_EXP_GBS = 3

type libusb_superspeedplus_sublink_attribute_link_protocol = int32

const LIBUSB_SSPLUS_ATTR_PROT_SS = 0
const LIBUSB_SSPLUS_ATTR_PROT_SSPLUS = 1

type libusb_ssplus_sublink_attribute = struct {
	Fssid      uint8_t
	Fexponent  libusb_superspeedplus_sublink_attribute_exponent
	Ftype1     libusb_superspeedplus_sublink_attribute_sublink_type
	Fdirection libusb_superspeedplus_sublink_attribute_sublink_direction
	Fprotocol  libusb_superspeedplus_sublink_attribute_link_protocol
	Fmantissa  uint16_t
}

type libusb_ssplus_usb_device_capability_descriptor = struct {
	F__ccgo_align              [0]uint32
	FnumSublinkSpeedAttributes uint8_t
	FnumSublinkSpeedIDs        uint8_t
	Fssid                      uint8_t
	FminRxLaneCount            uint8_t
	FminTxLaneCount            uint8_t
}

type libusb_container_id_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbDevCapabilityType uint8_t
	FbReserved          uint8_t
	FContainerID        [16]uint8_t
}

type libusb_platform_descriptor = struct {
	FbLength                uint8_t
	FbDescriptorType        uint8_t
	FbDevCapabilityType     uint8_t
	FbReserved              uint8_t
	FPlatformCapabilityUUID [16]uint8_t
}

type libusb_control_setup = struct {
	FbmRequestType uint8_t
	FbRequest      uint8_t
	FwValue        uint16_t
	FwIndex        uint16_t
	FwLength       uint16_t
}

type libusb_context = struct {
	Fdebug                  libusb_log_level
	Fdebug_fixed            int32
	Flog_handler            libusb_log_cb
	Fevent                  usbi_event_t
	Fusb_devs_lock          usbi_mutex_t
	Fusb_devs               list_head
	Fopen_devs_lock         usbi_mutex_t
	Fopen_devs              list_head
	Fhotplug_cbs_lock       usbi_mutex_t
	Fhotplug_cbs            list_head
	Fnext_hotplug_cb_handle libusb_hotplug_callback_handle
	Fhotplug_ready          usbi_atomic_t
	Fflying_transfers_lock  usbi_mutex_t
	Fflying_transfers       list_head
	Ffd_added_cb            libusb_pollfd_added_cb
	Ffd_removed_cb          libusb_pollfd_removed_cb
	Ffd_cb_user_data        uintptr
	Fevents_lock            usbi_mutex_t
	Fevent_handler_active   int32
	Fevent_handling_key     usbi_tls_key_t
	Fevent_waiters_lock     usbi_mutex_t
	Fevent_waiters_cond     usbi_cond_t
	Fevent_data_lock        usbi_mutex_t
	Fevent_flags            uint32
	Fdevice_close           uint32
	Fevent_sources          list_head
	Fremoved_event_sources  list_head
	Fevent_data             uintptr
	Fevent_data_cnt         uint32
	Fhotplug_msgs           list_head
	Fcompleted_transfers    list_head
	Flist                   list_head
}

type libusb_log_level = int32

const LIBUSB_LOG_LEVEL_NONE = 0
const LIBUSB_LOG_LEVEL_ERROR = 1
const LIBUSB_LOG_LEVEL_WARNING = 2
const LIBUSB_LOG_LEVEL_INFO = 3
const LIBUSB_LOG_LEVEL_DEBUG = 4

type libusb_device = struct {
	Frefcnt              usbi_atomic_t
	Fctx                 uintptr
	Fparent_dev          uintptr
	Fbus_number          uint8_t
	Fport_number         uint8_t
	Fdevice_address      uint8_t
	Fspeed               libusb_speed
	Flist                list_head
	Fsession_data        uint64
	Fdevice_descriptor   libusb_device_descriptor
	Fattached            usbi_atomic_t
	Fdevice_strings_utf8 [3]uintptr
}

type libusb_speed = int32

const LIBUSB_SPEED_UNKNOWN = 0
const LIBUSB_SPEED_LOW = 1
const LIBUSB_SPEED_FULL = 2
const LIBUSB_SPEED_HIGH = 3
const LIBUSB_SPEED_SUPER = 4
const LIBUSB_SPEED_SUPER_PLUS = 5
const LIBUSB_SPEED_SUPER_PLUS_X2 = 6

type libusb_device_handle = struct {
	Flock                      usbi_mutex_t
	Fclaimed_interfaces        uint64
	Flist                      list_head
	Fdev                       uintptr
	Fauto_detach_kernel_driver int32
}

type libusb_version = struct {
	Fmajor    uint16_t
	Fminor    uint16_t
	Fmicro    uint16_t
	Fnano     uint16_t
	Frc       uintptr
	Fdescribe uintptr
}

type libusb_error = int32

const LIBUSB_SUCCESS = 0
const LIBUSB_ERROR_IO = -1
const LIBUSB_ERROR_INVALID_PARAM = -2
const LIBUSB_ERROR_ACCESS = -3
const LIBUSB_ERROR_NO_DEVICE = -4
const LIBUSB_ERROR_NOT_FOUND = -5
const LIBUSB_ERROR_BUSY = -6
const LIBUSB_ERROR_TIMEOUT = -7
const LIBUSB_ERROR_OVERFLOW = -8
const LIBUSB_ERROR_PIPE = -9
const LIBUSB_ERROR_INTERRUPTED = -10
const LIBUSB_ERROR_NO_MEM = -11
const LIBUSB_ERROR_NOT_SUPPORTED = -12
const LIBUSB_ERROR_OTHER = -99

type libusb_transfer_type = int32

const LIBUSB_TRANSFER_TYPE_CONTROL = 0
const LIBUSB_TRANSFER_TYPE_ISOCHRONOUS = 1
const LIBUSB_TRANSFER_TYPE_BULK = 2
const LIBUSB_TRANSFER_TYPE_INTERRUPT = 3
const LIBUSB_TRANSFER_TYPE_BULK_STREAM = 4

type libusb_transfer_status = int32

const LIBUSB_TRANSFER_COMPLETED = 0
const LIBUSB_TRANSFER_ERROR = 1
const LIBUSB_TRANSFER_TIMED_OUT = 2
const LIBUSB_TRANSFER_CANCELLED = 3
const LIBUSB_TRANSFER_STALL = 4
const LIBUSB_TRANSFER_NO_DEVICE = 5
const LIBUSB_TRANSFER_OVERFLOW = 6

type libusb_transfer_flags = int32

const LIBUSB_TRANSFER_SHORT_NOT_OK = 1
const LIBUSB_TRANSFER_FREE_BUFFER = 2
const LIBUSB_TRANSFER_FREE_TRANSFER = 4
const LIBUSB_TRANSFER_ADD_ZERO_PACKET = 8

type libusb_iso_packet_descriptor = struct {
	Flength        uint32
	Factual_length uint32
	Fstatus        libusb_transfer_status
}

type libusb_transfer = struct {
	Fdev_handle      uintptr
	Fflags           uint8_t
	Fendpoint        uint8
	Ftype1           uint8
	Ftimeout         uint32
	Fstatus          libusb_transfer_status
	Flength          int32
	Factual_length   int32
	Fcallback        libusb_transfer_cb_fn
	Fuser_data       uintptr
	Fbuffer          uintptr
	Fnum_iso_packets int32
}

type libusb_transfer_cb_fn = uintptr

type libusb_capability = int32

const LIBUSB_CAP_HAS_CAPABILITY = 0
const LIBUSB_CAP_HAS_HOTPLUG = 1
const LIBUSB_CAP_HAS_HID_ACCESS = 256
const LIBUSB_CAP_SUPPORTS_DETACH_KERNEL_DRIVER = 257

type libusb_log_cb_mode = int32

const LIBUSB_LOG_CB_GLOBAL = 1
const LIBUSB_LOG_CB_CONTEXT = 2

type libusb_option = int32

const LIBUSB_OPTION_LOG_LEVEL = 0
const LIBUSB_OPTION_USE_USBDK = 1
const LIBUSB_OPTION_NO_DEVICE_DISCOVERY = 2
const LIBUSB_OPTION_LOG_CB = 3
const LIBUSB_OPTION_MAX = 4

type libusb_device_string_type = int32

const LIBUSB_DEVICE_STRING_MANUFACTURER = 0
const LIBUSB_DEVICE_STRING_PRODUCT = 1
const LIBUSB_DEVICE_STRING_SERIAL_NUMBER = 2
const LIBUSB_DEVICE_STRING_COUNT = 3

type libusb_log_cb = uintptr

type libusb_init_option = struct {
	Foption libusb_option
	Fvalue  struct {
		Flog_cbval   [0]libusb_log_cb
		Fival        int32
		F__ccgo_pad2 [4]byte
	}
}

func libusb_control_transfer_get_data(tls *libc.TLS, transfer uintptr) (r uintptr) {
	return (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer + uintptr(libc.Uint64FromInt64(8))
}

func libusb_control_transfer_get_setup(tls *libc.TLS, transfer uintptr) (r uintptr) {
	return (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer
}

func libusb_fill_control_setup(tls *libc.TLS, buffer uintptr, bmRequestType uint8_t, bRequest uint8_t, wValue uint16_t, wIndex uint16_t, wLength uint16_t) {
	var setup uintptr
	_ = setup
	setup = buffer
	(*libusb_control_setup)(unsafe.Pointer(setup)).FbmRequestType = bmRequestType
	(*libusb_control_setup)(unsafe.Pointer(setup)).FbRequest = bRequest
	(*libusb_control_setup)(unsafe.Pointer(setup)).FwValue = libusb_cpu_to_le16(tls, wValue)
	(*libusb_control_setup)(unsafe.Pointer(setup)).FwIndex = libusb_cpu_to_le16(tls, wIndex)
	(*libusb_control_setup)(unsafe.Pointer(setup)).FwLength = libusb_cpu_to_le16(tls, wLength)
}

func libusb_fill_control_transfer(tls *libc.TLS, transfer uintptr, dev_handle uintptr, buffer uintptr, __ccgo_fp_callback libusb_transfer_cb_fn, user_data uintptr, timeout uint32) {
	var setup uintptr
	_ = setup
	setup = buffer
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle = dev_handle
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint = uint8(0)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1 = uint8(LIBUSB_TRANSFER_TYPE_CONTROL)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftimeout = timeout
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer = buffer
	if setup != 0 {
		(*libusb_transfer)(unsafe.Pointer(transfer)).Flength = libc.Int32FromUint64(libc.Uint64FromInt64(8) + uint64(libusb_cpu_to_le16(tls, (*libusb_control_setup)(unsafe.Pointer(setup)).FwLength)))
	}
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data = user_data
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback = __ccgo_fp_callback
}

func libusb_fill_bulk_transfer(tls *libc.TLS, transfer uintptr, dev_handle uintptr, endpoint uint8, buffer uintptr, length int32, __ccgo_fp_callback libusb_transfer_cb_fn, user_data uintptr, timeout uint32) {
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle = dev_handle
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint = endpoint
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1 = uint8(LIBUSB_TRANSFER_TYPE_BULK)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftimeout = timeout
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer = buffer
	(*libusb_transfer)(unsafe.Pointer(transfer)).Flength = length
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data = user_data
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback = __ccgo_fp_callback
}

func libusb_fill_bulk_stream_transfer(tls *libc.TLS, transfer uintptr, dev_handle uintptr, endpoint uint8, stream_id uint32_t, buffer uintptr, length int32, __ccgo_fp_callback libusb_transfer_cb_fn, user_data uintptr, timeout uint32) {
	libusb_fill_bulk_transfer(tls, transfer, dev_handle, endpoint, buffer, length, __ccgo_fp_callback, user_data, timeout)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1 = uint8(LIBUSB_TRANSFER_TYPE_BULK_STREAM)
	libusb_transfer_set_stream_id(tls, transfer, stream_id)
}

func libusb_fill_interrupt_transfer(tls *libc.TLS, transfer uintptr, dev_handle uintptr, endpoint uint8, buffer uintptr, length int32, __ccgo_fp_callback libusb_transfer_cb_fn, user_data uintptr, timeout uint32) {
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle = dev_handle
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint = endpoint
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1 = uint8(LIBUSB_TRANSFER_TYPE_INTERRUPT)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftimeout = timeout
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer = buffer
	(*libusb_transfer)(unsafe.Pointer(transfer)).Flength = length
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data = user_data
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback = __ccgo_fp_callback
}

func libusb_fill_iso_transfer(tls *libc.TLS, transfer uintptr, dev_handle uintptr, endpoint uint8, buffer uintptr, length int32, num_iso_packets int32, __ccgo_fp_callback libusb_transfer_cb_fn, user_data uintptr, timeout uint32) {
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle = dev_handle
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint = endpoint
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1 = uint8(LIBUSB_TRANSFER_TYPE_ISOCHRONOUS)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftimeout = timeout
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer = buffer
	(*libusb_transfer)(unsafe.Pointer(transfer)).Flength = length
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fnum_iso_packets = num_iso_packets
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data = user_data
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback = __ccgo_fp_callback
}

func libusb_set_iso_packet_lengths(tls *libc.TLS, transfer uintptr, length uint32) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < (*libusb_transfer)(unsafe.Pointer(transfer)).Fnum_iso_packets) {
			break
		}
		(*(*libusb_iso_packet_descriptor)(unsafe.Pointer(transfer + 60 + uintptr(i)*12))).Flength = length
		goto _1
	_1:
		;
		i++
	}
}

func libusb_get_iso_packet_buffer(tls *libc.TLS, transfer uintptr, packet uint32) (r uintptr) {
	var _packet, i int32
	var offset size_t
	_, _, _ = _packet, i, offset
	offset = uint64(0)
	if packet > uint32(0x7fffffff) {
		return libc.UintptrFromInt32(0)
	}
	_packet = libc.Int32FromUint32(packet)
	if _packet >= (*libusb_transfer)(unsafe.Pointer(transfer)).Fnum_iso_packets {
		return libc.UintptrFromInt32(0)
	}
	i = 0
	for {
		if !(i < _packet) {
			break
		}
		offset += uint64((*(*libusb_iso_packet_descriptor)(unsafe.Pointer(transfer + 60 + uintptr(i)*12))).Flength)
		goto _1
	_1:
		;
		i++
	}
	return (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer + uintptr(offset)
}

func libusb_get_iso_packet_buffer_simple(tls *libc.TLS, transfer uintptr, packet uint32) (r uintptr) {
	var _packet int32
	_ = _packet
	if packet > uint32(0x7fffffff) {
		return libc.UintptrFromInt32(0)
	}
	_packet = libc.Int32FromUint32(packet)
	if _packet >= (*libusb_transfer)(unsafe.Pointer(transfer)).Fnum_iso_packets {
		return libc.UintptrFromInt32(0)
	}
	return (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer + uintptr(libc.Int32FromUint32((*(*libusb_iso_packet_descriptor)(unsafe.Pointer(transfer + 60))).Flength)*_packet)
}

func libusb_get_descriptor(tls *libc.TLS, dev_handle uintptr, desc_type uint8_t, desc_index uint8_t, data uintptr, length int32) (r int32) {
	return libusb_control_transfer(tls, dev_handle, uint8(LIBUSB_ENDPOINT_IN), uint8(LIBUSB_REQUEST_GET_DESCRIPTOR), libc.Uint16FromInt32(libc.Int32FromUint8(desc_type)<<libc.Int32FromInt32(8)|libc.Int32FromUint8(desc_index)), uint16(0), data, libc.Uint16FromInt32(length), uint32(1000))
}

func libusb_get_string_descriptor(tls *libc.TLS, dev_handle uintptr, desc_index uint8_t, langid uint16_t, data uintptr, length int32) (r int32) {
	return libusb_control_transfer(tls, dev_handle, uint8(LIBUSB_ENDPOINT_IN), uint8(LIBUSB_REQUEST_GET_DESCRIPTOR), libc.Uint16FromInt32(int32(LIBUSB_DT_STRING)<<libc.Int32FromInt32(8)|libc.Int32FromUint8(desc_index)), langid, data, libc.Uint16FromInt32(length), uint32(1000))
}

type libusb_pollfd = struct {
	Ffd     int32
	Fevents int16
}

type libusb_pollfd_added_cb = uintptr

type libusb_pollfd_removed_cb = uintptr

type libusb_hotplug_callback_handle = int32

type libusb_hotplug_event = int32

const LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED = 1
const LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT = 2

type libusb_hotplug_flag = int32

const LIBUSB_HOTPLUG_ENUMERATE = 1

type libusb_hotplug_callback_fn = uintptr

type memory_order = int32

const memory_order_relaxed = 0
const memory_order_consume = 1
const memory_order_acquire = 2
const memory_order_release = 3
const memory_order_acq_rel = 4
const memory_order_seq_cst = 5

type atomic_bool = uint8

type atomic_char = uint8

type atomic_schar = int8

type atomic_uchar = uint8

type atomic_short = int16

type atomic_ushort = uint16

type atomic_int = int32

type atomic_uint = uint32

type atomic_long = int64

type atomic_ulong = uint64

type atomic_llong = int64

type atomic_ullong = uint64

type atomic_char8_t = uint8

type atomic_char16_t = uint16

type atomic_char32_t = uint32

type atomic_wchar_t = int32

type atomic_int_least8_t = int8

type atomic_uint_least8_t = uint8

type atomic_int_least16_t = int16

type atomic_uint_least16_t = uint16

type atomic_int_least32_t = int32

type atomic_uint_least32_t = uint32

type atomic_int_least64_t = int64

type atomic_uint_least64_t = uint64

type atomic_int_fast8_t = int8

type atomic_uint_fast8_t = uint8

type atomic_int_fast16_t = int32

type atomic_uint_fast16_t = uint32

type atomic_int_fast32_t = int32

type atomic_uint_fast32_t = uint32

type atomic_int_fast64_t = int64

type atomic_uint_fast64_t = uint64

type atomic_intptr_t = int64

type atomic_uintptr_t = uint64

type atomic_size_t = uint64

type atomic_ptrdiff_t = int64

type atomic_intmax_t = int64

type atomic_uintmax_t = uint64

type atomic_flag = struct {
	F__val uint8
}

type usbi_atomic_t = int64

type usbi_os_handle_t = int32

type usbi_event_t = struct {
	Fpipefd [2]int32
}

type usbi_event = usbi_event_t

type sched_param = struct {
	Fsched_priority int32
	F__reserved1    int32
	F__reserved2    [2]struct {
		F__reserved1 time_t
		F__reserved2 int64
	}
	F__reserved3 int32
}

type cpu_set_t = struct {
	F__bits [16]uint64
}

func __CPU_AND_S(tls *libc.TLS, __size size_t, __dest uintptr, __src1 uintptr, __src2 uintptr) {
	var __i size_t
	_ = __i
	__i = uint64(0)
	for {
		if !(__i < __size/uint64(8)) {
			break
		}
		*(*uint64)(unsafe.Pointer(__dest + uintptr(__i)*8)) = *(*uint64)(unsafe.Pointer(__src1 + uintptr(__i)*8)) & *(*uint64)(unsafe.Pointer(__src2 + uintptr(__i)*8))
		goto _1
	_1:
		;
		__i++
	}
}

func __CPU_OR_S(tls *libc.TLS, __size size_t, __dest uintptr, __src1 uintptr, __src2 uintptr) {
	var __i size_t
	_ = __i
	__i = uint64(0)
	for {
		if !(__i < __size/uint64(8)) {
			break
		}
		*(*uint64)(unsafe.Pointer(__dest + uintptr(__i)*8)) = *(*uint64)(unsafe.Pointer(__src1 + uintptr(__i)*8)) | *(*uint64)(unsafe.Pointer(__src2 + uintptr(__i)*8))
		goto _1
	_1:
		;
		__i++
	}
}

func __CPU_XOR_S(tls *libc.TLS, __size size_t, __dest uintptr, __src1 uintptr, __src2 uintptr) {
	var __i size_t
	_ = __i
	__i = uint64(0)
	for {
		if !(__i < __size/uint64(8)) {
			break
		}
		*(*uint64)(unsafe.Pointer(__dest + uintptr(__i)*8)) = *(*uint64)(unsafe.Pointer(__src1 + uintptr(__i)*8)) ^ *(*uint64)(unsafe.Pointer(__src2 + uintptr(__i)*8))
		goto _1
	_1:
		;
		__i++
	}
}

type __ptcb = struct {
	F__f    uintptr
	F__x    uintptr
	F__next uintptr
}

type usbi_mutex_static_t = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

func usbi_mutex_static_lock(tls *libc.TLS, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_mutex_lock(tls, mutex) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+12, __ccgo_ts+43, int32(32), uintptr(unsafe.Pointer(&__func__)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__ = [23]uint8{'u', 's', 'b', 'i', '_', 'm', 'u', 't', 'e', 'x', '_', 's', 't', 'a', 't', 'i', 'c', '_', 'l', 'o', 'c', 'k'}

func usbi_mutex_static_unlock(tls *libc.TLS, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_mutex_unlock(tls, mutex) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+96, __ccgo_ts+43, int32(36), uintptr(unsafe.Pointer(&__func__1)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__1 = [25]uint8{'u', 's', 'b', 'i', '_', 'm', 'u', 't', 'e', 'x', '_', 's', 't', 'a', 't', 'i', 'c', '_', 'u', 'n', 'l', 'o', 'c', 'k'}

type usbi_mutex_t = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

func usbi_mutex_init(tls *libc.TLS, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_mutex_init(tls, mutex, libc.UintptrFromInt32(0)) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+129, __ccgo_ts+43, int32(42), uintptr(unsafe.Pointer(&__func__2)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__2 = [16]uint8{'u', 's', 'b', 'i', '_', 'm', 'u', 't', 'e', 'x', '_', 'i', 'n', 'i', 't'}

func usbi_mutex_lock(tls *libc.TLS, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_mutex_lock(tls, mutex) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+12, __ccgo_ts+43, int32(46), uintptr(unsafe.Pointer(&__func__3)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__3 = [16]uint8{'u', 's', 'b', 'i', '_', 'm', 'u', 't', 'e', 'x', '_', 'l', 'o', 'c', 'k'}

func usbi_mutex_unlock(tls *libc.TLS, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_mutex_unlock(tls, mutex) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+96, __ccgo_ts+43, int32(50), uintptr(unsafe.Pointer(&__func__4)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__4 = [18]uint8{'u', 's', 'b', 'i', '_', 'm', 'u', 't', 'e', 'x', '_', 'u', 'n', 'l', 'o', 'c', 'k'}

func usbi_mutex_trylock(tls *libc.TLS, mutex uintptr) (r int32) {
	var mutexIsLocked int32
	_ = mutexIsLocked
	mutexIsLocked = libc.BoolInt32(libc.Xpthread_mutex_trylock(tls, mutex) == 0)
	return mutexIsLocked
}

func usbi_mutex_destroy(tls *libc.TLS, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_mutex_destroy(tls, mutex) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+172, __ccgo_ts+43, int32(59), uintptr(unsafe.Pointer(&__func__5)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__5 = [19]uint8{'u', 's', 'b', 'i', '_', 'm', 'u', 't', 'e', 'x', '_', 'd', 'e', 's', 't', 'r', 'o', 'y'}

type usbi_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][6]uintptr
		F__i  [12]int32
	}
}

func usbi_cond_wait(tls *libc.TLS, cond uintptr, mutex uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_cond_wait(tls, cond, mutex) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+206, __ccgo_ts+43, int32(67), uintptr(unsafe.Pointer(&__func__6)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__6 = [15]uint8{'u', 's', 'b', 'i', '_', 'c', 'o', 'n', 'd', '_', 'w', 'a', 'i', 't'}

func usbi_cond_signal(tls *libc.TLS, cond uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_cond_signal(tls, cond) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+242, __ccgo_ts+43, int32(73), uintptr(unsafe.Pointer(&__func__7)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__7 = [17]uint8{'u', 's', 'b', 'i', '_', 'c', 'o', 'n', 'd', '_', 's', 'i', 'g', 'n', 'a', 'l'}

func usbi_cond_broadcast(tls *libc.TLS, cond uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_cond_broadcast(tls, cond) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+273, __ccgo_ts+43, int32(77), uintptr(unsafe.Pointer(&__func__8)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__8 = [20]uint8{'u', 's', 'b', 'i', '_', 'c', 'o', 'n', 'd', '_', 'b', 'r', 'o', 'a', 'd', 'c', 'a', 's', 't'}

func usbi_cond_destroy(tls *libc.TLS, cond uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_cond_destroy(tls, cond) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+307, __ccgo_ts+43, int32(81), uintptr(unsafe.Pointer(&__func__9)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__9 = [18]uint8{'u', 's', 'b', 'i', '_', 'c', 'o', 'n', 'd', '_', 'd', 'e', 's', 't', 'r', 'o', 'y'}

type usbi_tls_key_t = uint32

func usbi_tls_key_create(tls *libc.TLS, key uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_key_create(tls, key, libc.UintptrFromInt32(0)) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+339, __ccgo_ts+43, int32(87), uintptr(unsafe.Pointer(&__func__10)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__10 = [20]uint8{'u', 's', 'b', 'i', '_', 't', 'l', 's', '_', 'k', 'e', 'y', '_', 'c', 'r', 'e', 'a', 't', 'e'}

func usbi_tls_key_get(tls *libc.TLS, key usbi_tls_key_t) (r uintptr) {
	return libc.Xpthread_getspecific(tls, key)
}

func usbi_tls_key_set(tls *libc.TLS, key usbi_tls_key_t, ptr uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_setspecific(tls, key, ptr) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+380, __ccgo_ts+43, int32(95), uintptr(unsafe.Pointer(&__func__11)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__11 = [17]uint8{'u', 's', 'b', 'i', '_', 't', 'l', 's', '_', 'k', 'e', 'y', '_', 's', 'e', 't'}

func usbi_tls_key_delete(tls *libc.TLS, key usbi_tls_key_t) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_key_delete(tls, key) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+415, __ccgo_ts+43, int32(99), uintptr(unsafe.Pointer(&__func__12)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__12 = [20]uint8{'u', 's', 'b', 'i', '_', 't', 'l', 's', '_', 'k', 'e', 'y', '_', 'd', 'e', 'l', 'e', 't', 'e'}

type list_head = struct {
	Fprev uintptr
	Fnext uintptr
}

func list_init(tls *libc.TLS, entry uintptr) {
	var v1 uintptr
	_ = v1
	v1 = entry
	(*list_head)(unsafe.Pointer(entry)).Fnext = v1
	(*list_head)(unsafe.Pointer(entry)).Fprev = v1
}

func list_add(tls *libc.TLS, entry uintptr, head uintptr) {
	(*list_head)(unsafe.Pointer(entry)).Fnext = (*list_head)(unsafe.Pointer(head)).Fnext
	(*list_head)(unsafe.Pointer(entry)).Fprev = head
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(head)).Fnext)).Fprev = entry
	(*list_head)(unsafe.Pointer(head)).Fnext = entry
}

func list_add_tail(tls *libc.TLS, entry uintptr, head uintptr) {
	(*list_head)(unsafe.Pointer(entry)).Fnext = head
	(*list_head)(unsafe.Pointer(entry)).Fprev = (*list_head)(unsafe.Pointer(head)).Fprev
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(head)).Fprev)).Fnext = entry
	(*list_head)(unsafe.Pointer(head)).Fprev = entry
}

func list_del(tls *libc.TLS, entry uintptr) {
	var v1 uintptr
	_ = v1
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(entry)).Fnext)).Fprev = (*list_head)(unsafe.Pointer(entry)).Fprev
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(entry)).Fprev)).Fnext = (*list_head)(unsafe.Pointer(entry)).Fnext
	v1 = libc.UintptrFromInt32(0)
	(*list_head)(unsafe.Pointer(entry)).Fprev = v1
	(*list_head)(unsafe.Pointer(entry)).Fnext = v1
}

func list_cut(tls *libc.TLS, list uintptr, head uintptr) {
	if (*list_head)(unsafe.Pointer(head)).Fnext == head {
		list_init(tls, list)
		return
	}
	(*list_head)(unsafe.Pointer(list)).Fnext = (*list_head)(unsafe.Pointer(head)).Fnext
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(list)).Fnext)).Fprev = list
	(*list_head)(unsafe.Pointer(list)).Fprev = (*list_head)(unsafe.Pointer(head)).Fprev
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(list)).Fprev)).Fnext = list
	list_init(tls, head)
}

func list_splice_front(tls *libc.TLS, list uintptr, head uintptr) {
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(list)).Fnext)).Fprev = head
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(list)).Fprev)).Fnext = (*list_head)(unsafe.Pointer(head)).Fnext
	(*list_head)(unsafe.Pointer((*list_head)(unsafe.Pointer(head)).Fnext)).Fprev = (*list_head)(unsafe.Pointer(list)).Fprev
	(*list_head)(unsafe.Pointer(head)).Fnext = (*list_head)(unsafe.Pointer(list)).Fnext
}

func usbi_reallocf(tls *libc.TLS, ptr uintptr, size size_t) (r uintptr) {
	var ret uintptr
	_ = ret
	ret = libc.Xrealloc(tls, ptr, size)
	if !(ret != 0) {
		libc.Xfree(tls, ptr)
	}
	return ret
}

func usbi_get_context(tls *libc.TLS, ctx uintptr) (r uintptr) {
	if !(ctx != 0) {
		ctx = usbi_default_context
	}
	if !(ctx != 0) {
		ctx = usbi_fallback_context
		if ctx != 0 && warned == 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__13)), __ccgo_ts+444, 0)
			warned = int32(1)
		}
	}
	return ctx
}

var __func__13 = [17]uint8{'u', 's', 'b', 'i', '_', 'g', 'e', 't', '_', 'c', 'o', 'n', 't', 'e', 'x', 't'}

var warned int32

type usbi_event_flags = int32

const USBI_EVENT_EVENT_SOURCES_MODIFIED = 1
const USBI_EVENT_USER_INTERRUPT = 2
const USBI_EVENT_HOTPLUG_CB_DEREGISTERED = 4
const USBI_EVENT_HOTPLUG_MSG_PENDING = 8
const USBI_EVENT_TRANSFER_COMPLETED = 16
const USBI_EVENT_DEVICE_CLOSE = 32

func usbi_handling_events(tls *libc.TLS, ctx uintptr) (r int32) {
	return libc.BoolInt32(usbi_tls_key_get(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_handling_key) != libc.UintptrFromInt32(0))
}

func usbi_start_event_handling(tls *libc.TLS, ctx uintptr) {
	usbi_tls_key_set(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_handling_key, ctx)
}

func usbi_end_event_handling(tls *libc.TLS, ctx uintptr) {
	usbi_tls_key_set(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_handling_key, libc.UintptrFromInt32(0))
}

func usbi_localize_device_descriptor(tls *libc.TLS, desc uintptr) {
	(*libusb_device_descriptor)(unsafe.Pointer(desc)).FbcdUSB = libusb_cpu_to_le16(tls, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FbcdUSB)
	(*libusb_device_descriptor)(unsafe.Pointer(desc)).FidVendor = libusb_cpu_to_le16(tls, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FidVendor)
	(*libusb_device_descriptor)(unsafe.Pointer(desc)).FidProduct = libusb_cpu_to_le16(tls, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FidProduct)
	(*libusb_device_descriptor)(unsafe.Pointer(desc)).FbcdDevice = libusb_cpu_to_le16(tls, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FbcdDevice)
}

func usbi_get_monotonic_time(tls *libc.TLS, tp uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xclock_gettime(tls, int32(1), tp) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+503, __ccgo_ts+529, int32(554), uintptr(unsafe.Pointer(&__func__14)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__14 = [24]uint8{'u', 's', 'b', 'i', '_', 'g', 'e', 't', '_', 'm', 'o', 'n', 'o', 't', 'o', 'n', 'i', 'c', '_', 't', 'i', 'm', 'e'}

func usbi_get_real_time(tls *libc.TLS, tp uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xclock_gettime(tls, 0, tp) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+573, __ccgo_ts+529, int32(558), uintptr(unsafe.Pointer(&__func__15)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__15 = [19]uint8{'u', 's', 'b', 'i', '_', 'g', 'e', 't', '_', 'r', 'e', 'a', 'l', '_', 't', 'i', 'm', 'e'}

type usbi_transfer = struct {
	Flock            usbi_mutex_t
	Fnum_iso_packets int32
	Flist            list_head
	Fcompleted_list  list_head
	Ftimeout         timespec
	Ftransferred     int32
	Fstream_id       uint32_t
	Fstate_flags     uint32_t
	Ftimeout_flags   uint32_t
	Fdev             uintptr
	Fpriv            uintptr
}

type usbi_transfer_state_flags = int32

const USBI_TRANSFER_IN_FLIGHT = 1
const USBI_TRANSFER_CANCELLING = 2
const USBI_TRANSFER_DEVICE_DISAPPEARED = 4

type usbi_transfer_timeout_flags = int32

const USBI_TRANSFER_OS_HANDLES_TIMEOUT = 1
const USBI_TRANSFER_TIMEOUT_HANDLED = 2
const USBI_TRANSFER_TIMED_OUT = 4

type usbi_descriptor_header = struct {
	FbLength         uint8_t
	FbDescriptorType uint8_t
}

type usbi_device_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbcdUSB             uint16_t
	FbDeviceClass       uint8_t
	FbDeviceSubClass    uint8_t
	FbDeviceProtocol    uint8_t
	FbMaxPacketSize0    uint8_t
	FidVendor           uint16_t
	FidProduct          uint16_t
	FbcdDevice          uint16_t
	FiManufacturer      uint8_t
	FiProduct           uint8_t
	FiSerialNumber      uint8_t
	FbNumConfigurations uint8_t
}

type usbi_configuration_descriptor = struct {
	FbLength             uint8_t
	FbDescriptorType     uint8_t
	FwTotalLength        uint16_t
	FbNumInterfaces      uint8_t
	FbConfigurationValue uint8_t
	FiConfiguration      uint8_t
	FbmAttributes        uint8_t
	FbMaxPower           uint8_t
}

type usbi_interface_descriptor = struct {
	FbLength            uint8_t
	FbDescriptorType    uint8_t
	FbInterfaceNumber   uint8_t
	FbAlternateSetting  uint8_t
	FbNumEndpoints      uint8_t
	FbInterfaceClass    uint8_t
	FbInterfaceSubClass uint8_t
	FbInterfaceProtocol uint8_t
	FiInterface         uint8_t
}

type usbi_string_descriptor = struct {
	F__ccgo_align    [0]uint16
	FbLength         uint8_t
	FbDescriptorType uint8_t
}

type usbi_bos_descriptor = struct {
	FbLength         uint8_t
	FbDescriptorType uint8_t
	FwTotalLength    uint16_t
	FbNumDeviceCaps  uint8_t
}

type usbi_config_desc_buf = struct {
	Fbuf   [0][9]uint8_t
	Falign [0]uint16_t
	Fdesc  usbi_configuration_descriptor
}

type usbi_string_desc_buf = struct {
	Fbuf         [0][255]uint8_t
	Falign       [0]uint16_t
	Fdesc        usbi_string_descriptor
	F__ccgo_pad3 [254]byte
}

type usbi_bos_desc_buf = struct {
	Fbuf   [0][5]uint8_t
	Falign [0]uint16_t
	Fdesc  usbi_bos_descriptor
}

type usbi_hotplug_flags = int32

const USBI_HOTPLUG_DEVICE_ARRIVED = 1
const USBI_HOTPLUG_DEVICE_LEFT = 2
const USBI_HOTPLUG_VENDOR_ID_VALID = 8
const USBI_HOTPLUG_PRODUCT_ID_VALID = 16
const USBI_HOTPLUG_DEV_CLASS_VALID = 32
const USBI_HOTPLUG_NEEDS_FREE = 64

type usbi_hotplug_callback = struct {
	Fflags      uint8_t
	Fvendor_id  uint16_t
	Fproduct_id uint16_t
	Fdev_class  uint8_t
	Fcb         libusb_hotplug_callback_fn
	Fhandle     libusb_hotplug_callback_handle
	Fuser_data  uintptr
	Flist       list_head
}

type usbi_hotplug_message = struct {
	Fevent  libusb_hotplug_event
	Fdevice uintptr
	Flist   list_head
}

type usbi_event_source_data = struct {
	Fos_handle   usbi_os_handle_t
	Fpoll_events int16
}

type usbi_event_source = struct {
	Fdata usbi_event_source_data
	Flist list_head
}

type usbi_option = struct {
	Fis_set int32
	Farg    struct {
		Flog_cbval   [0]libusb_log_cb
		Fival        int32
		F__ccgo_pad2 [4]byte
	}
}

func usbi_using_timer(tls *libc.TLS, ctx uintptr) (r int32) {
	_ = ctx
	return 0
}

type usbi_reported_events = struct {
	F__ccgo0_0 struct {
		Fevent_bits [0]uint32
		F__ccgo0_0  struct {
			F__ccgo_align [0]uint32
			F__ccgo0      uint8
		}
	}
	Fevent_data       uintptr
	Fevent_data_count uint32
	Fnum_ready        uint32
}

func usbi_get_context_priv(tls *libc.TLS, ctx uintptr) (r uintptr) {
	return ctx + uintptr((libc.Uint64FromInt64(568)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
}

func usbi_get_device_priv(tls *libc.TLS, dev uintptr) (r uintptr) {
	return dev + uintptr((libc.Uint64FromInt64(112)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
}

func usbi_get_device_handle_priv(tls *libc.TLS, dev_handle uintptr) (r uintptr) {
	return dev_handle + uintptr((libc.Uint64FromInt64(80)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
}

func usbi_get_transfer_priv(tls *libc.TLS, itransfer uintptr) (r uintptr) {
	return (*usbi_transfer)(unsafe.Pointer(itransfer)).Fpriv
}

type discovered_devs = struct {
	Flen1     size_t
	Fcapacity size_t
}

type usbi_os_backend = struct {
	Fname                           uintptr
	Fcaps                           uint32_t
	Finit1                          uintptr
	Fexit                           uintptr
	Fset_option                     uintptr
	Fget_device_list                uintptr
	Fget_device_string              uintptr
	Fhotplug_poll                   uintptr
	Fwrap_sys_device                uintptr
	Fopen                           uintptr
	Fclose1                         uintptr
	Fget_active_config_descriptor   uintptr
	Fget_config_descriptor          uintptr
	Fget_config_descriptor_by_value uintptr
	Fget_configuration              uintptr
	Fset_configuration              uintptr
	Fclaim_interface                uintptr
	Frelease_interface              uintptr
	Fset_interface_altsetting       uintptr
	Fclear_halt                     uintptr
	Freset_device                   uintptr
	Falloc_streams                  uintptr
	Ffree_streams                   uintptr
	Fdev_mem_alloc                  uintptr
	Fdev_mem_free                   uintptr
	Fkernel_driver_active           uintptr
	Fdetach_kernel_driver           uintptr
	Fattach_kernel_driver           uintptr
	Fendpoint_supports_raw_io       uintptr
	Fendpoint_set_raw_io            uintptr
	Fget_max_raw_io_transfer_size   uintptr
	Fdestroy_device                 uintptr
	Fsubmit_transfer                uintptr
	Fcancel_transfer                uintptr
	Fclear_transfer_priv            uintptr
	Fhandle_events                  uintptr
	Fhandle_transfer_completion     uintptr
	Fcontext_priv_size              size_t
	Fdevice_priv_size               size_t
	Fdevice_handle_priv_size        size_t
	Ftransfer_priv_size             size_t
}

type iovec = struct {
	Fiov_base uintptr
	Fiov_len  size_t
}

type flock1 = struct {
	Fl_type   int16
	Fl_whence int16
	Fl_start  off_t
	Fl_len    off_t
	Fl_pid    pid_t
}

type file_handle = struct {
	Fhandle_bytes uint32
	Fhandle_type  int32
}

type f_owner_ex = struct {
	Ftype1 int32
	Fpid   pid_t
}

type usbi_nfds_t = uint64

func usbi_create_event(tls *libc.TLS, event uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ret int32
	_ = ret
	ret = libc.Xpipe2(tls, event, int32(02000000))
	if ret != 0 {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__16)), __ccgo_ts+599, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	ret = libc.Xfcntl(tls, *(*int32)(unsafe.Pointer(event + 1*4)), int32(3), 0)
	if ret == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__16)), __ccgo_ts+631, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		goto err_close_pipe
	}
	ret = libc.Xfcntl(tls, *(*int32)(unsafe.Pointer(event + 1*4)), int32(4), libc.VaList(bp+8, ret|int32(04000)))
	if ret == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__16)), __ccgo_ts+676, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		goto err_close_pipe
	}
	return 0
	goto err_close_pipe
err_close_pipe:
	;
	libc.Xclose(tls, *(*int32)(unsafe.Pointer(event + 1*4)))
	libc.Xclose(tls, *(*int32)(unsafe.Pointer(event)))
	return int32(LIBUSB_ERROR_OTHER)
}

var __func__16 = [18]uint8{'u', 's', 'b', 'i', '_', 'c', 'r', 'e', 'a', 't', 'e', '_', 'e', 'v', 'e', 'n', 't'}

func usbi_destroy_event(tls *libc.TLS, event uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	if libc.Xclose(tls, *(*int32)(unsafe.Pointer(event + 1*4))) == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__17)), __ccgo_ts+721, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
	}
	if libc.Xclose(tls, *(*int32)(unsafe.Pointer(event))) == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__17)), __ccgo_ts+762, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
	}
}

var __func__17 = [19]uint8{'u', 's', 'b', 'i', '_', 'd', 'e', 's', 't', 'r', 'o', 'y', '_', 'e', 'v', 'e', 'n', 't'}

func usbi_signal_event(tls *libc.TLS, event uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r ssize_t
	var _ /* dummy at bp+0 */ uint64_t
	_ = r
	*(*uint64_t)(unsafe.Pointer(bp)) = uint64(1)
	r = libc.Xwrite(tls, *(*int32)(unsafe.Pointer(event + 1*4)), bp, uint64(8))
	if libc.Uint64FromInt64(r) != uint64(8) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__18)), __ccgo_ts+802, 0)
	}
}

var __func__18 = [18]uint8{'u', 's', 'b', 'i', '_', 's', 'i', 'g', 'n', 'a', 'l', '_', 'e', 'v', 'e', 'n', 't'}

func usbi_clear_event(tls *libc.TLS, event uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r ssize_t
	var _ /* dummy at bp+0 */ uint64_t
	_ = r
	r = libc.Xread(tls, *(*int32)(unsafe.Pointer(event)), bp, uint64(8))
	if libc.Uint64FromInt64(r) != uint64(8) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__19)), __ccgo_ts+821, 0)
	}
}

var __func__19 = [17]uint8{'u', 's', 'b', 'i', '_', 'c', 'l', 'e', 'a', 'r', '_', 'e', 'v', 'e', 'n', 't'}

func usbi_alloc_event_data(tls *libc.TLS, ctx uintptr) (r int32) {
	var fds, ievent_source uintptr
	var i size_t
	_, _, _ = fds, i, ievent_source
	i = uint64(0)
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_data != 0 {
		libc.Xfree(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_data)
		(*libusb_context)(unsafe.Pointer(ctx)).Fevent_data = libc.UintptrFromInt32(0)
	}
	(*libusb_context)(unsafe.Pointer(ctx)).Fevent_data_cnt = uint32(0)
	ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+472)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	for {
		if !(ievent_source+8 != ctx+472) {
			break
		}
		(*libusb_context)(unsafe.Pointer(ctx)).Fevent_data_cnt++
		goto _1
	_1:
		;
		ievent_source = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	}
	fds = libc.Xcalloc(tls, uint64((*libusb_context)(unsafe.Pointer(ctx)).Fevent_data_cnt), uint64(8))
	if !(fds != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+472)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	for {
		if !(ievent_source+8 != ctx+472) {
			break
		}
		(*(*pollfd)(unsafe.Pointer(fds + uintptr(i)*8))).Ffd = (*usbi_event_source)(unsafe.Pointer(ievent_source)).Fdata.Fos_handle
		(*(*pollfd)(unsafe.Pointer(fds + uintptr(i)*8))).Fevents = (*usbi_event_source)(unsafe.Pointer(ievent_source)).Fdata.Fpoll_events
		i++
		goto _2
	_2:
		;
		ievent_source = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	}
	(*libusb_context)(unsafe.Pointer(ctx)).Fevent_data = fds
	return 0
}

func usbi_wait_for_events(tls *libc.TLS, ctx uintptr, reported_events uintptr, timeout_ms int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var fds, ievent_source uintptr
	var internal_fds, num_ready, v1 int32
	var n, nfds usbi_nfds_t
	var v4 bool
	_, _, _, _, _, _, _, _ = fds, ievent_source, internal_fds, n, nfds, num_ready, v1, v4
	fds = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_data
	nfds = uint64((*libusb_context)(unsafe.Pointer(ctx)).Fevent_data_cnt)
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__20)), __ccgo_ts+839, libc.VaList(bp+8, uint32(nfds), timeout_ms))
	num_ready = libc.Xpoll(tls, fds, nfds, timeout_ms)
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__20)), __ccgo_ts+874, libc.VaList(bp+8, num_ready))
	if num_ready == 0 {
		if usbi_using_timer(tls, ctx) != 0 {
			goto done
		}
		return int32(LIBUSB_ERROR_TIMEOUT)
	} else {
		if num_ready == -int32(1) {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(4) {
				return int32(LIBUSB_ERROR_INTERRUPTED)
			}
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__20)), __ccgo_ts+893, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			return int32(LIBUSB_ERROR_IO)
		}
	}
	if (*(*pollfd)(unsafe.Pointer(fds))).Frevents != 0 {
		libc.SetBitFieldPtr8Uint32(reported_events+0, libc.Uint32FromInt32(1), 0, 0x1)
		num_ready--
	} else {
		libc.SetBitFieldPtr8Uint32(reported_events+0, libc.Uint32FromInt32(0), 0, 0x1)
	}
	if !(num_ready != 0) {
		goto done
	}
	if usbi_using_timer(tls, ctx) != 0 {
		v1 = int32(2)
	} else {
		v1 = int32(1)
	}
	internal_fds = v1
	fds += uintptr(internal_fds) * 8
	nfds -= libc.Uint64FromInt32(internal_fds)
	usbi_mutex_lock(tls, ctx+424)
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_EVENT_SOURCES_MODIFIED) != 0 {
		ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+488)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
		for {
			if !(ievent_source+8 != ctx+488) {
				break
			}
			n = uint64(0)
			for {
				if !(n < nfds) {
					break
				}
				if (*usbi_event_source)(unsafe.Pointer(ievent_source)).Fdata.Fos_handle != (*(*pollfd)(unsafe.Pointer(fds + uintptr(n)*8))).Ffd {
					goto _3
				}
				if !((*(*pollfd)(unsafe.Pointer(fds + uintptr(n)*8))).Frevents != 0) {
					goto _3
				}
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__20)), __ccgo_ts+917, libc.VaList(bp+8, (*(*pollfd)(unsafe.Pointer(fds + uintptr(n)*8))).Ffd))
				(*(*pollfd)(unsafe.Pointer(fds + uintptr(n)*8))).Frevents = 0
				num_ready--
				break
				goto _3
			_3:
				;
				n++
			}
			goto _2
		_2:
			;
			ievent_source = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
		}
	}
	usbi_mutex_unlock(tls, ctx+424)
	if num_ready != 0 {
		if v4 = num_ready > 0; !v4 {
			libc.X__assert_fail(tls, __ccgo_ts+959, __ccgo_ts+973, int32(332), uintptr(unsafe.Pointer(&__func__20)))
		}
		_ = v4 || libc.Bool(libc.Int32FromInt32(0) != 0)
		(*usbi_reported_events)(unsafe.Pointer(reported_events)).Fevent_data = fds
		(*usbi_reported_events)(unsafe.Pointer(reported_events)).Fevent_data_count = uint32(nfds)
	}
	goto done
done:
	;
	(*usbi_reported_events)(unsafe.Pointer(reported_events)).Fnum_ready = libc.Uint32FromInt32(num_ready)
	return int32(LIBUSB_SUCCESS)
}

var __func__20 = [21]uint8{'u', 's', 'b', 'i', '_', 'w', 'a', 'i', 't', '_', 'f', 'o', 'r', '_', 'e', 'v', 'e', 'n', 't', 's'}

func usbi_cond_init(tls *libc.TLS, cond uintptr) {
	var v1 bool
	_ = v1
	if v1 = libc.Xpthread_cond_init(tls, cond, libc.UintptrFromInt32(0)) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+1025, __ccgo_ts+1066, int32(51), uintptr(unsafe.Pointer(&__func__21)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
}

var __func__21 = [15]uint8{'u', 's', 'b', 'i', '_', 'c', 'o', 'n', 'd', '_', 'i', 'n', 'i', 't'}

func usbi_cond_timedwait(tls *libc.TLS, cond uintptr, mutex uintptr, tv uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r int32
	var _ /* timeout at bp+0 */ timespec
	_ = r
	usbi_get_real_time(tls, bp)
	(*(*timespec)(unsafe.Pointer(bp))).Ftv_sec += (*timeval)(unsafe.Pointer(tv)).Ftv_sec
	(*(*timespec)(unsafe.Pointer(bp))).Ftv_nsec += (*timeval)(unsafe.Pointer(tv)).Ftv_usec * int64(1000)
	if (*(*timespec)(unsafe.Pointer(bp))).Ftv_nsec >= int64(1000000000) {
		(*(*timespec)(unsafe.Pointer(bp))).Ftv_nsec -= int64(1000000000)
		(*(*timespec)(unsafe.Pointer(bp))).Ftv_sec++
	}
	r = libc.Xpthread_cond_timedwait(tls, cond, mutex, bp)
	if r == 0 {
		return 0
	} else {
		if r == int32(110) {
			return int32(LIBUSB_ERROR_TIMEOUT)
		} else {
			return int32(LIBUSB_ERROR_OTHER)
		}
	}
	return r1
}

func usbi_get_tid(tls *libc.TLS) (r uint64) {
	var tid, v1 uint64
	_, _ = tid, v1
	if tl_tid != 0 {
		return tl_tid
	}
	tid = libc.Uint64FromInt64(libc.Xsyscall(tls, int64(186), 0))
	if tid == libc.Uint64FromUint64(2)*libc.Uint64FromInt64(0x7fffffffffffffff)+libc.Uint64FromInt32(1) {
		tid = uint64(libc.Xpthread_self(tls))
	}
	v1 = tid
	tl_tid = v1
	return v1
}

var tl_tid uint64

type __s8 = int8

type __u8 = uint8

type __s16 = int16

type __u16 = uint16

type __s32 = int32

type __u32 = uint32

type __s64 = int64

type __u64 = uint64

type __kernel_fd_set = struct {
	Ffds_bits [16]uint64
}

type __kernel_sighandler_t = uintptr

type __kernel_key_t = int32

type __kernel_mqd_t = int32

type __kernel_old_uid_t = uint16

type __kernel_old_gid_t = uint16

type __kernel_old_dev_t = uint64

type __kernel_long_t = int64

type __kernel_ulong_t = uint64

type __kernel_ino_t = uint64

type __kernel_mode_t = uint32

type __kernel_pid_t = int32

type __kernel_ipc_pid_t = int32

type __kernel_uid_t = uint32

type __kernel_gid_t = uint32

type __kernel_suseconds_t = int64

type __kernel_daddr_t = int32

type __kernel_uid32_t = uint32

type __kernel_gid32_t = uint32

type __kernel_size_t = uint64

type __kernel_ssize_t = int64

type __kernel_ptrdiff_t = int64

type __kernel_fsid_t = struct {
	Fval [2]int32
}

type __kernel_off_t = int64

type __kernel_loff_t = int64

type __kernel_uoff_t = uint64

type __kernel_old_time_t = int64

type __kernel_time_t = int64

type __kernel_time64_t = int64

type __kernel_clock_t = int64

type __kernel_timer_t = int32

type __kernel_clockid_t = int32

type __kernel_caddr_t = uintptr

type __kernel_uid16_t = uint16

type __kernel_gid16_t = uint16

type __le16 = uint16

type __be16 = uint16

type __le32 = uint32

type __be32 = uint32

type __le64 = uint64

type __be64 = uint64

type __sum16 = uint16

type __wsum = uint32

type __poll_t = uint32

type usbfs_ctrltransfer = struct {
	FbmRequestType __u8
	FbRequest      __u8
	FwValue        __u16
	FwIndex        __u16
	FwLength       __u16
	Ftimeout       __u32
	Fdata          uintptr
}

type usbfs_setinterface = struct {
	Finterface1 uint32
	Faltsetting uint32
}

type usbfs_getdriver = struct {
	Finterface1 uint32
	Fdriver     [256]uint8
}

type usbfs_iso_packet_desc = struct {
	Flength        uint32
	Factual_length uint32
	Fstatus        uint32
}

type usbfs_urb = struct {
	Ftype1         uint8
	Fendpoint      uint8
	Fstatus        int32
	Fflags         uint32
	Fbuffer        uintptr
	Fbuffer_length int32
	Factual_length int32
	Fstart_frame   int32
	F__ccgo8_36    struct {
		Fstream_id         [0]uint32
		Fnumber_of_packets int32
	}
	Ferror_count int32
	Fsignr       uint32
	Fusercontext uintptr
}

type usbfs_connectinfo = struct {
	Fdevnum uint32
	Fslow   uint8
}

type usbfs_ioctl = struct {
	Fifno       int32
	Fioctl_code int32
	Fdata       uintptr
}

type usbfs_disconnect_claim = struct {
	Finterface1 uint32
	Fflags      uint32
	Fdriver     [256]uint8
}

type usbfs_streams = struct {
	Fnum_streams uint32
	Fnum_eps     uint32
}

func linux_start_event_monitor(tls *libc.TLS) (r int32) {
	return linux_netlink_start_event_monitor(tls)
}

func linux_stop_event_monitor(tls *libc.TLS) {
	linux_netlink_stop_event_monitor(tls)
}

func linux_hotplug_poll(tls *libc.TLS) {
	linux_netlink_hotplug_poll(tls)
}

func __isspace(tls *libc.TLS, _c int32) (r int32) {
	return libc.BoolInt32(_c == int32(' ') || libc.Uint32FromInt32(_c)-uint32('\t') < uint32(5))
}

type dirent = struct {
	Fd_ino    ino_t
	Fd_off    off_t
	Fd_reclen uint16
	Fd_type   uint8
	Fd_name   [256]uint8
}

type reclen_t = uint16

type posix_dent = struct {
	Fd_ino    ino_t
	Fd_off    off_t
	Fd_reclen reclen_t
	Fd_type   uint8
}

type __isoc_va_list = uintptr

type fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]uint8
}

type _G_fpos64_t = fpos_t

type cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = cookie_io_functions_t

type winsize = struct {
	Fws_row    uint16
	Fws_col    uint16
	Fws_xpixel uint16
	Fws_ypixel uint16
}

type utsname = struct {
	Fsysname    [65]uint8
	Fnodename   [65]uint8
	Frelease    [65]uint8
	Fversion    [65]uint8
	Fmachine    [65]uint8
	Fdomainname [65]uint8
}

type statvfs1 = struct {
	Ff_bsize    uint64
	Ff_frsize   uint64
	Ff_blocks   fsblkcnt_t
	Ff_bfree    fsblkcnt_t
	Ff_bavail   fsblkcnt_t
	Ff_files    fsfilcnt_t
	Ff_ffree    fsfilcnt_t
	Ff_favail   fsfilcnt_t
	Ff_fsid     uint64
	Ff_flag     uint64
	Ff_namemax  uint64
	Ff_type     uint32
	F__reserved [5]int32
}

type fsid_t = struct {
	F__val [2]int32
}

type __fsid_t = fsid_t

type statfs1 = struct {
	Ff_type    uint64
	Ff_bsize   uint64
	Ff_blocks  fsblkcnt_t
	Ff_bfree   fsblkcnt_t
	Ff_bavail  fsblkcnt_t
	Ff_files   fsfilcnt_t
	Ff_ffree   fsfilcnt_t
	Ff_fsid    fsid_t
	Ff_namelen uint64
	Ff_frsize  uint64
	Ff_flags   uint64
	Ff_spare   [4]uint64
}

var usbdev_names = int32(0)
var max_iso_packet_len = uint32(0)
var sysfs_available = -int32(1)
var init_count = int32(0)

type kernel_version = struct {
	Fmajor    int32
	Fminor    int32
	Fsublevel int32
}

type config_descriptor = struct {
	Fdesc       uintptr
	Factual_len size_t
}

type linux_context_priv = struct {
	Fno_device_discovery int32
}

type linux_device_priv = struct {
	Fsysfs_dir          uintptr
	Fdescriptors        uintptr
	Fdescriptors_len    size_t
	Fconfig_descriptors uintptr
	Factive_config      int32
}

type linux_device_handle_priv = struct {
	Ffd         int32
	Ffd_removed int32
	Ffd_keep    int32
	Fcaps       uint32_t
}

type reap_action = int32

const NORMAL = 0
const SUBMIT_FAILED = 1
const CANCELLED = 2
const COMPLETED_EARLY = 3
const ERROR = 4

type linux_transfer_priv = struct {
	F__ccgo0_0 struct {
		Fiso_urbs [0]uintptr
		Furbs     uintptr
	}
	Freap_action       reap_action
	Fnum_urbs          int32
	Fnum_retired       int32
	Freap_status       libusb_transfer_status
	Fiso_packet_offset int32
}

func dev_has_config0(tls *libc.TLS, dev uintptr) (r int32) {
	var config, priv uintptr
	var idx uint8_t
	_, _, _ = config, idx, priv
	priv = usbi_get_device_priv(tls, dev)
	idx = uint8(0)
	for {
		if !(libc.Int32FromUint8(idx) < libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations)) {
			break
		}
		config = (*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors + uintptr(idx)*16
		if libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer((*config_descriptor)(unsafe.Pointer(config)).Fdesc)).FbConfigurationValue) == 0 {
			return int32(1)
		}
		goto _1
	_1:
		;
		idx++
	}
	return 0
}

func get_usbfs_fd(tls *libc.TLS, dev uintptr, access_mode int32, silent int32) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var ctx uintptr
	var delay_ms int64
	var fd int32
	var retry, v1 uint8_t
	var _ /* delay_ts at bp+24 */ timespec
	var _ /* path at bp+0 */ [24]uint8
	_, _, _, _, _ = ctx, delay_ms, fd, retry, v1
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	if usbdev_names != 0 {
		libc.X__builtin_snprintf(tls, bp, uint64(24), __ccgo_ts+1119, libc.VaList(bp+48, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_address)))
	} else {
		libc.X__builtin_snprintf(tls, bp, uint64(24), __ccgo_ts+1136, libc.VaList(bp+48, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_address)))
	}
	fd = libc.Xopen(tls, bp, access_mode|int32(02000000), 0)
	if fd != -int32(1) {
		return fd
	}
	if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(2) {
		delay_ms = int64(10)
		*(*timespec)(unsafe.Pointer(bp + 24)) = timespec{
			Ftv_nsec: delay_ms * int64(1000) * int64(1000),
		}
		retry = uint8(3)
		for {
			v1 = retry
			retry--
			if !(libc.Int32FromUint8(v1) > 0) {
				break
			}
			if !(silent != 0) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__22)), __ccgo_ts+1159, libc.VaList(bp+48, delay_ms))
			}
			libc.Xnanosleep(tls, bp+24, libc.UintptrFromInt32(0))
			fd = libc.Xopen(tls, bp, access_mode|int32(02000000), 0)
			if fd != -int32(1) {
				return fd
			}
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(2) {
				break
			}
		}
	}
	if !(silent != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__22)), __ccgo_ts+1205, libc.VaList(bp+48, bp, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(13) && access_mode == int32(02) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__22)), __ccgo_ts+1250, 0)
		}
	}
	if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(13) {
		return int32(LIBUSB_ERROR_ACCESS)
	}
	if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(2) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	return int32(LIBUSB_ERROR_IO)
}

var __func__22 = [13]uint8{'g', 'e', 't', '_', 'u', 's', 'b', 'f', 's', '_', 'f', 'd'}

func is_usbdev_entry(tls *libc.TLS, name uintptr, bus_p uintptr, dev_p uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* busnum at bp+0 */ int32
	var _ /* devnum at bp+4 */ int32
	if libc.Xsscanf(tls, name, __ccgo_ts+1299, libc.VaList(bp+16, bp, bp+4)) != int32(2) {
		return 0
	}
	if *(*int32)(unsafe.Pointer(bp)) < 0 || *(*int32)(unsafe.Pointer(bp)) > int32(0xff) || *(*int32)(unsafe.Pointer(bp + 4)) < 0 || *(*int32)(unsafe.Pointer(bp + 4)) > int32(0xff) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__23)), __ccgo_ts+1311, libc.VaList(bp+16, name))
		return 0
	}
	usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__23)), __ccgo_ts+1338, libc.VaList(bp+16, name))
	if bus_p != 0 {
		*(*uint8_t)(unsafe.Pointer(bus_p)) = libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(bp)))
	}
	if dev_p != 0 {
		*(*uint8_t)(unsafe.Pointer(dev_p)) = libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(bp + 4)))
	}
	return int32(1)
}

var __func__23 = [16]uint8{'i', 's', '_', 'u', 's', 'b', 'd', 'e', 'v', '_', 'e', 'n', 't', 'r', 'y'}

func find_usbfs_path(tls *libc.TLS) (r uintptr) {
	var dir, entry, path, v1, v2 uintptr
	_, _, _, _, _ = dir, entry, path, v1, v2
	path = __ccgo_ts + 1348
	dir = libc.Xopendir(tls, path)
	if dir != 0 {
		for {
			v1 = libc.Xreaddir(tls, dir)
			entry = v1
			if !(v1 != 0) {
				break
			}
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(entry + 19))) == int32('.') {
				continue
			}
			break
		}
		libc.Xclosedir(tls, dir)
		if entry != 0 {
			return path
		}
	}
	path = __ccgo_ts + 1361
	dir = libc.Xopendir(tls, path)
	if dir != 0 {
		for {
			v2 = libc.Xreaddir(tls, dir)
			entry = v2
			if !(v2 != 0) {
				break
			}
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(entry + 19))) == int32('.') {
				continue
			}
			if is_usbdev_entry(tls, entry+19, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0)) != 0 {
				break
			}
		}
		libc.Xclosedir(tls, dir)
		if entry != 0 {
			usbdev_names = int32(1)
			return path
		}
	}
	return libc.UintptrFromInt32(0)
}

func get_kernel_version(tls *libc.TLS, ctx uintptr, ver uintptr) (r int32) {
	bp := tls.Alloc(432)
	defer tls.Free(432)
	var atoms int32
	var _ /* uts at bp+0 */ utsname
	_ = atoms
	if libc.Xuname(tls, bp) < 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__24)), __ccgo_ts+1366, libc.VaList(bp+400, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return -int32(1)
	}
	atoms = libc.Xsscanf(tls, bp+130, __ccgo_ts+1389, libc.VaList(bp+400, ver, ver+4, ver+8))
	if atoms < int32(2) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__24)), __ccgo_ts+1398, libc.VaList(bp+400, bp+130))
		return -int32(1)
	}
	if atoms < int32(3) {
		(*kernel_version)(unsafe.Pointer(ver)).Fsublevel = -int32(1)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__24)), __ccgo_ts+1433, libc.VaList(bp+400, bp+130))
	return 0
}

var __func__24 = [19]uint8{'g', 'e', 't', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'v', 'e', 'r', 's', 'i', 'o', 'n'}

func kernel_version_ge(tls *libc.TLS, ver uintptr, major int32, minor int32, sublevel int32) (r int32) {
	if (*kernel_version)(unsafe.Pointer(ver)).Fmajor > major {
		return int32(1)
	} else {
		if (*kernel_version)(unsafe.Pointer(ver)).Fmajor < major {
			return 0
		}
	}
	if (*kernel_version)(unsafe.Pointer(ver)).Fminor > minor {
		return int32(1)
	} else {
		if (*kernel_version)(unsafe.Pointer(ver)).Fminor < minor {
			return 0
		}
	}
	if (*kernel_version)(unsafe.Pointer(ver)).Fsublevel == -int32(1) {
		return libc.BoolInt32(sublevel == 0)
	}
	return libc.BoolInt32((*kernel_version)(unsafe.Pointer(ver)).Fsublevel >= sublevel)
}

func op_init(tls *libc.TLS, ctx uintptr) (r1 int32) {
	bp := tls.Alloc(176)
	defer tls.Free(176)
	var cpriv, usbfs_path uintptr
	var r, v1 int32
	var _ /* kversion at bp+0 */ kernel_version
	var _ /* statfsbuf at bp+16 */ statfs1
	_, _, _, _ = cpriv, r, usbfs_path, v1
	cpriv = usbi_get_context_priv(tls, ctx)
	if get_kernel_version(tls, ctx, bp) < 0 {
		return int32(LIBUSB_ERROR_OTHER)
	}
	if !(kernel_version_ge(tls, bp, int32(2), int32(6), int32(32)) != 0) {
		if (*(*kernel_version)(unsafe.Pointer(bp))).Fsublevel != -int32(1) {
			v1 = (*(*kernel_version)(unsafe.Pointer(bp))).Fsublevel
		} else {
			v1 = 0
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1463, libc.VaList(bp+144, (*(*kernel_version)(unsafe.Pointer(bp))).Fmajor, (*(*kernel_version)(unsafe.Pointer(bp))).Fminor, v1))
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	usbfs_path = find_usbfs_path(tls)
	if !(usbfs_path != 0) {
		usbfs_path = __ccgo_ts + 1348
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1512, libc.VaList(bp+144, usbfs_path))
	} else {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1551, libc.VaList(bp+144, usbfs_path))
	}
	if !(max_iso_packet_len != 0) {
		if kernel_version_ge(tls, bp, int32(5), int32(2), 0) != 0 {
			max_iso_packet_len = uint32(98304)
		} else {
			if kernel_version_ge(tls, bp, int32(3), int32(10), 0) != 0 {
				max_iso_packet_len = uint32(49152)
			} else {
				max_iso_packet_len = uint32(8192)
			}
		}
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1569, libc.VaList(bp+144, max_iso_packet_len))
	if sysfs_available == -int32(1) {
		r = libc.Xstatfs(tls, __ccgo_ts+1612, bp+16)
		if r == 0 && (*(*statfs1)(unsafe.Pointer(bp + 16))).Ff_type == uint64(0x62656572) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1617, 0)
			sysfs_available = int32(1)
		} else {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1636, 0)
			sysfs_available = 0
		}
	}
	if (*linux_context_priv)(unsafe.Pointer(cpriv)).Fno_device_discovery != 0 {
		return int32(LIBUSB_SUCCESS)
	}
	r = int32(LIBUSB_SUCCESS)
	if init_count == 0 {
		r = linux_start_event_monitor(tls)
	}
	if r == int32(LIBUSB_SUCCESS) {
		r = linux_scan_devices(tls, ctx)
		if r == int32(LIBUSB_SUCCESS) {
			init_count++
		} else {
			if init_count == 0 {
				linux_stop_event_monitor(tls)
			}
		}
	} else {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__25)), __ccgo_ts+1654, 0)
	}
	return r
}

var __func__25 = [8]uint8{'o', 'p', '_', 'i', 'n', 'i', 't'}

func op_exit(tls *libc.TLS, ctx uintptr) {
	var cpriv uintptr
	var v1 bool
	var v2 int32
	_, _, _ = cpriv, v1, v2
	cpriv = usbi_get_context_priv(tls, ctx)
	if (*linux_context_priv)(unsafe.Pointer(cpriv)).Fno_device_discovery != 0 {
		return
	}
	if v1 = init_count != 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+1691, __ccgo_ts+1707, int32(436), uintptr(unsafe.Pointer(&__func__26)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	init_count--
	v2 = init_count
	if !(v2 != 0) {
		linux_stop_event_monitor(tls)
	}
}

var __func__26 = [8]uint8{'o', 'p', '_', 'e', 'x', 'i', 't'}

func op_set_option(tls *libc.TLS, ctx uintptr, option libusb_option, ap va_list) (r int32) {
	var cpriv uintptr
	_ = cpriv
	_ = ap
	if option == int32(LIBUSB_OPTION_NO_DEVICE_DISCOVERY) {
		cpriv = usbi_get_context_priv(tls, ctx)
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__27)), __ccgo_ts+1758, 0)
		(*linux_context_priv)(unsafe.Pointer(cpriv)).Fno_device_discovery = int32(1)
		return int32(LIBUSB_SUCCESS)
	}
	return int32(LIBUSB_ERROR_NOT_SUPPORTED)
}

var __func__27 = [14]uint8{'o', 'p', '_', 's', 'e', 't', '_', 'o', 'p', 't', 'i', 'o', 'n'}

func op_get_device_string(tls *libc.TLS, dev uintptr, string_type libusb_device_string_type, buffer uintptr, length int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var attr, ctx, priv uintptr
	var fd int32
	var r, v1 ssize_t
	_, _, _, _, _, _ = attr, ctx, fd, priv, r, v1
	priv = usbi_get_device_priv(tls, dev)
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	switch string_type {
	case int32(LIBUSB_DEVICE_STRING_MANUFACTURER):
		attr = __ccgo_ts + 1796
	case int32(LIBUSB_DEVICE_STRING_PRODUCT):
		attr = __ccgo_ts + 1809
	case int32(LIBUSB_DEVICE_STRING_SERIAL_NUMBER):
		attr = __ccgo_ts + 1817
	case int32(LIBUSB_DEVICE_STRING_COUNT):
		fallthrough
	default:
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	fd = open_sysfs_attr(tls, ctx, (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir, attr)
	if fd < 0 {
		return int32(LIBUSB_ERROR_IO)
	}
	r = libc.Xread(tls, fd, buffer, libc.Uint64FromInt32(length-int32(1)))
	if r < 0 {
		r = int64(*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))
		libc.Xclose(tls, fd)
		if r == int64(19) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__28)), __ccgo_ts+1824, libc.VaList(bp+8, attr, r))
		return int32(LIBUSB_ERROR_IO)
	}
	libc.Xclose(tls, fd)
	*(*uint8)(unsafe.Pointer(buffer + uintptr(r))) = uint8(0)
	for r != 0 && (libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(buffer + uintptr(r)))) == 0 || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(buffer + uintptr(r)))) == int32('\n')) {
		v1 = r
		r--
		*(*uint8)(unsafe.Pointer(buffer + uintptr(v1))) = uint8(0)
	}
	r++
	return int32(r)
}

var __func__28 = [21]uint8{'o', 'p', '_', 'g', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 's', 't', 'r', 'i', 'n', 'g'}

func linux_scan_devices(tls *libc.TLS, ctx uintptr) (r int32) {
	var ret int32
	_ = ret
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
	ret = linux_default_scan_devices(tls, ctx)
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
	return ret
}

func op_hotplug_poll(tls *libc.TLS) {
	linux_hotplug_poll(tls)
}

func open_sysfs_attr(tls *libc.TLS, ctx uintptr, sysfs_dir uintptr, attr uintptr) (r int32) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var fd int32
	var _ /* filename at bp+0 */ [256]uint8
	_ = fd
	libc.X__builtin_snprintf(tls, bp, uint64(256), __ccgo_ts+1860, libc.VaList(bp+264, sysfs_dir, attr))
	fd = libc.Xopen(tls, bp, libc.Int32FromInt32(00)|libc.Int32FromInt32(02000000), 0)
	if fd < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(2) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__29)), __ccgo_ts+1887, libc.VaList(bp+264, bp, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_IO)
	}
	return fd
}

var __func__29 = [16]uint8{'o', 'p', 'e', 'n', '_', 's', 'y', 's', 'f', 's', '_', 'a', 't', 't', 'r'}

func read_sysfs_attr(tls *libc.TLS, ctx uintptr, sysfs_dir uintptr, attr uintptr, max_value int32, value_p uintptr) (r1 int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var fd int32
	var r ssize_t
	var value int64
	var _ /* buf at bp+0 */ [20]uint8
	var _ /* endptr at bp+24 */ uintptr
	_, _, _ = fd, r, value
	fd = open_sysfs_attr(tls, ctx, sysfs_dir, attr)
	if fd < 0 {
		return fd
	}
	r = libc.Xread(tls, fd, bp, libc.Uint64FromInt64(20)-libc.Uint64FromInt32(1))
	if r < 0 {
		r = int64(*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))))
		libc.Xclose(tls, fd)
		if r == int64(19) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__30)), __ccgo_ts+1824, libc.VaList(bp+40, attr, r))
		return int32(LIBUSB_ERROR_IO)
	}
	libc.Xclose(tls, fd)
	if r == 0 {
		*(*int32)(unsafe.Pointer(value_p)) = -int32(1)
		return 0
	}
	if !(libc.BoolInt32(uint32((*(*[20]uint8)(unsafe.Pointer(bp)))[0])-libc.Uint32FromUint8('0') < libc.Uint32FromInt32(10)) != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__30)), __ccgo_ts+1912, libc.VaList(bp+40, attr))
		return int32(LIBUSB_ERROR_IO)
	} else {
		if libc.Int32FromUint8((*(*[20]uint8)(unsafe.Pointer(bp)))[r-int64(1)]) != int32('\n') {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__30)), __ccgo_ts+1953, libc.VaList(bp+40, attr))
		} else {
			r--
		}
	}
	(*(*[20]uint8)(unsafe.Pointer(bp)))[r] = uint8('\000')
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	value = libc.Xstrtol(tls, bp, bp+24, int32(10))
	if bp == *(*uintptr)(unsafe.Pointer(bp + 24)) || value < 0 || value > int64(max_value) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__30)), __ccgo_ts+1992, libc.VaList(bp+40, attr, bp))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	} else {
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24))))) != int32('\000') {
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24))))) == int32('.') && libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24)) + libc.UintptrFromInt32(1))))-uint32('0') < uint32(10)) != 0 {
				*(*uintptr)(unsafe.Pointer(bp + 24))++
				for libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24)))))-uint32('0') < uint32(10)) != 0 {
					*(*uintptr)(unsafe.Pointer(bp + 24))++
				}
			}
			if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 24))))) != int32('\000') {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__30)), __ccgo_ts+1992, libc.VaList(bp+40, attr, bp))
				return int32(LIBUSB_ERROR_INVALID_PARAM)
			}
		}
	}
	*(*int32)(unsafe.Pointer(value_p)) = int32(value)
	return 0
}

var __func__30 = [16]uint8{'r', 'e', 'a', 'd', '_', 's', 'y', 's', 'f', 's', '_', 'a', 't', 't', 'r'}

func sysfs_scan_device(tls *libc.TLS, ctx uintptr, devname uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ret int32
	var _ /* busnum at bp+0 */ uint8_t
	var _ /* devaddr at bp+1 */ uint8_t
	_ = ret
	ret = linux_get_device_address(tls, ctx, 0, bp, bp+1, libc.UintptrFromInt32(0), devname, -int32(1))
	if ret != int32(LIBUSB_SUCCESS) {
		return ret
	}
	return linux_enumerate_device(tls, ctx, *(*uint8_t)(unsafe.Pointer(bp)), *(*uint8_t)(unsafe.Pointer(bp + 1)), devname)
}

func sysfs_get_active_config(tls *libc.TLS, dev uintptr, config uintptr) (r int32) {
	var priv uintptr
	_ = priv
	priv = usbi_get_device_priv(tls, dev)
	return read_sysfs_attr(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir, __ccgo_ts+2037, int32(0xff), config)
}

func linux_get_device_address(tls *libc.TLS, ctx uintptr, detached int32, busnum uintptr, devaddr uintptr, dev_node uintptr, sys_name uintptr, fd int32) (r1 int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var fd_path uintptr
	var r int32
	var _ /* proc_path at bp+4 */ [32]uint8
	var _ /* sysfs_val at bp+0 */ int32
	_, _ = fd_path, r
	tls.AllocaEntry()
	defer tls.AllocaExit()
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__31)), __ccgo_ts+2057, libc.VaList(bp+48, sys_name, detached))
	if !(sysfs_available != 0) || detached != 0 || !(sys_name != 0) {
		if !(dev_node != 0) && fd >= 0 {
			fd_path = libc.X__builtin_alloca(tls, uint64(4096))
			libc.X__builtin_snprintf(tls, bp+4, uint64(32), __ccgo_ts+2101, libc.VaList(bp+48, fd))
			r = int32(libc.Xreadlink(tls, bp+4, fd_path, libc.Uint64FromInt32(libc.Int32FromInt32(4096)-libc.Int32FromInt32(1))))
			if r > 0 {
				*(*uint8)(unsafe.Pointer(fd_path + uintptr(r))) = uint8('\000')
				dev_node = fd_path
			}
		}
		if !(dev_node != 0) {
			return int32(LIBUSB_ERROR_OTHER)
		}
		if !(libc.Xstrncmp(tls, dev_node, __ccgo_ts+1348, uint64(12)) != 0) {
			libc.Xsscanf(tls, dev_node, __ccgo_ts+2118, libc.VaList(bp+48, busnum, devaddr))
		} else {
			return int32(LIBUSB_ERROR_OTHER)
		}
		return int32(LIBUSB_SUCCESS)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__31)), __ccgo_ts+2141, libc.VaList(bp+48, sys_name))
	r = read_sysfs_attr(tls, ctx, sys_name, __ccgo_ts+2149, int32(0xff), bp)
	if r < 0 {
		return r
	}
	*(*uint8_t)(unsafe.Pointer(busnum)) = libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(bp)))
	r = read_sysfs_attr(tls, ctx, sys_name, __ccgo_ts+2156, int32(0xff), bp)
	if r < 0 {
		return r
	}
	*(*uint8_t)(unsafe.Pointer(devaddr)) = libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(bp)))
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__31)), __ccgo_ts+2163, libc.VaList(bp+48, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(busnum))), libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(devaddr)))))
	return int32(LIBUSB_SUCCESS)
}

var __func__31 = [25]uint8{'l', 'i', 'n', 'u', 'x', '_', 'g', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'a', 'd', 'd', 'r', 'e', 's', 's'}

func seek_to_next_config(tls *libc.TLS, ctx uintptr, buffer uintptr, len1 size_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var header uintptr
	var offset int32
	_, _ = header, offset
	offset = int32(9)
	buffer += uintptr(9)
	len1 -= uint64(9)
	for len1 > uint64(0) {
		if len1 < uint64(2) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__32)), __ccgo_ts+2177, libc.VaList(bp+8, len1))
			return int32(LIBUSB_ERROR_IO)
		}
		header = buffer
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_CONFIG) {
			return offset
		}
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(2) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__32)), __ccgo_ts+2221, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
			return int32(LIBUSB_ERROR_IO)
		}
		if len1 < uint64((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__32)), __ccgo_ts+2253, libc.VaList(bp+8, uint64((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)-len1))
			return int32(LIBUSB_ERROR_IO)
		}
		offset += libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		len1 -= uint64((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__32)), __ccgo_ts+2283, 0)
	return int32(LIBUSB_ERROR_IO)
}

var __func__32 = [20]uint8{'s', 'e', 'e', 'k', '_', 't', 'o', '_', 'n', 'e', 'x', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g'}

func parse_config_descriptors(tls *libc.TLS, dev uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buffer, config_desc, ctx, device_desc, priv uintptr
	var config_len, sysfs_config_len uint16_t
	var idx, num_configs uint8_t
	var offset int32
	var remaining size_t
	_, _, _, _, _, _, _, _, _, _, _ = buffer, config_desc, config_len, ctx, device_desc, idx, num_configs, offset, priv, remaining, sysfs_config_len
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	priv = usbi_get_device_priv(tls, dev)
	device_desc = (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors
	num_configs = (*usbi_device_descriptor)(unsafe.Pointer(device_desc)).FbNumConfigurations
	if libc.Int32FromUint8(num_configs) == 0 {
		return 0
	}
	(*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors = libc.Xmalloc(tls, uint64(num_configs)*uint64(16))
	if !((*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	buffer = (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors + uintptr(18)
	remaining = (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors_len - uint64(18)
	idx = uint8(0)
	for {
		if !(libc.Int32FromUint8(idx) < libc.Int32FromUint8(num_configs)) {
			break
		}
		if remaining < uint64(9) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2311, libc.VaList(bp+8, remaining, int32(9)))
			return int32(LIBUSB_ERROR_IO)
		}
		config_desc = buffer
		if libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer(config_desc)).FbDescriptorType) != int32(LIBUSB_DT_CONFIG) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2340, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer(config_desc)).FbDescriptorType)))
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer(config_desc)).FbLength) < int32(9) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2386, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer(config_desc)).FbLength)))
				return int32(LIBUSB_ERROR_IO)
			}
		}
		config_len = libusb_cpu_to_le16(tls, (*usbi_configuration_descriptor)(unsafe.Pointer(config_desc)).FwTotalLength)
		if libc.Int32FromUint16(config_len) < int32(9) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2416, libc.VaList(bp+8, libc.Int32FromUint16(config_len)))
			return int32(LIBUSB_ERROR_IO)
		}
		if (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir != 0 {
			if libc.Int32FromUint8(num_configs) > int32(1) && libc.Int32FromUint8(idx) < libc.Int32FromUint8(num_configs)-int32(1) {
				offset = seek_to_next_config(tls, ctx, buffer, remaining)
				if offset < 0 {
					return offset
				}
				sysfs_config_len = libc.Uint16FromInt32(offset)
			} else {
				sysfs_config_len = uint16(remaining)
			}
			if libc.Int32FromUint16(config_len) != libc.Int32FromUint16(sysfs_config_len) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2440, libc.VaList(bp+8, libc.Int32FromUint16(config_len), libc.Int32FromUint16(sysfs_config_len)))
				config_len = sysfs_config_len
			}
		} else {
			if uint64(config_len) > remaining {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2487, libc.VaList(bp+8, remaining, libc.Int32FromUint16(config_len)))
				config_len = uint16(remaining)
			}
		}
		if libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer(config_desc)).FbConfigurationValue) == 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__33)), __ccgo_ts+2516, 0)
		}
		(*(*config_descriptor)(unsafe.Pointer((*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors + uintptr(idx)*16))).Fdesc = config_desc
		(*(*config_descriptor)(unsafe.Pointer((*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors + uintptr(idx)*16))).Factual_len = uint64(config_len)
		buffer += uintptr(config_len)
		remaining -= uint64(config_len)
		goto _1
	_1:
		;
		idx++
	}
	return int32(LIBUSB_SUCCESS)
}

var __func__33 = [25]uint8{'p', 'a', 'r', 's', 'e', '_', 'c', 'o', 'n', 'f', 'i', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', 's'}

func op_get_config_descriptor_by_value(tls *libc.TLS, dev uintptr, value uint8_t, buffer uintptr) (r int32) {
	var config, priv uintptr
	var idx uint8_t
	_, _, _ = config, idx, priv
	priv = usbi_get_device_priv(tls, dev)
	idx = uint8(0)
	for {
		if !(libc.Int32FromUint8(idx) < libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations)) {
			break
		}
		config = (*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors + uintptr(idx)*16
		if libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer((*config_descriptor)(unsafe.Pointer(config)).Fdesc)).FbConfigurationValue) == libc.Int32FromUint8(value) {
			*(*uintptr)(unsafe.Pointer(buffer)) = (*config_descriptor)(unsafe.Pointer(config)).Fdesc
			return libc.Int32FromUint64((*config_descriptor)(unsafe.Pointer(config)).Factual_len)
		}
		goto _1
	_1:
		;
		idx++
	}
	return int32(LIBUSB_ERROR_NOT_FOUND)
}

func op_get_active_config_descriptor(tls *libc.TLS, dev uintptr, buffer uintptr, len1 size_t) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var priv uintptr
	var r int32
	var v1 uint64
	var _ /* active_config at bp+8 */ int32
	var _ /* config_desc at bp+0 */ uintptr
	_, _, _ = priv, r, v1
	priv = usbi_get_device_priv(tls, dev)
	if (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir != 0 {
		r = sysfs_get_active_config(tls, dev, bp+8)
		if r < 0 {
			return r
		}
	} else {
		*(*int32)(unsafe.Pointer(bp + 8)) = (*linux_device_priv)(unsafe.Pointer(priv)).Factive_config
	}
	if *(*int32)(unsafe.Pointer(bp + 8)) == -int32(1) {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__34)), __ccgo_ts+2543, 0)
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	r = op_get_config_descriptor_by_value(tls, dev, libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(bp + 8))), bp)
	if r < 0 {
		return r
	}
	if len1 < libc.Uint64FromInt32(r) {
		v1 = len1
	} else {
		v1 = libc.Uint64FromInt32(r)
	}
	len1 = v1
	libc.Xmemcpy(tls, buffer, *(*uintptr)(unsafe.Pointer(bp)), len1)
	return libc.Int32FromUint64(len1)
}

var __func__34 = [32]uint8{'o', 'p', '_', 'g', 'e', 't', '_', 'a', 'c', 't', 'i', 'v', 'e', '_', 'c', 'o', 'n', 'f', 'i', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func op_get_config_descriptor(tls *libc.TLS, dev uintptr, config_index uint8_t, buffer uintptr, len1 size_t) (r int32) {
	var config, priv uintptr
	var v1 uint64
	_, _, _ = config, priv, v1
	priv = usbi_get_device_priv(tls, dev)
	if libc.Int32FromUint8(config_index) >= libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations) {
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	config = (*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors + uintptr(config_index)*16
	if len1 < (*config_descriptor)(unsafe.Pointer(config)).Factual_len {
		v1 = len1
	} else {
		v1 = (*config_descriptor)(unsafe.Pointer(config)).Factual_len
	}
	len1 = v1
	libc.Xmemcpy(tls, buffer, (*config_descriptor)(unsafe.Pointer(config)).Fdesc, len1)
	return libc.Int32FromUint64(len1)
}

func usbfs_get_active_config(tls *libc.TLS, dev uintptr, fd int32) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var priv uintptr
	var r int32
	var _ /* active_config at bp+0 */ uint8_t
	var _ /* ctrl at bp+8 */ usbfs_ctrltransfer
	_, _ = priv, r
	priv = usbi_get_device_priv(tls, dev)
	*(*uint8_t)(unsafe.Pointer(bp)) = uint8(0)
	*(*usbfs_ctrltransfer)(unsafe.Pointer(bp + 8)) = usbfs_ctrltransfer{
		FbmRequestType: uint8(LIBUSB_ENDPOINT_IN),
		FbRequest:      uint8(LIBUSB_REQUEST_GET_CONFIGURATION),
		FwLength:       uint16(1),
		Ftimeout:       uint32(1000),
		Fdata:          bp,
	}
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64((libc.Uint32FromUint32(2)|libc.Uint32FromUint32(1))<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(0)))|libc.Uint64FromInt64(24)<<libc.Int32FromInt32(16)), libc.VaList(bp+40, bp+8))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__35)), __ccgo_ts+2563, libc.VaList(bp+40, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		if (*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors != 0 {
			(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer((*(*config_descriptor)(unsafe.Pointer((*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors))).Fdesc)).FbConfigurationValue)
		} else {
			(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = -int32(1)
		}
	} else {
		if libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(bp))) == 0 {
			if dev_has_config0(tls, dev) != 0 {
				(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = 0
			} else {
				(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = -int32(1)
			}
		} else {
			(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(bp)))
		}
	}
	return int32(LIBUSB_SUCCESS)
}

var __func__35 = [24]uint8{'u', 's', 'b', 'f', 's', '_', 'g', 'e', 't', '_', 'a', 'c', 't', 'i', 'v', 'e', '_', 'c', 'o', 'n', 'f', 'i', 'g'}

func usbfs_get_speed(tls *libc.TLS, ctx uintptr, fd int32) (r1 libusb_speed) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r int32
	_ = r
	r = libc.Xioctl(tls, fd, libc.Int32FromUint32(libc.Uint32FromUint32(0)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(31))|libc.Uint32FromInt32(libc.Int32FromInt32(0)<<libc.Int32FromInt32(16))), libc.VaList(bp+8, libc.UintptrFromInt32(0)))
	switch r {
	case 0:
		return int32(LIBUSB_SPEED_UNKNOWN)
	case int32(1):
		return int32(LIBUSB_SPEED_LOW)
	case int32(2):
		return int32(LIBUSB_SPEED_FULL)
	case int32(3):
		return int32(LIBUSB_SPEED_HIGH)
	case int32(4):
		return int32(LIBUSB_SPEED_HIGH)
	case int32(5):
		return int32(LIBUSB_SPEED_SUPER)
	case int32(6):
		return int32(LIBUSB_SPEED_SUPER_PLUS)
	default:
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__36)), __ccgo_ts+2598, libc.VaList(bp+8, r))
	}
	return int32(LIBUSB_SPEED_UNKNOWN)
}

var __func__36 = [16]uint8{'u', 's', 'b', 'f', 's', '_', 'g', 'e', 't', '_', 's', 'p', 'e', 'e', 'd'}

func initialize_device(tls *libc.TLS, dev uintptr, busnum uint8_t, devaddr uint8_t, sysfs_dir uintptr, wrapped_fd int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var alloc_len, desc_read_length size_t
	var ctx, priv, read_ptr uintptr
	var fd, r int32
	var nb ssize_t
	var _ /* speed at bp+0 */ int32
	_, _, _, _, _, _, _, _ = alloc_len, ctx, desc_read_length, fd, nb, priv, r, read_ptr
	priv = usbi_get_device_priv(tls, dev)
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	(*libusb_device)(unsafe.Pointer(dev)).Fbus_number = busnum
	(*libusb_device)(unsafe.Pointer(dev)).Fdevice_address = devaddr
	if sysfs_dir != 0 {
		(*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir = libc.Xstrdup(tls, sysfs_dir)
		if !((*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir != 0) {
			return int32(LIBUSB_ERROR_NO_MEM)
		}
		if read_sysfs_attr(tls, ctx, sysfs_dir, __ccgo_ts+2629, int32(0x7fffffff), bp) == 0 {
			switch *(*int32)(unsafe.Pointer(bp)) {
			case int32(1):
				(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_LOW)
			case int32(12):
				(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_FULL)
			case int32(480):
				(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_HIGH)
			case int32(5000):
				(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_SUPER)
			case int32(10000):
				(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_SUPER_PLUS)
			case int32(20000):
				(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_SUPER_PLUS_X2)
			default:
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__37)), __ccgo_ts+2635, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(bp))))
			}
		}
	} else {
		if wrapped_fd >= 0 {
			(*libusb_device)(unsafe.Pointer(dev)).Fspeed = usbfs_get_speed(tls, ctx, wrapped_fd)
		}
	}
	if sysfs_dir != 0 {
		fd = open_sysfs_attr(tls, ctx, sysfs_dir, __ccgo_ts+2665)
	} else {
		if wrapped_fd < 0 {
			fd = get_usbfs_fd(tls, dev, 00, 0)
		} else {
			fd = wrapped_fd
			r = int32(libc.Xlseek(tls, fd, 0, 0))
			if r < 0 {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__37)), __ccgo_ts+2677, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
				return int32(LIBUSB_ERROR_IO)
			}
		}
	}
	if fd < 0 {
		return fd
	}
	alloc_len = uint64(0)
	for cond := true; cond; cond = (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors_len == alloc_len {
		desc_read_length = uint64(256)
		alloc_len += desc_read_length
		(*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors = usbi_reallocf(tls, (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors, alloc_len)
		if !((*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors != 0) {
			if fd != wrapped_fd {
				libc.Xclose(tls, fd)
			}
			return int32(LIBUSB_ERROR_NO_MEM)
		}
		read_ptr = (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors + uintptr((*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors_len)
		if !(sysfs_dir != 0) {
			libc.Xmemset(tls, read_ptr, 0, desc_read_length)
		}
		nb = libc.Xread(tls, fd, read_ptr, desc_read_length)
		if nb < 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__37)), __ccgo_ts+2700, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			if fd != wrapped_fd {
				libc.Xclose(tls, fd)
			}
			return int32(LIBUSB_ERROR_IO)
		}
		*(*size_t)(unsafe.Pointer(priv + 16)) += libc.Uint64FromInt64(nb)
	}
	if fd != wrapped_fd {
		libc.Xclose(tls, fd)
	}
	if (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors_len < uint64(18) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__37)), __ccgo_ts+2733, libc.VaList(bp+16, (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors_len))
		return int32(LIBUSB_ERROR_IO)
	}
	r = parse_config_descriptors(tls, dev)
	if r < 0 {
		return r
	}
	libc.Xmemcpy(tls, dev+56, (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors, uint64(18))
	if sysfs_dir != 0 {
		usbi_localize_device_descriptor(tls, dev+56)
		return int32(LIBUSB_SUCCESS)
	}
	if wrapped_fd < 0 {
		fd = get_usbfs_fd(tls, dev, int32(02), int32(1))
	} else {
		fd = wrapped_fd
	}
	if fd < 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__37)), __ccgo_ts+2761, 0)
		if (*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors != 0 {
			(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = libc.Int32FromUint8((*usbi_configuration_descriptor)(unsafe.Pointer((*(*config_descriptor)(unsafe.Pointer((*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors))).Fdesc)).FbConfigurationValue)
		} else {
			(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = -int32(1)
		}
		return int32(LIBUSB_SUCCESS)
	}
	r = usbfs_get_active_config(tls, dev, fd)
	if fd != wrapped_fd {
		libc.Xclose(tls, fd)
	}
	return r
}

var __func__37 = [18]uint8{'i', 'n', 'i', 't', 'i', 'a', 'l', 'i', 'z', 'e', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func linux_get_parent_info(tls *libc.TLS, dev uintptr, sysfs_dir uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var add_parent, ret int32
	var ctx, it, priv, start, tmp, v1, v2 uintptr
	var port_number int64
	var v3 bool
	var _ /* end at bp+8 */ uintptr
	var _ /* parent_sysfs_dir at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _ = add_parent, ctx, it, port_number, priv, ret, start, tmp, v1, v2, v3
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	add_parent = int32(1)
	if !(sysfs_dir != 0) || !(libc.Xstrncmp(tls, sysfs_dir, __ccgo_ts+2835, uint64(3)) != 0) {
		return int32(LIBUSB_SUCCESS)
	}
	*(*uintptr)(unsafe.Pointer(bp)) = libc.Xstrdup(tls, sysfs_dir)
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	v1 = libc.Xstrrchr(tls, *(*uintptr)(unsafe.Pointer(bp)), int32('.'))
	tmp = v1
	if v3 = v1 != 0; !v3 {
		v2 = libc.Xstrrchr(tls, *(*uintptr)(unsafe.Pointer(bp)), int32('-'))
		tmp = v2
	}
	if v3 || v2 != 0 {
		start = tmp + uintptr(1)
		port_number = libc.Xstrtol(tls, start, bp+8, int32(10))
		if port_number < 0 || port_number > int64(0x7fffffff) || start == *(*uintptr)(unsafe.Pointer(bp + 8)) || int32('\000') != libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 8))))) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__38)), __ccgo_ts+2839, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(bp))))
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp)))
			return int32(LIBUSB_ERROR_OTHER)
		} else {
			(*libusb_device)(unsafe.Pointer(dev)).Fport_number = libc.Uint8FromInt32(int32(port_number))
		}
		*(*uint8)(unsafe.Pointer(tmp)) = uint8('\000')
	} else {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__38)), __ccgo_ts+2891, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(bp))))
		libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp)))
		return int32(LIBUSB_SUCCESS)
	}
	if !(libc.Xstrchr(tls, *(*uintptr)(unsafe.Pointer(bp)), int32('-')) != 0) {
		tmp = *(*uintptr)(unsafe.Pointer(bp))
		ret = libc.Xasprintf(tls, bp, __ccgo_ts+2935, libc.VaList(bp+24, tmp))
		libc.Xfree(tls, tmp)
		if ret < 0 {
			return int32(LIBUSB_ERROR_NO_MEM)
		}
	}
	goto retry
retry:
	;
	usbi_mutex_lock(tls, ctx+24)
	it = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+64)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(it+32 != ctx+64) {
			break
		}
		priv = usbi_get_device_priv(tls, it)
		if (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir != 0 {
			if !(libc.Xstrcmp(tls, (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir, *(*uintptr)(unsafe.Pointer(bp))) != 0) {
				(*libusb_device)(unsafe.Pointer(dev)).Fparent_dev = libusb_ref_device(tls, it)
				break
			}
		}
		goto _4
	_4:
		;
		it = uintptr(uint64((*libusb_device)(unsafe.Pointer(it)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	usbi_mutex_unlock(tls, ctx+24)
	if !((*libusb_device)(unsafe.Pointer(dev)).Fparent_dev != 0) && add_parent != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__38)), __ccgo_ts+2941, libc.VaList(bp+24, *(*uintptr)(unsafe.Pointer(bp))))
		sysfs_scan_device(tls, ctx, *(*uintptr)(unsafe.Pointer(bp)))
		add_parent = 0
		goto retry
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__38)), __ccgo_ts+2991, libc.VaList(bp+24, dev, sysfs_dir, (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev, *(*uintptr)(unsafe.Pointer(bp)), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fport_number)))
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bp)))
	return int32(LIBUSB_SUCCESS)
}

var __func__38 = [22]uint8{'l', 'i', 'n', 'u', 'x', '_', 'g', 'e', 't', '_', 'p', 'a', 'r', 'e', 'n', 't', '_', 'i', 'n', 'f', 'o'}

func linux_enumerate_device(tls *libc.TLS, ctx uintptr, busnum uint8_t, devaddr uint8_t, sysfs_dir uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var dev uintptr
	var r int32
	var session_id uint64
	_, _, _ = dev, r, session_id
	session_id = libc.Uint64FromInt32(libc.Int32FromUint8(busnum)<<int32(8) | libc.Int32FromUint8(devaddr))
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__39)), __ccgo_ts+3030, libc.VaList(bp+8, libc.Int32FromUint8(busnum), libc.Int32FromUint8(devaddr), session_id))
	dev = usbi_get_device_by_session_id(tls, ctx, session_id)
	if dev != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__39)), __ccgo_ts+3066, libc.VaList(bp+8, session_id))
		libusb_unref_device(tls, dev)
		return int32(LIBUSB_SUCCESS)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__39)), __ccgo_ts+3096, libc.VaList(bp+8, libc.Int32FromUint8(busnum), libc.Int32FromUint8(devaddr), session_id))
	dev = usbi_alloc_device(tls, ctx, session_id)
	if !(dev != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = initialize_device(tls, dev, busnum, devaddr, sysfs_dir, -int32(1))
	if r < 0 {
		goto out
	}
	r = usbi_sanitize_device(tls, dev)
	if r < 0 {
		goto out
	}
	r = linux_get_parent_info(tls, dev, sysfs_dir)
	if r < 0 {
		goto out
	}
	goto out
out:
	;
	if r < 0 {
		libusb_unref_device(tls, dev)
	} else {
		usbi_connect_device(tls, dev)
	}
	return r
}

var __func__39 = [23]uint8{'l', 'i', 'n', 'u', 'x', '_', 'e', 'n', 'u', 'm', 'e', 'r', 'a', 't', 'e', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func linux_hotplug_enumerate(tls *libc.TLS, busnum uint8_t, devaddr uint8_t, sys_name uintptr) {
	var ctx uintptr
	_ = ctx
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	ctx = uintptr(uint64((*list_head)(unsafe.Pointer(uintptr(unsafe.Pointer(&active_contexts_list)))).Fnext) - uint64(libc.UintptrFromInt32(0)+552))
	for {
		if !(ctx+552 != uintptr(unsafe.Pointer(&active_contexts_list))) {
			break
		}
		linux_enumerate_device(tls, ctx, busnum, devaddr, sys_name)
		goto _1
	_1:
		;
		ctx = uintptr(uint64((*libusb_context)(unsafe.Pointer(ctx)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+552))
	}
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
}

func linux_device_disconnected(tls *libc.TLS, busnum uint8_t, devaddr uint8_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ctx, dev uintptr
	var session_id uint64
	_, _, _ = ctx, dev, session_id
	session_id = libc.Uint64FromInt32(libc.Int32FromUint8(busnum)<<int32(8) | libc.Int32FromUint8(devaddr))
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	ctx = uintptr(uint64((*list_head)(unsafe.Pointer(uintptr(unsafe.Pointer(&active_contexts_list)))).Fnext) - uint64(libc.UintptrFromInt32(0)+552))
	for {
		if !(ctx+552 != uintptr(unsafe.Pointer(&active_contexts_list))) {
			break
		}
		dev = usbi_get_device_by_session_id(tls, ctx, session_id)
		if dev != 0 {
			usbi_disconnect_device(tls, dev)
			libusb_unref_device(tls, dev)
		} else {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__40)), __ccgo_ts+3142, libc.VaList(bp+8, session_id))
		}
		goto _1
	_1:
		;
		ctx = uintptr(uint64((*libusb_context)(unsafe.Pointer(ctx)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+552))
	}
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
}

var __func__40 = [26]uint8{'l', 'i', 'n', 'u', 'x', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'd', 'i', 's', 'c', 'o', 'n', 'n', 'e', 'c', 't', 'e', 'd'}

func parse_u8(tls *libc.TLS, str uintptr, val_p uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var num int64
	var _ /* endptr at bp+0 */ uintptr
	_ = num
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	num = libc.Xstrtol(tls, str, bp, int32(10))
	if num < 0 || num > int64(libc.Int32FromInt32(0xff)) || *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
		return 0
	}
	if *(*uintptr)(unsafe.Pointer(bp)) == str || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp))))) != int32('\000') {
		return 0
	}
	*(*uint8_t)(unsafe.Pointer(val_p)) = libc.Uint8FromInt64(num)
	return int32(1)
}

func usbfs_scan_busdir(tls *libc.TLS, ctx uintptr, busnum uint8_t) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var dir, entry, v1 uintptr
	var r int32
	var _ /* devaddr at bp+20 */ uint8_t
	var _ /* dirpath at bp+0 */ [20]uint8
	_, _, _, _ = dir, entry, r, v1
	r = int32(LIBUSB_ERROR_IO)
	libc.X__builtin_snprintf(tls, bp, uint64(20), __ccgo_ts+3175, libc.VaList(bp+32, libc.Int32FromUint8(busnum)))
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__41)), __ccgo_ts+3193, libc.VaList(bp+32, bp))
	dir = libc.Xopendir(tls, bp)
	if !(dir != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__41)), __ccgo_ts+3196, libc.VaList(bp+32, bp, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return r
	}
	for {
		v1 = libc.Xreaddir(tls, dir)
		entry = v1
		if !(v1 != 0) {
			break
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(entry + 19))) == int32('.') {
			continue
		}
		if !(parse_u8(tls, entry+19, bp+20) != 0) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__41)), __ccgo_ts+3226, libc.VaList(bp+32, entry+19))
			continue
		}
		if linux_enumerate_device(tls, ctx, busnum, *(*uint8_t)(unsafe.Pointer(bp + 20)), libc.UintptrFromInt32(0)) != 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__41)), __ccgo_ts+3247, libc.VaList(bp+32, entry+19))
			continue
		}
		r = 0
	}
	libc.Xclosedir(tls, dir)
	return r
}

var __func__41 = [18]uint8{'u', 's', 'b', 'f', 's', '_', 's', 'c', 'a', 'n', '_', 'b', 'u', 's', 'd', 'i', 'r'}

func usbfs_get_device_list(tls *libc.TLS, ctx uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buses, entry, v1 uintptr
	var r int32
	var _ /* busnum at bp+0 */ uint8_t
	var _ /* devaddr at bp+1 */ uint8_t
	_, _, _, _ = buses, entry, r, v1
	r = 0
	if usbdev_names != 0 {
		buses = libc.Xopendir(tls, __ccgo_ts+1361)
	} else {
		buses = libc.Xopendir(tls, __ccgo_ts+1348)
	}
	if !(buses != 0) {
		if !(usbdev_names != 0) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(2) {
			return int32(LIBUSB_SUCCESS)
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__42)), __ccgo_ts+3280, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_IO)
	}
	for {
		v1 = libc.Xreaddir(tls, buses)
		entry = v1
		if !(v1 != 0) {
			break
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(entry + 19))) == int32('.') {
			continue
		}
		if usbdev_names != 0 {
			if !(is_usbdev_entry(tls, entry+19, bp, bp+1) != 0) {
				continue
			}
			r = linux_enumerate_device(tls, ctx, *(*uint8_t)(unsafe.Pointer(bp)), *(*uint8_t)(unsafe.Pointer(bp + 1)), libc.UintptrFromInt32(0))
			if r < 0 {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__42)), __ccgo_ts+3247, libc.VaList(bp+16, entry+19))
				continue
			}
		} else {
			if !(parse_u8(tls, entry+19, bp) != 0) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__42)), __ccgo_ts+3226, libc.VaList(bp+16, entry+19))
				continue
			}
			r = usbfs_scan_busdir(tls, ctx, *(*uint8_t)(unsafe.Pointer(bp)))
			if r < 0 {
				break
			}
		}
	}
	libc.Xclosedir(tls, buses)
	return r
}

var __func__42 = [22]uint8{'u', 's', 'b', 'f', 's', '_', 'g', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'l', 'i', 's', 't'}

func sysfs_get_device_list(tls *libc.TLS, ctx uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var devices, entry, v1 uintptr
	var num_devices, num_enumerated int32
	_, _, _, _, _ = devices, entry, num_devices, num_enumerated, v1
	devices = libc.Xopendir(tls, __ccgo_ts+3311)
	num_devices = 0
	num_enumerated = 0
	if !(devices != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__43)), __ccgo_ts+3332, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_IO)
	}
	for {
		v1 = libc.Xreaddir(tls, devices)
		entry = v1
		if !(v1 != 0) {
			break
		}
		if !(libc.BoolInt32(uint32(*(*uint8)(unsafe.Pointer(entry + 19)))-libc.Uint32FromUint8('0') < libc.Uint32FromInt32(10)) != 0) && libc.Xstrncmp(tls, entry+19, __ccgo_ts+2835, uint64(3)) != 0 || libc.Xstrchr(tls, entry+19, int32(':')) != 0 {
			continue
		}
		num_devices++
		if sysfs_scan_device(tls, ctx, entry+19) != 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__43)), __ccgo_ts+3247, libc.VaList(bp+8, entry+19))
			continue
		}
		num_enumerated++
	}
	libc.Xclosedir(tls, devices)
	if num_enumerated != 0 || !(num_devices != 0) {
		return int32(LIBUSB_SUCCESS)
	} else {
		return int32(LIBUSB_ERROR_IO)
	}
	return r
}

var __func__43 = [22]uint8{'s', 'y', 's', 'f', 's', '_', 'g', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'l', 'i', 's', 't'}

func linux_default_scan_devices(tls *libc.TLS, ctx uintptr) (r int32) {
	if sysfs_available != 0 && sysfs_get_device_list(tls, ctx) == int32(LIBUSB_SUCCESS) {
		return int32(LIBUSB_SUCCESS)
	}
	return usbfs_get_device_list(tls, ctx)
}

func initialize_handle(tls *libc.TLS, handle uintptr, fd int32) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var hpriv, v1, v2, v3 uintptr
	var r int32
	_, _, _, _, _ = hpriv, r, v1, v2, v3
	hpriv = usbi_get_device_handle_priv(tls, handle)
	(*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd = fd
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(26)))|libc.Uint64FromInt64(4)<<libc.Int32FromInt32(16)), libc.VaList(bp+8, hpriv+12))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(25) {
			if handle != 0 {
				v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
			} else {
				v1 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__44)), __ccgo_ts+3365, 0)
		} else {
			if handle != 0 {
				v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
			} else {
				v2 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__44)), __ccgo_ts+3386, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		}
		(*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Fcaps = uint32(0x02)
	}
	if handle != 0 {
		v3 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
	} else {
		v3 = libc.UintptrFromInt32(0)
	}
	return usbi_add_event_source(tls, v3, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, int16(0x004))
}

var __func__44 = [18]uint8{'i', 'n', 'i', 't', 'i', 'a', 'l', 'i', 'z', 'e', '_', 'h', 'a', 'n', 'd', 'l', 'e'}

func op_wrap_sys_device(tls *libc.TLS, ctx uintptr, handle uintptr, sys_dev intptr_t) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var __atomic_store_ptr, dev, hpriv uintptr
	var fd, r int32
	var _ /* __atomic_store_tmp at bp+16 */ usbi_atomic_t
	var _ /* busnum at bp+0 */ uint8_t
	var _ /* ci at bp+8 */ usbfs_connectinfo
	var _ /* devaddr at bp+1 */ uint8_t
	_, _, _, _, _ = __atomic_store_ptr, dev, fd, hpriv, r
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = int32(sys_dev)
	r = linux_get_device_address(tls, ctx, int32(1), bp, bp+1, libc.UintptrFromInt32(0), libc.UintptrFromInt32(0), fd)
	if r < 0 {
		r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(17)))|libc.Uint64FromInt64(8)<<libc.Int32FromInt32(16)), libc.VaList(bp+32, bp+8))
		if r < 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__45)), __ccgo_ts+3410, libc.VaList(bp+32, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			return int32(LIBUSB_ERROR_IO)
		}
		*(*uint8_t)(unsafe.Pointer(bp)) = uint8(0)
		*(*uint8_t)(unsafe.Pointer(bp + 1)) = uint8((*(*usbfs_connectinfo)(unsafe.Pointer(bp + 8))).Fdevnum)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__45)), __ccgo_ts+3439, libc.VaList(bp+32, fd))
	dev = usbi_alloc_device(tls, ctx, uint64(0))
	if !(dev != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = initialize_device(tls, dev, *(*uint8_t)(unsafe.Pointer(bp)), *(*uint8_t)(unsafe.Pointer(bp + 1)), libc.UintptrFromInt32(0), fd)
	if r < 0 {
		goto out
	}
	r = usbi_sanitize_device(tls, dev)
	if r < 0 {
		goto out
	}
	{
		__atomic_store_ptr = dev + 80
		*(*usbi_atomic_t)(unsafe.Pointer(bp + 16)) = int64(libc.Int32FromInt32(1))
		libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp+16, libc.Int32FromInt32(5))
	}
	(*libusb_device_handle)(unsafe.Pointer(handle)).Fdev = dev
	r = initialize_handle(tls, handle, fd)
	(*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd_keep = int32(1)
	goto out
out:
	;
	if r < 0 {
		libusb_unref_device(tls, dev)
	}
	return r
}

var __func__45 = [19]uint8{'o', 'p', '_', 'w', 'r', 'a', 'p', '_', 's', 'y', 's', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func op_open(tls *libc.TLS, handle uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_load_ptr, v2 uintptr
	var fd, r int32
	var v1 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _, _, _ = __atomic_load_ptr, fd, r, v1, v2
	fd = get_usbfs_fd(tls, (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev, int32(02), 0)
	if fd < 0 {
		if fd == int32(LIBUSB_ERROR_NO_DEVICE) {
			usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
			{
				__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev + 80
				libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
				v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
			}
			if v1 != 0 {
				if handle != 0 {
					v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
				} else {
					v2 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__46)), __ccgo_ts+3471, 0)
				linux_device_disconnected(tls, (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fbus_number, (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fdevice_address)
			}
			usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
		}
		return fd
	}
	r = initialize_handle(tls, handle, fd)
	if r < 0 {
		libc.Xclose(tls, fd)
	}
	return r
}

var __func__46 = [8]uint8{'o', 'p', '_', 'o', 'p', 'e', 'n'}

func op_close(tls *libc.TLS, dev_handle uintptr) {
	var hpriv, v1 uintptr
	_, _ = hpriv, v1
	hpriv = usbi_get_device_handle_priv(tls, dev_handle)
	if !((*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd_removed != 0) {
		if dev_handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_remove_event_source(tls, v1, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd)
	}
	if !((*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd_keep != 0) {
		libc.Xclose(tls, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd)
	}
}

func op_get_configuration(tls *libc.TLS, handle uintptr, config uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var hpriv, priv, v1 uintptr
	var r int32
	var _ /* active_config at bp+0 */ int32
	_, _, _, _ = hpriv, priv, r, v1
	priv = usbi_get_device_priv(tls, (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)
	*(*int32)(unsafe.Pointer(bp)) = -int32(1)
	if (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir != 0 {
		r = sysfs_get_active_config(tls, (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev, bp)
	} else {
		hpriv = usbi_get_device_handle_priv(tls, handle)
		r = usbfs_get_active_config(tls, (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd)
		if r == int32(LIBUSB_SUCCESS) {
			*(*int32)(unsafe.Pointer(bp)) = (*linux_device_priv)(unsafe.Pointer(priv)).Factive_config
		}
	}
	if r < 0 {
		return r
	}
	if *(*int32)(unsafe.Pointer(bp)) == -int32(1) {
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__47)), __ccgo_ts+2543, 0)
		*(*int32)(unsafe.Pointer(bp)) = 0
	}
	*(*uint8_t)(unsafe.Pointer(config)) = libc.Uint8FromInt32(*(*int32)(unsafe.Pointer(bp)))
	return 0
}

var __func__47 = [21]uint8{'o', 'p', '_', 'g', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'a', 't', 'i', 'o', 'n'}

func op_set_configuration(tls *libc.TLS, handle uintptr, _config int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*int32)(unsafe.Pointer(bp)) = _config
	var fd, r int32
	var hpriv, priv, v1 uintptr
	_, _, _, _, _ = fd, hpriv, priv, r, v1
	priv = usbi_get_device_priv(tls, (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(5)))|libc.Uint64FromInt64(4)<<libc.Int32FromInt32(16)), libc.VaList(bp+16, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(16) {
				return int32(LIBUSB_ERROR_BUSY)
			} else {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
					return int32(LIBUSB_ERROR_NO_DEVICE)
				}
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__48)), __ccgo_ts+3525, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	if !((*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir != 0) {
		if *(*int32)(unsafe.Pointer(bp)) == 0 && !(dev_has_config0(tls, (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev) != 0) {
			*(*int32)(unsafe.Pointer(bp)) = -int32(1)
		}
		(*linux_device_priv)(unsafe.Pointer(priv)).Factive_config = *(*int32)(unsafe.Pointer(bp))
	}
	return int32(LIBUSB_SUCCESS)
}

var __func__48 = [21]uint8{'o', 'p', '_', 's', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'a', 't', 'i', 'o', 'n'}

func claim_interface(tls *libc.TLS, handle uintptr, _iface uint32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*uint32)(unsafe.Pointer(bp)) = _iface
	var fd, r int32
	var hpriv, v1 uintptr
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(15)))|libc.Uint64FromInt64(4)<<libc.Int32FromInt32(16)), libc.VaList(bp+16, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(2) {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(16) {
				return int32(LIBUSB_ERROR_BUSY)
			} else {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
					return int32(LIBUSB_ERROR_NO_DEVICE)
				}
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__49)), __ccgo_ts+3560, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return 0
}

var __func__49 = [16]uint8{'c', 'l', 'a', 'i', 'm', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e'}

func release_interface(tls *libc.TLS, handle uintptr, _iface uint32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*uint32)(unsafe.Pointer(bp)) = _iface
	var fd, r int32
	var hpriv, v1 uintptr
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(16)))|libc.Uint64FromInt64(4)<<libc.Int32FromInt32(16)), libc.VaList(bp+16, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__50)), __ccgo_ts+3593, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return 0
}

var __func__50 = [18]uint8{'r', 'e', 'l', 'e', 'a', 's', 'e', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e'}

func op_set_interface(tls *libc.TLS, handle uintptr, interface1 uint8_t, altsetting uint8_t) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var fd, r int32
	var hpriv, v1 uintptr
	var _ /* setintf at bp+0 */ usbfs_setinterface
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	(*(*usbfs_setinterface)(unsafe.Pointer(bp))).Finterface1 = uint32(interface1)
	(*(*usbfs_setinterface)(unsafe.Pointer(bp))).Faltsetting = uint32(altsetting)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(4)))|libc.Uint64FromInt64(8)<<libc.Int32FromInt32(16)), libc.VaList(bp+16, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
				return int32(LIBUSB_ERROR_NO_DEVICE)
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__51)), __ccgo_ts+3628, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return 0
}

var __func__51 = [17]uint8{'o', 'p', '_', 's', 'e', 't', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e'}

func op_clear_halt(tls *libc.TLS, handle uintptr, endpoint uint8) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var fd, r int32
	var hpriv, v1 uintptr
	var _ /* _endpoint at bp+0 */ uint32
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	*(*uint32)(unsafe.Pointer(bp)) = uint32(endpoint)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(21)))|libc.Uint64FromInt64(4)<<libc.Int32FromInt32(16)), libc.VaList(bp+16, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(2) {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
				return int32(LIBUSB_ERROR_NO_DEVICE)
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__52)), __ccgo_ts+3659, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return 0
}

var __func__52 = [14]uint8{'o', 'p', '_', 'c', 'l', 'e', 'a', 'r', '_', 'h', 'a', 'l', 't'}

func op_reset_device(tls *libc.TLS, handle uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var fd, r, ret int32
	var hpriv, v2, v4 uintptr
	var i uint8_t
	_, _, _, _, _, _, _ = fd, hpriv, i, r, ret, v2, v4
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	ret = 0
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < int32(32)) {
			break
		}
		if (*libusb_device_handle)(unsafe.Pointer(handle)).Fclaimed_interfaces&(uint64(1)<<i) != 0 {
			release_interface(tls, handle, uint32(i))
		}
		goto _1
	_1:
		;
		i++
	}
	usbi_mutex_lock(tls, handle)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint32(libc.Uint32FromUint32(0)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(20))|libc.Uint32FromInt32(libc.Int32FromInt32(0)<<libc.Int32FromInt32(16))), libc.VaList(bp+8, libc.UintptrFromInt32(0)))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			ret = int32(LIBUSB_ERROR_NOT_FOUND)
			goto out
		}
		if handle != 0 {
			v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v2 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__53)), __ccgo_ts+3687, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		ret = int32(LIBUSB_ERROR_OTHER)
		goto out
	}
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < int32(32)) {
			break
		}
		if !((*libusb_device_handle)(unsafe.Pointer(handle)).Fclaimed_interfaces&(libc.Uint64FromUint64(1)<<i) != 0) {
			goto _3
		}
		r = detach_kernel_driver_and_claim(tls, handle, i)
		if r != 0 {
			if handle != 0 {
				v4 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
			} else {
				v4 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v4, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__53)), __ccgo_ts+3710, libc.VaList(bp+8, libc.Int32FromUint8(i), libusb_error_name(tls, r)))
			*(*uint64)(unsafe.Pointer(handle + 40)) &= ^(libc.Uint64FromUint64(1) << i)
			ret = int32(LIBUSB_ERROR_NOT_FOUND)
		}
		goto _3
	_3:
		;
		i++
	}
	goto out
out:
	;
	usbi_mutex_unlock(tls, handle)
	return ret
}

var __func__53 = [16]uint8{'o', 'p', '_', 'r', 'e', 's', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func do_streams_ioctl(tls *libc.TLS, handle uintptr, req uint64, num_streams uint32_t, endpoints uintptr, num_endpoints int32) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var fd, r int32
	var hpriv, streams, v1 uintptr
	_, _, _, _, _ = fd, hpriv, r, streams, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	if num_endpoints > int32(30) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	streams = libc.Xmalloc(tls, uint64(8)+libc.Uint64FromInt32(num_endpoints))
	if !(streams != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*usbfs_streams)(unsafe.Pointer(streams)).Fnum_streams = num_streams
	(*usbfs_streams)(unsafe.Pointer(streams)).Fnum_eps = libc.Uint32FromInt32(num_endpoints)
	libc.Xmemcpy(tls, streams+8, endpoints, libc.Uint64FromInt32(num_endpoints))
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(req), libc.VaList(bp+8, streams))
	libc.Xfree(tls, streams)
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(25) {
			return int32(LIBUSB_ERROR_NOT_SUPPORTED)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
				return int32(LIBUSB_ERROR_INVALID_PARAM)
			} else {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
					return int32(LIBUSB_ERROR_NO_DEVICE)
				}
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__54)), __ccgo_ts+3758, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return r
}

var __func__54 = [17]uint8{'d', 'o', '_', 's', 't', 'r', 'e', 'a', 'm', 's', '_', 'i', 'o', 'c', 't', 'l'}

func op_alloc_streams(tls *libc.TLS, handle uintptr, num_streams uint32_t, endpoints uintptr, num_endpoints int32) (r int32) {
	return do_streams_ioctl(tls, handle, uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(28)))|libc.Uint64FromInt64(8)<<libc.Int32FromInt32(16), num_streams, endpoints, num_endpoints)
}

func op_free_streams(tls *libc.TLS, handle uintptr, endpoints uintptr, num_endpoints int32) (r int32) {
	return do_streams_ioctl(tls, handle, uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(29)))|libc.Uint64FromInt64(8)<<libc.Int32FromInt32(16), uint32(0), endpoints, num_endpoints)
}

func op_dev_mem_alloc(tls *libc.TLS, handle uintptr, len1 size_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var buffer, hpriv, v1 uintptr
	_, _, _ = buffer, hpriv, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	buffer = libc.Xmmap(tls, libc.UintptrFromInt32(0), len1, libc.Int32FromInt32(1)|libc.Int32FromInt32(2), int32(0x01), (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, 0)
	if buffer == uintptr(-libc.Int32FromInt32(1)) {
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__55)), __ccgo_ts+3789, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return libc.UintptrFromInt32(0)
	}
	return buffer
}

var __func__55 = [17]uint8{'o', 'p', '_', 'd', 'e', 'v', '_', 'm', 'e', 'm', '_', 'a', 'l', 'l', 'o', 'c'}

func op_dev_mem_free(tls *libc.TLS, handle uintptr, buffer uintptr, len1 size_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uintptr
	_ = v1
	if libc.Xmunmap(tls, buffer, len1) != 0 {
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__56)), __ccgo_ts+3820, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	} else {
		return int32(LIBUSB_SUCCESS)
	}
	return r
}

var __func__56 = [16]uint8{'o', 'p', '_', 'd', 'e', 'v', '_', 'm', 'e', 'm', '_', 'f', 'r', 'e', 'e'}

func op_kernel_driver_active(tls *libc.TLS, handle uintptr, interface1 uint8_t) (r1 int32) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var fd, r int32
	var hpriv, v1 uintptr
	var _ /* getdrv at bp+0 */ usbfs_getdriver
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	(*(*usbfs_getdriver)(unsafe.Pointer(bp))).Finterface1 = uint32(interface1)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(8)))|libc.Uint64FromInt64(260)<<libc.Int32FromInt32(16)), libc.VaList(bp+272, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(61) {
			return 0
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
				return int32(LIBUSB_ERROR_NO_DEVICE)
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__57)), __ccgo_ts+3850, libc.VaList(bp+272, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return libc.BoolInt32(libc.Xstrcmp(tls, bp+4, __ccgo_ts+3878) != 0)
}

var __func__57 = [24]uint8{'o', 'p', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r', '_', 'a', 'c', 't', 'i', 'v', 'e'}

func op_detach_kernel_driver(tls *libc.TLS, handle uintptr, interface1 uint8_t) (r1 int32) {
	bp := tls.Alloc(304)
	defer tls.Free(304)
	var fd, r int32
	var hpriv, v1 uintptr
	var _ /* command at bp+0 */ usbfs_ioctl
	var _ /* getdrv at bp+16 */ usbfs_getdriver
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	(*(*usbfs_ioctl)(unsafe.Pointer(bp))).Fifno = libc.Int32FromUint8(interface1)
	(*(*usbfs_ioctl)(unsafe.Pointer(bp))).Fioctl_code = libc.Int32FromUint32(libc.Uint32FromUint32(0)<<libc.Int32FromInt32(30) | libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8)) | libc.Uint32FromInt32(libc.Int32FromInt32(22)) | libc.Uint32FromInt32(libc.Int32FromInt32(0)<<libc.Int32FromInt32(16)))
	(*(*usbfs_ioctl)(unsafe.Pointer(bp))).Fdata = libc.UintptrFromInt32(0)
	(*(*usbfs_getdriver)(unsafe.Pointer(bp + 16))).Finterface1 = uint32(interface1)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(8)))|libc.Uint64FromInt64(260)<<libc.Int32FromInt32(16)), libc.VaList(bp+288, bp+16))
	if r == 0 && !(libc.Xstrcmp(tls, bp+16+4, __ccgo_ts+3878) != 0) {
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64((libc.Uint32FromUint32(2)|libc.Uint32FromUint32(1))<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(18)))|libc.Uint64FromInt64(16)<<libc.Int32FromInt32(16)), libc.VaList(bp+288, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(61) {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
				return int32(LIBUSB_ERROR_INVALID_PARAM)
			} else {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
					return int32(LIBUSB_ERROR_NO_DEVICE)
				}
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__58)), __ccgo_ts+3884, libc.VaList(bp+288, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return 0
}

var __func__58 = [24]uint8{'o', 'p', '_', 'd', 'e', 't', 'a', 'c', 'h', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r'}

func op_attach_kernel_driver(tls *libc.TLS, handle uintptr, interface1 uint8_t) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var fd, r int32
	var hpriv, v1 uintptr
	var _ /* command at bp+0 */ usbfs_ioctl
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	(*(*usbfs_ioctl)(unsafe.Pointer(bp))).Fifno = libc.Int32FromUint8(interface1)
	(*(*usbfs_ioctl)(unsafe.Pointer(bp))).Fioctl_code = libc.Int32FromUint32(libc.Uint32FromUint32(0)<<libc.Int32FromInt32(30) | libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8)) | libc.Uint32FromInt32(libc.Int32FromInt32(23)) | libc.Uint32FromInt32(libc.Int32FromInt32(0)<<libc.Int32FromInt32(16)))
	(*(*usbfs_ioctl)(unsafe.Pointer(bp))).Fdata = libc.UintptrFromInt32(0)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64((libc.Uint32FromUint32(2)|libc.Uint32FromUint32(1))<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(18)))|libc.Uint64FromInt64(16)<<libc.Int32FromInt32(16)), libc.VaList(bp+24, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(61) {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
				return int32(LIBUSB_ERROR_INVALID_PARAM)
			} else {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
					return int32(LIBUSB_ERROR_NO_DEVICE)
				} else {
					if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(16) {
						return int32(LIBUSB_ERROR_BUSY)
					}
				}
			}
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__59)), __ccgo_ts+3908, libc.VaList(bp+24, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	} else {
		if r == 0 {
			return int32(LIBUSB_ERROR_NOT_FOUND)
		}
	}
	return 0
}

var __func__59 = [24]uint8{'o', 'p', '_', 'a', 't', 't', 'a', 'c', 'h', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r'}

func detach_kernel_driver_and_claim(tls *libc.TLS, handle uintptr, interface1 uint8_t) (r1 int32) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var fd, r int32
	var hpriv, v1 uintptr
	var _ /* dc at bp+0 */ usbfs_disconnect_claim
	_, _, _, _ = fd, hpriv, r, v1
	hpriv = usbi_get_device_handle_priv(tls, handle)
	fd = (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd
	(*(*usbfs_disconnect_claim)(unsafe.Pointer(bp))).Finterface1 = uint32(interface1)
	libc.Xstrcpy(tls, bp+8, __ccgo_ts+3878)
	(*(*usbfs_disconnect_claim)(unsafe.Pointer(bp))).Fflags = uint32(0x02)
	r = libc.Xioctl(tls, fd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(27)))|libc.Uint64FromInt64(264)<<libc.Int32FromInt32(16)), libc.VaList(bp+272, bp))
	if r == 0 {
		return 0
	}
	switch *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) {
	case int32(25):
	case int32(16):
		return int32(LIBUSB_ERROR_BUSY)
	case int32(22):
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	case int32(19):
		return int32(LIBUSB_ERROR_NO_DEVICE)
	default:
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__60)), __ccgo_ts+3932, libc.VaList(bp+272, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_OTHER)
	}
	r = op_detach_kernel_driver(tls, handle, interface1)
	if r != 0 && r != int32(LIBUSB_ERROR_NOT_FOUND) {
		return r
	}
	return claim_interface(tls, handle, uint32(interface1))
}

var __func__60 = [31]uint8{'d', 'e', 't', 'a', 'c', 'h', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r', '_', 'a', 'n', 'd', '_', 'c', 'l', 'a', 'i', 'm'}

func op_claim_interface(tls *libc.TLS, handle uintptr, interface1 uint8_t) (r int32) {
	if (*libusb_device_handle)(unsafe.Pointer(handle)).Fauto_detach_kernel_driver != 0 {
		return detach_kernel_driver_and_claim(tls, handle, interface1)
	} else {
		return claim_interface(tls, handle, uint32(interface1))
	}
	return r
}

func op_release_interface(tls *libc.TLS, handle uintptr, interface1 uint8_t) (r1 int32) {
	var r int32
	_ = r
	r = release_interface(tls, handle, uint32(interface1))
	if r != 0 {
		return r
	}
	if (*libusb_device_handle)(unsafe.Pointer(handle)).Fauto_detach_kernel_driver != 0 {
		op_attach_kernel_driver(tls, handle, interface1)
	}
	return 0
}

func op_destroy_device(tls *libc.TLS, dev uintptr) {
	var priv uintptr
	_ = priv
	priv = usbi_get_device_priv(tls, dev)
	libc.Xfree(tls, (*linux_device_priv)(unsafe.Pointer(priv)).Fconfig_descriptors)
	libc.Xfree(tls, (*linux_device_priv)(unsafe.Pointer(priv)).Fdescriptors)
	libc.Xfree(tls, (*linux_device_priv)(unsafe.Pointer(priv)).Fsysfs_dir)
}

func discard_urbs(tls *libc.TLS, itransfer uintptr, first int32, last_plus_one int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var hpriv, tpriv, transfer, urb, v2, v3, v4 uintptr
	var i, ret int32
	_, _, _, _, _, _, _, _, _ = hpriv, i, ret, tpriv, transfer, urb, v2, v3, v4
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	hpriv = usbi_get_device_handle_priv(tls, (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)
	ret = 0
	i = last_plus_one - int32(1)
	for {
		if !(i >= first) {
			break
		}
		if libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) == int32(LIBUSB_TRANSFER_TYPE_ISOCHRONOUS) {
			urb = *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(tpriv)) + uintptr(i)*8))
		} else {
			urb = (*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs + uintptr(i)*56
		}
		if libc.Xioctl(tls, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, libc.Int32FromUint32(libc.Uint32FromUint32(0)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(11))|libc.Uint32FromInt32(libc.Int32FromInt32(0)<<libc.Int32FromInt32(16))), libc.VaList(bp+8, urb)) == 0 {
			goto _1
		}
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v2 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v2 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__61)), __ccgo_ts+3970, 0)
			if i == last_plus_one-int32(1) {
				ret = int32(LIBUSB_ERROR_NOT_FOUND)
			}
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
				if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
					v3 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
				} else {
					v3 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__61)), __ccgo_ts+4016, 0)
				ret = int32(LIBUSB_ERROR_NO_DEVICE)
			} else {
				if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
					v4 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
				} else {
					v4 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v4, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__61)), __ccgo_ts+4073, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
				ret = int32(LIBUSB_ERROR_OTHER)
			}
		}
		goto _1
	_1:
		;
		i--
	}
	return ret
}

var __func__61 = [13]uint8{'d', 'i', 's', 'c', 'a', 'r', 'd', '_', 'u', 'r', 'b', 's'}

func free_iso_urbs(tls *libc.TLS, tpriv uintptr) {
	var i int32
	var urb uintptr
	_, _ = i, urb
	i = 0
	for {
		if !(i < (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs) {
			break
		}
		urb = *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(tpriv)) + uintptr(i)*8))
		if !(urb != 0) {
			break
		}
		libc.Xfree(tls, urb)
		goto _1
	_1:
		;
		i++
	}
	libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(tpriv)))
	*(*uintptr)(unsafe.Pointer(tpriv)) = libc.UintptrFromInt32(0)
}

func submit_bulk_transfer(tls *libc.TLS, itransfer uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bulk_buffer_len, i, is_out, last_urb_partial, num_urbs, r, use_bulk_continuation, v1, v2, v7 int32
	var hpriv, tpriv, transfer, urb, urbs, v3, v5, v6, v8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bulk_buffer_len, hpriv, i, is_out, last_urb_partial, num_urbs, r, tpriv, transfer, urb, urbs, use_bulk_continuation, v1, v2, v3, v5, v6, v7, v8
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	hpriv = usbi_get_device_handle_priv(tls, (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)
	is_out = libc.BoolInt32(!(libc.Int32FromInt32(0) != libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint)&int32(LIBUSB_ENDPOINT_IN)))
	last_urb_partial = 0
	if (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Fcaps&uint32(0x08) != 0 {
		if (*libusb_transfer)(unsafe.Pointer(transfer)).Flength != 0 {
			v1 = (*libusb_transfer)(unsafe.Pointer(transfer)).Flength
		} else {
			v1 = int32(1)
		}
		bulk_buffer_len = v1
		use_bulk_continuation = 0
	} else {
		if (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Fcaps&uint32(0x02) != 0 {
			bulk_buffer_len = int32(16384)
			use_bulk_continuation = int32(1)
		} else {
			if (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Fcaps&uint32(0x04) != 0 {
				if (*libusb_transfer)(unsafe.Pointer(transfer)).Flength != 0 {
					v2 = (*libusb_transfer)(unsafe.Pointer(transfer)).Flength
				} else {
					v2 = int32(1)
				}
				bulk_buffer_len = v2
				use_bulk_continuation = 0
			} else {
				bulk_buffer_len = int32(16384)
				use_bulk_continuation = 0
			}
		}
	}
	num_urbs = (*libusb_transfer)(unsafe.Pointer(transfer)).Flength / bulk_buffer_len
	if (*libusb_transfer)(unsafe.Pointer(transfer)).Flength == 0 {
		num_urbs = int32(1)
	} else {
		if (*libusb_transfer)(unsafe.Pointer(transfer)).Flength%bulk_buffer_len > 0 {
			last_urb_partial = int32(1)
			num_urbs++
		}
	}
	if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
		v3 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
	} else {
		v3 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__62)), __ccgo_ts+4103, libc.VaList(bp+8, num_urbs, (*libusb_transfer)(unsafe.Pointer(transfer)).Flength))
	urbs = libc.Xcalloc(tls, libc.Uint64FromInt32(num_urbs), uint64(56))
	if !(urbs != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = urbs
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs = num_urbs
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired = 0
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(NORMAL)
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status = int32(LIBUSB_TRANSFER_COMPLETED)
	i = 0
	for {
		if !(i < num_urbs) {
			break
		}
		urb = urbs + uintptr(i)*56
		(*usbfs_urb)(unsafe.Pointer(urb)).Fusercontext = itransfer
		switch libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) {
		case int32(LIBUSB_TRANSFER_TYPE_BULK):
			(*usbfs_urb)(unsafe.Pointer(urb)).Ftype1 = uint8(3)
			*(*uint32)(unsafe.Add(unsafe.Pointer(urb), 36)) = uint32(0)
		case int32(LIBUSB_TRANSFER_TYPE_BULK_STREAM):
			(*usbfs_urb)(unsafe.Pointer(urb)).Ftype1 = uint8(3)
			*(*uint32)(unsafe.Add(unsafe.Pointer(urb), 36)) = (*usbi_transfer)(unsafe.Pointer(itransfer)).Fstream_id
		case int32(LIBUSB_TRANSFER_TYPE_INTERRUPT):
			(*usbfs_urb)(unsafe.Pointer(urb)).Ftype1 = uint8(1)
			break
		}
		(*usbfs_urb)(unsafe.Pointer(urb)).Fendpoint = (*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint
		(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer = (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer + uintptr(i*bulk_buffer_len)
		if use_bulk_continuation != 0 && !(is_out != 0) && i < num_urbs-int32(1) {
			(*usbfs_urb)(unsafe.Pointer(urb)).Fflags = uint32(0x01)
		}
		if i == num_urbs-int32(1) && last_urb_partial != 0 {
			(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length = (*libusb_transfer)(unsafe.Pointer(transfer)).Flength % bulk_buffer_len
		} else {
			if (*libusb_transfer)(unsafe.Pointer(transfer)).Flength == 0 {
				(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length = 0
			} else {
				(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length = bulk_buffer_len
			}
		}
		if i > 0 && use_bulk_continuation != 0 {
			*(*uint32)(unsafe.Pointer(urb + 8)) |= uint32(0x04)
		}
		if is_out != 0 && i == num_urbs-int32(1) && libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Fflags)&int32(LIBUSB_TRANSFER_ADD_ZERO_PACKET) != 0 {
			*(*uint32)(unsafe.Pointer(urb + 8)) |= uint32(0x40)
		}
		r = libc.Xioctl(tls, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(10)))|libc.Uint64FromInt64(56)<<libc.Int32FromInt32(16)), libc.VaList(bp+8, urb))
		if r == 0 {
			goto _4
		}
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			r = int32(LIBUSB_ERROR_NO_DEVICE)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(12) {
				r = int32(LIBUSB_ERROR_NO_MEM)
			} else {
				if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
					v5 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
				} else {
					v5 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v5, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__62)), __ccgo_ts+4148, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
				r = int32(LIBUSB_ERROR_IO)
			}
		}
		if i == 0 {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v6 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v6 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v6, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__62)), __ccgo_ts+4175, 0)
			libc.Xfree(tls, urbs)
			(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = libc.UintptrFromInt32(0)
			return r
		}
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(121) {
			v7 = int32(COMPLETED_EARLY)
		} else {
			v7 = int32(SUBMIT_FAILED)
		}
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = v7
		*(*int32)(unsafe.Pointer(tpriv + 16)) += num_urbs - i
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(COMPLETED_EARLY) {
			return 0
		}
		discard_urbs(tls, itransfer, 0, i)
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v8 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v8 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v8, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__62)), __ccgo_ts+4204, libc.VaList(bp+8, i))
		return 0
		goto _4
	_4:
		;
		i++
	}
	return 0
}

var __func__62 = [21]uint8{'s', 'u', 'b', 'm', 'i', 't', '_', 'b', 'u', 'l', 'k', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func submit_iso_transfer(tls *libc.TLS, itransfer uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var alloc_size size_t
	var hpriv, tpriv, transfer, urb, urb_buffer, urbs, v10, v11, v12, v13, v2, v3, v9, p7 uintptr
	var i, j, k, num_packets, num_packets_in_urb, num_packets_remaining, num_urbs, r, v5 int32
	var packet_len, total_len uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alloc_size, hpriv, i, j, k, num_packets, num_packets_in_urb, num_packets_remaining, num_urbs, packet_len, r, total_len, tpriv, transfer, urb, urb_buffer, urbs, v10, v11, v12, v13, v2, v3, v5, v9, p7
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	hpriv = usbi_get_device_handle_priv(tls, (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)
	num_packets = (*libusb_transfer)(unsafe.Pointer(transfer)).Fnum_iso_packets
	total_len = uint32(0)
	urb_buffer = (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer
	if num_packets < int32(1) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	i = 0
	for {
		if !(i < num_packets) {
			break
		}
		packet_len = (*(*libusb_iso_packet_descriptor)(unsafe.Pointer(transfer + 60 + uintptr(i)*12))).Flength
		if packet_len > max_iso_packet_len {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v2 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v2 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4287, libc.VaList(bp+8, packet_len, max_iso_packet_len))
			return int32(LIBUSB_ERROR_INVALID_PARAM)
		}
		total_len += packet_len
		goto _1
	_1:
		;
		i++
	}
	if (*libusb_transfer)(unsafe.Pointer(transfer)).Flength < libc.Int32FromUint32(total_len) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	num_urbs = (num_packets + (libc.Int32FromInt32(128) - libc.Int32FromInt32(1))) / int32(128)
	if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
		v3 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
	} else {
		v3 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4103, libc.VaList(bp+8, num_urbs, (*libusb_transfer)(unsafe.Pointer(transfer)).Flength))
	urbs = libc.Xcalloc(tls, libc.Uint64FromInt32(num_urbs), uint64(8))
	if !(urbs != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	*(*uintptr)(unsafe.Pointer(tpriv)) = urbs
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs = num_urbs
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired = 0
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(NORMAL)
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fiso_packet_offset = 0
	num_packets_remaining = num_packets
	i = 0
	j = libc.Int32FromInt32(0)
	for {
		if !(i < num_urbs) {
			break
		}
		if num_packets_remaining < int32(128) {
			v5 = num_packets_remaining
		} else {
			v5 = int32(128)
		}
		num_packets_in_urb = v5
		alloc_size = uint64(56) + libc.Uint64FromInt32(num_packets_in_urb)*uint64(12)
		urb = libc.Xcalloc(tls, uint64(1), alloc_size)
		if !(urb != 0) {
			free_iso_urbs(tls, tpriv)
			return int32(LIBUSB_ERROR_NO_MEM)
		}
		*(*uintptr)(unsafe.Pointer(urbs + uintptr(i)*8)) = urb
		k = 0
		for {
			if !(k < num_packets_in_urb) {
				break
			}
			packet_len = (*(*libusb_iso_packet_descriptor)(unsafe.Pointer(transfer + 60 + uintptr(j)*12))).Flength
			p7 = urb + 24
			*(*int32)(unsafe.Pointer(p7)) = int32(uint32(*(*int32)(unsafe.Pointer(p7))) + packet_len)
			(*(*usbfs_iso_packet_desc)(unsafe.Pointer(urb + 56 + uintptr(k)*12))).Flength = packet_len
			goto _6
		_6:
			;
			j++
			k++
		}
		(*usbfs_urb)(unsafe.Pointer(urb)).Fusercontext = itransfer
		(*usbfs_urb)(unsafe.Pointer(urb)).Ftype1 = uint8(0)
		(*usbfs_urb)(unsafe.Pointer(urb)).Fflags = uint32(0x02)
		(*usbfs_urb)(unsafe.Pointer(urb)).Fendpoint = (*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint
		(*usbfs_urb)(unsafe.Pointer(urb)).F__ccgo8_36.Fnumber_of_packets = num_packets_in_urb
		(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer = urb_buffer
		urb_buffer += uintptr((*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length)
		num_packets_remaining -= num_packets_in_urb
		goto _4
	_4:
		;
		i++
	}
	i = 0
	for {
		if !(i < num_urbs) {
			break
		}
		r = libc.Xioctl(tls, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(10)))|libc.Uint64FromInt64(56)<<libc.Int32FromInt32(16)), libc.VaList(bp+8, *(*uintptr)(unsafe.Pointer(urbs + uintptr(i)*8))))
		if r == 0 {
			goto _8
		}
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			r = int32(LIBUSB_ERROR_NO_DEVICE)
		} else {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
				if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
					v9 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
				} else {
					v9 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v9, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4345, 0)
				r = int32(LIBUSB_ERROR_INVALID_PARAM)
			} else {
				if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(90) {
					if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
						v10 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
					} else {
						v10 = libc.UintptrFromInt32(0)
					}
					usbi_log(tls, v10, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4382, 0)
					r = int32(LIBUSB_ERROR_INVALID_PARAM)
				} else {
					if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
						v11 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
					} else {
						v11 = libc.UintptrFromInt32(0)
					}
					usbi_log(tls, v11, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4148, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
					r = int32(LIBUSB_ERROR_IO)
				}
			}
		}
		if i == 0 {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v12 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v12 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v12, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4175, 0)
			free_iso_urbs(tls, tpriv)
			return r
		}
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(SUBMIT_FAILED)
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired = num_urbs - i
		discard_urbs(tls, itransfer, 0, i)
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v13 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v13 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v13, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__63)), __ccgo_ts+4204, libc.VaList(bp+8, i))
		return 0
		goto _8
	_8:
		;
		i++
	}
	return 0
}

var __func__63 = [20]uint8{'s', 'u', 'b', 'm', 'i', 't', '_', 'i', 's', 'o', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func submit_control_transfer(tls *libc.TLS, itransfer uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var hpriv, tpriv, transfer, urb, v1 uintptr
	var r int32
	_, _, _, _, _, _ = hpriv, r, tpriv, transfer, urb, v1
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	hpriv = usbi_get_device_handle_priv(tls, (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)
	if libc.Uint64FromInt32((*libusb_transfer)(unsafe.Pointer(transfer)).Flength)-libc.Uint64FromInt64(8) > uint64(4096) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	urb = libc.Xcalloc(tls, uint64(1), uint64(56))
	if !(urb != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = urb
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs = int32(1)
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(NORMAL)
	(*usbfs_urb)(unsafe.Pointer(urb)).Fusercontext = itransfer
	(*usbfs_urb)(unsafe.Pointer(urb)).Ftype1 = uint8(2)
	(*usbfs_urb)(unsafe.Pointer(urb)).Fendpoint = (*libusb_transfer)(unsafe.Pointer(transfer)).Fendpoint
	(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer = (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer
	(*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length = (*libusb_transfer)(unsafe.Pointer(transfer)).Flength
	r = libc.Xioctl(tls, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(2)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(10)))|libc.Uint64FromInt64(56)<<libc.Int32FromInt32(16)), libc.VaList(bp+8, urb))
	if r < 0 {
		libc.Xfree(tls, urb)
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = libc.UintptrFromInt32(0)
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__64)), __ccgo_ts+4148, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_IO)
	}
	return 0
}

var __func__64 = [24]uint8{'s', 'u', 'b', 'm', 'i', 't', '_', 'c', 'o', 'n', 't', 'r', 'o', 'l', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func op_submit_transfer(tls *libc.TLS, itransfer uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var transfer, v1 uintptr
	_, _ = transfer, v1
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	switch libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) {
	case int32(LIBUSB_TRANSFER_TYPE_CONTROL):
		return submit_control_transfer(tls, itransfer)
	case int32(LIBUSB_TRANSFER_TYPE_BULK):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_BULK_STREAM):
		return submit_bulk_transfer(tls, itransfer)
	case int32(LIBUSB_TRANSFER_TYPE_INTERRUPT):
		return submit_bulk_transfer(tls, itransfer)
	case int32(LIBUSB_TRANSFER_TYPE_ISOCHRONOUS):
		return submit_iso_transfer(tls, itransfer)
	default:
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__65)), __ccgo_ts+4428, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1)))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	return r
}

var __func__65 = [19]uint8{'o', 'p', '_', 's', 'u', 'b', 'm', 'i', 't', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func op_cancel_transfer(tls *libc.TLS, itransfer uintptr) (r1 int32) {
	var r int32
	var tpriv, transfer uintptr
	_, _, _ = r, tpriv, transfer
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	if !((*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs != 0) {
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	r = discard_urbs(tls, itransfer, 0, (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs)
	if r != 0 {
		return r
	}
	switch libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) {
	case int32(LIBUSB_TRANSFER_TYPE_BULK):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_BULK_STREAM):
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(ERROR) {
			break
		}
		fallthrough
	default:
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(CANCELLED)
	}
	return 0
}

func op_clear_transfer_priv(tls *libc.TLS, itransfer uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var tpriv, transfer, v1 uintptr
	_, _, _ = tpriv, transfer, v1
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	switch libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) {
	case int32(LIBUSB_TRANSFER_TYPE_CONTROL):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_BULK):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_BULK_STREAM):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_INTERRUPT):
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs != 0 {
			libc.Xfree(tls, (*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs)
			(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = libc.UintptrFromInt32(0)
		}
	case int32(LIBUSB_TRANSFER_TYPE_ISOCHRONOUS):
		if *(*uintptr)(unsafe.Pointer(tpriv)) != 0 {
			free_iso_urbs(tls, tpriv)
			*(*uintptr)(unsafe.Pointer(tpriv)) = libc.UintptrFromInt32(0)
		}
	default:
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__66)), __ccgo_ts+4428, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1)))
	}
}

var __func__66 = [23]uint8{'o', 'p', '_', 'c', 'l', 'e', 'a', 'r', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r', '_', 'p', 'r', 'i', 'v'}

func handle_bulk_completion(tls *libc.TLS, itransfer uintptr, urb uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var target, tpriv, transfer, v1, v10, v11, v12, v2, v3, v4, v5, v6, v7, v8, v9 uintptr
	var urb_idx, v13 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = target, tpriv, transfer, urb_idx, v1, v10, v11, v12, v13, v2, v3, v4, v5, v6, v7, v8, v9
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	urb_idx = int32((int64(urb) - int64((*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs)) / 56)
	usbi_mutex_lock(tls, itransfer)
	if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4453, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus, urb_idx+int32(1), (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs))
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired++
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action != int32(NORMAL) {
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v2 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v2 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4501, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		if (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length > 0 {
			target = (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer + uintptr((*usbi_transfer)(unsafe.Pointer(itransfer)).Ftransferred)
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v3 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v3 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4530, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length))
			if (*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer != target {
				if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
					v4 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
				} else {
					v4 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v4, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4564, libc.VaList(bp+8, int64((*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer)-int64((*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer), int64(target)-int64((*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer)))
				libc.Xmemmove(tls, target, (*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer, libc.Uint64FromInt32((*usbfs_urb)(unsafe.Pointer(urb)).Factual_length))
			}
			*(*int32)(unsafe.Pointer(itransfer + 96)) += (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length
		}
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired == (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v5 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v5 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v5, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4614, 0)
			if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action != int32(COMPLETED_EARLY) && (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status == int32(LIBUSB_TRANSFER_COMPLETED) {
				(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status = int32(LIBUSB_TRANSFER_ERROR)
			}
			goto completed
		}
		goto out_unlock
	}
	*(*int32)(unsafe.Pointer(itransfer + 96)) += (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length
	switch (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus {
	case 0:
	case -int32(121):
	case -int32(2):
		fallthrough
	case -int32(104):
	case -int32(19):
		fallthrough
	case -int32(108):
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v6 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v6 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v6, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4657, 0)
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status = int32(LIBUSB_TRANSFER_NO_DEVICE)
		goto cancel_remaining
	case -int32(32):
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v7 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v7 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v7, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4672, 0)
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status == int32(LIBUSB_TRANSFER_COMPLETED) {
			(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status = int32(LIBUSB_TRANSFER_STALL)
		}
		goto cancel_remaining
	case -int32(75):
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v8 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v8 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v8, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4696, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length))
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status == int32(LIBUSB_TRANSFER_COMPLETED) {
			(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status = int32(LIBUSB_TRANSFER_OVERFLOW)
		}
		goto completed
	case -int32(62):
		fallthrough
	case -int32(71):
		fallthrough
	case -int32(84):
		fallthrough
	case -int32(70):
		fallthrough
	case -int32(63):
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v9 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v9 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v9, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4723, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(ERROR)
		goto cancel_remaining
	default:
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
			v10 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
		} else {
			v10 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v10, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4746, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(ERROR)
		goto cancel_remaining
	}
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired == (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs {
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v11 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v11 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v11, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4773, 0)
		goto completed
	} else {
		if (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length < (*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v12 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v12 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v12, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__67)), __ccgo_ts+4815, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length, (*usbfs_urb)(unsafe.Pointer(urb)).Fbuffer_length))
			if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(NORMAL) {
				(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action = int32(COMPLETED_EARLY)
			}
		} else {
			goto out_unlock
		}
	}
	goto cancel_remaining
cancel_remaining:
	;
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(ERROR) && (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status == int32(LIBUSB_TRANSFER_COMPLETED) {
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status = int32(LIBUSB_TRANSFER_ERROR)
	}
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired == (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs {
		goto completed
	}
	discard_urbs(tls, itransfer, urb_idx+int32(1), (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs)
	goto out_unlock
out_unlock:
	;
	usbi_mutex_unlock(tls, itransfer)
	return 0
	goto completed
completed:
	;
	libc.Xfree(tls, (*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs)
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = libc.UintptrFromInt32(0)
	usbi_mutex_unlock(tls, itransfer)
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(CANCELLED) {
		v13 = usbi_handle_transfer_cancellation(tls, itransfer)
	} else {
		v13 = usbi_handle_transfer_completion(tls, itransfer, (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_status)
	}
	return v13
}

var __func__67 = [23]uint8{'h', 'a', 'n', 'd', 'l', 'e', '_', 'b', 'u', 'l', 'k', '_', 'c', 'o', 'm', 'p', 'l', 'e', 't', 'i', 'o', 'n'}

func handle_iso_completion(tls *libc.TLS, itransfer uintptr, urb uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, num_urbs, urb_idx, v5 int32
	var lib_desc, tpriv, transfer, urb_desc, v10, v11, v12, v13, v14, v15, v16, v2, v3, v6, v7, v8, v9 uintptr
	var status libusb_transfer_status
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, lib_desc, num_urbs, status, tpriv, transfer, urb_desc, urb_idx, v10, v11, v12, v13, v14, v15, v16, v2, v3, v5, v6, v7, v8, v9
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	num_urbs = (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_urbs
	urb_idx = 0
	status = int32(LIBUSB_TRANSFER_COMPLETED)
	usbi_mutex_lock(tls, itransfer)
	i = 0
	for {
		if !(i < num_urbs) {
			break
		}
		if urb == *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(tpriv)) + uintptr(i)*8)) {
			urb_idx = i + int32(1)
			break
		}
		goto _1
	_1:
		;
		i++
	}
	if urb_idx == 0 {
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v2 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v2 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4850, 0)
		usbi_mutex_unlock(tls, itransfer)
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
		v3 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
	} else {
		v3 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4872, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus, urb_idx, num_urbs))
	i = 0
	for {
		if !(i < (*usbfs_urb)(unsafe.Pointer(urb)).F__ccgo8_36.Fnumber_of_packets) {
			break
		}
		urb_desc = urb + 56 + uintptr(i)*12
		v6 = tpriv + 24
		v5 = *(*int32)(unsafe.Pointer(v6))
		*(*int32)(unsafe.Pointer(v6))++
		lib_desc = transfer + 60 + uintptr(v5)*12
		(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Fstatus = int32(LIBUSB_TRANSFER_COMPLETED)
		switch (*usbfs_iso_packet_desc)(unsafe.Pointer(urb_desc)).Fstatus {
		case uint32(0):
		case libc.Uint32FromInt32(-libc.Int32FromInt32(2)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(104)):
		case libc.Uint32FromInt32(-libc.Int32FromInt32(19)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(108)):
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v7 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v7 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v7, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4919, libc.VaList(bp+8, i))
			(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Fstatus = int32(LIBUSB_TRANSFER_NO_DEVICE)
		case libc.Uint32FromInt32(-libc.Int32FromInt32(32)):
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v8 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v8 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v8, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4946, libc.VaList(bp+8, i))
			(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Fstatus = int32(LIBUSB_TRANSFER_STALL)
		case libc.Uint32FromInt32(-libc.Int32FromInt32(75)):
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v9 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v9 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v9, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4982, libc.VaList(bp+8, i))
			(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Fstatus = int32(LIBUSB_TRANSFER_OVERFLOW)
		case libc.Uint32FromInt32(-libc.Int32FromInt32(62)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(71)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(84)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(70)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(63)):
			fallthrough
		case libc.Uint32FromInt32(-libc.Int32FromInt32(18)):
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v10 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v10 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v10, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+5009, libc.VaList(bp+8, i, (*usbfs_iso_packet_desc)(unsafe.Pointer(urb_desc)).Fstatus))
			(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Fstatus = int32(LIBUSB_TRANSFER_ERROR)
		default:
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v11 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v11 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v11, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+5044, libc.VaList(bp+8, i, (*usbfs_iso_packet_desc)(unsafe.Pointer(urb_desc)).Fstatus))
			(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Fstatus = int32(LIBUSB_TRANSFER_ERROR)
			break
		}
		(*libusb_iso_packet_descriptor)(unsafe.Pointer(lib_desc)).Factual_length = (*usbfs_iso_packet_desc)(unsafe.Pointer(urb_desc)).Factual_length
		goto _4
	_4:
		;
		i++
	}
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired++
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action != int32(NORMAL) {
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v12 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v12 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v12, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+5083, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired == num_urbs {
			if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
				v13 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
			} else {
				v13 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v13, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+5105, 0)
			free_iso_urbs(tls, tpriv)
			if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(CANCELLED) {
				usbi_mutex_unlock(tls, itransfer)
				return usbi_handle_transfer_cancellation(tls, itransfer)
			} else {
				usbi_mutex_unlock(tls, itransfer)
				return usbi_handle_transfer_completion(tls, itransfer, int32(LIBUSB_TRANSFER_ERROR))
			}
		}
		goto out
	}
	switch (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus {
	case 0:
	case -int32(2):
		fallthrough
	case -int32(104):
	case -int32(108):
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v14 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v14 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v14, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4657, 0)
		status = int32(LIBUSB_TRANSFER_NO_DEVICE)
	default:
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v15 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v15 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v15, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4746, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		status = int32(LIBUSB_TRANSFER_ERROR)
		break
	}
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Fnum_retired == num_urbs {
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v16 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v16 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v16, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__68)), __ccgo_ts+4773, 0)
		free_iso_urbs(tls, tpriv)
		usbi_mutex_unlock(tls, itransfer)
		return usbi_handle_transfer_completion(tls, itransfer, status)
	}
	goto out
out:
	;
	usbi_mutex_unlock(tls, itransfer)
	return 0
}

var __func__68 = [22]uint8{'h', 'a', 'n', 'd', 'l', 'e', '_', 'i', 's', 'o', '_', 'c', 'o', 'm', 'p', 'l', 'e', 't', 'i', 'o', 'n'}

func handle_control_completion(tls *libc.TLS, itransfer uintptr, urb uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var status int32
	var tpriv, v1, v2, v3, v4, v5, v6, v7 uintptr
	_, _, _, _, _, _, _, _, _ = status, tpriv, v1, v2, v3, v4, v5, v6, v7
	tpriv = usbi_get_transfer_priv(tls, itransfer)
	usbi_mutex_lock(tls, itransfer)
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+5141, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
	*(*int32)(unsafe.Pointer(itransfer + 96)) += (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length
	if (*linux_transfer_priv)(unsafe.Pointer(tpriv)).Freap_action == int32(CANCELLED) {
		if (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus != 0 && (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus != -int32(2) {
			if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
				v2 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
			} else {
				v2 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+5171, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		}
		libc.Xfree(tls, (*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs)
		(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = libc.UintptrFromInt32(0)
		usbi_mutex_unlock(tls, itransfer)
		return usbi_handle_transfer_cancellation(tls, itransfer)
	}
	switch (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus {
	case 0:
		status = int32(LIBUSB_TRANSFER_COMPLETED)
	case -int32(2):
		status = int32(LIBUSB_TRANSFER_CANCELLED)
	case -int32(19):
		fallthrough
	case -int32(108):
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
			v3 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
		} else {
			v3 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+4657, 0)
		status = int32(LIBUSB_TRANSFER_NO_DEVICE)
	case -int32(32):
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
			v4 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
		} else {
			v4 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v4, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+5206, 0)
		status = int32(LIBUSB_TRANSFER_STALL)
	case -int32(75):
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
			v5 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
		} else {
			v5 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v5, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+4696, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Factual_length))
		status = int32(LIBUSB_TRANSFER_OVERFLOW)
	case -int32(62):
		fallthrough
	case -int32(71):
		fallthrough
	case -int32(84):
		fallthrough
	case -int32(70):
		fallthrough
	case -int32(63):
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
			v6 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
		} else {
			v6 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v6, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+4723, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		status = int32(LIBUSB_TRANSFER_ERROR)
	default:
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
			v7 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
		} else {
			v7 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v7, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__69)), __ccgo_ts+4746, libc.VaList(bp+8, (*usbfs_urb)(unsafe.Pointer(urb)).Fstatus))
		status = int32(LIBUSB_TRANSFER_ERROR)
		break
	}
	libc.Xfree(tls, (*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs)
	(*linux_transfer_priv)(unsafe.Pointer(tpriv)).F__ccgo0_0.Furbs = libc.UintptrFromInt32(0)
	usbi_mutex_unlock(tls, itransfer)
	return usbi_handle_transfer_completion(tls, itransfer, status)
}

var __func__69 = [26]uint8{'h', 'a', 'n', 'd', 'l', 'e', '_', 'c', 'o', 'n', 't', 'r', 'o', 'l', '_', 'c', 'o', 'm', 'p', 'l', 'e', 't', 'i', 'o', 'n'}

func reap_for_handle(tls *libc.TLS, handle uintptr) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var hpriv, itransfer, transfer, v1, v2, v3 uintptr
	var r int32
	var _ /* urb at bp+0 */ uintptr
	_, _, _, _, _, _, _ = hpriv, itransfer, r, transfer, v1, v2, v3
	hpriv = usbi_get_device_handle_priv(tls, handle)
	*(*uintptr)(unsafe.Pointer(bp)) = libc.UintptrFromInt32(0)
	r = libc.Xioctl(tls, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd, libc.Int32FromUint64(uint64(libc.Uint32FromUint32(1)<<libc.Int32FromInt32(30)|libc.Uint32FromInt32(libc.Int32FromUint8('U')<<libc.Int32FromInt32(8))|libc.Uint32FromInt32(libc.Int32FromInt32(13)))|libc.Uint64FromInt64(8)<<libc.Int32FromInt32(16)), libc.VaList(bp+16, bp))
	if r < 0 {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(11) {
			return int32(1)
		}
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(19) {
			return int32(LIBUSB_ERROR_NO_DEVICE)
		}
		if handle != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__70)), __ccgo_ts+5234, libc.VaList(bp+16, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		return int32(LIBUSB_ERROR_IO)
	}
	itransfer = (*usbfs_urb)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fusercontext
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	if handle != 0 {
		v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
	} else {
		v2 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__70)), __ccgo_ts+5256, libc.VaList(bp+16, libc.Int32FromUint8((*usbfs_urb)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Ftype1), (*usbfs_urb)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Fstatus, (*usbfs_urb)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Factual_length))
	switch libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) {
	case int32(LIBUSB_TRANSFER_TYPE_ISOCHRONOUS):
		return handle_iso_completion(tls, itransfer, *(*uintptr)(unsafe.Pointer(bp)))
	case int32(LIBUSB_TRANSFER_TYPE_BULK):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_BULK_STREAM):
		fallthrough
	case int32(LIBUSB_TRANSFER_TYPE_INTERRUPT):
		return handle_bulk_completion(tls, itransfer, *(*uintptr)(unsafe.Pointer(bp)))
	case int32(LIBUSB_TRANSFER_TYPE_CONTROL):
		return handle_control_completion(tls, itransfer, *(*uintptr)(unsafe.Pointer(bp)))
	default:
		if handle != 0 {
			v3 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
		} else {
			v3 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__70)), __ccgo_ts+5293, libc.VaList(bp+16, libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1)))
		return int32(LIBUSB_ERROR_OTHER)
	}
	return r1
}

var __func__70 = [16]uint8{'r', 'e', 'a', 'p', '_', 'f', 'o', 'r', '_', 'h', 'a', 'n', 'd', 'l', 'e'}

func op_handle_events(tls *libc.TLS, ctx uintptr, event_data uintptr, count uint32, num_ready uint32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, fds, handle, hpriv, pollfd1, v3 uintptr
	var n uint32
	var r, reap_count, v5 int32
	var v4 usbi_atomic_t
	var v6 bool
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _, _, _, _, _, _, _, _, _, _ = __atomic_load_ptr, fds, handle, hpriv, n, pollfd1, r, reap_count, v3, v4, v5, v6
	fds = event_data
	usbi_mutex_lock(tls, ctx+80)
	n = uint32(0)
	for {
		if !(n < count && num_ready > uint32(0)) {
			break
		}
		pollfd1 = fds + uintptr(n)*8
		hpriv = libc.UintptrFromInt32(0)
		if !((*pollfd)(unsafe.Pointer(pollfd1)).Frevents != 0) {
			goto _1
		}
		num_ready--
		handle = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+120)).Fnext) - uint64(libc.UintptrFromInt32(0)+48))
		for {
			if !(handle+48 != ctx+120) {
				break
			}
			hpriv = usbi_get_device_handle_priv(tls, handle)
			if (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd == (*pollfd)(unsafe.Pointer(pollfd1)).Ffd {
				break
			}
			goto _2
		_2:
			;
			handle = uintptr(uint64((*libusb_device_handle)(unsafe.Pointer(handle)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
		}
		if !(hpriv != 0) || (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd != (*pollfd)(unsafe.Pointer(pollfd1)).Ffd {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__71)), __ccgo_ts+5323, libc.VaList(bp+16, (*pollfd)(unsafe.Pointer(pollfd1)).Ffd))
			goto _1
		}
		if int32((*pollfd)(unsafe.Pointer(pollfd1)).Frevents)&int32(0x008) != 0 {
			if handle != 0 {
				v3 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fctx
			} else {
				v3 = libc.UintptrFromInt32(0)
			}
			usbi_remove_event_source(tls, v3, (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd)
			(*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Ffd_removed = int32(1)
			usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
			{
				__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(handle)).Fdev + 80
				libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
				v4 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
			}
			if v4 != 0 {
				linux_device_disconnected(tls, (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fbus_number, (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(handle)).Fdev)).Fdevice_address)
			}
			usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
			if (*linux_device_handle_priv)(unsafe.Pointer(hpriv)).Fcaps&uint32(0x10) != 0 {
				for cond := true; cond; cond = r == 0 {
					r = reap_for_handle(tls, handle)
				}
			}
			usbi_handle_disconnect(tls, ctx, handle)
			goto _1
		}
		reap_count = 0
		for {
			r = reap_for_handle(tls, handle)
			goto _7
		_7:
			;
			if v6 = r == 0; v6 {
				reap_count++
				v5 = reap_count
			}
			if !(v6 && v5 <= int32(25)) {
				break
			}
		}
		if r == int32(1) || r == int32(LIBUSB_ERROR_NO_DEVICE) {
			goto _1
		} else {
			if r < 0 {
				goto out
			}
		}
		goto _1
	_1:
		;
		n++
	}
	r = 0
	goto out
out:
	;
	usbi_mutex_unlock(tls, ctx+80)
	return r
}

var __func__71 = [17]uint8{'o', 'p', '_', 'h', 'a', 'n', 'd', 'l', 'e', '_', 'e', 'v', 'e', 'n', 't', 's'}

func init() {
	p := unsafe.Pointer(&usbi_backend)
	*(*uintptr)(unsafe.Add(p, 16)) = __ccgo_fp(op_init)
	*(*uintptr)(unsafe.Add(p, 24)) = __ccgo_fp(op_exit)
	*(*uintptr)(unsafe.Add(p, 32)) = __ccgo_fp(op_set_option)
	*(*uintptr)(unsafe.Add(p, 48)) = __ccgo_fp(op_get_device_string)
	*(*uintptr)(unsafe.Add(p, 56)) = __ccgo_fp(op_hotplug_poll)
	*(*uintptr)(unsafe.Add(p, 64)) = __ccgo_fp(op_wrap_sys_device)
	*(*uintptr)(unsafe.Add(p, 72)) = __ccgo_fp(op_open)
	*(*uintptr)(unsafe.Add(p, 80)) = __ccgo_fp(op_close)
	*(*uintptr)(unsafe.Add(p, 88)) = __ccgo_fp(op_get_active_config_descriptor)
	*(*uintptr)(unsafe.Add(p, 96)) = __ccgo_fp(op_get_config_descriptor)
	*(*uintptr)(unsafe.Add(p, 104)) = __ccgo_fp(op_get_config_descriptor_by_value)
	*(*uintptr)(unsafe.Add(p, 112)) = __ccgo_fp(op_get_configuration)
	*(*uintptr)(unsafe.Add(p, 120)) = __ccgo_fp(op_set_configuration)
	*(*uintptr)(unsafe.Add(p, 128)) = __ccgo_fp(op_claim_interface)
	*(*uintptr)(unsafe.Add(p, 136)) = __ccgo_fp(op_release_interface)
	*(*uintptr)(unsafe.Add(p, 144)) = __ccgo_fp(op_set_interface)
	*(*uintptr)(unsafe.Add(p, 152)) = __ccgo_fp(op_clear_halt)
	*(*uintptr)(unsafe.Add(p, 160)) = __ccgo_fp(op_reset_device)
	*(*uintptr)(unsafe.Add(p, 168)) = __ccgo_fp(op_alloc_streams)
	*(*uintptr)(unsafe.Add(p, 176)) = __ccgo_fp(op_free_streams)
	*(*uintptr)(unsafe.Add(p, 184)) = __ccgo_fp(op_dev_mem_alloc)
	*(*uintptr)(unsafe.Add(p, 192)) = __ccgo_fp(op_dev_mem_free)
	*(*uintptr)(unsafe.Add(p, 200)) = __ccgo_fp(op_kernel_driver_active)
	*(*uintptr)(unsafe.Add(p, 208)) = __ccgo_fp(op_detach_kernel_driver)
	*(*uintptr)(unsafe.Add(p, 216)) = __ccgo_fp(op_attach_kernel_driver)
	*(*uintptr)(unsafe.Add(p, 248)) = __ccgo_fp(op_destroy_device)
	*(*uintptr)(unsafe.Add(p, 256)) = __ccgo_fp(op_submit_transfer)
	*(*uintptr)(unsafe.Add(p, 264)) = __ccgo_fp(op_cancel_transfer)
	*(*uintptr)(unsafe.Add(p, 272)) = __ccgo_fp(op_clear_transfer_priv)
	*(*uintptr)(unsafe.Add(p, 280)) = __ccgo_fp(op_handle_events)
}

type socklen_t = uint32

type sa_family_t = uint16

type msghdr = struct {
	Fmsg_name       uintptr
	Fmsg_namelen    socklen_t
	Fmsg_iov        uintptr
	Fmsg_iovlen     int32
	F__pad1         int32
	Fmsg_control    uintptr
	Fmsg_controllen socklen_t
	F__pad2         int32
	Fmsg_flags      int32
}

type cmsghdr = struct {
	Fcmsg_len   socklen_t
	F__pad1     int32
	Fcmsg_level int32
	Fcmsg_type  int32
}

type ucred = struct {
	Fpid pid_t
	Fuid uid_t
	Fgid gid_t
}

type mmsghdr = struct {
	Fmsg_hdr msghdr
	Fmsg_len uint32
}

type linger = struct {
	Fl_onoff  int32
	Fl_linger int32
}

type sockaddr = struct {
	Fsa_family sa_family_t
	Fsa_data   [14]uint8
}

type sockaddr_storage = struct {
	Fss_family    sa_family_t
	F__ss_padding [118]uint8
	F__ss_align   uint64
}

type __kernel_sa_family_t = uint16

type __kernel_sockaddr_storage = struct {
	F__ccgo0_0 struct {
		F__align   [0]uintptr
		F__ccgo0_0 struct {
			Fss_family __kernel_sa_family_t
			F__data    [126]uint8
		}
	}
}

type sockaddr_nl = struct {
	Fnl_family __kernel_sa_family_t
	Fnl_pad    uint16
	Fnl_pid    __u32
	Fnl_groups __u32
}

type nlmsghdr = struct {
	Fnlmsg_len   __u32
	Fnlmsg_type  __u16
	Fnlmsg_flags __u16
	Fnlmsg_seq   __u32
	Fnlmsg_pid   __u32
}

type nlmsgerr = struct {
	Ferror1 int32
	Fmsg    nlmsghdr
}

type nlmsgerr_attrs = int32

const NLMSGERR_ATTR_UNUSED = 0
const NLMSGERR_ATTR_MSG = 1
const NLMSGERR_ATTR_OFFS = 2
const NLMSGERR_ATTR_COOKIE = 3
const NLMSGERR_ATTR_POLICY = 4
const NLMSGERR_ATTR_MISS_TYPE = 5
const NLMSGERR_ATTR_MISS_NEST = 6
const __NLMSGERR_ATTR_MAX = 7
const NLMSGERR_ATTR_MAX = 6

type nl_pktinfo = struct {
	Fgroup __u32
}

type nl_mmap_req = struct {
	Fnm_block_size uint32
	Fnm_block_nr   uint32
	Fnm_frame_size uint32
	Fnm_frame_nr   uint32
}

type nl_mmap_hdr = struct {
	Fnm_status uint32
	Fnm_len    uint32
	Fnm_group  __u32
	Fnm_pid    __u32
	Fnm_uid    __u32
	Fnm_gid    __u32
}

type nl_mmap_status = int32

const NL_MMAP_STATUS_UNUSED = 0
const NL_MMAP_STATUS_RESERVED = 1
const NL_MMAP_STATUS_VALID = 2
const NL_MMAP_STATUS_COPY = 3
const NL_MMAP_STATUS_SKIP = 4
const NETLINK_UNCONNECTED = 0
const NETLINK_CONNECTED = 1

type nlattr = struct {
	Fnla_len  __u16
	Fnla_type __u16
}

type nla_bitfield32 = struct {
	Fvalue    __u32
	Fselector __u32
}

type netlink_attribute_type = int32

const NL_ATTR_TYPE_INVALID = 0
const NL_ATTR_TYPE_FLAG = 1
const NL_ATTR_TYPE_U8 = 2
const NL_ATTR_TYPE_U16 = 3
const NL_ATTR_TYPE_U32 = 4
const NL_ATTR_TYPE_U64 = 5
const NL_ATTR_TYPE_S8 = 6
const NL_ATTR_TYPE_S16 = 7
const NL_ATTR_TYPE_S32 = 8
const NL_ATTR_TYPE_S64 = 9
const NL_ATTR_TYPE_BINARY = 10
const NL_ATTR_TYPE_STRING = 11
const NL_ATTR_TYPE_NUL_STRING = 12
const NL_ATTR_TYPE_NESTED = 13
const NL_ATTR_TYPE_NESTED_ARRAY = 14
const NL_ATTR_TYPE_BITFIELD32 = 15
const NL_ATTR_TYPE_SINT = 16
const NL_ATTR_TYPE_UINT = 17

type netlink_policy_type_attr = int32

const NL_POLICY_TYPE_ATTR_UNSPEC = 0
const NL_POLICY_TYPE_ATTR_TYPE = 1
const NL_POLICY_TYPE_ATTR_MIN_VALUE_S = 2
const NL_POLICY_TYPE_ATTR_MAX_VALUE_S = 3
const NL_POLICY_TYPE_ATTR_MIN_VALUE_U = 4
const NL_POLICY_TYPE_ATTR_MAX_VALUE_U = 5
const NL_POLICY_TYPE_ATTR_MIN_LENGTH = 6
const NL_POLICY_TYPE_ATTR_MAX_LENGTH = 7
const NL_POLICY_TYPE_ATTR_POLICY_IDX = 8
const NL_POLICY_TYPE_ATTR_POLICY_MAXTYPE = 9
const NL_POLICY_TYPE_ATTR_BITFIELD32_MASK = 10
const NL_POLICY_TYPE_ATTR_PAD = 11
const NL_POLICY_TYPE_ATTR_MASK = 12
const __NL_POLICY_TYPE_ATTR_MAX = 13
const NL_POLICY_TYPE_ATTR_MAX = 12

var linux_netlink_socket = -int32(1)
var netlink_control_event = usbi_event_t{
	Fpipefd: [2]int32{
		0: -int32(1),
		1: -int32(1),
	},
}
var libusb_linux_event_thread pthread_t

func set_fd_cloexec_nb(tls *libc.TLS, fd int32, socktype int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var flags int32
	_ = flags
	if !(socktype&libc.Int32FromInt32(02000000) != 0) {
		flags = libc.Xfcntl(tls, fd, int32(1), 0)
		if flags == -int32(1) {
			usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__72)), __ccgo_ts+5352, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			return -int32(1)
		}
		if libc.Xfcntl(tls, fd, int32(2), libc.VaList(bp+8, flags|int32(1))) == -int32(1) {
			usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__72)), __ccgo_ts+5393, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			return -int32(1)
		}
	}
	if !(socktype&libc.Int32FromInt32(04000) != 0) {
		flags = libc.Xfcntl(tls, fd, int32(3), 0)
		if flags == -int32(1) {
			usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__72)), __ccgo_ts+5434, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			return -int32(1)
		}
		if libc.Xfcntl(tls, fd, int32(4), libc.VaList(bp+8, flags|int32(04000))) == -int32(1) {
			usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__72)), __ccgo_ts+5482, libc.VaList(bp+8, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			return -int32(1)
		}
	}
	return 0
}

var __func__72 = [18]uint8{'s', 'e', 't', '_', 'f', 'd', '_', 'c', 'l', 'o', 'e', 'x', 'e', 'c', '_', 'n', 'b'}

func linux_netlink_start_event_monitor(tls *libc.TLS) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ret, socktype int32
	var _ /* opt at bp+12 */ int32
	var _ /* sa_nl at bp+0 */ sockaddr_nl
	_, _ = ret, socktype
	*(*sockaddr_nl)(unsafe.Pointer(bp)) = sockaddr_nl{
		Fnl_family: uint16(16),
		Fnl_groups: uint32(1),
	}
	socktype = libc.Int32FromInt32(3) | libc.Int32FromInt32(04000) | libc.Int32FromInt32(02000000)
	*(*int32)(unsafe.Pointer(bp + 12)) = int32(1)
	linux_netlink_socket = libc.Xsocket(tls, int32(16), socktype, int32(15))
	if linux_netlink_socket == -int32(1) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(22) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__73)), __ccgo_ts+5530, libc.VaList(bp+24, socktype))
		socktype = int32(3)
		linux_netlink_socket = libc.Xsocket(tls, int32(16), socktype, int32(15))
	}
	if linux_netlink_socket == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__73)), __ccgo_ts+5594, libc.VaList(bp+24, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		goto err
	}
	ret = set_fd_cloexec_nb(tls, linux_netlink_socket, socktype)
	if ret == -int32(1) {
		goto err_close_socket
	}
	ret = libc.Xbind(tls, linux_netlink_socket, bp, uint32(12))
	if ret == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__73)), __ccgo_ts+5636, libc.VaList(bp+24, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		goto err_close_socket
	}
	ret = libc.Xsetsockopt(tls, linux_netlink_socket, int32(1), int32(16), bp+12, uint32(4))
	if ret == -int32(1) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__73)), __ccgo_ts+5676, libc.VaList(bp+24, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		goto err_close_socket
	}
	ret = usbi_create_event(tls, uintptr(unsafe.Pointer(&netlink_control_event)))
	if ret != 0 {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__73)), __ccgo_ts+5734, 0)
		goto err_close_socket
	}
	ret = libc.Xpthread_create(tls, uintptr(unsafe.Pointer(&libusb_linux_event_thread)), libc.UintptrFromInt32(0), __ccgo_fp(linux_netlink_event_thread_main), libc.UintptrFromInt32(0))
	if ret != 0 {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__73)), __ccgo_ts+5773, libc.VaList(bp+24, ret))
		goto err_destroy_event
	}
	return int32(LIBUSB_SUCCESS)
	goto err_destroy_event
err_destroy_event:
	;
	usbi_destroy_event(tls, uintptr(unsafe.Pointer(&netlink_control_event)))
	netlink_control_event = usbi_event_t{
		Fpipefd: [2]int32{
			0: -int32(1),
			1: -int32(1),
		},
	}
	goto err_close_socket
err_close_socket:
	;
	libc.Xclose(tls, linux_netlink_socket)
	linux_netlink_socket = -int32(1)
	goto err
err:
	;
	return int32(LIBUSB_ERROR_OTHER)
	return r
}

var __func__73 = [34]uint8{'l', 'i', 'n', 'u', 'x', '_', 'n', 'e', 't', 'l', 'i', 'n', 'k', '_', 's', 't', 'a', 'r', 't', '_', 'e', 'v', 'e', 'n', 't', '_', 'm', 'o', 'n', 'i', 't', 'o', 'r'}

func linux_netlink_stop_event_monitor(tls *libc.TLS) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ret int32
	var v1 bool
	_, _ = ret, v1
	if v1 = linux_netlink_socket != -int32(1); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+5816, __ccgo_ts+5843, int32(156), uintptr(unsafe.Pointer(&__func__74)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	usbi_signal_event(tls, uintptr(unsafe.Pointer(&netlink_control_event)))
	ret = libc.Xpthread_join(tls, libusb_linux_event_thread, libc.UintptrFromInt32(0))
	if ret != 0 {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__74)), __ccgo_ts+5896, libc.VaList(bp+8, ret))
	}
	usbi_destroy_event(tls, uintptr(unsafe.Pointer(&netlink_control_event)))
	netlink_control_event = usbi_event_t{
		Fpipefd: [2]int32{
			0: -int32(1),
			1: -int32(1),
		},
	}
	libc.Xclose(tls, linux_netlink_socket)
	linux_netlink_socket = -int32(1)
	return int32(LIBUSB_SUCCESS)
}

var __func__74 = [33]uint8{'l', 'i', 'n', 'u', 'x', '_', 'n', 'e', 't', 'l', 'i', 'n', 'k', '_', 's', 't', 'o', 'p', '_', 'e', 'v', 'e', 'n', 't', '_', 'm', 'o', 'n', 'i', 't', 'o', 'r'}

func netlink_message_parse(tls *libc.TLS, buffer uintptr, len1 size_t, key uintptr) (r uintptr) {
	var end uintptr
	var keylen size_t
	_, _ = end, keylen
	end = buffer + uintptr(len1)
	keylen = libc.Xstrlen(tls, key)
	for buffer < end && *(*uint8)(unsafe.Pointer(buffer)) != 0 {
		if libc.Xstrncmp(tls, buffer, key, keylen) == 0 && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(buffer + uintptr(keylen)))) == int32('=') {
			return buffer + uintptr(keylen) + uintptr(1)
		}
		buffer += uintptr(libc.Xstrlen(tls, buffer) + uint64(1))
	}
	return libc.UintptrFromInt32(0)
}

func linux_netlink_parse(tls *libc.TLS, buffer uintptr, len1 size_t, detached uintptr, sys_name uintptr, busnum uintptr, devaddr uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var slash, tmp uintptr
	_, _ = slash, tmp
	*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
	*(*uintptr)(unsafe.Pointer(sys_name)) = libc.UintptrFromInt32(0)
	*(*int32)(unsafe.Pointer(detached)) = 0
	*(*uint8_t)(unsafe.Pointer(busnum)) = uint8(0)
	*(*uint8_t)(unsafe.Pointer(devaddr)) = uint8(0)
	tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+5937)
	if !(tmp != 0) {
		return -int32(1)
	} else {
		if libc.Xstrcmp(tls, tmp, __ccgo_ts+5944) == 0 {
			*(*int32)(unsafe.Pointer(detached)) = int32(1)
		} else {
			if libc.Xstrcmp(tls, tmp, __ccgo_ts+5951) != 0 {
				usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__75)), __ccgo_ts+5955, libc.VaList(bp+8, tmp))
				return -int32(1)
			}
		}
	}
	tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+5980)
	if !(tmp != 0) || libc.Xstrcmp(tls, tmp, __ccgo_ts+2835) != 0 {
		return -int32(1)
	}
	tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+5990)
	if !(tmp != 0) || libc.Xstrcmp(tls, tmp, __ccgo_ts+5998) != 0 {
		return -int32(1)
	}
	tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+6009)
	if tmp != 0 {
		*(*uint8_t)(unsafe.Pointer(busnum)) = uint8(libc.Xstrtoul(tls, tmp, libc.UintptrFromInt32(0), int32(10)) & libc.Uint64FromInt32(0xff))
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			return -int32(1)
		}
		tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+6016)
		if libc.UintptrFromInt32(0) == tmp {
			return -int32(1)
		}
		*(*uint8_t)(unsafe.Pointer(devaddr)) = uint8(libc.Xstrtoul(tls, tmp, libc.UintptrFromInt32(0), int32(10)) & libc.Uint64FromInt32(0xff))
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			return -int32(1)
		}
	} else {
		tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+6023)
		if !(tmp != 0) {
			return -int32(1)
		}
		slash = libc.Xstrrchr(tls, tmp, int32('/'))
		if !(slash != 0) {
			return -int32(1)
		}
		*(*uint8_t)(unsafe.Pointer(busnum)) = uint8(libc.Xstrtoul(tls, slash-uintptr(3), libc.UintptrFromInt32(0), int32(10)) & libc.Uint64FromInt32(0xff))
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			return -int32(1)
		}
		*(*uint8_t)(unsafe.Pointer(devaddr)) = uint8(libc.Xstrtoul(tls, slash+uintptr(1), libc.UintptrFromInt32(0), int32(10)) & libc.Uint64FromInt32(0xff))
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != 0 {
			*(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) = 0
			return -int32(1)
		}
		return 0
	}
	tmp = netlink_message_parse(tls, buffer, len1, __ccgo_ts+6030)
	if !(tmp != 0) {
		return -int32(1)
	}
	slash = libc.Xstrrchr(tls, tmp, int32('/'))
	if slash != 0 {
		*(*uintptr)(unsafe.Pointer(sys_name)) = slash + uintptr(1)
	}
	return 0
}

var __func__75 = [20]uint8{'l', 'i', 'n', 'u', 'x', '_', 'n', 'e', 't', 'l', 'i', 'n', 'k', '_', 'p', 'a', 'r', 's', 'e'}

func linux_netlink_read_message(tls *libc.TLS) (r1 int32) {
	bp := tls.Alloc(2224)
	defer tls.Free(2224)
	var cmsg, cred, v1, v2 uintptr
	var len1 ssize_t
	var r int32
	var _ /* busnum at bp+2088 */ uint8_t
	var _ /* cred_buffer at bp+0 */ [32]uint8
	var _ /* detached at bp+2092 */ int32
	var _ /* devaddr at bp+2089 */ uint8_t
	var _ /* iov at bp+2112 */ iovec
	var _ /* msg at bp+2128 */ msghdr
	var _ /* msg_buffer at bp+32 */ [2048]uint8
	var _ /* sa_nl at bp+2096 */ sockaddr_nl
	var _ /* sys_name at bp+2080 */ uintptr
	_, _, _, _, _, _ = cmsg, cred, len1, r, v1, v2
	*(*uintptr)(unsafe.Pointer(bp + 2080)) = libc.UintptrFromInt32(0)
	*(*iovec)(unsafe.Pointer(bp + 2112)) = iovec{
		Fiov_base: bp + 32,
		Fiov_len:  uint64(2048),
	}
	*(*msghdr)(unsafe.Pointer(bp + 2128)) = msghdr{
		Fmsg_name:       bp + 2096,
		Fmsg_namelen:    uint32(12),
		Fmsg_iov:        bp + 2112,
		Fmsg_iovlen:     int32(1),
		Fmsg_control:    bp,
		Fmsg_controllen: uint32(32),
	}
	len1 = libc.Xrecvmsg(tls, linux_netlink_socket, bp+2128, 0)
	if len1 == int64(-int32(1)) {
		if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(11) && *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) != int32(4) {
			usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__76)), __ccgo_ts+6038, libc.VaList(bp+2192, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
		}
		return -int32(1)
	}
	if len1 < int64(32) || (*(*msghdr)(unsafe.Pointer(bp + 2128))).Fmsg_flags&int32(0x0020) != 0 {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__76)), __ccgo_ts+6085, 0)
		return -int32(1)
	}
	if (*(*sockaddr_nl)(unsafe.Pointer(bp + 2096))).Fnl_groups != uint32(1) || (*(*sockaddr_nl)(unsafe.Pointer(bp + 2096))).Fnl_pid != uint32(0) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__76)), __ccgo_ts+6116, libc.VaList(bp+2192, (*(*sockaddr_nl)(unsafe.Pointer(bp + 2096))).Fnl_groups, (*(*sockaddr_nl)(unsafe.Pointer(bp + 2096))).Fnl_pid))
		return -int32(1)
	}
	if uint64((*msghdr)(unsafe.Pointer(bp+2128)).Fmsg_controllen) >= uint64(16) {
		v1 = (*msghdr)(unsafe.Pointer(bp + 2128)).Fmsg_control
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	cmsg = v1
	if !(cmsg != 0) || (*cmsghdr)(unsafe.Pointer(cmsg)).Fcmsg_type != int32(0x02) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__76)), __ccgo_ts+6172, 0)
		return -int32(1)
	}
	cred = cmsg + libc.UintptrFromInt32(1)*16
	if (*ucred)(unsafe.Pointer(cred)).Fuid != uint32(0) {
		usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__76)), __ccgo_ts+6224, libc.VaList(bp+2192, (*ucred)(unsafe.Pointer(cred)).Fuid))
		return -int32(1)
	}
	r = linux_netlink_parse(tls, bp+32, libc.Uint64FromInt64(len1), bp+2092, bp+2080, bp+2088, bp+2089)
	if r != 0 {
		return r
	}
	if *(*int32)(unsafe.Pointer(bp + 2092)) != 0 {
		v2 = __ccgo_ts + 6277
	} else {
		v2 = __ccgo_ts + 6281
	}
	usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__76)), __ccgo_ts+6284, libc.VaList(bp+2192, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(bp + 2088))), libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(bp + 2089))), *(*uintptr)(unsafe.Pointer(bp + 2080)), v2))
	if *(*int32)(unsafe.Pointer(bp + 2092)) != 0 {
		linux_device_disconnected(tls, *(*uint8_t)(unsafe.Pointer(bp + 2088)), *(*uint8_t)(unsafe.Pointer(bp + 2089)))
	} else {
		linux_hotplug_enumerate(tls, *(*uint8_t)(unsafe.Pointer(bp + 2088)), *(*uint8_t)(unsafe.Pointer(bp + 2089)), *(*uintptr)(unsafe.Pointer(bp + 2080)))
	}
	return 0
}

var __func__76 = [27]uint8{'l', 'i', 'n', 'u', 'x', '_', 'n', 'e', 't', 'l', 'i', 'n', 'k', '_', 'r', 'e', 'a', 'd', '_', 'm', 'e', 's', 's', 'a', 'g', 'e'}

func linux_netlink_event_thread_main(tls *libc.TLS, arg uintptr) (r1 uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var r int32
	var _ /* fds at bp+0 */ [2]pollfd
	_ = r
	*(*[2]pollfd)(unsafe.Pointer(bp)) = [2]pollfd{
		0: {
			Ffd:     *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&netlink_control_event)))),
			Fevents: int16(0x001),
		},
		1: {
			Ffd:     linux_netlink_socket,
			Fevents: int16(0x001),
		},
	}
	_ = arg
	usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__77)), __ccgo_ts+6368, 0)
	for int32(1) != 0 {
		r = libc.Xpoll(tls, bp, uint64(2), -int32(1))
		if r == -int32(1) {
			if *(*int32)(unsafe.Pointer(libc.X__errno_location(tls))) == int32(4) {
				continue
			}
			usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__77)), __ccgo_ts+893, libc.VaList(bp+24, *(*int32)(unsafe.Pointer(libc.X__errno_location(tls)))))
			break
		}
		if (*(*[2]pollfd)(unsafe.Pointer(bp)))[0].Frevents != 0 {
			break
		}
		if (*(*[2]pollfd)(unsafe.Pointer(bp)))[int32(1)].Frevents != 0 {
			usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
			linux_netlink_read_message(tls)
			usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
		}
	}
	usbi_log(tls, libc.UintptrFromInt32(0), int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__77)), __ccgo_ts+6398, 0)
	return libc.UintptrFromInt32(0)
}

var __func__77 = [32]uint8{'l', 'i', 'n', 'u', 'x', '_', 'n', 'e', 't', 'l', 'i', 'n', 'k', '_', 'e', 'v', 'e', 'n', 't', '_', 't', 'h', 'r', 'e', 'a', 'd', '_', 'm', 'a', 'i', 'n'}

func linux_netlink_hotplug_poll(tls *libc.TLS) {
	var r int32
	_ = r
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
	for cond := true; cond; cond = r == 0 {
		r = linux_netlink_read_message(tls)
	}
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&linux_hotplug_lock)))
}

var libusb_version_internal = libusb_version{
	Fmajor:    uint16(1),
	Fmicro:    uint16(30),
	Fnano:     uint16(12017),
	Frc:       __ccgo_ts + 6427,
	Fdescribe: __ccgo_ts + 6432,
}
var timestamp_origin timespec
var default_context_refcnt int32
var default_debug_level = int64(-int32(1))
var default_context_lock = usbi_mutex_static_t{}
var default_context_options [4]usbi_option

func discovered_devs_alloc(tls *libc.TLS) (r uintptr) {
	var ret uintptr
	_ = ret
	ret = libc.Xmalloc(tls, libc.Uint64FromInt64(16)+libc.Uint64FromInt64(8)*libc.Uint64FromInt32(16))
	if ret != 0 {
		(*discovered_devs)(unsafe.Pointer(ret)).Flen1 = uint64(0)
		(*discovered_devs)(unsafe.Pointer(ret)).Fcapacity = uint64(16)
	}
	return ret
}

func discovered_devs_free(tls *libc.TLS, discdevs uintptr) {
	var i size_t
	_ = i
	i = uint64(0)
	for {
		if !(i < (*discovered_devs)(unsafe.Pointer(discdevs)).Flen1) {
			break
		}
		libusb_unref_device(tls, *(*uintptr)(unsafe.Pointer(discdevs + 16 + uintptr(i)*8)))
		goto _1
	_1:
		;
		i++
	}
	libc.Xfree(tls, discdevs)
}

func discovered_devs_append(tls *libc.TLS, discdevs uintptr, dev uintptr) (r uintptr) {
	var capacity, len1 size_t
	var new_discdevs uintptr
	_, _, _ = capacity, len1, new_discdevs
	len1 = (*discovered_devs)(unsafe.Pointer(discdevs)).Flen1
	if len1 < (*discovered_devs)(unsafe.Pointer(discdevs)).Fcapacity {
		*(*uintptr)(unsafe.Pointer(discdevs + 16 + uintptr(len1)*8)) = libusb_ref_device(tls, dev)
		(*discovered_devs)(unsafe.Pointer(discdevs)).Flen1++
		return discdevs
	}
	usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__78)), __ccgo_ts+6452, 0)
	capacity = (*discovered_devs)(unsafe.Pointer(discdevs)).Fcapacity + uint64(16)
	new_discdevs = libc.Xrealloc(tls, discdevs, uint64(16)+uint64(8)*capacity)
	if !(new_discdevs != 0) {
		discovered_devs_free(tls, discdevs)
		return libc.UintptrFromInt32(0)
	}
	discdevs = new_discdevs
	(*discovered_devs)(unsafe.Pointer(discdevs)).Fcapacity = capacity
	*(*uintptr)(unsafe.Pointer(discdevs + 16 + uintptr(len1)*8)) = libusb_ref_device(tls, dev)
	(*discovered_devs)(unsafe.Pointer(discdevs)).Flen1++
	return discdevs
}

var __func__78 = [23]uint8{'d', 'i', 's', 'c', 'o', 'v', 'e', 'r', 'e', 'd', '_', 'd', 'e', 'v', 's', '_', 'a', 'p', 'p', 'e', 'n', 'd'}

func usbi_alloc_device(tls *libc.TLS, ctx uintptr, session_id uint64) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_store_ptr, dev uintptr
	var priv_size size_t
	var _ /* __atomic_store_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_store_ptr, dev, priv_size
	priv_size = usbi_backend.Fdevice_priv_size
	dev = libc.Xcalloc(tls, uint64(1), (libc.Uint64FromInt64(112)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)) + priv_size)
	if !(dev != 0) {
		return libc.UintptrFromInt32(0)
	}
	{
		__atomic_store_ptr = dev
		*(*usbi_atomic_t)(unsafe.Pointer(bp)) = int64(libc.Int32FromInt32(1))
		libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp, libc.Int32FromInt32(5))
	}
	(*libusb_device)(unsafe.Pointer(dev)).Fctx = ctx
	(*libusb_device)(unsafe.Pointer(dev)).Fsession_data = session_id
	(*libusb_device)(unsafe.Pointer(dev)).Fspeed = int32(LIBUSB_SPEED_UNKNOWN)
	if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
		usbi_connect_device(tls, dev)
	}
	return dev
}

func usbi_attach_device(tls *libc.TLS, dev uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_store_ptr, ctx uintptr
	var _ /* __atomic_store_tmp at bp+0 */ usbi_atomic_t
	_, _ = __atomic_store_ptr, ctx
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	{
		__atomic_store_ptr = dev + 80
		*(*usbi_atomic_t)(unsafe.Pointer(bp)) = int64(libc.Int32FromInt32(1))
		libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp, libc.Int32FromInt32(5))
	}
	usbi_mutex_lock(tls, ctx+24)
	list_add(tls, dev+32, ctx+64)
	usbi_mutex_unlock(tls, ctx+24)
}

func usbi_connect_device(tls *libc.TLS, dev uintptr) {
	usbi_attach_device(tls, dev)
	usbi_hotplug_notification(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, dev, int32(LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED))
}

func usbi_detach_device(tls *libc.TLS, dev uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_store_ptr, ctx uintptr
	var _ /* __atomic_store_tmp at bp+0 */ usbi_atomic_t
	_, _ = __atomic_store_ptr, ctx
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	{
		__atomic_store_ptr = dev + 80
		*(*usbi_atomic_t)(unsafe.Pointer(bp)) = int64(libc.Int32FromInt32(0))
		libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp, libc.Int32FromInt32(5))
	}
	usbi_mutex_lock(tls, ctx+24)
	list_del(tls, dev+32)
	usbi_mutex_unlock(tls, ctx+24)
}

func usbi_disconnect_device(tls *libc.TLS, dev uintptr) {
	usbi_detach_device(tls, dev)
	usbi_hotplug_notification(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, dev, int32(LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT))
}

func usbi_sanitize_device(tls *libc.TLS, dev uintptr) (r int32) {
	var num_configurations uint8_t
	_ = num_configurations
	if libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbLength) != int32(18) || libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbDescriptorType) != int32(LIBUSB_DT_DEVICE) {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__79)), __ccgo_ts+6478, 0)
		return int32(LIBUSB_ERROR_IO)
	}
	num_configurations = (*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations
	if libc.Int32FromUint8(num_configurations) > int32(8) {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__79)), __ccgo_ts+6504, 0)
		return int32(LIBUSB_ERROR_IO)
	} else {
		if 0 == libc.Int32FromUint8(num_configurations) {
			usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__79)), __ccgo_ts+6528, 0)
		}
	}
	return 0
}

var __func__79 = [21]uint8{'u', 's', 'b', 'i', '_', 's', 'a', 'n', 'i', 't', 'i', 'z', 'e', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func usbi_get_device_by_session_id(tls *libc.TLS, ctx uintptr, session_id uint64) (r uintptr) {
	var dev, ret uintptr
	_, _ = dev, ret
	ret = libc.UintptrFromInt32(0)
	usbi_mutex_lock(tls, ctx+24)
	dev = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+64)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(dev+32 != ctx+64) {
			break
		}
		if (*libusb_device)(unsafe.Pointer(dev)).Fsession_data == session_id {
			ret = libusb_ref_device(tls, dev)
			break
		}
		goto _1
	_1:
		;
		dev = uintptr(uint64((*libusb_device)(unsafe.Pointer(dev)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	usbi_mutex_unlock(tls, ctx+24)
	return ret
}

func libusb_get_device_list(tls *libc.TLS, ctx uintptr, list uintptr) (r1 ssize_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var dev, dev1, ret uintptr
	var i, len1 ssize_t
	var r int32
	var _ /* discdevs at bp+0 */ uintptr
	_, _, _, _, _, _ = dev, dev1, i, len1, r, ret
	*(*uintptr)(unsafe.Pointer(bp)) = discovered_devs_alloc(tls)
	r = 0
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__80)), __ccgo_ts+6578, 0)
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
		return int64(LIBUSB_ERROR_NO_MEM)
	}
	ctx = usbi_get_context(tls, ctx)
	if libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0 {
		if usbi_backend.Fhotplug_poll != 0 {
			(*(*func(*libc.TLS))(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fhotplug_poll})))(tls)
		}
		usbi_mutex_lock(tls, ctx+24)
		dev = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+64)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
		for {
			if !(dev+32 != ctx+64) {
				break
			}
			*(*uintptr)(unsafe.Pointer(bp)) = discovered_devs_append(tls, *(*uintptr)(unsafe.Pointer(bp)), dev)
			if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
				r = int32(LIBUSB_ERROR_NO_MEM)
				break
			}
			goto _1
		_1:
			;
			dev = uintptr(uint64((*libusb_device)(unsafe.Pointer(dev)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
		}
		usbi_mutex_unlock(tls, ctx+24)
	} else {
		r = (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_device_list})))(tls, ctx, bp)
	}
	if r < 0 {
		len1 = int64(r)
		goto out
	}
	len1 = libc.Int64FromUint64((*discovered_devs)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Flen1)
	ret = libc.Xcalloc(tls, libc.Uint64FromInt64(len1)+uint64(1), uint64(8))
	if !(ret != 0) {
		len1 = int64(LIBUSB_ERROR_NO_MEM)
		goto out
	}
	*(*uintptr)(unsafe.Pointer(ret + uintptr(len1)*8)) = libc.UintptrFromInt32(0)
	i = 0
	for {
		if !(i < len1) {
			break
		}
		dev1 = *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + 16 + uintptr(i)*8))
		*(*uintptr)(unsafe.Pointer(ret + uintptr(i)*8)) = libusb_ref_device(tls, dev1)
		goto _2
	_2:
		;
		i++
	}
	*(*uintptr)(unsafe.Pointer(list)) = ret
	goto out
out:
	;
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		discovered_devs_free(tls, *(*uintptr)(unsafe.Pointer(bp)))
	}
	return len1
}

var __func__80 = [23]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'l', 'i', 's', 't'}

func libusb_free_device_list(tls *libc.TLS, list uintptr, unref_devices int32) {
	var dev, v1 uintptr
	var i, v2 int32
	_, _, _, _ = dev, i, v1, v2
	if !(list != 0) {
		return
	}
	if unref_devices != 0 {
		i = 0
		for {
			v2 = i
			i++
			v1 = *(*uintptr)(unsafe.Pointer(list + uintptr(v2)*8))
			dev = v1
			if !(v1 != libc.UintptrFromInt32(0)) {
				break
			}
			libusb_unref_device(tls, dev)
		}
	}
	libc.Xfree(tls, list)
}

func libusb_get_session_data(tls *libc.TLS, dev uintptr) (r uint64) {
	return (*libusb_device)(unsafe.Pointer(dev)).Fsession_data
}

func libusb_get_bus_number(tls *libc.TLS, dev uintptr) (r uint8_t) {
	return (*libusb_device)(unsafe.Pointer(dev)).Fbus_number
}

func libusb_get_port_number(tls *libc.TLS, dev uintptr) (r uint8_t) {
	return (*libusb_device)(unsafe.Pointer(dev)).Fport_number
}

func libusb_get_port_numbers(tls *libc.TLS, dev uintptr, port_numbers uintptr, port_numbers_len int32) (r int32) {
	var ctx uintptr
	var i, v1 int32
	_, _, _ = ctx, i, v1
	i = port_numbers_len
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	if port_numbers_len <= 0 {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	for dev != 0 && libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fport_number) != 0 {
		i--
		v1 = i
		if v1 < 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__81)), __ccgo_ts+6580, 0)
			return int32(LIBUSB_ERROR_OVERFLOW)
		}
		*(*uint8_t)(unsafe.Pointer(port_numbers + uintptr(i))) = (*libusb_device)(unsafe.Pointer(dev)).Fport_number
		dev = (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev
	}
	if i < port_numbers_len {
		libc.Xmemmove(tls, port_numbers, port_numbers+uintptr(i), libc.Uint64FromInt32(port_numbers_len-i))
	}
	return port_numbers_len - i
}

var __func__81 = [24]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'p', 'o', 'r', 't', '_', 'n', 'u', 'm', 'b', 'e', 'r', 's'}

func libusb_get_port_path(tls *libc.TLS, ctx uintptr, dev uintptr, port_numbers uintptr, port_numbers_len uint8_t) (r int32) {
	_ = ctx
	return libusb_get_port_numbers(tls, dev, port_numbers, libc.Int32FromUint8(port_numbers_len))
}

func libusb_get_parent(tls *libc.TLS, dev uintptr) (r uintptr) {
	return (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev
}

func libusb_get_device_address(tls *libc.TLS, dev uintptr) (r uint8_t) {
	return (*libusb_device)(unsafe.Pointer(dev)).Fdevice_address
}

func libusb_get_device_speed(tls *libc.TLS, dev uintptr) (r int32) {
	return (*libusb_device)(unsafe.Pointer(dev)).Fspeed
}

func find_endpoint(tls *libc.TLS, config uintptr, endpoint uint8) (r uintptr) {
	var altsetting, ep, iface uintptr
	var altsetting_idx, ep_idx, iface_idx int32
	_, _, _, _, _, _ = altsetting, altsetting_idx, ep, ep_idx, iface, iface_idx
	iface_idx = 0
	for {
		if !(iface_idx < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces)) {
			break
		}
		iface = (*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1 + uintptr(iface_idx)*16
		altsetting_idx = 0
		for {
			if !(altsetting_idx < (*libusb_interface)(unsafe.Pointer(iface)).Fnum_altsetting) {
				break
			}
			altsetting = (*libusb_interface)(unsafe.Pointer(iface)).Faltsetting + uintptr(altsetting_idx)*40
			ep_idx = 0
			for {
				if !(ep_idx < libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(altsetting)).FbNumEndpoints)) {
					break
				}
				ep = (*libusb_interface_descriptor)(unsafe.Pointer(altsetting)).Fendpoint + uintptr(ep_idx)*32
				if libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbEndpointAddress) == libc.Int32FromUint8(endpoint) {
					return ep
				}
				goto _3
			_3:
				;
				ep_idx++
			}
			goto _2
		_2:
			;
			altsetting_idx++
		}
		goto _1
	_1:
		;
		iface_idx++
	}
	return libc.UintptrFromInt32(0)
}

func libusb_get_max_packet_size(tls *libc.TLS, dev uintptr, endpoint uint8) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ep uintptr
	var r int32
	var _ /* config at bp+0 */ uintptr
	_, _ = ep, r
	r = libusb_get_active_config_descriptor(tls, dev, bp)
	if r < 0 {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__82)), __ccgo_ts+6612, 0)
		return int32(LIBUSB_ERROR_OTHER)
	}
	ep = find_endpoint(tls, *(*uintptr)(unsafe.Pointer(bp)), endpoint)
	if !(ep != 0) {
		r = int32(LIBUSB_ERROR_NOT_FOUND)
		goto out
	}
	r = libc.Int32FromUint16((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FwMaxPacketSize)
	goto out
out:
	;
	libusb_free_config_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp)))
	return r
}

var __func__82 = [27]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'm', 'a', 'x', '_', 'p', 'a', 'c', 'k', 'e', 't', '_', 's', 'i', 'z', 'e'}

func find_alt_endpoint(tls *libc.TLS, config uintptr, iface_idx int32, altsetting_idx int32, endpoint uint8) (r uintptr) {
	var altsetting, ep, iface uintptr
	var ep_idx int32
	_, _, _, _ = altsetting, ep, ep_idx, iface
	if iface_idx >= libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces) {
		return libc.UintptrFromInt32(0)
	}
	iface = (*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1 + uintptr(iface_idx)*16
	if altsetting_idx >= (*libusb_interface)(unsafe.Pointer(iface)).Fnum_altsetting {
		return libc.UintptrFromInt32(0)
	}
	altsetting = (*libusb_interface)(unsafe.Pointer(iface)).Faltsetting + uintptr(altsetting_idx)*40
	ep_idx = 0
	for {
		if !(ep_idx < libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(altsetting)).FbNumEndpoints)) {
			break
		}
		ep = (*libusb_interface_descriptor)(unsafe.Pointer(altsetting)).Fendpoint + uintptr(ep_idx)*32
		if libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbEndpointAddress) == libc.Int32FromUint8(endpoint) {
			return ep
		}
		goto _1
	_1:
		;
		ep_idx++
	}
	return libc.UintptrFromInt32(0)
}

func get_endpoint_max_packet_size(tls *libc.TLS, dev uintptr, ep uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ep_type libusb_endpoint_transfer_type
	var r, speed int32
	var val uint16_t
	var _ /* ss_ep_cmp at bp+0 */ uintptr
	_, _, _, _ = ep_type, r, speed, val
	r = 0
	speed = libusb_get_device_speed(tls, dev)
	if speed >= int32(LIBUSB_SPEED_SUPER) {
		r = libusb_get_ss_endpoint_companion_descriptor(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, ep, bp)
		if r == int32(LIBUSB_SUCCESS) {
			r = libc.Int32FromUint16((*libusb_ss_endpoint_companion_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).FwBytesPerInterval)
			libusb_free_ss_endpoint_companion_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp)))
		}
	}
	if speed < int32(LIBUSB_SPEED_SUPER) || r < 0 {
		val = (*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FwMaxPacketSize
		ep_type = libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbmAttributes) & libc.Int32FromInt32(0x3)
		r = libc.Int32FromUint16(val) & int32(0x07ff)
		if ep_type == int32(LIBUSB_ENDPOINT_TRANSFER_TYPE_ISOCHRONOUS) || ep_type == int32(LIBUSB_ENDPOINT_TRANSFER_TYPE_INTERRUPT) {
			r *= int32(1) + libc.Int32FromUint16(val)>>int32(11)&int32(3)
		}
	}
	return r
}

func libusb_get_max_iso_packet_size(tls *libc.TLS, dev uintptr, endpoint uint8) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ep uintptr
	var r int32
	var _ /* config at bp+0 */ uintptr
	_, _ = ep, r
	r = libusb_get_active_config_descriptor(tls, dev, bp)
	if r < 0 {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__83)), __ccgo_ts+6612, 0)
		return int32(LIBUSB_ERROR_OTHER)
	}
	ep = find_endpoint(tls, *(*uintptr)(unsafe.Pointer(bp)), endpoint)
	if !(ep != 0) {
		r = int32(LIBUSB_ERROR_NOT_FOUND)
		goto out
	}
	r = get_endpoint_max_packet_size(tls, dev, ep)
	goto out
out:
	;
	libusb_free_config_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp)))
	return r
}

var __func__83 = [31]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'm', 'a', 'x', '_', 'i', 's', 'o', '_', 'p', 'a', 'c', 'k', 'e', 't', '_', 's', 'i', 'z', 'e'}

func libusb_get_max_alt_packet_size(tls *libc.TLS, dev uintptr, interface_number int32, alternate_setting int32, endpoint uint8) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ep uintptr
	var r int32
	var _ /* config at bp+0 */ uintptr
	_, _ = ep, r
	r = libusb_get_active_config_descriptor(tls, dev, bp)
	if r < 0 {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__84)), __ccgo_ts+6612, 0)
		return int32(LIBUSB_ERROR_OTHER)
	}
	ep = find_alt_endpoint(tls, *(*uintptr)(unsafe.Pointer(bp)), interface_number, alternate_setting, endpoint)
	if !(ep != 0) {
		r = int32(LIBUSB_ERROR_NOT_FOUND)
		goto out
	}
	r = get_endpoint_max_packet_size(tls, dev, ep)
	goto out
out:
	;
	libusb_free_config_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp)))
	return r
}

var __func__84 = [31]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'm', 'a', 'x', '_', 'a', 'l', 't', '_', 'p', 'a', 'c', 'k', 'e', 't', '_', 's', 'i', 'z', 'e'}

func libusb_ref_device(tls *libc.TLS, dev uintptr) (r uintptr) {
	var refcnt int64
	var v1 bool
	_, _ = refcnt, v1
	refcnt = int64(int32(libc.X__atomic_fetch_addInt64(tls, dev, int64(libc.Int32FromInt32(1)), libc.Int32FromInt32(5))) + libc.Int32FromInt32(1))
	if v1 = refcnt >= int64(2); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+6656, __ccgo_ts+6668, int32(1312), uintptr(unsafe.Pointer(&__func__85)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	_ = refcnt
	return dev
}

var __func__85 = [18]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'r', 'e', 'f', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func libusb_unref_device(tls *libc.TLS, dev uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var idx int32
	var refcnt int64
	var v1 bool
	_, _, _ = idx, refcnt, v1
	if !(dev != 0) {
		return
	}
	refcnt = int64(int32(libc.X__atomic_fetch_addInt64(tls, dev, int64(-libc.Int32FromInt32(1)), libc.Int32FromInt32(5))) - libc.Int32FromInt32(1))
	if v1 = refcnt >= 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+6709, __ccgo_ts+6668, int32(1331), uintptr(unsafe.Pointer(&__func__86)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if refcnt == 0 {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__86)), __ccgo_ts+6721, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_address)))
		libusb_unref_device(tls, (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev)
		if usbi_backend.Fdestroy_device != 0 {
			(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fdestroy_device})))(tls, dev)
		}
		if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
			usbi_disconnect_device(tls, dev)
		}
		idx = 0
		for {
			if !(idx < int32(LIBUSB_DEVICE_STRING_COUNT)) {
				break
			}
			libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(dev + 88 + uintptr(idx)*8)))
			goto _2
		_2:
			;
			idx++
		}
		libc.Xfree(tls, dev)
	}
}

var __func__86 = [20]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'u', 'n', 'r', 'e', 'f', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func libusb_wrap_sys_device(tls *libc.TLS, ctx uintptr, sys_dev intptr_t, dev_handle uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _dev_handle uintptr
	var priv_size size_t
	var r int32
	_, _, _ = _dev_handle, priv_size, r
	priv_size = usbi_backend.Fdevice_handle_priv_size
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__87)), __ccgo_ts+6742, libc.VaList(bp+8, libc.Uint64FromInt64(sys_dev)))
	ctx = usbi_get_context(tls, ctx)
	if !(usbi_backend.Fwrap_sys_device != 0) {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	_dev_handle = libc.Xcalloc(tls, uint64(1), (libc.Uint64FromInt64(80)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)) + priv_size)
	if !(_dev_handle != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	usbi_mutex_init(tls, _dev_handle)
	r = (*(*func(*libc.TLS, uintptr, uintptr, intptr_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fwrap_sys_device})))(tls, ctx, _dev_handle, sys_dev)
	if r < 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__87)), __ccgo_ts+6764, libc.VaList(bp+8, libc.Uint64FromInt64(sys_dev), r))
		usbi_mutex_destroy(tls, _dev_handle)
		libc.Xfree(tls, _dev_handle)
		return r
	}
	usbi_mutex_lock(tls, ctx+80)
	list_add(tls, _dev_handle+48, ctx+120)
	usbi_mutex_unlock(tls, ctx+80)
	*(*uintptr)(unsafe.Pointer(dev_handle)) = _dev_handle
	return 0
}

var __func__87 = [23]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'w', 'r', 'a', 'p', '_', 's', 'y', 's', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func libusb_open(tls *libc.TLS, dev uintptr, dev_handle uintptr) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var __atomic_load_ptr, _dev_handle, ctx uintptr
	var priv_size size_t
	var r int32
	var v1 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _, _, _, _ = __atomic_load_ptr, _dev_handle, ctx, priv_size, r, v1
	ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
	priv_size = usbi_backend.Fdevice_handle_priv_size
	usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__88)), __ccgo_ts+6797, libc.VaList(bp+16, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_address)))
	{
		__atomic_load_ptr = dev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v1 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	_dev_handle = libc.Xcalloc(tls, uint64(1), (libc.Uint64FromInt64(80)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)) + priv_size)
	if !(_dev_handle != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	usbi_mutex_init(tls, _dev_handle)
	(*libusb_device_handle)(unsafe.Pointer(_dev_handle)).Fdev = libusb_ref_device(tls, dev)
	r = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fopen})))(tls, _dev_handle)
	if r < 0 {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__88)), __ccgo_ts+6808, libc.VaList(bp+16, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_address), r))
		libusb_unref_device(tls, dev)
		usbi_mutex_destroy(tls, _dev_handle)
		libc.Xfree(tls, _dev_handle)
		return r
	}
	usbi_mutex_lock(tls, ctx+80)
	list_add(tls, _dev_handle+48, ctx+120)
	usbi_mutex_unlock(tls, ctx+80)
	*(*uintptr)(unsafe.Pointer(dev_handle)) = _dev_handle
	return 0
}

var __func__88 = [12]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'o', 'p', 'e', 'n'}

func libusb_open_device_with_vid_pid(tls *libc.TLS, ctx uintptr, vendor_id uint16_t, product_id uint16_t) (r1 uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var dev, found, v1 uintptr
	var i, v2 size_t
	var r int32
	var _ /* desc at bp+16 */ libusb_device_descriptor
	var _ /* dev_handle at bp+8 */ uintptr
	var _ /* devs at bp+0 */ uintptr
	_, _, _, _, _, _ = dev, found, i, r, v1, v2
	found = libc.UintptrFromInt32(0)
	*(*uintptr)(unsafe.Pointer(bp + 8)) = libc.UintptrFromInt32(0)
	i = uint64(0)
	if libusb_get_device_list(tls, ctx, bp) < 0 {
		return libc.UintptrFromInt32(0)
	}
	for {
		v2 = i
		i++
		v1 = *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + uintptr(v2)*8))
		dev = v1
		if !(v1 != libc.UintptrFromInt32(0)) {
			break
		}
		r = libusb_get_device_descriptor(tls, dev, bp+16)
		if r < 0 {
			goto out
		}
		if libc.Int32FromUint16((*(*libusb_device_descriptor)(unsafe.Pointer(bp + 16))).FidVendor) == libc.Int32FromUint16(vendor_id) && libc.Int32FromUint16((*(*libusb_device_descriptor)(unsafe.Pointer(bp + 16))).FidProduct) == libc.Int32FromUint16(product_id) {
			found = dev
			break
		}
	}
	if found != 0 {
		r = libusb_open(tls, found, bp+8)
		if r < 0 {
			*(*uintptr)(unsafe.Pointer(bp + 8)) = libc.UintptrFromInt32(0)
		}
	}
	goto out
out:
	;
	libusb_free_device_list(tls, *(*uintptr)(unsafe.Pointer(bp)), int32(1))
	return *(*uintptr)(unsafe.Pointer(bp + 8))
}

func do_close(tls *libc.TLS, ctx uintptr, dev_handle uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var itransfer, tmp, transfer uintptr
	var state_flags uint32_t
	_, _, _, _ = itransfer, state_flags, tmp, transfer
	usbi_mutex_lock(tls, ctx+208)
	itransfer = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+248)).Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	tmp = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(itransfer)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	for {
		if !(itransfer+48 != ctx+248) {
			break
		}
		transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
		if (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle != dev_handle {
			goto _1
		}
		usbi_mutex_lock(tls, itransfer)
		state_flags = (*usbi_transfer)(unsafe.Pointer(itransfer)).Fstate_flags
		usbi_mutex_unlock(tls, itransfer)
		if !(state_flags&uint32(USBI_TRANSFER_DEVICE_DISAPPEARED) != 0) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__89)), __ccgo_ts+6830, 0)
			if state_flags&uint32(USBI_TRANSFER_CANCELLING) != 0 {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__89)), __ccgo_ts+6945, 0)
			} else {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__89)), __ccgo_ts+7033, 0)
			}
		}
		list_del(tls, itransfer+48)
		(*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle = libc.UintptrFromInt32(0)
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__89)), __ccgo_ts+7123, libc.VaList(bp+8, transfer, dev_handle))
		goto _1
	_1:
		;
		itransfer = tmp
		tmp = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(tmp)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	}
	usbi_mutex_unlock(tls, ctx+208)
	usbi_mutex_lock(tls, ctx+80)
	list_del(tls, dev_handle+48)
	usbi_mutex_unlock(tls, ctx+80)
	(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fclose1})))(tls, dev_handle)
	libusb_unref_device(tls, (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)
	usbi_mutex_destroy(tls, dev_handle)
	libc.Xfree(tls, dev_handle)
}

var __func__89 = [9]uint8{'d', 'o', '_', 'c', 'l', 'o', 's', 'e'}

func libusb_close(tls *libc.TLS, dev_handle uintptr) {
	var ctx, v1, v3, v5 uintptr
	var event_flags, v2, v4 uint32
	var handling_events int32
	_, _, _, _, _, _, _, _ = ctx, event_flags, handling_events, v1, v2, v3, v4, v5
	if !(dev_handle != 0) {
		return
	}
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__90)), __ccgo_ts+6578, 0)
	handling_events = usbi_handling_events(tls, ctx)
	if !(handling_events != 0) {
		usbi_mutex_lock(tls, ctx+424)
		event_flags = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags
		v3 = ctx + 468
		v2 = *(*uint32)(unsafe.Pointer(v3))
		*(*uint32)(unsafe.Pointer(v3))++
		if !(v2 != 0) {
			*(*uint32)(unsafe.Pointer(ctx + 464)) |= uint32(USBI_EVENT_DEVICE_CLOSE)
		}
		if !(event_flags != 0) {
			usbi_signal_event(tls, ctx+16)
		}
		usbi_mutex_unlock(tls, ctx+424)
		libusb_lock_events(tls, ctx)
	}
	do_close(tls, ctx, dev_handle)
	if !(handling_events != 0) {
		usbi_mutex_lock(tls, ctx+424)
		v5 = ctx + 468
		*(*uint32)(unsafe.Pointer(v5))--
		v4 = *(*uint32)(unsafe.Pointer(v5))
		if !(v4 != 0) {
			*(*uint32)(unsafe.Pointer(ctx + 464)) &= libc.Uint32FromInt32(^int32(USBI_EVENT_DEVICE_CLOSE))
		}
		if !((*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags != 0) {
			usbi_clear_event(tls, ctx+16)
		}
		usbi_mutex_unlock(tls, ctx+424)
		libusb_unlock_events(tls, ctx)
	}
}

var __func__90 = [13]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'c', 'l', 'o', 's', 'e'}

func libusb_get_device(tls *libc.TLS, dev_handle uintptr) (r uintptr) {
	return (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev
}

func libusb_get_configuration(tls *libc.TLS, dev_handle uintptr, config uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ctx, v1 uintptr
	var r int32
	var _ /* tmp at bp+0 */ uint8_t
	_, _, _ = ctx, r, v1
	r = int32(LIBUSB_ERROR_NOT_SUPPORTED)
	*(*uint8_t)(unsafe.Pointer(bp)) = uint8(0)
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__91)), __ccgo_ts+6578, 0)
	if usbi_backend.Fget_configuration != 0 {
		r = (*(*func(*libc.TLS, uintptr, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_configuration})))(tls, dev_handle, bp)
	}
	if r == int32(LIBUSB_ERROR_NOT_SUPPORTED) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__91)), __ccgo_ts+7199, 0)
		r = libusb_control_transfer(tls, dev_handle, uint8(LIBUSB_ENDPOINT_IN), uint8(LIBUSB_REQUEST_GET_CONFIGURATION), uint16(0), uint16(0), bp, uint16(1), uint32(1000))
		if r == int32(1) {
			r = 0
		} else {
			if r == 0 {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__91)), __ccgo_ts+7231, 0)
				r = int32(LIBUSB_ERROR_IO)
			} else {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__91)), __ccgo_ts+7269, libc.VaList(bp+16, r))
			}
		}
	}
	if r == 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__91)), __ccgo_ts+7294, libc.VaList(bp+16, libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(bp)))))
		*(*int32)(unsafe.Pointer(config)) = libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(bp)))
	}
	return r
}

var __func__91 = [25]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'a', 't', 'i', 'o', 'n'}

func libusb_set_configuration(tls *libc.TLS, dev_handle uintptr, configuration int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var v1 uintptr
	_ = v1
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__92)), __ccgo_ts+7311, libc.VaList(bp+8, configuration))
	if configuration < -int32(1) || configuration > libc.Int32FromInt32(0xff) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	return (*(*func(*libc.TLS, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fset_configuration})))(tls, dev_handle, configuration)
}

var __func__92 = [25]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 's', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'a', 't', 'i', 'o', 'n'}

func libusb_claim_interface(tls *libc.TLS, dev_handle uintptr, interface_number int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var r int32
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _, _ = __atomic_load_ptr, r, v1, v2
	r = 0
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__93)), __ccgo_ts+7328, libc.VaList(bp+16, interface_number))
	if interface_number < 0 || interface_number >= int32(32) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	usbi_mutex_lock(tls, dev_handle)
	if (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fclaimed_interfaces&uint64(libc.Uint32FromUint32(1)<<interface_number) != 0 {
		goto out
	}
	r = (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fclaim_interface})))(tls, dev_handle, libc.Uint8FromInt32(interface_number))
	if r == 0 {
		*(*uint64)(unsafe.Pointer(dev_handle + 40)) |= uint64(uint32(1) << interface_number)
	}
	goto out
out:
	;
	usbi_mutex_unlock(tls, dev_handle)
	return r
}

var __func__93 = [23]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'c', 'l', 'a', 'i', 'm', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e'}

func libusb_release_interface(tls *libc.TLS, dev_handle uintptr, interface_number int32) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r int32
	var v1 uintptr
	_, _ = r, v1
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__94)), __ccgo_ts+7328, libc.VaList(bp+8, interface_number))
	if interface_number < 0 || interface_number >= int32(32) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	usbi_mutex_lock(tls, dev_handle)
	if !((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fclaimed_interfaces&uint64(libc.Uint32FromUint32(1)<<interface_number) != 0) {
		r = int32(LIBUSB_ERROR_NOT_FOUND)
		goto out
	}
	r = (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Frelease_interface})))(tls, dev_handle, libc.Uint8FromInt32(interface_number))
	if r == 0 {
		*(*uint64)(unsafe.Pointer(dev_handle + 40)) &= uint64(^(libc.Uint32FromUint32(1) << interface_number))
	}
	goto out
out:
	;
	usbi_mutex_unlock(tls, dev_handle)
	return r
}

var __func__94 = [25]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'r', 'e', 'l', 'e', 'a', 's', 'e', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e'}

func libusb_set_interface_alt_setting(tls *libc.TLS, dev_handle uintptr, interface_number int32, alternate_setting int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__95)), __ccgo_ts+7341, libc.VaList(bp+16, interface_number, alternate_setting))
	if interface_number < 0 || interface_number >= int32(32) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	if alternate_setting < 0 || alternate_setting > libc.Int32FromInt32(0xff) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	usbi_mutex_lock(tls, dev_handle)
	if !((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fclaimed_interfaces&uint64(libc.Uint32FromUint32(1)<<interface_number) != 0) {
		usbi_mutex_unlock(tls, dev_handle)
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	usbi_mutex_unlock(tls, dev_handle)
	return (*(*func(*libc.TLS, uintptr, uint8_t, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fset_interface_altsetting})))(tls, dev_handle, libc.Uint8FromInt32(interface_number), libc.Uint8FromInt32(alternate_setting))
}

var __func__95 = [33]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 's', 'e', 't', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e', '_', 'a', 'l', 't', '_', 's', 'e', 't', 't', 'i', 'n', 'g'}

func libusb_clear_halt(tls *libc.TLS, dev_handle uintptr, endpoint uint8) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__96)), __ccgo_ts+7368, libc.VaList(bp+16, libc.Int32FromUint8(endpoint)))
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	return (*(*func(*libc.TLS, uintptr, uint8) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fclear_halt})))(tls, dev_handle, endpoint)
}

var __func__96 = [18]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'c', 'l', 'e', 'a', 'r', '_', 'h', 'a', 'l', 't'}

func libusb_reset_device(tls *libc.TLS, dev_handle uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__97)), __ccgo_ts+6578, 0)
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	if usbi_backend.Freset_device != 0 {
		return (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Freset_device})))(tls, dev_handle)
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

var __func__97 = [20]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'r', 'e', 's', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e'}

func libusb_alloc_streams(tls *libc.TLS, dev_handle uintptr, num_streams uint32_t, endpoints uintptr, num_endpoints int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__98)), __ccgo_ts+7382, libc.VaList(bp+16, num_streams, num_endpoints))
	if !(num_streams != 0) || !(endpoints != 0) || num_endpoints <= 0 {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	if usbi_backend.Falloc_streams != 0 {
		return (*(*func(*libc.TLS, uintptr, uint32_t, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Falloc_streams})))(tls, dev_handle, num_streams, endpoints, num_endpoints)
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

var __func__98 = [21]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'a', 'l', 'l', 'o', 'c', '_', 's', 't', 'r', 'e', 'a', 'm', 's'}

func libusb_free_streams(tls *libc.TLS, dev_handle uintptr, endpoints uintptr, num_endpoints int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__99)), __ccgo_ts+7400, libc.VaList(bp+16, num_endpoints))
	if !(endpoints != 0) || num_endpoints <= 0 {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	if usbi_backend.Ffree_streams != 0 {
		return (*(*func(*libc.TLS, uintptr, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Ffree_streams})))(tls, dev_handle, endpoints, num_endpoints)
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

var __func__99 = [20]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'f', 'r', 'e', 'e', '_', 's', 't', 'r', 'e', 'a', 'm', 's'}

func libusb_dev_mem_alloc(tls *libc.TLS, dev_handle uintptr, length size_t) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_load_ptr uintptr
	var v1 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _ = __atomic_load_ptr, v1
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v1 != 0) {
		return libc.UintptrFromInt32(0)
	}
	if usbi_backend.Fdev_mem_alloc != 0 {
		return (*(*func(*libc.TLS, uintptr, size_t) uintptr)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fdev_mem_alloc})))(tls, dev_handle, length)
	} else {
		return libc.UintptrFromInt32(0)
	}
	return r
}

func libusb_dev_mem_free(tls *libc.TLS, dev_handle uintptr, buffer uintptr, length size_t) (r int32) {
	if usbi_backend.Fdev_mem_free != 0 {
		return (*(*func(*libc.TLS, uintptr, uintptr, size_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fdev_mem_free})))(tls, dev_handle, buffer, length)
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

func libusb_kernel_driver_active(tls *libc.TLS, dev_handle uintptr, interface_number int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__100)), __ccgo_ts+7328, libc.VaList(bp+16, interface_number))
	if interface_number < 0 || interface_number >= int32(32) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	if usbi_backend.Fkernel_driver_active != 0 {
		return (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fkernel_driver_active})))(tls, dev_handle, libc.Uint8FromInt32(interface_number))
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

var __func__100 = [28]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r', '_', 'a', 'c', 't', 'i', 'v', 'e'}

func libusb_detach_kernel_driver(tls *libc.TLS, dev_handle uintptr, interface_number int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__101)), __ccgo_ts+7328, libc.VaList(bp+16, interface_number))
	if interface_number < 0 || interface_number >= int32(32) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	if usbi_backend.Fdetach_kernel_driver != 0 {
		return (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fdetach_kernel_driver})))(tls, dev_handle, libc.Uint8FromInt32(interface_number))
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

var __func__101 = [28]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'd', 'e', 't', 'a', 'c', 'h', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r'}

func libusb_attach_kernel_driver(tls *libc.TLS, dev_handle uintptr, interface_number int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var __atomic_load_ptr, v1 uintptr
	var v2 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__102)), __ccgo_ts+7328, libc.VaList(bp+16, interface_number))
	if interface_number < 0 || interface_number >= int32(32) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	{
		__atomic_load_ptr = (*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev + 80
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v2 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v2 != 0) {
		return int32(LIBUSB_ERROR_NO_DEVICE)
	}
	if usbi_backend.Fattach_kernel_driver != 0 {
		return (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fattach_kernel_driver})))(tls, dev_handle, libc.Uint8FromInt32(interface_number))
	} else {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return r
}

var __func__102 = [28]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'a', 't', 't', 'a', 'c', 'h', '_', 'k', 'e', 'r', 'n', 'e', 'l', '_', 'd', 'r', 'i', 'v', 'e', 'r'}

func libusb_set_auto_detach_kernel_driver(tls *libc.TLS, dev_handle uintptr, enable int32) (r int32) {
	if !(usbi_backend.Fcaps&libc.Uint32FromInt32(0x00020000) != 0) {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	(*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fauto_detach_kernel_driver = enable
	return int32(LIBUSB_SUCCESS)
}

func libusb_endpoint_supports_raw_io(tls *libc.TLS, dev_handle uintptr, endpoint uint8_t) (r int32) {
	var v1, v2 bool
	_, _ = v1, v2
	if usbi_backend.Fendpoint_supports_raw_io == libc.UintptrFromInt32(0) {
		return 0
	}
	if v1 = usbi_backend.Fendpoint_set_raw_io != libc.UintptrFromInt32(0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+7407, __ccgo_ts+6668, int32(2274), uintptr(unsafe.Pointer(&__func__103)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if v2 = usbi_backend.Fget_max_raw_io_transfer_size != libc.UintptrFromInt32(0); !v2 {
		libc.X__assert_fail(tls, __ccgo_ts+7448, __ccgo_ts+6668, int32(2275), uintptr(unsafe.Pointer(&__func__103)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	return (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fendpoint_supports_raw_io})))(tls, dev_handle, endpoint)
}

var __func__103 = [32]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'e', 'n', 'd', 'p', 'o', 'i', 'n', 't', '_', 's', 'u', 'p', 'p', 'o', 'r', 't', 's', '_', 'r', 'a', 'w', '_', 'i', 'o'}

func libusb_endpoint_set_raw_io(tls *libc.TLS, dev_handle uintptr, endpoint uint8_t, enable int32) (r int32) {
	if !(usbi_backend.Fendpoint_set_raw_io != 0) {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return (*(*func(*libc.TLS, uintptr, uint8_t, int32) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fendpoint_set_raw_io})))(tls, dev_handle, endpoint, enable)
}

func libusb_get_max_raw_io_transfer_size(tls *libc.TLS, dev_handle uintptr, endpoint uint8_t) (r int32) {
	if !(usbi_backend.Fget_max_raw_io_transfer_size != 0) {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return (*(*func(*libc.TLS, uintptr, uint8_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_max_raw_io_transfer_size})))(tls, dev_handle, endpoint)
}

func libusb_set_debug(tls *libc.TLS, ctx uintptr, level int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	libusb_set_option(tls, ctx, int32(LIBUSB_OPTION_LOG_LEVEL), libc.VaList(bp+8, level))
}

func libusb_set_log_cb_internal(tls *libc.TLS, ctx uintptr, __ccgo_fp_cb libusb_log_cb, mode int32) {
	if mode&int32(LIBUSB_LOG_CB_CONTEXT) != 0 {
		ctx = usbi_get_context(tls, ctx)
		(*libusb_context)(unsafe.Pointer(ctx)).Flog_handler = __ccgo_fp_cb
	}
}

type __ccgo_fp__Xlibusb_set_log_cb_1 = func(*libc.TLS, uintptr, int32, uintptr)

func libusb_set_log_cb(tls *libc.TLS, ctx uintptr, __ccgo_fp_cb libusb_log_cb, mode int32) {
	libusb_set_log_cb_internal(tls, ctx, __ccgo_fp_cb, mode)
}

func libusb_set_option(tls *libc.TLS, ctx uintptr, option libusb_option, va uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_store_ptr uintptr
	var ap va_list
	var arg, is_default_context, r, v3, v4 int32
	var log_cb libusb_log_cb
	var _ /* __atomic_store_tmp at bp+0 */ usbi_atomic_t
	_, _, _, _, _, _, _, _ = __atomic_store_ptr, ap, arg, is_default_context, log_cb, r, v3, v4
	arg = 0
	r = int32(LIBUSB_SUCCESS)
	log_cb = libc.UintptrFromInt32(0)
	is_default_context = libc.BoolInt32(libc.UintptrFromInt32(0) == ctx)
	ap = va
	if int32(LIBUSB_OPTION_LOG_LEVEL) == option {
		arg = libc.VaInt32(&ap)
		if arg < int32(LIBUSB_LOG_LEVEL_NONE) || arg > int32(LIBUSB_LOG_LEVEL_DEBUG) {
			r = int32(LIBUSB_ERROR_INVALID_PARAM)
		}
	}
	if int32(LIBUSB_OPTION_LOG_CB) == option {
		log_cb = libc.VaUintptr(&ap)
	}
	if int32(LIBUSB_SUCCESS) != r {
		goto _1
	}
	if option >= int32(LIBUSB_OPTION_MAX) {
		r = int32(LIBUSB_ERROR_INVALID_PARAM)
		goto _1
	}
	if libc.UintptrFromInt32(0) == ctx {
		usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
		default_context_options[option].Fis_set = int32(1)
		if int32(LIBUSB_OPTION_LOG_LEVEL) == option {
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&default_context_options)) + uintptr(option)*16 + 8)) = arg
		} else {
			if int32(LIBUSB_OPTION_LOG_CB) == option {
				*(*libusb_log_cb)(unsafe.Pointer(uintptr(unsafe.Pointer(&default_context_options)) + uintptr(option)*16 + 8)) = log_cb
				libusb_set_log_cb_internal(tls, libc.UintptrFromInt32(0), log_cb, int32(LIBUSB_LOG_CB_GLOBAL))
			}
		}
		usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
	}
	ctx = usbi_get_context(tls, ctx)
	if libc.UintptrFromInt32(0) == ctx {
		goto _1
	}
	switch option {
	case int32(LIBUSB_OPTION_LOG_LEVEL):
		if !((*libusb_context)(unsafe.Pointer(ctx)).Fdebug_fixed != 0) {
			(*libusb_context)(unsafe.Pointer(ctx)).Fdebug = arg
			if is_default_context != 0 {
				{
					__atomic_store_ptr = uintptr(unsafe.Pointer(&default_debug_level))
					if arg < int32(LIBUSB_LOG_LEVEL_NONE) {
						v3 = int32(LIBUSB_LOG_LEVEL_NONE)
					} else {
						if arg > int32(LIBUSB_LOG_LEVEL_DEBUG) {
							v4 = int32(LIBUSB_LOG_LEVEL_DEBUG)
						} else {
							v4 = arg
						}
						v3 = v4
					}
					*(*usbi_atomic_t)(unsafe.Pointer(bp)) = int64(v3)
					libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp, libc.Int32FromInt32(5))
				}
			}
		}
	case int32(LIBUSB_OPTION_USE_USBDK):
		fallthrough
	case int32(LIBUSB_OPTION_NO_DEVICE_DISCOVERY):
		if usbi_backend.Fset_option != 0 {
			r = (*(*func(*libc.TLS, uintptr, libusb_option, va_list) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fset_option})))(tls, ctx, option, ap)
			break
		}
		r = int32(LIBUSB_ERROR_NOT_SUPPORTED)
	case int32(LIBUSB_OPTION_LOG_CB):
		libusb_set_log_cb_internal(tls, ctx, log_cb, int32(LIBUSB_LOG_CB_CONTEXT))
	case int32(LIBUSB_OPTION_MAX):
		fallthrough
	default:
		r = int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	goto _2
_2:
	;
	goto _1
_1: /**/
	; //
	_ = ap
	return r
}

func get_env_debug_level(tls *libc.TLS) (r libusb_log_level) {
	var dbg uintptr
	var dbg_level, v1, v2 int64
	var level libusb_log_level
	_, _, _, _, _ = dbg, dbg_level, level, v1, v2
	level = int32(LIBUSB_LOG_LEVEL_NONE)
	dbg = libc.Xgetenv(tls, __ccgo_ts+7498)
	if dbg != 0 {
		dbg_level = libc.Xstrtol(tls, dbg, libc.UintptrFromInt32(0), int32(10))
		if dbg_level < int64(int32(LIBUSB_LOG_LEVEL_NONE)) {
			v1 = int64(int32(LIBUSB_LOG_LEVEL_NONE))
		} else {
			if dbg_level > int64(int32(LIBUSB_LOG_LEVEL_DEBUG)) {
				v2 = int64(int32(LIBUSB_LOG_LEVEL_DEBUG))
			} else {
				v2 = dbg_level
			}
			v1 = v2
		}
		dbg_level = v1
		level = int32(dbg_level)
	}
	return level
}

func libusb_init(tls *libc.TLS, ctx uintptr) (r int32) {
	return libusb_init_context(tls, ctx, libc.UintptrFromInt32(0), 0)
}

func libusb_init_context(tls *libc.TLS, ctx uintptr, options uintptr, num_options int32) (r1 int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var __atomic_load_ptr, __atomic_store_ptr, __atomic_store_ptr1, _ctx uintptr
	var i, r int32
	var option libusb_option
	var priv_size size_t
	var v3 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+8 */ usbi_atomic_t
	var _ /* __atomic_store_tmp at bp+0 */ usbi_atomic_t
	var _ /* __atomic_store_tmp at bp+16 */ usbi_atomic_t
	_, _, _, _, _, _, _, _, _ = __atomic_load_ptr, __atomic_store_ptr, __atomic_store_ptr1, _ctx, i, option, priv_size, r, v3
	priv_size = usbi_backend.Fcontext_priv_size
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
	if !(ctx != 0) && default_context_refcnt > 0 {
		usbi_log(tls, usbi_default_context, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__104)), __ccgo_ts+7511, 0)
		default_context_refcnt++
		usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
		return 0
	}
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	if !(active_contexts_list.Fnext != 0) {
		list_init(tls, uintptr(unsafe.Pointer(&active_contexts_list)))
		usbi_get_monotonic_time(tls, uintptr(unsafe.Pointer(&timestamp_origin)))
	}
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	_ctx = libc.Xcalloc(tls, uint64(1), (libc.Uint64FromInt64(568)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)) + priv_size)
	if !(_ctx != 0) {
		usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_context)(unsafe.Pointer(_ctx)).Fdebug = int32(LIBUSB_LOG_LEVEL_NONE)
	if libc.Xgetenv(tls, __ccgo_ts+7498) != 0 {
		(*libusb_context)(unsafe.Pointer(_ctx)).Fdebug = get_env_debug_level(tls)
		(*libusb_context)(unsafe.Pointer(_ctx)).Fdebug_fixed = int32(1)
	} else {
		if default_context_options[int32(LIBUSB_OPTION_LOG_LEVEL)].Fis_set != 0 {
			(*libusb_context)(unsafe.Pointer(_ctx)).Fdebug = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&default_context_options)) + uintptr(LIBUSB_OPTION_LOG_LEVEL)*16 + 8))
		}
	}
	usbi_mutex_init(tls, _ctx+24)
	usbi_mutex_init(tls, _ctx+80)
	list_init(tls, _ctx+64)
	list_init(tls, _ctx+120)
	option = 0
	for {
		if !(option < int32(LIBUSB_OPTION_MAX)) {
			break
		}
		if int32(LIBUSB_OPTION_LOG_LEVEL) == option || !(default_context_options[option].Fis_set != 0) {
			goto _1
		}
		if int32(LIBUSB_OPTION_LOG_CB) != option {
			r = libusb_set_option(tls, _ctx, option, 0)
		} else {
			r = libusb_set_option(tls, _ctx, option, libc.VaList(bp+32, *(*libusb_log_cb)(unsafe.Pointer(uintptr(unsafe.Pointer(&default_context_options)) + uintptr(option)*16 + 8))))
		}
		if int32(LIBUSB_SUCCESS) != r {
			goto err_free_ctx
		}
		goto _1
	_1:
		;
		option++
	}
	i = 0
	for {
		if !(i < num_options) {
			break
		}
		switch (*(*libusb_init_option)(unsafe.Pointer(options + uintptr(i)*16))).Foption {
		case int32(LIBUSB_OPTION_LOG_CB):
			r = libusb_set_option(tls, _ctx, (*(*libusb_init_option)(unsafe.Pointer(options + uintptr(i)*16))).Foption, libc.VaList(bp+32, *(*libusb_log_cb)(unsafe.Pointer(options + uintptr(i)*16 + 8))))
		case int32(LIBUSB_OPTION_LOG_LEVEL):
			fallthrough
		case int32(LIBUSB_OPTION_USE_USBDK):
			fallthrough
		case int32(LIBUSB_OPTION_NO_DEVICE_DISCOVERY):
			fallthrough
		case int32(LIBUSB_OPTION_MAX):
			fallthrough
		default:
			r = libusb_set_option(tls, _ctx, (*(*libusb_init_option)(unsafe.Pointer(options + uintptr(i)*16))).Foption, libc.VaList(bp+32, *(*int32)(unsafe.Pointer(options + uintptr(i)*16 + 8))))
		}
		if int32(LIBUSB_SUCCESS) != r {
			goto err_free_ctx
		}
		goto _2
	_2:
		;
		i++
	}
	if !(ctx != 0) {
		usbi_default_context = _ctx
		default_context_refcnt = int32(1)
		{
			__atomic_store_ptr = uintptr(unsafe.Pointer(&default_debug_level))
			*(*usbi_atomic_t)(unsafe.Pointer(bp)) = int64((*libusb_context)(unsafe.Pointer(_ctx)).Fdebug)
			libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp, libc.Int32FromInt32(5))
		}
		usbi_log(tls, usbi_default_context, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__104)), __ccgo_ts+7535, 0)
	}
	usbi_log(tls, _ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__104)), __ccgo_ts+7559, libc.VaList(bp+32, libc.Int32FromUint16(libusb_version_internal.Fmajor), libc.Int32FromUint16(libusb_version_internal.Fminor), libc.Int32FromUint16(libusb_version_internal.Fmicro), libc.Int32FromUint16(libusb_version_internal.Fnano), libusb_version_internal.Frc))
	r = usbi_io_init(tls, _ctx)
	if r < 0 {
		goto err_free_ctx
	}
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	list_add(tls, _ctx+552, uintptr(unsafe.Pointer(&active_contexts_list)))
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	if usbi_backend.Finit1 != 0 {
		r = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Finit1})))(tls, _ctx)
		if r != 0 {
			goto err_io_exit
		}
	}
	usbi_hotplug_init(tls, _ctx)
	if ctx != 0 {
		*(*uintptr)(unsafe.Pointer(ctx)) = _ctx
		if !(usbi_fallback_context != 0) {
			{
				__atomic_load_ptr = uintptr(unsafe.Pointer(&default_debug_level))
				libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp+8, libc.Int32FromInt32(5))
				v3 = *(*usbi_atomic_t)(unsafe.Pointer(bp + 8))
			}
			if v3 == int64(-int32(1)) {
				{
					__atomic_store_ptr1 = uintptr(unsafe.Pointer(&default_debug_level))
					*(*usbi_atomic_t)(unsafe.Pointer(bp + 16)) = int64((*libusb_context)(unsafe.Pointer(_ctx)).Fdebug)
					libc.X__atomic_storeInt64(tls, __atomic_store_ptr1, bp+16, libc.Int32FromInt32(5))
				}
			}
			usbi_fallback_context = _ctx
			usbi_log(tls, usbi_fallback_context, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__104)), __ccgo_ts+7581, 0)
		}
	}
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
	return 0
	goto err_io_exit
err_io_exit:
	;
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	list_del(tls, _ctx+552)
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	usbi_hotplug_exit(tls, _ctx)
	usbi_io_exit(tls, _ctx)
	goto err_free_ctx
err_free_ctx:
	;
	if !(ctx != 0) {
		usbi_default_context = libc.UintptrFromInt32(0)
		default_context_refcnt = 0
	}
	usbi_mutex_destroy(tls, _ctx+80)
	usbi_mutex_destroy(tls, _ctx+24)
	libc.Xfree(tls, _ctx)
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
	return r
}

var __func__104 = [20]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'i', 'n', 'i', 't', '_', 'c', 'o', 'n', 't', 'e', 'x', 't'}

func libusb_exit(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ctx, dev uintptr
	var v1 int32
	_, _, _ = _ctx, dev, v1
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
	if !(ctx != 0) {
		if !(usbi_default_context != 0) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__105)), __ccgo_ts+7624, 0)
			usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
			return
		}
		default_context_refcnt--
		v1 = default_context_refcnt
		if v1 > 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__105)), __ccgo_ts+7661, 0)
			usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
			return
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__105)), __ccgo_ts+7692, 0)
		_ctx = usbi_default_context
	} else {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__105)), __ccgo_ts+6578, 0)
		_ctx = ctx
	}
	usbi_mutex_static_lock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	list_del(tls, _ctx+552)
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&active_contexts_lock)))
	usbi_hotplug_exit(tls, _ctx)
	if usbi_backend.Fexit != 0 {
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fexit})))(tls, _ctx)
	}
	if !(ctx != 0) {
		usbi_default_context = libc.UintptrFromInt32(0)
	}
	if ctx == usbi_fallback_context {
		usbi_fallback_context = libc.UintptrFromInt32(0)
	}
	usbi_mutex_static_unlock(tls, uintptr(unsafe.Pointer(&default_context_lock)))
	usbi_io_exit(tls, _ctx)
	dev = uintptr(uint64((*list_head)(unsafe.Pointer(_ctx+64)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(dev+32 != _ctx+64) {
			break
		}
		usbi_log(tls, _ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__105)), __ccgo_ts+7719, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_address)))
		(*libusb_device)(unsafe.Pointer(dev)).Fctx = libc.UintptrFromInt32(0)
		goto _2
	_2:
		;
		dev = uintptr(uint64((*libusb_device)(unsafe.Pointer(dev)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	if !((*list_head)(unsafe.Pointer(_ctx+120)).Fnext == _ctx+120) {
		usbi_log(tls, _ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__105)), __ccgo_ts+7749, 0)
	}
	usbi_mutex_destroy(tls, _ctx+80)
	usbi_mutex_destroy(tls, _ctx+24)
	libc.Xfree(tls, _ctx)
}

var __func__105 = [12]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'e', 'x', 'i', 't'}

func libusb_has_capability(tls *libc.TLS, capability uint32_t) (r int32) {
	switch capability {
	case uint32(LIBUSB_CAP_HAS_CAPABILITY):
		return int32(1)
	case uint32(LIBUSB_CAP_HAS_HOTPLUG):
		return libc.BoolInt32(!(usbi_backend.Fget_device_list != 0))
	case uint32(LIBUSB_CAP_HAS_HID_ACCESS):
		return libc.Int32FromUint32(usbi_backend.Fcaps & libc.Uint32FromInt32(0x00010000))
	case uint32(LIBUSB_CAP_SUPPORTS_DETACH_KERNEL_DRIVER):
		return libc.Int32FromUint32(usbi_backend.Fcaps & libc.Uint32FromInt32(0x00020000))
	}
	return 0
}

func log_str(tls *libc.TLS, level libusb_log_level, str uintptr) {
	_ = level
	libc.Xfputs(tls, str, libc.Xstderr)
}

func log_v(tls *libc.TLS, ctx uintptr, level libusb_log_level, function uintptr, format uintptr, args va_list) {
	bp := tls.Alloc(1104)
	defer tls.Free(1104)
	var __atomic_load_ptr, prefix uintptr
	var ctx_level libusb_log_level
	var default_level_value int64
	var global_debug, header_len, text_len, v2 int32
	var v1 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+1024 */ usbi_atomic_t
	var _ /* buf at bp+0 */ [1024]uint8
	var _ /* timestamp at bp+1032 */ timespec
	_, _, _, _, _, _, _, _, _ = __atomic_load_ptr, ctx_level, default_level_value, global_debug, header_len, prefix, text_len, v1, v2
	if ctx != 0 {
		ctx_level = (*libusb_context)(unsafe.Pointer(ctx)).Fdebug
	} else {
		{
			__atomic_load_ptr = uintptr(unsafe.Pointer(&default_debug_level))
			libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp+1024, libc.Int32FromInt32(5))
			v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp + 1024))
		}
		default_level_value = v1
		if default_level_value < 0 {
			v2 = get_env_debug_level(tls)
		} else {
			v2 = int32(default_level_value)
		}
		ctx_level = v2
	}
	if ctx_level < level {
		return
	}
	global_debug = libc.BoolInt32(ctx_level == int32(LIBUSB_LOG_LEVEL_DEBUG))
	switch level {
	case int32(LIBUSB_LOG_LEVEL_NONE):
		return
	case int32(LIBUSB_LOG_LEVEL_ERROR):
		prefix = __ccgo_ts + 7784
	case int32(LIBUSB_LOG_LEVEL_WARNING):
		prefix = __ccgo_ts + 7790
	case int32(LIBUSB_LOG_LEVEL_INFO):
		prefix = __ccgo_ts + 7798
	case int32(LIBUSB_LOG_LEVEL_DEBUG):
		prefix = __ccgo_ts + 7803
	default:
		prefix = __ccgo_ts + 7809
		break
	}
	if global_debug != 0 {
		if !(has_debug_header_been_displayed != 0) {
			has_debug_header_been_displayed = int32(1)
			log_str(tls, int32(LIBUSB_LOG_LEVEL_DEBUG), __ccgo_ts+7817)
			log_str(tls, int32(LIBUSB_LOG_LEVEL_DEBUG), __ccgo_ts+7882)
		}
		usbi_get_monotonic_time(tls, bp+1032)
		(*timespec)(unsafe.Pointer(bp + 1032)).Ftv_sec = (*timespec)(unsafe.Pointer(bp+1032)).Ftv_sec - (*timespec)(unsafe.Pointer(uintptr(unsafe.Pointer(&timestamp_origin)))).Ftv_sec
		(*timespec)(unsafe.Pointer(bp + 1032)).Ftv_nsec = (*timespec)(unsafe.Pointer(bp+1032)).Ftv_nsec - (*timespec)(unsafe.Pointer(uintptr(unsafe.Pointer(&timestamp_origin)))).Ftv_nsec
		if (*timespec)(unsafe.Pointer(bp+1032)).Ftv_nsec < 0 {
			(*timespec)(unsafe.Pointer(bp+1032)).Ftv_sec--
			*(*int64)(unsafe.Pointer(bp + 1032 + 8)) += int64(1000000000)
		}
		header_len = libc.X__builtin_snprintf(tls, bp, uint64(1024), __ccgo_ts+7964, libc.VaList(bp+1056, (*(*timespec)(unsafe.Pointer(bp + 1032))).Ftv_sec, (*(*timespec)(unsafe.Pointer(bp + 1032))).Ftv_nsec/libc.Int64FromInt64(1000), usbi_get_tid(tls), prefix, function))
	} else {
		header_len = libc.X__builtin_snprintf(tls, bp, uint64(1024), __ccgo_ts+8002, libc.VaList(bp+1056, prefix, function))
	}
	if header_len < 0 || header_len >= libc.Int32FromInt64(1024) {
		header_len = 0
	}
	text_len = libc.X__builtin_vsnprintf(tls, bp+uintptr(header_len), uint64(1024)-libc.Uint64FromInt32(header_len), format, args)
	if text_len < 0 || text_len+header_len >= libc.Int32FromInt64(1024) {
		text_len = libc.Int32FromInt64(1024) - header_len
	}
	if header_len+text_len+libc.Int32FromInt64(2) >= libc.Int32FromInt64(1024) {
		text_len -= header_len + text_len + libc.Int32FromInt64(2) - libc.Int32FromInt64(1024)
	}
	libc.Xstrcpy(tls, bp+uintptr(header_len)+uintptr(text_len), __ccgo_ts+8019)
	log_str(tls, level, bp)
	if ctx != 0 && (*libusb_context)(unsafe.Pointer(ctx)).Flog_handler != 0 {
		(*(*func(*libc.TLS, uintptr, libusb_log_level, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*libusb_context)(unsafe.Pointer(ctx)).Flog_handler})))(tls, ctx, level, bp)
	}
}

var has_debug_header_been_displayed int32

func usbi_log(tls *libc.TLS, ctx uintptr, level libusb_log_level, function uintptr, format uintptr, va uintptr) {
	var args va_list
	_ = args
	args = va
	log_v(tls, ctx, level, function, format, args)
	_ = args
}

func libusb_error_name(tls *libc.TLS, error_code int32) (r uintptr) {
	switch error_code {
	case int32(LIBUSB_ERROR_IO):
		return __ccgo_ts + 8021
	case int32(LIBUSB_ERROR_INVALID_PARAM):
		return __ccgo_ts + 8037
	case int32(LIBUSB_ERROR_ACCESS):
		return __ccgo_ts + 8064
	case int32(LIBUSB_ERROR_NO_DEVICE):
		return __ccgo_ts + 8084
	case int32(LIBUSB_ERROR_NOT_FOUND):
		return __ccgo_ts + 8107
	case int32(LIBUSB_ERROR_BUSY):
		return __ccgo_ts + 8130
	case int32(LIBUSB_ERROR_TIMEOUT):
		return __ccgo_ts + 8148
	case int32(LIBUSB_ERROR_OVERFLOW):
		return __ccgo_ts + 8169
	case int32(LIBUSB_ERROR_PIPE):
		return __ccgo_ts + 8191
	case int32(LIBUSB_ERROR_INTERRUPTED):
		return __ccgo_ts + 8209
	case int32(LIBUSB_ERROR_NO_MEM):
		return __ccgo_ts + 8234
	case int32(LIBUSB_ERROR_NOT_SUPPORTED):
		return __ccgo_ts + 8254
	case int32(LIBUSB_ERROR_OTHER):
		return __ccgo_ts + 8281
	case int32(LIBUSB_TRANSFER_ERROR):
		return __ccgo_ts + 8300
	case int32(LIBUSB_TRANSFER_TIMED_OUT):
		return __ccgo_ts + 8322
	case int32(LIBUSB_TRANSFER_CANCELLED):
		return __ccgo_ts + 8348
	case int32(LIBUSB_TRANSFER_STALL):
		return __ccgo_ts + 8374
	case int32(LIBUSB_TRANSFER_NO_DEVICE):
		return __ccgo_ts + 8396
	case int32(LIBUSB_TRANSFER_OVERFLOW):
		return __ccgo_ts + 8422
	case 0:
		return __ccgo_ts + 8447
	default:
		return __ccgo_ts + 8490
	}
	return r
}

func libusb_get_version(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&libusb_version_internal))
}

func ReadLittleEndian16(tls *libc.TLS, p uintptr) (r uint16_t) {
	return libc.Uint16FromInt32(libc.Int32FromUint16(uint16(*(*uint8_t)(unsafe.Pointer(p + 1))))<<libc.Int32FromInt32(8) | libc.Int32FromUint16(uint16(*(*uint8_t)(unsafe.Pointer(p)))))
}

func ReadLittleEndian32(tls *libc.TLS, p uintptr) (r uint32_t) {
	return uint32(*(*uint8_t)(unsafe.Pointer(p + 3)))<<libc.Int32FromInt32(24) | uint32(*(*uint8_t)(unsafe.Pointer(p + 2)))<<libc.Int32FromInt32(16) | uint32(*(*uint8_t)(unsafe.Pointer(p + 1)))<<libc.Int32FromInt32(8) | uint32(*(*uint8_t)(unsafe.Pointer(p)))
}

func clear_endpoint(tls *libc.TLS, endpoint uintptr) {
	libc.Xfree(tls, (*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).Fextra)
}

func parse_endpoint(tls *libc.TLS, ctx uintptr, endpoint uintptr, buffer uintptr, size int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var begin, extra, header uintptr
	var len1 ptrdiff_t
	var parsed int32
	_, _, _, _, _ = begin, extra, header, len1, parsed
	parsed = 0
	if size < int32(2) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8502, libc.VaList(bp+8, size, int32(2)))
		return int32(LIBUSB_ERROR_IO)
	}
	header = buffer
	if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) != int32(LIBUSB_DT_ENDPOINT) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8539, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType), int32(LIBUSB_DT_ENDPOINT)))
		return parsed
	} else {
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(7) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8582, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) > size {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8612, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				return parsed
			}
		}
	}
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbEndpointAddress = *(*uint8_t)(unsafe.Pointer(buffer + 2))
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbmAttributes = *(*uint8_t)(unsafe.Pointer(buffer + 3))
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FwMaxPacketSize = ReadLittleEndian16(tls, buffer+4)
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbInterval = *(*uint8_t)(unsafe.Pointer(buffer + 6))
	if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) >= int32(9) {
		(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbRefresh = *(*uint8_t)(unsafe.Pointer(buffer + 7))
		(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).FbSynchAddress = *(*uint8_t)(unsafe.Pointer(buffer + 8))
	}
	buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
	size -= libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
	parsed += libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
	begin = buffer
	for size >= int32(2) {
		header = buffer
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(2) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8649, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) > size {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8680, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				return parsed
			}
		}
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_ENDPOINT) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_INTERFACE) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_CONFIG) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_DEVICE) {
			break
		}
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__106)), __ccgo_ts+8711, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType)))
		buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		size -= libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		parsed += libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
	}
	len1 = int64(buffer) - int64(begin)
	if len1 <= 0 {
		return parsed
	}
	extra = libc.Xmalloc(tls, libc.Uint64FromInt64(len1))
	if !(extra != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	libc.Xmemcpy(tls, extra, begin, libc.Uint64FromInt64(len1))
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).Fextra = extra
	(*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).Fextra_length = int32(len1)
	return parsed
}

var __func__106 = [15]uint8{'p', 'a', 'r', 's', 'e', '_', 'e', 'n', 'd', 'p', 'o', 'i', 'n', 't'}

func clear_interface(tls *libc.TLS, usb_interface uintptr) {
	var i int32
	var ifp uintptr
	var j uint8_t
	_, _, _ = i, ifp, j
	if (*libusb_interface)(unsafe.Pointer(usb_interface)).Faltsetting != 0 {
		i = 0
		for {
			if !(i < (*libusb_interface)(unsafe.Pointer(usb_interface)).Fnum_altsetting) {
				break
			}
			ifp = (*libusb_interface)(unsafe.Pointer(usb_interface)).Faltsetting + uintptr(i)*40
			libc.Xfree(tls, (*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra)
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra = libc.UintptrFromInt32(0)
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra_length = 0
			if (*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fendpoint != 0 {
				j = uint8(0)
				for {
					if !(libc.Int32FromUint8(j) < libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints)) {
						break
					}
					clear_endpoint(tls, (*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fendpoint+uintptr(j)*32)
					goto _2
				_2:
					;
					j++
				}
			}
			libc.Xfree(tls, (*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fendpoint)
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fendpoint = libc.UintptrFromInt32(0)
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints = uint8(0)
			goto _1
		_1:
			;
			i++
		}
	}
	libc.Xfree(tls, (*libusb_interface)(unsafe.Pointer(usb_interface)).Faltsetting)
	(*libusb_interface)(unsafe.Pointer(usb_interface)).Faltsetting = libc.UintptrFromInt32(0)
	(*libusb_interface)(unsafe.Pointer(usb_interface)).Fnum_altsetting = 0
}

func parse_interface(tls *libc.TLS, ctx uintptr, usb_interface uintptr, buffer uintptr, size int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var altsetting, begin, endpoint, extra, header, if_desc, ifp uintptr
	var i uint8_t
	var interface_number, parsed, r int32
	var len1 ptrdiff_t
	_, _, _, _, _, _, _, _, _, _, _, _ = altsetting, begin, endpoint, extra, header, i, if_desc, ifp, interface_number, len1, parsed, r
	parsed = 0
	interface_number = -int32(1)
	for size >= int32(9) {
		altsetting = libc.Xrealloc(tls, (*libusb_interface)(unsafe.Pointer(usb_interface)).Faltsetting, uint64(40)*libc.Uint64FromInt32((*libusb_interface)(unsafe.Pointer(usb_interface)).Fnum_altsetting+libc.Int32FromInt32(1)))
		if !(altsetting != 0) {
			r = int32(LIBUSB_ERROR_NO_MEM)
			goto err
		}
		(*libusb_interface)(unsafe.Pointer(usb_interface)).Faltsetting = altsetting
		ifp = altsetting + uintptr((*libusb_interface)(unsafe.Pointer(usb_interface)).Fnum_altsetting)*40
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbInterfaceNumber = *(*uint8_t)(unsafe.Pointer(buffer + 2))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbAlternateSetting = *(*uint8_t)(unsafe.Pointer(buffer + 3))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints = *(*uint8_t)(unsafe.Pointer(buffer + 4))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbInterfaceClass = *(*uint8_t)(unsafe.Pointer(buffer + 5))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbInterfaceSubClass = *(*uint8_t)(unsafe.Pointer(buffer + 6))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbInterfaceProtocol = *(*uint8_t)(unsafe.Pointer(buffer + 7))
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FiInterface = *(*uint8_t)(unsafe.Pointer(buffer + 8))
		if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbDescriptorType) != int32(LIBUSB_DT_INTERFACE) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__107)), __ccgo_ts+8539, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbDescriptorType), int32(LIBUSB_DT_INTERFACE)))
			return parsed
		} else {
			if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength) < int32(9) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__107)), __ccgo_ts+8736, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength)))
				r = int32(LIBUSB_ERROR_IO)
				goto err
			} else {
				if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength) > size {
					usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__107)), __ccgo_ts+8767, libc.VaList(bp+8, size, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength)))
					return parsed
				} else {
					if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints) > int32(32) {
						usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__107)), __ccgo_ts+8800, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints)))
						r = int32(LIBUSB_ERROR_IO)
						goto err
					}
				}
			}
		}
		(*libusb_interface)(unsafe.Pointer(usb_interface)).Fnum_altsetting++
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra = libc.UintptrFromInt32(0)
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra_length = 0
		(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fendpoint = libc.UintptrFromInt32(0)
		if interface_number == -int32(1) {
			interface_number = libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbInterfaceNumber)
		}
		buffer += uintptr((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength)
		parsed += libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength)
		size -= libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbLength)
		begin = buffer
		for size >= int32(2) {
			header = buffer
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(2) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__107)), __ccgo_ts+8824, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				r = int32(LIBUSB_ERROR_IO)
				goto err
			} else {
				if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) > size {
					usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__107)), __ccgo_ts+8857, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
					return parsed
				}
			}
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_INTERFACE) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_ENDPOINT) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_CONFIG) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_DEVICE) {
				break
			}
			buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
			parsed += libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
			size -= libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		}
		len1 = int64(buffer) - int64(begin)
		if len1 > 0 {
			extra = libc.Xmalloc(tls, libc.Uint64FromInt64(len1))
			if !(extra != 0) {
				r = int32(LIBUSB_ERROR_NO_MEM)
				goto err
			}
			libc.Xmemcpy(tls, extra, begin, libc.Uint64FromInt64(len1))
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra = extra
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fextra_length = int32(len1)
		}
		if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints) > 0 {
			endpoint = libc.Xcalloc(tls, uint64((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints), uint64(32))
			if !(endpoint != 0) {
				r = int32(LIBUSB_ERROR_NO_MEM)
				goto err
			}
			(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).Fendpoint = endpoint
			i = uint8(0)
			for {
				if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints)) {
					break
				}
				r = parse_endpoint(tls, ctx, endpoint+uintptr(i)*32, buffer, size)
				if r < 0 {
					goto err
				}
				if r == 0 {
					(*libusb_interface_descriptor)(unsafe.Pointer(ifp)).FbNumEndpoints = i
					break
				}
				buffer += uintptr(r)
				parsed += r
				size -= r
				goto _1
			_1:
				;
				i++
			}
		}
		if_desc = buffer
		if size < int32(9) || libc.Int32FromUint8((*usbi_interface_descriptor)(unsafe.Pointer(if_desc)).FbDescriptorType) != int32(LIBUSB_DT_INTERFACE) || libc.Int32FromUint8((*usbi_interface_descriptor)(unsafe.Pointer(if_desc)).FbInterfaceNumber) != interface_number {
			return parsed
		}
	}
	return parsed
	goto err
err:
	;
	clear_interface(tls, usb_interface)
	return r
}

var __func__107 = [16]uint8{'p', 'a', 'r', 's', 'e', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e'}

func clear_configuration(tls *libc.TLS, config uintptr) {
	var i uint8_t
	_ = i
	if (*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1 != 0 {
		i = uint8(0)
		for {
			if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces)) {
				break
			}
			clear_interface(tls, (*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1+uintptr(i)*16)
			goto _1
		_1:
			;
			i++
		}
	}
	libc.Xfree(tls, (*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1)
	(*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1 = libc.UintptrFromInt32(0)
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces = uint8(0)
	libc.Xfree(tls, (*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra)
	(*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra = libc.UintptrFromInt32(0)
	(*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra_length = 0
}

func parse_configuration(tls *libc.TLS, ctx uintptr, config uintptr, buffer uintptr, size int32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var begin, extra, header, usb_interface uintptr
	var i uint8_t
	var len1 ptrdiff_t
	var r int32
	_, _, _, _, _, _, _ = begin, extra, header, i, len1, r, usb_interface
	if size < int32(9) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+8890, libc.VaList(bp+8, size, int32(9)))
		return int32(LIBUSB_ERROR_IO)
	}
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FwTotalLength = ReadLittleEndian16(tls, buffer+2)
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces = *(*uint8_t)(unsafe.Pointer(buffer + 4))
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FbConfigurationValue = *(*uint8_t)(unsafe.Pointer(buffer + 5))
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FiConfiguration = *(*uint8_t)(unsafe.Pointer(buffer + 6))
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FbmAttributes = *(*uint8_t)(unsafe.Pointer(buffer + 7))
	(*libusb_config_descriptor)(unsafe.Pointer(config)).FMaxPower = *(*uint8_t)(unsafe.Pointer(buffer + 8))
	if libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbDescriptorType) != int32(LIBUSB_DT_CONFIG) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+8539, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbDescriptorType), int32(LIBUSB_DT_CONFIG)))
		return int32(LIBUSB_ERROR_IO)
	} else {
		if libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength) < int32(9) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+8925, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength)))
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength) > size {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+8953, libc.VaList(bp+8, size, libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength)))
				return int32(LIBUSB_ERROR_IO)
			} else {
				if libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces) > int32(32) {
					usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+8988, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces)))
					return int32(LIBUSB_ERROR_IO)
				}
			}
		}
	}
	usb_interface = libc.Xcalloc(tls, uint64((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces), uint64(16))
	if !(usb_interface != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_config_descriptor)(unsafe.Pointer(config)).Finterface1 = usb_interface
	buffer += uintptr((*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength)
	size -= libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbLength)
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces)) {
			break
		}
		begin = buffer
		for size >= int32(2) {
			header = buffer
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(2) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+9013, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				r = int32(LIBUSB_ERROR_IO)
				goto err
			} else {
				if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) > size {
					usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+9048, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
					(*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces = i
					return size
				}
			}
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_ENDPOINT) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_INTERFACE) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_CONFIG) || libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) == int32(LIBUSB_DT_DEVICE) {
				break
			}
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__108)), __ccgo_ts+8711, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType)))
			buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
			size -= libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		}
		len1 = int64(buffer) - int64(begin)
		if len1 > 0 {
			extra = libc.Xrealloc(tls, (*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra, libc.Uint64FromInt32((*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra_length)+libc.Uint64FromInt64(len1))
			if !(extra != 0) {
				r = int32(LIBUSB_ERROR_NO_MEM)
				goto err
			}
			libc.Xmemcpy(tls, extra+uintptr((*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra_length), begin, libc.Uint64FromInt64(len1))
			(*libusb_config_descriptor)(unsafe.Pointer(config)).Fextra = extra
			*(*int32)(unsafe.Pointer(config + 32)) += int32(len1)
		}
		r = parse_interface(tls, ctx, usb_interface+uintptr(i)*16, buffer, size)
		if r < 0 {
			goto err
		}
		if r == 0 {
			(*libusb_config_descriptor)(unsafe.Pointer(config)).FbNumInterfaces = i
			break
		}
		buffer += uintptr(r)
		size -= r
		goto _1
	_1:
		;
		i++
	}
	return size
	goto err
err:
	;
	clear_configuration(tls, config)
	return r
}

var __func__108 = [20]uint8{'p', 'a', 'r', 's', 'e', '_', 'c', 'o', 'n', 'f', 'i', 'g', 'u', 'r', 'a', 't', 'i', 'o', 'n'}

func raw_desc_to_config(tls *libc.TLS, ctx uintptr, buf uintptr, size int32, config uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _config uintptr
	var r int32
	_, _ = _config, r
	_config = libc.Xcalloc(tls, uint64(1), uint64(40))
	if !(_config != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = parse_configuration(tls, ctx, _config, buf, size)
	if r < 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__109)), __ccgo_ts+9083, libc.VaList(bp+8, r))
		libc.Xfree(tls, _config)
		return r
	} else {
		if r > 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__109)), __ccgo_ts+9124, libc.VaList(bp+8, r))
		}
	}
	*(*uintptr)(unsafe.Pointer(config)) = _config
	return int32(LIBUSB_SUCCESS)
}

var __func__109 = [19]uint8{'r', 'a', 'w', '_', 'd', 'e', 's', 'c', '_', 't', 'o', '_', 'c', 'o', 'n', 'f', 'i', 'g'}

func get_active_config_descriptor(tls *libc.TLS, dev uintptr, buffer uintptr, size size_t) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var r int32
	_ = r
	r = (*(*func(*libc.TLS, uintptr, uintptr, size_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_active_config_descriptor})))(tls, dev, buffer, size)
	if r < 0 {
		return r
	}
	if r < int32(9) {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__110)), __ccgo_ts+8890, libc.VaList(bp+8, r, int32(9)))
		return int32(LIBUSB_ERROR_IO)
	} else {
		if r != libc.Int32FromUint64(size) {
			usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__110)), __ccgo_ts+8890, libc.VaList(bp+8, r, libc.Int32FromUint64(size)))
		}
	}
	return r
}

var __func__110 = [29]uint8{'g', 'e', 't', '_', 'a', 'c', 't', 'i', 'v', 'e', '_', 'c', 'o', 'n', 'f', 'i', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func get_config_descriptor(tls *libc.TLS, dev uintptr, config_idx uint8_t, buffer uintptr, size size_t) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var r int32
	_ = r
	r = (*(*func(*libc.TLS, uintptr, uint8_t, uintptr, size_t) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_config_descriptor})))(tls, dev, config_idx, buffer, size)
	if r < 0 {
		return r
	}
	if r < int32(9) {
		usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__111)), __ccgo_ts+8890, libc.VaList(bp+8, r, int32(9)))
		return int32(LIBUSB_ERROR_IO)
	} else {
		if r != libc.Int32FromUint64(size) {
			usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__111)), __ccgo_ts+8890, libc.VaList(bp+8, r, libc.Int32FromUint64(size)))
		}
	}
	return r
}

var __func__111 = [22]uint8{'g', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_get_device_descriptor(tls *libc.TLS, dev uintptr, desc uintptr) (r int32) {
	usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__112)), __ccgo_ts+6578, 0)
	*(*libusb_device_descriptor)(unsafe.Pointer(desc)) = (*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor
	return 0
}

var __func__112 = [29]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_get_active_config_descriptor(tls *libc.TLS, dev uintptr, config uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var buf uintptr
	var config_len uint16_t
	var r int32
	var _ /* _config at bp+0 */ usbi_config_desc_buf
	_, _, _ = buf, config_len, r
	r = get_active_config_descriptor(tls, dev, bp, uint64(9))
	if r < 0 {
		return r
	}
	config_len = libusb_cpu_to_le16(tls, (*(*usbi_configuration_descriptor)(unsafe.Pointer(bp))).FwTotalLength)
	buf = libc.Xmalloc(tls, uint64(config_len))
	if !(buf != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = get_active_config_descriptor(tls, dev, buf, uint64(config_len))
	if r >= 0 {
		r = raw_desc_to_config(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, buf, r, config)
	}
	libc.Xfree(tls, buf)
	return r
}

func libusb_get_config_descriptor(tls *libc.TLS, dev uintptr, config_index uint8_t, config uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buf uintptr
	var config_len uint16_t
	var r int32
	var _ /* _config at bp+0 */ usbi_config_desc_buf
	_, _, _ = buf, config_len, r
	usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__113)), __ccgo_ts+9163, libc.VaList(bp+24, libc.Int32FromUint8(config_index)))
	if libc.Int32FromUint8(config_index) >= libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations) {
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	r = get_config_descriptor(tls, dev, config_index, bp, uint64(9))
	if r < 0 {
		return r
	}
	config_len = libusb_cpu_to_le16(tls, (*(*usbi_configuration_descriptor)(unsafe.Pointer(bp))).FwTotalLength)
	buf = libc.Xmalloc(tls, uint64(config_len))
	if !(buf != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = get_config_descriptor(tls, dev, config_index, buf, uint64(config_len))
	if r >= 0 {
		r = raw_desc_to_config(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, buf, r, config)
	}
	libc.Xfree(tls, buf)
	return r
}

var __func__113 = [29]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_get_config_descriptor_by_value(tls *libc.TLS, dev uintptr, bConfigurationValue uint8_t, config uintptr) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var idx uint8_t
	var r int32
	var _ /* _config at bp+8 */ usbi_config_desc_buf
	var _ /* buf at bp+0 */ uintptr
	_, _ = idx, r
	if usbi_backend.Fget_config_descriptor_by_value != 0 {
		r = (*(*func(*libc.TLS, uintptr, uint8_t, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_config_descriptor_by_value})))(tls, dev, bConfigurationValue, bp)
		if r < 0 {
			return r
		}
		return raw_desc_to_config(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, *(*uintptr)(unsafe.Pointer(bp)), r, config)
	}
	usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__114)), __ccgo_ts+9172, libc.VaList(bp+32, libc.Int32FromUint8(bConfigurationValue)))
	idx = uint8(0)
	for {
		if !(libc.Int32FromUint8(idx) < libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations)) {
			break
		}
		r = get_config_descriptor(tls, dev, idx, bp+8, uint64(9))
		if r < 0 {
			return r
		}
		if libc.Int32FromUint8((*(*usbi_configuration_descriptor)(unsafe.Pointer(bp + 8))).FbConfigurationValue) == libc.Int32FromUint8(bConfigurationValue) {
			return libusb_get_config_descriptor(tls, dev, idx, config)
		}
		goto _1
	_1:
		;
		idx++
	}
	return int32(LIBUSB_ERROR_NOT_FOUND)
}

var __func__114 = [38]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'c', 'o', 'n', 'f', 'i', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', '_', 'b', 'y', '_', 'v', 'a', 'l', 'u', 'e'}

func libusb_free_config_descriptor(tls *libc.TLS, config uintptr) {
	if !(config != 0) {
		return
	}
	clear_configuration(tls, config)
	libc.Xfree(tls, config)
}

func libusb_get_ss_endpoint_companion_descriptor(tls *libc.TLS, ctx uintptr, endpoint uintptr, ep_comp uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buffer, header uintptr
	var size int32
	_, _, _ = buffer, header, size
	buffer = (*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).Fextra
	size = (*libusb_endpoint_descriptor)(unsafe.Pointer(endpoint)).Fextra_length
	*(*uintptr)(unsafe.Pointer(ep_comp)) = libc.UintptrFromInt32(0)
	for size >= int32(2) {
		header = buffer
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) != int32(LIBUSB_DT_SS_ENDPOINT_COMPANION) {
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(2) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__115)), __ccgo_ts+9181, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				return int32(LIBUSB_ERROR_IO)
			}
			buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
			size -= libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
			continue
		} else {
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(6) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__115)), __ccgo_ts+9210, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				return int32(LIBUSB_ERROR_IO)
			} else {
				if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) > size {
					usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__115)), __ccgo_ts+9244, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
					return int32(LIBUSB_ERROR_IO)
				}
			}
		}
		*(*uintptr)(unsafe.Pointer(ep_comp)) = libc.Xmalloc(tls, uint64(6))
		if !(*(*uintptr)(unsafe.Pointer(ep_comp)) != 0) {
			return int32(LIBUSB_ERROR_NO_MEM)
		}
		(*libusb_ss_endpoint_companion_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ep_comp)))).FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
		(*libusb_ss_endpoint_companion_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ep_comp)))).FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
		(*libusb_ss_endpoint_companion_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ep_comp)))).FbMaxBurst = *(*uint8_t)(unsafe.Pointer(buffer + 2))
		(*libusb_ss_endpoint_companion_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ep_comp)))).FbmAttributes = *(*uint8_t)(unsafe.Pointer(buffer + 3))
		(*libusb_ss_endpoint_companion_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ep_comp)))).FwBytesPerInterval = ReadLittleEndian16(tls, buffer+4)
		return int32(LIBUSB_SUCCESS)
	}
	return int32(LIBUSB_ERROR_NOT_FOUND)
}

var __func__115 = [44]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 's', 's', '_', 'e', 'n', 'd', 'p', 'o', 'i', 'n', 't', '_', 'c', 'o', 'm', 'p', 'a', 'n', 'i', 'o', 'n', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_free_ss_endpoint_companion_descriptor(tls *libc.TLS, ep_comp uintptr) {
	libc.Xfree(tls, ep_comp)
}

func parse_bos(tls *libc.TLS, ctx uintptr, bos uintptr, buffer uintptr, size int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _bos, bos_desc, header uintptr
	var i uint8_t
	_, _, _, _ = _bos, bos_desc, header, i
	if size < int32(5) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+9277, libc.VaList(bp+8, size, int32(5)))
		return int32(LIBUSB_ERROR_IO)
	}
	bos_desc = buffer
	if libc.Int32FromUint8((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbDescriptorType) != int32(LIBUSB_DT_BOS) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+8539, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbDescriptorType), int32(LIBUSB_DT_BOS)))
		return int32(LIBUSB_ERROR_IO)
	} else {
		if libc.Int32FromUint8((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbLength) < int32(5) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+9309, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbLength)))
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbLength) > size {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+9334, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbLength)))
				return int32(LIBUSB_ERROR_IO)
			}
		}
	}
	_bos = libc.Xcalloc(tls, uint64(1), uint64(8)+uint64((*usbi_bos_descriptor)(unsafe.Pointer(bos_desc)).FbNumDeviceCaps)*uint64(8))
	if !(_bos != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
	(*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
	(*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FwTotalLength = ReadLittleEndian16(tls, buffer+2)
	(*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbNumDeviceCaps = *(*uint8_t)(unsafe.Pointer(buffer + 4))
	buffer += uintptr((*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbLength)
	size -= libc.Int32FromUint8((*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbLength)
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbNumDeviceCaps)) {
			break
		}
		if size < int32(3) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+9366, libc.VaList(bp+8, size, int32(3)))
			break
		}
		header = buffer
		if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType) != int32(LIBUSB_DT_DEVICE_CAPABILITY) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+8539, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbDescriptorType), int32(LIBUSB_DT_DEVICE_CAPABILITY)))
			break
		} else {
			if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) < int32(3) {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+9402, libc.VaList(bp+8, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
				libusb_free_bos_descriptor(tls, _bos)
				return int32(LIBUSB_ERROR_IO)
			} else {
				if libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength) > size {
					usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__116)), __ccgo_ts+9431, libc.VaList(bp+8, size, libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)))
					break
				}
			}
		}
		*(*uintptr)(unsafe.Pointer(_bos + 8 + uintptr(i)*8)) = libc.Xmalloc(tls, uint64((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength))
		if !(*(*uintptr)(unsafe.Pointer(_bos + 8 + uintptr(i)*8)) != 0) {
			libusb_free_bos_descriptor(tls, _bos)
			return int32(LIBUSB_ERROR_NO_MEM)
		}
		libc.Xmemcpy(tls, *(*uintptr)(unsafe.Pointer(_bos + 8 + uintptr(i)*8)), buffer, uint64((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength))
		buffer += uintptr((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		size -= libc.Int32FromUint8((*usbi_descriptor_header)(unsafe.Pointer(header)).FbLength)
		goto _1
	_1:
		;
		i++
	}
	(*libusb_bos_descriptor)(unsafe.Pointer(_bos)).FbNumDeviceCaps = i
	*(*uintptr)(unsafe.Pointer(bos)) = _bos
	return int32(LIBUSB_SUCCESS)
}

var __func__116 = [10]uint8{'p', 'a', 'r', 's', 'e', '_', 'b', 'o', 's'}

func libusb_get_bos_descriptor(tls *libc.TLS, dev_handle uintptr, bos uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var bos_data, ctx, v1, v2 uintptr
	var bos_len uint16_t
	var r int32
	var _ /* _bos at bp+0 */ usbi_bos_desc_buf
	_, _, _, _, _, _ = bos_data, bos_len, ctx, r, v1, v2
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	r = libusb_get_descriptor(tls, dev_handle, uint8(LIBUSB_DT_BOS), uint8(0), bp, int32(5))
	if r < 0 {
		if r != int32(LIBUSB_ERROR_PIPE) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__117)), __ccgo_ts+9467, libc.VaList(bp+16, r))
		}
		return r
	}
	if r < int32(5) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__117)), __ccgo_ts+9491, libc.VaList(bp+16, r, int32(5)))
		return int32(LIBUSB_ERROR_IO)
	}
	bos_len = libusb_cpu_to_le16(tls, (*(*usbi_bos_descriptor)(unsafe.Pointer(bp))).FwTotalLength)
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__117)), __ccgo_ts+9512, libc.VaList(bp+16, libc.Int32FromUint16(bos_len), libc.Int32FromUint8((*(*usbi_bos_descriptor)(unsafe.Pointer(bp))).FbNumDeviceCaps)))
	bos_data = libc.Xcalloc(tls, uint64(1), uint64(bos_len))
	if !(bos_data != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = libusb_get_descriptor(tls, dev_handle, uint8(LIBUSB_DT_BOS), uint8(0), bos_data, libc.Int32FromUint16(bos_len))
	if r >= 0 {
		if r != libc.Int32FromUint16(bos_len) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__117)), __ccgo_ts+9565, libc.VaList(bp+16, r, libc.Int32FromUint16(bos_len)))
		}
		if dev_handle != 0 {
			v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
		} else {
			v2 = libc.UintptrFromInt32(0)
		}
		r = parse_bos(tls, v2, bos, bos_data, r)
	} else {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__117)), __ccgo_ts+9467, libc.VaList(bp+16, r))
	}
	libc.Xfree(tls, bos_data)
	return r
}

var __func__117 = [26]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'b', 'o', 's', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_free_bos_descriptor(tls *libc.TLS, bos uintptr) {
	var i uint8_t
	_ = i
	if !(bos != 0) {
		return
	}
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*libusb_bos_descriptor)(unsafe.Pointer(bos)).FbNumDeviceCaps)) {
			break
		}
		libc.Xfree(tls, *(*uintptr)(unsafe.Pointer(bos + 8 + uintptr(i)*8)))
		goto _1
	_1:
		;
		i++
	}
	libc.Xfree(tls, bos)
}

func libusb_get_usb_2_0_extension_descriptor(tls *libc.TLS, ctx uintptr, dev_cap uintptr, usb_2_0_extension uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _usb_2_0_extension uintptr
	_ = _usb_2_0_extension
	if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType) != int32(LIBUSB_BT_USB_2_0_EXTENSION) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__118)), __ccgo_ts+9586, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType), int32(LIBUSB_BT_USB_2_0_EXTENSION)))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	} else {
		if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) < int32(7) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__118)), __ccgo_ts+9637, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength), int32(7)))
			return int32(LIBUSB_ERROR_IO)
		}
	}
	_usb_2_0_extension = libc.Xmalloc(tls, uint64(8))
	if !(_usb_2_0_extension != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_usb_2_0_extension_descriptor)(unsafe.Pointer(_usb_2_0_extension)).FbLength = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength
	(*libusb_usb_2_0_extension_descriptor)(unsafe.Pointer(_usb_2_0_extension)).FbDescriptorType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDescriptorType
	(*libusb_usb_2_0_extension_descriptor)(unsafe.Pointer(_usb_2_0_extension)).FbDevCapabilityType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType
	(*libusb_usb_2_0_extension_descriptor)(unsafe.Pointer(_usb_2_0_extension)).FbmAttributes = ReadLittleEndian32(tls, dev_cap+3)
	*(*uintptr)(unsafe.Pointer(usb_2_0_extension)) = _usb_2_0_extension
	return int32(LIBUSB_SUCCESS)
}

var __func__118 = [40]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'u', 's', 'b', '_', '2', '_', '0', '_', 'e', 'x', 't', 'e', 'n', 's', 'i', 'o', 'n', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_free_usb_2_0_extension_descriptor(tls *libc.TLS, usb_2_0_extension uintptr) {
	libc.Xfree(tls, usb_2_0_extension)
}

func libusb_get_ss_usb_device_capability_descriptor(tls *libc.TLS, ctx uintptr, dev_cap uintptr, ss_usb_device_cap uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ss_usb_device_cap uintptr
	_ = _ss_usb_device_cap
	if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType) != int32(LIBUSB_BT_SS_USB_DEVICE_CAPABILITY) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__119)), __ccgo_ts+9586, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType), int32(LIBUSB_BT_SS_USB_DEVICE_CAPABILITY)))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	} else {
		if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) < int32(10) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__119)), __ccgo_ts+9637, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength), int32(10)))
			return int32(LIBUSB_ERROR_IO)
		}
	}
	_ss_usb_device_cap = libc.Xmalloc(tls, uint64(10))
	if !(_ss_usb_device_cap != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbLength = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbDescriptorType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDescriptorType
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbDevCapabilityType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbmAttributes = *(*uint8_t)(unsafe.Pointer(dev_cap + 3))
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FwSpeedSupported = ReadLittleEndian16(tls, dev_cap+3+1)
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbFunctionalitySupport = *(*uint8_t)(unsafe.Pointer(dev_cap + 3 + 3))
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbU1DevExitLat = *(*uint8_t)(unsafe.Pointer(dev_cap + 3 + 4))
	(*libusb_ss_usb_device_capability_descriptor)(unsafe.Pointer(_ss_usb_device_cap)).FbU2DevExitLat = ReadLittleEndian16(tls, dev_cap+3+5)
	*(*uintptr)(unsafe.Pointer(ss_usb_device_cap)) = _ss_usb_device_cap
	return int32(LIBUSB_SUCCESS)
}

var __func__119 = [47]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 's', 's', '_', 'u', 's', 'b', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'c', 'a', 'p', 'a', 'b', 'i', 'l', 'i', 't', 'y', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

type internal_ssplus_capability_descriptor = struct {
	FbLength               uint8_t
	FbDescriptorType       uint8_t
	FbDevCapabilityType    uint8_t
	FbReserved             uint8_t
	FbmAttributes          uint32_t
	FwFunctionalitySupport uint16_t
	FwReserved             uint16_t
}

func libusb_get_ssplus_usb_device_capability_descriptor(tls *libc.TLS, ctx uintptr, dev_cap uintptr, ssplus_usb_device_cap uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ssplus_cap, base, dev_capability_data uintptr
	var attr uint32_t
	var i, numSublikSpeedAttributes uint8_t
	var parsedDescriptor internal_ssplus_capability_descriptor
	var v2, v3, v4 int32
	_, _, _, _, _, _, _, _, _, _ = _ssplus_cap, attr, base, dev_capability_data, i, numSublikSpeedAttributes, parsedDescriptor, v2, v3, v4
	if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType) != int32(LIBUSB_BT_SUPERSPEED_PLUS_CAPABILITY) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__120)), __ccgo_ts+9586, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType), int32(LIBUSB_BT_SUPERSPEED_PLUS_CAPABILITY)))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	} else {
		if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) < int32(12) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__120)), __ccgo_ts+9637, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength), int32(12)))
			return int32(LIBUSB_ERROR_IO)
		}
	}
	dev_capability_data = dev_cap + 3
	parsedDescriptor.FbLength = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength
	parsedDescriptor.FbDescriptorType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDescriptorType
	parsedDescriptor.FbDevCapabilityType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType
	parsedDescriptor.FbReserved = *(*uint8_t)(unsafe.Pointer(dev_capability_data))
	parsedDescriptor.FbmAttributes = ReadLittleEndian32(tls, dev_capability_data+1)
	parsedDescriptor.FwFunctionalitySupport = ReadLittleEndian16(tls, dev_capability_data+5)
	parsedDescriptor.FwReserved = ReadLittleEndian16(tls, dev_capability_data+7)
	numSublikSpeedAttributes = uint8(parsedDescriptor.FbmAttributes&uint32(0xF) + uint32(1))
	_ssplus_cap = libc.Xmalloc(tls, uint64(8)+uint64(numSublikSpeedAttributes)*uint64(24))
	if !(_ssplus_cap != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).FnumSublinkSpeedAttributes = numSublikSpeedAttributes
	(*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).FnumSublinkSpeedIDs = uint8(parsedDescriptor.FbmAttributes&uint32(0xF0)>>int32(4) + uint32(1))
	(*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).Fssid = libc.Uint8FromInt32(libc.Int32FromUint16(parsedDescriptor.FwFunctionalitySupport) & int32(0xF))
	(*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).FminRxLaneCount = libc.Uint8FromInt32(libc.Int32FromUint16(parsedDescriptor.FwFunctionalitySupport) & int32(0x0F00) >> int32(8))
	(*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).FminTxLaneCount = libc.Uint8FromInt32(libc.Int32FromUint16(parsedDescriptor.FwFunctionalitySupport) & int32(0xF000) >> int32(12))
	if uint64((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) < uint64(12)+uint64((*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).FnumSublinkSpeedAttributes)*uint64(4) {
		libc.Xfree(tls, _ssplus_cap)
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__120)), __ccgo_ts+9673, 0)
		return int32(LIBUSB_ERROR_IO)
	}
	base = dev_cap + uintptr(12)
	i = uint8(0)
	for {
		if !(libc.Int32FromUint8(i) < libc.Int32FromUint8((*libusb_ssplus_usb_device_capability_descriptor)(unsafe.Pointer(_ssplus_cap)).FnumSublinkSpeedAttributes)) {
			break
		}
		attr = ReadLittleEndian32(tls, base+uintptr(uint64(i)*libc.Uint64FromInt64(4)))
		(*(*libusb_ssplus_sublink_attribute)(unsafe.Pointer(_ssplus_cap + 8 + uintptr(i)*24))).Fssid = uint8(attr & uint32(0x0f))
		(*(*libusb_ssplus_sublink_attribute)(unsafe.Pointer(_ssplus_cap + 8 + uintptr(i)*24))).Fmantissa = uint16(attr >> int32(16))
		(*(*libusb_ssplus_sublink_attribute)(unsafe.Pointer(_ssplus_cap + 8 + uintptr(i)*24))).Fexponent = libc.Int32FromUint32(attr >> libc.Int32FromInt32(4) & uint32(0x3))
		if attr&uint32(0x40) != 0 {
			v2 = int32(LIBUSB_SSPLUS_ATTR_TYPE_ASYM)
		} else {
			v2 = int32(LIBUSB_SSPLUS_ATTR_TYPE_SYM)
		}
		(*(*libusb_ssplus_sublink_attribute)(unsafe.Pointer(_ssplus_cap + 8 + uintptr(i)*24))).Ftype1 = v2
		if attr&uint32(0x80) != 0 {
			v3 = int32(LIBUSB_SSPLUS_ATTR_DIR_TX)
		} else {
			v3 = int32(LIBUSB_SSPLUS_ATTR_DIR_RX)
		}
		(*(*libusb_ssplus_sublink_attribute)(unsafe.Pointer(_ssplus_cap + 8 + uintptr(i)*24))).Fdirection = v3
		if attr&uint32(0x4000) != 0 {
			v4 = int32(LIBUSB_SSPLUS_ATTR_PROT_SSPLUS)
		} else {
			v4 = int32(LIBUSB_SSPLUS_ATTR_PROT_SS)
		}
		(*(*libusb_ssplus_sublink_attribute)(unsafe.Pointer(_ssplus_cap + 8 + uintptr(i)*24))).Fprotocol = v4
		goto _1
	_1:
		;
		i++
	}
	*(*uintptr)(unsafe.Pointer(ssplus_usb_device_cap)) = _ssplus_cap
	return int32(LIBUSB_SUCCESS)
}

var __func__120 = [51]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 's', 's', 'p', 'l', 'u', 's', '_', 'u', 's', 'b', '_', 'd', 'e', 'v', 'i', 'c', 'e', '_', 'c', 'a', 'p', 'a', 'b', 'i', 'l', 'i', 't', 'y', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_free_ssplus_usb_device_capability_descriptor(tls *libc.TLS, ssplus_usb_device_cap uintptr) {
	libc.Xfree(tls, ssplus_usb_device_cap)
}

func libusb_free_ss_usb_device_capability_descriptor(tls *libc.TLS, ss_usb_device_cap uintptr) {
	libc.Xfree(tls, ss_usb_device_cap)
}

func libusb_get_container_id_descriptor(tls *libc.TLS, ctx uintptr, dev_cap uintptr, container_id uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _container_id uintptr
	_ = _container_id
	if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType) != int32(LIBUSB_BT_CONTAINER_ID) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__121)), __ccgo_ts+9586, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType), int32(LIBUSB_BT_CONTAINER_ID)))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	} else {
		if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) < int32(20) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__121)), __ccgo_ts+9637, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength), int32(20)))
			return int32(LIBUSB_ERROR_IO)
		}
	}
	_container_id = libc.Xmalloc(tls, uint64(20))
	if !(_container_id != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_container_id_descriptor)(unsafe.Pointer(_container_id)).FbLength = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength
	(*libusb_container_id_descriptor)(unsafe.Pointer(_container_id)).FbDescriptorType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDescriptorType
	(*libusb_container_id_descriptor)(unsafe.Pointer(_container_id)).FbDevCapabilityType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType
	(*libusb_container_id_descriptor)(unsafe.Pointer(_container_id)).FbReserved = *(*uint8_t)(unsafe.Pointer(dev_cap + 3))
	libc.Xmemcpy(tls, _container_id+4, dev_cap+3+1, uint64(16))
	*(*uintptr)(unsafe.Pointer(container_id)) = _container_id
	return int32(LIBUSB_SUCCESS)
}

var __func__121 = [35]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'c', 'o', 'n', 't', 'a', 'i', 'n', 'e', 'r', '_', 'i', 'd', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_free_container_id_descriptor(tls *libc.TLS, container_id uintptr) {
	libc.Xfree(tls, container_id)
}

func libusb_get_platform_descriptor(tls *libc.TLS, ctx uintptr, dev_cap uintptr, platform_descriptor uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _platform_descriptor, capability_data uintptr
	var capability_data_length size_t
	_, _, _ = _platform_descriptor, capability_data, capability_data_length
	if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType) != int32(LIBUSB_BT_PLATFORM_DESCRIPTOR) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__122)), __ccgo_ts+9586, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType), int32(LIBUSB_BT_PLATFORM_DESCRIPTOR)))
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	} else {
		if libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) < int32(20) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__122)), __ccgo_ts+9637, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength), int32(20)))
			return int32(LIBUSB_ERROR_IO)
		}
	}
	_platform_descriptor = libc.Xmalloc(tls, uint64((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength))
	if !(_platform_descriptor != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*libusb_platform_descriptor)(unsafe.Pointer(_platform_descriptor)).FbLength = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength
	(*libusb_platform_descriptor)(unsafe.Pointer(_platform_descriptor)).FbDescriptorType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDescriptorType
	(*libusb_platform_descriptor)(unsafe.Pointer(_platform_descriptor)).FbDevCapabilityType = (*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbDevCapabilityType
	(*libusb_platform_descriptor)(unsafe.Pointer(_platform_descriptor)).FbReserved = *(*uint8_t)(unsafe.Pointer(dev_cap + 3))
	libc.Xmemcpy(tls, _platform_descriptor+4, dev_cap+3+1, uint64(16))
	capability_data = dev_cap + 3 + uintptr(1) + uintptr(16)
	capability_data_length = libc.Uint64FromInt32(libc.Int32FromUint8((*libusb_bos_dev_capability_descriptor)(unsafe.Pointer(dev_cap)).FbLength) - (libc.Int32FromInt32(3) + libc.Int32FromInt32(1) + libc.Int32FromInt32(16)))
	libc.Xmemcpy(tls, _platform_descriptor+20, capability_data, capability_data_length)
	*(*uintptr)(unsafe.Pointer(platform_descriptor)) = _platform_descriptor
	return int32(LIBUSB_SUCCESS)
}

var __func__122 = [31]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'p', 'l', 'a', 't', 'f', 'o', 'r', 'm', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r'}

func libusb_free_platform_descriptor(tls *libc.TLS, platform_descriptor uintptr) {
	libc.Xfree(tls, platform_descriptor)
}

func libusb_get_string_descriptor_ascii(tls *libc.TLS, dev_handle uintptr, desc_index uint8_t, data uintptr, length int32) (r1 int32) {
	bp := tls.Alloc(288)
	defer tls.Free(288)
	var dest_max, idx, idx_max, r, src_max, v4 int32
	var langid, wdata uint16_t
	var v1, v2, v3 uintptr
	var _ /* str at bp+0 */ usbi_string_desc_buf
	_, _, _, _, _, _, _, _, _, _, _ = dest_max, idx, idx_max, langid, r, src_max, wdata, v1, v2, v3, v4
	if libc.Int32FromUint8(desc_index) == 0 {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	r = libusb_get_string_descriptor(tls, dev_handle, uint8(0), uint16(0), bp, int32(4))
	if r < 0 {
		return r
	} else {
		if r != int32(4) || libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength) < int32(4) || libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbDescriptorType) != int32(LIBUSB_DT_STRING) {
			if dev_handle != 0 {
				v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
			} else {
				v1 = libc.UintptrFromInt32(0)
			}
			usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__123)), __ccgo_ts+9750, 0)
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength)&int32(1) != 0 {
				if dev_handle != 0 {
					v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
				} else {
					v2 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__123)), __ccgo_ts+9788, libc.VaList(bp+264, libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength)))
			}
		}
	}
	langid = libusb_cpu_to_le16(tls, *(*uint16_t)(unsafe.Pointer(bp + 2)))
	r = libusb_get_string_descriptor(tls, dev_handle, desc_index, langid, bp, int32(255))
	if r < 0 {
		return r
	} else {
		if r < int32(2) || libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength) > r || libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbDescriptorType) != int32(LIBUSB_DT_STRING) {
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength)&int32(1) != 0 || libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength) != r {
				if dev_handle != 0 {
					v3 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
				} else {
					v3 = libc.UintptrFromInt32(0)
				}
				usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__123)), __ccgo_ts+9844, libc.VaList(bp+264, libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength), r))
			}
		}
	}
	dest_max = length - int32(1)
	src_max = (libc.Int32FromUint8((*(*usbi_string_desc_buf)(unsafe.Pointer(bp))).Fdesc.FbLength) - int32(1) - int32(1)) / int32(2)
	if dest_max < src_max {
		v4 = dest_max
	} else {
		v4 = src_max
	}
	idx_max = v4
	idx = 0
	for {
		if !(idx < idx_max) {
			break
		}
		wdata = libusb_cpu_to_le16(tls, *(*uint16_t)(unsafe.Pointer(bp + 2 + uintptr(idx)*2)))
		if libc.Int32FromUint16(wdata) < int32(0x80) {
			*(*uint8)(unsafe.Pointer(data + uintptr(idx))) = uint8(wdata)
		} else {
			*(*uint8)(unsafe.Pointer(data + uintptr(idx))) = uint8('?')
		}
		goto _5
	_5:
		;
		idx++
	}
	*(*uint8)(unsafe.Pointer(data + uintptr(idx))) = uint8(0)
	return idx
}

var __func__123 = [35]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 's', 't', 'r', 'i', 'n', 'g', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', '_', 'a', 's', 'c', 'i', 'i'}

func parse_iad_array(tls *libc.TLS, ctx uintptr, iad_array uintptr, buffer uintptr, size int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buf, iad uintptr
	var consumed, remaining int32
	var header usbi_descriptor_header
	var i uint8_t
	_, _, _, _, _, _ = buf, consumed, header, i, iad, remaining
	consumed = 0
	buf = buffer
	if size < int32(9) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__124)), __ccgo_ts+8890, libc.VaList(bp+8, size, int32(9)))
		return int32(LIBUSB_ERROR_IO)
	}
	(*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Flength = 0
	for consumed < size {
		header.FbLength = *(*uint8_t)(unsafe.Pointer(buf))
		header.FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buf + 1))
		if libc.Int32FromUint8(header.FbLength) < int32(2) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__124)), __ccgo_ts+9898, libc.VaList(bp+8, libc.Int32FromUint8(header.FbLength)))
			return int32(LIBUSB_ERROR_IO)
		} else {
			if libc.Int32FromUint8(header.FbLength) > size {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__124)), __ccgo_ts+8953, libc.VaList(bp+8, size, libc.Int32FromUint8(header.FbLength)))
				return int32(LIBUSB_ERROR_IO)
			}
		}
		if libc.Int32FromUint8(header.FbDescriptorType) == int32(LIBUSB_DT_INTERFACE_ASSOCIATION) {
			(*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Flength++
		}
		buf += uintptr(header.FbLength)
		consumed += libc.Int32FromUint8(header.FbLength)
	}
	(*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Fiad = libc.UintptrFromInt32(0)
	if (*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Flength > 0 {
		iad = libc.Xcalloc(tls, libc.Uint64FromInt32((*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Flength), uint64(8))
		if !(iad != 0) {
			return int32(LIBUSB_ERROR_NO_MEM)
		}
		(*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Fiad = iad
		remaining = size
		i = uint8(0)
		for cond := true; cond; cond = int32(1) != 0 {
			header.FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
			header.FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
			if libc.Int32FromUint8(header.FbDescriptorType) == int32(LIBUSB_DT_INTERFACE_ASSOCIATION) && remaining >= int32(8) {
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbLength = *(*uint8_t)(unsafe.Pointer(buffer))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbDescriptorType = *(*uint8_t)(unsafe.Pointer(buffer + 1))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbFirstInterface = *(*uint8_t)(unsafe.Pointer(buffer + 2))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbInterfaceCount = *(*uint8_t)(unsafe.Pointer(buffer + 3))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbFunctionClass = *(*uint8_t)(unsafe.Pointer(buffer + 4))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbFunctionSubClass = *(*uint8_t)(unsafe.Pointer(buffer + 5))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FbFunctionProtocol = *(*uint8_t)(unsafe.Pointer(buffer + 6))
				(*(*libusb_interface_association_descriptor)(unsafe.Pointer(iad + uintptr(i)*8))).FiFunction = *(*uint8_t)(unsafe.Pointer(buffer + 7))
				i++
			}
			remaining -= libc.Int32FromUint8(header.FbLength)
			if remaining < int32(2) {
				break
			}
			buffer += uintptr(header.FbLength)
		}
	}
	return int32(LIBUSB_SUCCESS)
}

var __func__124 = [16]uint8{'p', 'a', 'r', 's', 'e', '_', 'i', 'a', 'd', '_', 'a', 'r', 'r', 'a', 'y'}

func raw_desc_to_iad_array(tls *libc.TLS, ctx uintptr, buf uintptr, size int32, iad_array uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _iad_array uintptr
	var r int32
	_, _ = _iad_array, r
	_iad_array = libc.Xcalloc(tls, uint64(1), uint64(16))
	if !(_iad_array != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = parse_iad_array(tls, ctx, _iad_array, buf, size)
	if r < 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__125)), __ccgo_ts+9928, libc.VaList(bp+8, r))
		libc.Xfree(tls, _iad_array)
		return r
	}
	*(*uintptr)(unsafe.Pointer(iad_array)) = _iad_array
	return int32(LIBUSB_SUCCESS)
}

var __func__125 = [22]uint8{'r', 'a', 'w', '_', 'd', 'e', 's', 'c', '_', 't', 'o', '_', 'i', 'a', 'd', '_', 'a', 'r', 'r', 'a', 'y'}

func libusb_get_interface_association_descriptors(tls *libc.TLS, dev uintptr, config_index uint8_t, iad_array uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buf uintptr
	var config_len uint16_t
	var r int32
	var _ /* _config at bp+0 */ usbi_config_desc_buf
	_, _, _ = buf, config_len, r
	if !(iad_array != 0) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	usbi_log(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__126)), __ccgo_ts+9965, libc.VaList(bp+24, libc.Int32FromUint8(config_index)))
	if libc.Int32FromUint8(config_index) >= libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbNumConfigurations) {
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	r = get_config_descriptor(tls, dev, config_index, bp, uint64(9))
	if r < 0 {
		return r
	}
	config_len = libusb_cpu_to_le16(tls, (*(*usbi_configuration_descriptor)(unsafe.Pointer(bp))).FwTotalLength)
	buf = libc.Xmalloc(tls, uint64(config_len))
	if !(buf != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = get_config_descriptor(tls, dev, config_index, buf, uint64(config_len))
	if r >= 0 {
		r = raw_desc_to_iad_array(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, buf, r, iad_array)
	}
	libc.Xfree(tls, buf)
	return r
}

var __func__126 = [45]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'i', 'n', 't', 'e', 'r', 'f', 'a', 'c', 'e', '_', 'a', 's', 's', 'o', 'c', 'i', 'a', 't', 'i', 'o', 'n', '_', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', 's'}

func libusb_get_active_interface_association_descriptors(tls *libc.TLS, dev uintptr, iad_array uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var buf uintptr
	var config_len uint16_t
	var r int32
	var _ /* _config at bp+0 */ usbi_config_desc_buf
	_, _, _ = buf, config_len, r
	if !(iad_array != 0) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	r = get_active_config_descriptor(tls, dev, bp, uint64(9))
	if r < 0 {
		return r
	}
	config_len = libusb_cpu_to_le16(tls, (*(*usbi_configuration_descriptor)(unsafe.Pointer(bp))).FwTotalLength)
	buf = libc.Xmalloc(tls, uint64(config_len))
	if !(buf != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	r = get_active_config_descriptor(tls, dev, buf, uint64(config_len))
	if r >= 0 {
		r = raw_desc_to_iad_array(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, buf, r, iad_array)
	}
	libc.Xfree(tls, buf)
	return r
}

func libusb_free_interface_association_descriptors(tls *libc.TLS, iad_array uintptr) {
	if !(iad_array != 0) {
		return
	}
	if (*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Fiad != 0 {
		libc.Xfree(tls, (*libusb_interface_association_descriptor_array)(unsafe.Pointer(iad_array)).Fiad)
	}
	libc.Xfree(tls, iad_array)
}

func usbi_utf8_copy(tls *libc.TLS, tgt uintptr, src uintptr, tgt_size int32) (r int32) {
	var idx, k, v2 int32
	var s, t uintptr
	_, _, _, _, _ = idx, k, s, t, v2
	t = tgt
	s = src
	if libc.UintptrFromInt32(0) == src {
		if libc.UintptrFromInt32(0) != tgt && tgt_size > 0 {
			*(*uint8)(unsafe.Pointer(tgt)) = uint8(0)
		}
		return int32(1)
	}
	if libc.UintptrFromInt32(0) == tgt || tgt_size <= 0 {
		return libc.Int32FromUint64(libc.Xstrlen(tls, src) + libc.Uint64FromInt32(1))
	}
	k = 0
	k = 0
	for {
		if !(*(*uint8_t)(unsafe.Pointer(s + uintptr(k))) != 0) {
			break
		}
		if k < tgt_size {
			*(*uint8_t)(unsafe.Pointer(t + uintptr(k))) = *(*uint8_t)(unsafe.Pointer(s + uintptr(k)))
		} else {
			break
		}
		goto _1
	_1:
		;
		k++
	}
	if k >= tgt_size {
		idx = tgt_size - int32(1)
		for idx != 0 && int32(0x80) == libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(t + uintptr(idx))))&int32(0xC0) {
			idx--
		}
		*(*uint8_t)(unsafe.Pointer(t + uintptr(idx))) = uint8(0)
		return libc.Int32FromUint64(libc.Xstrlen(tls, src) + libc.Uint64FromInt32(1))
	} else {
		v2 = k
		k++
		*(*uint8_t)(unsafe.Pointer(t + uintptr(v2))) = uint8(0)
		return k
	}
	return r
}

func libusb_get_device_string(tls *libc.TLS, dev uintptr, string_type libusb_device_string_type, data uintptr, length int32) (r int32) {
	var rv int32
	var s uintptr
	_, _ = rv, s
	if libc.UintptrFromInt32(0) == dev {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	if string_type < 0 || string_type >= int32(LIBUSB_DEVICE_STRING_COUNT) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	if length < 0 {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	if libc.UintptrFromInt32(0) == data {
		length = 0
		data = libc.UintptrFromInt32(0)
	} else {
		if length > 0 {
			*(*uint8)(unsafe.Pointer(data)) = uint8(0)
		}
	}
	if libc.UintptrFromInt32(0) == *(*uintptr)(unsafe.Pointer(dev + 88 + uintptr(string_type)*8)) {
		if usbi_backend.Fget_device_string != 0 {
			s = libc.Xmalloc(tls, uint64(libc.Uint32FromUint32(384)))
			rv = (*(*func(*libc.TLS, uintptr, libusb_device_string_type, uintptr, int32) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fget_device_string})))(tls, dev, string_type, s, libc.Int32FromUint32(libc.Uint32FromUint32(384)))
			if rv < 0 {
				libc.Xfree(tls, s)
				return rv
			} else {
				*(*uintptr)(unsafe.Pointer(dev + 88 + uintptr(string_type)*8)) = s
			}
		} else {
			return int32(LIBUSB_ERROR_NOT_SUPPORTED)
		}
	}
	s = *(*uintptr)(unsafe.Pointer(dev + 88 + uintptr(string_type)*8))
	if libc.UintptrFromInt32(0) == s {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	return usbi_utf8_copy(tls, data, s, length)
}

func usbi_hotplug_init(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_store_ptr uintptr
	var _ /* __atomic_store_tmp at bp+0 */ usbi_atomic_t
	_ = __atomic_store_ptr
	if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
		return
	}
	usbi_mutex_init(tls, ctx+136)
	list_init(tls, ctx+176)
	(*libusb_context)(unsafe.Pointer(ctx)).Fnext_hotplug_cb_handle = int32(1)
	{
		__atomic_store_ptr = ctx + 200
		*(*usbi_atomic_t)(unsafe.Pointer(bp)) = int64(libc.Int32FromInt32(1))
		libc.X__atomic_storeInt64(tls, __atomic_store_ptr, bp, libc.Int32FromInt32(5))
	}
}

func usbi_recursively_remove_parents(tls *libc.TLS, dev uintptr, next_dev uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_load_ptr uintptr
	var v1 usbi_atomic_t
	var v2 bool
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _ = __atomic_load_ptr, v1, v2
	if dev != 0 && (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev != 0 {
		{
			__atomic_load_ptr = (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev
			libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
			v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
		}
		if v1 == int64(1) {
			if v2 = (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev != next_dev; !v2 {
				libc.X__assert_fail(tls, __ccgo_ts+9990, __ccgo_ts+10018, int32(175), uintptr(unsafe.Pointer(&__func__127)))
			}
			_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
			if (*libusb_device)(unsafe.Pointer((*libusb_device)(unsafe.Pointer(dev)).Fparent_dev)).Flist.Fnext != 0 && (*libusb_device)(unsafe.Pointer((*libusb_device)(unsafe.Pointer(dev)).Fparent_dev)).Flist.Fprev != 0 {
				list_del(tls, (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev+32)
			}
		}
		usbi_recursively_remove_parents(tls, (*libusb_device)(unsafe.Pointer(dev)).Fparent_dev, next_dev)
	}
}

var __func__127 = [32]uint8{'u', 's', 'b', 'i', '_', 'r', 'e', 'c', 'u', 'r', 's', 'i', 'v', 'e', 'l', 'y', '_', 'r', 'e', 'm', 'o', 'v', 'e', '_', 'p', 'a', 'r', 'e', 'n', 't', 's'}

func usbi_hotplug_exit(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_load_ptr, __atomic_load_ptr1, dev, hotplug_cb, msg, next_cb, next_dev uintptr
	var v1, v4 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+8 */ usbi_atomic_t
	_, _, _, _, _, _, _, _, _ = __atomic_load_ptr, __atomic_load_ptr1, dev, hotplug_cb, msg, next_cb, next_dev, v1, v4
	if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
		return
	}
	{
		__atomic_load_ptr = ctx + 200
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v1 != 0) {
		return
	}
	hotplug_cb = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+176)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	next_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(hotplug_cb+32 != ctx+176) {
			break
		}
		list_del(tls, hotplug_cb+32)
		libc.Xfree(tls, hotplug_cb)
		goto _2
	_2:
		;
		hotplug_cb = next_cb
		next_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(next_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	for !((*list_head)(unsafe.Pointer(ctx+520)).Fnext == ctx+520) {
		msg = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+520)).Fnext) - uint64(libc.UintptrFromInt32(0)+16))
		if (*usbi_hotplug_message)(unsafe.Pointer(msg)).Fevent == int32(LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT) {
			libusb_unref_device(tls, (*usbi_hotplug_message)(unsafe.Pointer(msg)).Fdevice)
		}
		list_del(tls, msg+16)
		libc.Xfree(tls, msg)
	}
	usbi_mutex_lock(tls, ctx+24)
	dev = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+64)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	next_dev = uintptr(uint64((*libusb_device)(unsafe.Pointer(dev)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(dev+32 != ctx+64) {
			break
		}
		{
			__atomic_load_ptr1 = dev
			libc.X__atomic_loadInt64(tls, __atomic_load_ptr1, bp+8, libc.Int32FromInt32(5))
			v4 = *(*usbi_atomic_t)(unsafe.Pointer(bp + 8))
		}
		if v4 == int64(1) {
			list_del(tls, dev+32)
		}
		usbi_recursively_remove_parents(tls, dev, next_dev)
		libusb_unref_device(tls, dev)
		goto _3
	_3:
		;
		dev = next_dev
		next_dev = uintptr(uint64((*libusb_device)(unsafe.Pointer(next_dev)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	usbi_mutex_unlock(tls, ctx+24)
	usbi_mutex_destroy(tls, ctx+136)
}

func usbi_hotplug_match_cb(tls *libc.TLS, dev uintptr, event libusb_hotplug_event, hotplug_cb uintptr) (r int32) {
	if !(libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags)&event != 0) {
		return 0
	}
	if libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags)&int32(USBI_HOTPLUG_VENDOR_ID_VALID) != 0 && libc.Int32FromUint16((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fvendor_id) != libc.Int32FromUint16((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FidVendor) {
		return 0
	}
	if libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags)&int32(USBI_HOTPLUG_PRODUCT_ID_VALID) != 0 && libc.Int32FromUint16((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fproduct_id) != libc.Int32FromUint16((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FidProduct) {
		return 0
	}
	if libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags)&int32(USBI_HOTPLUG_DEV_CLASS_VALID) != 0 && libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fdev_class) != libc.Int32FromUint8((*libusb_device)(unsafe.Pointer(dev)).Fdevice_descriptor.FbDeviceClass) {
		return 0
	}
	return (*(*func(*libc.TLS, uintptr, uintptr, libusb_hotplug_event, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fcb})))(tls, (*libusb_device)(unsafe.Pointer(dev)).Fctx, dev, event, (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fuser_data)
}

func usbi_hotplug_notification(tls *libc.TLS, ctx uintptr, dev uintptr, event libusb_hotplug_event) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var __atomic_load_ptr, msg uintptr
	var event_flags uint32
	var v1 usbi_atomic_t
	var _ /* __atomic_load_tmp at bp+0 */ usbi_atomic_t
	_, _, _, _ = __atomic_load_ptr, event_flags, msg, v1
	{
		__atomic_load_ptr = ctx + 200
		libc.X__atomic_loadInt64(tls, __atomic_load_ptr, bp, libc.Int32FromInt32(5))
		v1 = *(*usbi_atomic_t)(unsafe.Pointer(bp))
	}
	if !(v1 != 0) {
		return
	}
	msg = libc.Xcalloc(tls, uint64(1), uint64(32))
	if !(msg != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__128)), __ccgo_ts+10062, 0)
		return
	}
	(*usbi_hotplug_message)(unsafe.Pointer(msg)).Fevent = event
	(*usbi_hotplug_message)(unsafe.Pointer(msg)).Fdevice = dev
	usbi_mutex_lock(tls, ctx+424)
	event_flags = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags
	*(*uint32)(unsafe.Pointer(ctx + 464)) |= uint32(USBI_EVENT_HOTPLUG_MSG_PENDING)
	list_add_tail(tls, msg+16, ctx+520)
	if !(event_flags != 0) {
		usbi_signal_event(tls, ctx+16)
	}
	usbi_mutex_unlock(tls, ctx+424)
}

var __func__128 = [26]uint8{'u', 's', 'b', 'i', '_', 'h', 'o', 't', 'p', 'l', 'u', 'g', '_', 'n', 'o', 't', 'i', 'f', 'i', 'c', 'a', 't', 'i', 'o', 'n'}

func usbi_hotplug_process(tls *libc.TLS, ctx uintptr, hotplug_msgs uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var hotplug_cb, msg, next_cb uintptr
	var r int32
	_, _, _, _ = hotplug_cb, msg, next_cb, r
	usbi_mutex_lock(tls, ctx+136)
	for !((*list_head)(unsafe.Pointer(hotplug_msgs)).Fnext == hotplug_msgs) {
		msg = uintptr(uint64((*list_head)(unsafe.Pointer(hotplug_msgs)).Fnext) - uint64(libc.UintptrFromInt32(0)+16))
		hotplug_cb = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+176)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
		next_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
		for {
			if !(hotplug_cb+32 != ctx+176) {
				break
			}
			if libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags)&int32(USBI_HOTPLUG_NEEDS_FREE) != 0 {
				goto _1
			}
			usbi_mutex_unlock(tls, ctx+136)
			r = usbi_hotplug_match_cb(tls, (*usbi_hotplug_message)(unsafe.Pointer(msg)).Fdevice, (*usbi_hotplug_message)(unsafe.Pointer(msg)).Fevent, hotplug_cb)
			usbi_mutex_lock(tls, ctx+136)
			if r != 0 {
				list_del(tls, hotplug_cb+32)
				libc.Xfree(tls, hotplug_cb)
			}
			goto _1
		_1:
			;
			hotplug_cb = next_cb
			next_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(next_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
		}
		if (*usbi_hotplug_message)(unsafe.Pointer(msg)).Fevent == int32(LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT) {
			libusb_unref_device(tls, (*usbi_hotplug_message)(unsafe.Pointer(msg)).Fdevice)
		}
		list_del(tls, msg+16)
		libc.Xfree(tls, msg)
	}
	hotplug_cb = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+176)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	next_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(hotplug_cb+32 != ctx+176) {
			break
		}
		if libc.Int32FromUint8((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags)&int32(USBI_HOTPLUG_NEEDS_FREE) != 0 {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__129)), __ccgo_ts+10095, libc.VaList(bp+8, hotplug_cb, (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle))
			list_del(tls, hotplug_cb+32)
			libc.Xfree(tls, hotplug_cb)
		}
		goto _2
	_2:
		;
		hotplug_cb = next_cb
		next_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(next_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	usbi_mutex_unlock(tls, ctx+136)
}

var __func__129 = [21]uint8{'u', 's', 'b', 'i', '_', 'h', 'o', 't', 'p', 'l', 'u', 'g', '_', 'p', 'r', 'o', 'c', 'e', 's', 's'}

type __ccgo_fp__Xlibusb_hotplug_register_callback_6 = func(*libc.TLS, uintptr, uintptr, int32, uintptr) int32

func libusb_hotplug_register_callback(tls *libc.TLS, ctx uintptr, events int32, flags int32, vendor_id int32, product_id int32, dev_class int32, __ccgo_fp_cb_fn libusb_hotplug_callback_fn, user_data uintptr, callback_handle uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var hotplug_cb, v5, p1, p2, p3 uintptr
	var i, len1 ssize_t
	var v4 libusb_hotplug_callback_handle
	var _ /* devs at bp+0 */ uintptr
	_, _, _, _, _, _, _, _ = hotplug_cb, i, len1, v4, v5, p1, p2, p3
	if !(events != 0) || ^(int32(LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED)|int32(LIBUSB_HOTPLUG_EVENT_DEVICE_LEFT))&events != 0 || ^int32(LIBUSB_HOTPLUG_ENUMERATE)&flags != 0 || -int32(1) != vendor_id && ^libc.Int32FromInt32(0xffff)&vendor_id != 0 || -int32(1) != product_id && ^libc.Int32FromInt32(0xffff)&product_id != 0 || -int32(1) != dev_class && ^libc.Int32FromInt32(0xff)&dev_class != 0 || !(__ccgo_fp_cb_fn != 0) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
		return int32(LIBUSB_ERROR_NOT_SUPPORTED)
	}
	ctx = usbi_get_context(tls, ctx)
	hotplug_cb = libc.Xcalloc(tls, uint64(1), uint64(48))
	if !(hotplug_cb != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fflags = libc.Uint8FromInt32(events)
	if -int32(1) != vendor_id {
		p1 = hotplug_cb
		*(*uint8_t)(unsafe.Pointer(p1)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(p1))) | int32(USBI_HOTPLUG_VENDOR_ID_VALID))
		(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fvendor_id = libc.Uint16FromInt32(vendor_id)
	}
	if -int32(1) != product_id {
		p2 = hotplug_cb
		*(*uint8_t)(unsafe.Pointer(p2)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(p2))) | int32(USBI_HOTPLUG_PRODUCT_ID_VALID))
		(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fproduct_id = libc.Uint16FromInt32(product_id)
	}
	if -int32(1) != dev_class {
		p3 = hotplug_cb
		*(*uint8_t)(unsafe.Pointer(p3)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(p3))) | int32(USBI_HOTPLUG_DEV_CLASS_VALID))
		(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fdev_class = libc.Uint8FromInt32(dev_class)
	}
	(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fcb = __ccgo_fp_cb_fn
	(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fuser_data = user_data
	usbi_mutex_lock(tls, ctx+136)
	v5 = ctx + 192
	v4 = *(*libusb_hotplug_callback_handle)(unsafe.Pointer(v5))
	*(*libusb_hotplug_callback_handle)(unsafe.Pointer(v5))++
	(*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle = v4
	if (*libusb_context)(unsafe.Pointer(ctx)).Fnext_hotplug_cb_handle < 0 {
		(*libusb_context)(unsafe.Pointer(ctx)).Fnext_hotplug_cb_handle = int32(1)
	}
	list_add(tls, hotplug_cb+32, ctx+176)
	usbi_mutex_unlock(tls, ctx+136)
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__130)), __ccgo_ts+10132, libc.VaList(bp+16, hotplug_cb, (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle))
	if flags&int32(LIBUSB_HOTPLUG_ENUMERATE) != 0 && events&int32(LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED) != 0 {
		len1 = libusb_get_device_list(tls, ctx, bp)
		if len1 < 0 {
			libusb_hotplug_deregister_callback(tls, ctx, (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle)
			return int32(len1)
		}
		i = 0
		for {
			if !(i < len1) {
				break
			}
			usbi_hotplug_match_cb(tls, *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + uintptr(i)*8)), int32(LIBUSB_HOTPLUG_EVENT_DEVICE_ARRIVED), hotplug_cb)
			goto _6
		_6:
			;
			i++
		}
		libusb_free_device_list(tls, *(*uintptr)(unsafe.Pointer(bp)), int32(1))
	}
	if callback_handle != 0 {
		*(*libusb_hotplug_callback_handle)(unsafe.Pointer(callback_handle)) = (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle
	}
	return int32(LIBUSB_SUCCESS)
}

var __func__130 = [33]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'h', 'o', 't', 'p', 'l', 'u', 'g', '_', 'r', 'e', 'g', 'i', 's', 't', 'e', 'r', '_', 'c', 'a', 'l', 'l', 'b', 'a', 'c', 'k'}

func libusb_hotplug_deregister_callback(tls *libc.TLS, ctx uintptr, callback_handle libusb_hotplug_callback_handle) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var deregistered int32
	var event_flags uint32
	var hotplug_cb, p2 uintptr
	_, _, _, _ = deregistered, event_flags, hotplug_cb, p2
	deregistered = 0
	if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
		return
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__131)), __ccgo_ts+10165, libc.VaList(bp+8, callback_handle))
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+136)
	hotplug_cb = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+176)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(hotplug_cb+32 != ctx+176) {
			break
		}
		if callback_handle == (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle {
			p2 = hotplug_cb
			*(*uint8_t)(unsafe.Pointer(p2)) = uint8_t(int32(*(*uint8_t)(unsafe.Pointer(p2))) | int32(USBI_HOTPLUG_NEEDS_FREE))
			deregistered = int32(1)
			break
		}
		goto _1
	_1:
		;
		hotplug_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	usbi_mutex_unlock(tls, ctx+136)
	if deregistered != 0 {
		usbi_mutex_lock(tls, ctx+424)
		event_flags = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags
		*(*uint32)(unsafe.Pointer(ctx + 464)) |= uint32(USBI_EVENT_HOTPLUG_CB_DEREGISTERED)
		if !(event_flags != 0) {
			usbi_signal_event(tls, ctx+16)
		}
		usbi_mutex_unlock(tls, ctx+424)
	}
}

var __func__131 = [35]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'h', 'o', 't', 'p', 'l', 'u', 'g', '_', 'd', 'e', 'r', 'e', 'g', 'i', 's', 't', 'e', 'r', '_', 'c', 'a', 'l', 'l', 'b', 'a', 'c', 'k'}

func libusb_hotplug_get_user_data(tls *libc.TLS, ctx uintptr, callback_handle libusb_hotplug_callback_handle) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var hotplug_cb, user_data uintptr
	_, _ = hotplug_cb, user_data
	user_data = libc.UintptrFromInt32(0)
	if !(libusb_has_capability(tls, uint32(LIBUSB_CAP_HAS_HOTPLUG)) != 0) {
		return libc.UintptrFromInt32(0)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__132)), __ccgo_ts+10190, libc.VaList(bp+8, callback_handle))
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+136)
	hotplug_cb = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+176)).Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	for {
		if !(hotplug_cb+32 != ctx+176) {
			break
		}
		if callback_handle == (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fhandle {
			user_data = (*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Fuser_data
			break
		}
		goto _1
	_1:
		;
		hotplug_cb = uintptr(uint64((*usbi_hotplug_callback)(unsafe.Pointer(hotplug_cb)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+32))
	}
	usbi_mutex_unlock(tls, ctx+136)
	return user_data
}

var __func__132 = [29]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'h', 'o', 't', 'p', 'l', 'u', 'g', '_', 'g', 'e', 't', '_', 'u', 's', 'e', 'r', '_', 'd', 'a', 't', 'a'}

func usbi_io_init(tls *libc.TLS, ctx uintptr) (r1 int32) {
	var r int32
	_ = r
	usbi_mutex_init(tls, ctx+208)
	usbi_mutex_init(tls, ctx+288)
	usbi_mutex_init(tls, ctx+336)
	usbi_cond_init(tls, ctx+376)
	usbi_mutex_init(tls, ctx+424)
	usbi_tls_key_create(tls, ctx+332)
	list_init(tls, ctx+248)
	list_init(tls, ctx+472)
	list_init(tls, ctx+488)
	list_init(tls, ctx+520)
	list_init(tls, ctx+536)
	r = usbi_create_event(tls, ctx+16)
	if r < 0 {
		goto err
	}
	r = usbi_add_event_source(tls, ctx, *(*int32)(unsafe.Pointer(ctx + 16)), int16(0x001))
	if r < 0 {
		goto err_destroy_event
	}
	return 0
	goto err_destroy_event
err_destroy_event:
	;
	usbi_destroy_event(tls, ctx+16)
	goto err
err:
	;
	usbi_mutex_destroy(tls, ctx+208)
	usbi_mutex_destroy(tls, ctx+288)
	usbi_mutex_destroy(tls, ctx+336)
	usbi_cond_destroy(tls, ctx+376)
	usbi_mutex_destroy(tls, ctx+424)
	usbi_tls_key_delete(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_handling_key)
	return r
}

func cleanup_removed_event_sources(tls *libc.TLS, ctx uintptr) {
	var ievent_source, tmp uintptr
	_, _ = ievent_source, tmp
	ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+488)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	tmp = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	for {
		if !(ievent_source+8 != ctx+488) {
			break
		}
		list_del(tls, ievent_source+8)
		libc.Xfree(tls, ievent_source)
		goto _1
	_1:
		;
		ievent_source = tmp
		tmp = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(tmp)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	}
}

func usbi_io_exit(tls *libc.TLS, ctx uintptr) {
	usbi_remove_event_source(tls, ctx, *(*int32)(unsafe.Pointer(ctx + 16)))
	usbi_destroy_event(tls, ctx+16)
	usbi_mutex_destroy(tls, ctx+208)
	usbi_mutex_destroy(tls, ctx+288)
	usbi_mutex_destroy(tls, ctx+336)
	usbi_cond_destroy(tls, ctx+376)
	usbi_mutex_destroy(tls, ctx+424)
	usbi_tls_key_delete(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_handling_key)
	cleanup_removed_event_sources(tls, ctx)
	libc.Xfree(tls, (*libusb_context)(unsafe.Pointer(ctx)).Fevent_data)
}

func calculate_timeout(tls *libc.TLS, itransfer uintptr) {
	var timeout uint32
	var transfer uintptr
	var v1 int64
	_, _, _ = timeout, transfer, v1
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	timeout = (*libusb_transfer)(unsafe.Pointer(transfer)).Ftimeout
	if !(timeout != 0) {
		v1 = libc.Int64FromInt32(0)
		(*timespec)(unsafe.Pointer(itransfer + 80)).Ftv_nsec = v1
		(*timespec)(unsafe.Pointer(itransfer + 80)).Ftv_sec = v1
		return
	}
	usbi_get_monotonic_time(tls, itransfer+80)
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout.Ftv_sec += libc.Int64FromUint32(timeout / uint32(1000))
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout.Ftv_nsec += libc.Int64FromUint32(timeout%libc.Uint32FromUint32(1000)) * int64(1000000)
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout.Ftv_nsec >= int64(1000000000) {
		(*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout.Ftv_sec++
		(*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout.Ftv_nsec -= int64(1000000000)
	}
}

func libusb_alloc_transfer(tls *libc.TLS, iso_packets int32) (r uintptr) {
	var alloc_size, iso_packets_size, libusb_transfer_size, priv_size, usbi_transfer_size size_t
	var itransfer, ptr, transfer uintptr
	var v1 bool
	_, _, _, _, _, _, _, _, _ = alloc_size, iso_packets_size, itransfer, libusb_transfer_size, priv_size, ptr, transfer, usbi_transfer_size, v1
	if v1 = iso_packets >= 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+10218, __ccgo_ts+10235, int32(1292), uintptr(unsafe.Pointer(&__func__133)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if iso_packets < 0 {
		return libc.UintptrFromInt32(0)
	}
	priv_size = (usbi_backend.Ftransfer_priv_size + (libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))
	usbi_transfer_size = (libc.Uint64FromInt64(128) + (libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))
	libusb_transfer_size = (libc.Uint64FromInt64(64) + (libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8) - libc.Uint64FromInt32(1))
	iso_packets_size = uint64(12) * libc.Uint64FromInt32(iso_packets)
	alloc_size = priv_size + usbi_transfer_size + libusb_transfer_size + iso_packets_size
	ptr = libc.Xcalloc(tls, uint64(1), alloc_size)
	if !(ptr != 0) {
		return libc.UintptrFromInt32(0)
	}
	itransfer = ptr + uintptr(priv_size)
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Fnum_iso_packets = iso_packets
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Fpriv = ptr
	usbi_mutex_init(tls, itransfer)
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	return transfer
}

var __func__133 = [22]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'a', 'l', 'l', 'o', 'c', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func libusb_free_transfer(tls *libc.TLS, transfer uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var itransfer, ptr, v1 uintptr
	var v2 bool
	_, _, _, _ = itransfer, ptr, v1, v2
	if !(transfer != 0) {
		return
	}
	if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__134)), __ccgo_ts+10274, libc.VaList(bp+8, transfer))
	if libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Fflags)&int32(LIBUSB_TRANSFER_FREE_BUFFER) != 0 {
		libc.Xfree(tls, (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer)
	}
	itransfer = transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	usbi_mutex_destroy(tls, itransfer)
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		libusb_unref_device(tls, (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)
	}
	ptr = itransfer - uintptr((usbi_backend.Ftransfer_priv_size+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	if v2 = ptr == (*usbi_transfer)(unsafe.Pointer(itransfer)).Fpriv; !v2 {
		libc.X__assert_fail(tls, __ccgo_ts+10286, __ccgo_ts+10235, int32(1346), uintptr(unsafe.Pointer(&__func__134)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	libc.Xfree(tls, ptr)
}

var __func__134 = [21]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'f', 'r', 'e', 'e', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func arm_timer_for_next_timeout(tls *libc.TLS, ctx uintptr) (r int32) {
	_ = ctx
	return 0
}

func add_to_flying_list(tls *libc.TLS, itransfer uintptr) (r1 int32) {
	var ctx, cur, cur_ts, timeout, v1 uintptr
	var first, r, v3 int32
	var v4 bool
	_, _, _, _, _, _, _, _, _ = ctx, cur, cur_ts, first, r, timeout, v1, v3, v4
	timeout = itransfer + 80
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	r = 0
	first = int32(1)
	calculate_timeout(tls, itransfer)
	if (*list_head)(unsafe.Pointer(ctx+248)).Fnext == ctx+248 {
		list_add(tls, itransfer+48, ctx+248)
		goto out
	}
	if !((*timespec)(unsafe.Pointer(timeout)).Ftv_sec != 0 || (*timespec)(unsafe.Pointer(timeout)).Ftv_nsec != 0) {
		list_add_tail(tls, itransfer+48, ctx+248)
		goto out
	}
	cur = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+248)).Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	for {
		if !(cur+48 != ctx+248) {
			break
		}
		cur_ts = cur + 80
		if v4 = !((*timespec)(unsafe.Pointer(cur_ts)).Ftv_sec != 0 || (*timespec)(unsafe.Pointer(cur_ts)).Ftv_nsec != 0); !v4 {
			if (*timespec)(unsafe.Pointer(cur_ts)).Ftv_sec == (*timespec)(unsafe.Pointer(timeout)).Ftv_sec {
				v3 = libc.BoolInt32((*timespec)(unsafe.Pointer(cur_ts)).Ftv_nsec > (*timespec)(unsafe.Pointer(timeout)).Ftv_nsec)
			} else {
				v3 = libc.BoolInt32((*timespec)(unsafe.Pointer(cur_ts)).Ftv_sec > (*timespec)(unsafe.Pointer(timeout)).Ftv_sec)
			}
		}
		if v4 || v3 != 0 {
			list_add_tail(tls, itransfer+48, cur+48)
			goto out
		}
		first = 0
		goto _2
	_2:
		;
		cur = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(cur)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	}
	list_add_tail(tls, itransfer+48, ctx+248)
	goto out
out:
	;
	_ = first
	if r != 0 {
		list_del(tls, itransfer+48)
	}
	return r
}

func remove_from_flying_list(tls *libc.TLS, itransfer uintptr) (r1 int32) {
	var ctx, v1 uintptr
	var r, rearm_timer int32
	_, _, _, _ = ctx, r, rearm_timer, v1
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	r = 0
	rearm_timer = libc.BoolInt32(((*timespec)(unsafe.Pointer(itransfer+80)).Ftv_sec != 0 || (*timespec)(unsafe.Pointer(itransfer+80)).Ftv_nsec != 0) && uintptr(uint64((*list_head)(unsafe.Pointer(ctx+248)).Fnext)-uint64(libc.UintptrFromInt32(0)+48)) == itransfer)
	list_del(tls, itransfer+48)
	if rearm_timer != 0 {
		r = arm_timer_for_next_timeout(tls, ctx)
	}
	return r
}

func libusb_submit_transfer(tls *libc.TLS, transfer uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ctx, itransfer, v2 uintptr
	var r int32
	var v1 bool
	_, _, _, _, _ = ctx, itransfer, r, v1, v2
	itransfer = transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	if v1 = (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle != 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+10309, __ccgo_ts+10235, int32(1497), uintptr(unsafe.Pointer(&__func__135)))
	}
	_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		libusb_unref_device(tls, (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)
	}
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev = libusb_ref_device(tls, (*libusb_device_handle)(unsafe.Pointer((*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)).Fdev)
	if (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle != 0 {
		v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer((*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)).Fdev)).Fctx
	} else {
		v2 = libc.UintptrFromInt32(0)
	}
	ctx = v2
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__135)), __ccgo_ts+10274, libc.VaList(bp+8, transfer))
	usbi_mutex_lock(tls, ctx+208)
	usbi_mutex_lock(tls, itransfer)
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fstate_flags&uint32(USBI_TRANSFER_IN_FLIGHT) != 0 {
		usbi_mutex_unlock(tls, ctx+208)
		usbi_mutex_unlock(tls, itransfer)
		return int32(LIBUSB_ERROR_BUSY)
	}
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Ftransferred = 0
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Fstate_flags = uint32(0)
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout_flags = uint32(0)
	r = add_to_flying_list(tls, itransfer)
	if r != 0 {
		usbi_mutex_unlock(tls, ctx+208)
		usbi_mutex_unlock(tls, itransfer)
		return r
	}
	usbi_mutex_unlock(tls, ctx+208)
	r = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fsubmit_transfer})))(tls, itransfer)
	if r == int32(LIBUSB_SUCCESS) {
		*(*uint32_t)(unsafe.Pointer(itransfer + 104)) |= uint32(USBI_TRANSFER_IN_FLIGHT)
	}
	usbi_mutex_unlock(tls, itransfer)
	if r != int32(LIBUSB_SUCCESS) {
		usbi_mutex_lock(tls, ctx+208)
		remove_from_flying_list(tls, itransfer)
		usbi_mutex_unlock(tls, ctx+208)
	}
	return r
}

var __func__135 = [23]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 's', 'u', 'b', 'm', 'i', 't', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func libusb_cancel_transfer(tls *libc.TLS, transfer uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ctx, itransfer, v1 uintptr
	var r int32
	_, _, _, _ = ctx, itransfer, r, v1
	itransfer = transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__136)), __ccgo_ts+10274, libc.VaList(bp+8, transfer))
	usbi_mutex_lock(tls, itransfer)
	if !((*usbi_transfer)(unsafe.Pointer(itransfer)).Fstate_flags&uint32(USBI_TRANSFER_IN_FLIGHT) != 0) || (*usbi_transfer)(unsafe.Pointer(itransfer)).Fstate_flags&uint32(USBI_TRANSFER_CANCELLING) != 0 {
		r = int32(LIBUSB_ERROR_NOT_FOUND)
		goto out
	}
	r = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fcancel_transfer})))(tls, itransfer)
	if r < 0 {
		if r != int32(LIBUSB_ERROR_NOT_FOUND) && r != int32(LIBUSB_ERROR_NO_DEVICE) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__136)), __ccgo_ts+10330, libc.VaList(bp+8, r))
		} else {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__136)), __ccgo_ts+10330, libc.VaList(bp+8, r))
		}
		if r == int32(LIBUSB_ERROR_NO_DEVICE) {
			*(*uint32_t)(unsafe.Pointer(itransfer + 104)) |= uint32(USBI_TRANSFER_DEVICE_DISAPPEARED)
		}
	}
	*(*uint32_t)(unsafe.Pointer(itransfer + 104)) |= uint32(USBI_TRANSFER_CANCELLING)
	goto out
out:
	;
	usbi_mutex_unlock(tls, itransfer)
	return r
}

var __func__136 = [23]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'c', 'a', 'n', 'c', 'e', 'l', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func libusb_transfer_set_stream_id(tls *libc.TLS, transfer uintptr, stream_id uint32_t) {
	var itransfer uintptr
	_ = itransfer
	itransfer = transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	(*usbi_transfer)(unsafe.Pointer(itransfer)).Fstream_id = stream_id
}

func libusb_transfer_get_stream_id(tls *libc.TLS, transfer uintptr) (r uint32_t) {
	var itransfer uintptr
	_ = itransfer
	itransfer = transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	return (*usbi_transfer)(unsafe.Pointer(itransfer)).Fstream_id
}

func usbi_handle_transfer_completion(tls *libc.TLS, itransfer uintptr, status libusb_transfer_status) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ctx, transfer, v1 uintptr
	var flags uint8_t
	var r, rqlen int32
	var v2 bool
	_, _, _, _, _, _, _ = ctx, flags, r, rqlen, transfer, v1, v2
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	usbi_mutex_lock(tls, ctx+208)
	r = remove_from_flying_list(tls, itransfer)
	usbi_mutex_unlock(tls, ctx+208)
	if r < 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__137)), __ccgo_ts+10362, 0)
	}
	usbi_mutex_lock(tls, itransfer)
	*(*uint32_t)(unsafe.Pointer(itransfer + 104)) &= libc.Uint32FromInt32(^int32(USBI_TRANSFER_IN_FLIGHT))
	usbi_mutex_unlock(tls, itransfer)
	if status == int32(LIBUSB_TRANSFER_COMPLETED) && libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Fflags)&int32(LIBUSB_TRANSFER_SHORT_NOT_OK) != 0 {
		rqlen = (*libusb_transfer)(unsafe.Pointer(transfer)).Flength
		if libc.Int32FromUint8((*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1) == int32(LIBUSB_TRANSFER_TYPE_CONTROL) {
			rqlen = int32(uint64(rqlen) - libc.Uint64FromInt64(8))
		}
		if rqlen != (*usbi_transfer)(unsafe.Pointer(itransfer)).Ftransferred {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__137)), __ccgo_ts+10399, 0)
			status = int32(LIBUSB_TRANSFER_ERROR)
		}
	}
	flags = (*libusb_transfer)(unsafe.Pointer(transfer)).Fflags
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus = status
	(*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length = (*usbi_transfer)(unsafe.Pointer(itransfer)).Ftransferred
	if v2 = (*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length >= 0; !v2 {
		libc.X__assert_fail(tls, __ccgo_ts+10436, __ccgo_ts+10235, int32(1721), uintptr(unsafe.Pointer(&__func__137)))
	}
	_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__137)), __ccgo_ts+10465, libc.VaList(bp+8, transfer, (*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback))
	if (*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback != 0 {
		libusb_lock_event_waiters(tls, ctx)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*libusb_transfer)(unsafe.Pointer(transfer)).Fcallback})))(tls, transfer)
		libusb_unlock_event_waiters(tls, ctx)
	}
	if libc.Int32FromUint8(flags)&int32(LIBUSB_TRANSFER_FREE_TRANSFER) != 0 {
		libusb_free_transfer(tls, transfer)
	}
	return r
}

var __func__137 = [32]uint8{'u', 's', 'b', 'i', '_', 'h', 'a', 'n', 'd', 'l', 'e', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r', '_', 'c', 'o', 'm', 'p', 'l', 'e', 't', 'i', 'o', 'n'}

func usbi_handle_transfer_cancellation(tls *libc.TLS, itransfer uintptr) (r int32) {
	var ctx, v1 uintptr
	var timed_out uint8_t
	_, _, _ = ctx, timed_out, v1
	if (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	usbi_mutex_lock(tls, ctx+208)
	timed_out = uint8((*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout_flags & uint32(USBI_TRANSFER_TIMED_OUT))
	usbi_mutex_unlock(tls, ctx+208)
	if timed_out != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__138)), __ccgo_ts+10493, 0)
		return usbi_handle_transfer_completion(tls, itransfer, int32(LIBUSB_TRANSFER_TIMED_OUT))
	}
	return usbi_handle_transfer_completion(tls, itransfer, int32(LIBUSB_TRANSFER_CANCELLED))
}

var __func__138 = [34]uint8{'u', 's', 'b', 'i', '_', 'h', 'a', 'n', 'd', 'l', 'e', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r', '_', 'c', 'a', 'n', 'c', 'e', 'l', 'l', 'a', 't', 'i', 'o', 'n'}

func usbi_signal_transfer_completion(tls *libc.TLS, itransfer uintptr) {
	var ctx, dev uintptr
	var event_flags uint32
	_, _, _ = ctx, dev, event_flags
	dev = (*usbi_transfer)(unsafe.Pointer(itransfer)).Fdev
	if dev != 0 {
		ctx = (*libusb_device)(unsafe.Pointer(dev)).Fctx
		usbi_mutex_lock(tls, ctx+424)
		event_flags = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags
		*(*uint32)(unsafe.Pointer(ctx + 464)) |= uint32(USBI_EVENT_TRANSFER_COMPLETED)
		list_add_tail(tls, itransfer+64, ctx+536)
		if !(event_flags != 0) {
			usbi_signal_event(tls, ctx+16)
		}
		usbi_mutex_unlock(tls, ctx+424)
	}
}

func libusb_try_lock_events(tls *libc.TLS, ctx uintptr) (r1 int32) {
	var r int32
	var ru uint32
	_, _ = r, ru
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+424)
	ru = (*libusb_context)(unsafe.Pointer(ctx)).Fdevice_close
	usbi_mutex_unlock(tls, ctx+424)
	if ru != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__139)), __ccgo_ts+10523, 0)
		return int32(1)
	}
	r = usbi_mutex_trylock(tls, ctx+288)
	if !(r != 0) {
		return int32(1)
	}
	(*libusb_context)(unsafe.Pointer(ctx)).Fevent_handler_active = int32(1)
	return 0
}

var __func__139 = [23]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 't', 'r', 'y', '_', 'l', 'o', 'c', 'k', '_', 'e', 'v', 'e', 'n', 't', 's'}

func libusb_lock_events(tls *libc.TLS, ctx uintptr) {
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+288)
	(*libusb_context)(unsafe.Pointer(ctx)).Fevent_handler_active = int32(1)
}

func libusb_unlock_events(tls *libc.TLS, ctx uintptr) {
	ctx = usbi_get_context(tls, ctx)
	(*libusb_context)(unsafe.Pointer(ctx)).Fevent_handler_active = 0
	usbi_mutex_unlock(tls, ctx+288)
	usbi_mutex_lock(tls, ctx+336)
	usbi_cond_broadcast(tls, ctx+376)
	usbi_mutex_unlock(tls, ctx+336)
}

func libusb_event_handling_ok(tls *libc.TLS, ctx uintptr) (r1 int32) {
	var r uint32
	_ = r
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+424)
	r = (*libusb_context)(unsafe.Pointer(ctx)).Fdevice_close
	usbi_mutex_unlock(tls, ctx+424)
	if r != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__140)), __ccgo_ts+10523, 0)
		return 0
	}
	return int32(1)
}

var __func__140 = [25]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'e', 'v', 'e', 'n', 't', '_', 'h', 'a', 'n', 'd', 'l', 'i', 'n', 'g', '_', 'o', 'k'}

func libusb_event_handler_active(tls *libc.TLS, ctx uintptr) (r1 int32) {
	var r uint32
	_ = r
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+424)
	r = (*libusb_context)(unsafe.Pointer(ctx)).Fdevice_close
	usbi_mutex_unlock(tls, ctx+424)
	if r != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__141)), __ccgo_ts+10523, 0)
		return int32(1)
	}
	return (*libusb_context)(unsafe.Pointer(ctx)).Fevent_handler_active
}

var __func__141 = [28]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'e', 'v', 'e', 'n', 't', '_', 'h', 'a', 'n', 'd', 'l', 'e', 'r', '_', 'a', 'c', 't', 'i', 'v', 'e'}

func libusb_interrupt_event_handler(tls *libc.TLS, ctx uintptr) {
	var event_flags uint32
	_ = event_flags
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__142)), __ccgo_ts+6578, 0)
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+424)
	event_flags = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags
	*(*uint32)(unsafe.Pointer(ctx + 464)) |= uint32(USBI_EVENT_USER_INTERRUPT)
	if !(event_flags != 0) {
		usbi_signal_event(tls, ctx+16)
	}
	usbi_mutex_unlock(tls, ctx+424)
}

var __func__142 = [31]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'i', 'n', 't', 'e', 'r', 'r', 'u', 'p', 't', '_', 'e', 'v', 'e', 'n', 't', '_', 'h', 'a', 'n', 'd', 'l', 'e', 'r'}

func libusb_lock_event_waiters(tls *libc.TLS, ctx uintptr) {
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+336)
}

func libusb_unlock_event_waiters(tls *libc.TLS, ctx uintptr) {
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_unlock(tls, ctx+336)
}

func libusb_wait_for_event(tls *libc.TLS, ctx uintptr, tv uintptr) (r1 int32) {
	var r int32
	_ = r
	ctx = usbi_get_context(tls, ctx)
	if !(tv != 0) {
		usbi_cond_wait(tls, ctx+376, ctx+336)
		return 0
	}
	if !((*timeval)(unsafe.Pointer(tv)).Ftv_sec >= 0 && (*timeval)(unsafe.Pointer(tv)).Ftv_usec >= 0 && (*timeval)(unsafe.Pointer(tv)).Ftv_usec < int64(1000000)) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	r = usbi_cond_timedwait(tls, ctx+376, ctx+336, tv)
	if r < 0 {
		return libc.BoolInt32(r == int32(LIBUSB_ERROR_TIMEOUT))
	}
	return 0
}

func handle_timeout(tls *libc.TLS, itransfer uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r int32
	var transfer, v1 uintptr
	_, _, _ = r, transfer, v1
	transfer = itransfer + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
	*(*uint32_t)(unsafe.Pointer(itransfer + 108)) |= uint32(USBI_TRANSFER_TIMEOUT_HANDLED)
	r = libusb_cancel_transfer(tls, transfer)
	if r == int32(LIBUSB_SUCCESS) {
		*(*uint32_t)(unsafe.Pointer(itransfer + 108)) |= uint32(USBI_TRANSFER_TIMED_OUT)
	} else {
		if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
			v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
		} else {
			v1 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__143)), __ccgo_ts+10556, libc.VaList(bp+8, r))
	}
}

var __func__143 = [15]uint8{'h', 'a', 'n', 'd', 'l', 'e', '_', 't', 'i', 'm', 'e', 'o', 'u', 't'}

func handle_timeouts_locked(tls *libc.TLS, ctx uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var cur_ts, itransfer uintptr
	var v2 int32
	var _ /* systime at bp+0 */ timespec
	_, _, _ = cur_ts, itransfer, v2
	if (*list_head)(unsafe.Pointer(ctx+248)).Fnext == ctx+248 {
		return
	}
	usbi_get_monotonic_time(tls, bp)
	itransfer = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+248)).Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	for {
		if !(itransfer+48 != ctx+248) {
			break
		}
		cur_ts = itransfer + 80
		if !((*timespec)(unsafe.Pointer(cur_ts)).Ftv_sec != 0 || (*timespec)(unsafe.Pointer(cur_ts)).Ftv_nsec != 0) {
			return
		}
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout_flags&libc.Uint32FromInt32(int32(USBI_TRANSFER_TIMEOUT_HANDLED)|int32(USBI_TRANSFER_OS_HANDLES_TIMEOUT)) != 0 {
			goto _1
		}
		if (*timespec)(unsafe.Pointer(cur_ts)).Ftv_sec == (*timespec)(unsafe.Pointer(bp)).Ftv_sec {
			v2 = libc.BoolInt32((*timespec)(unsafe.Pointer(cur_ts)).Ftv_nsec > (*timespec)(unsafe.Pointer(bp)).Ftv_nsec)
		} else {
			v2 = libc.BoolInt32((*timespec)(unsafe.Pointer(cur_ts)).Ftv_sec > (*timespec)(unsafe.Pointer(bp)).Ftv_sec)
		}
		if v2 != 0 {
			return
		}
		handle_timeout(tls, itransfer)
		goto _1
	_1:
		;
		itransfer = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(itransfer)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	}
}

func handle_timeouts(tls *libc.TLS, ctx uintptr) {
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+208)
	handle_timeouts_locked(tls, ctx)
	usbi_mutex_unlock(tls, ctx+208)
}

func handle_event_trigger(tls *libc.TLS, ctx uintptr) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var hotplug_event, r int32
	var itransfer, tmp uintptr
	var v1, v2 bool
	var _ /* completed_transfers at bp+16 */ list_head
	var _ /* hotplug_msgs at bp+0 */ list_head
	_, _, _, _, _, _ = hotplug_event, itransfer, r, tmp, v1, v2
	hotplug_event = 0
	r = 0
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10579, 0)
	list_init(tls, bp)
	usbi_mutex_lock(tls, ctx+424)
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_EVENT_SOURCES_MODIFIED) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10595, 0)
	}
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_USER_INTERRUPT) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10629, 0)
		*(*uint32)(unsafe.Pointer(ctx + 464)) &= libc.Uint32FromInt32(^int32(USBI_EVENT_USER_INTERRUPT))
	}
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_HOTPLUG_CB_DEREGISTERED) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10662, 0)
		*(*uint32)(unsafe.Pointer(ctx + 464)) &= libc.Uint32FromInt32(^int32(USBI_EVENT_HOTPLUG_CB_DEREGISTERED))
		hotplug_event = int32(1)
	}
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_DEVICE_CLOSE) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10696, 0)
	}
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_HOTPLUG_MSG_PENDING) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10724, 0)
		*(*uint32)(unsafe.Pointer(ctx + 464)) &= libc.Uint32FromInt32(^int32(USBI_EVENT_HOTPLUG_MSG_PENDING))
		hotplug_event = int32(1)
		if v1 = !((*list_head)(unsafe.Pointer(ctx+520)).Fnext == ctx+520); !v1 {
			libc.X__assert_fail(tls, __ccgo_ts+10749, __ccgo_ts+10235, int32(2147), uintptr(unsafe.Pointer(&__func__144)))
		}
		_ = v1 || libc.Bool(libc.Int32FromInt32(0) != 0)
		list_cut(tls, bp, ctx+520)
	}
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_TRANSFER_COMPLETED) != 0 {
		if v2 = !((*list_head)(unsafe.Pointer(ctx+536)).Fnext == ctx+536); !v2 {
			libc.X__assert_fail(tls, __ccgo_ts+10781, __ccgo_ts+10235, int32(2156), uintptr(unsafe.Pointer(&__func__144)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		list_cut(tls, bp+16, ctx+536)
		usbi_mutex_unlock(tls, ctx+424)
		itransfer = uintptr(uint64((*list_head)(unsafe.Pointer(bp+16)).Fnext) - uint64(libc.UintptrFromInt32(0)+64))
		tmp = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(itransfer)).Fcompleted_list.Fnext) - uint64(libc.UintptrFromInt32(0)+64))
		for {
			if !(itransfer+64 != bp+16) {
				break
			}
			list_del(tls, itransfer+64)
			r = (*(*func(*libc.TLS, uintptr) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fhandle_transfer_completion})))(tls, itransfer)
			if r != 0 {
				usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__144)), __ccgo_ts+10820, libc.VaList(bp+40, r))
				break
			}
			goto _3
		_3:
			;
			itransfer = tmp
			tmp = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(tmp)).Fcompleted_list.Fnext) - uint64(libc.UintptrFromInt32(0)+64))
		}
		usbi_mutex_lock(tls, ctx+424)
		if !((*list_head)(unsafe.Pointer(bp+16)).Fnext == bp+16) {
			list_splice_front(tls, bp+16, ctx+536)
		} else {
			if (*list_head)(unsafe.Pointer(ctx+536)).Fnext == ctx+536 {
				*(*uint32)(unsafe.Pointer(ctx + 464)) &= libc.Uint32FromInt32(^int32(USBI_EVENT_TRANSFER_COMPLETED))
			}
		}
	}
	if !((*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags != 0) {
		usbi_clear_event(tls, ctx+16)
	}
	usbi_mutex_unlock(tls, ctx+424)
	if hotplug_event != 0 {
		usbi_hotplug_process(tls, ctx, bp)
	}
	return r
}

var __func__144 = [21]uint8{'h', 'a', 'n', 'd', 'l', 'e', '_', 'e', 'v', 'e', 'n', 't', '_', 't', 'r', 'i', 'g', 'g', 'e', 'r'}

func handle_events(tls *libc.TLS, ctx uintptr, tv uintptr) (r1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var r, timeout_ms int32
	var _ /* reported_events at bp+0 */ usbi_reported_events
	_, _ = r, timeout_ms
	if usbi_handling_events(tls, ctx) != 0 {
		return int32(LIBUSB_ERROR_BUSY)
	}
	usbi_mutex_lock(tls, ctx+424)
	if (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags&uint32(USBI_EVENT_EVENT_SOURCES_MODIFIED) != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__145)), __ccgo_ts+10876, 0)
		cleanup_removed_event_sources(tls, ctx)
		r = usbi_alloc_event_data(tls, ctx)
		if r != 0 {
			usbi_mutex_unlock(tls, ctx+424)
			return r
		}
		*(*uint32)(unsafe.Pointer(ctx + 464)) &= libc.Uint32FromInt32(^int32(USBI_EVENT_EVENT_SOURCES_MODIFIED))
		if !((*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags != 0) {
			usbi_clear_event(tls, ctx+16)
		}
	}
	usbi_mutex_unlock(tls, ctx+424)
	timeout_ms = int32(int64(int32((*timeval)(unsafe.Pointer(tv)).Ftv_sec*libc.Int64FromInt32(1000))) + (*timeval)(unsafe.Pointer(tv)).Ftv_usec/int64(1000))
	if (*timeval)(unsafe.Pointer(tv)).Ftv_usec%int64(1000) != 0 {
		timeout_ms++
	}
	*(*uint32)(unsafe.Pointer(&*(*usbi_reported_events)(unsafe.Pointer(bp)))) = uint32(0)
	usbi_start_event_handling(tls, ctx)
	r = usbi_wait_for_events(tls, ctx, bp, timeout_ms)
	if r != int32(LIBUSB_SUCCESS) {
		if r == int32(LIBUSB_ERROR_TIMEOUT) {
			handle_timeouts(tls, ctx)
			r = int32(LIBUSB_SUCCESS)
		}
		goto done
	}
	if int32(uint32(*(*uint8)(unsafe.Pointer(bp + 0))&0x1>>0)) != 0 {
		r = handle_event_trigger(tls, ctx)
		if r != 0 {
			goto done
		}
	}
	if !((*(*usbi_reported_events)(unsafe.Pointer(bp))).Fnum_ready != 0) {
		goto done
	}
	r = (*(*func(*libc.TLS, uintptr, uintptr, uint32, uint32) int32)(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fhandle_events})))(tls, ctx, (*(*usbi_reported_events)(unsafe.Pointer(bp))).Fevent_data, (*(*usbi_reported_events)(unsafe.Pointer(bp))).Fevent_data_count, (*(*usbi_reported_events)(unsafe.Pointer(bp))).Fnum_ready)
	if r != 0 {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__145)), __ccgo_ts+10924, libc.VaList(bp+32, r))
	}
	goto done
done:
	;
	usbi_end_event_handling(tls, ctx)
	return r
}

var __func__145 = [14]uint8{'h', 'a', 'n', 'd', 'l', 'e', '_', 'e', 'v', 'e', 'n', 't', 's'}

func get_next_timeout(tls *libc.TLS, ctx uintptr, tv uintptr, out uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r, v1 int32
	var _ /* timeout at bp+0 */ timeval
	_, _ = r, v1
	r = libusb_get_next_timeout(tls, ctx, bp)
	if r != 0 {
		if !((*timeval)(unsafe.Pointer(bp)).Ftv_sec != 0 || (*timeval)(unsafe.Pointer(bp)).Ftv_usec != 0) {
			return int32(1)
		}
		if (*timeval)(unsafe.Pointer(bp)).Ftv_sec == (*timeval)(unsafe.Pointer(tv)).Ftv_sec {
			v1 = libc.BoolInt32((*timeval)(unsafe.Pointer(bp)).Ftv_usec < (*timeval)(unsafe.Pointer(tv)).Ftv_usec)
		} else {
			v1 = libc.BoolInt32((*timeval)(unsafe.Pointer(bp)).Ftv_sec < (*timeval)(unsafe.Pointer(tv)).Ftv_sec)
		}
		if v1 != 0 {
			*(*timeval)(unsafe.Pointer(out)) = *(*timeval)(unsafe.Pointer(bp))
		} else {
			*(*timeval)(unsafe.Pointer(out)) = *(*timeval)(unsafe.Pointer(tv))
		}
	} else {
		*(*timeval)(unsafe.Pointer(out)) = *(*timeval)(unsafe.Pointer(tv))
	}
	return 0
}

func libusb_handle_events_timeout_completed(tls *libc.TLS, ctx uintptr, tv uintptr, completed uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r int32
	var _ /* poll_timeout at bp+0 */ timeval
	_ = r
	if !((*timeval)(unsafe.Pointer(tv)).Ftv_sec >= 0 && (*timeval)(unsafe.Pointer(tv)).Ftv_usec >= 0 && (*timeval)(unsafe.Pointer(tv)).Ftv_usec < int64(1000000)) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	ctx = usbi_get_context(tls, ctx)
	r = get_next_timeout(tls, ctx, tv, bp)
	if r != 0 {
		handle_timeouts(tls, ctx)
		return 0
	}
	goto retry
retry:
	;
	if libusb_try_lock_events(tls, ctx) == 0 {
		if completed == libc.UintptrFromInt32(0) || !(*(*int32)(unsafe.Pointer(completed)) != 0) {
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__146)), __ccgo_ts+10967, 0)
			r = handle_events(tls, ctx, bp)
		}
		libusb_unlock_events(tls, ctx)
		return r
	}
	libusb_lock_event_waiters(tls, ctx)
	if completed != 0 && *(*int32)(unsafe.Pointer(completed)) != 0 {
		goto already_done
	}
	if !(libusb_event_handler_active(tls, ctx) != 0) {
		libusb_unlock_event_waiters(tls, ctx)
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__146)), __ccgo_ts+10996, 0)
		goto retry
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__146)), __ccgo_ts+11045, 0)
	r = libusb_wait_for_event(tls, ctx, bp)
	goto already_done
already_done:
	;
	libusb_unlock_event_waiters(tls, ctx)
	if r < 0 {
		return r
	} else {
		if r == int32(1) {
			handle_timeouts(tls, ctx)
		}
	}
	return 0
}

var __func__146 = [39]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'h', 'a', 'n', 'd', 'l', 'e', '_', 'e', 'v', 'e', 'n', 't', 's', '_', 't', 'i', 'm', 'e', 'o', 'u', 't', '_', 'c', 'o', 'm', 'p', 'l', 'e', 't', 'e', 'd'}

func libusb_handle_events_timeout(tls *libc.TLS, ctx uintptr, tv uintptr) (r int32) {
	return libusb_handle_events_timeout_completed(tls, ctx, tv, libc.UintptrFromInt32(0))
}

func libusb_handle_events(tls *libc.TLS, ctx uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* tv at bp+0 */ timeval
	(*(*timeval)(unsafe.Pointer(bp))).Ftv_sec = int64(60)
	(*(*timeval)(unsafe.Pointer(bp))).Ftv_usec = 0
	return libusb_handle_events_timeout_completed(tls, ctx, bp, libc.UintptrFromInt32(0))
}

func libusb_handle_events_completed(tls *libc.TLS, ctx uintptr, completed uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _ /* tv at bp+0 */ timeval
	(*(*timeval)(unsafe.Pointer(bp))).Ftv_sec = int64(60)
	(*(*timeval)(unsafe.Pointer(bp))).Ftv_usec = 0
	return libusb_handle_events_timeout_completed(tls, ctx, bp, completed)
}

func libusb_handle_events_locked(tls *libc.TLS, ctx uintptr, tv uintptr) (r1 int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var r int32
	var _ /* poll_timeout at bp+0 */ timeval
	_ = r
	if !((*timeval)(unsafe.Pointer(tv)).Ftv_sec >= 0 && (*timeval)(unsafe.Pointer(tv)).Ftv_usec >= 0 && (*timeval)(unsafe.Pointer(tv)).Ftv_usec < int64(1000000)) {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	ctx = usbi_get_context(tls, ctx)
	r = get_next_timeout(tls, ctx, tv, bp)
	if r != 0 {
		handle_timeouts(tls, ctx)
		return 0
	}
	return handle_events(tls, ctx, bp)
}

func libusb_pollfds_handle_timeouts(tls *libc.TLS, ctx uintptr) (r int32) {
	ctx = usbi_get_context(tls, ctx)
	return usbi_using_timer(tls, ctx)
}

func libusb_get_next_timeout(tls *libc.TLS, ctx uintptr, tv uintptr) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var itransfer uintptr
	var v2 int32
	var v3 suseconds_t
	var _ /* next_timeout at bp+16 */ timespec
	var _ /* systime at bp+0 */ timespec
	_, _, _ = itransfer, v2, v3
	*(*timespec)(unsafe.Pointer(bp + 16)) = timespec{}
	ctx = usbi_get_context(tls, ctx)
	if usbi_using_timer(tls, ctx) != 0 {
		return 0
	}
	usbi_mutex_lock(tls, ctx+208)
	if (*list_head)(unsafe.Pointer(ctx+248)).Fnext == ctx+248 {
		usbi_mutex_unlock(tls, ctx+208)
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__147)), __ccgo_ts+11084, 0)
		return 0
	}
	itransfer = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+248)).Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	for {
		if !(itransfer+48 != ctx+248) {
			break
		}
		if (*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout_flags&libc.Uint32FromInt32(int32(USBI_TRANSFER_TIMEOUT_HANDLED)|int32(USBI_TRANSFER_OS_HANDLES_TIMEOUT)) != 0 {
			goto _1
		}
		if !((*timespec)(unsafe.Pointer(itransfer+80)).Ftv_sec != 0 || (*timespec)(unsafe.Pointer(itransfer+80)).Ftv_nsec != 0) {
			break
		}
		*(*timespec)(unsafe.Pointer(bp + 16)) = (*usbi_transfer)(unsafe.Pointer(itransfer)).Ftimeout
		break
		goto _1
	_1:
		;
		itransfer = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(itransfer)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
	}
	usbi_mutex_unlock(tls, ctx+208)
	if !((*timespec)(unsafe.Pointer(bp+16)).Ftv_sec != 0 || (*timespec)(unsafe.Pointer(bp+16)).Ftv_nsec != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__147)), __ccgo_ts+11105, 0)
		return 0
	}
	usbi_get_monotonic_time(tls, bp)
	if (*timespec)(unsafe.Pointer(bp)).Ftv_sec == (*timespec)(unsafe.Pointer(bp+16)).Ftv_sec {
		v2 = libc.BoolInt32((*timespec)(unsafe.Pointer(bp)).Ftv_nsec < (*timespec)(unsafe.Pointer(bp+16)).Ftv_nsec)
	} else {
		v2 = libc.BoolInt32((*timespec)(unsafe.Pointer(bp)).Ftv_sec < (*timespec)(unsafe.Pointer(bp+16)).Ftv_sec)
	}
	if !(v2 != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__147)), __ccgo_ts+11159, 0)
		v3 = libc.Int64FromInt32(0)
		(*timeval)(unsafe.Pointer(tv)).Ftv_usec = v3
		(*timeval)(unsafe.Pointer(tv)).Ftv_sec = v3
	} else {
		(*timespec)(unsafe.Pointer(bp + 16)).Ftv_sec = (*timespec)(unsafe.Pointer(bp+16)).Ftv_sec - (*timespec)(unsafe.Pointer(bp)).Ftv_sec
		(*timespec)(unsafe.Pointer(bp + 16)).Ftv_nsec = (*timespec)(unsafe.Pointer(bp+16)).Ftv_nsec - (*timespec)(unsafe.Pointer(bp)).Ftv_nsec
		if (*timespec)(unsafe.Pointer(bp+16)).Ftv_nsec < 0 {
			(*timespec)(unsafe.Pointer(bp+16)).Ftv_sec--
			*(*int64)(unsafe.Pointer(bp + 16 + 8)) += int64(1000000000)
		}
		(*timeval)(unsafe.Pointer(tv)).Ftv_sec = (*timespec)(unsafe.Pointer(bp + 16)).Ftv_sec
		(*timeval)(unsafe.Pointer(tv)).Ftv_usec = (*timespec)(unsafe.Pointer(bp+16)).Ftv_nsec / int64(1000)
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__147)), __ccgo_ts+11189, libc.VaList(bp+40, (*timeval)(unsafe.Pointer(tv)).Ftv_sec, (*timeval)(unsafe.Pointer(tv)).Ftv_usec))
	}
	return int32(1)
}

var __func__147 = [24]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'g', 'e', 't', '_', 'n', 'e', 'x', 't', '_', 't', 'i', 'm', 'e', 'o', 'u', 't'}

type __ccgo_fp__Xlibusb_set_pollfd_notifiers_1 = func(*libc.TLS, int32, int16, uintptr)

type __ccgo_fp__Xlibusb_set_pollfd_notifiers_2 = func(*libc.TLS, int32, uintptr)

func libusb_set_pollfd_notifiers(tls *libc.TLS, ctx uintptr, __ccgo_fp_added_cb libusb_pollfd_added_cb, __ccgo_fp_removed_cb libusb_pollfd_removed_cb, user_data uintptr) {
	ctx = usbi_get_context(tls, ctx)
	(*libusb_context)(unsafe.Pointer(ctx)).Ffd_added_cb = __ccgo_fp_added_cb
	(*libusb_context)(unsafe.Pointer(ctx)).Ffd_removed_cb = __ccgo_fp_removed_cb
	(*libusb_context)(unsafe.Pointer(ctx)).Ffd_cb_user_data = user_data
}

func usbi_event_source_notification(tls *libc.TLS, ctx uintptr) {
	var event_flags uint32
	_ = event_flags
	event_flags = (*libusb_context)(unsafe.Pointer(ctx)).Fevent_flags
	*(*uint32)(unsafe.Pointer(ctx + 464)) |= uint32(USBI_EVENT_EVENT_SOURCES_MODIFIED)
	if !(event_flags != 0) {
		usbi_signal_event(tls, ctx+16)
	}
}

func usbi_add_event_source(tls *libc.TLS, ctx uintptr, os_handle usbi_os_handle_t, poll_events int16) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ievent_source uintptr
	_ = ievent_source
	ievent_source = libc.Xmalloc(tls, uint64(24))
	if !(ievent_source != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__148)), __ccgo_ts+11216, libc.VaList(bp+8, os_handle, int32(poll_events)))
	(*usbi_event_source)(unsafe.Pointer(ievent_source)).Fdata.Fos_handle = os_handle
	(*usbi_event_source)(unsafe.Pointer(ievent_source)).Fdata.Fpoll_events = poll_events
	usbi_mutex_lock(tls, ctx+424)
	list_add_tail(tls, ievent_source+8, ctx+472)
	usbi_event_source_notification(tls, ctx)
	usbi_mutex_unlock(tls, ctx+424)
	if (*libusb_context)(unsafe.Pointer(ctx)).Ffd_added_cb != 0 {
		(*(*func(*libc.TLS, int32, int16, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*libusb_context)(unsafe.Pointer(ctx)).Ffd_added_cb})))(tls, os_handle, poll_events, (*libusb_context)(unsafe.Pointer(ctx)).Ffd_cb_user_data)
	}
	return 0
}

var __func__148 = [22]uint8{'u', 's', 'b', 'i', '_', 'a', 'd', 'd', '_', 'e', 'v', 'e', 'n', 't', '_', 's', 'o', 'u', 'r', 'c', 'e'}

func usbi_remove_event_source(tls *libc.TLS, ctx uintptr, os_handle usbi_os_handle_t) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var found int32
	var ievent_source uintptr
	_, _ = found, ievent_source
	found = 0
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__149)), __ccgo_ts+11236, libc.VaList(bp+8, os_handle))
	usbi_mutex_lock(tls, ctx+424)
	ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+472)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	for {
		if !(ievent_source+8 != ctx+472) {
			break
		}
		if (*usbi_event_source)(unsafe.Pointer(ievent_source)).Fdata.Fos_handle == os_handle {
			found = int32(1)
			break
		}
		goto _1
	_1:
		;
		ievent_source = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	}
	if !(found != 0) {
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__149)), __ccgo_ts+11249, libc.VaList(bp+8, os_handle))
		usbi_mutex_unlock(tls, ctx+424)
		return
	}
	list_del(tls, ievent_source+8)
	list_add_tail(tls, ievent_source+8, ctx+488)
	usbi_event_source_notification(tls, ctx)
	usbi_mutex_unlock(tls, ctx+424)
	if (*libusb_context)(unsafe.Pointer(ctx)).Ffd_removed_cb != 0 {
		(*(*func(*libc.TLS, int32, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*libusb_context)(unsafe.Pointer(ctx)).Ffd_removed_cb})))(tls, os_handle, (*libusb_context)(unsafe.Pointer(ctx)).Ffd_cb_user_data)
	}
}

var __func__149 = [25]uint8{'u', 's', 'b', 'i', '_', 'r', 'e', 'm', 'o', 'v', 'e', '_', 'e', 'v', 'e', 'n', 't', '_', 's', 'o', 'u', 'r', 'c', 'e'}

func libusb_get_pollfds(tls *libc.TLS, ctx uintptr) (r uintptr) {
	var i, v3 size_t
	var ievent_source, ret uintptr
	_, _, _, _ = i, ievent_source, ret, v3
	ret = libc.UintptrFromInt32(0)
	ctx = usbi_get_context(tls, ctx)
	usbi_mutex_lock(tls, ctx+424)
	i = uint64(0)
	ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+472)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	for {
		if !(ievent_source+8 != ctx+472) {
			break
		}
		i++
		goto _1
	_1:
		;
		ievent_source = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	}
	ret = libc.Xcalloc(tls, i+uint64(1), uint64(8))
	if !(ret != 0) {
		goto out
	}
	i = uint64(0)
	ievent_source = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+472)).Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	for {
		if !(ievent_source+8 != ctx+472) {
			break
		}
		v3 = i
		i++
		*(*uintptr)(unsafe.Pointer(ret + uintptr(v3)*8)) = ievent_source
		goto _2
	_2:
		;
		ievent_source = uintptr(uint64((*usbi_event_source)(unsafe.Pointer(ievent_source)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+8))
	}
	goto out
out:
	;
	usbi_mutex_unlock(tls, ctx+424)
	return ret
}

func libusb_free_pollfds(tls *libc.TLS, pollfds uintptr) {
	libc.Xfree(tls, pollfds)
}

func usbi_handle_disconnect(tls *libc.TLS, ctx uintptr, dev_handle uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var cur, cur_transfer, to_cancel, transfer_to_cancel uintptr
	_, _, _, _ = cur, cur_transfer, to_cancel, transfer_to_cancel
	usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__150)), __ccgo_ts+11279, libc.VaList(bp+8, libc.Int32FromUint8((*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fbus_number), libc.Int32FromUint8((*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fdevice_address)))
	for int32(1) != 0 {
		to_cancel = libc.UintptrFromInt32(0)
		usbi_mutex_lock(tls, ctx+208)
		cur = uintptr(uint64((*list_head)(unsafe.Pointer(ctx+248)).Fnext) - uint64(libc.UintptrFromInt32(0)+48))
		for {
			if !(cur+48 != ctx+248) {
				break
			}
			cur_transfer = cur + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
			if (*libusb_transfer)(unsafe.Pointer(cur_transfer)).Fdev_handle == dev_handle {
				usbi_mutex_lock(tls, cur)
				if (*usbi_transfer)(unsafe.Pointer(cur)).Fstate_flags&uint32(USBI_TRANSFER_IN_FLIGHT) != 0 {
					to_cancel = cur
				}
				usbi_mutex_unlock(tls, cur)
				if to_cancel != 0 {
					break
				}
			}
			goto _1
		_1:
			;
			cur = uintptr(uint64((*usbi_transfer)(unsafe.Pointer(cur)).Flist.Fnext) - uint64(libc.UintptrFromInt32(0)+48))
		}
		usbi_mutex_unlock(tls, ctx+208)
		if !(to_cancel != 0) {
			break
		}
		transfer_to_cancel = to_cancel + uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1)))
		usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__150)), __ccgo_ts+11292, libc.VaList(bp+8, transfer_to_cancel))
		usbi_mutex_lock(tls, to_cancel)
		(*(*func(*libc.TLS, uintptr))(unsafe.Pointer(&struct{ uintptr }{usbi_backend.Fclear_transfer_priv})))(tls, to_cancel)
		usbi_mutex_unlock(tls, to_cancel)
		usbi_handle_transfer_completion(tls, to_cancel, int32(LIBUSB_TRANSFER_NO_DEVICE))
	}
}

var __func__150 = [23]uint8{'u', 's', 'b', 'i', '_', 'h', 'a', 'n', 'd', 'l', 'e', '_', 'd', 'i', 's', 'c', 'o', 'n', 'n', 'e', 'c', 't'}

var usbi_locale_supported = [6]uintptr{
	0: __ccgo_ts + 11331,
	1: __ccgo_ts + 11334,
	2: __ccgo_ts + 11337,
	3: __ccgo_ts + 11340,
	4: __ccgo_ts + 11343,
	5: __ccgo_ts + 11346,
}
var usbi_localized_errors = [6][14]uintptr{
	0: {
		0:  __ccgo_ts + 11349,
		1:  __ccgo_ts + 11357,
		2:  __ccgo_ts + 11376,
		3:  __ccgo_ts + 11394,
		4:  __ccgo_ts + 11435,
		5:  __ccgo_ts + 11482,
		6:  __ccgo_ts + 11499,
		7:  __ccgo_ts + 11513,
		8:  __ccgo_ts + 11533,
		9:  __ccgo_ts + 11542,
		10: __ccgo_ts + 11553,
		11: __ccgo_ts + 11601,
		12: __ccgo_ts + 11621,
		13: __ccgo_ts + 11679,
	},
	1: {
		0:  __ccgo_ts + 11691,
		1:  __ccgo_ts + 11698,
		2:  __ccgo_ts + 11718,
		3:  __ccgo_ts + 11736,
		4:  __ccgo_ts + 11784,
		5:  __ccgo_ts + 11843,
		6:  __ccgo_ts + 11857,
		7:  __ccgo_ts + 11887,
		8:  __ccgo_ts + 11906,
		9:  __ccgo_ts + 11925,
		10: __ccgo_ts + 11939,
		11: __ccgo_ts + 11966,
		12: __ccgo_ts + 11999,
		13: __ccgo_ts + 12032,
	},
	2: {
		0:  __ccgo_ts + 12044,
		1:  __ccgo_ts + 12052,
		2:  __ccgo_ts + 12076,
		3:  __ccgo_ts + 12096,
		4:  __ccgo_ts + 12139,
		5:  __ccgo_ts + 12192,
		6:  __ccgo_ts + 12213,
		7:  __ccgo_ts + 12238,
		8:  __ccgo_ts + 12257,
		9:  __ccgo_ts + 12270,
		10: __ccgo_ts + 12285,
		11: __ccgo_ts + 12347,
		12: __ccgo_ts + 12369,
		13: __ccgo_ts + 12437,
	},
	3: {
		0:  __ccgo_ts + 12450,
		1:  __ccgo_ts + 12461,
		2:  __ccgo_ts + 12498,
		3:  __ccgo_ts + 12532,
		4:  __ccgo_ts + 12593,
		5:  __ccgo_ts + 12696,
		6:  __ccgo_ts + 12729,
		7:  __ccgo_ts + 12753,
		8:  __ccgo_ts + 12813,
		9:  __ccgo_ts + 12838,
		10: __ccgo_ts + 12864,
		11: __ccgo_ts + 12946,
		12: __ccgo_ts + 12978,
		13: __ccgo_ts + 13063,
	},
	4: {
		0:  __ccgo_ts + 13099,
		1:  __ccgo_ts + 13111,
		2:  __ccgo_ts + 13134,
		3:  __ccgo_ts + 13156,
		4:  __ccgo_ts + 13199,
		5:  __ccgo_ts + 13264,
		6:  __ccgo_ts + 13288,
		7:  __ccgo_ts + 13313,
		8:  __ccgo_ts + 13361,
		9:  __ccgo_ts + 13395,
		10: __ccgo_ts + 13437,
		11: __ccgo_ts + 13486,
		12: __ccgo_ts + 13527,
		13: __ccgo_ts + 13614,
	},
	5: {
		0:  __ccgo_ts + 13633,
		1:  __ccgo_ts + 13641,
		2:  __ccgo_ts + 13659,
		3:  __ccgo_ts + 13684,
		4:  __ccgo_ts + 13709,
		5:  __ccgo_ts + 13757,
		6:  __ccgo_ts + 13773,
		7:  __ccgo_ts + 13796,
		8:  __ccgo_ts + 13812,
		9:  __ccgo_ts + 13827,
		10: __ccgo_ts + 13849,
		11: __ccgo_ts + 13878,
		12: __ccgo_ts + 13899,
		13: __ccgo_ts + 13944,
	},
}
var usbi_error_strings = uintptr(unsafe.Pointer(&usbi_localized_errors))

func libusb_setlocale(tls *libc.TLS, locale uintptr) (r int32) {
	var i size_t
	_ = i
	if !(locale != 0) || libc.Xstrlen(tls, locale) < uint64(2) || libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(locale + 2))) != int32('\000') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(locale + 2))) != int32('-') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(locale + 2))) != int32('_') && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(locale + 2))) != int32('.') {
		return int32(LIBUSB_ERROR_INVALID_PARAM)
	}
	i = uint64(0)
	for {
		if !(i < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(8)) {
			break
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(usbi_locale_supported[i]))) == libc.Xtolower(tls, libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(locale)))) && libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(usbi_locale_supported[i] + 1))) == libc.Xtolower(tls, libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(locale + 1)))) {
			break
		}
		goto _1
	_1:
		;
		i++
	}
	if i == libc.Uint64FromInt64(48)/libc.Uint64FromInt64(8) {
		return int32(LIBUSB_ERROR_NOT_FOUND)
	}
	usbi_error_strings = uintptr(unsafe.Pointer(&usbi_localized_errors)) + uintptr(i)*112
	return int32(LIBUSB_SUCCESS)
}

func libusb_strerror(tls *libc.TLS, errcode int32) (r uintptr) {
	var errcode_index int32
	_ = errcode_index
	errcode_index = -errcode
	if errcode_index < 0 || errcode_index >= int32(14) {
		errcode_index = libc.Int32FromInt32(14) - libc.Int32FromInt32(1)
	}
	return *(*uintptr)(unsafe.Pointer(usbi_error_strings + uintptr(errcode_index)*8))
}

func sync_transfer_cb(tls *libc.TLS, transfer uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var completed, v1 uintptr
	_, _ = completed, v1
	if (*usbi_transfer)(unsafe.Pointer(transfer-uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*usbi_transfer)(unsafe.Pointer(transfer - uintptr((libc.Uint64FromInt64(128)+(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))) & ^(libc.Uint64FromInt64(8)-libc.Uint64FromInt32(1))))).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	usbi_log(tls, v1, int32(LIBUSB_LOG_LEVEL_DEBUG), uintptr(unsafe.Pointer(&__func__151)), __ccgo_ts+13961, libc.VaList(bp+8, (*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length))
	completed = (*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data
	*(*int32)(unsafe.Pointer(completed)) = int32(1)
}

var __func__151 = [17]uint8{'s', 'y', 'n', 'c', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r', '_', 'c', 'b'}

func sync_transfer_wait_for_completion(tls *libc.TLS, transfer uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var completed, ctx, v1 uintptr
	var r int32
	_, _, _, _ = completed, ctx, r, v1
	completed = (*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data
	if (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer((*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	ctx = v1
	for !(*(*int32)(unsafe.Pointer(completed)) != 0) {
		r = libusb_handle_events_completed(tls, ctx, completed)
		if r < 0 {
			if r == int32(LIBUSB_ERROR_INTERRUPTED) {
				continue
			}
			usbi_log(tls, ctx, int32(LIBUSB_LOG_LEVEL_ERROR), uintptr(unsafe.Pointer(&__func__152)), __ccgo_ts+13978, libc.VaList(bp+8, libusb_error_name(tls, r)))
			libusb_cancel_transfer(tls, transfer)
			continue
		}
		if libc.UintptrFromInt32(0) == (*libusb_transfer)(unsafe.Pointer(transfer)).Fdev_handle {
			(*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus = int32(LIBUSB_TRANSFER_NO_DEVICE)
			*(*int32)(unsafe.Pointer(completed)) = int32(1)
		}
	}
}

var __func__152 = [34]uint8{'s', 'y', 'n', 'c', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r', '_', 'w', 'a', 'i', 't', '_', 'f', 'o', 'r', '_', 'c', 'o', 'm', 'p', 'l', 'e', 't', 'i', 'o', 'n'}

func libusb_control_transfer(tls *libc.TLS, dev_handle uintptr, bmRequestType uint8_t, bRequest uint8_t, wValue uint16_t, wIndex uint16_t, data uintptr, wLength uint16_t, timeout uint32) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var buffer, transfer, v1, v2 uintptr
	var r int32
	var _ /* completed at bp+0 */ int32
	_, _, _, _, _ = buffer, r, transfer, v1, v2
	*(*int32)(unsafe.Pointer(bp)) = 0
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	if usbi_handling_events(tls, v1) != 0 {
		return int32(LIBUSB_ERROR_BUSY)
	}
	transfer = libusb_alloc_transfer(tls, 0)
	if !(transfer != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	buffer = libc.Xmalloc(tls, libc.Uint64FromInt64(8)+uint64(wLength))
	if !(buffer != 0) {
		libusb_free_transfer(tls, transfer)
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	libusb_fill_control_setup(tls, buffer, bmRequestType, bRequest, wValue, wIndex, wLength)
	if libc.Int32FromUint8(bmRequestType)&int32(0x80) == int32(LIBUSB_ENDPOINT_OUT) {
		libc.Xmemcpy(tls, buffer+uintptr(libc.Uint64FromInt64(8)), data, uint64(wLength))
	}
	libusb_fill_control_transfer(tls, transfer, dev_handle, buffer, __ccgo_fp(sync_transfer_cb), bp, timeout)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Fflags = uint8(LIBUSB_TRANSFER_FREE_BUFFER)
	r = libusb_submit_transfer(tls, transfer)
	if r < 0 {
		libusb_free_transfer(tls, transfer)
		return r
	}
	sync_transfer_wait_for_completion(tls, transfer)
	if libc.Int32FromUint8(bmRequestType)&int32(0x80) == int32(LIBUSB_ENDPOINT_IN) {
		libc.Xmemcpy(tls, data, libusb_control_transfer_get_data(tls, transfer), libc.Uint64FromInt32((*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length))
	}
	switch (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus {
	case int32(LIBUSB_TRANSFER_COMPLETED):
		r = (*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length
	case int32(LIBUSB_TRANSFER_TIMED_OUT):
		r = int32(LIBUSB_ERROR_TIMEOUT)
	case int32(LIBUSB_TRANSFER_STALL):
		r = int32(LIBUSB_ERROR_PIPE)
	case int32(LIBUSB_TRANSFER_NO_DEVICE):
		r = int32(LIBUSB_ERROR_NO_DEVICE)
	case int32(LIBUSB_TRANSFER_OVERFLOW):
		r = int32(LIBUSB_ERROR_OVERFLOW)
	case int32(LIBUSB_TRANSFER_ERROR):
		fallthrough
	case int32(LIBUSB_TRANSFER_CANCELLED):
		r = int32(LIBUSB_ERROR_IO)
	default:
		if dev_handle != 0 {
			v2 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
		} else {
			v2 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v2, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__153)), __ccgo_ts+14044, libc.VaList(bp+16, (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus))
		r = int32(LIBUSB_ERROR_OTHER)
	}
	libusb_free_transfer(tls, transfer)
	return r
}

var __func__153 = [24]uint8{'l', 'i', 'b', 'u', 's', 'b', '_', 'c', 'o', 'n', 't', 'r', 'o', 'l', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func do_sync_bulk_transfer(tls *libc.TLS, dev_handle uintptr, endpoint uint8, buffer uintptr, length int32, transferred uintptr, timeout uint32, type1 uint8) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var r int32
	var transfer, v1, v3 uintptr
	var v2 bool
	var _ /* completed at bp+0 */ int32
	_, _, _, _, _ = r, transfer, v1, v2, v3
	*(*int32)(unsafe.Pointer(bp)) = 0
	if dev_handle != 0 {
		v1 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	if usbi_handling_events(tls, v1) != 0 {
		return int32(LIBUSB_ERROR_BUSY)
	}
	transfer = libusb_alloc_transfer(tls, 0)
	if !(transfer != 0) {
		return int32(LIBUSB_ERROR_NO_MEM)
	}
	libusb_fill_bulk_transfer(tls, transfer, dev_handle, endpoint, buffer, length, __ccgo_fp(sync_transfer_cb), bp, timeout)
	(*libusb_transfer)(unsafe.Pointer(transfer)).Ftype1 = type1
	r = libusb_submit_transfer(tls, transfer)
	if r < 0 {
		libusb_free_transfer(tls, transfer)
		return r
	}
	sync_transfer_wait_for_completion(tls, transfer)
	if transferred != 0 {
		if v2 = (*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length >= 0; !v2 {
			libc.X__assert_fail(tls, __ccgo_ts+10436, __ccgo_ts+14072, int32(203), uintptr(unsafe.Pointer(&__func__154)))
		}
		_ = v2 || libc.Bool(libc.Int32FromInt32(0) != 0)
		*(*int32)(unsafe.Pointer(transferred)) = (*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length
	}
	switch (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus {
	case int32(LIBUSB_TRANSFER_COMPLETED):
		r = 0
	case int32(LIBUSB_TRANSFER_TIMED_OUT):
		r = int32(LIBUSB_ERROR_TIMEOUT)
	case int32(LIBUSB_TRANSFER_STALL):
		r = int32(LIBUSB_ERROR_PIPE)
	case int32(LIBUSB_TRANSFER_OVERFLOW):
		r = int32(LIBUSB_ERROR_OVERFLOW)
	case int32(LIBUSB_TRANSFER_NO_DEVICE):
		r = int32(LIBUSB_ERROR_NO_DEVICE)
	case int32(LIBUSB_TRANSFER_ERROR):
		fallthrough
	case int32(LIBUSB_TRANSFER_CANCELLED):
		r = int32(LIBUSB_ERROR_IO)
	default:
		if dev_handle != 0 {
			v3 = (*libusb_device)(unsafe.Pointer((*libusb_device_handle)(unsafe.Pointer(dev_handle)).Fdev)).Fctx
		} else {
			v3 = libc.UintptrFromInt32(0)
		}
		usbi_log(tls, v3, int32(LIBUSB_LOG_LEVEL_WARNING), uintptr(unsafe.Pointer(&__func__154)), __ccgo_ts+14044, libc.VaList(bp+16, (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus))
		r = int32(LIBUSB_ERROR_OTHER)
	}
	libusb_free_transfer(tls, transfer)
	return r
}

var __func__154 = [22]uint8{'d', 'o', '_', 's', 'y', 'n', 'c', '_', 'b', 'u', 'l', 'k', '_', 't', 'r', 'a', 'n', 's', 'f', 'e', 'r'}

func libusb_bulk_transfer(tls *libc.TLS, dev_handle uintptr, endpoint uint8, data uintptr, length int32, transferred uintptr, timeout uint32) (r int32) {
	return do_sync_bulk_transfer(tls, dev_handle, endpoint, data, length, transferred, timeout, uint8(LIBUSB_TRANSFER_TYPE_BULK))
}

func libusb_interrupt_transfer(tls *libc.TLS, dev_handle uintptr, endpoint uint8, data uintptr, length int32, transferred uintptr, timeout uint32) (r int32) {
	return do_sync_bulk_transfer(tls, dev_handle, endpoint, data, length, transferred, timeout, uint8(LIBUSB_TRANSFER_TYPE_INTERRUPT))
}

type lconv = struct {
	Fdecimal_point      uintptr
	Fthousands_sep      uintptr
	Fgrouping           uintptr
	Fint_curr_symbol    uintptr
	Fcurrency_symbol    uintptr
	Fmon_decimal_point  uintptr
	Fmon_thousands_sep  uintptr
	Fmon_grouping       uintptr
	Fpositive_sign      uintptr
	Fnegative_sign      uintptr
	Fint_frac_digits    uint8
	Ffrac_digits        uint8
	Fp_cs_precedes      uint8
	Fp_sep_by_space     uint8
	Fn_cs_precedes      uint8
	Fn_sep_by_space     uint8
	Fp_sign_posn        uint8
	Fn_sign_posn        uint8
	Fint_p_cs_precedes  uint8
	Fint_p_sep_by_space uint8
	Fint_n_cs_precedes  uint8
	Fint_n_sep_by_space uint8
	Fint_p_sign_posn    uint8
	Fint_n_sign_posn    uint8
}

type stat1 = struct {
	Fst_dev     dev_t
	Fst_ino     ino_t
	Fst_nlink   nlink_t
	Fst_mode    mode_t
	Fst_uid     uid_t
	Fst_gid     gid_t
	F__pad0     uint32
	Fst_rdev    dev_t
	Fst_size    off_t
	Fst_blksize blksize_t
	Fst_blocks  blkcnt_t
	Fst_atim    timespec
	Fst_mtim    timespec
	Fst_ctim    timespec
	F__unused   [3]int64
}

type statx_timestamp = struct {
	Ftv_sec  int64_t
	Ftv_nsec uint32_t
	F__pad   uint32_t
}

type statx1 = struct {
	Fstx_mask                      uint32_t
	Fstx_blksize                   uint32_t
	Fstx_attributes                uint64_t
	Fstx_nlink                     uint32_t
	Fstx_uid                       uint32_t
	Fstx_gid                       uint32_t
	Fstx_mode                      uint16_t
	F__pad0                        [1]uint16_t
	Fstx_ino                       uint64_t
	Fstx_size                      uint64_t
	Fstx_blocks                    uint64_t
	Fstx_attributes_mask           uint64_t
	Fstx_atime                     statx_timestamp
	Fstx_btime                     statx_timestamp
	Fstx_ctime                     statx_timestamp
	Fstx_mtime                     statx_timestamp
	Fstx_rdev_major                uint32_t
	Fstx_rdev_minor                uint32_t
	Fstx_dev_major                 uint32_t
	Fstx_dev_minor                 uint32_t
	Fstx_mnt_id                    uint64_t
	Fstx_dio_mem_align             uint32_t
	Fstx_dio_offset_align          uint32_t
	Fstx_subvol                    uint64_t
	Fstx_atomic_write_unit_min     uint32_t
	Fstx_atomic_write_unit_max     uint32_t
	Fstx_atomic_write_segments_max uint32_t
	F__pad1                        [1]uint32_t
	F__pad2                        [9]uint64_t
}

type wint_t = uint32

type wctype_t = uint64

type mbstate_t = struct {
	F__opaque1 uint32
	F__opaque2 uint32
}

type __mbstate_t = mbstate_t

type iconv_t = uintptr

type hid_api_version = struct {
	Fmajor int32
	Fminor int32
	Fpatch int32
}

type hid_device_ = struct {
	Fdevice_handle            uintptr
	Fconfig_number            int32
	Finterface1               int32
	Freport_descriptor_size   uint16_t
	Finput_endpoint           int32
	Foutput_endpoint          int32
	Finput_ep_max_packet_size int32
	Fmanufacturer_index       int32
	Fproduct_index            int32
	Fserial_index             int32
	Fdevice_info              uintptr
	Fblocking                 int32
	Fthread_state             hidapi_thread_state
	Fshutdown_thread          int32
	Ftransfer_loop_finished   int32
	Ftransfer                 uintptr
	Finput_reports            uintptr
	Fis_driver_detached       int32
}

type hid_device = struct {
	Fdevice_handle            uintptr
	Fconfig_number            int32
	Finterface1               int32
	Freport_descriptor_size   uint16_t
	Finput_endpoint           int32
	Foutput_endpoint          int32
	Finput_ep_max_packet_size int32
	Fmanufacturer_index       int32
	Fproduct_index            int32
	Fserial_index             int32
	Fdevice_info              uintptr
	Fblocking                 int32
	Fthread_state             hidapi_thread_state
	Fshutdown_thread          int32
	Ftransfer_loop_finished   int32
	Ftransfer                 uintptr
	Finput_reports            uintptr
	Fis_driver_detached       int32
}

type hid_bus_type = int32

const HID_API_BUS_UNKNOWN = 0
const HID_API_BUS_USB = 1
const HID_API_BUS_BLUETOOTH = 2
const HID_API_BUS_I2C = 3
const HID_API_BUS_SPI = 4
const HID_API_BUS_VIRTUAL = 5

type hid_device_info = struct {
	Fpath                uintptr
	Fvendor_id           uint16
	Fproduct_id          uint16
	Fserial_number       uintptr
	Frelease_number      uint16
	Fmanufacturer_string uintptr
	Fproduct_string      uintptr
	Fusage_page          uint16
	Fusage               uint16
	Finterface_number    int32
	Fnext                uintptr
	Fbus_type            hid_bus_type
}

type hidapi_timespec = struct {
	Ftv_sec  time_t
	Ftv_nsec int64
}

type hidapi_thread_state = struct {
	Fthread    pthread_t
	Fmutex     pthread_mutex_t
	Fcondition pthread_cond_t
	Fbarrier   pthread_barrier_t
}

func hidapi_thread_state_init(tls *libc.TLS, state uintptr) {
	libc.Xpthread_mutex_init(tls, state+8, libc.UintptrFromInt32(0))
	libc.Xpthread_cond_init(tls, state+48, libc.UintptrFromInt32(0))
	libc.Xpthread_barrier_init(tls, state+96, libc.UintptrFromInt32(0), uint32(2))
}

func hidapi_thread_state_destroy(tls *libc.TLS, state uintptr) {
	libc.Xpthread_barrier_destroy(tls, state+96)
	libc.Xpthread_cond_destroy(tls, state+48)
	libc.Xpthread_mutex_destroy(tls, state+8)
}

func hidapi_thread_mutex_lock(tls *libc.TLS, state uintptr) {
	libc.Xpthread_mutex_lock(tls, state+8)
}

func hidapi_thread_mutex_unlock(tls *libc.TLS, state uintptr) {
	libc.Xpthread_mutex_unlock(tls, state+8)
}

func hidapi_thread_cond_wait(tls *libc.TLS, state uintptr) {
	libc.Xpthread_cond_wait(tls, state+48, state+8)
}

func hidapi_thread_cond_timedwait(tls *libc.TLS, state uintptr, ts uintptr) (r int32) {
	return libc.Xpthread_cond_timedwait(tls, state+48, state+8, ts)
}

func hidapi_thread_cond_signal(tls *libc.TLS, state uintptr) {
	libc.Xpthread_cond_signal(tls, state+48)
}

func hidapi_thread_cond_broadcast(tls *libc.TLS, state uintptr) {
	libc.Xpthread_cond_broadcast(tls, state+48)
}

func hidapi_thread_barrier_wait(tls *libc.TLS, state uintptr) {
	libc.Xpthread_barrier_wait(tls, state+96)
}

func hidapi_thread_create(tls *libc.TLS, state uintptr, __ccgo_fp_func uintptr, func_arg uintptr) {
	libc.Xpthread_create(tls, state, libc.UintptrFromInt32(0), __ccgo_fp_func, func_arg)
}

func hidapi_thread_join(tls *libc.TLS, state uintptr) {
	libc.Xpthread_join(tls, (*hidapi_thread_state)(unsafe.Pointer(state)).Fthread, libc.UintptrFromInt32(0))
}

func hidapi_thread_gettime(tls *libc.TLS, ts uintptr) {
	libc.Xclock_gettime(tls, 0, ts)
}

func hidapi_thread_addtime(tls *libc.TLS, ts uintptr, milliseconds int32) {
	*(*time_t)(unsafe.Pointer(ts)) += int64(milliseconds / int32(1000))
	*(*int64)(unsafe.Pointer(ts + 8)) += int64(milliseconds % int32(1000) * int32(1000000))
	if (*hidapi_timespec)(unsafe.Pointer(ts)).Ftv_nsec >= int64(1000000000) {
		(*hidapi_timespec)(unsafe.Pointer(ts)).Ftv_sec++
		*(*int64)(unsafe.Pointer(ts + 8)) -= int64(1000000000)
	}
}

type input_report = struct {
	Fdata uintptr
	Flen1 size_t
	Fnext uintptr
}

var api_version = hid_api_version{
	Fminor: int32(16),
}
var usb_context = libc.UintptrFromInt32(0)

func new_hid_device(tls *libc.TLS) (r uintptr) {
	var dev uintptr
	_ = dev
	dev = libc.Xcalloc(tls, uint64(1), uint64(224))
	if !(dev != 0) {
		return libc.UintptrFromInt32(0)
	}
	(*hid_device)(unsafe.Pointer(dev)).Fblocking = int32(1)
	hidapi_thread_state_init(tls, dev+64)
	return dev
}

func free_hid_device(tls *libc.TLS, dev uintptr) {
	hidapi_thread_state_destroy(tls, dev+64)
	hid_free_enumeration(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_info)
	libc.Xfree(tls, dev)
}

func get_bytes(tls *libc.TLS, rpt uintptr, len1 size_t, num_bytes size_t, cur size_t) (r uint32_t) {
	if cur+num_bytes >= len1 {
		return uint32(0)
	}
	if num_bytes == uint64(0) {
		return uint32(0)
	} else {
		if num_bytes == uint64(1) {
			return uint32(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(1)))))
		} else {
			if num_bytes == uint64(2) {
				return libc.Uint32FromInt32(libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(2)))))*libc.Int32FromInt32(256) + libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(1))))))
			} else {
				if num_bytes == uint64(4) {
					return libc.Uint32FromInt32(libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(4)))))*libc.Int32FromInt32(0x01000000) + libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(3)))))*libc.Int32FromInt32(0x00010000) + libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(2)))))*libc.Int32FromInt32(0x00000100) + libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(rpt + uintptr(cur+uint64(1)))))*libc.Int32FromInt32(0x00000001))
				} else {
					return uint32(0)
				}
			}
		}
	}
	return r
}

func get_usage(tls *libc.TLS, report_descriptor uintptr, size size_t, usage_page uintptr, usage uintptr) (r int32) {
	var data_len, key, key_cmd, key_size, size_code, usage_found, usage_page_found int32
	var i uint32
	_, _, _, _, _, _, _, _ = data_len, i, key, key_cmd, key_size, size_code, usage_found, usage_page_found
	i = uint32(0)
	usage_found = 0
	usage_page_found = 0
	for uint64(i) < size {
		key = libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(report_descriptor + uintptr(i))))
		key_cmd = key & int32(0xfc)
		if key&int32(0xf0) == int32(0xf0) {
			if uint64(i+uint32(1)) < size {
				data_len = libc.Int32FromUint8(*(*uint8_t)(unsafe.Pointer(report_descriptor + uintptr(i+uint32(1)))))
			} else {
				data_len = 0
			}
			key_size = int32(3)
		} else {
			size_code = key & int32(0x3)
			switch size_code {
			case 0:
				fallthrough
			case int32(1):
				fallthrough
			case int32(2):
				data_len = size_code
			case int32(3):
				data_len = int32(4)
			default:
				data_len = 0
				break
			}
			key_size = int32(1)
		}
		if key_cmd == int32(0x4) {
			*(*uint16)(unsafe.Pointer(usage_page)) = uint16(get_bytes(tls, report_descriptor, size, libc.Uint64FromInt32(data_len), uint64(i)))
			usage_page_found = int32(1)
		}
		if key_cmd == int32(0x8) {
			if data_len == int32(4) {
				*(*uint16)(unsafe.Pointer(usage_page)) = uint16(get_bytes(tls, report_descriptor, size, uint64(2), uint64(i+uint32(2))))
				usage_page_found = int32(1)
				*(*uint16)(unsafe.Pointer(usage)) = uint16(get_bytes(tls, report_descriptor, size, uint64(2), uint64(i)))
				usage_found = int32(1)
			} else {
				*(*uint16)(unsafe.Pointer(usage)) = uint16(get_bytes(tls, report_descriptor, size, libc.Uint64FromInt32(data_len), uint64(i)))
				usage_found = int32(1)
			}
		}
		if usage_page_found != 0 && usage_found != 0 {
			return 0
		}
		i += libc.Uint32FromInt32(data_len + key_size)
	}
	return -int32(1)
}

func get_first_language(tls *libc.TLS, dev uintptr) (r uint16_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var len1 int32
	var _ /* buf at bp+0 */ [32]uint16_t
	_ = len1
	len1 = libusb_get_string_descriptor(tls, dev, uint8(0x0), uint16(0x0), bp, int32(64))
	if len1 < int32(4) {
		return uint16(0x0)
	}
	return (*(*[32]uint16_t)(unsafe.Pointer(bp)))[int32(1)]
}

func is_language_supported(tls *libc.TLS, dev uintptr, lang uint16_t) (r int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var i, len1 int32
	var _ /* buf at bp+0 */ [32]uint16_t
	_, _ = i, len1
	len1 = libusb_get_string_descriptor(tls, dev, uint8(0x0), uint16(0x0), bp, int32(64))
	if len1 < int32(4) {
		return 0x0
	}
	len1 /= int32(2)
	i = int32(1)
	for {
		if !(i < len1) {
			break
		}
		if libc.Int32FromUint16((*(*[32]uint16_t)(unsafe.Pointer(bp)))[i]) == libc.Int32FromUint16(lang) {
			return int32(1)
		}
		goto _1
	_1:
		;
		i++
	}
	return 0
}

func get_usb_string(tls *libc.TLS, dev uintptr, idx uint8_t) (r uintptr) {
	bp := tls.Alloc(1568)
	defer tls.Free(1568)
	var ic iconv_t
	var lang uint16_t
	var len1 int32
	var res size_t
	var str uintptr
	var _ /* buf at bp+0 */ [512]uint8
	var _ /* inbytes at bp+1536 */ size_t
	var _ /* inptr at bp+1552 */ uintptr
	var _ /* outbytes at bp+1544 */ size_t
	var _ /* outptr at bp+1560 */ uintptr
	var _ /* wbuf at bp+512 */ [256]wchar_t
	_, _, _, _, _ = ic, lang, len1, res, str
	str = libc.UintptrFromInt32(0)
	lang = get_usb_code_for_current_locale(tls)
	if !(is_language_supported(tls, dev, lang) != 0) {
		lang = get_first_language(tls, dev)
	}
	len1 = libusb_get_string_descriptor(tls, dev, idx, lang, bp, int32(512))
	if len1 < int32(2) {
		return libc.UintptrFromInt32(0)
	}
	ic = libc.Xiconv_open(tls, __ccgo_ts+14113, __ccgo_ts+14121)
	if ic == uintptr(-libc.Int32FromInt32(1)) {
		return libc.UintptrFromInt32(0)
	}
	*(*uintptr)(unsafe.Pointer(bp + 1552)) = bp + uintptr(2)
	*(*size_t)(unsafe.Pointer(bp + 1536)) = libc.Uint64FromInt32(len1 - int32(2))
	*(*uintptr)(unsafe.Pointer(bp + 1560)) = bp + 512
	*(*size_t)(unsafe.Pointer(bp + 1544)) = uint64(1024)
	res = libc.Xiconv(tls, ic, bp+1552, bp+1536, bp+1560, bp+1544)
	if res == libc.Uint64FromInt32(-libc.Int32FromInt32(1)) {
		goto err
	}
	(*(*[256]wchar_t)(unsafe.Pointer(bp + 512)))[libc.Uint64FromInt64(1024)/libc.Uint64FromInt64(4)-libc.Uint64FromInt32(1)] = 0x00000000
	if *(*size_t)(unsafe.Pointer(bp + 1544)) >= uint64(4) {
		*(*wchar_t)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 1560)))) = 0x00000000
	}
	str = libc.Xwcsdup(tls, bp+512)
	goto err
err:
	;
	libc.Xiconv_close(tls, ic)
	return str
}

func get_path(tls *libc.TLS, result uintptr, dev uintptr, config_number int32, interface_number int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i uint8_t
	var n, num_ports int32
	var str uintptr
	var _ /* port_numbers at bp+0 */ [8]uint8_t
	_, _, _, _ = i, n, num_ports, str
	str = result
	*(*[8]uint8_t)(unsafe.Pointer(bp)) = [8]uint8_t{}
	num_ports = libusb_get_port_numbers(tls, dev, bp, int32(8))
	if num_ports > 0 {
		n = libc.X__builtin_snprintf(tls, str, uint64(8), __ccgo_ts+14130, libc.VaList(bp+16, libc.Int32FromUint8(libusb_get_bus_number(tls, dev)), libc.Int32FromUint8((*(*[8]uint8_t)(unsafe.Pointer(bp)))[0])))
		i = uint8(1)
		for {
			if !(libc.Int32FromUint8(i) < num_ports) {
				break
			}
			n += libc.X__builtin_snprintf(tls, str+uintptr(n), uint64(5), __ccgo_ts+14136, libc.VaList(bp+16, libc.Int32FromUint8((*(*[8]uint8_t)(unsafe.Pointer(bp)))[i])))
			goto _1
		_1:
			;
			i++
		}
		n += libc.X__builtin_snprintf(tls, str+uintptr(n), uint64(9), __ccgo_ts+14140, libc.VaList(bp+16, libc.Int32FromUint8(libc.Uint8FromInt32(config_number)), libc.Int32FromUint8(libc.Uint8FromInt32(interface_number))))
		*(*uint8)(unsafe.Pointer(str + uintptr(n))) = uint8('\000')
	} else {
		if num_ports == int32(LIBUSB_ERROR_OVERFLOW) {
		} else {
		}
		*(*uint8)(unsafe.Pointer(str)) = uint8('\000')
	}
}

func make_path(tls *libc.TLS, dev uintptr, config_number int32, interface_number int32) (r uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var _ /* str at bp+0 */ [64]uint8
	get_path(tls, bp, dev, config_number, interface_number)
	return libc.Xstrdup(tls, bp)
}

func hid_version(tls *libc.TLS) (r uintptr) {
	return uintptr(unsafe.Pointer(&api_version))
}

func hid_version_str(tls *libc.TLS) (r uintptr) {
	return __ccgo_ts + 14147
}

func hid_init(tls *libc.TLS) (r int32) {
	var locale uintptr
	_ = locale
	if !(usb_context != 0) {
		if libusb_init(tls, uintptr(unsafe.Pointer(&usb_context))) != 0 {
			return -int32(1)
		}
		locale = libc.Xsetlocale(tls, 0, libc.UintptrFromInt32(0))
		if !(locale != 0) {
			libc.Xsetlocale(tls, 0, __ccgo_ts+14154)
		}
	}
	return 0
}

func hid_exit(tls *libc.TLS) (r int32) {
	if usb_context != 0 {
		libusb_exit(tls, usb_context)
		usb_context = libc.UintptrFromInt32(0)
	}
	return 0
}

func hid_get_report_descriptor_libusb(tls *libc.TLS, handle uintptr, interface_num int32, expected_report_descriptor_size uint16_t, buf uintptr, buf_size size_t) (r int32) {
	bp := tls.Alloc(4096)
	defer tls.Free(4096)
	var res int32
	var _ /* tmp at bp+0 */ [4096]uint8
	_ = res
	if libc.Int32FromUint16(expected_report_descriptor_size) > int32(4096) {
		expected_report_descriptor_size = uint16(4096)
	}
	res = libusb_control_transfer(tls, handle, libc.Uint8FromInt32(int32(LIBUSB_ENDPOINT_IN)|int32(LIBUSB_RECIPIENT_INTERFACE)), uint8(LIBUSB_REQUEST_GET_DESCRIPTOR), libc.Uint16FromInt32(int32(LIBUSB_DT_REPORT)<<libc.Int32FromInt32(8)), libc.Uint16FromInt32(interface_num), bp, expected_report_descriptor_size, uint32(5000))
	if res < 0 {
		return -int32(1)
	}
	if res > libc.Int32FromUint64(buf_size) {
		res = libc.Int32FromUint64(buf_size)
	}
	libc.Xmemcpy(tls, buf, bp, libc.Uint64FromInt32(res))
	return res
}

func fill_device_info_usage(tls *libc.TLS, cur_dev uintptr, handle uintptr, interface_num int32, expected_report_descriptor_size uint16_t) {
	bp := tls.Alloc(4112)
	defer tls.Free(4112)
	var res int32
	var _ /* hid_report_descriptor at bp+0 */ [4096]uint8
	var _ /* page at bp+4096 */ uint16
	var _ /* usage at bp+4098 */ uint16
	_ = res
	*(*uint16)(unsafe.Pointer(bp + 4096)) = uint16(0)
	*(*uint16)(unsafe.Pointer(bp + 4098)) = uint16(0)
	res = hid_get_report_descriptor_libusb(tls, handle, interface_num, expected_report_descriptor_size, bp, uint64(4096))
	if res >= 0 {
		get_usage(tls, bp, libc.Uint64FromInt32(res), bp+4096, bp+4098)
	}
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Fusage_page = *(*uint16)(unsafe.Pointer(bp + 4096))
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Fusage = *(*uint16)(unsafe.Pointer(bp + 4098))
}

func create_device_info_for_device(tls *libc.TLS, device uintptr, handle uintptr, desc uintptr, config_number int32, interface_num int32) (r uintptr) {
	var cur_dev uintptr
	_ = cur_dev
	cur_dev = libc.Xcalloc(tls, uint64(1), uint64(72))
	if cur_dev == libc.UintptrFromInt32(0) {
		return libc.UintptrFromInt32(0)
	}
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Fvendor_id = (*libusb_device_descriptor)(unsafe.Pointer(desc)).FidVendor
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Fproduct_id = (*libusb_device_descriptor)(unsafe.Pointer(desc)).FidProduct
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Frelease_number = (*libusb_device_descriptor)(unsafe.Pointer(desc)).FbcdDevice
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Finterface_number = interface_num
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Fbus_type = int32(HID_API_BUS_USB)
	(*hid_device_info)(unsafe.Pointer(cur_dev)).Fpath = make_path(tls, device, config_number, interface_num)
	if !(handle != 0) {
		return cur_dev
	}
	if libc.Int32FromUint8((*libusb_device_descriptor)(unsafe.Pointer(desc)).FiSerialNumber) > 0 {
		(*hid_device_info)(unsafe.Pointer(cur_dev)).Fserial_number = get_usb_string(tls, handle, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FiSerialNumber)
	}
	if libc.Int32FromUint8((*libusb_device_descriptor)(unsafe.Pointer(desc)).FiManufacturer) > 0 {
		(*hid_device_info)(unsafe.Pointer(cur_dev)).Fmanufacturer_string = get_usb_string(tls, handle, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FiManufacturer)
	}
	if libc.Int32FromUint8((*libusb_device_descriptor)(unsafe.Pointer(desc)).FiProduct) > 0 {
		(*hid_device_info)(unsafe.Pointer(cur_dev)).Fproduct_string = get_usb_string(tls, handle, (*libusb_device_descriptor)(unsafe.Pointer(desc)).FiProduct)
	}
	return cur_dev
}

func get_report_descriptor_size_from_interface_descriptors(tls *libc.TLS, intf_desc uintptr) (r uint16_t) {
	var bNumDescriptors uint8
	var extra uintptr
	var extra_length, found_hid_report_descriptor, i int32
	var result uint16_t
	_, _, _, _, _, _ = bNumDescriptors, extra, extra_length, found_hid_report_descriptor, i, result
	i = 0
	found_hid_report_descriptor = 0
	result = uint16(4096)
	extra = (*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).Fextra
	extra_length = (*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).Fextra_length
	for extra_length >= int32(2) {
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(extra + 1))) == int32(LIBUSB_DT_HID) {
			if extra_length < int32(6) {
				break
			}
			bNumDescriptors = *(*uint8)(unsafe.Pointer(extra + 5))
			if extra_length < int32(6)+int32(3)*libc.Int32FromUint8(bNumDescriptors) {
				break
			}
			i = 0
			for {
				if !(i < libc.Int32FromUint8(bNumDescriptors)) {
					break
				}
				if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(extra + uintptr(int32(6)+int32(3)*i)))) == int32(LIBUSB_DT_REPORT) {
					result = libc.Uint16FromInt32(libc.Int32FromUint16(uint16(*(*uint8)(unsafe.Pointer(extra + uintptr(int32(6)+int32(3)*i+int32(2))))))<<int32(8) | libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(extra + uintptr(int32(6)+int32(3)*i+int32(1))))))
					found_hid_report_descriptor = int32(1)
					break
				}
				goto _1
			_1:
				;
				i++
			}
			if !(found_hid_report_descriptor != 0) {
			}
			break
		}
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(extra))) == 0 {
			break
		}
		extra_length -= libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(extra)))
		extra += uintptr(*(*uint8)(unsafe.Pointer(extra)))
	}
	return result
}

func is_xbox360(tls *libc.TLS, vendor_id uint16, intf_desc uintptr) (r int32) {
	var i size_t
	_ = i
	if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceClass) == int32(LIBUSB_CLASS_VENDOR_SPEC) && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceSubClass) == xb360_iface_subclass && (libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceProtocol) == xb360_iface_protocol || libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceProtocol) == xb360w_iface_protocol) {
		i = uint64(0)
		for {
			if !(i < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(4)) {
				break
			}
			if libc.Int32FromUint16(vendor_id) == supported_vendors[i] {
				return int32(1)
			}
			goto _1
		_1:
			;
			i++
		}
	}
	return 0
}

var xb360_iface_subclass = int32(93)

var xb360_iface_protocol = int32(1)

var xb360w_iface_protocol = int32(129)

var supported_vendors = [26]int32{
	0:  int32(0x0079),
	1:  int32(0x044f),
	2:  int32(0x045e),
	3:  int32(0x046d),
	4:  int32(0x056e),
	5:  int32(0x06a3),
	6:  int32(0x0738),
	7:  int32(0x07ff),
	8:  int32(0x0e6f),
	9:  int32(0x0f0d),
	10: int32(0x1038),
	11: int32(0x11c9),
	12: int32(0x12ab),
	13: int32(0x1430),
	14: int32(0x146b),
	15: int32(0x1532),
	16: int32(0x15e4),
	17: int32(0x162e),
	18: int32(0x1689),
	19: int32(0x1949),
	20: int32(0x1bad),
	21: int32(0x20d6),
	22: int32(0x24c6),
	23: int32(0x2c22),
	24: int32(0x2dc8),
	25: int32(0x9886),
}

func is_xboxone(tls *libc.TLS, vendor_id uint16, intf_desc uintptr) (r int32) {
	var i size_t
	_ = i
	if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber) == 0 && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceClass) == int32(LIBUSB_CLASS_VENDOR_SPEC) && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceSubClass) == xb1_iface_subclass && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceProtocol) == xb1_iface_protocol {
		i = uint64(0)
		for {
			if !(i < libc.Uint64FromInt64(48)/libc.Uint64FromInt64(4)) {
				break
			}
			if libc.Int32FromUint16(vendor_id) == supported_vendors1[i] {
				return int32(1)
			}
			goto _1
		_1:
			;
			i++
		}
	}
	return 0
}

var xb1_iface_subclass = int32(71)

var xb1_iface_protocol = int32(208)

var supported_vendors1 = [12]int32{
	0:  int32(0x044f),
	1:  int32(0x045e),
	2:  int32(0x0738),
	3:  int32(0x0e6f),
	4:  int32(0x0f0d),
	5:  int32(0x10f5),
	6:  int32(0x1532),
	7:  int32(0x20d6),
	8:  int32(0x24c6),
	9:  int32(0x2dc8),
	10: int32(0x2e24),
	11: int32(0x3537),
}

func should_enumerate_interface(tls *libc.TLS, vendor_id uint16, intf_desc uintptr) (r int32) {
	if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceClass) == int32(LIBUSB_CLASS_HID) {
		return int32(1)
	}
	if is_xbox360(tls, vendor_id, intf_desc) != 0 {
		return int32(1)
	}
	if is_xboxone(tls, vendor_id, intf_desc) != 0 {
		return int32(1)
	}
	return 0
}

func hid_enumerate(tls *libc.TLS, vendor_id uint16, product_id uint16) (r uintptr) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var cur_dev, dev, intf, intf_desc, root, tmp, v1 uintptr
	var dev_pid, dev_vid uint16
	var i, j, k, res, v2 int32
	var num_devs ssize_t
	var _ /* conf_desc at bp+40 */ uintptr
	var _ /* desc at bp+16 */ libusb_device_descriptor
	var _ /* devs at bp+0 */ uintptr
	var _ /* handle at bp+8 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cur_dev, dev, dev_pid, dev_vid, i, intf, intf_desc, j, k, num_devs, res, root, tmp, v1, v2
	*(*uintptr)(unsafe.Pointer(bp + 8)) = libc.UintptrFromInt32(0)
	i = 0
	root = libc.UintptrFromInt32(0)
	cur_dev = libc.UintptrFromInt32(0)
	if hid_init(tls) < 0 {
		return libc.UintptrFromInt32(0)
	}
	num_devs = libusb_get_device_list(tls, usb_context, bp)
	if num_devs < 0 {
		return libc.UintptrFromInt32(0)
	}
	for {
		v2 = i
		i++
		v1 = *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + uintptr(v2)*8))
		dev = v1
		if !(v1 != libc.UintptrFromInt32(0)) {
			break
		}
		*(*uintptr)(unsafe.Pointer(bp + 40)) = libc.UintptrFromInt32(0)
		res = libusb_get_device_descriptor(tls, dev, bp+16)
		if res < 0 {
			continue
		}
		dev_vid = (*(*libusb_device_descriptor)(unsafe.Pointer(bp + 16))).FidVendor
		dev_pid = (*(*libusb_device_descriptor)(unsafe.Pointer(bp + 16))).FidProduct
		if libc.Int32FromUint16(vendor_id) != 0x0 && libc.Int32FromUint16(vendor_id) != libc.Int32FromUint16(dev_vid) || libc.Int32FromUint16(product_id) != 0x0 && libc.Int32FromUint16(product_id) != libc.Int32FromUint16(dev_pid) {
			continue
		}
		res = libusb_get_active_config_descriptor(tls, dev, bp+40)
		if res < 0 {
			libusb_get_config_descriptor(tls, dev, uint8(0), bp+40)
		}
		if *(*uintptr)(unsafe.Pointer(bp + 40)) != 0 {
			j = 0
			for {
				if !(j < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).FbNumInterfaces)) {
					break
				}
				intf = (*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).Finterface1 + uintptr(j)*16
				k = 0
				for {
					if !(k < (*libusb_interface)(unsafe.Pointer(intf)).Fnum_altsetting) {
						break
					}
					intf_desc = (*libusb_interface)(unsafe.Pointer(intf)).Faltsetting + uintptr(k)*40
					if should_enumerate_interface(tls, dev_vid, intf_desc) != 0 {
						res = libusb_open(tls, dev, bp+8)
						tmp = create_device_info_for_device(tls, dev, *(*uintptr)(unsafe.Pointer(bp + 8)), bp+16, libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 40)))).FbConfigurationValue), libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
						if tmp != 0 {
							if cur_dev != 0 {
								(*hid_device_info)(unsafe.Pointer(cur_dev)).Fnext = tmp
							} else {
								root = tmp
							}
							cur_dev = tmp
						}
						if res >= 0 {
							libusb_close(tls, *(*uintptr)(unsafe.Pointer(bp + 8)))
							*(*uintptr)(unsafe.Pointer(bp + 8)) = libc.UintptrFromInt32(0)
						}
						break
					}
					goto _4
				_4:
					;
					k++
				}
				goto _3
			_3:
				;
				j++
			}
			libusb_free_config_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp + 40)))
		}
	}
	libusb_free_device_list(tls, *(*uintptr)(unsafe.Pointer(bp)), int32(1))
	return root
}

func hid_free_enumeration(tls *libc.TLS, devs uintptr) {
	var d, next uintptr
	_, _ = d, next
	d = devs
	for d != 0 {
		next = (*hid_device_info)(unsafe.Pointer(d)).Fnext
		libc.Xfree(tls, (*hid_device_info)(unsafe.Pointer(d)).Fpath)
		libc.Xfree(tls, (*hid_device_info)(unsafe.Pointer(d)).Fserial_number)
		libc.Xfree(tls, (*hid_device_info)(unsafe.Pointer(d)).Fmanufacturer_string)
		libc.Xfree(tls, (*hid_device_info)(unsafe.Pointer(d)).Fproduct_string)
		libc.Xfree(tls, d)
		d = next
	}
}

func hid_open(tls *libc.TLS, vendor_id uint16, product_id uint16, serial_number uintptr) (r uintptr) {
	var cur_dev, devs, handle, path_to_open uintptr
	_, _, _, _ = cur_dev, devs, handle, path_to_open
	path_to_open = libc.UintptrFromInt32(0)
	handle = libc.UintptrFromInt32(0)
	devs = hid_enumerate(tls, vendor_id, product_id)
	cur_dev = devs
	for cur_dev != 0 {
		if libc.Int32FromUint16((*hid_device_info)(unsafe.Pointer(cur_dev)).Fvendor_id) == libc.Int32FromUint16(vendor_id) && libc.Int32FromUint16((*hid_device_info)(unsafe.Pointer(cur_dev)).Fproduct_id) == libc.Int32FromUint16(product_id) {
			if serial_number != 0 {
				if (*hid_device_info)(unsafe.Pointer(cur_dev)).Fserial_number != 0 && libc.Xwcscmp(tls, serial_number, (*hid_device_info)(unsafe.Pointer(cur_dev)).Fserial_number) == 0 {
					path_to_open = (*hid_device_info)(unsafe.Pointer(cur_dev)).Fpath
					break
				}
			} else {
				path_to_open = (*hid_device_info)(unsafe.Pointer(cur_dev)).Fpath
				break
			}
		}
		cur_dev = (*hid_device_info)(unsafe.Pointer(cur_dev)).Fnext
	}
	if path_to_open != 0 {
		handle = hid_open_path(tls, path_to_open)
	}
	hid_free_enumeration(tls, devs)
	return handle
}

func read_callback(tls *libc.TLS, transfer uintptr) {
	var cur, dev, rpt uintptr
	var num_queued, res int32
	_, _, _, _, _ = cur, dev, num_queued, res, rpt
	dev = (*libusb_transfer)(unsafe.Pointer(transfer)).Fuser_data
	if (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus == int32(LIBUSB_TRANSFER_COMPLETED) {
		rpt = libc.Xmalloc(tls, uint64(24))
		(*input_report)(unsafe.Pointer(rpt)).Fdata = libc.Xmalloc(tls, libc.Uint64FromInt32((*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length))
		libc.Xmemcpy(tls, (*input_report)(unsafe.Pointer(rpt)).Fdata, (*libusb_transfer)(unsafe.Pointer(transfer)).Fbuffer, libc.Uint64FromInt32((*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length))
		(*input_report)(unsafe.Pointer(rpt)).Flen1 = libc.Uint64FromInt32((*libusb_transfer)(unsafe.Pointer(transfer)).Factual_length)
		(*input_report)(unsafe.Pointer(rpt)).Fnext = libc.UintptrFromInt32(0)
		hidapi_thread_mutex_lock(tls, dev+64)
		if (*hid_device)(unsafe.Pointer(dev)).Finput_reports == libc.UintptrFromInt32(0) {
			(*hid_device)(unsafe.Pointer(dev)).Finput_reports = rpt
			hidapi_thread_cond_signal(tls, dev+64)
		} else {
			cur = (*hid_device)(unsafe.Pointer(dev)).Finput_reports
			num_queued = 0
			for (*input_report)(unsafe.Pointer(cur)).Fnext != libc.UintptrFromInt32(0) {
				cur = (*input_report)(unsafe.Pointer(cur)).Fnext
				num_queued++
			}
			(*input_report)(unsafe.Pointer(cur)).Fnext = rpt
			if num_queued > int32(30) {
				return_data(tls, dev, libc.UintptrFromInt32(0), uint64(0))
			}
		}
		hidapi_thread_mutex_unlock(tls, dev+64)
	} else {
		if (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus == int32(LIBUSB_TRANSFER_CANCELLED) {
			(*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread = int32(1)
		} else {
			if (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus == int32(LIBUSB_TRANSFER_NO_DEVICE) {
				(*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread = int32(1)
			} else {
				if (*libusb_transfer)(unsafe.Pointer(transfer)).Fstatus == int32(LIBUSB_TRANSFER_TIMED_OUT) {
				} else {
				}
			}
		}
	}
	if (*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread != 0 {
		(*hid_device)(unsafe.Pointer(dev)).Ftransfer_loop_finished = int32(1)
		return
	}
	res = libusb_submit_transfer(tls, transfer)
	if res != 0 {
		(*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread = int32(1)
		(*hid_device)(unsafe.Pointer(dev)).Ftransfer_loop_finished = int32(1)
	}
}

func read_thread(tls *libc.TLS, param uintptr) (r uintptr) {
	var buf, dev uintptr
	var length size_t
	var res int32
	_, _, _, _ = buf, dev, length, res
	dev = param
	length = libc.Uint64FromInt32((*hid_device)(unsafe.Pointer(dev)).Finput_ep_max_packet_size)
	buf = libc.Xmalloc(tls, length)
	(*hid_device)(unsafe.Pointer(dev)).Ftransfer = libusb_alloc_transfer(tls, 0)
	libusb_fill_interrupt_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Ftransfer, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32((*hid_device)(unsafe.Pointer(dev)).Finput_endpoint), buf, libc.Int32FromUint64(length), __ccgo_fp(read_callback), dev, uint32(5000))
	res = libusb_submit_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Ftransfer)
	if res < 0 {
		(*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread = int32(1)
		(*hid_device)(unsafe.Pointer(dev)).Ftransfer_loop_finished = int32(1)
	}
	hidapi_thread_barrier_wait(tls, dev+64)
	for !((*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread != 0) {
		res = libusb_handle_events(tls, usb_context)
		if res < 0 {
			if res != int32(LIBUSB_ERROR_BUSY) && res != int32(LIBUSB_ERROR_TIMEOUT) && res != int32(LIBUSB_ERROR_OVERFLOW) && res != int32(LIBUSB_ERROR_INTERRUPTED) {
				(*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread = int32(1)
				break
			}
		}
	}
	libusb_cancel_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Ftransfer)
	for !((*hid_device)(unsafe.Pointer(dev)).Ftransfer_loop_finished != 0) {
		libusb_handle_events_completed(tls, usb_context, dev+196)
	}
	hidapi_thread_mutex_lock(tls, dev+64)
	hidapi_thread_cond_broadcast(tls, dev+64)
	hidapi_thread_mutex_unlock(tls, dev+64)
	return libc.UintptrFromInt32(0)
}

func init_xbox360(tls *libc.TLS, device_handle uintptr, idVendor uint16, idProduct uint16, conf_desc uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _ /* data at bp+0 */ [20]uint8
	_ = conf_desc
	if libc.Int32FromUint16(idVendor) == int32(0x05ac) && libc.Int32FromUint16(idProduct) == int32(0x055b) || libc.Int32FromUint16(idVendor) == int32(0x0f0d) {
		libc.Xmemset(tls, bp, 0, uint64(20))
		libusb_control_transfer(tls, device_handle, uint8(0xC1), uint8(0x01), uint16(0x100), uint16(0x0), bp, uint16(20), uint32(100))
	}
}

func init_xboxone(tls *libc.TLS, device_handle uintptr, idVendor uint16, idProduct uint16, conf_desc uintptr) {
	var bSetAlternateSetting, j, k, res int32
	var intf, intf_desc uintptr
	_, _, _, _, _, _ = bSetAlternateSetting, intf, intf_desc, j, k, res
	_ = idProduct
	j = 0
	for {
		if !(j < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(conf_desc)).FbNumInterfaces)) {
			break
		}
		intf = (*libusb_config_descriptor)(unsafe.Pointer(conf_desc)).Finterface1 + uintptr(j)*16
		k = 0
		for {
			if !(k < (*libusb_interface)(unsafe.Pointer(intf)).Fnum_altsetting) {
				break
			}
			intf_desc = (*libusb_interface)(unsafe.Pointer(intf)).Faltsetting + uintptr(k)*40
			if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceClass) == int32(LIBUSB_CLASS_VENDOR_SPEC) && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceSubClass) == xb1_iface_subclass1 && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceProtocol) == xb1_iface_protocol1 {
				bSetAlternateSetting = 0
				if libc.Int32FromUint16(idVendor) == vendor_microsoft && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber) == 0 && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbAlternateSetting) == int32(1) {
					bSetAlternateSetting = int32(1)
				} else {
					if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber) != 0 && libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbAlternateSetting) == 0 {
						bSetAlternateSetting = int32(1)
					}
				}
				if bSetAlternateSetting != 0 {
					res = libusb_claim_interface(tls, device_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
					if res < 0 {
						goto _2
					}
					res = libusb_set_interface_alt_setting(tls, device_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber), libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbAlternateSetting))
					if res < 0 {
					}
					libusb_release_interface(tls, device_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
				}
			}
			goto _2
		_2:
			;
			k++
		}
		goto _1
	_1:
		;
		j++
	}
}

var vendor_microsoft = int32(0x045e)

var xb1_iface_subclass1 = int32(71)

var xb1_iface_protocol1 = int32(208)

func hidapi_initialize_device(tls *libc.TLS, dev uintptr, intf_desc uintptr, conf_desc uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ep uintptr
	var i, is_input, is_interrupt, is_output, res int32
	var _ /* desc at bp+0 */ libusb_device_descriptor
	_, _, _, _, _, _ = ep, i, is_input, is_interrupt, is_output, res
	i = 0
	res = 0
	libusb_get_device_descriptor(tls, libusb_get_device(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle), bp)
	(*hid_device)(unsafe.Pointer(dev)).Fis_driver_detached = 0
	if libusb_kernel_driver_active(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber)) == int32(1) {
		res = libusb_detach_kernel_driver(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
		if res < 0 {
			return 0
		} else {
			(*hid_device)(unsafe.Pointer(dev)).Fis_driver_detached = int32(1)
		}
	}
	res = libusb_claim_interface(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
	if res < 0 {
		if (*hid_device)(unsafe.Pointer(dev)).Fis_driver_detached != 0 {
			res = libusb_attach_kernel_driver(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
			if res < 0 {
			}
		}
		return 0
	}
	if is_xbox360(tls, (*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FidVendor, intf_desc) != 0 {
		init_xbox360(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, (*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FidVendor, (*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FidProduct, conf_desc)
	}
	if is_xboxone(tls, (*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FidVendor, intf_desc) != 0 {
		init_xboxone(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, (*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FidVendor, (*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FidProduct, conf_desc)
	}
	(*hid_device)(unsafe.Pointer(dev)).Fmanufacturer_index = libc.Int32FromUint8((*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FiManufacturer)
	(*hid_device)(unsafe.Pointer(dev)).Fproduct_index = libc.Int32FromUint8((*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FiProduct)
	(*hid_device)(unsafe.Pointer(dev)).Fserial_index = libc.Int32FromUint8((*(*libusb_device_descriptor)(unsafe.Pointer(bp))).FiSerialNumber)
	(*hid_device)(unsafe.Pointer(dev)).Fconfig_number = libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(conf_desc)).FbConfigurationValue)
	(*hid_device)(unsafe.Pointer(dev)).Finterface1 = libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber)
	(*hid_device)(unsafe.Pointer(dev)).Freport_descriptor_size = get_report_descriptor_size_from_interface_descriptors(tls, intf_desc)
	(*hid_device)(unsafe.Pointer(dev)).Finput_endpoint = 0
	(*hid_device)(unsafe.Pointer(dev)).Finput_ep_max_packet_size = 0
	(*hid_device)(unsafe.Pointer(dev)).Foutput_endpoint = 0
	i = 0
	for {
		if !(i < libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbNumEndpoints)) {
			break
		}
		ep = (*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).Fendpoint + uintptr(i)*32
		is_interrupt = libc.BoolInt32(libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbmAttributes)&int32(0x03) == int32(LIBUSB_TRANSFER_TYPE_INTERRUPT))
		is_output = libc.BoolInt32(libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbEndpointAddress)&int32(0x80) == int32(LIBUSB_ENDPOINT_OUT))
		is_input = libc.BoolInt32(libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbEndpointAddress)&int32(0x80) == int32(LIBUSB_ENDPOINT_IN))
		if (*hid_device)(unsafe.Pointer(dev)).Finput_endpoint == 0 && is_interrupt != 0 && is_input != 0 {
			(*hid_device)(unsafe.Pointer(dev)).Finput_endpoint = libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbEndpointAddress)
			(*hid_device)(unsafe.Pointer(dev)).Finput_ep_max_packet_size = libc.Int32FromUint16((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FwMaxPacketSize)
		}
		if (*hid_device)(unsafe.Pointer(dev)).Foutput_endpoint == 0 && is_interrupt != 0 && is_output != 0 {
			(*hid_device)(unsafe.Pointer(dev)).Foutput_endpoint = libc.Int32FromUint8((*libusb_endpoint_descriptor)(unsafe.Pointer(ep)).FbEndpointAddress)
		}
		goto _1
	_1:
		;
		i++
	}
	hidapi_thread_create(tls, dev+64, __ccgo_fp(read_thread), dev)
	hidapi_thread_barrier_wait(tls, dev+64)
	return int32(1)
}

func hid_open_path(tls *libc.TLS, path uintptr) (r uintptr) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var d, good_open, j, k, res, v2 int32
	var dev, intf, intf_desc, usb_dev, v1 uintptr
	var _ /* conf_desc at bp+32 */ uintptr
	var _ /* desc at bp+8 */ libusb_device_descriptor
	var _ /* dev_path at bp+40 */ [64]uint8
	var _ /* devs at bp+0 */ uintptr
	_, _, _, _, _, _, _, _, _, _, _ = d, dev, good_open, intf, intf_desc, j, k, res, usb_dev, v1, v2
	dev = libc.UintptrFromInt32(0)
	*(*uintptr)(unsafe.Pointer(bp)) = libc.UintptrFromInt32(0)
	usb_dev = libc.UintptrFromInt32(0)
	res = 0
	d = 0
	good_open = 0
	if hid_init(tls) < 0 {
		return libc.UintptrFromInt32(0)
	}
	dev = new_hid_device(tls)
	if !(dev != 0) {
		return libc.UintptrFromInt32(0)
	}
	libusb_get_device_list(tls, usb_context, bp)
	for {
		v2 = d
		d++
		v1 = *(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)) + uintptr(v2)*8))
		usb_dev = v1
		if !(v1 != libc.UintptrFromInt32(0) && !(good_open != 0)) {
			break
		}
		*(*uintptr)(unsafe.Pointer(bp + 32)) = libc.UintptrFromInt32(0)
		res = libusb_get_device_descriptor(tls, usb_dev, bp+8)
		if res < 0 {
			continue
		}
		res = libusb_get_active_config_descriptor(tls, usb_dev, bp+32)
		if res < 0 {
			libusb_get_config_descriptor(tls, usb_dev, uint8(0), bp+32)
		}
		if !(*(*uintptr)(unsafe.Pointer(bp + 32)) != 0) {
			continue
		}
		j = 0
		for {
			if !(j < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).FbNumInterfaces) && !(good_open != 0)) {
				break
			}
			intf = (*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).Finterface1 + uintptr(j)*16
			k = 0
			for {
				if !(k < (*libusb_interface)(unsafe.Pointer(intf)).Fnum_altsetting && !(good_open != 0)) {
					break
				}
				intf_desc = (*libusb_interface)(unsafe.Pointer(intf)).Faltsetting + uintptr(k)*40
				if should_enumerate_interface(tls, (*(*libusb_device_descriptor)(unsafe.Pointer(bp + 8))).FidVendor, intf_desc) != 0 {
					get_path(tls, bp+40, usb_dev, libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 32)))).FbConfigurationValue), libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber))
					if !(libc.Xstrcmp(tls, bp+40, path) != 0) {
						res = libusb_open(tls, usb_dev, dev)
						if res < 0 {
							break
						}
						good_open = hidapi_initialize_device(tls, dev, intf_desc, *(*uintptr)(unsafe.Pointer(bp + 32)))
						if !(good_open != 0) {
							libusb_close(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle)
						}
					}
				}
				goto _4
			_4:
				;
				k++
			}
			goto _3
		_3:
			;
			j++
		}
		libusb_free_config_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp + 32)))
	}
	libusb_free_device_list(tls, *(*uintptr)(unsafe.Pointer(bp)), int32(1))
	if good_open != 0 {
		return dev
	} else {
		free_hid_device(tls, dev)
		return libc.UintptrFromInt32(0)
	}
	return r
}

func hid_libusb_wrap_sys_device(tls *libc.TLS, sys_dev intptr_t, interface_num int32) (r uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var dev, intf, intf_desc, selected_intf_desc uintptr
	var j, k, res int32
	var _ /* conf_desc at bp+0 */ uintptr
	_, _, _, _, _, _, _ = dev, intf, intf_desc, j, k, res, selected_intf_desc
	dev = libc.UintptrFromInt32(0)
	*(*uintptr)(unsafe.Pointer(bp)) = libc.UintptrFromInt32(0)
	selected_intf_desc = libc.UintptrFromInt32(0)
	res = 0
	j = 0
	k = 0
	if hid_init(tls) < 0 {
		return libc.UintptrFromInt32(0)
	}
	dev = new_hid_device(tls)
	if !(dev != 0) {
		return libc.UintptrFromInt32(0)
	}
	res = libusb_wrap_sys_device(tls, usb_context, sys_dev, dev)
	if res < 0 {
		goto err
	}
	res = libusb_get_active_config_descriptor(tls, libusb_get_device(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle), bp)
	if res < 0 {
		libusb_get_config_descriptor(tls, libusb_get_device(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle), uint8(0), bp)
	}
	if !(*(*uintptr)(unsafe.Pointer(bp)) != 0) {
		goto err
	}
	j = 0
	for {
		if !(j < libc.Int32FromUint8((*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).FbNumInterfaces) && !(selected_intf_desc != 0)) {
			break
		}
		intf = (*libusb_config_descriptor)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp)))).Finterface1 + uintptr(j)*16
		k = 0
		for {
			if !(k < (*libusb_interface)(unsafe.Pointer(intf)).Fnum_altsetting) {
				break
			}
			intf_desc = (*libusb_interface)(unsafe.Pointer(intf)).Faltsetting + uintptr(k)*40
			if libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceClass) == int32(LIBUSB_CLASS_HID) {
				if interface_num < 0 || interface_num == libc.Int32FromUint8((*libusb_interface_descriptor)(unsafe.Pointer(intf_desc)).FbInterfaceNumber) {
					selected_intf_desc = intf_desc
					break
				}
			}
			goto _2
		_2:
			;
			k++
		}
		goto _1
	_1:
		;
		j++
	}
	if !(selected_intf_desc != 0) {
		if interface_num < 0 {
		} else {
		}
		goto err
	}
	if !(hidapi_initialize_device(tls, dev, selected_intf_desc, *(*uintptr)(unsafe.Pointer(bp))) != 0) {
		goto err
	}
	return dev
	goto err
err:
	;
	if *(*uintptr)(unsafe.Pointer(bp)) != 0 {
		libusb_free_config_descriptor(tls, *(*uintptr)(unsafe.Pointer(bp)))
	}
	if (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle != 0 {
		libusb_close(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle)
	}
	free_hid_device(tls, dev)
	return libc.UintptrFromInt32(0)
}

func hid_write(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var report_number, res, skipped_report_id int32
	var _ /* actual_length at bp+0 */ int32
	_, _, _ = report_number, res, skipped_report_id
	skipped_report_id = 0
	if (*hid_device)(unsafe.Pointer(dev)).Foutput_endpoint <= 0 {
		return hid_send_output_report(tls, dev, data, length)
	}
	if !(data != 0) || length == uint64(0) {
		return -int32(1)
	}
	report_number = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(data)))
	if report_number == 0x0 {
		data++
		length--
		skipped_report_id = int32(1)
	}
	res = libusb_interrupt_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32((*hid_device)(unsafe.Pointer(dev)).Foutput_endpoint), data, libc.Int32FromUint64(length), bp, uint32(1000))
	if res < 0 {
		return -int32(1)
	}
	if skipped_report_id != 0 {
		*(*int32)(unsafe.Pointer(bp))++
	}
	return *(*int32)(unsafe.Pointer(bp))
}

func return_data(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	var len1 size_t
	var rpt uintptr
	var v1 uint64
	_, _, _ = len1, rpt, v1
	rpt = (*hid_device)(unsafe.Pointer(dev)).Finput_reports
	if length < (*input_report)(unsafe.Pointer(rpt)).Flen1 {
		v1 = length
	} else {
		v1 = (*input_report)(unsafe.Pointer(rpt)).Flen1
	}
	len1 = v1
	if len1 > uint64(0) {
		libc.Xmemcpy(tls, data, (*input_report)(unsafe.Pointer(rpt)).Fdata, len1)
	}
	(*hid_device)(unsafe.Pointer(dev)).Finput_reports = (*input_report)(unsafe.Pointer(rpt)).Fnext
	libc.Xfree(tls, (*input_report)(unsafe.Pointer(rpt)).Fdata)
	libc.Xfree(tls, rpt)
	return libc.Int32FromUint64(len1)
}

func cleanup_mutex(tls *libc.TLS, param uintptr) {
	var dev uintptr
	_ = dev
	dev = param
	hidapi_thread_mutex_unlock(tls, dev+64)
}

func hid_read_timeout(tls *libc.TLS, dev uintptr, data uintptr, length size_t, milliseconds int32) (r int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var bytes_read, res int32
	var _ /* __cb at bp+0 */ __ptcb
	var _ /* ts at bp+24 */ hidapi_timespec
	_, _ = bytes_read, res
	hidapi_thread_mutex_lock(tls, dev+64)
	libc.X_pthread_cleanup_push(tls, bp, __ccgo_fp(cleanup_mutex), dev)
	bytes_read = -int32(1)
	if (*hid_device)(unsafe.Pointer(dev)).Finput_reports != 0 {
		bytes_read = return_data(tls, dev, data, length)
		goto ret
	}
	if (*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread != 0 {
		bytes_read = -int32(1)
		goto ret
	}
	if milliseconds == -int32(1) {
		for !((*hid_device)(unsafe.Pointer(dev)).Finput_reports != 0) && !((*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread != 0) {
			hidapi_thread_cond_wait(tls, dev+64)
		}
		if (*hid_device)(unsafe.Pointer(dev)).Finput_reports != 0 {
			bytes_read = return_data(tls, dev, data, length)
		}
	} else {
		if milliseconds > 0 {
			hidapi_thread_gettime(tls, bp+24)
			hidapi_thread_addtime(tls, bp+24, milliseconds)
			for !((*hid_device)(unsafe.Pointer(dev)).Finput_reports != 0) && !((*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread != 0) {
				res = hidapi_thread_cond_timedwait(tls, dev+64, bp+24)
				if res == 0 {
					if (*hid_device)(unsafe.Pointer(dev)).Finput_reports != 0 {
						bytes_read = return_data(tls, dev, data, length)
						break
					}
				} else {
					if res == int32(110) {
						bytes_read = 0
						break
					} else {
						bytes_read = -int32(1)
						break
					}
				}
			}
		} else {
			bytes_read = 0
		}
	}
	goto ret
ret:
	;
	hidapi_thread_mutex_unlock(tls, dev+64)
	libc.X_pthread_cleanup_pop(tls, bp, 0)
	return bytes_read
}

func hid_read(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	var v1 int32
	_ = v1
	if (*hid_device)(unsafe.Pointer(dev)).Fblocking != 0 {
		v1 = -int32(1)
	} else {
		v1 = 0
	}
	return hid_read_timeout(tls, dev, data, length, v1)
}

func hid_read_error(tls *libc.TLS, dev uintptr) (r uintptr) {
	_ = dev
	return __ccgo_ts + 14155
}

func hid_set_nonblocking(tls *libc.TLS, dev uintptr, nonblock int32) (r int32) {
	(*hid_device)(unsafe.Pointer(dev)).Fblocking = libc.BoolInt32(!(nonblock != 0))
	return 0
}

func hid_send_feature_report(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	var report_number, res, skipped_report_id int32
	_, _, _ = report_number, res, skipped_report_id
	res = -int32(1)
	skipped_report_id = 0
	report_number = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(data)))
	if report_number == 0x0 {
		data++
		length--
		skipped_report_id = int32(1)
	}
	res = libusb_control_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32(int32(LIBUSB_REQUEST_TYPE_CLASS)|int32(LIBUSB_RECIPIENT_INTERFACE)|int32(LIBUSB_ENDPOINT_OUT)), uint8(0x09), libc.Uint16FromInt32(libc.Int32FromInt32(3)<<libc.Int32FromInt32(8)|report_number), libc.Uint16FromInt32((*hid_device)(unsafe.Pointer(dev)).Finterface1), data, uint16(length), uint32(1000))
	if res < 0 {
		return -int32(1)
	}
	if skipped_report_id != 0 {
		length++
	}
	return libc.Int32FromUint64(length)
}

func hid_get_feature_report(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	var report_number, res, skipped_report_id int32
	_, _, _ = report_number, res, skipped_report_id
	res = -int32(1)
	skipped_report_id = 0
	report_number = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(data)))
	if report_number == 0x0 {
		data++
		length--
		skipped_report_id = int32(1)
	}
	res = libusb_control_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32(int32(LIBUSB_REQUEST_TYPE_CLASS)|int32(LIBUSB_RECIPIENT_INTERFACE)|int32(LIBUSB_ENDPOINT_IN)), uint8(0x01), libc.Uint16FromInt32(libc.Int32FromInt32(3)<<libc.Int32FromInt32(8)|report_number), libc.Uint16FromInt32((*hid_device)(unsafe.Pointer(dev)).Finterface1), data, uint16(length), uint32(1000))
	if res < 0 {
		return -int32(1)
	}
	if skipped_report_id != 0 {
		res++
	}
	return res
}

func hid_send_output_report(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	var report_number, res, skipped_report_id int32
	_, _, _ = report_number, res, skipped_report_id
	res = -int32(1)
	skipped_report_id = 0
	report_number = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(data)))
	if report_number == 0x0 {
		data++
		length--
		skipped_report_id = int32(1)
	}
	res = libusb_control_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32(int32(LIBUSB_REQUEST_TYPE_CLASS)|int32(LIBUSB_RECIPIENT_INTERFACE)|int32(LIBUSB_ENDPOINT_OUT)), uint8(0x09), libc.Uint16FromInt32(libc.Int32FromInt32(2)<<libc.Int32FromInt32(8)|report_number), libc.Uint16FromInt32((*hid_device)(unsafe.Pointer(dev)).Finterface1), data, uint16(length), uint32(1000))
	if res < 0 {
		return -int32(1)
	}
	if skipped_report_id != 0 {
		length++
	}
	return libc.Int32FromUint64(length)
}

func hid_get_input_report(tls *libc.TLS, dev uintptr, data uintptr, length size_t) (r int32) {
	var report_number, res, skipped_report_id int32
	_, _, _ = report_number, res, skipped_report_id
	res = -int32(1)
	skipped_report_id = 0
	report_number = libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(data)))
	if report_number == 0x0 {
		data++
		length--
		skipped_report_id = int32(1)
	}
	res = libusb_control_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32(int32(LIBUSB_REQUEST_TYPE_CLASS)|int32(LIBUSB_RECIPIENT_INTERFACE)|int32(LIBUSB_ENDPOINT_IN)), uint8(0x01), libc.Uint16FromInt32(libc.Int32FromInt32(1)<<libc.Int32FromInt32(8)|report_number), libc.Uint16FromInt32((*hid_device)(unsafe.Pointer(dev)).Finterface1), data, uint16(length), uint32(1000))
	if res < 0 {
		return -int32(1)
	}
	if skipped_report_id != 0 {
		res++
	}
	return res
}

func hid_close(tls *libc.TLS, dev uintptr) {
	var res int32
	_ = res
	if !(dev != 0) {
		return
	}
	(*hid_device)(unsafe.Pointer(dev)).Fshutdown_thread = int32(1)
	libusb_cancel_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Ftransfer)
	hidapi_thread_join(tls, dev+64)
	libc.Xfree(tls, (*libusb_transfer)(unsafe.Pointer((*hid_device)(unsafe.Pointer(dev)).Ftransfer)).Fbuffer)
	(*libusb_transfer)(unsafe.Pointer((*hid_device)(unsafe.Pointer(dev)).Ftransfer)).Fbuffer = libc.UintptrFromInt32(0)
	libusb_free_transfer(tls, (*hid_device)(unsafe.Pointer(dev)).Ftransfer)
	libusb_release_interface(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, (*hid_device)(unsafe.Pointer(dev)).Finterface1)
	if (*hid_device)(unsafe.Pointer(dev)).Fis_driver_detached != 0 {
		res = libusb_attach_kernel_driver(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, (*hid_device)(unsafe.Pointer(dev)).Finterface1)
		if res < 0 {
		}
	}
	libusb_close(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle)
	hidapi_thread_mutex_lock(tls, dev+64)
	for (*hid_device)(unsafe.Pointer(dev)).Finput_reports != 0 {
		return_data(tls, dev, libc.UintptrFromInt32(0), uint64(0))
	}
	hidapi_thread_mutex_unlock(tls, dev+64)
	free_hid_device(tls, dev)
}

func hid_get_manufacturer_string(tls *libc.TLS, dev uintptr, string1 uintptr, maxlen size_t) (r int32) {
	return hid_get_indexed_string(tls, dev, (*hid_device)(unsafe.Pointer(dev)).Fmanufacturer_index, string1, maxlen)
}

func hid_get_product_string(tls *libc.TLS, dev uintptr, string1 uintptr, maxlen size_t) (r int32) {
	return hid_get_indexed_string(tls, dev, (*hid_device)(unsafe.Pointer(dev)).Fproduct_index, string1, maxlen)
}

func hid_get_serial_number_string(tls *libc.TLS, dev uintptr, string1 uintptr, maxlen size_t) (r int32) {
	return hid_get_indexed_string(tls, dev, (*hid_device)(unsafe.Pointer(dev)).Fserial_index, string1, maxlen)
}

func hid_get_device_info(tls *libc.TLS, dev uintptr) (r uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var usb_device uintptr
	var _ /* desc at bp+0 */ libusb_device_descriptor
	_ = usb_device
	if !((*hid_device)(unsafe.Pointer(dev)).Fdevice_info != 0) {
		usb_device = libusb_get_device(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle)
		libusb_get_device_descriptor(tls, usb_device, bp)
		(*hid_device)(unsafe.Pointer(dev)).Fdevice_info = create_device_info_for_device(tls, usb_device, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, bp, (*hid_device)(unsafe.Pointer(dev)).Fconfig_number, (*hid_device)(unsafe.Pointer(dev)).Finterface1)
		if (*hid_device)(unsafe.Pointer(dev)).Fdevice_info != 0 {
			fill_device_info_usage(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_info, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, (*hid_device)(unsafe.Pointer(dev)).Finterface1, (*hid_device)(unsafe.Pointer(dev)).Freport_descriptor_size)
		}
	}
	return (*hid_device)(unsafe.Pointer(dev)).Fdevice_info
}

func hid_get_indexed_string(tls *libc.TLS, dev uintptr, string_index int32, string1 uintptr, maxlen size_t) (r int32) {
	var str uintptr
	_ = str
	str = get_usb_string(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, libc.Uint8FromInt32(string_index))
	if str != 0 {
		libc.Xwcsncpy(tls, string1, str, maxlen)
		*(*wchar_t)(unsafe.Pointer(string1 + uintptr(maxlen-uint64(1))*4)) = int32('\000')
		libc.Xfree(tls, str)
		return 0
	} else {
		return -int32(1)
	}
	return r
}

func hid_get_report_descriptor(tls *libc.TLS, dev uintptr, buf uintptr, buf_size size_t) (r int32) {
	return hid_get_report_descriptor_libusb(tls, (*hid_device)(unsafe.Pointer(dev)).Fdevice_handle, (*hid_device)(unsafe.Pointer(dev)).Finterface1, (*hid_device)(unsafe.Pointer(dev)).Freport_descriptor_size, buf, buf_size)
}

func hid_error(tls *libc.TLS, dev uintptr) (r uintptr) {
	_ = dev
	return __ccgo_ts + 14307
}

type lang_map_entry = struct {
	Fname        uintptr
	Fstring_code uintptr
	Fusb_code    uint16_t
}

var lang_map = [134]lang_map_entry{
	0: {
		Fname:        __ccgo_ts + 14439,
		Fstring_code: __ccgo_ts + 14449,
		Fusb_code:    uint16(0x0436),
	},
	1: {
		Fname:        __ccgo_ts + 14452,
		Fstring_code: __ccgo_ts + 14461,
		Fusb_code:    uint16(0x041C),
	},
	2: {
		Fname:        __ccgo_ts + 14464,
		Fstring_code: __ccgo_ts + 14494,
		Fusb_code:    uint16(0x3801),
	},
	3: {
		Fname:        __ccgo_ts + 14500,
		Fstring_code: __ccgo_ts + 14517,
		Fusb_code:    uint16(0x3C01),
	},
	4: {
		Fname:        __ccgo_ts + 14523,
		Fstring_code: __ccgo_ts + 14540,
		Fusb_code:    uint16(0x1401),
	},
	5: {
		Fname:        __ccgo_ts + 14546,
		Fstring_code: __ccgo_ts + 14561,
		Fusb_code:    uint16(0x0C01),
	},
	6: {
		Fname:        __ccgo_ts + 14567,
		Fstring_code: __ccgo_ts + 14581,
		Fusb_code:    uint16(0x0801),
	},
	7: {
		Fname:        __ccgo_ts + 14587,
		Fstring_code: __ccgo_ts + 14603,
		Fusb_code:    uint16(0x2C01),
	},
	8: {
		Fname:        __ccgo_ts + 14609,
		Fstring_code: __ccgo_ts + 14625,
		Fusb_code:    uint16(0x3401),
	},
	9: {
		Fname:        __ccgo_ts + 14631,
		Fstring_code: __ccgo_ts + 14648,
		Fusb_code:    uint16(0x3001),
	},
	10: {
		Fname:        __ccgo_ts + 14654,
		Fstring_code: __ccgo_ts + 14669,
		Fusb_code:    uint16(0x1001),
	},
	11: {
		Fname:        __ccgo_ts + 14675,
		Fstring_code: __ccgo_ts + 14692,
		Fusb_code:    uint16(0x1801),
	},
	12: {
		Fname:        __ccgo_ts + 14698,
		Fstring_code: __ccgo_ts + 14712,
		Fusb_code:    uint16(0x2001),
	},
	13: {
		Fname:        __ccgo_ts + 14718,
		Fstring_code: __ccgo_ts + 14733,
		Fusb_code:    uint16(0x4001),
	},
	14: {
		Fname:        __ccgo_ts + 14739,
		Fstring_code: __ccgo_ts + 14761,
		Fusb_code:    uint16(0x0401),
	},
	15: {
		Fname:        __ccgo_ts + 14767,
		Fstring_code: __ccgo_ts + 14782,
		Fusb_code:    uint16(0x2801),
	},
	16: {
		Fname:        __ccgo_ts + 14788,
		Fstring_code: __ccgo_ts + 14805,
		Fusb_code:    uint16(0x1C01),
	},
	17: {
		Fname:        __ccgo_ts + 14811,
		Fstring_code: __ccgo_ts + 14826,
		Fusb_code:    uint16(0x2401),
	},
	18: {
		Fname:        __ccgo_ts + 14832,
		Fstring_code: __ccgo_ts + 14841,
		Fusb_code:    uint16(0x042B),
	},
	19: {
		Fname:        __ccgo_ts + 14844,
		Fstring_code: __ccgo_ts + 14858,
		Fusb_code:    uint16(0x042C),
	},
	20: {
		Fname:        __ccgo_ts + 14864,
		Fstring_code: __ccgo_ts + 14858,
		Fusb_code:    uint16(0x082C),
	},
	21: {
		Fname:        __ccgo_ts + 14881,
		Fstring_code: __ccgo_ts + 14888,
		Fusb_code:    uint16(0x042D),
	},
	22: {
		Fname:        __ccgo_ts + 14891,
		Fstring_code: __ccgo_ts + 14902,
		Fusb_code:    uint16(0x0423),
	},
	23: {
		Fname:        __ccgo_ts + 14905,
		Fstring_code: __ccgo_ts + 14915,
		Fusb_code:    uint16(0x0402),
	},
	24: {
		Fname:        __ccgo_ts + 14918,
		Fstring_code: __ccgo_ts + 14926,
		Fusb_code:    uint16(0x0403),
	},
	25: {
		Fname:        __ccgo_ts + 14929,
		Fstring_code: __ccgo_ts + 14945,
		Fusb_code:    uint16(0x0804),
	},
	26: {
		Fname:        __ccgo_ts + 14951,
		Fstring_code: __ccgo_ts + 14975,
		Fusb_code:    uint16(0x0C04),
	},
	27: {
		Fname:        __ccgo_ts + 14981,
		Fstring_code: __ccgo_ts + 15001,
		Fusb_code:    uint16(0x1404),
	},
	28: {
		Fname:        __ccgo_ts + 15007,
		Fstring_code: __ccgo_ts + 15027,
		Fusb_code:    uint16(0x1004),
	},
	29: {
		Fname:        __ccgo_ts + 15033,
		Fstring_code: __ccgo_ts + 15050,
		Fusb_code:    uint16(0x0404),
	},
	30: {
		Fname:        __ccgo_ts + 15056,
		Fstring_code: __ccgo_ts + 15065,
		Fusb_code:    uint16(0x041A),
	},
	31: {
		Fname:        __ccgo_ts + 15068,
		Fstring_code: __ccgo_ts + 15074,
		Fusb_code:    uint16(0x0405),
	},
	32: {
		Fname:        __ccgo_ts + 15077,
		Fstring_code: __ccgo_ts + 15084,
		Fusb_code:    uint16(0x0406),
	},
	33: {
		Fname:        __ccgo_ts + 15087,
		Fstring_code: __ccgo_ts + 15107,
		Fusb_code:    uint16(0x0413),
	},
	34: {
		Fname:        __ccgo_ts + 15113,
		Fstring_code: __ccgo_ts + 15129,
		Fusb_code:    uint16(0x0813),
	},
	35: {
		Fname:        __ccgo_ts + 15135,
		Fstring_code: __ccgo_ts + 15155,
		Fusb_code:    uint16(0x0C09),
	},
	36: {
		Fname:        __ccgo_ts + 15161,
		Fstring_code: __ccgo_ts + 15178,
		Fusb_code:    uint16(0x2809),
	},
	37: {
		Fname:        __ccgo_ts + 15184,
		Fstring_code: __ccgo_ts + 15201,
		Fusb_code:    uint16(0x1009),
	},
	38: {
		Fname:        __ccgo_ts + 15207,
		Fstring_code: __ccgo_ts + 15227,
		Fusb_code:    uint16(0x2409),
	},
	39: {
		Fname:        __ccgo_ts + 15233,
		Fstring_code: __ccgo_ts + 15251,
		Fusb_code:    uint16(0x1809),
	},
	40: {
		Fname:        __ccgo_ts + 15257,
		Fstring_code: __ccgo_ts + 15275,
		Fusb_code:    uint16(0x2009),
	},
	41: {
		Fname:        __ccgo_ts + 15281,
		Fstring_code: __ccgo_ts + 15303,
		Fusb_code:    uint16(0x1409),
	},
	42: {
		Fname:        __ccgo_ts + 15309,
		Fstring_code: __ccgo_ts + 15331,
		Fusb_code:    uint16(0x3409),
	},
	43: {
		Fname:        __ccgo_ts + 15337,
		Fstring_code: __ccgo_ts + 15363,
		Fusb_code:    uint16(0x1C09),
	},
	44: {
		Fname:        __ccgo_ts + 15369,
		Fstring_code: __ccgo_ts + 15388,
		Fusb_code:    uint16(0x2C09),
	},
	45: {
		Fname:        __ccgo_ts + 15394,
		Fstring_code: __ccgo_ts + 15418,
		Fusb_code:    uint16(0x0809),
	},
	46: {
		Fname:        __ccgo_ts + 15424,
		Fstring_code: __ccgo_ts + 15448,
		Fusb_code:    uint16(0x0409),
	},
	47: {
		Fname:        __ccgo_ts + 15454,
		Fstring_code: __ccgo_ts + 15463,
		Fusb_code:    uint16(0x0425),
	},
	48: {
		Fname:        __ccgo_ts + 15466,
		Fstring_code: __ccgo_ts + 15472,
		Fusb_code:    uint16(0x0429),
	},
	49: {
		Fname:        __ccgo_ts + 15475,
		Fstring_code: __ccgo_ts + 15483,
		Fusb_code:    uint16(0x040B),
	},
	50: {
		Fname:        __ccgo_ts + 15486,
		Fstring_code: __ccgo_ts + 15494,
		Fusb_code:    uint16(0x0438),
	},
	51: {
		Fname:        __ccgo_ts + 15497,
		Fstring_code: __ccgo_ts + 15513,
		Fusb_code:    uint16(0x040C),
	},
	52: {
		Fname:        __ccgo_ts + 15519,
		Fstring_code: __ccgo_ts + 15536,
		Fusb_code:    uint16(0x080C),
	},
	53: {
		Fname:        __ccgo_ts + 15542,
		Fstring_code: __ccgo_ts + 15558,
		Fusb_code:    uint16(0x0C0C),
	},
	54: {
		Fname:        __ccgo_ts + 15564,
		Fstring_code: __ccgo_ts + 15584,
		Fusb_code:    uint16(0x140C),
	},
	55: {
		Fname:        __ccgo_ts + 15590,
		Fstring_code: __ccgo_ts + 15611,
		Fusb_code:    uint16(0x100C),
	},
	56: {
		Fname:        __ccgo_ts + 15617,
		Fstring_code: __ccgo_ts + 15634,
		Fusb_code:    uint16(0x083C),
	},
	57: {
		Fname:        __ccgo_ts + 15640,
		Fstring_code: __ccgo_ts + 15658,
		Fusb_code:    uint16(0x043C),
	},
	58: {
		Fname:        __ccgo_ts + 15661,
		Fstring_code: __ccgo_ts + 15678,
		Fusb_code:    uint16(0x0407),
	},
	59: {
		Fname:        __ccgo_ts + 15684,
		Fstring_code: __ccgo_ts + 15701,
		Fusb_code:    uint16(0x0C07),
	},
	60: {
		Fname:        __ccgo_ts + 15707,
		Fstring_code: __ccgo_ts + 15730,
		Fusb_code:    uint16(0x1407),
	},
	61: {
		Fname:        __ccgo_ts + 15736,
		Fstring_code: __ccgo_ts + 15756,
		Fusb_code:    uint16(0x1007),
	},
	62: {
		Fname:        __ccgo_ts + 15762,
		Fstring_code: __ccgo_ts + 15783,
		Fusb_code:    uint16(0x0807),
	},
	63: {
		Fname:        __ccgo_ts + 15789,
		Fstring_code: __ccgo_ts + 15795,
		Fusb_code:    uint16(0x0408),
	},
	64: {
		Fname:        __ccgo_ts + 15798,
		Fstring_code: __ccgo_ts + 15805,
		Fusb_code:    uint16(0x040D),
	},
	65: {
		Fname:        __ccgo_ts + 15808,
		Fstring_code: __ccgo_ts + 15814,
		Fusb_code:    uint16(0x0439),
	},
	66: {
		Fname:        __ccgo_ts + 15817,
		Fstring_code: __ccgo_ts + 11346,
		Fusb_code:    uint16(0x040E),
	},
	67: {
		Fname:        __ccgo_ts + 15827,
		Fstring_code: __ccgo_ts + 15837,
		Fusb_code:    uint16(0x040F),
	},
	68: {
		Fname:        __ccgo_ts + 15840,
		Fstring_code: __ccgo_ts + 15851,
		Fusb_code:    uint16(0x0421),
	},
	69: {
		Fname:        __ccgo_ts + 15854,
		Fstring_code: __ccgo_ts + 15870,
		Fusb_code:    uint16(0x0410),
	},
	70: {
		Fname:        __ccgo_ts + 15876,
		Fstring_code: __ccgo_ts + 15898,
		Fusb_code:    uint16(0x0810),
	},
	71: {
		Fname:        __ccgo_ts + 15904,
		Fstring_code: __ccgo_ts + 15913,
		Fusb_code:    uint16(0x0411),
	},
	72: {
		Fname:        __ccgo_ts + 15916,
		Fstring_code: __ccgo_ts + 15923,
		Fusb_code:    uint16(0x0412),
	},
	73: {
		Fname:        __ccgo_ts + 15926,
		Fstring_code: __ccgo_ts + 15934,
		Fusb_code:    uint16(0x0426),
	},
	74: {
		Fname:        __ccgo_ts + 15937,
		Fstring_code: __ccgo_ts + 15948,
		Fusb_code:    uint16(0x0427),
	},
	75: {
		Fname:        __ccgo_ts + 15951,
		Fstring_code: __ccgo_ts + 15970,
		Fusb_code:    uint16(0x042F),
	},
	76: {
		Fname:        __ccgo_ts + 15973,
		Fstring_code: __ccgo_ts + 15990,
		Fusb_code:    uint16(0x043E),
	},
	77: {
		Fname:        __ccgo_ts + 15996,
		Fstring_code: __ccgo_ts + 16013,
		Fusb_code:    uint16(0x083E),
	},
	78: {
		Fname:        __ccgo_ts + 16019,
		Fstring_code: __ccgo_ts + 16027,
		Fusb_code:    uint16(0x043A),
	},
	79: {
		Fname:        __ccgo_ts + 16030,
		Fstring_code: __ccgo_ts + 16038,
		Fusb_code:    uint16(0x044E),
	},
	80: {
		Fname:        __ccgo_ts + 16041,
		Fstring_code: __ccgo_ts + 16059,
		Fusb_code:    uint16(0x0414),
	},
	81: {
		Fname:        __ccgo_ts + 16065,
		Fstring_code: __ccgo_ts + 16059,
		Fusb_code:    uint16(0x0814),
	},
	82: {
		Fname:        __ccgo_ts + 16085,
		Fstring_code: __ccgo_ts + 16092,
		Fusb_code:    uint16(0x0415),
	},
	83: {
		Fname:        __ccgo_ts + 16095,
		Fstring_code: __ccgo_ts + 16117,
		Fusb_code:    uint16(0x0816),
	},
	84: {
		Fname:        __ccgo_ts + 16123,
		Fstring_code: __ccgo_ts + 16143,
		Fusb_code:    uint16(0x0416),
	},
	85: {
		Fname:        __ccgo_ts + 16149,
		Fstring_code: __ccgo_ts + 16163,
		Fusb_code:    uint16(0x0417),
	},
	86: {
		Fname:        __ccgo_ts + 16166,
		Fstring_code: __ccgo_ts + 16185,
		Fusb_code:    uint16(0x0418),
	},
	87: {
		Fname:        __ccgo_ts + 16188,
		Fstring_code: __ccgo_ts + 16219,
		Fusb_code:    uint16(0x0818),
	},
	88: {
		Fname:        __ccgo_ts + 16225,
		Fstring_code: __ccgo_ts + 11340,
		Fusb_code:    uint16(0x0419),
	},
	89: {
		Fname:        __ccgo_ts + 16233,
		Fstring_code: __ccgo_ts + 16263,
		Fusb_code:    uint16(0x0819),
	},
	90: {
		Fname:        __ccgo_ts + 16269,
		Fstring_code: __ccgo_ts + 16278,
		Fusb_code:    uint16(0x044F),
	},
	91: {
		Fname:        __ccgo_ts + 16281,
		Fstring_code: __ccgo_ts + 16300,
		Fusb_code:    uint16(0x0C1A),
	},
	92: {
		Fname:        __ccgo_ts + 16306,
		Fstring_code: __ccgo_ts + 16300,
		Fusb_code:    uint16(0x081A),
	},
	93: {
		Fname:        __ccgo_ts + 16322,
		Fstring_code: __ccgo_ts + 16331,
		Fusb_code:    uint16(0x0432),
	},
	94: {
		Fname:        __ccgo_ts + 16334,
		Fstring_code: __ccgo_ts + 16344,
		Fusb_code:    uint16(0x0424),
	},
	95: {
		Fname:        __ccgo_ts + 16347,
		Fstring_code: __ccgo_ts + 16354,
		Fusb_code:    uint16(0x041B),
	},
	96: {
		Fname:        __ccgo_ts + 16357,
		Fstring_code: __ccgo_ts + 16365,
		Fusb_code:    uint16(0x042E),
	},
	97: {
		Fname:        __ccgo_ts + 16368,
		Fstring_code: __ccgo_ts + 16398,
		Fusb_code:    uint16(0x040A),
	},
	98: {
		Fname:        __ccgo_ts + 16404,
		Fstring_code: __ccgo_ts + 16424,
		Fusb_code:    uint16(0x2C0A),
	},
	99: {
		Fname:        __ccgo_ts + 16430,
		Fstring_code: __ccgo_ts + 16448,
		Fusb_code:    uint16(0x400A),
	},
	100: {
		Fname:        __ccgo_ts + 16454,
		Fstring_code: __ccgo_ts + 16470,
		Fusb_code:    uint16(0x340A),
	},
	101: {
		Fname:        __ccgo_ts + 16476,
		Fstring_code: __ccgo_ts + 16495,
		Fusb_code:    uint16(0x240A),
	},
	102: {
		Fname:        __ccgo_ts + 16501,
		Fstring_code: __ccgo_ts + 16522,
		Fusb_code:    uint16(0x140A),
	},
	103: {
		Fname:        __ccgo_ts + 16528,
		Fstring_code: __ccgo_ts + 16557,
		Fusb_code:    uint16(0x1C0A),
	},
	104: {
		Fname:        __ccgo_ts + 16563,
		Fstring_code: __ccgo_ts + 16581,
		Fusb_code:    uint16(0x300A),
	},
	105: {
		Fname:        __ccgo_ts + 16587,
		Fstring_code: __ccgo_ts + 16607,
		Fusb_code:    uint16(0x100A),
	},
	106: {
		Fname:        __ccgo_ts + 16613,
		Fstring_code: __ccgo_ts + 16632,
		Fusb_code:    uint16(0x480A),
	},
	107: {
		Fname:        __ccgo_ts + 16638,
		Fstring_code: __ccgo_ts + 16655,
		Fusb_code:    uint16(0x080A),
	},
	108: {
		Fname:        __ccgo_ts + 16661,
		Fstring_code: __ccgo_ts + 16681,
		Fusb_code:    uint16(0x4C0A),
	},
	109: {
		Fname:        __ccgo_ts + 16687,
		Fstring_code: __ccgo_ts + 16704,
		Fusb_code:    uint16(0x180A),
	},
	110: {
		Fname:        __ccgo_ts + 16710,
		Fstring_code: __ccgo_ts + 16725,
		Fusb_code:    uint16(0x280A),
	},
	111: {
		Fname:        __ccgo_ts + 16731,
		Fstring_code: __ccgo_ts + 16753,
		Fusb_code:    uint16(0x500A),
	},
	112: {
		Fname:        __ccgo_ts + 16759,
		Fstring_code: __ccgo_ts + 16778,
		Fusb_code:    uint16(0x3C0A),
	},
	113: {
		Fname:        __ccgo_ts + 16784,
		Fstring_code: __ccgo_ts + 16806,
		Fusb_code:    uint16(0x440A),
	},
	114: {
		Fname:        __ccgo_ts + 16812,
		Fstring_code: __ccgo_ts + 16830,
		Fusb_code:    uint16(0x380A),
	},
	115: {
		Fname:        __ccgo_ts + 16836,
		Fstring_code: __ccgo_ts + 16856,
		Fusb_code:    uint16(0x200A),
	},
	116: {
		Fname:        __ccgo_ts + 16862,
		Fstring_code: __ccgo_ts + 16877,
		Fusb_code:    uint16(0x0430),
	},
	117: {
		Fname:        __ccgo_ts + 16880,
		Fstring_code: __ccgo_ts + 16888,
		Fusb_code:    uint16(0x0441),
	},
	118: {
		Fname:        __ccgo_ts + 16891,
		Fstring_code: __ccgo_ts + 16908,
		Fusb_code:    uint16(0x041D),
	},
	119: {
		Fname:        __ccgo_ts + 16914,
		Fstring_code: __ccgo_ts + 16932,
		Fusb_code:    uint16(0x081D),
	},
	120: {
		Fname:        __ccgo_ts + 16938,
		Fstring_code: __ccgo_ts + 16944,
		Fusb_code:    uint16(0x0449),
	},
	121: {
		Fname:        __ccgo_ts + 16947,
		Fstring_code: __ccgo_ts + 16953,
		Fusb_code:    uint16(0x0444),
	},
	122: {
		Fname:        __ccgo_ts + 16956,
		Fstring_code: __ccgo_ts + 16961,
		Fusb_code:    uint16(0x041E),
	},
	123: {
		Fname:        __ccgo_ts + 16964,
		Fstring_code: __ccgo_ts + 16972,
		Fusb_code:    uint16(0x041F),
	},
	124: {
		Fname:        __ccgo_ts + 16975,
		Fstring_code: __ccgo_ts + 16982,
		Fusb_code:    uint16(0x0431),
	},
	125: {
		Fname:        __ccgo_ts + 16985,
		Fstring_code: __ccgo_ts + 16995,
		Fusb_code:    uint16(0x0422),
	},
	126: {
		Fname:        __ccgo_ts + 16998,
		Fstring_code: __ccgo_ts + 17003,
		Fusb_code:    uint16(0x0420),
	},
	127: {
		Fname:        __ccgo_ts + 17006,
		Fstring_code: __ccgo_ts + 17023,
		Fusb_code:    uint16(0x0843),
	},
	128: {
		Fname:        __ccgo_ts + 17029,
		Fstring_code: __ccgo_ts + 17023,
		Fusb_code:    uint16(0x0443),
	},
	129: {
		Fname:        __ccgo_ts + 17045,
		Fstring_code: __ccgo_ts + 17056,
		Fusb_code:    uint16(0x042A),
	},
	130: {
		Fname:        __ccgo_ts + 17059,
		Fstring_code: __ccgo_ts + 17065,
		Fusb_code:    uint16(0x0434),
	},
	131: {
		Fname:        __ccgo_ts + 17068,
		Fstring_code: __ccgo_ts + 17076,
		Fusb_code:    uint16(0x043D),
	},
	132: {
		Fname:        __ccgo_ts + 17079,
		Fstring_code: __ccgo_ts + 17084,
		Fusb_code:    uint16(0x0435),
	},
	133: {},
}

func get_usb_code_for_current_locale(tls *libc.TLS) (r uint16_t) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var lang, locale, ptr uintptr
	var _ /* search_string at bp+0 */ [64]uint8
	_, _, _ = lang, locale, ptr
	locale = libc.Xsetlocale(tls, 0, libc.UintptrFromInt32(0))
	if !(locale != 0) {
		return uint16(0x0)
	}
	libc.Xstrncpy(tls, bp, locale, libc.Uint64FromInt64(64)-libc.Uint64FromInt32(1))
	(*(*[64]uint8)(unsafe.Pointer(bp)))[libc.Uint64FromInt64(64)-libc.Uint64FromInt32(1)] = uint8('\000')
	ptr = bp
	for *(*uint8)(unsafe.Pointer(ptr)) != 0 {
		*(*uint8)(unsafe.Pointer(ptr)) = libc.Uint8FromInt32(libc.Xtolower(tls, libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ptr)))))
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ptr))) == int32('.') {
			*(*uint8)(unsafe.Pointer(ptr)) = uint8('\000')
			break
		}
		ptr++
	}
	lang = uintptr(unsafe.Pointer(&lang_map))
	for (*lang_map_entry)(unsafe.Pointer(lang)).Fstring_code != 0 {
		if !(libc.Xstrcmp(tls, (*lang_map_entry)(unsafe.Pointer(lang)).Fstring_code, bp) != 0) {
			return (*lang_map_entry)(unsafe.Pointer(lang)).Fusb_code
		}
		lang += 24
	}
	ptr = bp
	for *(*uint8)(unsafe.Pointer(ptr)) != 0 {
		*(*uint8)(unsafe.Pointer(ptr)) = libc.Uint8FromInt32(libc.Xtolower(tls, libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ptr)))))
		if libc.Int32FromUint8(*(*uint8)(unsafe.Pointer(ptr))) == int32('_') {
			*(*uint8)(unsafe.Pointer(ptr)) = uint8('\000')
			break
		}
		ptr++
	}
	return uint16(0x0)
}

func __ccgo_fp(f interface{}) uintptr {
	type iface [2]uintptr
	return (*iface)(unsafe.Pointer(&f))[1]
}

var active_contexts_list list_head

var active_contexts_lock = usbi_mutex_static_t{}

var linux_hotplug_lock = usbi_mutex_static_t{}

var usbi_backend = usbi_os_backend{
	Fname:                    __ccgo_ts,
	Fcaps:                    libc.Uint32FromInt32(libc.Int32FromInt32(0x00010000) | libc.Int32FromInt32(0x00020000)),
	Fcontext_priv_size:       uint64(4),
	Fdevice_priv_size:        uint64(40),
	Fdevice_handle_priv_size: uint64(16),
	Ftransfer_priv_size:      uint64(32),
}

var usbi_default_context uintptr

var usbi_fallback_context uintptr

var __ccgo_ts = (*reflect.StringHeader)(unsafe.Pointer(&__ccgo_ts1)).Data

var __ccgo_ts1 = "Linux usbfs\x00pthread_mutex_lock(mutex) == 0\x00/home/builduser/hid/libusb/libusb/os/threads_posix.h\x00pthread_mutex_unlock(mutex) == 0\x00pthread_mutex_init(mutex, ((void*)0)) == 0\x00pthread_mutex_destroy(mutex) == 0\x00pthread_cond_wait(cond, mutex) == 0\x00pthread_cond_signal(cond) == 0\x00pthread_cond_broadcast(cond) == 0\x00pthread_cond_destroy(cond) == 0\x00pthread_key_create(key, ((void*)0)) == 0\x00pthread_setspecific(key, ptr) == 0\x00pthread_key_delete(key) == 0\x00API misuse! Using non-default context as implicit default.\x00clock_gettime(1, tp) == 0\x00/home/builduser/hid/libusb/libusb/libusbi.h\x00clock_gettime(0, tp) == 0\x00failed to create pipe, errno=%d\x00failed to get pipe fd status flags, errno=%d\x00failed to set pipe fd status flags, errno=%d\x00failed to close pipe write end, errno=%d\x00failed to close pipe read end, errno=%d\x00event write failed\x00event read failed\x00poll() %u fds with timeout in %dms\x00poll() returned %d\x00poll() failed, errno=%d\x00fd %d was removed, ignoring raised events\x00num_ready > 0\x00/home/builduser/hid/libusb/libusb/os/events_posix.c\x00pthread_cond_init(cond, ((void*)0)) == 0\x00/home/builduser/hid/libusb/libusb/os/threads_posix.c\x00/dev/usbdev%u.%u\x00/dev/bus/usb/%03u/%03u\x00File doesn't exist, wait %ld ms and try again\x00libusb couldn't open USB device %s, errno=%d\x00libusb requires write access to USB device nodes\x00usbdev%d.%d\x00invalid usbdev format '%s'\x00found: %s\x00/dev/bus/usb\x00/dev\x00uname failed, errno=%d\x00%d.%d.%d\x00failed to parse uname release '%s'\x00reported kernel version is %s\x00kernel version is too old (reported as %d.%d.%d)\x00could not find usbfs, defaulting to %s\x00found usbfs at %s\x00max iso packet length is (likely) %u bytes\x00/sys\x00sysfs is available\x00sysfs not mounted\x00error starting hotplug event monitor\x00init_count != 0\x00/home/builduser/hid/libusb/libusb/os/linux_usbfs.c\x00no device discovery will be performed\x00manufacturer\x00product\x00serial\x00attribute %s read failed, errno=%zd\x00/sys/bus/usb/devices/%s/%s\x00open %s failed, errno=%d\x00attribute %s doesn't have numeric value?\x00attribute %s doesn't end with newline?\x00attribute %s contains an invalid value: '%s'\x00bConfigurationValue\x00getting address for device: %s detached: %d\x00/proc/self/fd/%d\x00/dev/bus/usb/%hhu/%hhu\x00scan %s\x00busnum\x00devnum\x00bus=%u dev=%u\x00remaining descriptor length too small %zu/2\x00invalid descriptor bLength %hhu\x00bLength overflow by %zu bytes\x00config descriptor not found\x00short descriptor read %zu/%d\x00descriptor is not a config desc (type 0x%02x)\x00invalid descriptor bLength %u\x00invalid wTotalLength %u\x00config length mismatch wTotalLength %u real %u\x00short descriptor read %zu/%u\x00device has configuration 0\x00device unconfigured\x00get configuration failed, errno=%d\x00Error getting device speed: %d\x00speed\x00unknown device speed: %d Mbps\x00descriptors\x00lseek failed, errno=%d\x00read descriptor failed, errno=%d\x00short descriptor read (%zu)\x00Missing rw usbfs access; cannot determine active configuration descriptor\x00usb\x00Can not parse sysfs_dir: %s, unexpected parent info\x00Can not parse sysfs_dir: %s, no parent info\x00usb%s\x00parent_dev %s not enumerated yet, enumerating now\x00dev %p (%s) has parent %p (%s) port %u\x00busnum %u devaddr %u session_id %lu\x00session_id %lu already exists\x00allocating new device for %u/%u (session %lu)\x00device not found for session %lx\x00/dev/bus/usb/%03u\x00%s\x00opendir '%s' failed, errno=%d\x00unknown dir entry %s\x00failed to enumerate dir entry %s\x00opendir buses failed, errno=%d\x00/sys/bus/usb/devices\x00opendir devices failed, errno=%d\x00getcap not available\x00getcap failed, errno=%d\x00connectinfo failed, errno=%d\x00allocating new device for fd %d\x00open failed with no device, but device still attached\x00set configuration failed, errno=%d\x00claim interface failed, errno=%d\x00release interface failed, errno=%d\x00set interface failed, errno=%d\x00clear halt failed, errno=%d\x00reset failed, errno=%d\x00failed to re-claim interface %u after reset: %s\x00streams-ioctl failed, errno=%d\x00alloc dev mem failed, errno=%d\x00free dev mem failed, errno=%d\x00get driver failed, errno=%d\x00usbfs\x00detach failed, errno=%d\x00attach failed, errno=%d\x00disconnect-and-claim failed, errno=%d\x00URB not found --> assuming ready to be reaped\x00Device not found for URB --> assuming ready to be reaped\x00unrecognised discard errno %d\x00need %d urbs for new transfer with length %d\x00submiturb failed, errno=%d\x00first URB failed, easy peasy\x00reporting successful submission but waiting for %d discards before reporting error\x00iso packet length of %u bytes exceeds maximum of %u bytes\x00submiturb failed, transfer too large\x00submiturb failed, iso packet length too large\x00unknown transfer type %u\x00handling completion status %d of bulk urb %d/%d\x00abnormal reap: urb status %d\x00received %d bytes of surplus data\x00moving surplus data from offset %zu to offset %zu\x00abnormal reap: last URB handled, reporting\x00device removed\x00detected endpoint stall\x00overflow, actual_length=%d\x00low-level bus error %d\x00unrecognised urb status %d\x00all URBs in transfer reaped --> complete!\x00short transfer %d/%d --> complete!\x00could not locate urb!\x00handling completion status %d of iso urb %d/%d\x00packet %d - device removed\x00packet %d - detected endpoint stall\x00packet %d - overflow error\x00packet %d - low-level USB error %d\x00packet %d - unrecognised urb status %d\x00CANCEL: urb status %d\x00CANCEL: last URB handled, reporting\x00handling completion status %d\x00cancel: unrecognised urb status %d\x00unsupported control request\x00reap failed, errno=%d\x00urb type=%u status=%d transferred=%d\x00unrecognised transfer type %u\x00cannot find handle for fd %d\x00failed to get netlink fd flags, errno=%d\x00failed to set netlink fd flags, errno=%d\x00failed to get netlink fd status flags, errno=%d\x00failed to set netlink fd status flags, errno=%d\x00failed to create netlink socket of type %d, attempting SOCK_RAW\x00failed to create netlink socket, errno=%d\x00failed to bind netlink socket, errno=%d\x00failed to set netlink socket SO_PASSCRED option, errno=%d\x00failed to create netlink control event\x00failed to create netlink event thread (%d)\x00linux_netlink_socket != -1\x00/home/builduser/hid/libusb/libusb/os/linux_netlink.c\x00failed to join netlink event thread (%d)\x00ACTION\x00remove\x00add\x00unknown device action %s\x00SUBSYSTEM\x00DEVTYPE\x00usb_device\x00BUSNUM\x00DEVNUM\x00DEVICE\x00DEVPATH\x00error receiving message from netlink, errno=%d\x00invalid netlink message length\x00ignoring netlink message from unknown group/PID (%u/%u)\x00ignoring netlink message with no sender credentials\x00ignoring netlink message with non-zero sender UID %u\x00yes\x00no\x00netlink hotplug found device busnum: %hhu, devaddr: %hhu, sys_name: %s, removed: %s\x00netlink event thread entering\x00netlink event thread exiting\x00-rc1\x00https://libusb.info\x00need to increase capacity\x00invalid device descriptor\x00too many configurations\x00zero configurations, maybe an unauthorized device\x00 \x00port numbers array is too small\x00could not retrieve active config descriptor\x00refcnt >= 2\x00/home/builduser/hid/libusb/libusb/core.c\x00refcnt >= 0\x00destroy device %d.%d\x00wrap_sys_device 0x%lx\x00wrap_sys_device 0x%lx returns %d\x00open %d.%d\x00open %d.%d returns %d\x00Device handle closed while transfer was still being processed, but the device is still connected as far as we know\x00A cancellation for an in-flight transfer hasn't completed but closing the device handle\x00A cancellation hasn't even been scheduled on the transfer for which the device is closing\x00Removed transfer %p from the in-flight list because device handle %p closed\x00falling back to control message\x00zero bytes returned in ctrl transfer?\x00control failed, error %d\x00active config %u\x00configuration %d\x00interface %d\x00interface %d altsetting %d\x00endpoint 0x%x\x00streams %u eps %d\x00eps %d\x00usbi_backend.endpoint_set_raw_io != NULL\x00usbi_backend.get_max_raw_io_transfer_size != NULL\x00LIBUSB_DEBUG\x00reusing default context\x00created default context\x00libusb v%u.%u.%u.%u%s\x00installing new context as implicit default\x00no default context, not initialized?\x00not destroying default context\x00destroying default context\x00device %d.%d still referenced\x00application left some devices open\x00error\x00warning\x00info\x00debug\x00unknown\x00[timestamp] [threadID] facility level [function call] <message>\n\x00--------------------------------------------------------------------------------\n\x00[%2ld.%06ld] [%08lx] libusb: %s [%s] \x00libusb: %s [%s] \x00\n\x00LIBUSB_ERROR_IO\x00LIBUSB_ERROR_INVALID_PARAM\x00LIBUSB_ERROR_ACCESS\x00LIBUSB_ERROR_NO_DEVICE\x00LIBUSB_ERROR_NOT_FOUND\x00LIBUSB_ERROR_BUSY\x00LIBUSB_ERROR_TIMEOUT\x00LIBUSB_ERROR_OVERFLOW\x00LIBUSB_ERROR_PIPE\x00LIBUSB_ERROR_INTERRUPTED\x00LIBUSB_ERROR_NO_MEM\x00LIBUSB_ERROR_NOT_SUPPORTED\x00LIBUSB_ERROR_OTHER\x00LIBUSB_TRANSFER_ERROR\x00LIBUSB_TRANSFER_TIMED_OUT\x00LIBUSB_TRANSFER_CANCELLED\x00LIBUSB_TRANSFER_STALL\x00LIBUSB_TRANSFER_NO_DEVICE\x00LIBUSB_TRANSFER_OVERFLOW\x00LIBUSB_SUCCESS / LIBUSB_TRANSFER_COMPLETED\x00**UNKNOWN**\x00short endpoint descriptor read %d/%d\x00unexpected descriptor 0x%x (expected 0x%x)\x00invalid endpoint bLength (%u)\x00short endpoint descriptor read %d/%u\x00invalid extra ep desc len (%u)\x00short extra ep desc read %d/%u\x00skipping descriptor 0x%x\x00invalid interface bLength (%u)\x00short intf descriptor read %d/%u\x00too many endpoints (%u)\x00invalid extra intf desc len (%u)\x00short extra intf desc read %d/%u\x00short config descriptor read %d/%d\x00invalid config bLength (%u)\x00short config descriptor read %d/%u\x00too many interfaces (%u)\x00invalid extra config desc len (%u)\x00short extra config desc read %d/%u\x00parse_configuration failed with error %d\x00still %d bytes of descriptor data left\x00index %u\x00value %u\x00invalid descriptor length %u\x00invalid ss-ep-comp-desc length %u\x00short ss-ep-comp-desc read %d/%u\x00short bos descriptor read %d/%d\x00invalid bos bLength (%u)\x00short bos descriptor read %d/%u\x00short dev-cap descriptor read %d/%d\x00invalid dev-cap bLength (%u)\x00short dev-cap descriptor read %d/%u\x00failed to read BOS (%d)\x00short BOS read %d/%d\x00found BOS descriptor: size %u bytes, %u capabilities\x00short BOS read %d/%u\x00unexpected bDevCapabilityType 0x%x (expected 0x%x)\x00short dev-cap descriptor read %u/%d\x00short ssplus capability descriptor, unable to read sublinks: Not enough data\x00invalid language ID string descriptor\x00suspicious bLength %u for language ID string descriptor\x00suspicious bLength %u for string descriptor (read %d)\x00invalid descriptor bLength %d\x00parse_iad_array failed with error %d\x00IADs for config index %u\x00dev->parent_dev != next_dev\x00/home/builduser/hid/libusb/libusb/hotplug.c\x00error allocating hotplug message\x00freeing hotplug cb %p with handle %d\x00new hotplug cb %p with handle %d\x00deregister hotplug cb %d\x00get hotplug cb %d user data\x00iso_packets >= 0\x00/home/builduser/hid/libusb/libusb/io.c\x00transfer %p\x00ptr == itransfer->priv\x00transfer->dev_handle\x00cancel transfer failed error %d\x00failed to set timer for next timeout\x00interpreting short transfer as error\x00transfer->actual_length >= 0\x00transfer %p has callback %p\x00detected timeout cancellation\x00someone else is closing a device\x00async cancel failed %d\x00event triggered\x00someone updated the event sources\x00someone purposefully interrupted\x00someone unregistered a hotplug cb\x00someone is closing a device\x00hotplug message received\x00!list_empty(&ctx->hotplug_msgs)\x00!list_empty(&ctx->completed_transfers)\x00backend handle_transfer_completion failed with error %d\x00event sources modified, reallocating event data\x00backend handle_events failed with error %d\x00doing our own event handling\x00event handler was active but went away, retrying\x00another thread is doing event handling\x00no URBs, no timeout!\x00no URB with timeout or all handled by OS; no timeout!\x00first timeout already expired\x00next timeout in %ld.%06lds\x00add fd %d events %d\x00remove fd %d\x00couldn't find fd %d to remove\x00device %d.%d\x00cancelling transfer %p from disconnect\x00en\x00nl\x00fr\x00ru\x00de\x00hu\x00Success\x00Input/Output Error\x00Invalid parameter\x00Access denied (insufficient permissions)\x00No such device (it may have been disconnected)\x00Entity not found\x00Resource busy\x00Operation timed out\x00Overflow\x00Pipe error\x00System call interrupted (perhaps due to signal)\x00Insufficient memory\x00Operation not supported or unimplemented on this platform\x00Other error\x00Gelukt\x00Invoer-/uitvoerfout\x00Ongeldig argument\x00Toegang geweigerd (onvoldoende toegangsrechten)\x00Apparaat bestaat niet (verbinding met apparaat verbroken?)\x00Niet gevonden\x00Apparaat of hulpbron is bezig\x00Bewerking verlopen\x00Waarde is te groot\x00Gebroken pijp\x00Onderbroken systeemaanroep\x00Onvoldoende geheugen beschikbaar\x00Bewerking wordt niet ondersteund\x00Andere fout\x00Succès\x00Erreur d'entrée/sortie\x00Paramètre invalide\x00Accès refusé (permissions insuffisantes)\x00Périphérique introuvable (peut-être déconnecté)\x00Elément introuvable\x00Resource déjà occupée\x00Operation expirée\x00Débordement\x00Erreur de pipe\x00Appel système abandonné (peut-être à cause d’un signal)\x00Mémoire insuffisante\x00Opération non supportée or non implémentée sur cette plateforme\x00Autre erreur\x00Успех\x00Ошибка ввода/вывода\x00Неверный параметр\x00Доступ запрещён (не хватает прав)\x00Устройство отсутствует (возможно, оно было отсоединено)\x00Элемент не найден\x00Ресурс занят\x00Истекло время ожидания операции\x00Переполнение\x00Ошибка канала\x00Системный вызов прерван (возможно, сигналом)\x00Память исчерпана\x00Операция не поддерживается данной платформой\x00Неизвестная ошибка\x00Erfolgreich\x00Eingabe-/Ausgabefehler\x00Ungültiger Parameter\x00Keine Berechtigung (Zugriffsrechte fehlen)\x00Kein passendes Gerät gefunden (es könnte entfernt worden sein)\x00Entität nicht gefunden\x00Die Ressource ist belegt\x00Die Wartezeit für die Operation ist abgelaufen\x00Mehr Daten empfangen als erwartet\x00Datenübergabe unterbrochen (broken pipe)\x00Unterbrechung während des Betriebssystemaufrufs\x00Nicht genügend Hauptspeicher verfügbar\x00Die Operation wird nicht unterstützt oder ist auf dieser Platform nicht implementiert\x00Allgemeiner Fehler\x00Sikeres\x00Be-/kimeneti hiba\x00Érvénytelen paraméter\x00Hozzáférés megtagadva\x00Az eszköz nem található (eltávolították?)\x00Nem található\x00Az erőforrás foglalt\x00Időtúllépés\x00Túlcsordulás\x00Törött adatcsatorna\x00Rendszerhívás megszakítva\x00Nincs elég memória\x00A művelet nem támogatott ezen a rendszeren\x00Általános hiba\x00actual_length=%d\x00libusb_handle_events failed: %s, cancelling transfer and retrying\x00unrecognised status code %d\x00/home/builduser/hid/libusb/libusb/sync.c\x00WCHAR_T\x00UTF-16LE\x00%u-%u\x00.%u\x00:%u.%u\x000.16.0\x00\x00\x00\x00\x00h\x00\x00\x00i\x00\x00\x00d\x00\x00\x00_\x00\x00\x00r\x00\x00\x00e\x00\x00\x00a\x00\x00\x00d\x00\x00\x00_\x00\x00\x00e\x00\x00\x00r\x00\x00\x00r\x00\x00\x00o\x00\x00\x00r\x00\x00\x00 \x00\x00\x00i\x00\x00\x00s\x00\x00\x00 \x00\x00\x00n\x00\x00\x00o\x00\x00\x00t\x00\x00\x00 \x00\x00\x00i\x00\x00\x00m\x00\x00\x00p\x00\x00\x00l\x00\x00\x00e\x00\x00\x00m\x00\x00\x00e\x00\x00\x00n\x00\x00\x00t\x00\x00\x00e\x00\x00\x00d\x00\x00\x00 \x00\x00\x00y\x00\x00\x00e\x00\x00\x00t\x00\x00\x00\x00\x00\x00\x00h\x00\x00\x00i\x00\x00\x00d\x00\x00\x00_\x00\x00\x00e\x00\x00\x00r\x00\x00\x00r\x00\x00\x00o\x00\x00\x00r\x00\x00\x00 \x00\x00\x00i\x00\x00\x00s\x00\x00\x00 \x00\x00\x00n\x00\x00\x00o\x00\x00\x00t\x00\x00\x00 \x00\x00\x00i\x00\x00\x00m\x00\x00\x00p\x00\x00\x00l\x00\x00\x00e\x00\x00\x00m\x00\x00\x00e\x00\x00\x00n\x00\x00\x00t\x00\x00\x00e\x00\x00\x00d\x00\x00\x00 \x00\x00\x00y\x00\x00\x00e\x00\x00\x00t\x00\x00\x00\x00Afrikaans\x00af\x00Albanian\x00sq\x00Arabic - United Arab Emirates\x00ar_ae\x00Arabic - Bahrain\x00ar_bh\x00Arabic - Algeria\x00ar_dz\x00Arabic - Egypt\x00ar_eg\x00Arabic - Iraq\x00ar_iq\x00Arabic - Jordan\x00ar_jo\x00Arabic - Kuwait\x00ar_kw\x00Arabic - Lebanon\x00ar_lb\x00Arabic - Libya\x00ar_ly\x00Arabic - Morocco\x00ar_ma\x00Arabic - Oman\x00ar_om\x00Arabic - Qatar\x00ar_qa\x00Arabic - Saudi Arabia\x00ar_sa\x00Arabic - Syria\x00ar_sy\x00Arabic - Tunisia\x00ar_tn\x00Arabic - Yemen\x00ar_ye\x00Armenian\x00hy\x00Azeri - Latin\x00az_az\x00Azeri - Cyrillic\x00Basque\x00eu\x00Belarusian\x00be\x00Bulgarian\x00bg\x00Catalan\x00ca\x00Chinese - China\x00zh_cn\x00Chinese - Hong Kong SAR\x00zh_hk\x00Chinese - Macau SAR\x00zh_mo\x00Chinese - Singapore\x00zh_sg\x00Chinese - Taiwan\x00zh_tw\x00Croatian\x00hr\x00Czech\x00cs\x00Danish\x00da\x00Dutch - Netherlands\x00nl_nl\x00Dutch - Belgium\x00nl_be\x00English - Australia\x00en_au\x00English - Belize\x00en_bz\x00English - Canada\x00en_ca\x00English - Caribbean\x00en_cb\x00English - Ireland\x00en_ie\x00English - Jamaica\x00en_jm\x00English - New Zealand\x00en_nz\x00English - Philippines\x00en_ph\x00English - Southern Africa\x00en_za\x00English - Trinidad\x00en_tt\x00English - Great Britain\x00en_gb\x00English - United States\x00en_us\x00Estonian\x00et\x00Farsi\x00fa\x00Finnish\x00fi\x00Faroese\x00fo\x00French - France\x00fr_fr\x00French - Belgium\x00fr_be\x00French - Canada\x00fr_ca\x00French - Luxembourg\x00fr_lu\x00French - Switzerland\x00fr_ch\x00Gaelic - Ireland\x00gd_ie\x00Gaelic - Scotland\x00gd\x00German - Germany\x00de_de\x00German - Austria\x00de_at\x00German - Liechtenstein\x00de_li\x00German - Luxembourg\x00de_lu\x00German - Switzerland\x00de_ch\x00Greek\x00el\x00Hebrew\x00he\x00Hindi\x00hi\x00Hungarian\x00Icelandic\x00is\x00Indonesian\x00id\x00Italian - Italy\x00it_it\x00Italian - Switzerland\x00it_ch\x00Japanese\x00ja\x00Korean\x00ko\x00Latvian\x00lv\x00Lithuanian\x00lt\x00F.Y.R.O. Macedonia\x00mk\x00Malay - Malaysia\x00ms_my\x00Malay – Brunei\x00ms_bn\x00Maltese\x00mt\x00Marathi\x00mr\x00Norwegian - Bokml\x00no_no\x00Norwegian - Nynorsk\x00Polish\x00pl\x00Portuguese - Portugal\x00pt_pt\x00Portuguese - Brazil\x00pt_br\x00Raeto-Romance\x00rm\x00Romanian - Romania\x00ro\x00Romanian - Republic of Moldova\x00ro_mo\x00Russian\x00Russian - Republic of Moldova\x00ru_mo\x00Sanskrit\x00sa\x00Serbian - Cyrillic\x00sr_sp\x00Serbian - Latin\x00Setsuana\x00tn\x00Slovenian\x00sl\x00Slovak\x00sk\x00Sorbian\x00sb\x00Spanish - Spain (Traditional)\x00es_es\x00Spanish - Argentina\x00es_ar\x00Spanish - Bolivia\x00es_bo\x00Spanish - Chile\x00es_cl\x00Spanish - Colombia\x00es_co\x00Spanish - Costa Rica\x00es_cr\x00Spanish - Dominican Republic\x00es_do\x00Spanish - Ecuador\x00es_ec\x00Spanish - Guatemala\x00es_gt\x00Spanish - Honduras\x00es_hn\x00Spanish - Mexico\x00es_mx\x00Spanish - Nicaragua\x00es_ni\x00Spanish - Panama\x00es_pa\x00Spanish - Peru\x00es_pe\x00Spanish - Puerto Rico\x00es_pr\x00Spanish - Paraguay\x00es_py\x00Spanish - El Salvador\x00es_sv\x00Spanish - Uruguay\x00es_uy\x00Spanish - Venezuela\x00es_ve\x00Southern Sotho\x00st\x00Swahili\x00sw\x00Swedish - Sweden\x00sv_se\x00Swedish - Finland\x00sv_fi\x00Tamil\x00ta\x00Tatar\x00tt\x00Thai\x00th\x00Turkish\x00tr\x00Tsonga\x00ts\x00Ukrainian\x00uk\x00Urdu\x00ur\x00Uzbek - Cyrillic\x00uz_uz\x00Uzbek – Latin\x00Vietnamese\x00vi\x00Xhosa\x00xh\x00Yiddish\x00yi\x00Zulu\x00zu\x00"
