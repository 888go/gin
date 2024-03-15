# # 在本地开发服务器上运行 Gin 的指南（App Engine 本地开发服务器）

1. 在您的计算机上下载、安装并设置 Go 环境（包括设置 `$GOPATH`）。
2. 从[这里](https://cloud.google.com/appengine/docs/standard/go/download)下载适用于您平台的 SDK：`https://cloud.google.com/appengine/docs/standard/go/download`
3. 使用以下命令下载 Gin 源代码：`$ go get github.com/gin-gonic/examples`
4. 导航到 examples 文件夹：`$ cd $GOPATH/src/github.com/gin-gonic/examples/app-engine/`
5. 运行它：`$ dev_appserver.py .` （请注意，您需要使用 Python2 来运行这个脚本）
