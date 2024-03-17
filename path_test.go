// 版权所有2013朱利安施密特
// 版权所有
// 基于路径包，版权归the Go Authors所有
// 此源代码的使用受bsd风格的许可证的约束，该许可证可在https://github.com/julienschmidt/httprouter/blob/master/LICENSE上找到

package gin

import (
	"strings"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

type cleanPathTest struct {
	path, result string
}

var cleanTests = []cleanPathTest{
// 已经清洁
	{"/", "/"},
	{"/abc", "/abc"},
	{"/a/b/c", "/a/b/c"},
	{"/abc/", "/abc/"},
	{"/a/b/c/", "/a/b/c/"},

// 缺失的根源
	{"", "/"},
	{"a/", "/a/"},
	{"abc", "/abc"},
	{"abc/def", "/abc/def"},
	{"a/b/c", "/a/b/c"},

// 删除双斜线
	{"//", "/"},
	{"/abc//", "/abc/"},
	{"/abc/def//", "/abc/def/"},
	{"/a/b/c//", "/a/b/c/"},
	{"/abc//def//ghi", "/abc/def/ghi"},
	{"//abc", "/abc"},
	{"///abc", "/abc"},
	{"//abc//", "/abc/"},

// 删除
// 元素
	{".", "/"},
	{"./", "/"},
	{"/abc/./def", "/abc/def"},
	{"/./abc/def", "/abc/def"},
	{"/abc/.", "/abc/"},

// 删除. .元素
	{"..", "/"},
	{"../", "/"},
	{"../../", "/"},
	{"../..", "/"},
	{"../../abc", "/abc"},
	{"/abc/def/ghi/../jkl", "/abc/def/jkl"},
	{"/abc/def/../ghi/../jkl", "/abc/jkl"},
	{"/abc/def/..", "/abc"},
	{"/abc/def/../..", "/"},
	{"/abc/def/../../..", "/"},
	{"/abc/def/../../..", "/"},
	{"/abc/def/../../../ghi/jkl/../../../mno", "/mno"},

// 组合
	{"abc/./../def", "/def"},
	{"abc//./../def", "/def"},
	{"abc/../../././../def", "/def"},
}


// ff:
// t:
func TestPathClean(t *testing.T) {
	for _, test := range cleanTests {
		assert.Equal(t, test.result, cleanPath(test.path))
		assert.Equal(t, test.result, cleanPath(test.result))
	}
}


// ff:
// t:
func TestPathCleanMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping malloc count in short mode")
	}

	for _, test := range cleanTests {
		allocs := testing.AllocsPerRun(100, func() { cleanPath(test.result) })
		assert.EqualValues(t, allocs, 0)
	}
}


// ff:
// b:
func BenchmarkPathClean(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, test := range cleanTests {
			cleanPath(test.path)
		}
	}
}

func genLongPaths() (testPaths []cleanPathTest) {
	for i := 1; i <= 1234; i++ {
		ss := strings.Repeat("a", i)

		correctPath := "/" + ss
		testPaths = append(testPaths, cleanPathTest{
			path:   correctPath,
			result: correctPath,
		}, cleanPathTest{
			path:   ss,
			result: correctPath,
		}, cleanPathTest{
			path:   "//" + ss,
			result: correctPath,
		}, cleanPathTest{
			path:   "/" + ss + "/b/..",
			result: correctPath,
		})
	}
	return
}


// ff:
// t:
func TestPathCleanLong(t *testing.T) {
	cleanTests := genLongPaths()

	for _, test := range cleanTests {
		assert.Equal(t, test.result, cleanPath(test.path))
		assert.Equal(t, test.result, cleanPath(test.result))
	}
}


// ff:
// b:
func BenchmarkPathCleanLong(b *testing.B) {
	cleanTests := genLongPaths()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, test := range cleanTests {
			cleanPath(test.path)
		}
	}
}
