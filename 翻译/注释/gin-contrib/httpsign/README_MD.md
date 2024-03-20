
<原文开始>
httpsign

[![Run Tests](https://github.com/gin-contrib/httpsign/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/httpsign/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/httpsign/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/httpsign)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/httpsign)](https://goreportcard.com/report/github.com/gin-contrib/httpsign)
[![GoDoc](https://godoc.org/github.com/gin-contrib/httpsign?status.svg)](https://godoc.org/github.com/gin-contrib/httpsign)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Signing HTTP Messages Middleware base on [HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures).


<原文结束>

# <翻译开始>
# httpsign

[![运行测试](https://github.com/gin-contrib/httpsign/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/httpsign/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gin-contrib/httpsign/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/httpsign)
[![Go 代码质量报告](https://goreportcard.com/badge/github.com/gin-contrib/httpsign)](https://goreportcard.com/report/github.com/gin-contrib/httpsign)
[![GoDoc 文档](https://godoc.org/github.com/gin-contrib/httpsign?status.svg)](https://godoc.org/github.com/gin-contrib/httpsign)
[![加入聊天室](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

基于[HTTP Signatures](https://tools.ietf.org/html/draft-cavage-http-signatures)的HTTP消息签名中间件。

# <翻译结束>


<原文开始>
Example

``` go

package main

import (
  "github.com/gin-contrib/httpsign"
  "github.com/gin-contrib/httpsign/crypto"
  "github.com/gin-gonic/gin"
)

func main() {
  // Define algorithm
  hmacsha256 := &crypto.HmacSha256{}
  hmacsha512 := &crypto.HmacSha512{}
  // Init define secret params
  readKeyID := httpsign.KeyID("read")
  writeKeyID := httpsign.KeyID("write")
  secrets := httpsign.Secrets{
    readKeyID: &httpsign.Secret{
      Key:       "HMACSHA256-SecretKey",
      Algorithm: hmacsha256, // You could using other algo with interface Crypto
    },
    writeKeyID: &httpsign.Secret{
      Key:       "HMACSHA512-SecretKey",
      Algorithm: hmacsha512,
    },
  }

  // Init server
  r := gin.Default()

  //Create middleware with default rule. Could modify by parse Option func
  auth := httpsign.NewAuthenticator(secrets)

  r.Use(auth.Authenticated())
  r.GET("/a", a)
  r.POST("/b", b)
  r.POST("/c", c)

  r.Run(":8080")
}
```

<原文结束>

# <翻译开始>
# 示例

``` go
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
      Key:       "HMACSHA256-SecretKey", // HMACSHA256密钥
      Algorithm: hmacsha256,            // 可以使用其他符合Crypto接口的算法
    },
    writeKeyID: &httpsign.Secret{
      Key:       "HMACSHA512-SecretKey", // HMACSHA512密钥
      Algorithm: hmacsha512,
    },
  }

// 初始化服务器
  r := gin.Default()

// 创建中间件，采用默认规则。可以通过解析Option函数进行修改
  auth := httpsign.NewAuthenticator(secrets)

  r.Use(auth.Authenticated())
  r.GET("/a", a)
  r.POST("/b", b)
  r.POST("/c", c)

  r.Run(":8080") // 运行服务器在8080端口
}
```

这段代码是Go语言的一个示例，使用gin框架和httpsign库实现HTTP请求签名验证功能。首先定义了两种哈希算法（HMACSHA256和HMACSHA512）以及对应的密钥标识符（"read" 和 "write"），并设置了相应的密钥和算法。接着初始化了一个gin服务器，并创建一个基于预定义密钥和算法的认证中间件。最后将该中间件应用于特定路由，并运行服务器在8080端口监听请求。

# <翻译结束>

