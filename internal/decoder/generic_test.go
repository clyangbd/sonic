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

package decoder

import (
    `testing`

    `github.com/bytedance/sonic/internal/native/types`
)



//go:nosplit
func decodeValueStub(st *_Stack, sp string, ic int, vp *interface{}, df uint64) (int, types.ParsingError)

func decodeValue(k *_Stack, s string, i int, f uint64) (p int, v interface{}, e types.ParsingError) {
    p, e = decodeValueStub(k, s, i, &v, f)
    return
}

func decodeGeneric(s string, i int, f uint64) (p int, v interface{}, e types.ParsingError) {
    t := newStack()
    p, e = decodeValueStub(t, s, i, &v, f)
    freeStack(t)
    return
}

func BenchmarkGeneric_DecodeGeneric(b *testing.B) {
    t := newStack()
    _, _, _ = decodeValue(t, TwitterJson, 0, 0)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _, _ = decodeValue(t, TwitterJson, 0, 0)
    }
    freeStack(t)
}

func BenchmarkGeneric_Parallel_DecodeGeneric(b *testing.B) {
    _, _, _ = decodeGeneric(TwitterJson, 0, 0)
    b.SetBytes(int64(len(TwitterJson)))
    b.ResetTimer()
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            _, _, _ = decodeGeneric(TwitterJson, 0, 0)
        }
    })
}
