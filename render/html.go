// 版权声明 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package render

import (
	"html/template"
	"net/http"
)

// Delims 表示用于 HTML 模板渲染的一组左（Left）和右（Right）定界符。
type Delims struct {
// 左侧分隔符，默认为 {{.
	Left string
// 右侧分隔符，默认为 }}。
	Right string
}

// HTMLRender 接口需要由 HTMLProduction 和 HTMLDebug 实现。
type HTMLRender interface {
// Instance 返回一个HTML实例。
	Instance(string, any) Render
}

// HTMLProduction 包含模板引用及其分隔符。
type HTMLProduction struct {
	Template *template.Template
	Delims   Delims
}

// HTMLDebug 包含模板分隔符、模式以及带有文件列表的函数，主要用于调试HTML。
type HTMLDebug struct {
	Files   []string
	Glob    string
	Delims  Delims
	FuncMap template.FuncMap
}

// HTML 包含模板引用及其名称，以及给定的接口对象。
type HTML struct {
	Template *template.Template
	Name     string
	Data     any
}

var htmlContentType = []string{"text/html; charset=utf-8"}

// Instance (HTMLProduction) 返回一个实现了Render接口的HTML实例。
func (r HTMLProduction) Instance(name string, data any) Render {
	return HTML{
		Template: r.Template,
		Name:     name,
		Data:     data,
	}
}

// Instance (HTMLDebug) 返回一个实现了Render接口的HTML实例。
func (r HTMLDebug) Instance(name string, data any) Render {
	return HTML{
		Template: r.loadTemplate(),
		Name:     name,
		Data:     data,
	}
}
func (r HTMLDebug) loadTemplate() *template.Template {
	if r.FuncMap == nil {
		r.FuncMap = template.FuncMap{}
	}
	if len(r.Files) > 0 {
		return template.Must(template.New("").Delims(r.Delims.Left, r.Delims.Right).Funcs(r.FuncMap).ParseFiles(r.Files...))
	}
	if r.Glob != "" {
		return template.Must(template.New("").Delims(r.Delims.Left, r.Delims.Right).Funcs(r.FuncMap).ParseGlob(r.Glob))
	}
	panic("the HTML debug render was created without files or glob pattern")
}

// Render (HTML) 执行模板并使用自定义 ContentType 将其结果写入响应。
func (r HTML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	if r.Name == "" {
		return r.Template.Execute(w, r.Data)
	}
	return r.Template.ExecuteTemplate(w, r.Name, r.Data)
}

// WriteContentType (HTML) 写入 HTML ContentType。
func (r HTML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, htmlContentType)
}
