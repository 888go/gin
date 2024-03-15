// 版权所有2022 Gin Core团队
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

//go:build go1.19

package gin

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestContextFormFileFailed19(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	mw.Close()
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "/", nil)
	c.Request.Header.Set("Content-Type", mw.FormDataContentType())
	c.engine.MaxMultipartMemory = 8 << 20
	f, err := c.FormFile("file")
	assert.Error(t, err)
	assert.Nil(t, f)
}
