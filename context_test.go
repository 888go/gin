// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin类

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
	
	"github.com/888go/gin/binding"
	testdata "github.com/888go/gin/testdata/protoexample"
	"github.com/gin-contrib/sse"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

var _ context.Context = (*Context)(nil)

var errTestRender = errors.New("TestRender")

// Unit tests TODO
// func (c *Context) File(filepath string) {
// func (c *Context) Negotiate(code int, config Negotiate) {
// BAD case: func (c *Context) Render(code int, render render.Render, obj ...any) {
// test that information is not leaked when reusing Contexts (using the Pool)

func createMultipartRequest() *http.Request {
	boundary := "--testboundary"
	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)
	defer mw.Close()

	must(mw.SetBoundary(boundary))
	must(mw.WriteField("foo", "bar"))
	must(mw.WriteField("bar", "10"))
	must(mw.WriteField("bar", "foo2"))
	must(mw.WriteField("array", "first"))
	must(mw.WriteField("array", "second"))
	must(mw.WriteField("id", ""))
	must(mw.WriteField("time_local", "31/12/2016 14:55"))
	must(mw.WriteField("time_utc", "31/12/2016 14:55"))
	must(mw.WriteField("time_location", "31/12/2016 14:55"))
	must(mw.WriteField("names[a]", "thinkerou"))
	must(mw.WriteField("names[b]", "tianou"))
	req, err := http.NewRequest("POST", "/", body)
	must(err)
	req.Header.Set("Content-Type", MIMEMultipartPOSTForm+"; boundary="+boundary)
	return req
}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func TestContextFormFile(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	w, err := mw.CreateFormFile("file", "test")
	if assert.NoError(t, err) {
		_, err = w.Write([]byte("test"))
		assert.NoError(t, err)
	}
	mw.Close()
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", buf)
	c.X请求.Header.Set("Content-Type", mw.FormDataContentType())
	f, err := c.X取表单上传文件("file")
	if assert.NoError(t, err) {
		assert.Equal(t, "test", f.Filename)
	}

	assert.NoError(t, c.X保存上传文件(f, "test"))
}

func TestContextMultipartForm(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	assert.NoError(t, mw.WriteField("foo", "bar"))
	w, err := mw.CreateFormFile("file", "test")
	if assert.NoError(t, err) {
		_, err = w.Write([]byte("test"))
		assert.NoError(t, err)
	}
	mw.Close()
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", buf)
	c.X请求.Header.Set("Content-Type", mw.FormDataContentType())
	f, err := c.X取表单multipart对象()
	if assert.NoError(t, err) {
		assert.NotNil(t, f)
	}

	assert.NoError(t, c.X保存上传文件(f.File["file"][0], "test"))
}

func TestSaveUploadedOpenFailed(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	mw.Close()

	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", buf)
	c.X请求.Header.Set("Content-Type", mw.FormDataContentType())

	f := &multipart.FileHeader{
		Filename: "file",
	}
	assert.Error(t, c.X保存上传文件(f, "test"))
}

func TestSaveUploadedCreateFailed(t *testing.T) {
	buf := new(bytes.Buffer)
	mw := multipart.NewWriter(buf)
	w, err := mw.CreateFormFile("file", "test")
	if assert.NoError(t, err) {
		_, err = w.Write([]byte("test"))
		assert.NoError(t, err)
	}
	mw.Close()
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", buf)
	c.X请求.Header.Set("Content-Type", mw.FormDataContentType())
	f, err := c.X取表单上传文件("file")
	if assert.NoError(t, err) {
		assert.Equal(t, "test", f.Filename)
	}

	assert.Error(t, c.X保存上传文件(f, "/"))
}

func TestContextReset(t *testing.T) {
	router := X创建()
	c := router.allocateContext(0)
	assert.Equal(t, c.engine, router)

	c.index = 2
	c.Writer = &responseWriter{ResponseWriter: httptest.NewRecorder()}
	c.X参数 = Params{Param{}}
	c.X错误(errors.New("test")) //nolint: errcheck
	c.X设置值("foo", "bar")
	c.reset()

	assert.False(t, c.X是否已停止())
	assert.Nil(t, c.X上下文设置值Map)
	assert.Nil(t, c.Accepted)
	assert.Len(t, c.X错误s, 0)
	assert.Empty(t, c.X错误s.Errors())
	assert.Empty(t, c.X错误s.ByType(ErrorTypeAny))
	assert.Len(t, c.X参数, 0)
	assert.EqualValues(t, c.index, -1)
	assert.Equal(t, c.Writer.(*responseWriter), &c.writermem)
}

func TestContextHandlers(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	assert.Nil(t, c.handlers)
	assert.Nil(t, c.handlers.X取最后一个处理函数())

	c.handlers = HandlersChain{}
	assert.NotNil(t, c.handlers)
	assert.Nil(t, c.handlers.X取最后一个处理函数())

	f := func(c *Context) {}
	g := func(c *Context) {}

	c.handlers = HandlersChain{f}
	compareFunc(t, f, c.handlers.X取最后一个处理函数())

	c.handlers = HandlersChain{f, g}
	compareFunc(t, g, c.handlers.X取最后一个处理函数())
}

// TestContextSetGet tests that a parameter is set correctly on the
// current context and can be retrieved using Get.
func TestContextSetGet(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("foo", "bar")

	value, err := c.X取值("foo")
	assert.Equal(t, "bar", value)
	assert.True(t, err)

	value, err = c.X取值("foo2")
	assert.Nil(t, value)
	assert.False(t, err)

	assert.Equal(t, "bar", c.X取值PANI("foo"))
	assert.Panics(t, func() { c.X取值PANI("no_exist") })
}

func TestContextSetGetValues(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("string", "this is a string")
	c.X设置值("int32", int32(-42))
	c.X设置值("int64", int64(42424242424242))
	c.X设置值("uint64", uint64(42))
	c.X设置值("float32", float32(4.2))
	c.X设置值("float64", 4.2)
	var a any = 1
	c.X设置值("intInterface", a)

	assert.Exactly(t, c.X取值PANI("string").(string), "this is a string")
	assert.Exactly(t, c.X取值PANI("int32").(int32), int32(-42))
	assert.Exactly(t, c.X取值PANI("int64").(int64), int64(42424242424242))
	assert.Exactly(t, c.X取值PANI("uint64").(uint64), uint64(42))
	assert.Exactly(t, c.X取值PANI("float32").(float32), float32(4.2))
	assert.Exactly(t, c.X取值PANI("float64").(float64), 4.2)
	assert.Exactly(t, c.X取值PANI("intInterface").(int), 1)
}

func TestContextGetString(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("string", "this is a string")
	assert.Equal(t, "this is a string", c.X取文本值("string"))
}

func TestContextSetGetBool(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("bool", true)
	assert.True(t, c.X取布尔值("bool"))
}

func TestContextGetInt(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("int", 1)
	assert.Equal(t, 1, c.X取整数值("int"))
}

func TestContextGetInt64(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("int64", int64(42424242424242))
	assert.Equal(t, int64(42424242424242), c.X取整数64位值("int64"))
}

func TestContextGetUint(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("uint", uint(1))
	assert.Equal(t, uint(1), c.X取正整数值("uint"))
}

func TestContextGetUint64(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("uint64", uint64(18446744073709551615))
	assert.Equal(t, uint64(18446744073709551615), c.X取正整数64位值("uint64"))
}

func TestContextGetFloat64(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("float64", 4.2)
	assert.Equal(t, 4.2, c.X取小数64位值("float64"))
}

func TestContextGetTime(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	t1, _ := time.Parse("1/2/2006 15:04:05", "01/01/2017 12:00:00")
	c.X设置值("time", t1)
	assert.Equal(t, t1, c.X取时间值("time"))
}

func TestContextGetDuration(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("duration", time.Second)
	assert.Equal(t, time.Second, c.X取时长值("duration"))
}

func TestContextGetStringSlice(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置值("slice", []string{"foo"})
	assert.Equal(t, []string{"foo"}, c.X取数组值("slice"))
}

func TestContextGetStringMap(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	m := make(map[string]any)
	m["foo"] = 1
	c.X设置值("map", m)

	assert.Equal(t, m, c.X取Map值("map"))
	assert.Equal(t, 1, c.X取Map值("map")["foo"])
}

func TestContextGetStringMapString(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	m := make(map[string]string)
	m["foo"] = "bar"
	c.X设置值("map", m)

	assert.Equal(t, m, c.X取文本Map值("map"))
	assert.Equal(t, "bar", c.X取文本Map值("map")["foo"])
}

func TestContextGetStringMapStringSlice(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	m := make(map[string][]string)
	m["foo"] = []string{"foo"}
	c.X设置值("map", m)

	assert.Equal(t, m, c.X取数组Map值("map"))
	assert.Equal(t, []string{"foo"}, c.X取数组Map值("map")["foo"])
}

