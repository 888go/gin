
<原文开始>
Building a single binary containing templates

This is a complete example to create a single binary with the
[gin-gonic/gin][gin] Web Server with HTML templates.

[gin]: https://github.com/gin-gonic/gin


<原文结束>

# <翻译开始>
# 构建包含模板的单个二进制文件

这是一个完整的示例，演示如何使用 [gin-gonic/gin][gin] Web 服务器（含 HTML 模板）创建一个单一的二进制文件。

[gin]: https://github.com/gin-gonic/gin

（翻译：这个例子旨在说明如何将 gin-gonic/gin 库中的 Web 服务器与 HTML 模板结合，以生成一个单一的二进制程序。）

（注：[gin] 是指向 gin-gonic/gin GitHub 仓库的链接。）

# <翻译结束>


<原文开始>
Prepare Packages

```sh
go get github.com/gin-gonic/gin
go install github.com/jessevdk/go-assets-builder@latest
```

#
<原文结束>

# <翻译开始>
# 准备包

```sh
go get github.com/gin-gonic/gin
go install github.com/jessevdk/go-assets-builder@latest
```

#

# <翻译结束>


<原文开始>
Generate assets.go

```sh
go-assets-builder html -o assets.go
```

#
<原文结束>

# <翻译开始>
# 生成assets.go

```sh
go-assets-builder html -o assets.go
```

#

# <翻译结束>


<原文开始>
Build the server

```sh
go build -o assets-in-binary
```

#
<原文结束>

# <翻译开始>
# 构建服务器

```sh
go build -o assets-in-binary
```

#

# <翻译结束>


<原文开始>
Run

```sh
./assets-in-binary
```

<原文结束>

# <翻译开始>
# 运行

```shell
./assets-in-binary
```

# <翻译结束>

