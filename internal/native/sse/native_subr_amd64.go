// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__count_elems = 26192
    _entry__skip_one_fast = 22288
    _entry__count_elems_fast = 26784
    _entry__f32toa = 33136
    _entry__f64toa = 160
    _entry__format_significand = 37472
    _entry__format_integer = 2960
    _entry__fsm_exec = 18016
    _entry__advance_string = 14352
    _entry__advance_string_default = 38864
    _entry__do_skip_number = 20608
    _entry__get_by_path = 27696
    _entry__html_escape = 8912
    _entry__i64toa = 3392
    _entry__u64toa = 3520
    _entry__lspace = 16
    _entry__quote = 4832
    _entry__skip_array = 17984
    _entry__skip_number = 21904
    _entry__skip_object = 20256
    _entry__skip_one = 22064
    _entry__unquote = 6576
    _entry__validate_one = 22112
    _entry__validate_utf8 = 31904
    _entry__validate_utf8_fast = 32576
    _entry__value = 12352
    _entry__vnumber = 15744
    _entry__atof_eisel_lemire64 = 10192
    _entry__atof_native = 11744
    _entry__decimal_to_f64 = 10560
    _entry__right_shift = 38432
    _entry__left_shift = 37936
    _entry__vsigned = 17296
    _entry__vstring = 14176
    _entry__vunsigned = 17632
)

const (
    _stack__count_elems = 192
    _stack__skip_one_fast = 136
    _stack__count_elems_fast = 208
    _stack__f32toa = 48
    _stack__f64toa = 80
    _stack__format_significand = 24
    _stack__format_integer = 16
    _stack__fsm_exec = 168
    _stack__advance_string = 64
    _stack__advance_string_default = 64
    _stack__do_skip_number = 48
    _stack__get_by_path = 272
    _stack__html_escape = 72
    _stack__i64toa = 16
    _stack__u64toa = 8
    _stack__lspace = 8
    _stack__quote = 64
    _stack__skip_array = 176
    _stack__skip_number = 104
    _stack__skip_object = 176
    _stack__skip_one = 176
    _stack__unquote = 88
    _stack__validate_one = 176
    _stack__validate_utf8 = 48
    _stack__validate_utf8_fast = 24
    _stack__value = 328
    _stack__vnumber = 240
    _stack__atof_eisel_lemire64 = 32
    _stack__atof_native = 136
    _stack__decimal_to_f64 = 80
    _stack__right_shift = 8
    _stack__left_shift = 24
    _stack__vsigned = 16
    _stack__vstring = 120
    _stack__vunsigned = 8
)

const (
    _size__count_elems = 544
    _size__skip_one_fast = 3404
    _size__count_elems_fast = 912
    _size__f32toa = 3328
    _size__f64toa = 2800
    _size__format_significand = 464
    _size__format_integer = 432
    _size__fsm_exec = 1692
    _size__advance_string = 1344
    _size__advance_string_default = 960
    _size__do_skip_number = 956
    _size__get_by_path = 4208
    _size__html_escape = 1280
    _size__i64toa = 48
    _size__u64toa = 1264
    _size__lspace = 128
    _size__quote = 1728
    _size__skip_array = 32
    _size__skip_number = 160
    _size__skip_object = 32
    _size__skip_one = 48
    _size__unquote = 2272
    _size__validate_one = 48
    _size__validate_utf8 = 672
    _size__validate_utf8_fast = 544
    _size__value = 1308
    _size__vnumber = 1552
    _size__atof_eisel_lemire64 = 368
    _size__atof_native = 608
    _size__decimal_to_f64 = 1184
    _size__right_shift = 400
    _size__left_shift = 496
    _size__vsigned = 336
    _size__vstring = 128
    _size__vunsigned = 336
)

