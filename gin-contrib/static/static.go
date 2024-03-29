package static

import (
	"net/http"
	"os"
	"path"
	"strings"
	
	"github.com/888go/gin"
)

const INDEX = "index.html"

type ServeFileSystem interface {
	http.FileSystem
	Exists(prefix string, path string) bool
}

type localFileSystem struct {
	http.FileSystem
	root    string
	indexes bool
}

func LocalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin类.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}

func (l *localFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		name := path.Join(l.root, p)
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			if !l.indexes {
				index := path.Join(name, INDEX)
				_, err := os.Stat(index)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

func ServeRoot(urlPrefix, root string) gin类.HandlerFunc {
	return Serve(urlPrefix, LocalFile(root, false))
}

// Static 返回一个中间件处理器，用于在指定目录下提供静态文件服务。
func Serve(urlPrefix string, fs ServeFileSystem) gin类.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin类.Context) {
		if fs.Exists(urlPrefix, c.X请求.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.X请求)
			c.X停止()
		}
	}
}