func TestContextCopy(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.index = 2
	c.X请求, _ = http.NewRequest("POST", "/hola", nil)
	c.handlers = HandlersChain{func(c *Context) {}}
	c.X参数 = Params{Param{Key: "foo", Value: "bar"}}
	c.X设置值("foo", "bar")

	cp := c.X取副本()
	assert.Nil(t, cp.handlers)
	assert.Nil(t, cp.writermem.ResponseWriter)
	assert.Equal(t, &cp.writermem, cp.Writer.(*responseWriter))
	assert.Equal(t, cp.X请求, c.X请求)
	assert.Equal(t, cp.index, abortIndex)
	assert.Equal(t, cp.X上下文设置值Map, c.X上下文设置值Map)
	assert.Equal(t, cp.engine, c.engine)
	assert.Equal(t, cp.X参数, c.X参数)
	cp.X设置值("foo", "notBar")
	assert.False(t, cp.X上下文设置值Map["foo"] == c.X上下文设置值Map["foo"])
}

func TestContextHandlerName(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.handlers = HandlersChain{func(c *Context) {}, handlerNameTest}

	assert.Regexp(t, "^(.*/vendor/)?github.com/888go/gin.handlerNameTest$", c.X取主处理程序名称())
}

func TestContextHandlerNames(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.handlers = HandlersChain{func(c *Context) {}, handlerNameTest, func(c *Context) {}, handlerNameTest2}

	names := c.X取处理程序数组()

	assert.True(t, len(names) == 4)
	for _, name := range names {
		assert.Regexp(t, `^(.*/vendor/)?(github\.com/888go/gin\.){1}(TestContextHandlerNames\.func.*){0,1}(handlerNameTest.*){0,1}`, name)
	}
}

func handlerNameTest(c *Context) {
}

func handlerNameTest2(c *Context) {
}

var handlerTest HandlerFunc = func(c *Context) {
}

func TestContextHandler(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.handlers = HandlersChain{func(c *Context) {}, handlerTest}

	assert.Equal(t, reflect.ValueOf(handlerTest).Pointer(), reflect.ValueOf(c.X取主处理程序()).Pointer())
}

func TestContextQuery(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("GET", "http://example.com/?foo=bar&page=10&id=", nil)

	value, ok := c.X取URL参数值2("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", value)
	assert.Equal(t, "bar", c.X取URL参数值并带默认("foo", "none"))
	assert.Equal(t, "bar", c.X取URL参数值("foo"))

	value, ok = c.X取URL参数值2("page")
	assert.True(t, ok)
	assert.Equal(t, "10", value)
	assert.Equal(t, "10", c.X取URL参数值并带默认("page", "0"))
	assert.Equal(t, "10", c.X取URL参数值("page"))

	value, ok = c.X取URL参数值2("id")
	assert.True(t, ok)
	assert.Empty(t, value)
	assert.Empty(t, c.X取URL参数值并带默认("id", "nada"))
	assert.Empty(t, c.X取URL参数值("id"))

	value, ok = c.X取URL参数值2("NoKey")
	assert.False(t, ok)
	assert.Empty(t, value)
	assert.Equal(t, "nada", c.X取URL参数值并带默认("NoKey", "nada"))
	assert.Empty(t, c.X取URL参数值("NoKey"))

	// postform should not mess
	value, ok = c.X取表单参数值2("page")
	assert.False(t, ok)
	assert.Empty(t, value)
	assert.Empty(t, c.X取表单参数值("foo"))
}

func TestContextDefaultQueryOnEmptyRequest(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder()) // here c.Request == nil
	assert.NotPanics(t, func() {
		value, ok := c.X取URL参数值2("NoKey")
		assert.False(t, ok)
		assert.Empty(t, value)
	})
	assert.NotPanics(t, func() {
		assert.Equal(t, "nada", c.X取URL参数值并带默认("NoKey", "nada"))
	})
	assert.NotPanics(t, func() {
		assert.Empty(t, c.X取URL参数值("NoKey"))
	})
}

func TestContextQueryAndPostForm(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString("foo=bar&page=11&both=&foo=second")
	c.X请求, _ = http.NewRequest("POST",
		"/?both=GET&id=main&id=omit&array[]=first&array[]=second&ids[a]=hi&ids[b]=3.14", body)
	c.X请求.Header.Add("Content-Type", MIMEPOSTForm)

	assert.Equal(t, "bar", c.X取表单参数值并带默认("foo", "none"))
	assert.Equal(t, "bar", c.X取表单参数值("foo"))
	assert.Empty(t, c.X取URL参数值("foo"))

	value, ok := c.X取表单参数值2("page")
	assert.True(t, ok)
	assert.Equal(t, "11", value)
	assert.Equal(t, "11", c.X取表单参数值并带默认("page", "0"))
	assert.Equal(t, "11", c.X取表单参数值("page"))
	assert.Empty(t, c.X取URL参数值("page"))

	value, ok = c.X取表单参数值2("both")
	assert.True(t, ok)
	assert.Empty(t, value)
	assert.Empty(t, c.X取表单参数值("both"))
	assert.Empty(t, c.X取表单参数值并带默认("both", "nothing"))
	assert.Equal(t, "GET", c.X取URL参数值("both"), "GET")

	value, ok = c.X取URL参数值2("id")
	assert.True(t, ok)
	assert.Equal(t, "main", value)
	assert.Equal(t, "000", c.X取表单参数值并带默认("id", "000"))
	assert.Equal(t, "main", c.X取URL参数值("id"))
	assert.Empty(t, c.X取表单参数值("id"))

	value, ok = c.X取URL参数值2("NoKey")
	assert.False(t, ok)
	assert.Empty(t, value)
	value, ok = c.X取表单参数值2("NoKey")
	assert.False(t, ok)
	assert.Empty(t, value)
	assert.Equal(t, "nada", c.X取表单参数值并带默认("NoKey", "nada"))
	assert.Equal(t, "nothing", c.X取URL参数值并带默认("NoKey", "nothing"))
	assert.Empty(t, c.X取表单参数值("NoKey"))
	assert.Empty(t, c.X取URL参数值("NoKey"))

	var obj struct {
		Foo   string   `form:"foo"`
		ID    string   `form:"id"`
		Page  int      `form:"page"`
		Both  string   `form:"both"`
		Array []string `form:"array[]"`
	}
	assert.NoError(t, c.X取参数到指针PANI(&obj))
	assert.Equal(t, "bar", obj.Foo, "bar")
	assert.Equal(t, "main", obj.ID, "main")
	assert.Equal(t, 11, obj.Page, 11)
	assert.Empty(t, obj.Both)
	assert.Equal(t, []string{"first", "second"}, obj.Array)

	values, ok := c.X取URL参数数组值2("array[]")
	assert.True(t, ok)
	assert.Equal(t, "first", values[0])
	assert.Equal(t, "second", values[1])

	values = c.X取URL参数数组值("array[]")
	assert.Equal(t, "first", values[0])
	assert.Equal(t, "second", values[1])

	values = c.X取URL参数数组值("nokey")
	assert.Equal(t, 0, len(values))

	values = c.X取URL参数数组值("both")
	assert.Equal(t, 1, len(values))
	assert.Equal(t, "GET", values[0])

	dicts, ok := c.X取URL参数Map值2("ids")
	assert.True(t, ok)
	assert.Equal(t, "hi", dicts["a"])
	assert.Equal(t, "3.14", dicts["b"])

	dicts, ok = c.X取URL参数Map值2("nokey")
	assert.False(t, ok)
	assert.Equal(t, 0, len(dicts))

	dicts, ok = c.X取URL参数Map值2("both")
	assert.False(t, ok)
	assert.Equal(t, 0, len(dicts))

	dicts, ok = c.X取URL参数Map值2("array")
	assert.False(t, ok)
	assert.Equal(t, 0, len(dicts))

	dicts = c.X取URL参数Map值("ids")
	assert.Equal(t, "hi", dicts["a"])
	assert.Equal(t, "3.14", dicts["b"])

	dicts = c.X取URL参数Map值("nokey")
	assert.Equal(t, 0, len(dicts))
}

