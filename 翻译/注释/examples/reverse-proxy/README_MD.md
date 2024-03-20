
<原文开始>
A simple reverse proxy

We can see real server in real_server.go and proxy server in reverse_server.go

Run this two file and if we do some request like `curl 'http://localhost:2002/something`

we will get a response in JSON contains ip of client and path we requested.

<原文结束>

# <翻译开始>
# 一个简单的反向代理

在 real_server.go 文件中，我们可以看到真实服务器的代码，在 reverse_server.go 文件中则是代理服务器的代码。

运行这两个文件，如果我们执行类似 `curl 'http://localhost:2002/something'` 的请求，

我们将得到一个 JSON 格式的响应，其中包含了客户端的 IP 地址和我们请求的路径。

# <翻译结束>

