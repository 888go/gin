package main

import (
	"log"
	"net/http"
	"strings"
	
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/888go/gin/gin-contrib/static"
	"github.com/888go/gin"
)

type binaryFileSystem struct {
	fs http.FileSystem
}


// ff:
// http.File:
// name:
func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}


// ff:
// filepath:
// prefix:
func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}


// ff:
// root:
func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
	return &binaryFileSystem{
		fs,
	}
}

// 使用方法
// $ go-bindata data/
// $ go build && ./bindata
// 
// 这段注释的中文翻译是：
// 
// 用法
// $ 运行命令 go-bindata 并指定数据目录：data/
// $ 执行构建命令 go build，然后运行生成的可执行文件 ./bindata
func main() {
	r := gin.Default()

	r.Use(static.Serve("/static", BinaryFileSystem("data")))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	// 在0.0.0.0:8080监听并服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
