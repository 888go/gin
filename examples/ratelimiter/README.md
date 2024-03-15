# # 在 Gin 中实现速率限制

该项目是一个使用漏桶算法实现速率限制的示例。尽管 Golang 官方包提供了一个基于令牌桶算法的实现 [time/rate](https://pkg.go.dev/golang.org/x/time/rate?tab=doc)，

但你也可以通过 Gin 的函数式 `Use()` 方法来集成自定义中间件，以实现额外的速率限制功能。
## Effect

```go
// You can assign the ratelimit of the server
// rps: requests per second
go run rate.go -rps=100
```

- Let's hava a simple test by ab with 3000 mock requests, not surprisingly，it will takes 10ms each request.

```bash
ab -n 3000 -c 1 http://localhost:8080/rate
```

- Gin Log Output

```bash
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    8.307734ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |   10.512913ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |     8.54681ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    8.356436ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    9.677276ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    7.536156ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    11.57084ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |       7.802ms |       127.0.0.1 | GET      /rate
[GIN] 10ms
[GIN] 2020/07/14 - 15:07:49 | 200 |    9.602394ms |       127.0.0.1 | GET      /rate
```

- AB Test Reporter

```java
Concurrency Level:      1
Time taken for tests:   30.00 seconds
Complete requests:      3000
Requests per second:    100.00 [#/sec] (mean)
Time per request:       10.001 [ms] (mean)
Time per request:       10.001 [ms] (mean, across all concurrent requests)
```
