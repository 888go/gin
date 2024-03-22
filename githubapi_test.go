// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin类

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

// 参考网址：GitHub开发者文档 v3
var githubAPI = []route{
	// OAuth Authorizations
	{http.MethodGet, "/authorizations"},
	{http.MethodGet, "/authorizations/:id"},
	{http.MethodPost, "/authorizations"},
// {http.MethodPut, "/authorizations/clients/:client_id"}：  
// 对于HTTP方法PUT，路由为"/authorizations/clients/:client_id" 
// 
// {http.MethodPatch, "/authorizations/:id"}：  
// 对于HTTP方法PATCH，路由为"/authorizations/:id" 
// 
// 这里，`http.MethodPut`和`http.MethodPatch`分别表示HTTP请求的PUT和PATCH方法，冒号（:`client_id`和`:id`）表示路由参数，这些参数在实际使用时会被动态值替换。具体来说：
// 
// - `/authorizations/clients/:client_id`：用于处理与特定客户端ID相关的授权更新或创建操作。
// - `/authorizations/:id`：用于处理通过其ID进行部分更新的操作，这里的授权是指某个具体的授权资源。
	{http.MethodDelete, "/authorizations/:id"},
	{http.MethodGet, "/applications/:client_id/tokens/:access_token"},
	{http.MethodDelete, "/applications/:client_id/tokens"},
	{http.MethodDelete, "/applications/:client_id/tokens/:access_token"},

	// Activity
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
	//{http.MethodPatch, "/notifications/threads/:id"} 的翻译为：
//（HTTP方法）PATCH，"/notifications/threads/:id" 
// 
// 这段注释描述的是在Go语言中一个HTTP路由的定义，具体含义如下：
// 
// - `http.MethodPatch`：表示HTTP请求方法为PATCH，主要用于更新资源的部分内容。
// - `"/notifications/threads/:id"`：这是一个URL路径模板，其中`:id`是一个占位符参数，表示通知线程的ID。所以这个路由用于处理针对特定通知线程进行PATCH操作的请求，例如更新某个通知线程的相关信息。
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

	// Gists
	{http.MethodGet, "/users/:user/gists"},
	{http.MethodGet, "/gists"},
// {http.MethodGet, "/gists/public"}： 这是Go语言中的HTTP路由注释，表示处理GET请求的方法和对应的路径。翻译为中文为：`处理GET方法的请求，路径为"/gists/public"`。
// 
// {http.MethodGet, "/gists/starred"}： 同样也是HTTP路由注释，表示处理GET请求的方法和对应的路径。翻译为中文为：`处理GET方法的请求，路径为"/gists/starred"`。
// 
// 这两个注释分别描述了两个不同的接口路由，一个是获取公开的Gist资源，另一个是获取已加星标的Gist资源。
	{http.MethodGet, "/gists/:id"},
	{http.MethodPost, "/gists"},
	// {http.MethodPatch, "/gists/:id"} 的中文翻译是：
// 
// (HTTP方法)PATCH, 路由路径"/gists/（id变量）"，
// 
// 这段注释描述了HTTP服务的一个路由规则，其中：
// - `http.MethodPatch` 表示请求方法为 PATCH，
// - `"/gists/:id"` 表示URL路径，其中 `:id` 是一个动态参数，表示"gists"资源的ID。这个路由用于处理针对特定ID的gists资源的PATCH请求。
	{http.MethodPut, "/gists/:id/star"},
	{http.MethodDelete, "/gists/:id/star"},
	{http.MethodGet, "/gists/:id/star"},
	{http.MethodPost, "/gists/:id/forks"},
	{http.MethodDelete, "/gists/:id"},

	// Git Data
	{http.MethodGet, "/repos/:owner/:repo/git/blobs/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/blobs"},
	{http.MethodGet, "/repos/:owner/:repo/git/commits/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/commits"},
	//{http.MethodGet, "/repos/:owner/:repo/git/refs/*ref"} 的翻译为：
// 
// GET 方法，请求路径为 "/repos/:owner/:repo/git/refs/*ref"
// 
// 其中：
// - `http.MethodGet` 表示HTTP的GET方法
// - `"/repos/:owner/:repo/git/refs/*ref"` 是API的路径模板，具体含义是：
//   - `:owner` 和 `:repo` 分别代表仓库所有者和仓库名，它们是动态参数
//   - `/git/refs/*ref` 表示获取Git引用（如分支、标签等），这里的 `*ref` 也是一个动态参数，表示可以匹配任何引用名称
	{http.MethodGet, "/repos/:owner/:repo/git/refs"},
	{http.MethodPost, "/repos/:owner/:repo/git/refs"},
// {http.MethodPatch, "/repos/:owner/:repo/git/refs/*ref"}：该注释描述了一个HTTP路由规则，表示当接收到PATCH方法的请求时，路径应匹配"/repos/{owner}/{repo}/git/refs/*ref"。其中`:owner`和`:repo`是动态参数，代表仓库所有者和仓库名，`*ref`也是一个动态参数，代表Git引用（如分支、标签）的完整名称。
// 
// {http.MethodDelete, "/repos/:owner/:repo/git/refs/*ref"}：该注释描述了另一个HTTP路由规则，表示当接收到DELETE方法的请求时，路径同样应匹配"/repos/{owner}/{repo}/git/refs/*ref"。这个路由用于处理删除指定Git仓库中特定引用的操作。
	{http.MethodGet, "/repos/:owner/:repo/git/tags/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/tags"},
	{http.MethodGet, "/repos/:owner/:repo/git/trees/:sha"},
	{http.MethodPost, "/repos/:owner/:repo/git/trees"},

	// Issues
	{http.MethodGet, "/issues"},
	{http.MethodGet, "/user/issues"},
	{http.MethodGet, "/orgs/:org/issues"},
	{http.MethodGet, "/repos/:owner/:repo/issues"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number"},
	{http.MethodPost, "/repos/:owner/:repo/issues"},
	//{http.MethodPatch, "/repos/:owner/:repo/issues/:number"} 的中文翻译是：
// 
// （HTTP方法）PATCH，"/repos/（仓库拥有者）/:owner/（仓库名）/:repo/issues/（问题编号）/:number"
	{http.MethodGet, "/repos/:owner/:repo/assignees"},
	{http.MethodGet, "/repos/:owner/:repo/assignees/:assignee"},
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/comments"},
// GET 请求，访问路径为 "/repos/:owner/:repo/issues/comments"，其中 ":owner" 和 ":repo" 分别代表仓库所有者和仓库名，用于获取指定仓库的issue评论列表
// GET 请求，访问路径为 "/repos/:owner/:repo/issues/comments/:id"，其中 ":owner"、":repo" 和 ":id" 分别代表仓库所有者、仓库名和评论ID，用于获取指定仓库中特定ID的issue评论
	{http.MethodPost, "/repos/:owner/:repo/issues/:number/comments"},
//{http.MethodPatch, "/repos/:owner/:repo/issues/comments/:id"}： 
//PATCH方法，用于更新仓库中某条Issue的指定评论，路径为"/repos/拥有者/:repo/issues/comments/评论ID"
// 
//{http.MethodDelete, "/repos/:owner/:repo/issues/comments/:id"}：
//DELETE方法，用于删除仓库中某条Issue的指定评论，路径为"/repos/拥有者/:repo/issues/comments/评论ID"
	{http.MethodGet, "/repos/:owner/:repo/issues/:number/events"},
// GET 请求，获取仓库中指定拥有者和仓库名的议题事件
// GET 请求，通过拥有者、仓库名和ID获取特定的议题事件
	{http.MethodGet, "/repos/:owner/:repo/labels"},
	{http.MethodGet, "/repos/:owner/:repo/labels/:name"},
	{http.MethodPost, "/repos/:owner/:repo/labels"},
	// {http.MethodPatch, "/repos/:owner/:repo/labels/:name"}
// 
// 这段Go语言注释是对一个HTTP路由的描述，翻译为：
// 
// 请求方法为PATCH，路由路径为"/repos/:owner/:repo/labels/:name"
// 
// 其中：
// - `http.MethodPatch` 表示HTTP请求方法为PATCH，通常用于更新资源的部分内容。
// - 路由路径 `/repos/:owner/:repo/labels/:name` 中的 `:owner`、`:repo` 和 `:name` 是动态参数，分别表示仓库所有者名称、仓库名称和标签名称。这条路由主要用于更新指定仓库中某个标签的信息。
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
	// {http.MethodPatch, "/repos/:owner/:repo/milestones/:number"} 的中文翻译是：
// 
// (HTTP方法) PATCH, 路由路径: "/repos/仓库拥有者/:owner/仓库名称/:repo/milestones/里程碑编号/:number"
// 
// 这条注释描述了Go语言中一个HTTP路由规则，其中：
// - `http.MethodPatch` 表示这是一个PATCH类型的HTTP请求方法。
// - 路由路径 `/repos/:owner/:repo/milestones/:number` 指定了处理该请求的URL路径模板，其中`:owner`、`:repo`和`:number`是参数占位符，分别代表仓库拥有者的用户名、仓库名称和要操作的里程碑编号。
	{http.MethodDelete, "/repos/:owner/:repo/milestones/:number"},

	// Miscellaneous
	{http.MethodGet, "/emojis"},
	{http.MethodGet, "/gitignore/templates"},
	{http.MethodGet, "/gitignore/templates/:name"},
	{http.MethodPost, "/markdown"},
	{http.MethodPost, "/markdown/raw"},
	{http.MethodGet, "/meta"},
	{http.MethodGet, "/rate_limit"},

	// Organizations
	{http.MethodGet, "/users/:user/orgs"},
	{http.MethodGet, "/user/orgs"},
	{http.MethodGet, "/orgs/:org"},
	//{http.MethodPatch, "/orgs/:org"} 的中文翻译是：
// 
// (HTTP方法) PATCH, 路由 "/orgs/:org"
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
	//{http.MethodPatch, "/teams/:id"} 的翻译为：
// 方法：PATCH，路径：/teams/:id
// 
// 这里是在Go语言中定义HTTP路由的一种方式，`http.MethodPatch` 表示HTTP请求方法为PATCH，`"/teams/:id"` 是URL路径，其中`:id` 是一个参数占位符，表示团队的ID，可以根据不同的ID值匹配到相应的团队资源进行PATCH操作。
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

	// Pull Requests
	{http.MethodGet, "/repos/:owner/:repo/pulls"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number"},
	{http.MethodPost, "/repos/:owner/:repo/pulls"},
	//{http.MethodPatch, "/repos/:owner/:repo/pulls/:number"}，
// 
// 这个Go语言的注释表示一个HTTP路由规则，翻译为：
// 
// 使用HTTP方法PATCH，路径为"/repos/:owner/:repo/pulls/:number"
// 
// 其中，`:owner`、`:repo`和`:number`是参数占位符，分别代表仓库所有者名称、仓库名称和拉取请求（Pull Request）的编号。在实际应用中，这三个部分会被具体的值所替换，以实现对特定资源的PATCH操作。
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/commits"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/files"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/merge"},
	{http.MethodPut, "/repos/:owner/:repo/pulls/:number/merge"},
	{http.MethodGet, "/repos/:owner/:repo/pulls/:number/comments"},
// GET 请求，访问路径为 "/repos/:owner/:repo/pulls/comments"，其中 ":owner" 和 ":repo" 代表仓库所有者和仓库名，用于获取指定仓库拉取请求的全部评论。
// GET 请求，访问路径为 "/repos/:owner/:repo/pulls/comments/:number"，其中 ":owner"、":repo" 和 ":number" 分别代表仓库所有者、仓库名和评论编号，用于获取指定仓库拉取请求中特定编号的评论。
	{http.MethodPut, "/repos/:owner/:repo/pulls/:number/comments"},
// {http.MethodPatch, "/repos/:owner/:repo/pulls/comments/:number"}：对repos仓库下指定owner和repo的pull requests评论（编号为:number）进行更新或修改操作。
// {http.MethodDelete, "/repos/:owner/:repo/pulls/comments/:number"}：删除repos仓库下指定owner和repo的pull requests评论（编号为:number）。

	// Repositories
	{http.MethodGet, "/user/repos"},
	{http.MethodGet, "/users/:user/repos"},
	{http.MethodGet, "/orgs/:org/repos"},
	{http.MethodGet, "/repositories"},
	{http.MethodPost, "/user/repos"},
	{http.MethodPost, "/orgs/:org/repos"},
	{http.MethodGet, "/repos/:owner/:repo"},
	// {http.MethodPatch, "/repos/:owner/:repo"} 的中文翻译是：
// 
// `{http.MethodPatch, "/repos/:owner/:repo"}` 表示：
// 
// 这是一个HTTP路由注释，描述了API接口的请求方法和路径。
// 
// - `http.MethodPatch`：表示HTTP请求方法为PATCH，通常用于更新资源的部分内容。
//   
// - `"/repos/:owner/:repo"`：表示URL路径，其中`:owner`和`:repo`是占位符参数，分别代表仓库所有者和仓库名称。这条路径通常用于访问或修改GitHub等平台上的特定用户的所有仓库中的某个仓库。
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
	// {http.MethodPatch, "/repos/:owner/:repo/comments/:id"} 翻译为：
// 
// 使用HTTP方法PATCH，访问"/repos/:owner/:repo/comments/:id"路径
// 
// 其中：
// - `http.MethodPatch`表示使用HTTP的PATCH方法，用于更新资源的部分内容。
// - `"/repos/:owner/:repo/comments/:id"`是一个URL路径模板，其中`:owner`、`:repo`和`:id`是参数占位符，分别代表仓库所有者名称、仓库名称和评论ID。在实际请求时，这三个占位符会被具体的值替换。
	{http.MethodDelete, "/repos/:owner/:repo/comments/:id"},
	{http.MethodGet, "/repos/:owner/:repo/commits"},
	{http.MethodGet, "/repos/:owner/:repo/commits/:sha"},
	{http.MethodGet, "/repos/:owner/:repo/readme"},
// {http.MethodGet, "/repos/:owner/:repo/contents/*path"}：GET 方法，访问路径为 /repos/仓库拥有者/仓库名/contents/任意路径
// {http.MethodPut, "/repos/:owner/:repo/contents/*path"}：PUT 方法，访问路径为 /repos/仓库拥有者/仓库名/contents/任意路径
// {http.MethodDelete, "/repos/:owner/:repo/contents/*path"}：DELETE 方法，访问路径为 /repos/仓库拥有者/仓库名/contents/任意路径
// {http.MethodGet, "/repos/:owner/:repo/:archive_format/:ref"}：GET 方法，访问路径为 /repos/仓库拥有者/仓库名/指定归档格式/分支或提交引用
// 
// 这些注释描述了在Go语言中HTTP请求处理的路由规则，其中`:owner`、`:repo`、`:path` 和 `:archive_format` 是动态参数，分别表示仓库拥有者名称、仓库名称、任意路径和归档格式，`:ref` 表示特定的分支或提交引用。
	{http.MethodGet, "/repos/:owner/:repo/keys"},
	{http.MethodGet, "/repos/:owner/:repo/keys/:id"},
	{http.MethodPost, "/repos/:owner/:repo/keys"},
	// {http.MethodPatch, "/repos/:owner/:repo/keys/:id"} 的中文注释可以这样表述：
// 
// 使用HTTP方法Patch，请求路径为"/repos/:owner/:repo/keys/:id"，
// 
// 这里的":owner"、":repo"和":id"是动态参数，分别代表仓库所有者用户名、仓库名和密钥ID。这个API路由主要用于更新指定仓库的特定密钥信息。
	{http.MethodDelete, "/repos/:owner/:repo/keys/:id"},
	{http.MethodGet, "/repos/:owner/:repo/downloads"},
	{http.MethodGet, "/repos/:owner/:repo/downloads/:id"},
	{http.MethodDelete, "/repos/:owner/:repo/downloads/:id"},
	{http.MethodGet, "/repos/:owner/:repo/forks"},
	{http.MethodPost, "/repos/:owner/:repo/forks"},
	{http.MethodGet, "/repos/:owner/:repo/hooks"},
	{http.MethodGet, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/hooks"},
	// {http.MethodPatch, "/repos/:owner/:repo/hooks/:id"} 
// 
// 这段Go注释的翻译为：
// 
// 使用HTTP方法PATCH，访问路径为"/repos/:owner/:repo/hooks/:id"
	{http.MethodPost, "/repos/:owner/:repo/hooks/:id/tests"},
	{http.MethodDelete, "/repos/:owner/:repo/hooks/:id"},
	{http.MethodPost, "/repos/:owner/:repo/merges"},
	{http.MethodGet, "/repos/:owner/:repo/releases"},
	{http.MethodGet, "/repos/:owner/:repo/releases/:id"},
	{http.MethodPost, "/repos/:owner/:repo/releases"},
	// {http.MethodPatch, "/repos/:owner/:repo/releases/:id"}
// 
// 这段Go语言的注释表示的是一个HTTP路由映射，翻译为：
// 
// 使用HTTP方法PATCH，路径为"/repos/:owner/:repo/releases/:id"
// 
// 其中：
// 
// - `http.MethodPatch` 表示HTTP请求方法为PATCH，主要用于更新资源的部分内容。
// - 路径`/repos/:owner/:repo/releases/:id`中，`:owner`、`:repo`和`:id`是动态参数，分别代表仓库所有者用户名、仓库名以及要操作的发布版本ID。这个路由通常用于更新GitHub或其他类似平台上的某个仓库的指定版本信息。
	{http.MethodDelete, "/repos/:owner/:repo/releases/:id"},
	{http.MethodGet, "/repos/:owner/:repo/releases/:id/assets"},
	{http.MethodGet, "/repos/:owner/:repo/stats/contributors"},
	{http.MethodGet, "/repos/:owner/:repo/stats/commit_activity"},
	{http.MethodGet, "/repos/:owner/:repo/stats/code_frequency"},
	{http.MethodGet, "/repos/:owner/:repo/stats/participation"},
	{http.MethodGet, "/repos/:owner/:repo/stats/punch_card"},
	{http.MethodGet, "/repos/:owner/:repo/statuses/:ref"},
	{http.MethodPost, "/repos/:owner/:repo/statuses/:ref"},

	// Search
	{http.MethodGet, "/search/repositories"},
	{http.MethodGet, "/search/code"},
	{http.MethodGet, "/search/issues"},
	{http.MethodGet, "/search/users"},
	{http.MethodGet, "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{http.MethodGet, "/legacy/repos/search/:keyword"},
	{http.MethodGet, "/legacy/user/search/:keyword"},
	{http.MethodGet, "/legacy/user/email/:email"},

	// Users
	{http.MethodGet, "/users/:user"},
	{http.MethodGet, "/user"},
	//{http.MethodPatch, "/user"} 的翻译为：
// 
// （HTTP方法）PATCH，"/user"（URL路径）
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
	//{http.MethodPatch, "/user/keys/:id"} 的中文注释翻译为：
// 
// 使用HTTP方法Patch，请求路径为"/user/keys/:id"
	{http.MethodDelete, "/user/keys/:id"},
}

func TestShouldBindUri(t *testing.T) {
	DefaultWriter = os.Stdout
	router := X创建()

	type Person struct {
		Name string `uri:"name" binding:"required"`
		ID   string `uri:"id" binding:"required"`
	}
	router.X绑定(http.MethodGet, "/rest/:name/:id", func(c *Context) {
		var person Person
		assert.NoError(t, c.X取Uri参数到指针(&person))
		assert.True(t, person.Name != "")
		assert.True(t, person.ID != "")
		c.X输出文本(http.StatusOK, "ShouldBindUri test OK")
	})

	path, _ := exampleFromPath("/rest/:name/:id")
	w := PerformRequest(router, http.MethodGet, path)
	assert.Equal(t, "ShouldBindUri test OK", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBindUri(t *testing.T) {
	DefaultWriter = os.Stdout
	router := X创建()

	type Person struct {
		Name string `uri:"name" binding:"required"`
		ID   string `uri:"id" binding:"required"`
	}
	router.X绑定(http.MethodGet, "/rest/:name/:id", func(c *Context) {
		var person Person
		assert.NoError(t, c.X取Uri参数到指针PANI(&person))
		assert.True(t, person.Name != "")
		assert.True(t, person.ID != "")
		c.X输出文本(http.StatusOK, "BindUri test OK")
	})

	path, _ := exampleFromPath("/rest/:name/:id")
	w := PerformRequest(router, http.MethodGet, path)
	assert.Equal(t, "BindUri test OK", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBindUriError(t *testing.T) {
	DefaultWriter = os.Stdout
	router := X创建()

	type Member struct {
		Number string `uri:"num" binding:"required,uuid"`
	}
	router.X绑定(http.MethodGet, "/new/rest/:num", func(c *Context) {
		var m Member
		assert.Error(t, c.X取Uri参数到指针PANI(&m))
	})

	path1, _ := exampleFromPath("/new/rest/:num")
	w1 := PerformRequest(router, http.MethodGet, path1)
	assert.Equal(t, http.StatusBadRequest, w1.Code)
}

func TestRaceContextCopy(t *testing.T) {
	DefaultWriter = os.Stdout
	router := X创建默认对象()
	router.X绑定GET("/test/copy/race", func(c *Context) {
		c.X设置值("1", 0)
		c.X设置值("2", 0)

		// 向两个独立的goroutine发送Context的副本
		go readWriteKeys(c.X取副本())
		go readWriteKeys(c.X取副本())
		c.X输出文本(http.StatusOK, "run OK, no panics")
	})
	w := PerformRequest(router, http.MethodGet, "/test/copy/race")
	assert.Equal(t, "run OK, no panics", w.Body.String())
}

func readWriteKeys(c *Context) {
	for {
		c.X设置值("1", rand.Int())
		c.X设置值("2", c.Value("1"))
	}
}

func githubConfigRouter(router *Engine) {
	for _, route := range githubAPI {
		router.X绑定(route.method, route.path, func(c *Context) {
			output := make(map[string]string, len(c.X参数)+1)
			output["status"] = "good"
			for _, param := range c.X参数 {
				output[param.Key] = param.Value
			}
			c.X输出JSON(http.StatusOK, output)
		})
	}
}

func TestGithubAPI(t *testing.T) {
	DefaultWriter = os.Stdout
	router := X创建()
	githubConfigRouter(router)

	for _, route := range githubAPI {
		path, values := exampleFromPath(route.path)
		w := PerformRequest(router, route.method, path)

		// TEST
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
	router := X创建()
	githubConfigRouter(router)
	runRequest(b, router, http.MethodGet, "/legacy/issues/search/:owner/:repository/:state/:keyword")
}

func BenchmarkParallelGithub(b *testing.B) {
	DefaultWriter = os.Stdout
	router := X创建()
	githubConfigRouter(router)

	req, _ := http.NewRequest(http.MethodPost, "/repos/manucorporat/sse/git/blobs", nil)

	b.RunParallel(func(pb *testing.PB) {
		// 每个goroutine都有自己独立的bytes.Buffer。
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}

func BenchmarkParallelGithubDefault(b *testing.B) {
	DefaultWriter = os.Stdout
	router := X创建()
	githubConfigRouter(router)

	req, _ := http.NewRequest(http.MethodPost, "/repos/manucorporat/sse/git/blobs", nil)

	b.RunParallel(func(pb *testing.PB) {
		// 每个goroutine都有自己独立的bytes.Buffer。
		for pb.Next() {
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}
	})
}
