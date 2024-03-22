// 版权所有 ? 2022 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

//go:build sonic && avx && (linux || windows || darwin) && amd64

package json

import (
	"github.com/bytedance/sonic"
)

var (
	json = sonic.ConfigStd
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
