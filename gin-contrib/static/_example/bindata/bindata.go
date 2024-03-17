package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func data_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x44, 0xce,
		0xbd, 0xae, 0xc2, 0x30, 0x0c, 0x05, 0xe0, 0xbd, 0x4f, 0xe1, 0x9b, 0xbd,
		0xaa, 0xba, 0xdd, 0x21, 0xcd, 0xc2, 0xef, 0x06, 0x43, 0x19, 0x18, 0x5d,
		0x62, 0x35, 0x41, 0x4e, 0x22, 0x15, 0x4b, 0x88, 0xb7, 0x27, 0x21, 0x45,
		0x4c, 0x39, 0xb1, 0xf5, 0x1d, 0x59, 0xff, 0x6d, 0x4f, 0x9b, 0xf1, 0x7a,
		0xde, 0x81, 0x93, 0xc0, 0xa6, 0xd1, 0xe5, 0x01, 0xc6, 0x38, 0x0f, 0xea,
		0x8e, 0xca, 0x34, 0x00, 0xda, 0x11, 0xda, 0x12, 0x72, 0x0c, 0x24, 0x08,
		0x37, 0x87, 0xcb, 0x83, 0x64, 0x50, 0x97, 0x71, 0xdf, 0xfe, 0x2b, 0xe8,
		0xd6, 0xa5, 0x78, 0x61, 0x32, 0x73, 0x6a, 0x27, 0x1f, 0x2d, 0x0a, 0xea,
		0xae, 0x4e, 0x4a, 0x47, 0xf7, 0x2d, 0xd1, 0x53, 0xb2, 0xaf, 0x15, 0xb8,
		0xde, 0x1c, 0x89, 0x39, 0xc1, 0xc1, 0x47, 0xf8, 0x39, 0x08, 0xde, 0x5a,
		0xa6, 0x27, 0x2e, 0x94, 0x5d, 0x5f, 0x7d, 0x65, 0xf9, 0xff, 0x39, 0xf3,
		0x1d, 0x00, 0x00, 0xff, 0xff, 0x51, 0x69, 0x85, 0x27, 0xb7, 0x00, 0x00,
		0x00,
	},
		"data/index.html",
	)
}

// Asset 函数加载并返回指定名称的资源。
// 如果无法找到该资源或无法加载，则返回错误。

// ff:
// name:

// ff:
// name:
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames 返回资产的名称列表。

// ff:

// ff:
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata 是一个表，用于存储每个资产生成器，并将其映射到对应的名称。
var _bindata = map[string]func() ([]byte, error){
	"data/index.html": data_index_html,
}

// AssetDir 返回在由 go-bindata 嵌入到文件中的特定目录下的文件名。
// 例如，如果你运行 go-bindata 对 data/... 进行处理，且 data 包含以下层次结构：
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// 那么 AssetDir("data") 将返回 []string{"foo.txt", "img"}
// AssetDir("data/img") 将返回 []string{"a.png", "b.png"}
// 而 AssetDir("foo.txt") 和 AssetDir("notexist") 则会返回错误

// ff:
// name:

// ff:
// name:
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"data": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{data_index_html, map[string]*_bintree_t{}},
	}},
}}

// AssetInfo 返回指定路径的文件信息

// ff:
// os.FileInfo:
// path:

// ff:
// os.FileInfo:
// path:
func AssetInfo(path string) (os.FileInfo, error) {
	return os.Stat(path)
}
