
<原文开始>
tracing

[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/opengintracing)](https://goreportcard.com/report/github.com/gin-contrib/opengintracing)
[![GoDoc](https://godoc.org/github.com/gin-contrib/opengintracing?status.png)](https://godoc.org/github.com/gin-contrib/opengintracing)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple implementation of the api gateway, server1, server2 and server3 examples.

To build:

```shell
go build api_gateway.go
go build server1.go
go build server2.go
go build server3.go
```

To run, open four terminals, and execute the following:

```shell
./api_gateway
./server1
./server2
./server3
```

To test:

```shell
curl -X POST http://localhost:8000/service1 -v
curl -X POST http://localhost:8000/service2 -v
```

Header information is printed to stdout. You should see headers propagated from service to service.

On the API gateway:

```sh
Incoming Headers
User-Agent: [curl/7.47.0]
Accept: [*/*]
Outgoing Headers
X-B3-Traceid: [34a0552a76c40432]
X-B3-Spanid: [34a0552a76c40432]
X-B3-Sampled: [1]
```

On service1:

```sh
Incoming headers
User-Agent: [Go-http-client/1.1]
Content-Length: [0]
X-B3-Sampled: [1]
X-B3-Spanid: [65025274cfd25c6b]
X-B3-Traceid: [65025274cfd25c6b]
Accept-Encoding: [gzip]
Outgoing headers
X-B3-Traceid: [65025274cfd25c6b]
X-B3-Parentspanid: [65025274cfd25c6b]
X-B3-Spanid: [f8ef4e9c50cad67]
X-B3-Sampled: [1]
```

On service3:

```sh
Incoming Headers
X-B3-Spanid: [aa66150e951c54c]
X-B3-Traceid: [10386b198f22ca04]
Accept-Encoding: [gzip]
User-Agent: [Go-http-client/1.1]
Content-Length: [0]
X-B3-Parentspanid: [10386b198f22ca04]
X-B3-Sampled: [1]
```

<原文结束>

# <翻译开始>
# 这是一个简单的API网关、server1、server2和server3实例的实现。

构建方法：

```shell
go build api_gateway.go
go build server1.go
go build server2.go
go build server3.go
```

运行方式：打开四个终端，分别执行以下命令：

```shell
./api_gateway
./server1
./server2
./server3
```

测试方法：

```shell
curl -X POST http://localhost:8000/service1 -v
curl -X POST http://localhost:8000/service2 -v
```

Header信息将被打印到标准输出。您应能看到请求头从一个服务传递到另一个服务的过程。

在API网关端：

```sh
传入请求头
User-Agent: [curl/7.47.0]
Accept: [*/*]
传出请求头
X-B3-Traceid: [34a0552a76c40432]
X-B3-Spanid: [34a0552a76c40432]
X-B3-Sampled: [1]
```

在service1上：

```sh
传入请求头
User-Agent: [Go-http-client/1.1]
Content-Length: [0]
X-B3-Sampled: [1]
X-B3-Spanid: [65025274cfd25c6b]
X-B3-Traceid: [65025274cfd25c6b]
Accept-Encoding: [gzip]
传出请求头
X-B3-Traceid: [65025274cfd25c6b]
X-B3-Parentspanid: [65025274cfd25c6b]
X-B3-Spanid: [f8ef4e9c50cad67]
X-B3-Sampled: [1]
```

在service3上：

```sh
传入请求头
X-B3-Spanid: [aa66150e951c54c]
X-B3-Traceid: [10386b198f22ca04]
Accept-Encoding: [gzip]
User-Agent: [Go-http-client/1.1]
Content-Length: [0]
X-B3-Parentspanid: [10386b198f22ca04]
X-B3-Sampled: [1]
``` 

这些内容展示了如何构建和运行一个简单的分布式追踪系统，并通过curl命令进行测试。其中，`X-B3-Traceid`、`X-B3-Spanid` 和 `X-B3-Parentspanid` 等请求头用于追踪跨服务调用链路。

# <翻译结束>

