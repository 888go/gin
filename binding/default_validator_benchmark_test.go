// 版权所有 ? 2022 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package binding

import (
	"errors"
	"strconv"
	"testing"
)

func BenchmarkSliceValidationError(b *testing.B) {
	const size int = 100
	for i := 0; i < b.N; i++ {
		e := make(SliceValidationError, size)
		for j := 0; j < size; j++ {
			e[j] = errors.New(strconv.Itoa(j))
		}
		if len(e.Error()) == 0 {
			b.Errorf("error")
		}
	}
}
