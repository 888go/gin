# # Authz

[![CodeQL](https://github.com/gin-contrib/authz/actions/workflows/codeql.yml/badge.svg)](https://github.com/gin-contrib/authz/actions/workflows/codeql.yml)
[![运行测试](https://github.com/gin-contrib/authz/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gin-contrib/authz/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/authz/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/authz)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/authz)](https://goreportcard.com/report/github.com/gin-contrib/authz)
[![GoDoc](https://godoc.org/github.com/gin-contrib/authz?status.svg)](https://godoc.org/github.com/gin-contrib/authz)
[![加入Gitter聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Authz 是 Gin 框架的一款授权中间件，基于 [https://github.com/casbin/casbin](https://github.com/casbin/casbin) 开发。
## # 安装

```bash
go get github.com/gin-contrib/authz
```
## # 简单示例

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

// 定义你的路由，并使用Casbin权限认证中间件。
// 被authz拒绝的访问将会返回HTTP 403错误。
  router := gin.New()
  router.Use(authz.NewAuthorizer(e))
}
```
## # 文档

授权基于``{主体, 对象, 动作}``来决定一个请求，即确定某个``主体``能够在什么``对象``上执行什么``动作``。在这个插件中，它们的具体含义如下：

1. ``主体``：已登录的用户名
2. ``对象``：Web资源对应的URL路径，如"dataset1/item1"
3. ``动作``：HTTP方法，如GET、POST、PUT、DELETE，或者是您定义的高级别动作，如"读取文件"、"写博客"

有关如何编写授权策略及其他详细信息，请参考[Casbin的文档](https://github.com/casbin/casbin)。
## # 获取帮助

- [Casbin](https://github.com/casbin/casbin)
## # 许可证

本项目采用 MIT 许可证。您可以在 [LICENSE](LICENSE) 文件中查看完整的许可证文本。
