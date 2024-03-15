# # Gin websocket 客户端和服务端示例

这个示例展示了一个简单的客户端和服务器。

服务器会回显发送给它的消息。客户端每秒发送一条消息，并打印接收到的所有消息。

要运行此示例，请先启动服务器：

```bash
go run server/server.go
```

接下来，启动客户端：

```bash
go run client/client.go
```

服务器包含一个简易的网页客户端。要使用该客户端，请在浏览器中打开 [URL](http://127.0.0.1:8080) 并按照页面上的指示操作。
## # deps

- [gorilla/websocket](https://github.com/gorilla/websocket) (依赖项)

 Gorilla/websocket (翻译：大猩猩/websocket)
