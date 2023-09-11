// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package neon

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __native_entry__() uintptr

var (
    _subr__f32toa             uintptr = __native_entry__() + 104412
    _subr__f64toa             uintptr = __native_entry__() + 96
    _subr__get_by_path        uintptr = __native_entry__() + 86928
    _subr__html_escape        uintptr = __native_entry__() + 11520
    _subr__i64toa             uintptr = __native_entry__() + 4144
    _subr__lspace             uintptr = __native_entry__() + 16
    _subr__quote              uintptr = __native_entry__() + 8256
    _subr__skip_array         uintptr = __native_entry__() + 49544
    _subr__skip_number        uintptr = __native_entry__() + 67048
    _subr__skip_object        uintptr = __native_entry__() + 58304
    _subr__skip_one           uintptr = __native_entry__() + 68116
    _subr__skip_one_fast      uintptr = __native_entry__() + 83348
    _subr__u64toa             uintptr = __native_entry__() + 6848
    _subr__unquote            uintptr = __native_entry__() + 9728
    _subr__validate_one       uintptr = __native_entry__() + 76876
    _subr__validate_utf8      uintptr = __native_entry__() + 103244
    _subr__validate_utf8_fast uintptr = __native_entry__() + 103900
    _subr__value              uintptr = __native_entry__() + 17456
    _subr__vnumber            uintptr = __native_entry__() + 29736
    _subr__vsigned            uintptr = __native_entry__() + 35948
    _subr__vstring            uintptr = __native_entry__() + 27728
    _subr__vunsigned          uintptr = __native_entry__() + 36316
)

const (
    _stack__f32toa = 0
    _stack__f64toa = 16
    _stack__get_by_path = 192
    _stack__html_escape = 16
    _stack__i64toa = 0
    _stack__lspace = 0
    _stack__quote = 0
    _stack__skip_array = 128
    _stack__skip_number = 16
    _stack__skip_object = 128
    _stack__skip_one = 128
    _stack__skip_one_fast = 112
    _stack__u64toa = 0
    _stack__unquote = 96
    _stack__validate_one = 128
    _stack__validate_utf8 = 32
    _stack__validate_utf8_fast = 16
    _stack__value = 64
    _stack__vnumber = 80
    _stack__vsigned = 0
    _stack__vstring = 16
    _stack__vunsigned = 0
)

var (
    _ = _subr__f32toa
    _ = _subr__f64toa
    _ = _subr__get_by_path
    _ = _subr__html_escape
    _ = _subr__i64toa
    _ = _subr__lspace
    _ = _subr__quote
    _ = _subr__skip_array
    _ = _subr__skip_number
    _ = _subr__skip_object
    _ = _subr__skip_one
    _ = _subr__skip_one_fast
    _ = _subr__u64toa
    _ = _subr__unquote
    _ = _subr__validate_one
    _ = _subr__validate_utf8
    _ = _subr__validate_utf8_fast
    _ = _subr__value
    _ = _subr__vnumber
    _ = _subr__vsigned
    _ = _subr__vstring
    _ = _subr__vunsigned
)

const (
    _ = _stack__f32toa
    _ = _stack__f64toa
    _ = _stack__get_by_path
    _ = _stack__html_escape
    _ = _stack__i64toa
    _ = _stack__lspace
    _ = _stack__quote
    _ = _stack__skip_array
    _ = _stack__skip_number
    _ = _stack__skip_object
    _ = _stack__skip_one
    _ = _stack__skip_one_fast
    _ = _stack__u64toa
    _ = _stack__unquote
    _ = _stack__validate_one
    _ = _stack__validate_utf8
    _ = _stack__validate_utf8_fast
    _ = _stack__value
    _ = _stack__vnumber
    _ = _stack__vsigned
    _ = _stack__vstring
    _ = _stack__vunsigned
)
