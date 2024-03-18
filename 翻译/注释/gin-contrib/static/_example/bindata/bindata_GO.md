
<原文开始>
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
<原文结束>

# <翻译开始>
// Asset 函数加载并返回指定名称的资源。
// 如果无法找到该资源或无法加载，则返回错误。
# <翻译结束>


<原文开始>
// AssetNames returns the names of the assets.
<原文结束>

# <翻译开始>
// AssetNames 返回资产的名称列表。
# <翻译结束>


<原文开始>
// _bindata is a table, holding each asset generator, mapped to its name.
<原文结束>

# <翻译开始>
// _bindata 是一个表，用于存储每个资产生成器，并将其映射到对应的名称。
# <翻译结束>


<原文开始>
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
// AssetInfo returns file info of given path
<原文结束>

# <翻译开始>
// AssetInfo 返回指定路径的文件信息
# <翻译结束>


<原文开始>
// ff:
// os.FileInfo:
// path:
<原文结束>

# <翻译开始>
// ff:
// os.FileInfo:
// path:
# <翻译结束>

