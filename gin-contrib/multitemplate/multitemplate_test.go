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


// ff:
// t:

// ff:
// t:
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


// ff:
// t:

// ff:
// t:
func TestAddFromFiles(t *testing.T) {
	router := gin.New()
	router.HTMLRender = createFromFile()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Test Multiple Template",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is article template\n", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestAddFromGlob(t *testing.T) {
	router := gin.New()
	router.HTMLRender = createFromGlob()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Test Multiple Template",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<p>Test Multiple Template</p>\nHi, this is login template\n", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestAddFromString(t *testing.T) {
	router := gin.New()
	router.HTMLRender = createFromString()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestAddFromStringsFruncs(t *testing.T) {
	router := gin.New()
	router.HTMLRender = createFromStringsWithFuncs()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestAddFromFilesFruncs(t *testing.T) {
	router := gin.New()
	router.HTMLRender = createFromFilesWithFuncs()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"name": "index",
		})
	})

	w := performRequest(router, "GET", "/")
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Welcome to index template\n", w.Body.String())
}


// ff:
// t:

// ff:
// t:
func TestDuplicateTemplate(t *testing.T) {
	assert.Panics(t, func() {
		r := New()
		r.AddFromString("index", "Welcome to {{ .name }} template")
		r.AddFromString("index", "Welcome to {{ .name }} template")
	})
}
