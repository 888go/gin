# # 构建包含模板的单个二进制文件

这是一个完整的示例，演示如何使用 [gin-gonic/gin][gin] Web 服务器（含 HTML 模板）创建一个单一的二进制文件。

[gin]: https://github.com/gin-gonic/gin

（翻译：这个例子旨在说明如何将 gin-gonic/gin 库中的 Web 服务器与 HTML 模板结合，以生成一个单一的二进制程序。）

（注：[gin] 是指向 gin-gonic/gin GitHub 仓库的链接。）
## How to use

### # 准备包

```sh
go get github.com/gin-gonic/gin
go install github.com/jessevdk/go-assets-builder@latest
```

#
## # 生成assets.go

```sh
go-assets-builder html -o assets.go
```

#
## # 构建服务器

```sh
go build -o assets-in-binary
```

#
## # 运行

```shell
./assets-in-binary
```
