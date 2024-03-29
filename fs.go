// 版权所有 ? 2017 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证协议约束，
// 该协议可在 LICENSE 文件中查阅。

package gin类

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

// Dir 返回一个可用于 http.FileServer() 的 http.FileSystem。它在 router.Static() 内部使用。
// 如果 listDirectory 为 true，则其行为与 http.Dir() 相同；否则，它将返回一个文件系统，
// 阻止 http.FileServer() 列出目录中的文件。
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyFilesFS{fs}
}

// Open 符合 http.Filesystem 接口。
//
// 注意!!! 此方法不能翻译, 因为是http包的接口实现
func (fs onlyFilesFS) Open(名称 string) (http.File, error) {
	f, err := fs.fs.Open(名称)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir 重写（覆盖）了 http.File 的默认实现。
func (f neuteredReaddirFile) Readdir(_ int) ([]os.FileInfo, error) {
	// 这将禁用目录列表
	return nil, nil
}
