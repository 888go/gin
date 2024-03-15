// 版权所有 2017 Bo-Yi Wu。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build !jsoniter && !go_json && !(sonic && avx && (linux || windows || darwin) && amd64)

package json

import (
	"encoding/json"
)

var (
// Marshal 是由 gin/json 包导出的。
	Marshal = json.Marshal
// Unmarshal 是由 gin/json 包导出的。
	Unmarshal = json.Unmarshal
// MarshalIndent 是 gin/json 包导出的方法。
// （该方法用于）将数据按照指定的缩进格式转换为 JSON 格式的字符串。
	MarshalIndent = json.MarshalIndent
// NewDecoder 是 gin/json 包导出的方法。
	NewDecoder = json.NewDecoder
// NewEncoder 由 gin/json 包导出。
	NewEncoder = json.NewEncoder
)
