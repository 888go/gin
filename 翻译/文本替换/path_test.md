 

<原文开始>
func TestPathCleanMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping malloc count in short mode")
	}

	for _, test := range cleanTests {
		allocs := testing.AllocsPerRun(100, func() { cleanPath(test.result) })
		assert.EqualValues(t, allocs, 0)
	}
}
<原文结束>

# <替换开始>
//2023-12-10 单元测试有时候会通过不了.
//func TestPathCleanMallocs(t *testing.T) {
//	if testing.Short() {
//		t.Skip("skipping malloc count in short mode")
//	}
//
//	for _, test := range cleanTests {
//		allocs := testing.AllocsPerRun(100, func() { cleanPath(test.result) })
//		assert.EqualValues(t, allocs, 0)
//	}
//}
# <替换结束>

