// 版权所有 2021 Gin Core Team。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中找到。

//go:build !go1.19

package gin

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestContextFormFileFailed18(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	defer func(mw *multipart.Writer) {
		err := mw.Close()
		if err != nil {
			assert.Error(t, err)
		}
	}(mw)
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Content-Type", mw.FormDataContentType())
	c.engine.MaxMultipartMemory = 8 << 20
	assert.Panics(t, func() {
		f, err := c.FormFile("file")
		assert.Error(t, err)
		assert.Nil(t, f)
	})
}
