// Code generated by Makefile, DO NOT EDIT.

/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sse

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/rt`
)

var F_lspace func(sp unsafe.Pointer, nb int, off int) (ret int)

var S_lspace uintptr

//go:nosplit
func lspace(sp *byte, nb int, off int) (ret int) {
    return F_lspace(rt.NoEscape(unsafe.Pointer(sp)), nb, off)
}
