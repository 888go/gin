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

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
	return &binaryFileSystem{
		fs,
	}
}

// Usage
// $ go-bindata data/
// $ go build && ./bindata
//
func main() {
	r := gin类.X创建默认对象()

	r.X中间件(static.Serve("/static", BinaryFileSystem("data")))
	r.X绑定GET("/ping", func(c *gin类.Context) {
		c.X输出文本(200, "test")
	})
	// Listen and Server in 0.0.0.0:8080
	if err := r.X监听(":8080"); err != nil {
		log.Fatal(err)
	}
}
