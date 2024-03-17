// 版权所有 ? 2020 Gin Core Team。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build !go1.20

package bytesconv

import (
	"unsafe"
)

// StringToBytes 将字符串转换为字节切片，且无需进行内存分配。

// ff:
// s:
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString 将字节切片转换为字符串，无需分配内存。

// ff:
// b:
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
