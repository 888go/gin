// 版权所有2017马努·马丁内斯-阿尔梅达
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"net/http"
	"os"
)

type onlyFilesFS struct {
	fs http.FileSystem
}

type neuteredReaddirFile struct {
	http.File
}

// Dir返回一个http
// http.FileServer()可以使用的文件系统
// 它在router.Static()内部使用
// 如果listDirectory == true，那么它的工作方式与http.Dir()相同，否则它返回一个文件系统，阻止http.FileServer()列出目录文件

// ff:
// listDirectory:
// root:
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyFilesFS{fs}
}

// Open符合http.Filesystem

// ff:
// http.File:
// name:
func (fs onlyFilesFS) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir覆盖http
// 文件默认实现

// ff:
// []os.FileInfo:
// _:
func (f neuteredReaddirFile) Readdir(_ int) ([]os.FileInfo, error) {
// 这将禁用目录列表
	return nil, nil
}