var (
    _pcsp__count_elems = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {521, 56},
        {525, 48},
        {526, 40},
        {528, 32},
        {530, 24},
        {532, 16},
        {534, 8},
        {536, 0},
    }
    _pcsp__skip_one_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {658, 136},
        {662, 48},
        {663, 40},
        {665, 32},
        {667, 24},
        {669, 16},
        {671, 8},
        {672, 0},
        {3404, 136},
    }
    _pcsp__count_elems_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {889, 72},
        {893, 40},
        {894, 32},
        {896, 24},
        {898, 16},
        {900, 8},
        {902, 0},
    }
    _pcsp__f32toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {3286, 48},
        {3287, 40},
        {3289, 32},
        {3291, 24},
        {3293, 16},
        {3295, 8},
        {3296, 0},
        {3318, 48},
    }
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {2740, 56},
        {2744, 48},
        {2745, 40},
        {2747, 32},
        {2749, 24},
        {2751, 16},
        {2753, 8},
        {2754, 0},
        {2792, 56},
    }
    _pcsp__format_significand = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {452, 24},
        {453, 16},
        {455, 8},
        {457, 0},
    }
    _pcsp__format_integer = [][2]uint32{
        {1, 0},
        {4, 8},
        {412, 16},
        {413, 8},
        {414, 0},
        {423, 16},
        {424, 8},
        {426, 0},
    }
    _pcsp__fsm_exec = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1342, 104},
        {1346, 48},
        {1347, 40},
        {1349, 32},
        {1351, 24},
        {1353, 16},
        {1355, 8},
        {1356, 0},
        {1692, 104},
    }
    _pcsp__advance_string = [][2]uint32{
        {14, 0},
        {18, 8},
        {20, 16},
        {22, 24},
        {24, 32},
        {26, 40},
        {27, 48},
        {614, 56},
        {618, 48},
        {619, 40},
        {621, 32},
        {623, 24},
        {625, 16},
        {627, 8},
        {628, 0},
        {1339, 56},
    }
    _pcsp__advance_string_default = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {576, 64},
        {580, 48},
        {581, 40},
        {583, 32},
        {585, 24},
        {587, 16},
        {589, 8},
        {590, 0},
        {955, 64},
    }
    _pcsp__do_skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {881, 48},
        {882, 40},
        {884, 32},
        {886, 24},
        {888, 16},
        {890, 8},
        {891, 0},
        {956, 48},
    }
    _pcsp__get_by_path = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {4012, 104},
        {4016, 48},
        {4017, 40},
        {4019, 32},
        {4021, 24},
        {4023, 16},
        {4025, 8},
        {4026, 0},
        {4194, 104},
    }
    _pcsp__html_escape = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1256, 72},
        {1260, 48},
        {1261, 40},
        {1263, 32},
        {1265, 24},
        {1267, 16},
        {1269, 8},
        {1271, 0},
    }
    _pcsp__i64toa = [][2]uint32{
        {14, 0},
        {34, 8},
        {36, 0},
    }
    _pcsp__u64toa = [][2]uint32{
        {1, 0},
        {161, 8},
        {162, 0},
        {457, 8},
        {458, 0},
        {772, 8},
        {773, 0},
        {1249, 8},
        {1251, 0},
    }
    _pcsp__lspace = [][2]uint32{
        {1, 0},
        {89, 8},
        {90, 0},
        {103, 8},
        {104, 0},
        {111, 8},
        {113, 0},
    }
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1681, 64},
        {1685, 48},
        {1686, 40},
        {1688, 32},
        {1690, 24},
        {1692, 16},
        {1694, 8},
        {1695, 0},
        {1722, 64},
    }
    _pcsp__skip_array = [][2]uint32{
        {1, 0},
        {26, 8},
        {32, 0},
    }
    _pcsp__skip_number = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {107, 56},
        {111, 48},
        {112, 40},
        {114, 32},
        {116, 24},
        {118, 16},
        {120, 8},
        {121, 0},
        {145, 56},
    }
    _pcsp__skip_object = [][2]uint32{
        {1, 0},
        {26, 8},
        {32, 0},
    }
    _pcsp__skip_one = [][2]uint32{
        {1, 0},
        {30, 8},
        {36, 0},
    }
    _pcsp__unquote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1684, 88},
        {1688, 48},
        {1689, 40},
        {1691, 32},
        {1693, 24},
        {1695, 16},
        {1697, 8},
        {1698, 0},
        {2270, 88},
    }
    _pcsp__validate_one = [][2]uint32{
        {1, 0},
        {35, 8},
        {41, 0},
    }
    _pcsp__validate_utf8 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {623, 48},
        {627, 40},
        {628, 32},
        {630, 24},
        {632, 16},
        {634, 8},
        {635, 0},
        {666, 48},
    }
    _pcsp__validate_utf8_fast = [][2]uint32{
        {1, 0},
        {4, 8},
        {5, 16},
        {247, 24},
        {251, 16},
        {252, 8},
        {253, 0},
        {527, 24},
        {531, 16},
        {532, 8},
        {534, 0},
    }
    _pcsp__value = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {495, 88},
        {499, 48},
        {500, 40},
        {502, 32},
        {504, 24},
        {506, 16},
        {508, 8},
        {509, 0},
        {1308, 88},
    }
    _pcsp__vnumber = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {803, 104},
        {807, 48},
        {808, 40},
        {810, 32},
        {812, 24},
        {814, 16},
        {816, 8},
        {817, 0},
        {1551, 104},
    }
    _pcsp__atof_eisel_lemire64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {292, 32},
        {293, 24},
        {295, 16},
        {297, 8},
        {298, 0},
        {362, 32},
    }
    _pcsp__atof_native = [][2]uint32{
        {1, 0},
        {4, 8},
        {587, 56},
        {591, 8},
        {593, 0},
    }
    _pcsp__decimal_to_f64 = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1144, 56},
        {1148, 48},
        {1149, 40},
        {1151, 32},
        {1153, 24},
        {1155, 16},
        {1157, 8},
        {1158, 0},
        {1169, 56},
    }
    _pcsp__right_shift = [][2]uint32{
        {1, 0},
        {318, 8},
        {319, 0},
        {387, 8},
        {388, 0},
        {396, 8},
        {398, 0},
    }
    _pcsp__left_shift = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {363, 24},
        {364, 16},
        {366, 8},
        {367, 0},
        {470, 24},
        {471, 16},
        {473, 8},
        {474, 0},
        {486, 24},
    }
    _pcsp__vsigned = [][2]uint32{
        {1, 0},
        {4, 8},
        {119, 16},
        {120, 8},
        {121, 0},
        {132, 16},
        {133, 8},
        {134, 0},
        {276, 16},
        {277, 8},
        {278, 0},
        {282, 16},
        {283, 8},
        {284, 0},
        {322, 16},
        {323, 8},
        {324, 0},
        {332, 16},
        {333, 8},
        {335, 0},
    }
    _pcsp__vstring = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {11, 40},
        {105, 56},
        {109, 40},
        {110, 32},
        {112, 24},
        {114, 16},
        {116, 8},
        {118, 0},
    }
    _pcsp__vunsigned = [][2]uint32{
        {1, 0},
        {78, 8},
        {79, 0},
        {90, 8},
        {91, 0},
        {114, 8},
        {115, 0},
        {273, 8},
        {274, 0},
        {312, 8},
        {313, 0},
        {320, 8},
        {322, 0},
    }
)

