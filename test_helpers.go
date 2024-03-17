// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"net/http"
)

// CreateTestContext为测试目的返回一个新的引擎和上下文

// ff:
// r:
// c:
// w:
func CreateTestContext(w http.ResponseWriter) (c *Context, r *Engine) {
	r = New()
	c = r.allocateContext(0)
	c.reset()
	c.writermem.reset(w)
	return
}

// CreateTestContextOnly返回基于引擎的新上下文，用于测试目的

// ff:
// c:
// r:
// w:
func CreateTestContextOnly(w http.ResponseWriter, r *Engine) (c *Context) {
	c = r.allocateContext(r.maxParams)
	c.reset()
	c.writermem.reset(w)
	return
}
