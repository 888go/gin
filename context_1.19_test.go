// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build go1.19

package gin类

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
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Set("Content-Type", mw.FormDataContentType())
	c.engine.X最大Multipart内存 = 8 << 20
	f, err := c.X取表单上传文件("file")
	assert.Error(t, err)
	assert.Nil(t, f)
}
