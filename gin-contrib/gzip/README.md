# # GZIP中间件（用于gin）

[![运行测试](https://github.com/gin-contrib/gzip/actions/workflows/go.yml/badge.svg)](https://github.com/gin-contrib/gzip/actions/workflows/go.yml)
[![codecov代码覆盖率](https://codecov.io/gh/gin-contrib/gzip/branch/master/graph/badge.svg)](https://codecov.io/gh/gin-contrib/gzip)
[![Go 语言报告卡](https://goreportcard.com/badge/github.com/gin-contrib/gzip)](https://goreportcard.com/report/github.com/gin-contrib/gzip)
[![GoDoc文档](https://godoc.org/github.com/gin-contrib/gzip?status.svg)](https://godoc.org/github.com/gin-contrib/gzip)
[![加入Gitter聊天](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

这是一个用于gin的中间件，用于启用`GZIP`支持。
## Usage

Download and install it:

```sh
go get github.com/gin-contrib/gzip
```

Import it in your code:

```go
import "github.com/gin-contrib/gzip"
```

Canonical example:

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/gin-contrib/gzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Extensions

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/gin-contrib/gzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Paths

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/gin-contrib/gzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Paths

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/gin-contrib/gzip"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```
