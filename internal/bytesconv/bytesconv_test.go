// 版权所有 ? 2020 Gin 核心团队。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package bytesconv

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"
	"time"
)

var testString = "Albert Einstein: Logic will get you from A to B. Imagination will take you everywhere."
var testBytes = []byte(testString)

func rawBytesToStr(b []byte) string {
	return string(b)
}

func rawStrToBytes(s string) []byte {
	return []byte(s)
}

// go test -v

func TestBytesToString(t *testing.T) {
	data := make([]byte, 1024)
	for i := 0; i < 100; i++ {
		rand.Read(data)
		if rawBytesToStr(data) != BytesToString(data) {
			t.Fatal("don't match")
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 使用6位表示字母索引
	letterIdxMask = 1<<letterIdxBits - 1 // 所有为1的位，数量与letterIdxBits相同
	letterIdxMax  = 63 / letterIdxBits   // 符合63位的字母索引数量
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() 生成63个随机位，足够用于letterIdxMax个字符！
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func TestStringToBytes(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := RandStringBytesMaskImprSrcSB(64)
		if !bytes.Equal(rawStrToBytes(s), StringToBytes(s)) {
			t.Fatal("don't match")
		}
	}
}

// 运行命令：go test -v（详细模式）-run=none（不运行任何正常测试用例）-bench=^BenchmarkBytesConv（仅运行名称以"BenchmarkBytesConv"开头的基准测试）-benchmem=true（在基准测试中包含内存分配统计信息）

func BenchmarkBytesConvBytesToStrRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawBytesToStr(testBytes)
	}
}

func BenchmarkBytesConvBytesToStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BytesToString(testBytes)
	}
}

func BenchmarkBytesConvStrToBytesRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawStrToBytes(testString)
	}
}

func BenchmarkBytesConvStrToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(testString)
	}
}
