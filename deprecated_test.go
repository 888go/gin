// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

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

	c.Request, _ = http.NewRequest("POST", "/?foo=bar&bar=foo", bytes.NewBufferString("foo=unused"))

	var obj struct {
		Foo string `form:"foo"`
		Bar string `form:"bar"`
	}
	captureOutput(t, func() {
		assert.NoError(t, c.BindWith(&obj, binding.Form))
	})
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}
