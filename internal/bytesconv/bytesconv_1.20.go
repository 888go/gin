// 版权所有 2023 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build go1.20

package bytesconv

import (
	"unsafe"
)

// StringToBytes 将字符串转换为字节切片，且无需进行内存分配。
// 有关更多详细信息，请参见 https://github.com/golang/go/issues/53003#issuecomment-1140276077。
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString 将字节切片转换为字符串，无需进行内存分配。
// 有关更多详细信息，请参见 https://github.com/golang/go/issues/53003#issuecomment-1140276077。
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