var Funcs = []loader.CFunc{
    {"__native_entry__", 0, 67, 0, nil},
    {"_count_elems", _entry__count_elems, _size__count_elems, _stack__count_elems, _pcsp__count_elems},
    {"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
    {"_count_elems_fast", _entry__count_elems_fast, _size__count_elems_fast, _stack__count_elems_fast, _pcsp__count_elems_fast},
    {"_f32toa", _entry__f32toa, _size__f32toa, _stack__f32toa, _pcsp__f32toa},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
    {"_format_significand", _entry__format_significand, _size__format_significand, _stack__format_significand, _pcsp__format_significand},
    {"_format_integer", _entry__format_integer, _size__format_integer, _stack__format_integer, _pcsp__format_integer},
    {"_fsm_exec", _entry__fsm_exec, _size__fsm_exec, _stack__fsm_exec, _pcsp__fsm_exec},
    {"_advance_string", _entry__advance_string, _size__advance_string, _stack__advance_string, _pcsp__advance_string},
    {"_advance_string_default", _entry__advance_string_default, _size__advance_string_default, _stack__advance_string_default, _pcsp__advance_string_default},
    {"_do_skip_number", _entry__do_skip_number, _size__do_skip_number, _stack__do_skip_number, _pcsp__do_skip_number},
    {"_get_by_path", _entry__get_by_path, _size__get_by_path, _stack__get_by_path, _pcsp__get_by_path},
    {"_html_escape", _entry__html_escape, _size__html_escape, _stack__html_escape, _pcsp__html_escape},
    {"_i64toa", _entry__i64toa, _size__i64toa, _stack__i64toa, _pcsp__i64toa},
    {"_u64toa", _entry__u64toa, _size__u64toa, _stack__u64toa, _pcsp__u64toa},
    {"_lspace", _entry__lspace, _size__lspace, _stack__lspace, _pcsp__lspace},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
    {"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
    {"_skip_number", _entry__skip_number, _size__skip_number, _stack__skip_number, _pcsp__skip_number},
    {"_skip_object", _entry__skip_object, _size__skip_object, _stack__skip_object, _pcsp__skip_object},
    {"_skip_one", _entry__skip_one, _size__skip_one, _stack__skip_one, _pcsp__skip_one},
    {"_unquote", _entry__unquote, _size__unquote, _stack__unquote, _pcsp__unquote},
    {"_validate_one", _entry__validate_one, _size__validate_one, _stack__validate_one, _pcsp__validate_one},
    {"_validate_utf8", _entry__validate_utf8, _size__validate_utf8, _stack__validate_utf8, _pcsp__validate_utf8},
    {"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
    {"_value", _entry__value, _size__value, _stack__value, _pcsp__value},
    {"_vnumber", _entry__vnumber, _size__vnumber, _stack__vnumber, _pcsp__vnumber},
    {"_atof_eisel_lemire64", _entry__atof_eisel_lemire64, _size__atof_eisel_lemire64, _stack__atof_eisel_lemire64, _pcsp__atof_eisel_lemire64},
    {"_atof_native", _entry__atof_native, _size__atof_native, _stack__atof_native, _pcsp__atof_native},
    {"_decimal_to_f64", _entry__decimal_to_f64, _size__decimal_to_f64, _stack__decimal_to_f64, _pcsp__decimal_to_f64},
    {"_right_shift", _entry__right_shift, _size__right_shift, _stack__right_shift, _pcsp__right_shift},
    {"_left_shift", _entry__left_shift, _size__left_shift, _stack__left_shift, _pcsp__left_shift},
    {"_vsigned", _entry__vsigned, _size__vsigned, _stack__vsigned, _pcsp__vsigned},
    {"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
    {"_vunsigned", _entry__vunsigned, _size__vunsigned, _stack__vunsigned, _pcsp__vunsigned},
}
