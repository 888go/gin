// 版权所有 ? 2020 Gin Core Team。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中找到。

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

// 使用以下命令运行测试并显示详细信息：go test -v

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
	letterIdxBits = 6                    // 用6位来表示一个字母索引
	letterIdxMask = 1<<letterIdxBits - 1 // 生成包含1-bits的切片，数量与letterIdxBits相同
	letterIdxMax  = 63 / letterIdxBits   // 符合63位大小的字母索引数量
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
// A src.Int63() 生成63个随机位，足以生成letterIdxMax个字符！
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

// 使用以下命令运行测试并显示详细信息：go test -v -run=none -bench=^BenchmarkBytesConv -benchmem=true

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
