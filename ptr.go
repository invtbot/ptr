// Copyright (c) 2025 inventor.bot & Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package ptr provides helpers to convert basic types to pointers.
package ptr

// Bool converts a bool to a pointer
func Bool(x bool) *bool {
	return &x
}

// Int8 converts an int8 to a pointer
func Int8(x int8) *int8 {
	return &x
}

// Int16 converts an int16 to a pointer
func Int16(x int16) *int16 {
	return &x
}

// Int32 converts an int32 to a pointer
func Int32(x int32) *int32 {
	return &x
}

// Int64 converts an int64 to a pointer
func Int64(x int64) *int64 {
	return &x
}

// Int converts an int to a pointer
func Int(x int) *int {
	return &x
}

// Uint8 converts a uint8 to a pointer
func Uint8(x uint8) *uint8 {
	return &x
}

// Uint16 converts a uint16 to a pointer
func Uint16(x uint16) *uint16 {
	return &x
}

// Uint32 converts a uint32 to a pointer
func Uint32(x uint32) *uint32 {
	return &x
}

// Uint64 converts a uint64 to a pointer
func Uint64(x uint64) *uint64 {
	return &x
}

// Uint converts a uint to a pointer
func Uint(x uint) *uint {
	return &x
}

// Uintptr converts a uintptr to a pointer
func Uintptr(x uintptr) *uintptr {
	return &x
}

// Float32 converts a float32 to a pointer
func Float32(x float32) *float32 {
	return &x
}

// Float64 converts a float64 to a pointer
func Float64(x float64) *float64 {
	return &x
}

// Complex64 converts a complex64 to a pointer
func Complex64(x complex64) *complex64 {
	return &x
}

// Complex128 converts a complex128 to a pointer
func Complex128(x complex128) *complex128 {
	return &x
}

// String converts a string to a pointer
func String(x string) *string {
	return &x
}

// Rune converts a rune to a pointer
func Rune(x rune) *rune {
	return &x
}

// Byte converts a byte to a pointer
func Byte(x byte) *byte {
	return &x
}
