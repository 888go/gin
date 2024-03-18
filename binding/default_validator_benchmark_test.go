// 版权所有2022 Gin Core团队
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

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
