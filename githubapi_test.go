// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

type route struct {
	method string
	path   string
}

// 参考链接：GitHub开发者文档v3版本
var githubAPI = []route{
// OAuth授权
	{http.MethodGet, "/authorizations"},
	{http.MethodGet, "/authorizations/:id"},
	{http.MethodPost, "/authorizations"},
// {http
// MethodPut， "/authorizations/clients/:client_id"}， {http;“MethodPatch; /授权/:id"},
	{http.MethodDelete, "/authorizations/:id"},
	{http.MethodGet, "/applications/:client_id/tokens/:access_token"},
	{http.MethodDelete, "/applications/:client_id/tokens"},
	{http.MethodDelete, "/applications/:client_id/tokens/:access_token"},

// 活动
	{http.MethodGet, "/events"},
	{http.MethodGet, "/repos/:owner/:repo/events"},
	{http.MethodGet, "/networks/:owner/:repo/events"},
	{http.MethodGet, "/orgs/:org/events"},
	{http.MethodGet, "/users/:user/received_events"},
	{http.MethodGet, "/users/:user/received_events/public"},
	{http.MethodGet, "/users/:user/events"},
	{http.MethodGet, "/users/:user/events/public"},
	{http.MethodGet, "/users/:user/events/orgs/:org"},
	{http.MethodGet, "/feeds"},
	{http.MethodGet, "/notifications"},
	{http.MethodGet, "/repos/:owner/:repo/notifications"},
	{http.MethodPut, "/notifications"},
	{http.MethodPut, "/repos/:owner/:repo/notifications"},
	{http.MethodGet, "/notifications/threads/:id"},
// {http
// “MethodPatch; / /线程/通知:id"},
	{http.MethodGet, "/notifications/threads/:id/subscription"},
	{http.MethodPut, "/notifications/threads/:id/subscription"},
	{http.MethodDelete, "/notifications/threads/:id/subscription"},
	{http.MethodGet, "/repos/:owner/:repo/stargazers"},
	{http.MethodGet, "/users/:user/starred"},
	{http.MethodGet, "/user/starred"},
	{http.MethodGet, "/user/starred/:owner/:repo"},
	{http.MethodPut, "/user/starred/:owner/:repo"},
	{http.MethodDelete, "/user/starred/:owner/:repo"},
	{http.MethodGet, "/repos/:owner/:repo/subscribers"},
	{http.MethodGet, "/users/:user/subscriptions"},
	{http.MethodGet, "/user/subscriptions"},
	{http.MethodGet, "/repos/:owner/:repo/subscription"},
	{http.MethodPut, "/repos/:owner/:repo/subscription"},
	{http.MethodDelete, "/repos/:owner/:repo/subscription"},
	{http.MethodGet, "/user/subscriptions/:owner/:repo"},
	{http.MethodPut, "/user/subscriptions/:owner/:repo"},
	{http.MethodDelete, "/user/subscriptions/:owner/:repo"},

// 依据
	{http.MethodGet, "/users/:user/gists"},
	{http.MethodGet, "/gists"},
// {http
// MethodGet， "/gist /public"}， {http;MethodGet,“/丰子恺/ starred"},
	{http.MethodGet, "/gists/:id"},
	{http.MethodPost, "/gists"},
// {http
// MethodPatch,“丰子恺/:id"},
	{http.MethodPut, "/gists/:id/star"},
	{http.MethodDelete, "/gists/:id/star"},
	{http.MethodGet, "/gists/:id/star"},
	{http.MethodPost, "/gists/:id/forks"},
	{http.MethodDelete, "/gists/:id"},

// Git数据
	{http.MethodGet, "/repos/:owner/:repo/git/blobs/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/blobs"},
	{http.MethodGet, "/repos/:owner/:repo/git/commits/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/commits"},
// {http
// MethodGet " / /:业主休息,回购:git / refs ref"} / *,
	{http.MethodGet, "/repos/:owner/:repo/git/refs"},
	{http.MethodPost, "/repos/:owner/:repo/git/refs"},
// {http
// MethodPatch， "/repos/:owner/:repo/git/refs/* refquot;}， {http;MethodDelete,“/回购:所有者/:回购/ git / refs / * ref"},
	{http.MethodGet, "/repos/:owner/:repo/git/tags/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/tags"},
	{http.MethodGet, "/repos/:owner/:repo/git/trees/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/trees"},

// 问题
	{http.MethodGet, "/issues"},
	{http.MethodGet, "/user/issues"},
	{http.MethodGet, "/orgs/:org/issues"},
	{http.MethodGet, "/repos/:owner/:repo/issues"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number"},
	{http.MethodPost, "/repos/:owner/:repo/issues"},
// {http
// MethodPatch,“/回购:所有者/:回购/问题/:number"},
	{http.MethodGet, "/repos/:owner/:repo/assignees"},
	{http.MethodGet, "/repos/:owner/:repo/assignees/:assignee"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/comments"},
// {http
// MethodGet， "/repos/:owner/:repo/issues/comments"}， {http;MethodGet,“/回购:所有者/:回购/问题/评论/:id"},
	{http.MethodPost, "/repos/:owner/:repo/issues/:number/comments"},
// {http
// MethodPatch， "/repos/:owner/:repo/issues/comments/:id"}， {http;MethodDelete,“/回购:所有者/:回购/问题/评论/:id"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/events"},
// {http
// MethodGet， "/repos/:owner/:repo/issues/events"}， {http;MethodGet,“/回购:所有者/:回购/问题/事件/:id"},
	{http.MethodGet, "/repos/:owner/:repo/labels"},
	{http.MethodGet, "/repos/:owner/:repo/labels/:name"},
	{http.MethodPost, "/repos/:owner/:repo/labels"},
// {http
// MethodPatch,“/回购:所有者/:回购/标签/:name"},
	{http.MethodDelete, "/repos/:owner/:repo/labels/:name"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodPost, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodDelete, "/repos/:owner/:repo/issues/:number/labels/:name"},
	{http.MethodPut, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodDelete, "/repos/:owner/:repo/issues/:number/labels"},
	{http.MethodGet, "/repos/:owner/:repo/milestones/:number/labels"},
	{http.MethodGet, "/repos/:owner/:repo/milestones"},
	{http.MethodGet, "/repos/:owner/:repo/milestones/:number"},
	{http.MethodPost, "/repos/:owner/:repo/milestones"},
// {http
// MethodPatch,“/回购:所有者/:回购/里程碑:number"},
	{http.MethodDelete, "/repos/:owner/:repo/milestones/:number"},

// 杂项
	{http.MethodGet, "/emojis"},
	{http.MethodGet, "/gitignore/templates"},
	{http.MethodGet, "/gitignore/templates/:name"},
	{http.MethodPost, "/markdown"},
	{http.MethodPost, "/markdown/raw"},
	{http.MethodGet, "/meta"},
	{http.MethodGet, "/rate_limit"},

// 组织
	{http.MethodGet, "/users/:user/orgs"},
	{http.MethodGet, "/user/orgs"},
	{http.MethodGet, "/orgs/:org"},
// {http
// MethodPatch &quot / orgs: org"},,
	{http.MethodGet, "/orgs/:org/members"},
	{http.MethodGet, "/orgs/:org/members/:user"},
	{http.MethodDelete, "/orgs/:org/members/:user"},
	{http.MethodGet, "/orgs/:org/public_members"},
	{http.MethodGet, "/orgs/:org/public_members/:user"},
	{http.MethodPut, "/orgs/:org/public_members/:user"},
	{http.MethodDelete, "/orgs/:org/public_members/:user"},
	{http.MethodGet, "/orgs/:org/teams"},
	{http.MethodGet, "/teams/:id"},
	{http.MethodPost, "/orgs/:org/teams"},
// {http
// “MethodPatch; /团队/:id"},
	{http.MethodDelete, "/teams/:id"},
	{http.MethodGet, "/teams/:id/members"},
	{http.MethodGet, "/teams/:id/members/:user"},
	{http.MethodPut, "/teams/:id/members/:user"},
	{http.MethodDelete, "/teams/:id/members/:user"},
	{http.MethodGet, "/teams/:id/repos"},
	{http.MethodGet, "/teams/:id/repos/:owner/:repo"},
	{http.MethodPut, "/teams/:id/repos/:owner/:repo"},
	{http.MethodDelete, "/teams/:id/repos/:owner/:repo"},
	{http.MethodGet, "/user/teams"},

// 把请求
	{http.MethodGet, "/repos/:owner/:repo/pulls"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number"},
	{http.MethodPost, "/repos/:owner/:repo/pulls"},
// {http
// MethodPatch,“/回购:所有者/:回购/拉/:number"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/commits"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/files"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/merge"},
	{http.MethodPut, "/repos/:owner/:repo/pulls/:number/merge"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/comments"},
// {http
// MethodGet， "/repos/:owner/:repo/拉/评论"}， {http;MethodGet,“/回购:所有者/:回购/拉/评论/:number"},
	{http.MethodPut, "/repos/:owner/:repo/pulls/:number/comments"},
// {http
// MethodPatch， "/repos/:owner/:repo/pull /comments/:number"}， {http;MethodDelete,“/回购:所有者/:回购/拉/评论/:number"},

// 存储库
	{http.MethodGet, "/user/repos"},
	{http.MethodGet, "/users/:user/repos"},
	{http.MethodGet, "/orgs/:org/repos"},
	{http.MethodGet, "/repositories"},
	{http.MethodPost, "/user/repos"},
	{http.MethodPost, "/orgs/:org/repos"},
	{http.MethodGet, "/repos/:owner/:repo"},
// {http
// MethodPatch,“/回购:所有者/:repo"},
	{http.MethodGet, "/repos/:owner/:repo/contributors"},
	{http.MethodGet, "/repos/:owner/:repo/languages"},
	{http.MethodGet, "/repos/:owner/:repo/teams"},
	{http.MethodGet, "/repos/:owner/:repo/tags"},
	{http.MethodGet, "/repos/:owner/:repo/branches"},
	{http.MethodGet, "/repos/:owner/:repo/branches/:branch"},
	{http.MethodDelete, "/repos/:owner/:repo"},
	{http.MethodGet, "/repos/:owner/:repo/collaborators"},
	{http.MethodGet, "/repos/:owner/:repo/collaborators/:user"},
	{http.MethodPut, "/repos/:owner/:repo/collaborators/:user"},
	{http.MethodDelete, "/repos/:owner/:repo/collaborators/:user"},
	{http.MethodGet, "/repos/:owner/:repo/comments"},
	{http.MethodGet, "/repos/:owner/:repo/commits/:sha/comments"},
	{http.MethodPost, "/repos/:owner/:repo/commits/:sha/comments"},
	{http.MethodGet, "/repos/:owner/:repo/comments/:id"},
// {http
// MethodPatch,“/回购:所有者/:回购/评论/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/comments/:id"},
	{http.MethodGet, "/repos/:owner/:repo/commits"},
	{http.MethodGet, "/repos/:owner/:repo/commits/:sha"},
	{http.MethodGet, "/repos/:owner/:repo/readme"},
// {http
// MethodGet， "/repos/:owner/:repo/contents/*path"}， {http;MethodPut， "/repos/:owner/:repo/contents/*path"}， {http;MethodDelete， "/repos/:owner/:repo/contents/*path"}， {http;MethodGet,“/回购:所有者/:回购/:archive_format /: ref"},
	{http.MethodGet, "/repos/:owner/:repo/keys"},
	{http.MethodGet, "/repos/:owner/:repo/keys/:id"},
	{http.MethodPost, "/repos/:owner/:repo/keys"},
// {http
// MethodPatch,“/回购:所有者/:回购/键/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/keys/:id"},
	{http.MethodGet, "/repos/:owner/:repo/downloads"},
	{http.MethodGet, "/repos/:owner/:repo/downloads/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/downloads/:id"},
	{http.MethodGet, "/repos/:owner/:repo/forks"},
	{http.MethodPost, "/repos/:owner/:repo/forks"},
	{http.MethodGet, "/repos/:owner/:repo/hooks"},
	{http.MethodGet, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/hooks"},
// {http
// MethodPatch,“/回购:所有者/:回购/钩子:id"},
	{http.MethodPost, "/repos/:owner/:repo/hooks/:id/tests"},
	{http.MethodDelete, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/merges"},
	{http.MethodGet, "/repos/:owner/:repo/releases"},
	{http.MethodGet, "/repos/:owner/:repo/releases/:id"},
	{http.MethodPost, "/repos/:owner/:repo/releases"},
// {http
// MethodPatch,“/回购:所有者/:回购/版本/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/releases/:id"},
	{http.MethodGet, "/repos/:owner/:repo/releases/:id/assets"},
	{http.MethodGet, "/repos/:owner/:repo/stats/contributors"},
	{http.MethodGet, "/repos/:owner/:repo/stats/commit_activity"},
	{http.MethodGet, "/repos/:owner/:repo/stats/code_frequency"},
	{http.MethodGet, "/repos/:owner/:repo/stats/participation"},
	{http.MethodGet, "/repos/:owner/:repo/stats/punch_card"},
	{http.MethodGet, "/repos/:owner/:repo/statuses/:ref"},
	{http.MethodPost, "/repos/:owner/:repo/statuses/:ref"},

// 搜索
	{http.MethodGet, "/search/repositories"},
	{http.MethodGet, "/search/code"},
	{http.MethodGet, "/search/issues"},
	{http.MethodGet, "/search/users"},
	{http.MethodGet, "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{http.MethodGet, "/legacy/repos/search/:keyword"},
	{http.MethodGet, "/legacy/user/search/:keyword"},
	{http.MethodGet, "/legacy/user/email/:email"},

// 用户
	{http.MethodGet, "/users/:user"},
	{http.MethodGet, "/user"},
// {http
// MethodPatch,“/ user"},
	{http.MethodGet, "/users"},
	{http.MethodGet, "/user/emails"},
	{http.MethodPost, "/user/emails"},
	{http.MethodDelete, "/user/emails"},
	{http.MethodGet, "/users/:user/followers"},
	{http.MethodGet, "/user/followers"},
	{http.MethodGet, "/users/:user/following"},
	{http.MethodGet, "/user/following"},
	{http.MethodGet, "/user/following/:user"},
	{http.MethodGet, "/users/:user/following/:target_user"},
	{http.MethodPut, "/user/following/:user"},
	{http.MethodDelete, "/user/following/:user"},
	{http.MethodGet, "/users/:user/keys"},
	{http.MethodGet, "/user/keys"},
	{http.MethodGet, "/user/keys/:id"},
	{http.MethodPost, "/user/keys"},
// {http
// MethodPatch“/用户/键/:id"},
	{http.MethodDelete, "/user/keys/:id"},
}

func TestShouldBindUri(t *testing.T) {
	DefaultWriter = os.Stdout
	router := New()

	type Person struct {
		Name string `uri:"name" binding:"required"`
		ID   string `uri:"id" binding:"required"`
	}
	router.Handle(http.MethodGet, "/rest/:name/:id", func(c *Context) {
		var person Person
		assert.NoError(t, c.ShouldBindUri(&person))
		assert.True(t, person.Name != "")
		assert.True(t, person.ID != "")
		c.String(http.StatusOK, "ShouldBindUri test OK")
	})

	path, _ := exampleFromPath("/rest/:name/:id")
	w := PerformRequest(router, http.MethodGet, path)
	assert.Equal(t, "ShouldBindUri test OK", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBindUri(t *testing.T) {
	DefaultWriter = os.Stdout
	router := New()

	type Person struct {
		Name string `uri:"name" binding:"required"`
		ID   string `uri:"id" binding:"required"`
	}
	router.Handle(http.MethodGet, "/rest/:name/:id", func(c *Context) {
		var person Person
		assert.NoError(t, c.BindUri(&person))
		assert.True(t, person.Name != "")
		assert.True(t, person.ID != "")
		c.String(http.StatusOK, "BindUri test OK")
	})

	path, _ := exampleFromPath("/rest/:name/:id")
	w := PerformRequest(router, http.MethodGet, path)
	assert.Equal(t, "BindUri test OK", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBindUriError(t *testing.T) {
	DefaultWriter = os.Stdout
	router := New()

	type Member struct {
		Number string `uri:"num" binding:"required,uuid"`
	}
	router.Handle(http.MethodGet, "/new/rest/:num", func(c *Context) {
		var m Member
		assert.Error(t, c.BindUri(&m))
	})

	path1, _ := exampleFromPath("/new/rest/:num")
	w1 := PerformRequest(router, http.MethodGet, path1)
	assert.Equal(t, http.StatusBadRequest, w1.Code)
}

func TestRaceContextCopy(t *testing.T) {
	DefaultWriter = os.Stdout
	router := Default()
	router.GET("/test/copy/race", func(c *Context) {
		c.Set("1", 0)
		c.Set("2", 0)

// 将Context的副本发送给两个独立的例程
		go readWriteKeys(c.Copy())
		go readWriteKeys(c.Copy())
		c.String(http.StatusOK, "run OK, no panics")
	})
	w := PerformRequest(router, http.MethodGet, "/test/copy/race")
	assert.Equal(t, "run OK, no panics", w.Body.String())
}

func readWriteKeys(c *Context) {
	for {
		c.Set("1", rand.Int())
		c.Set("2", c.Value("1"))
	}
}

func githubConfigRouter(router *Engine) {
	for _, route := range githubAPI {
		router.Handle(route.method, route.path, func(c *Context) {
			output := make(map[string]string, len(c.Params)+1)
			output["status"] = "good"
			for _, param := range c.Params {
				output[param.Key] = param.Value
			}
			c.JSON(http.StatusOK, output)
		})
	}
}

func TestGithubAPI(t *testing.T) {
	DefaultWriter = os.Stdout
	router := New()
	githubConfigRouter(router)

	for _, route := range githubAPI {
		path, values := exampleFromPath(route.path)
		w := PerformRequest(router, route.method, path)

// 测试
		assert.Contains(t, w.Body.String(), "\"status\":\"good\"")
		for _, value := range values {
			str := fmt.Sprintf("\"%s\":\"%s\"", value.Key, value.Value)
			assert.Contains(t, w.Body.String(), str)
		}
	}
}

func exampleFromPath(path string) (string, Params) {
	output := new(strings.Builder)
	params := make(Params, 0, 6)
	start := -1
	for i, c := range path {
		if c == ':' {
			start = i + 1
		}
		if start >= 0 {
			if c == '/' {
				value := fmt.Sprint(rand.Intn(100000))
				params = append(params, Param{
					Key:   path[start:i],
					Value: value,
				})
				output.WriteString(value)
				output.WriteRune(c)
				start = -1
			}
		} else {
			output.WriteRune(c)
		}
	}
	if start >= 0 {
		value := fmt.Sprint(rand.Intn(100000))
		params = append(params, Param{
			Key:   path[start:],
			Value: value,
		})
		output.WriteString(value)
	}

	return output.String(), params
}

func BenchmarkGithub(b *testing.B) {
	router := New()
	githubConfigRouter(router)
	runRequest(b, router, http.MethodGet, "/legacy/issues/search/:owner/:repository/:state/:keyword")
}

func BenchmarkParallelGithub(b *testing.B) {
	DefaultWriter = os.Stdout
	router := New()
	githubConfigRouter(router)

	req, _ := http.NewRequest(http.MethodPost, "/repos/manucorporat/sse/git/blobs", nil)

	b.RunParallel(func(pb *testing.PB) {
// 每个程序都有自己的bytes.Buffer
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkParallelGithubDefault(b *testing.B) {
	DefaultWriter = os.Stdout
	router := New()
	githubConfigRouter(router)

	req, _ := http.NewRequest(http.MethodPost, "/repos/manucorporat/sse/git/blobs", nil)

	b.RunParallel(func(pb *testing.PB) {
// 每个程序都有自己的bytes.Buffer
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}
