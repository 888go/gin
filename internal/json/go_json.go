// 版权所有 ? 2017 Bo-Yi Wu。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build go_json

package json

import (
	json "github.com/goccy/go-json"
)

var (
	// Marshal 是由 gin/json 包导出的。
	Marshal = json.Marshal
	// Unmarshal 是由 gin/json 包导出的。
	Unmarshal = json.Unmarshal
	// MarshalIndent 是 gin/json 包导出的方法。
	MarshalIndent = json.MarshalIndent
	// NewDecoder 由 gin/json 包导出。
	NewDecoder = json.NewDecoder
	// NewEncoder 是由 gin/json 包导出的。
	NewEncoder = json.NewEncoder
)
