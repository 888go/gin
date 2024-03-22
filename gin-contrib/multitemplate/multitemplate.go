package multitemplate

import (
	"fmt"
	"html/template"
	"path/filepath"
	
	"github.com/888go/gin/render"
)

// Render type
type Render map[string]*template.Template

var (
	_ render.HTMLRender = Render{}
	_ Renderer          = Render{}
)

// New instance
func New() Render {
	return make(Render)
}

// Add new template
func (r Render) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	if _, ok := r[name]; ok {
		panic(fmt.Sprintf("template %s already exists", name))
	}
	r[name] = tmpl
}

// AddFromFiles 从文件中加载并添加模板
func (r Render) AddFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

// AddFromGlob 从全局路径提供添加模板的功能
func (r Render) AddFromGlob(name, glob string) *template.Template {
	tmpl := template.Must(template.ParseGlob(glob))
	r.Add(name, tmpl)
	return tmpl
}

// AddFromString 从字符串中提供添加模板
func (r Render) AddFromString(name, templateString string) *template.Template {
	tmpl := template.Must(template.New(name).Parse(templateString))
	r.Add(name, tmpl)
	return tmpl
}

// AddFromStringsFuncs 从字符串提供添加模板功能
func (r Render) AddFromStringsFuncs(name string, funcMap template.FuncMap, templateStrings ...string) *template.Template {
	tmpl := template.New(name).Funcs(funcMap)

	for _, ts := range templateStrings {
		tmpl = template.Must(tmpl.Parse(ts))
	}

	r.Add(name, tmpl)
	return tmpl
}

// AddFromFilesFuncs 用于提供从文件添加模板的回调函数
func (r Render) AddFromFilesFuncs(name string, funcMap template.FuncMap, files ...string) *template.Template {
	tname := filepath.Base(files[0])
	tmpl := template.Must(template.New(tname).Funcs(funcMap).ParseFiles(files...))
	r.Add(name, tmpl)
	return tmpl
}

// 实例提供渲染字符串
func (r Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r[name],
		Data:     data,
	}
}
