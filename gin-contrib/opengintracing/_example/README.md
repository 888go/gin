# # 这是一个简单的API网关、server1、server2和server3实例的实现。

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
