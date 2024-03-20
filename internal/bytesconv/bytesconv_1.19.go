// 版权所有 ? 2020 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

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

// BytesToString 将字节切片转换为字符串，无需进行内存分配。

// ff:
// b:
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
