# # httpsign

[![运行测试](https://github.com/gin-contrib/httpsign/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/httpsign/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/httpsign/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/httpsign)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/httpsign)](https://goreportcard.com/report/github.com/gin-contrib/httpsign)
[![GoDoc](https://godoc.org/github.com/gin-contrib/httpsign?status.svg)](https://godoc.org/github.com/gin-contrib/httpsign)
[![加入 Gitter 聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

基于 [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures) 的 HTTP 消息签名中间件。
## # 示例

```go
package main

import (
    "github.com/gin-contrib/httpsign"
    "github.com/gin-contrib/httpsign/crypto"
    "github.com/gin-gonic/gin"
)

func main() {
    // 定义算法
    hmacsha256 := &crypto.HmacSha256{}
    hmacsha512 := &crypto.HmacSha512{}

    // 初始化密钥参数
    readKeyID := httpsign.KeyID("read")
    writeKeyID := httpsign.KeyID("write")
    secrets := httpsign.Secrets{
        readKeyID: &httpsign.Secret{
            Key:       "HMACSHA256-SecretKey",
            Algorithm: hmacsha256, // 可以使用其他实现Crypto接口的算法
        },
        writeKeyID: &httpsign.Secret{
            Key:       "HMACSHA512-SecretKey",
            Algorithm: hmacsha512,
        },
    }

    // 初始化服务器
    r := gin.Default()

    // 创建中间件，使用默认规则。可以通过解析Option函数进行修改
    auth := httpsign.NewAuthenticator(secrets)

    r.Use(auth.Authenticated())
    r.GET("/a", a)
    r.POST("/b", b)
    r.POST("/c", c)

    r.Run(":8080")
}
```

这段代码是一个Go语言示例，它定义了一个基于gin框架的HTTP服务器，该服务器使用了`httpsign`库来处理HTTP签名认证。首先导入所需包，并定义了两种哈希算法（HMAC SHA256和HMAC SHA512）。接着初始化了两个密钥ID以及对应的密钥和加密算法。然后创建了一个gin实例并配置了一个基于已定义密钥和算法的认证中间件。最后为三个路由分别设置了GET和POST方法，并启动服务器监听8080端口。
