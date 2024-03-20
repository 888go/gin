
<原文开始>
// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。
# <翻译结束>


<原文开始>
// params[0]=url example:http://127.0.0.1:8080/index (cannot be empty)
// params[1]=response status (custom compare status) default:"200 OK"
// params[2]=response body (custom compare content)  default:"it worked"
<原文结束>

# <翻译开始>
// params[0] = url 示例：http://127.0.0.1:8080/index （不能为空）
// params[1] = 响应状态（自定义比较状态）默认值："200 OK"
// params[2] = 响应体内容（自定义比较内容）默认值："it worked"
# <翻译结束>


<原文开始>
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
<原文结束>

# <翻译开始>
// 必须等待 goroutine 启动并运行服务器
// 否则主线程将提前完成
# <翻译结束>


<原文开始>
// not supported by windows, it is unimplemented now
<原文结束>

# <翻译开始>
// 不支持Windows系统，目前尚未实现
# <翻译结束>


<原文开始>
// func TestWithHttptestWithSpecifiedPort(t *testing.T) {
// 	router := New()
// 	router.GET("/example", func(c *Context) { c.String(http.StatusOK, "it worked") })
<原文结束>

# <翻译开始>
// 函数TestWithHttptestWithSpecifiedPort用于测试指定端口的功能
// (接收一个*testing.T类型的参数t)
// 
// 创建一个新的路由实例router
// 在router上注册GET方法的"/example"路由处理函数，当该路由被访问时，
// 会调用传入的函数，向Context写入http.StatusOK状态码及字符串"it worked"作为响应内容
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
// 创建一个监听TCP端口8033的网络监听器，忽略可能的错误
// 初始化一个httptest.Server结构体，其中Listener字段设置为上述创建的监听器，Config字段设置为一个http.Server指针，其Handler字段设置为router
// 调用ts.Start()方法启动服务器
// 使用defer关键字确保在函数结束时调用ts.Close()方法关闭服务器
# <翻译结束>


<原文开始>
// 	testRequest(t, "http://localhost:8033/example")
// }
<原文结束>

# <翻译开始>
// 测试请求(t, "http://localhost:8033/example")
// }
# <翻译结束>