func TestContextPostFormMultipart(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求 = createMultipartRequest()

	var obj struct {
		Foo          string    `form:"foo"`
		Bar          string    `form:"bar"`
		BarAsInt     int       `form:"bar"`
		Array        []string  `form:"array"`
		ID           string    `form:"id"`
		TimeLocal    time.Time `form:"time_local" time_format:"02/01/2006 15:04"`
		TimeUTC      time.Time `form:"time_utc" time_format:"02/01/2006 15:04" time_utc:"1"`
		TimeLocation time.Time `form:"time_location" time_format:"02/01/2006 15:04" time_location:"Asia/Tokyo"`
		BlankTime    time.Time `form:"blank_time" time_format:"02/01/2006 15:04"`
	}
	assert.NoError(t, c.X取参数到指针PANI(&obj))
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "10", obj.Bar)
	assert.Equal(t, 10, obj.BarAsInt)
	assert.Equal(t, []string{"first", "second"}, obj.Array)
	assert.Empty(t, obj.ID)
	assert.Equal(t, "31/12/2016 14:55", obj.TimeLocal.Format("02/01/2006 15:04"))
	assert.Equal(t, time.Local, obj.TimeLocal.Location())
	assert.Equal(t, "31/12/2016 14:55", obj.TimeUTC.Format("02/01/2006 15:04"))
	assert.Equal(t, time.UTC, obj.TimeUTC.Location())
	loc, _ := time.LoadLocation("Asia/Tokyo")
	assert.Equal(t, "31/12/2016 14:55", obj.TimeLocation.Format("02/01/2006 15:04"))
	assert.Equal(t, loc, obj.TimeLocation.Location())
	assert.True(t, obj.BlankTime.IsZero())

	value, ok := c.X取URL参数值2("foo")
	assert.False(t, ok)
	assert.Empty(t, value)
	assert.Empty(t, c.X取URL参数值("bar"))
	assert.Equal(t, "nothing", c.X取URL参数值并带默认("id", "nothing"))

	value, ok = c.X取表单参数值2("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", value)
	assert.Equal(t, "bar", c.X取表单参数值("foo"))

	value, ok = c.X取表单参数值2("array")
	assert.True(t, ok)
	assert.Equal(t, "first", value)
	assert.Equal(t, "first", c.X取表单参数值("array"))

	assert.Equal(t, "10", c.X取表单参数值并带默认("bar", "nothing"))

	value, ok = c.X取表单参数值2("id")
	assert.True(t, ok)
	assert.Empty(t, value)
	assert.Empty(t, c.X取表单参数值("id"))
	assert.Empty(t, c.X取表单参数值并带默认("id", "nothing"))

	value, ok = c.X取表单参数值2("nokey")
	assert.False(t, ok)
	assert.Empty(t, value)
	assert.Equal(t, "nothing", c.X取表单参数值并带默认("nokey", "nothing"))

	values, ok := c.X取参数数组值("array")
	assert.True(t, ok)
	assert.Equal(t, "first", values[0])
	assert.Equal(t, "second", values[1])

	values = c.X取表单参数数组值("array")
	assert.Equal(t, "first", values[0])
	assert.Equal(t, "second", values[1])

	values = c.X取表单参数数组值("nokey")
	assert.Equal(t, 0, len(values))

	values = c.X取表单参数数组值("foo")
	assert.Equal(t, 1, len(values))
	assert.Equal(t, "bar", values[0])

	dicts, ok := c.X取参数Map值("names")
	assert.True(t, ok)
	assert.Equal(t, "thinkerou", dicts["a"])
	assert.Equal(t, "tianou", dicts["b"])

	dicts, ok = c.X取参数Map值("nokey")
	assert.False(t, ok)
	assert.Equal(t, 0, len(dicts))

	dicts = c.X取表单参数Map值("names")
	assert.Equal(t, "thinkerou", dicts["a"])
	assert.Equal(t, "tianou", dicts["b"])

	dicts = c.X取表单参数Map值("nokey")
	assert.Equal(t, 0, len(dicts))
}

func TestContextSetCookie(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置cookie跨站(http.SameSiteLaxMode)
	c.X设置cookie值("user", "gin", 1, "/", "localhost", true, true)
	assert.Equal(t, "user=gin; Path=/; Domain=localhost; Max-Age=1; HttpOnly; Secure; SameSite=Lax", c.Writer.Header().Get("Set-Cookie"))
}

func TestContextSetCookiePathEmpty(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置cookie跨站(http.SameSiteLaxMode)
	c.X设置cookie值("user", "gin", 1, "", "localhost", true, true)
	assert.Equal(t, "user=gin; Path=/; Domain=localhost; Max-Age=1; HttpOnly; Secure; SameSite=Lax", c.Writer.Header().Get("Set-Cookie"))
}

func TestContextGetCookie(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("GET", "/get", nil)
	c.X请求.Header.Set("Cookie", "user=gin")
	cookie, _ := c.X取cookie值("user")
	assert.Equal(t, "gin", cookie)

	_, err := c.X取cookie值("nokey")
	assert.Error(t, err)
}

func TestContextBodyAllowedForStatus(t *testing.T) {
	assert.False(t, false, bodyAllowedForStatus(http.StatusProcessing))
	assert.False(t, false, bodyAllowedForStatus(http.StatusNoContent))
	assert.False(t, false, bodyAllowedForStatus(http.StatusNotModified))
	assert.True(t, true, bodyAllowedForStatus(http.StatusInternalServerError))
}

type TestRender struct{}

func (*TestRender) Render(http.ResponseWriter) error {
	return errTestRender
}

func (*TestRender) WriteContentType(http.ResponseWriter) {}

func TestContextRenderIfErr(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.Render底层方法(http.StatusOK, &TestRender{})

	assert.Equal(t, errorMsgs{&Error{Err: errTestRender, Type: 1}}, c.X错误s)
}

// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
// and special HTML characters are escaped
func TestContextRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出JSON(http.StatusCreated, H{"foo": "bar", "html": "<b>"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"foo\":\"bar\",\"html\":\"\\u003cb\\u003e\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that the response is serialized as JSONP
// and Content-Type is set to application/javascript
func TestContextRenderJSONP(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("GET", "http://example.com/?callback=x", nil)

	c.X输出JSONP(http.StatusCreated, H{"foo": "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "x({\"foo\":\"bar\"});", w.Body.String())
	assert.Equal(t, "application/javascript; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that the response is serialized as JSONP
// and Content-Type is set to application/json
func TestContextRenderJSONPWithoutCallback(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("GET", "http://example.com", nil)

	c.X输出JSONP(http.StatusCreated, H{"foo": "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"foo\":\"bar\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no JSON is rendered if code is 204
func TestContextRenderNoContentJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出JSON(http.StatusNoContent, H{"foo": "bar"})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that the response is serialized as JSON
// we change the content-type before
func TestContextRenderAPIJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X设置响应协议头值("Content-Type", "application/vnd.api+json")
	c.X输出JSON(http.StatusCreated, H{"foo": "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"foo\":\"bar\"}", w.Body.String())
	assert.Equal(t, "application/vnd.api+json", w.Header().Get("Content-Type"))
}

// Tests that no Custom JSON is rendered if code is 204
func TestContextRenderNoContentAPIJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X设置响应协议头值("Content-Type", "application/vnd.api+json")
	c.X输出JSON(http.StatusNoContent, H{"foo": "bar"})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, w.Header().Get("Content-Type"), "application/vnd.api+json")
}

// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
func TestContextRenderIndentedJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出JSON并美化(http.StatusCreated, H{"foo": "bar", "bar": "foo", "nested": H{"foo": "bar"}})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\n    \"bar\": \"foo\",\n    \"foo\": \"bar\",\n    \"nested\": {\n        \"foo\": \"bar\"\n    }\n}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no Custom JSON is rendered if code is 204
func TestContextRenderNoContentIndentedJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出JSON并美化(http.StatusNoContent, H{"foo": "bar", "bar": "foo", "nested": H{"foo": "bar"}})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that the response is serialized as Secure JSON
// and Content-Type is set to application/json
func TestContextRenderSecureJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, router := CreateTestContext(w)

	router.X设置Json防劫持前缀("&&&START&&&")
	c.X输出JSON并防劫持(http.StatusCreated, []string{"foo", "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "&&&START&&&[\"foo\",\"bar\"]", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no Custom JSON is rendered if code is 204
func TestContextRenderNoContentSecureJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出JSON并防劫持(http.StatusNoContent, []string{"foo", "bar"})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextRenderNoContentAsciiJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出JSON并按ASCII(http.StatusNoContent, []string{"lang", "Go语言"})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}

// Tests that the response is serialized as JSON
// and Content-Type is set to application/json
// and special HTML characters are preserved
func TestContextRenderPureJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X输出JSON并按原文(http.StatusCreated, H{"foo": "bar", "html": "<b>"})
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"foo\":\"bar\",\"html\":\"<b>\"}\n", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that the response executes the templates
// and responds with Content-Type set to text/html
func TestContextRenderHTML(t *testing.T) {
	w := httptest.NewRecorder()
	c, router := CreateTestContext(w)

	templ := template.Must(template.New("t").Parse(`Hello {{.name}}`))
	router.X设置Template模板(templ)

	c.X输出html模板(http.StatusCreated, "t", H{"name": "alexandernyquist"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Hello alexandernyquist", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextRenderHTML2(t *testing.T) {
	w := httptest.NewRecorder()
	c, router := CreateTestContext(w)

	// print debug warning log when Engine.trees > 0
	router.addRoute("GET", "/", HandlersChain{func(_ *Context) {}})
	assert.Len(t, router.trees, 1)

	templ := template.Must(template.New("t").Parse(`Hello {{.name}}`))
	re := captureOutput(t, func() {
		X设置运行模式(X常量_运行模式_调试)
		router.X设置Template模板(templ)
		X设置运行模式(X常量_运行模式_测试)
	})

	assert.Equal(t, "[GIN-debug] [WARNING] Since SetHTMLTemplate() is NOT thread-safe. It should only be called\nat initialization. ie. before any route is registered or the router is listening in a socket:\n\n\trouter := gin.Default()\n\trouter.SetHTMLTemplate(template) // << good place\n\n", re)

	c.X输出html模板(http.StatusCreated, "t", H{"name": "alexandernyquist"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Hello alexandernyquist", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no HTML is rendered if code is 204
func TestContextRenderNoContentHTML(t *testing.T) {
	w := httptest.NewRecorder()
	c, router := CreateTestContext(w)
	templ := template.Must(template.New("t").Parse(`Hello {{.name}}`))
	router.X设置Template模板(templ)

	c.X输出html模板(http.StatusNoContent, "t", H{"name": "alexandernyquist"})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestContextXML tests that the response is serialized as XML
// and Content-Type is set to application/xml
func TestContextRenderXML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出XML(http.StatusCreated, H{"foo": "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "<map><foo>bar</foo></map>", w.Body.String())
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no XML is rendered if code is 204
func TestContextRenderNoContentXML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出XML(http.StatusNoContent, H{"foo": "bar"})

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestContextString tests that the response is returned
// with Content-Type set to text/plain
func TestContextRenderString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出文本(http.StatusCreated, "test %s %d", "string", 2)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "test string 2", w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no String is rendered if code is 204
func TestContextRenderNoContentString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出文本(http.StatusNoContent, "test %s %d", "string", 2)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestContextString tests that the response is returned
// with Content-Type set to text/html
func TestContextRenderHTMLString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X设置响应协议头值("Content-Type", "text/html; charset=utf-8")
	c.X输出文本(http.StatusCreated, "<html>%s %d</html>", "string", 3)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "<html>string 3</html>", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

// Tests that no HTML String is rendered if code is 204
func TestContextRenderNoContentHTMLString(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X设置响应协议头值("Content-Type", "text/html; charset=utf-8")
	c.X输出文本(http.StatusNoContent, "<html>%s %d</html>", "string", 3)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestContextData tests that the response can be written from `bytestring`
// with specified MIME type
func TestContextRenderData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出字节集(http.StatusCreated, "text/csv", []byte(`foo,bar`))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "foo,bar", w.Body.String())
	assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
}

// Tests that no Custom Data is rendered if code is 204
func TestContextRenderNoContentData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出字节集(http.StatusNoContent, "text/csv", []byte(`foo,bar`))

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
	assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
}

func TestContextRenderSSE(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.SSEvent("float", 1.5)
	c.Render底层方法(-1, sse.Event{
		Id:   "123",
		Data: "text",
	})
	c.SSEvent("chat", H{
		"foo": "bar",
		"bar": "foo",
	})

	assert.Equal(t, strings.Replace(w.Body.String(), " ", "", -1), strings.Replace("event:float\ndata:1.5\n\nid:123\ndata:text\n\nevent:chat\ndata:{\"bar\":\"foo\",\"foo\":\"bar\"}\n\n", " ", "", -1))
}

func TestContextRenderFile(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("GET", "/", nil)
	c.X下载文件("./gin.go")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "func X创建() *Engine {") //th:assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
	// Content-Type='text/plain; charset=utf-8' when go version <= 1.16,
	// else, Content-Type='text/x-go; charset=utf-8'
	assert.NotEqual(t, "", w.Header().Get("Content-Type"))
}

func TestContextRenderFileFromFS(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("GET", "/some/path", nil)
	c.X下载文件FS("./gin.go", Dir(".", false))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "func X创建() *Engine {") //th:assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
	// Content-Type='text/plain; charset=utf-8' when go version <= 1.16,
	// else, Content-Type='text/x-go; charset=utf-8'
	assert.NotEqual(t, "", w.Header().Get("Content-Type"))
	assert.Equal(t, "/some/path", c.X请求.URL.Path)
}

func TestContextRenderAttachment(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	newFilename := "new_filename.go"

	c.X请求, _ = http.NewRequest("GET", "/", nil)
	c.X下载文件并带文件名("./gin.go", newFilename)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "func X创建() *Engine {") //th:assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
	assert.Equal(t, fmt.Sprintf("attachment; filename=\"%s\"", newFilename), w.Header().Get("Content-Disposition"))
}

func TestContextRenderAndEscapeAttachment(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	maliciousFilename := "tampering_field.sh\"; \\\"; dummy=.go"
	actualEscapedResponseFilename := "tampering_field.sh\\\"; \\\\\\\"; dummy=.go"

	c.X请求, _ = http.NewRequest("GET", "/", nil)
	c.X下载文件并带文件名("./gin.go", maliciousFilename)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "func X创建() *Engine {") //th:assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
	assert.Equal(t, fmt.Sprintf("attachment; filename=\"%s\"", actualEscapedResponseFilename), w.Header().Get("Content-Disposition"))
}

func TestContextRenderUTF8Attachment(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	newFilename := "new🧡_filename.go"

	c.X请求, _ = http.NewRequest("GET", "/", nil)
	c.X下载文件并带文件名("./gin.go", newFilename)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "func X创建() *Engine {") //th:assert.Contains(t, w.Body.String(), "func X创建() *Engine {")
	assert.Equal(t, `attachment; filename*=UTF-8''`+url.QueryEscape(newFilename), w.Header().Get("Content-Disposition"))
}

// TestContextRenderYAML tests that the response is serialized as YAML
// and Content-Type is set to application/x-yaml
func TestContextRenderYAML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出YAML(http.StatusCreated, H{"foo": "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "foo: bar\n", w.Body.String())
	assert.Equal(t, "application/x-yaml; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestContextRenderTOML tests that the response is serialized as TOML
// and Content-Type is set to application/toml
func TestContextRenderTOML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X输出TOML(http.StatusCreated, H{"foo": "bar"})

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "foo = 'bar'\n", w.Body.String())
	assert.Equal(t, "application/toml; charset=utf-8", w.Header().Get("Content-Type"))
}

// TestContextRenderProtoBuf tests that the response is serialized as ProtoBuf
// and Content-Type is set to application/x-protobuf
// and we just use the example protobuf to check if the response is correct
func TestContextRenderProtoBuf(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	reps := []int64{int64(1), int64(2)}
	label := "test"
	data := &testdata.Test{
		Label: &label,
		Reps:  reps,
	}

	c.X输出ProtoBuf(http.StatusCreated, data)

	protoData, err := proto.Marshal(data)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, string(protoData), w.Body.String())
	assert.Equal(t, "application/x-protobuf", w.Header().Get("Content-Type"))
}

func TestContextHeaders(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X设置响应协议头值("Content-Type", "text/plain")
	c.X设置响应协议头值("X-Custom", "value")

	assert.Equal(t, "text/plain", c.Writer.Header().Get("Content-Type"))
	assert.Equal(t, "value", c.Writer.Header().Get("X-Custom"))

	c.X设置响应协议头值("Content-Type", "text/html")
	c.X设置响应协议头值("X-Custom", "")

	assert.Equal(t, "text/html", c.Writer.Header().Get("Content-Type"))
	_, exist := c.Writer.Header()["X-Custom"]
	assert.False(t, exist)
}

// TODO
func TestContextRenderRedirectWithRelativePath(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "http://example.com", nil)
	assert.Panics(t, func() { c.X重定向(299, "/new_path") })
	assert.Panics(t, func() { c.X重定向(309, "/new_path") })

	c.X重定向(http.StatusMovedPermanently, "/path")
	c.Writer.WriteHeaderNow()
	assert.Equal(t, http.StatusMovedPermanently, w.Code)
	assert.Equal(t, "/path", w.Header().Get("Location"))
}

func TestContextRenderRedirectWithAbsolutePath(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "http://example.com", nil)
	c.X重定向(http.StatusFound, "http://google.com")
	c.Writer.WriteHeaderNow()

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "http://google.com", w.Header().Get("Location"))
}

func TestContextRenderRedirectWith201(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "http://example.com", nil)
	c.X重定向(http.StatusCreated, "/resource")
	c.Writer.WriteHeaderNow()

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "/resource", w.Header().Get("Location"))
}

func TestContextRenderRedirectAll(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "http://example.com", nil)
	assert.Panics(t, func() { c.X重定向(http.StatusOK, "/resource") })
	assert.Panics(t, func() { c.X重定向(http.StatusAccepted, "/resource") })
	assert.Panics(t, func() { c.X重定向(299, "/resource") })
	assert.Panics(t, func() { c.X重定向(309, "/resource") })
	assert.NotPanics(t, func() { c.X重定向(http.StatusMultipleChoices, "/resource") })
	assert.NotPanics(t, func() { c.X重定向(http.StatusPermanentRedirect, "/resource") })
}

func TestContextNegotiationWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("POST", "", nil)

	c.Negotiate底层方法(http.StatusOK, Negotiate{
		Offered: []string{MIMEJSON, MIMEXML, MIMEYAML},
		Data:    H{"foo": "bar"},
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"foo\":\"bar\"}", w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextNegotiationWithXML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("POST", "", nil)

	c.Negotiate底层方法(http.StatusOK, Negotiate{
		Offered: []string{MIMEXML, MIMEJSON, MIMEYAML},
		Data:    H{"foo": "bar"},
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "<map><foo>bar</foo></map>", w.Body.String())
	assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextNegotiationWithYAML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("POST", "", nil)

	c.Negotiate底层方法(http.StatusOK, Negotiate{
		Offered: []string{MIMEYAML, MIMEXML, MIMEJSON, MIMETOML},
		Data:    H{"foo": "bar"},
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "foo: bar\n", w.Body.String())
	assert.Equal(t, "application/x-yaml; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextNegotiationWithTOML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("POST", "", nil)

	c.Negotiate底层方法(http.StatusOK, Negotiate{
		Offered: []string{MIMETOML, MIMEXML, MIMEJSON, MIMEYAML},
		Data:    H{"foo": "bar"},
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "foo = 'bar'\n", w.Body.String())
	assert.Equal(t, "application/toml; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextNegotiationWithHTML(t *testing.T) {
	w := httptest.NewRecorder()
	c, router := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("POST", "", nil)
	templ := template.Must(template.New("t").Parse(`Hello {{.name}}`))
	router.X设置Template模板(templ)

	c.Negotiate底层方法(http.StatusOK, Negotiate{
		Offered:  []string{MIMEHTML},
		Data:     H{"name": "gin"},
		HTMLName: "t",
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello gin", w.Body.String())
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContextNegotiationNotSupport(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.X请求, _ = http.NewRequest("POST", "", nil)

	c.Negotiate底层方法(http.StatusOK, Negotiate{
		Offered: []string{MIMEPOSTForm},
	})

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
	assert.Equal(t, c.index, abortIndex)
	assert.True(t, c.X是否已停止())
}

func TestContextNegotiationFormat(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "", nil)

	assert.Panics(t, func() { c.NegotiateFormat底层方法() })
	assert.Equal(t, MIMEJSON, c.NegotiateFormat底层方法(MIMEJSON, MIMEXML))
	assert.Equal(t, MIMEHTML, c.NegotiateFormat底层方法(MIMEHTML, MIMEJSON))
}

func TestContextNegotiationFormatWithAccept(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9;q=0.8")

	assert.Equal(t, MIMEXML, c.NegotiateFormat底层方法(MIMEJSON, MIMEXML))
	assert.Equal(t, MIMEHTML, c.NegotiateFormat底层方法(MIMEXML, MIMEHTML))
	assert.Empty(t, c.NegotiateFormat底层方法(MIMEJSON))
}

func TestContextNegotiationFormatWithWildcardAccept(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("Accept", "*/*")

	assert.Equal(t, c.NegotiateFormat底层方法("*/*"), "*/*")
	assert.Equal(t, c.NegotiateFormat底层方法("text/*"), "text/*")
	assert.Equal(t, c.NegotiateFormat底层方法("application/*"), "application/*")
	assert.Equal(t, c.NegotiateFormat底层方法(MIMEJSON), MIMEJSON)
	assert.Equal(t, c.NegotiateFormat底层方法(MIMEXML), MIMEXML)
	assert.Equal(t, c.NegotiateFormat底层方法(MIMEHTML), MIMEHTML)

	c, _ = CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("Accept", "text/*")

	assert.Equal(t, c.NegotiateFormat底层方法("*/*"), "*/*")
	assert.Equal(t, c.NegotiateFormat底层方法("text/*"), "text/*")
	assert.Equal(t, c.NegotiateFormat底层方法("application/*"), "")
	assert.Equal(t, c.NegotiateFormat底层方法(MIMEJSON), "")
	assert.Equal(t, c.NegotiateFormat底层方法(MIMEXML), "")
	assert.Equal(t, c.NegotiateFormat底层方法(MIMEHTML), MIMEHTML)
}

func TestContextNegotiationFormatCustom(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9;q=0.8")

	c.Accepted = nil
	c.SetAccepted底层方法(MIMEJSON, MIMEXML)

	assert.Equal(t, MIMEJSON, c.NegotiateFormat底层方法(MIMEJSON, MIMEXML))
	assert.Equal(t, MIMEXML, c.NegotiateFormat底层方法(MIMEXML, MIMEHTML))
	assert.Equal(t, MIMEJSON, c.NegotiateFormat底层方法(MIMEJSON))
}

func TestContextNegotiationFormat2(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("Accept", "image/tiff-fx")

	assert.Equal(t, "", c.NegotiateFormat底层方法("image/tiff"))
}

func TestContextIsAborted(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	assert.False(t, c.X是否已停止())

	c.X停止()
	assert.True(t, c.X是否已停止())

	c.X中间件继续()
	assert.True(t, c.X是否已停止())

	c.index++
	assert.True(t, c.X是否已停止())
}

// TestContextData tests that the response can be written from `bytestring`
// with specified MIME type
func TestContextAbortWithStatus(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.index = 4
	c.X停止并带状态码(http.StatusUnauthorized)

	assert.Equal(t, abortIndex, c.index)
	assert.Equal(t, http.StatusUnauthorized, c.Writer.Status())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.True(t, c.X是否已停止())
}

type testJSONAbortMsg struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func TestContextAbortWithStatusJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)
	c.index = 4

	in := new(testJSONAbortMsg)
	in.Bar = "barValue"
	in.Foo = "fooValue"

	c.X停止并带状态码且返回JSON(http.StatusUnsupportedMediaType, in)

	assert.Equal(t, abortIndex, c.index)
	assert.Equal(t, http.StatusUnsupportedMediaType, c.Writer.Status())
	assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
	assert.True(t, c.X是否已停止())

	contentType := w.Header().Get("Content-Type")
	assert.Equal(t, "application/json; charset=utf-8", contentType)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(w.Body)
	assert.NoError(t, err)
	jsonStringBody := buf.String()
	assert.Equal(t, "{\"foo\":\"fooValue\",\"bar\":\"barValue\"}", jsonStringBody)
}

func TestContextError(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	assert.Empty(t, c.X错误s)

	firstErr := errors.New("first error")
	c.X错误(firstErr) //nolint: errcheck
	assert.Len(t, c.X错误s, 1)
	assert.Equal(t, "Error #01: first error\n", c.X错误s.String())

	secondErr := errors.New("second error")
	c.X错误(&Error{ //nolint: errcheck
		Err:  secondErr,
		Meta: "some data 2",
		Type: ErrorTypePublic,
	})
	assert.Len(t, c.X错误s, 2)

	assert.Equal(t, firstErr, c.X错误s[0].Err)
	assert.Nil(t, c.X错误s[0].Meta)
	assert.Equal(t, ErrorTypePrivate, c.X错误s[0].Type)

	assert.Equal(t, secondErr, c.X错误s[1].Err)
	assert.Equal(t, "some data 2", c.X错误s[1].Meta)
	assert.Equal(t, ErrorTypePublic, c.X错误s[1].Type)

	assert.Equal(t, c.X错误s.Last(), c.X错误s[1])

	defer func() {
		if recover() == nil {
			t.Error("didn't panic")
		}
	}()
	c.X错误(nil) //nolint: errcheck
}

func TestContextTypedError(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X错误(errors.New("externo 0")).SetType(ErrorTypePublic)  //nolint: errcheck
	c.X错误(errors.New("interno 0")).SetType(ErrorTypePrivate) //nolint: errcheck

	for _, err := range c.X错误s.ByType(ErrorTypePublic) {
		assert.Equal(t, ErrorTypePublic, err.Type)
	}
	for _, err := range c.X错误s.ByType(ErrorTypePrivate) {
		assert.Equal(t, ErrorTypePrivate, err.Type)
	}
	assert.Equal(t, []string{"externo 0", "interno 0"}, c.X错误s.Errors())
}

func TestContextAbortWithError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X停止并带状态码与错误(http.StatusUnauthorized, errors.New("bad input")).SetMeta("some input") //nolint: errcheck

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, abortIndex, c.index)
	assert.True(t, c.X是否已停止())
}

func TestContextClientIP(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.engine.trustedCIDRs, _ = c.engine.prepareTrustedCIDRs()
	resetContextForClientIPTests(c)

	// Legacy tests (validating that the defaults don't break the
	// (insecure!) old behaviour)
	assert.Equal(t, "20.20.20.20", c.X取客户端ip())

	c.X请求.Header.Del("X-Forwarded-For")
	assert.Equal(t, "10.10.10.10", c.X取客户端ip())

	c.X请求.Header.Set("X-Forwarded-For", "30.30.30.30  ")
	assert.Equal(t, "30.30.30.30", c.X取客户端ip())

	c.X请求.Header.Del("X-Forwarded-For")
	c.X请求.Header.Del("X-Real-IP")
	c.engine.TrustedPlatform = PlatformGoogleAppEngine
	assert.Equal(t, "50.50.50.50", c.X取客户端ip())

	c.X请求.Header.Del("X-Appengine-Remote-Addr")
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// no port
	c.X请求.RemoteAddr = "50.50.50.50"
	assert.Empty(t, c.X取客户端ip())

	// Tests exercising the TrustedProxies functionality
	resetContextForClientIPTests(c)

	// IPv6 support
	c.X请求.RemoteAddr = "[::1]:12345"
	assert.Equal(t, "20.20.20.20", c.X取客户端ip())

	resetContextForClientIPTests(c)
	// No trusted proxies
	_ = c.engine.X设置受信任代理([]string{})
	c.engine.RemoteIPHeaders = []string{"X-Forwarded-For"}
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// Disabled TrustedProxies feature
	_ = c.engine.X设置受信任代理(nil)
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// Last proxy is trusted, but the RemoteAddr is not
	_ = c.engine.X设置受信任代理([]string{"30.30.30.30"})
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// Only trust RemoteAddr
	_ = c.engine.X设置受信任代理([]string{"40.40.40.40"})
	assert.Equal(t, "30.30.30.30", c.X取客户端ip())

	// All steps are trusted
	_ = c.engine.X设置受信任代理([]string{"40.40.40.40", "30.30.30.30", "20.20.20.20"})
	assert.Equal(t, "20.20.20.20", c.X取客户端ip())

	// Use CIDR
	_ = c.engine.X设置受信任代理([]string{"40.40.25.25/16", "30.30.30.30"})
	assert.Equal(t, "20.20.20.20", c.X取客户端ip())

	// Use hostname that resolves to all the proxies
	_ = c.engine.X设置受信任代理([]string{"foo"})
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// Use hostname that returns an error
	_ = c.engine.X设置受信任代理([]string{"bar"})
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// X-Forwarded-For has a non-IP element
	_ = c.engine.X设置受信任代理([]string{"40.40.40.40"})
	c.X请求.Header.Set("X-Forwarded-For", " blah ")
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// Result from LookupHost has non-IP element. This should never
	// happen, but we should test it to make sure we handle it
	// gracefully.
	_ = c.engine.X设置受信任代理([]string{"baz"})
	c.X请求.Header.Set("X-Forwarded-For", " 30.30.30.30 ")
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	_ = c.engine.X设置受信任代理([]string{"40.40.40.40"})
	c.X请求.Header.Del("X-Forwarded-For")
	c.engine.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}
	assert.Equal(t, "10.10.10.10", c.X取客户端ip())

	c.engine.RemoteIPHeaders = []string{}
	c.engine.TrustedPlatform = PlatformGoogleAppEngine
	assert.Equal(t, "50.50.50.50", c.X取客户端ip())

	// Use custom TrustedPlatform header
	c.engine.TrustedPlatform = "X-CDN-IP"
	c.X请求.Header.Set("X-CDN-IP", "80.80.80.80")
	assert.Equal(t, "80.80.80.80", c.X取客户端ip())
	// wrong header
	c.engine.TrustedPlatform = "X-Wrong-Header"
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	c.X请求.Header.Del("X-CDN-IP")
	// TrustedPlatform is empty
	c.engine.TrustedPlatform = ""
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	// Test the legacy flag
	c.engine.AppEngine弃用 = true
	assert.Equal(t, "50.50.50.50", c.X取客户端ip())
	c.engine.AppEngine弃用 = false
	c.engine.TrustedPlatform = PlatformGoogleAppEngine

	c.X请求.Header.Del("X-Appengine-Remote-Addr")
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	c.engine.TrustedPlatform = PlatformCloudflare
	assert.Equal(t, "60.60.60.60", c.X取客户端ip())

	c.X请求.Header.Del("CF-Connecting-IP")
	assert.Equal(t, "40.40.40.40", c.X取客户端ip())

	c.engine.TrustedPlatform = ""

	// no port
	c.X请求.RemoteAddr = "50.50.50.50"
	assert.Empty(t, c.X取客户端ip())
}

func resetContextForClientIPTests(c *Context) {
	c.X请求.Header.Set("X-Real-IP", " 10.10.10.10  ")
	c.X请求.Header.Set("X-Forwarded-For", "  20.20.20.20, 30.30.30.30")
	c.X请求.Header.Set("X-Appengine-Remote-Addr", "50.50.50.50")
	c.X请求.Header.Set("CF-Connecting-IP", "60.60.60.60")
	c.X请求.RemoteAddr = "  40.40.40.40:42123 "
	c.engine.TrustedPlatform = ""
	c.engine.trustedCIDRs = defaultTrustedCIDRs
	c.engine.AppEngine弃用 = false
}

func TestContextContentType(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Set("Content-Type", "application/json; charset=utf-8")

	assert.Equal(t, "application/json", c.X取协议头ContentType())
}

func TestContextAutoBindJSON(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.X请求.Header.Add("Content-Type", MIMEJSON)

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	assert.NoError(t, c.X取参数到指针PANI(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Empty(t, c.X错误s)
}

func TestContextBindWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	assert.NoError(t, c.X取JSON参数到指针PANI(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBindWithXML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString(`<?xml version="1.0" encoding="UTF-8"?>
		<root>
			<foo>FOO</foo>
		   	<bar>BAR</bar>
		</root>`))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `xml:"foo"`
		Bar string `xml:"bar"`
	}
	assert.NoError(t, c.X取XML参数到指针PANI(&obj))
	assert.Equal(t, "FOO", obj.Foo)
	assert.Equal(t, "BAR", obj.Bar)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBindHeader(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("rate", "8000")
	c.X请求.Header.Add("domain", "music")
	c.X请求.Header.Add("limit", "1000")

	var testHeader struct {
		Rate   int    `header:"Rate"`
		Domain string `header:"Domain"`
		Limit  int    `header:"limit"`
	}

	assert.NoError(t, c.X取Header参数到指针PANI(&testHeader))
	assert.Equal(t, 8000, testHeader.Rate)
	assert.Equal(t, "music", testHeader.Domain)
	assert.Equal(t, 1000, testHeader.Limit)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBindWithQuery(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/?foo=bar&bar=foo", bytes.NewBufferString("foo=unused"))

	var obj struct {
		Foo string `form:"foo"`
		Bar string `form:"bar"`
	}
	assert.NoError(t, c.X取URL参数到指针PANI(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBindWithYAML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("foo: bar\nbar: foo"))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `yaml:"foo"`
		Bar string `yaml:"bar"`
	}
	assert.NoError(t, c.X取YAML参数到指针PANI(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBindWithTOML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("foo = 'bar'\nbar = 'foo'"))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `toml:"foo"`
		Bar string `toml:"bar"`
	}
	assert.NoError(t, c.X取TOML参数到指针PANI(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBadAutoBind(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "http://example.com", bytes.NewBufferString("\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.X请求.Header.Add("Content-Type", MIMEJSON)
	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	assert.False(t, c.X是否已停止())
	assert.Error(t, c.X取参数到指针PANI(&obj))
	c.Writer.WriteHeaderNow()

	assert.Empty(t, obj.Bar)
	assert.Empty(t, obj.Foo)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, c.X是否已停止())
}

func TestContextAutoShouldBindJSON(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.X请求.Header.Add("Content-Type", MIMEJSON)

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	assert.NoError(t, c.X取参数到指针(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Empty(t, c.X错误s)
}

func TestContextShouldBindWithJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	assert.NoError(t, c.X取JSON参数到指针(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextShouldBindWithXML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString(`<?xml version="1.0" encoding="UTF-8"?>
		<root>
			<foo>FOO</foo>
			<bar>BAR</bar>
		</root>`))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `xml:"foo"`
		Bar string `xml:"bar"`
	}
	assert.NoError(t, c.X取XML参数到指针(&obj))
	assert.Equal(t, "FOO", obj.Foo)
	assert.Equal(t, "BAR", obj.Bar)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextShouldBindHeader(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.Header.Add("rate", "8000")
	c.X请求.Header.Add("domain", "music")
	c.X请求.Header.Add("limit", "1000")

	var testHeader struct {
		Rate   int    `header:"Rate"`
		Domain string `header:"Domain"`
		Limit  int    `header:"limit"`
	}

	assert.NoError(t, c.X取Header参数到指针(&testHeader))
	assert.Equal(t, 8000, testHeader.Rate)
	assert.Equal(t, "music", testHeader.Domain)
	assert.Equal(t, 1000, testHeader.Limit)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextShouldBindWithQuery(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/?foo=bar&bar=foo&Foo=bar1&Bar=foo1", bytes.NewBufferString("foo=unused"))

	var obj struct {
		Foo  string `form:"foo"`
		Bar  string `form:"bar"`
		Foo1 string `form:"Foo"`
		Bar1 string `form:"Bar"`
	}
	assert.NoError(t, c.X取URL参数到指针(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, "foo1", obj.Bar1)
	assert.Equal(t, "bar1", obj.Foo1)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextShouldBindWithYAML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("foo: bar\nbar: foo"))
	c.X请求.Header.Add("Content-Type", MIMEXML) // set fake content-type

	var obj struct {
		Foo string `yaml:"foo"`
		Bar string `yaml:"bar"`
	}
	assert.NoError(t, c.X取YAML参数到指针(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextShouldBindWithTOML(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("foo='bar'\nbar= 'foo'"))
	c.X请求.Header.Add("Content-Type", MIMETOML) // set fake content-type

	var obj struct {
		Foo string `toml:"foo"`
		Bar string `toml:"bar"`
	}
	assert.NoError(t, c.X取TOML参数到指针(&obj))
	assert.Equal(t, "foo", obj.Bar)
	assert.Equal(t, "bar", obj.Foo)
	assert.Equal(t, 0, w.Body.Len())
}

func TestContextBadAutoShouldBind(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	c.X请求, _ = http.NewRequest("POST", "http://example.com", bytes.NewBufferString("\"foo\":\"bar\", \"bar\":\"foo\"}"))
	c.X请求.Header.Add("Content-Type", MIMEJSON)
	var obj struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}

	assert.False(t, c.X是否已停止())
	assert.Error(t, c.X取参数到指针(&obj))

	assert.Empty(t, obj.Bar)
	assert.Empty(t, obj.Foo)
	assert.False(t, c.X是否已停止())
}

func TestContextShouldBindBodyWith(t *testing.T) {
	type typeA struct {
		Foo string `json:"foo" xml:"foo" binding:"required"`
	}
	type typeB struct {
		Bar string `json:"bar" xml:"bar" binding:"required"`
	}
	for _, tt := range []struct {
		name               string
		bindingA, bindingB binding.BindingBody
		bodyA, bodyB       string
	}{
		{
			name:     "JSON & JSON",
			bindingA: binding.JSON,
			bindingB: binding.JSON,
			bodyA:    `{"foo":"FOO"}`,
			bodyB:    `{"bar":"BAR"}`,
		},
		{
			name:     "JSON & XML",
			bindingA: binding.JSON,
			bindingB: binding.XML,
			bodyA:    `{"foo":"FOO"}`,
			bodyB: `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <bar>BAR</bar>
</root>`,
		},
		{
			name:     "XML & XML",
			bindingA: binding.XML,
			bindingB: binding.XML,
			bodyA: `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <foo>FOO</foo>
</root>`,
			bodyB: `<?xml version="1.0" encoding="UTF-8"?>
<root>
   <bar>BAR</bar>
</root>`,
		},
	} {
		t.Logf("testing: %s", tt.name)
		// bodyA to typeA and typeB
		{
			w := httptest.NewRecorder()
			c, _ := CreateTestContext(w)
			c.X请求, _ = http.NewRequest(
				"POST", "http://example.com", bytes.NewBufferString(tt.bodyA),
			)
			// When it binds to typeA and typeB, it finds the body is
			// not typeB but typeA.
			objA := typeA{}
			assert.NoError(t, c.X取参数到指针并按类型且缓存(&objA, tt.bindingA))
			assert.Equal(t, typeA{"FOO"}, objA)
			objB := typeB{}
			assert.Error(t, c.X取参数到指针并按类型且缓存(&objB, tt.bindingB))
			assert.NotEqual(t, typeB{"BAR"}, objB)
		}
		// bodyB to typeA and typeB
		{
			// When it binds to typeA and typeB, it finds the body is
			// not typeA but typeB.
			w := httptest.NewRecorder()
			c, _ := CreateTestContext(w)
			c.X请求, _ = http.NewRequest(
				"POST", "http://example.com", bytes.NewBufferString(tt.bodyB),
			)
			objA := typeA{}
			assert.Error(t, c.X取参数到指针并按类型且缓存(&objA, tt.bindingA))
			assert.NotEqual(t, typeA{"FOO"}, objA)
			objB := typeB{}
			assert.NoError(t, c.X取参数到指针并按类型且缓存(&objB, tt.bindingB))
			assert.Equal(t, typeB{"BAR"}, objB)
		}
	}
}

func TestContextGolangContext(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", bytes.NewBufferString("{\"foo\":\"bar\", \"bar\":\"foo\"}"))
	assert.NoError(t, c.Err())
	assert.Nil(t, c.Done())
	ti, ok := c.Deadline()
	assert.Equal(t, ti, time.Time{})
	assert.False(t, ok)
	assert.Equal(t, c.Value(0), c.X请求)
	assert.Equal(t, c.Value(ContextKey), c)
	assert.Nil(t, c.Value("foo"))

	c.X设置值("foo", "bar")
	assert.Equal(t, "bar", c.Value("foo"))
	assert.Nil(t, c.Value(1))
}

func TestWebsocketsRequired(t *testing.T) {
	// Example request from spec: https://tools.ietf.org/html/rfc6455#section-1.2
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("GET", "/chat", nil)
	c.X请求.Header.Set("Host", "server.example.com")
	c.X请求.Header.Set("Upgrade", "websocket")
	c.X请求.Header.Set("Connection", "Upgrade")
	c.X请求.Header.Set("Sec-WebSocket-Key", "dGhlIHNhbXBsZSBub25jZQ==")
	c.X请求.Header.Set("Origin", "http://example.com")
	c.X请求.Header.Set("Sec-WebSocket-Protocol", "chat, superchat")
	c.X请求.Header.Set("Sec-WebSocket-Version", "13")

	assert.True(t, c.X是否为Websocket请求())

	// Normal request, no websocket required.
	c, _ = CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("GET", "/chat", nil)
	c.X请求.Header.Set("Host", "server.example.com")

	assert.False(t, c.X是否为Websocket请求())
}

func TestGetRequestHeaderValue(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("GET", "/chat", nil)
	c.X请求.Header.Set("Gin-Version", "1.0.0")

	assert.Equal(t, "1.0.0", c.X取请求协议头值("Gin-Version"))
	assert.Empty(t, c.X取请求协议头值("Connection"))
}

func TestContextGetRawData(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString("Fetch binary post data")
	c.X请求, _ = http.NewRequest("POST", "/", body)
	c.X请求.Header.Add("Content-Type", MIMEPOSTForm)

	data, err := c.X取流数据()
	assert.Nil(t, err)
	assert.Equal(t, "Fetch binary post data", string(data))
}

func TestContextRenderDataFromReader(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	body := "#!PNG some raw data"
	reader := strings.NewReader(body)
	contentLength := int64(len(body))
	contentType := "image/png"
	extraHeaders := map[string]string{"Content-Disposition": `attachment; filename="gopher.png"`}

	c.X输出字节集并按IO(http.StatusOK, contentLength, contentType, reader, extraHeaders)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, w.Body.String())
	assert.Equal(t, contentType, w.Header().Get("Content-Type"))
	assert.Equal(t, fmt.Sprintf("%d", contentLength), w.Header().Get("Content-Length"))
	assert.Equal(t, extraHeaders["Content-Disposition"], w.Header().Get("Content-Disposition"))
}

func TestContextRenderDataFromReaderNoHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := CreateTestContext(w)

	body := "#!PNG some raw data"
	reader := strings.NewReader(body)
	contentLength := int64(len(body))
	contentType := "image/png"

	c.X输出字节集并按IO(http.StatusOK, contentLength, contentType, reader, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, body, w.Body.String())
	assert.Equal(t, contentType, w.Header().Get("Content-Type"))
	assert.Equal(t, fmt.Sprintf("%d", contentLength), w.Header().Get("Content-Length"))
}

type TestResponseRecorder struct {
	*httptest.ResponseRecorder
	closeChannel chan bool
}

func (r *TestResponseRecorder) CloseNotify() <-chan bool {
	return r.closeChannel
}

func (r *TestResponseRecorder) closeClient() {
	r.closeChannel <- true
}

func CreateTestResponseRecorder() *TestResponseRecorder {
	return &TestResponseRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func TestContextStream(t *testing.T) {
	w := CreateTestResponseRecorder()
	c, _ := CreateTestContext(w)

	stopStream := true
	c.Stream(func(w io.Writer) bool {
		defer func() {
			stopStream = false
		}()

		_, err := w.Write([]byte("test"))
		assert.NoError(t, err)

		return stopStream
	})

	assert.Equal(t, "testtest", w.Body.String())
}

func TestContextStreamWithClientGone(t *testing.T) {
	w := CreateTestResponseRecorder()
	c, _ := CreateTestContext(w)

	c.Stream(func(writer io.Writer) bool {
		defer func() {
			w.closeClient()
		}()

		_, err := writer.Write([]byte("test"))
		assert.NoError(t, err)

		return true
	})

	assert.Equal(t, "test", w.Body.String())
}

func TestContextResetInHandler(t *testing.T) {
	w := CreateTestResponseRecorder()
	c, _ := CreateTestContext(w)

	c.handlers = []HandlerFunc{
		func(c *Context) { c.reset() },
	}
	assert.NotPanics(t, func() {
		c.X中间件继续()
	})
}

func TestRaceParamsContextCopy(t *testing.T) {
	DefaultWriter = os.Stdout
	router := X创建默认对象()
	nameGroup := router.X创建分组路由("/:name")
	var wg sync.WaitGroup
	wg.Add(2)
	{
		nameGroup.X绑定GET("/api", func(c *Context) {
			go func(c *Context, param string) {
				defer wg.Done()
				// First assert must be executed after the second request
				time.Sleep(50 * time.Millisecond)
				assert.Equal(t, c.X取API参数值("name"), param)
			}(c.X取副本(), c.X取API参数值("name"))
		})
	}
	PerformRequest(router, "GET", "/name1/api")
	PerformRequest(router, "GET", "/name2/api")
	wg.Wait()
}

func TestContextWithKeysMutex(t *testing.T) {
	c := &Context{}
	c.X设置值("foo", "bar")

	value, err := c.X取值("foo")
	assert.Equal(t, "bar", value)
	assert.True(t, err)

	value, err = c.X取值("foo2")
	assert.Nil(t, value)
	assert.False(t, err)
}

func TestRemoteIPFail(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	c.X请求, _ = http.NewRequest("POST", "/", nil)
	c.X请求.RemoteAddr = "[:::]:80"
	ip := net.ParseIP(c.X取协议头ip())
	trust := c.engine.isTrustedProxy(ip)
	assert.Nil(t, ip)
	assert.False(t, trust)
}

func TestHasRequestContext(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	assert.False(t, c.hasRequestContext(), "no request, no fallback")
	c.engine.ContextWithFallback = true
	assert.False(t, c.hasRequestContext(), "no request, has fallback")
	c.X请求, _ = http.NewRequest(http.MethodGet, "/", nil)
	assert.True(t, c.hasRequestContext(), "has request, has fallback")
	c.X请求, _ = http.NewRequestWithContext(nil, "", "", nil) //nolint:staticcheck
	assert.False(t, c.hasRequestContext(), "has request with nil ctx, has fallback")
	c.engine.ContextWithFallback = false
	assert.False(t, c.hasRequestContext(), "has request, no fallback")

	c = &Context{}
	assert.False(t, c.hasRequestContext(), "no request, no engine")
	c.X请求, _ = http.NewRequest(http.MethodGet, "/", nil)
	assert.False(t, c.hasRequestContext(), "has request, no engine")
}

func TestContextWithFallbackDeadlineFromRequestContext(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	// enable ContextWithFallback feature flag
	c.engine.ContextWithFallback = true

	deadline, ok := c.Deadline()
	assert.Zero(t, deadline)
	assert.False(t, ok)

	c2, _ := CreateTestContext(httptest.NewRecorder())
	// enable ContextWithFallback feature flag
	c2.engine.ContextWithFallback = true

	c2.X请求, _ = http.NewRequest(http.MethodGet, "/", nil)
	d := time.Now().Add(time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	c2.X请求 = c2.X请求.WithContext(ctx)
	deadline, ok = c2.Deadline()
	assert.Equal(t, d, deadline)
	assert.True(t, ok)
}

func TestContextWithFallbackDoneFromRequestContext(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	// enable ContextWithFallback feature flag
	c.engine.ContextWithFallback = true

	assert.Nil(t, c.Done())

	c2, _ := CreateTestContext(httptest.NewRecorder())
	// enable ContextWithFallback feature flag
	c2.engine.ContextWithFallback = true

	c2.X请求, _ = http.NewRequest(http.MethodGet, "/", nil)
	ctx, cancel := context.WithCancel(context.Background())
	c2.X请求 = c2.X请求.WithContext(ctx)
	cancel()
	assert.NotNil(t, <-c2.Done())
}

func TestContextWithFallbackErrFromRequestContext(t *testing.T) {
	c, _ := CreateTestContext(httptest.NewRecorder())
	// enable ContextWithFallback feature flag
	c.engine.ContextWithFallback = true

	assert.Nil(t, c.Err())

	c2, _ := CreateTestContext(httptest.NewRecorder())
	// enable ContextWithFallback feature flag
	c2.engine.ContextWithFallback = true

	c2.X请求, _ = http.NewRequest(http.MethodGet, "/", nil)
	ctx, cancel := context.WithCancel(context.Background())
	c2.X请求 = c2.X请求.WithContext(ctx)
	cancel()

	assert.EqualError(t, c2.Err(), context.Canceled.Error())
}

func TestContextWithFallbackValueFromRequestContext(t *testing.T) {
	type contextKey string

	tests := []struct {
		name             string
		getContextAndKey func() (*Context, any)
		value            any
	}{
		{
			name: "c with struct context key",
			getContextAndKey: func() (*Context, any) {
				var key struct{}
				c, _ := CreateTestContext(httptest.NewRecorder())
				// enable ContextWithFallback feature flag
				c.engine.ContextWithFallback = true
				c.X请求, _ = http.NewRequest("POST", "/", nil)
				c.X请求 = c.X请求.WithContext(context.WithValue(context.TODO(), key, "value"))
				return c, key
			},
			value: "value",
		},
		{
			name: "c with string context key",
			getContextAndKey: func() (*Context, any) {
				c, _ := CreateTestContext(httptest.NewRecorder())
				// enable ContextWithFallback feature flag
				c.engine.ContextWithFallback = true
				c.X请求, _ = http.NewRequest("POST", "/", nil)
				c.X请求 = c.X请求.WithContext(context.WithValue(context.TODO(), contextKey("key"), "value"))
				return c, contextKey("key")
			},
			value: "value",
		},
		{
			name: "c with nil http.Request",
			getContextAndKey: func() (*Context, any) {
				c, _ := CreateTestContext(httptest.NewRecorder())
				// enable ContextWithFallback feature flag
				c.engine.ContextWithFallback = true
				c.X请求 = nil
				return c, "key"
			},
			value: nil,
		},
		{
			name: "c with nil http.Request.Context()",
			getContextAndKey: func() (*Context, any) {
				c, _ := CreateTestContext(httptest.NewRecorder())
				// enable ContextWithFallback feature flag
				c.engine.ContextWithFallback = true
				c.X请求, _ = http.NewRequest("POST", "/", nil)
				return c, "key"
			},
			value: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, key := tt.getContextAndKey()
			assert.Equal(t, tt.value, c.Value(key))
		})
	}
}

func TestContextCopyShouldNotCancel(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	ensureRequestIsOver := make(chan struct{})

	wg := &sync.WaitGroup{}

	r := X创建()
	r.X绑定GET("/", func(ginctx *Context) {
		wg.Add(1)

		ginctx = ginctx.X取副本()

		// start async goroutine for calling srv
		go func() {
			defer wg.Done()

			<-ensureRequestIsOver // ensure request is done

			req, err := http.NewRequestWithContext(ginctx, http.MethodGet, srv.URL, nil)
			must(err)

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Error(fmt.Errorf("request error: %w", err))
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Error(fmt.Errorf("unexpected status code: %s", res.Status))
			}
		}()
	})

	l, err := net.Listen("tcp", ":0")
	must(err)
	go func() {
		s := &http.Server{
			Handler: r,
		}

		must(s.Serve(l))
	}()

	addr := strings.Split(l.Addr().String(), ":")
	res, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/", addr[len(addr)-1]))
	if err != nil {
		t.Error(fmt.Errorf("request error: %w", err))
		return
	}

	close(ensureRequestIsOver)

	if res.StatusCode != http.StatusOK {
		t.Error(fmt.Errorf("unexpected status code: %s", res.Status))
		return
	}

	wg.Wait()
}

func TestContextAddParam(t *testing.T) {
	c := &Context{}
	id := "id"
	value := "1"
	c.X设置API参数值(id, value)

	v, ok := c.X参数.Get(id)
	assert.Equal(t, ok, true)
	assert.Equal(t, value, v)
}

func TestCreateTestContextWithRouteParams(t *testing.T) {
	w := httptest.NewRecorder()
	engine := X创建()
	engine.X绑定GET("/:action/:name", func(ctx *Context) {
		ctx.X输出文本(http.StatusOK, "%s %s", ctx.X取API参数值("action"), ctx.X取API参数值("name"))
	})
	c := CreateTestContextOnly(w, engine)
	c.X请求, _ = http.NewRequest(http.MethodGet, "/hello/gin", nil)
	engine.HandleContext底层方法(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin", w.Body.String())
}

type interceptedWriter struct {
	ResponseWriter
	b *bytes.Buffer
}

func (i interceptedWriter) WriteHeader(code int) {
	i.Header().Del("X-Test")
	i.ResponseWriter.WriteHeader(code)
}

func TestInterceptedHeader(t *testing.T) {
	w := httptest.NewRecorder()
	c, r := CreateTestContext(w)

	r.X中间件(func(c *Context) {
		i := interceptedWriter{
			ResponseWriter: c.Writer,
			b:              bytes.NewBuffer(nil),
		}
		c.Writer = i
		c.X中间件继续()
		c.X设置响应协议头值("X-Test", "overridden")
		c.Writer = i.ResponseWriter
	})
	r.X绑定GET("/", func(c *Context) {
		c.X设置响应协议头值("X-Test", "original")
		c.X设置响应协议头值("X-Test-2", "present")
		c.X输出文本(http.StatusOK, "hello world")
	})
	c.X请求 = httptest.NewRequest("GET", "/", nil)
	r.HandleContext底层方法(c)
	// Result() has headers frozen when WriteHeaderNow() has been called
	// Compared to this time, this is when the response headers will be flushed
	// As response is flushed on c.String, the Header cannot be set by the first
	// middleware. Assert this
	assert.Equal(t, "", w.Result().Header.Get("X-Test"))
	assert.Equal(t, "present", w.Result().Header.Get("X-Test-2"))
}
