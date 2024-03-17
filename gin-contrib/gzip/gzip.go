package gzip

import (
	"compress/gzip"
	
	"github.com/888go/gin"
)

const (
	BestCompression    = gzip.BestCompression
	BestSpeed          = gzip.BestSpeed
	DefaultCompression = gzip.DefaultCompression
	NoCompression      = gzip.NoCompression
)


// ff:
// options:
// level:

// ff:
// options:
// level:
func Gzip(level int, options ...Option) gin.HandlerFunc {
	return newGzipHandler(level, options...).Handle
}

type gzipWriter struct {
	gin.ResponseWriter
	writer *gzip.Writer
}


// ff:
// s:

// ff:
// s:
func (g *gzipWriter) WriteString(s string) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write([]byte(s))
}


// ff:
// data:

// ff:
// data:
func (g *gzipWriter) Write(data []byte) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write(data)
}

// 修复：https://github.com/mholt/caddy/issues/38
// （该注释表明该代码片段是为了修复Caddy项目（作者为mholt）在GitHub上的第38号问题。）

// ff:
// code:

// ff:
// code:
func (g *gzipWriter) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}
