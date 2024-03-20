# # 在本地开发服务器上运行Gin的指南（App Engine环境）

1. 在您的计算机上下载、安装并设置Go环境。（这包括设置您的`$GOPATH`。）
2. 从[这里](https://cloud.google.com/appengine/docs/standard/go/download)为您的平台下载SDK：`https://cloud.google.com/appengine/docs/standard/go/download`
3. 使用以下命令下载Gin源代码：`$ go get github.com/gin-gonic/examples`
4. 导航到examples文件夹：`$ cd $GOPATH/src/github.com/gin-gonic/examples/app-engine/`
5. 运行它：`$ dev_appserver.py .`（请注意，您必须使用Python2来运行这个脚本）
