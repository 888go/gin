
<原文开始>
// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// params[0]=url example:http://127.0.0.1:8080/index (cannot be empty)
// params[1]=response status (custom compare status) default:"200 OK"
// params[2]=response body (custom compare content)  default:"it worked"
<原文结束>

# <翻译开始>
// params[0]=url示例:http://127.0.0.1:8080/index(不能为空)params[1]=response status(自定义比较状态)默认值:"200 OK"Params[2]=响应体(自定义比较内容)默认值:"它工作"
# <翻译结束>


<原文开始>
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
<原文结束>

# <翻译开始>
// 必须等待程序启动并运行服务器，否则主线程将完成
# <翻译结束>


<原文开始>
		// not supported by windows, it is unimplemented now
<原文结束>

# <翻译开始>
// windows不支持，目前未实现
# <翻译结束>


<原文开始>
// func TestWithHttptestWithSpecifiedPort(t *testing.T) {
// 	router := New()
// 	router.GET("/example", func(c *Context) { c.String(http.StatusOK, "it worked") })
<原文结束>

# <翻译开始>
// func TestWithHttptestWithSpecifiedPort(t *testing.T) {router:= New() router. get ("/example"， func(c *Context) {c. string (http. string)StatusOK， "它工作")})
# <翻译结束>


<原文开始>
// 	l, _ := net.Listen("tcp", ":8033")
// 	ts := httptest.Server{
// 		Listener: l,
// 		Config:   &http.Server{Handler: router},
// 	}
// 	ts.Start()
// 	defer ts.Close()
<原文结束>

# <翻译开始>
// 1， _:= net.Listen("tcp"， ":8033") ts:= httptest
// 服务器{监听器:1，配置:&http
// Server{Handler: router}，} ts.Start() defer ts.Close()
# <翻译结束>


<原文开始>
// 	testRequest(t, "http://localhost:8033/example")
// }
<原文结束>

# <翻译开始>
// testquest (t， "http://localhost:8033/example")}
# <翻译结束>


<原文开始>
	// 404 not found
<原文结束>

# <翻译开始>
// 404未找到
# <翻译结束>

