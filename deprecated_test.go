// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin/binding"
	"github.com/stretchr/testify/assert"
)

func TestBindWith(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/?foo=bar&bar=foo", bytes.NewBufferString("foo=unused"))

	var obj struct {
		Foo string `form:"foo"`
		Bar string `form:"bar"`
	}
	captureOutput(t, func() {
		assert.NoError(t, c.X弃用BindWith(&obj, binding.Form))
	})
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}
