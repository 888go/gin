
<原文开始>
gRPC Example

This guide gets you started with gRPC in Go with a simple working example.


<原文结束>

# <翻译开始>
# gRPC 示例

本指南通过一个简单的实际示例帮助您开始在 Go 中使用 gRPC。

# <翻译结束>


<原文开始>
Prerequisites

Install the protocol compiler plugins for Go using the following commands:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Update your `PATH` so that the `protoc` compiler can find the plugins:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```


<原文结束>

# <翻译开始>
# 前置条件

使用以下命令安装Go的协议编译器插件：

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

更新你的`PATH`环境变量，以便`protoc`编译器能够找到这些插件：

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

# <翻译结束>


<原文开始>
Regenerate gRPC code

```sh
protoc --go_out=gen --go_opt=paths=source_relative \
  --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
  -I=$PWD pb/helloworld.proto
```


<原文结束>

# <翻译开始>
# 重新生成gRPC代码

```sh
protoc --go_out=gen --go_opt=paths=source_relative \
  --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
  -I=$PWD pb/helloworld.proto
```

这段命令用于通过`protoc`编译器重新生成gRPC的Go语言代码。参数说明如下：

- `--go_out=gen`: 指定Go语言代码输出目录为`gen`。
- `--go_opt=paths=source_relative`: 设置Go代码中的import路径为相对于源文件的相对路径。
- `--go-grpc_out=gen`: 同样指定gRPC Go代码（基于protobuf的服务定义）的输出目录为`gen`。
- `--go-grpc_opt=paths=source_relative`: 设置gRPC Go代码中的import路径也为相对于源文件的相对路径。
- `-I=$PWD`: 指定包含.proto文件的搜索目录，这里使用当前工作目录（$PWD）作为搜索路径。
- `pb/helloworld.proto`: 需要编译的protobuf接口描述文件，位于`pb`目录下，文件名为`helloworld.proto`。

# <翻译结束>


<原文开始>
Runing

First Step: run grpc server

```sh
go run grpc/server.go
```

Second Step: run gin server

```sh
go run gin/main.go
```


<原文结束>

# <翻译开始>
# 运行

第一步：启动 gRPC 服务器

```sh
go run grpc/server.go
```

第二步：启动 Gin 服务器

```sh
go run gin/main.go
```

# <翻译结束>


<原文开始>
Testing

Send data to gin server:

```sh
curl -v 'http://localhost:8080/rest/n/gin'
```

or using [grpcurl](https://github.com/fullstorydev/grpcurl) command:

```sh
grpcurl -d '{"name": "gin"}' \
  -plaintext localhost:50051 helloworld.v1.Greeter/SayHello
```

<原文结束>

# <翻译开始>
# 测试

向gin服务器发送数据：

```sh
curl -v 'http://localhost:8080/rest/n/gin'
```

或者使用[grpcurl](https://github.com/fullstorydev/grpcurl)命令：

```sh
grpcurl -d '{"name": "gin"}' \
  -plaintext localhost:50051 helloworld.v1.Greeter/SayHello
```

# <翻译结束>

