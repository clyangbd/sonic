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

package avx2

import (
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
    `github.com/bytedance/sonic/internal/rt`
)

var F_parse_lazy func(s unsafe.Pointer, p unsafe.Pointer, node unsafe.Pointer, path unsafe.Pointer) (ret int)

var S_parse_lazy uintptr

//go:nosplit
func parse_lazy(s *string, p *int, node *types.Node, path *[]interface{}) (ret int) {
    return F_parse_lazy(rt.NoEscape(unsafe.Pointer(s)), rt.NoEscape(unsafe.Pointer(p)), rt.NoEscape(unsafe.Pointer(node)), rt.NoEscape(unsafe.Pointer(path)))
}
