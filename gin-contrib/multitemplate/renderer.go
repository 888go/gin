package multitemplate

import (
	"html/template"
	
	"github.com/888go/gin/render"
)

// Renderer 类型是用于多模板的通用渲染器。
// 当 gin 处于调试模式时，所有多模板功能支持热重载，这意味着您可以即时修改文件模板并查看更改。
// Renderer 应该使用 multitemplate.NewRenderer() 构造函数进行创建。
type Renderer interface {
	render.HTMLRender
	Add(name string, tmpl *template.Template)
	AddFromFiles(name string, files ...string) *template.Template
	AddFromGlob(name, glob string) *template.Template
	AddFromString(name, templateString string) *template.Template
	AddFromStringsFuncs(name string, funcMap template.FuncMap, templateStrings ...string) *template.Template
	AddFromFilesFuncs(name string, funcMap template.FuncMap, files ...string) *template.Template
}
