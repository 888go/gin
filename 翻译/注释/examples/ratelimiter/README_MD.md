
<原文开始>
Ratelimit in Gin

This project is a sample for ratelimit using Leaky Bucket. Although the golang official pkg provide a implement with Token Bucket [time/rate](https://pkg.go.dev/golang.org/x/time/rate?tab=doc),

you can also make your owns via gin's functional `Use()` to integrate extra middlewares.


<原文结束>

# <翻译开始>
# 在 Gin 中实现速率限制

该项目是一个使用漏桶算法实现速率限制的示例。尽管 Golang 官方包提供了一个基于令牌桶算法的实现 [time/rate](https://pkg.go.dev/golang.org/x/time/rate?tab=doc)，

但你也可以通过 Gin 的函数式 `Use()` 方法来集成自定义中间件，以实现额外的速率限制功能。

# <翻译结束>

