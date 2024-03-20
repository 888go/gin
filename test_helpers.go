// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。

package gin

import (
	"net/http"
)

// CreateTestContext 返回一个用于测试的新鲜（即初始化的）引擎和上下文

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

// CreateTestContextOnly 根据引擎为测试目的返回一个新的上下文副本

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
