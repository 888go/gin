package multitemplate

import (
	"context"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/888go/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequestWithContext(context.Background(), method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func createFromFile() Render {
	r := New()
	r.AddFromFiles("index", "tests/base.html", "tests/article.html")

	return r
}

func createFromGlob() Render {
	r := New()
	r.AddFromGlob("index", "tests/global/*")

	return r
}

func createFromString() Render {
	r := New()
	r.AddFromString("index", "Welcome to {{ .name }} template")

	return r
}

func createFromStringsWithFuncs() Render {
	r := New()
	r.AddFromStringsFuncs("index", template.FuncMap{}, `Welcome to {{ .name }} {{template "content"}}`, `{{define "content"}}template{{end}}`)

	return r
}

func createFromFilesWithFuncs() Render {
	r := New()
	r.AddFromFilesFuncs("index", template.FuncMap{}, "tests/welcome.html", "tests/content.html")

	return r
}

func TestMissingTemplateOrName(t *testing.T) {
	r := New()
	tmpl := template.Must(template.New("test").Parse("Welcome to {{ .name }} template"))
	assert.Panics(t, func() {
		r.Add("", tmpl)
	}, "template name cannot be empty")

	assert.Panics(t, func() {
		r.Add("test", nil)
	}, "template can not be nil")
}

func TestAddFromFiles(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromFile()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"title": "Test Multiple Template",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is article template\n", w.Body.String())
}

func TestAddFromGlob(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromGlob()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"title": "Test Multiple Template",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is login template\n", w.Body.String())
}

func TestAddFromString(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromString()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}

func TestAddFromStringsFruncs(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromStringsWithFuncs()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}

func TestAddFromFilesFruncs(t *testing.T) {
	router := gin类.X创建()
	router.HTMLRender = createFromFilesWithFuncs()
	router.X绑定GET("/", func(c *gin类.Context) {
		c.X输出html模板(200, "index", gin类.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template\n", w.Body.String())
}

func TestDuplicateTemplate(t *testing.T) {
	assert.Panics(t, func() {
		r := New()
		r.AddFromString("index", "Welcome to {{ .name }} template")
		r.AddFromString("index", "Welcome to {{ .name }} template")
	})
}
