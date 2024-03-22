package multitemplate

import (
	"html/template"
	"testing"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func createFromFileDynamic() Renderer {
	r := NewRenderer()
	r.AddFromFiles("index", "tests/base.html", "tests/article.html")

	return r
}

func createFromGlobDynamic() Renderer {
	r := NewRenderer()
	r.AddFromGlob("index", "tests/global/*")

	return r
}

func createFromStringDynamic() Renderer {
	r := NewRenderer()
	r.AddFromString("index", "Welcome to {{ .name }} template")

	return r
}

func createFromStringsWithFuncsDynamic() Renderer {
	r := NewRenderer()
	r.AddFromStringsFuncs("index", template.FuncMap{}, `Welcome to {{ .name }} {{template "content"}}`, `{{define "content"}}template{{end}}`)

	return r
}

func createFromFilesWithFuncsDynamic() Renderer {
	r := NewRenderer()
	r.AddFromFilesFuncs("index", template.FuncMap{}, "tests/welcome.html", "tests/content.html")

	return r
}

func createFromTemplatesDynamic() Renderer {
	tmpl := template.Must(template.New("test").Parse("Welcome to {{ .name }} template"))
	r := NewRenderer()
	r.Add("test", tmpl)
	return r
}

func TestMissingTemplateOrNameDynamic(t *testing.T) {
	r := NewRenderer()
	tmpl := template.Must(template.New("test").Parse("Welcome to {{ .name }} template"))
	assert.Panics(t, func() {
		r.Add("", tmpl)
	}, "template name cannot be empty")

	assert.Panics(t, func() {
		r.Add("test", nil)
	}, "template can not be nil")
}

func TestAddFromFilesDynamic(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromFileDynamic()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"title": "Test Multiple Template",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is article template\n", w.Body.String())
}

func TestAddFromGlobDynamic(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromGlobDynamic()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"title": "Test Multiple Template",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is login template\n", w.Body.String())
}

func TestAddFromStringDynamic(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromStringDynamic()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}

func TestAddFromStringsFruncsDynamic(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromStringsWithFuncsDynamic()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}

func TestAddFromFilesFruncsDynamic(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromFilesWithFuncsDynamic()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template\n", w.Body.String())
}

func TestPanicInvalidTypeBuilder(t *testing.T) {
	assert.Panics(t, func() {
		b := templateBuilder{}
		b.buildType = 10
		b.buildTemplate()
	})
}

func TestTemplateNotFound(t *testing.T) {
	r := make(DynamicRender)
	r.AddFromString("index", "This is a test template")
	assert.Panics(t, func() {
		r.Instance("NotFoundTemplate", nil)
	})
}

func TestNotDynamicMode(t *testing.T) {
	gin类.X设置运行模式("test")
	TestAddFromFilesDynamic(t)
	gin类.X设置运行模式("debug")
}

func TestAddTemplate(t *testing.T) {
	tmpl := template.Must(template.ParseFiles("tests/base.html", "tests/article.html"))
	b := templateBuilder{}
	b.buildType = templateType
	b.tmpl = tmpl
	b.buildTemplate()
	assert.NotPanics(t, func() {
		b.buildTemplate()
	})
}

func TestAddingTemplate(t *testing.T) {
	assert.NotPanics(t, func() {
		createFromTemplatesDynamic()
	})
}
