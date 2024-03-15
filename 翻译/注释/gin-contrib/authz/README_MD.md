
<原文开始>
Authz

[![CodeQL](https://github.com/gin-contrib/authz/actions/workflows/codeql.yml/badge.svg)](https://github.com/gin-contrib/authz/actions/workflows/codeql.yml)
[![Run Tests](https://github.com/gin-contrib/authz/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/authz/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/authz/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/authz)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/authz)](https://goreportcard.com/report/github.com/gin-contrib/authz)
[![GoDoc](https://godoc.org/github.com/gin-contrib/authz?status.svg)](https://godoc.org/github.com/gin-contrib/authz)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Authz is an authorization middleware for [Gin](https://github.com/gin-gonic/gin), it's based on [https://github.com/casbin/casbin](https://github.com/casbin/casbin).


<原文结束>

# <翻译开始>
# Authz

[![CodeQL](https://github.com/gin-contrib/authz/actions/workflows/codeql.yml/badge.svg)](https://github.com/gin-contrib/authz/actions/workflows/codeql.yml)
[![运行测试](https://github.com/gin-contrib/authz/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/authz/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/authz/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/authz)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/authz)](https://goreportcard.com/report/github.com/gin-contrib/authz)
[![GoDoc](https://godoc.org/github.com/gin-contrib/authz?status.svg)](https://godoc.org/github.com/gin-contrib/authz)
[![加入Gitter聊天](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Authz 是 Gin（[https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)）的一个授权中间件，基于 [https://github.com/casbin/casbin](https://github.com/casbin/casbin) 构建。

# <翻译结束>


<原文开始>
Installation

```bash
go get github.com/gin-contrib/authz
```


<原文结束>

# <翻译开始>
# 安装

```bash
go get github.com/gin-contrib/authz
```

# <翻译结束>


<原文开始>
Simple Example

```Go
package main

import (
  "net/http"

  "github.com/casbin/casbin/v2"
  "github.com/gin-contrib/authz"
  "github.com/gin-gonic/gin"
)

func main() {
  // load the casbin model and policy from files, database is also supported.
  e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

  // define your router, and use the Casbin authz middleware.
  // the access that is denied by authz will return HTTP 403 error.
  router := gin.New()
  router.Use(authz.NewAuthorizer(e))
}
```


<原文结束>

# <翻译开始>
# 简单示例

```Go
package main

import (
  "net/http"
  "github.com/casbin/casbin/v2"
  "github.com/gin-contrib/authz"
  "github.com/gin-gonic/gin"
)

func main() {
// 从文件加载casbin模型和策略，也支持从数据库加载。
  e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

// 定义你的路由，并使用Casbin授权中间件。
// 被authz拒绝的访问将返回HTTP 403错误。
  router := gin.New()
  router.Use(authz.NewAuthorizer(e))
}
```

# <翻译结束>


<原文开始>
Documentation

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user name
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"

For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).


<原文结束>

# <翻译开始>
# 文档

授权基于``{主体, 对象, 动作}``来判断请求，即确定“主体”可以在“对象”上执行何种“动作”。在这个插件中，各部分含义如下：

1. ``主体``：已登录的用户名
2. ``对象``：Web资源对应的URL路径，如"dataset1/item1"
3. ``动作``：HTTP方法，如GET、POST、PUT、DELETE，或者是您定义的高级别操作，如"read-file"、"write-blog"

关于如何编写授权策略及其他详细信息，请参阅[Casbin的文档](https://github.com/casbin/casbin)。

# <翻译结束>


<原文开始>
Getting Help

- [Casbin](https://github.com/casbin/casbin)


<原文结束>

# <翻译开始>
# 获取帮助

- [Casbin](https://github.com/casbin/casbin)

# <翻译结束>


<原文开始>
License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.

<原文结束>

# <翻译开始>
# 许可证

本项目遵循MIT许可协议。有关完整许可文本，请参阅[LICENSE](LICENSE)文件。

# <翻译结束>

