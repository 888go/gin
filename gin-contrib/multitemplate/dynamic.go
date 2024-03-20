package multitemplate

import (
	"fmt"
	"html/template"
	"path/filepath"
	
	"github.com/888go/gin"
	"github.com/888go/gin/render"
)

// DynamicRender type
type DynamicRender map[string]*templateBuilder

var (
	_ render.HTMLRender = DynamicRender{}
	_ Renderer          = DynamicRender{}
)

// NewDynamic 是用于创建动态模板的构造函数
func NewDynamic() DynamicRender {
	return make(DynamicRender)
}

// NewRenderer 允许创建一个基于启用的 gin 模式无关的多模板渲染器
func NewRenderer() Renderer {
	if gin.IsDebugging() {
		return NewDynamic()
	}
	return New()
}

// 动态构建器的类型
type builderType int

// 动态构建器类型
const (
	templateType builderType = iota
	filesTemplateType
	globTemplateType
	stringTemplateType
	stringFuncTemplateType
	filesFuncTemplateType
)

// 动态模板构建器
type templateBuilder struct {
	buildType       builderType
	tmpl            *template.Template
	templateName    string
	files           []string
	glob            string
	templateString  string
	funcMap         template.FuncMap
	templateStrings []string
}

func (tb templateBuilder) buildTemplate() *template.Template {
	switch tb.buildType {
	case templateType:
		return tb.tmpl
	case filesTemplateType:
		return template.Must(template.ParseFiles(tb.files...))
	case globTemplateType:
		return template.Must(template.ParseGlob(tb.glob))
	case stringTemplateType:
		return template.Must(template.New(tb.templateName).Parse(tb.templateString))
	case stringFuncTemplateType:
		tmpl := template.New(tb.templateName).Funcs(tb.funcMap)
		for _, ts := range tb.templateStrings {
			tmpl = template.Must(tmpl.Parse(ts))
		}
		return tmpl
	case filesFuncTemplateType:
		return template.Must(template.New(tb.templateName).Funcs(tb.funcMap).ParseFiles(tb.files...))
	default:
		panic("Invalid builder type for dynamic template")
	}
}

// Add new template
func (r DynamicRender) Add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template cannot be nil")
	}
	if len(name) == 0 {
		panic("template name cannot be empty")
	}
	builder := &templateBuilder{templateName: name, tmpl: tmpl}
	builder.buildType = templateType
	r[name] = builder
}

// AddFromFiles 从文件中加载并添加模板
func (r DynamicRender) AddFromFiles(name string, files ...string) *template.Template {
	builder := &templateBuilder{templateName: name, files: files}
	builder.buildType = filesTemplateType
	r[name] = builder
	return builder.buildTemplate()
}

// AddFromGlob 从全局路径提供添加模板的功能
func (r DynamicRender) AddFromGlob(name, glob string) *template.Template {
	builder := &templateBuilder{templateName: name, glob: glob}
	builder.buildType = globTemplateType
	r[name] = builder
	return builder.buildTemplate()
}

// AddFromString 从字符串中提供添加模板
func (r DynamicRender) AddFromString(name, templateString string) *template.Template {
	builder := &templateBuilder{templateName: name, templateString: templateString}
	builder.buildType = stringTemplateType
	r[name] = builder
	return builder.buildTemplate()
}

// AddFromStringsFuncs 从字符串提供添加模板功能
func (r DynamicRender) AddFromStringsFuncs(name string, funcMap template.FuncMap, templateStrings ...string) *template.Template {
	builder := &templateBuilder{
		templateName: name, funcMap: funcMap,
		templateStrings: templateStrings,
	}
	builder.buildType = stringFuncTemplateType
	r[name] = builder
	return builder.buildTemplate()
}

// AddFromFilesFuncs 用于提供从文件添加模板的回调函数
func (r DynamicRender) AddFromFilesFuncs(name string, funcMap template.FuncMap, files ...string) *template.Template {
	tname := filepath.Base(files[0])
	builder := &templateBuilder{templateName: tname, funcMap: funcMap, files: files}
	builder.buildType = filesFuncTemplateType
	r[name] = builder
	return builder.buildTemplate()
}

// 实例提供渲染字符串
func (r DynamicRender) Instance(name string, data interface{}) render.Render {
	builder, ok := r[name]
	if !ok {
		panic(fmt.Sprintf("Dynamic template with name %s not found", name))
	}
	return render.HTML{
		Template: builder.buildTemplate(),
		Data:     data,
	}
}
