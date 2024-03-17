package gzip

import (
	"compress/gzip"
	"net/http"
	"regexp"
	"strings"
	
	"github.com/888go/gin"
)

var (
	DefaultExcludedExtentions = NewExcludedExtensions([]string{
		".png", ".gif", ".jpeg", ".jpg",
	})
	DefaultOptions = &Options{
		ExcludedExtensions: DefaultExcludedExtentions,
	}
)

type Options struct {
	ExcludedExtensions   ExcludedExtensions
	ExcludedPaths        ExcludedPaths
	ExcludedPathesRegexs ExcludedPathesRegexs
	DecompressFn         func(c *gin.Context)
}

type Option func(*Options)


// ff:
// args:

// ff:
// args:
func WithExcludedExtensions(args []string) Option {
	return func(o *Options) {
		o.ExcludedExtensions = NewExcludedExtensions(args)
	}
}


// ff:
// args:

// ff:
// args:
func WithExcludedPaths(args []string) Option {
	return func(o *Options) {
		o.ExcludedPaths = NewExcludedPaths(args)
	}
}


// ff:
// args:

// ff:
// args:
func WithExcludedPathsRegexs(args []string) Option {
	return func(o *Options) {
		o.ExcludedPathesRegexs = NewExcludedPathesRegexs(args)
	}
}


// ff:
// decompressFn:
// c:

// ff:
// decompressFn:
// c:
func WithDecompressFn(decompressFn func(c *gin.Context)) Option {
	return func(o *Options) {
		o.DecompressFn = decompressFn
	}
}

// 使用map以获得更好的查找性能
type ExcludedExtensions map[string]bool


// ff:
// extensions:

// ff:
// extensions:
func NewExcludedExtensions(extensions []string) ExcludedExtensions {
	res := make(ExcludedExtensions)
	for _, e := range extensions {
		res[e] = true
	}
	return res
}


// ff:
// target:

// ff:
// target:
func (e ExcludedExtensions) Contains(target string) bool {
	_, ok := e[target]
	return ok
}

type ExcludedPaths []string


// ff:
// paths:

// ff:
// paths:
func NewExcludedPaths(paths []string) ExcludedPaths {
	return ExcludedPaths(paths)
}


// ff:
// requestURI:

// ff:
// requestURI:
func (e ExcludedPaths) Contains(requestURI string) bool {
	for _, path := range e {
		if strings.HasPrefix(requestURI, path) {
			return true
		}
	}
	return false
}

type ExcludedPathesRegexs []*regexp.Regexp


// ff:
// regexs:

// ff:
// regexs:
func NewExcludedPathesRegexs(regexs []string) ExcludedPathesRegexs {
	result := make([]*regexp.Regexp, len(regexs))
	for i, reg := range regexs {
		result[i] = regexp.MustCompile(reg)
	}
	return result
}


// ff:
// requestURI:

// ff:
// requestURI:
func (e ExcludedPathesRegexs) Contains(requestURI string) bool {
	for _, reg := range e {
		if reg.MatchString(requestURI) {
			return true
		}
	}
	return false
}


// ff:
// c:

// ff:
// c:
func DefaultDecompressHandle(c *gin.Context) {
	if c.Request.Body == nil {
		return
	}
	r, err := gzip.NewReader(c.Request.Body)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Request.Header.Del("Content-Encoding")
	c.Request.Header.Del("Content-Length")
	c.Request.Body = r
}
