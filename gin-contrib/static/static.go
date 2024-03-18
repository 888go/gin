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


// ff:
// indexes:
// root:

// ff:
// indexes:
// root:

// ff:
// indexes:
// root:

// ff:
// indexes:
// root:

// ff:
// indexes:
// root:

// ff:
// indexes:
// root:

// ff:
// indexes:
// root:
func LocalFile(root string, indexes bool) *localFileSystem {
	return &localFileSystem{
		FileSystem: gin.Dir(root, indexes),
		root:       root,
		indexes:    indexes,
	}
}


// ff:
// filepath:
// prefix:

// ff:
// filepath:
// prefix:

// ff:
// filepath:
// prefix:

// ff:
// filepath:
// prefix:

// ff:
// filepath:
// prefix:

// ff:
// filepath:
// prefix:

// ff:
// filepath:
// prefix:
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


// ff:
// root:
// urlPrefix:

// ff:
// root:
// urlPrefix:

// ff:
// root:
// urlPrefix:

// ff:
// root:
// urlPrefix:

// ff:
// root:
// urlPrefix:

// ff:
// root:
// urlPrefix:

// ff:
// root:
// urlPrefix:
func ServeRoot(urlPrefix, root string) gin.HandlerFunc {
	return Serve(urlPrefix, LocalFile(root, false))
}

// Static 返回一个中间件处理程序，用于在指定目录中提供静态文件服务。

// ff:
// fs:
// urlPrefix:

// ff:
// fs:
// urlPrefix:

// ff:
// fs:
// urlPrefix:

// ff:
// fs:
// urlPrefix:

// ff:
// fs:
// urlPrefix:

// ff:
// fs:
// urlPrefix:

// ff:
// fs:
// urlPrefix:
func Serve(urlPrefix string, fs ServeFileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if fs.Exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
