# # gRPC示例

本指南通过一个简单的实际示例，帮助您开始在Go中使用gRPC。
## # 前置条件

使用以下命令安装Go的协议编译器插件：

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

更新你的`PATH`环境变量，以便`protoc`编译器能够找到这些插件：

```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```
## # 重新生成gRPC代码

```sh
protoc --go_out=gen --go_opt=paths=source_relative \
  --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
  -I=$PWD pb/helloworld.proto
```

这段命令用于通过`protoc`编译器重新生成gRPC的Go语言代码。具体步骤如下：

1. `protoc`：这是Google Protocol Buffers的编译器，用于将`.proto`文件编译为目标语言的代码。

2. `--go_out=gen`：指定Go语言代码的输出目录为`gen`。

3. `--go_opt=paths=source_relative`：设置Go代码中导入路径为相对路径，相对于源`.proto`文件的位置。

4. `--go-grpc_out=gen`：同时生成gRPC的Go语言代码，并将其输出到`gen`目录下。

5. `--go-grpc_opt=paths=source_relative`：设置gRPC Go代码的导入路径同样采用相对路径。

6. `-I=$PWD`：指定搜索包含`.proto`文件的目录，这里使用当前工作目录（$PWD）。

7. `pb/helloworld.proto`：指定要编译的`.proto`文件路径，本例中是`pb`目录下的`helloworld.proto`文件。
## # 运行

第一步：启动 gRPC 服务器

```sh
go run grpc/server.go
```

第二步：启动 Gin 服务器

```sh
go run gin/main.go
```
## # 测试

向 gin 服务器发送数据：

```sh
curl -v 'http://localhost:8080/rest/n/gin'
```

或者使用 [grpcurl](https://github.com/fullstorydev/grpcurl) 命令：

```sh
grpcurl -d '{"name": "gin"}' \
  -plaintext localhost:50051 helloworld.v1.Greeter/SayHello
```
