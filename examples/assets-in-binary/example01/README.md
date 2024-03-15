# # 构建包含模板的单个二进制文件

这是一个完整的示例，用于使用 [gin-gonic/gin][gin] Web 服务器和 HTML 模板创建一个单个二进制文件。

[gin]: https://github.com/gin-gonic/gin

（翻译：）

构建包含模板的单一可执行文件

这是一个完整的实例，展示如何将 [gin-gonic/gin][gin] Web 服务器与 HTML 模板结合，生成一个单一的可执行文件。

[gin]: https://github.com/gin-gonic/gin
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

```sh
./assets-in-binary
```
