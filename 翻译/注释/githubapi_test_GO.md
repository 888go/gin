
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// http://developer.github.com/v3/
<原文结束>

# <翻译开始>
// 参考网址：GitHub开发者文档 v3
# <翻译结束>


<原文开始>
	//{http.MethodPut, "/authorizations/clients/:client_id"},
	//{http.MethodPatch, "/authorizations/:id"},
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/notifications/threads/:id"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/notifications/threads/:id"} 的翻译为：
//（HTTP方法）PATCH，"/notifications/threads/:id" 
// 
// 这段注释描述的是在Go语言中一个HTTP路由的定义，具体含义如下：
// 
// - `http.MethodPatch`：表示HTTP请求方法为PATCH，主要用于更新资源的部分内容。
// - `"/notifications/threads/:id"`：这是一个URL路径模板，其中`:id`是一个占位符参数，表示通知线程的ID。所以这个路由用于处理针对特定通知线程进行PATCH操作的请求，例如更新某个通知线程的相关信息。
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/gists/public"},
	//{http.MethodGet, "/gists/starred"},
<原文结束>

# <翻译开始>
// {http.MethodGet, "/gists/public"}： 这是Go语言中的HTTP路由注释，表示处理GET请求的方法和对应的路径。翻译为中文为：`处理GET方法的请求，路径为"/gists/public"`。
// 
// {http.MethodGet, "/gists/starred"}： 同样也是HTTP路由注释，表示处理GET请求的方法和对应的路径。翻译为中文为：`处理GET方法的请求，路径为"/gists/starred"`。
// 
// 这两个注释分别描述了两个不同的接口路由，一个是获取公开的Gist资源，另一个是获取已加星标的Gist资源。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/gists/:id"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/gists/:id"} 的中文翻译是：
// 
// (HTTP方法)PATCH, 路由路径"/gists/（id变量）"，
// 
// 这段注释描述了HTTP服务的一个路由规则，其中：
// - `http.MethodPatch` 表示请求方法为 PATCH，
// - `"/gists/:id"` 表示URL路径，其中 `:id` 是一个动态参数，表示"gists"资源的ID。这个路由用于处理针对特定ID的gists资源的PATCH请求。
# <翻译结束>


<原文开始>
//{http.MethodGet, "/repos/:owner/:repo/git/refs/*ref"},
<原文结束>

# <翻译开始>
//{http.MethodGet, "/repos/:owner/:repo/git/refs/*ref"} 的翻译为：
// 
// GET 方法，请求路径为 "/repos/:owner/:repo/git/refs/*ref"
// 
// 其中：
// - `http.MethodGet` 表示HTTP的GET方法
// - `"/repos/:owner/:repo/git/refs/*ref"` 是API的路径模板，具体含义是：
//   - `:owner` 和 `:repo` 分别代表仓库所有者和仓库名，它们是动态参数
//   - `/git/refs/*ref` 表示获取Git引用（如分支、标签等），这里的 `*ref` 也是一个动态参数，表示可以匹配任何引用名称
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/git/refs/*ref"},
	//{http.MethodDelete, "/repos/:owner/:repo/git/refs/*ref"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/git/refs/*ref"}：该注释描述了一个HTTP路由规则，表示当接收到PATCH方法的请求时，路径应匹配"/repos/{owner}/{repo}/git/refs/*ref"。其中`:owner`和`:repo`是动态参数，代表仓库所有者和仓库名，`*ref`也是一个动态参数，代表Git引用（如分支、标签）的完整名称。
// 
// {http.MethodDelete, "/repos/:owner/:repo/git/refs/*ref"}：该注释描述了另一个HTTP路由规则，表示当接收到DELETE方法的请求时，路径同样应匹配"/repos/{owner}/{repo}/git/refs/*ref"。这个路由用于处理删除指定Git仓库中特定引用的操作。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/issues/:number"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/repos/:owner/:repo/issues/:number"} 的中文翻译是：
// 
// （HTTP方法）PATCH，"/repos/（仓库拥有者）/:owner/（仓库名）/:repo/issues/（问题编号）/:number"
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/issues/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/comments/:id"},
<原文结束>

# <翻译开始>
// GET 请求，访问路径为 "/repos/:owner/:repo/issues/comments"，其中 ":owner" 和 ":repo" 分别代表仓库所有者和仓库名，用于获取指定仓库的issue评论列表
// GET 请求，访问路径为 "/repos/:owner/:repo/issues/comments/:id"，其中 ":owner"、":repo" 和 ":id" 分别代表仓库所有者、仓库名和评论ID，用于获取指定仓库中特定ID的issue评论
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/issues/comments/:id"},
	//{http.MethodDelete, "/repos/:owner/:repo/issues/comments/:id"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/repos/:owner/:repo/issues/comments/:id"}： 
//PATCH方法，用于更新仓库中某条Issue的指定评论，路径为"/repos/拥有者/:repo/issues/comments/评论ID"
// 
//{http.MethodDelete, "/repos/:owner/:repo/issues/comments/:id"}：
//DELETE方法，用于删除仓库中某条Issue的指定评论，路径为"/repos/拥有者/:repo/issues/comments/评论ID"
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/issues/events"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/events/:id"},
<原文结束>

# <翻译开始>
// GET 请求，获取仓库中指定拥有者和仓库名的议题事件
// GET 请求，通过拥有者、仓库名和ID获取特定的议题事件
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/labels/:name"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/labels/:name"}
// 
// 这段Go语言注释是对一个HTTP路由的描述，翻译为：
// 
// 请求方法为PATCH，路由路径为"/repos/:owner/:repo/labels/:name"
// 
// 其中：
// - `http.MethodPatch` 表示HTTP请求方法为PATCH，通常用于更新资源的部分内容。
// - 路由路径 `/repos/:owner/:repo/labels/:name` 中的 `:owner`、`:repo` 和 `:name` 是动态参数，分别表示仓库所有者名称、仓库名称和标签名称。这条路由主要用于更新指定仓库中某个标签的信息。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/milestones/:number"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/milestones/:number"} 的中文翻译是：
// 
// (HTTP方法) PATCH, 路由路径: "/repos/仓库拥有者/:owner/仓库名称/:repo/milestones/里程碑编号/:number"
// 
// 这条注释描述了Go语言中一个HTTP路由规则，其中：
// - `http.MethodPatch` 表示这是一个PATCH类型的HTTP请求方法。
// - 路由路径 `/repos/:owner/:repo/milestones/:number` 指定了处理该请求的URL路径模板，其中`:owner`、`:repo`和`:number`是参数占位符，分别代表仓库拥有者的用户名、仓库名称和要操作的里程碑编号。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/orgs/:org"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/orgs/:org"} 的中文翻译是：
// 
// (HTTP方法) PATCH, 路由 "/orgs/:org"
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/teams/:id"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/teams/:id"} 的翻译为：
// 方法：PATCH，路径：/teams/:id
// 
// 这里是在Go语言中定义HTTP路由的一种方式，`http.MethodPatch` 表示HTTP请求方法为PATCH，`"/teams/:id"` 是URL路径，其中`:id` 是一个参数占位符，表示团队的ID，可以根据不同的ID值匹配到相应的团队资源进行PATCH操作。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/pulls/:number"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/repos/:owner/:repo/pulls/:number"}，
// 
// 这个Go语言的注释表示一个HTTP路由规则，翻译为：
// 
// 使用HTTP方法PATCH，路径为"/repos/:owner/:repo/pulls/:number"
// 
// 其中，`:owner`、`:repo`和`:number`是参数占位符，分别代表仓库所有者名称、仓库名称和拉取请求（Pull Request）的编号。在实际应用中，这三个部分会被具体的值所替换，以实现对特定资源的PATCH操作。
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/pulls/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/pulls/comments/:number"},
<原文结束>

# <翻译开始>
// GET 请求，访问路径为 "/repos/:owner/:repo/pulls/comments"，其中 ":owner" 和 ":repo" 代表仓库所有者和仓库名，用于获取指定仓库拉取请求的全部评论。
// GET 请求，访问路径为 "/repos/:owner/:repo/pulls/comments/:number"，其中 ":owner"、":repo" 和 ":number" 分别代表仓库所有者、仓库名和评论编号，用于获取指定仓库拉取请求中特定编号的评论。
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/pulls/comments/:number"},
	//{http.MethodDelete, "/repos/:owner/:repo/pulls/comments/:number"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/pulls/comments/:number"}：对repos仓库下指定owner和repo的pull requests评论（编号为:number）进行更新或修改操作。
// {http.MethodDelete, "/repos/:owner/:repo/pulls/comments/:number"}：删除repos仓库下指定owner和repo的pull requests评论（编号为:number）。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo"} 的中文翻译是：
// 
// `{http.MethodPatch, "/repos/:owner/:repo"}` 表示：
// 
// 这是一个HTTP路由注释，描述了API接口的请求方法和路径。
// 
// - `http.MethodPatch`：表示HTTP请求方法为PATCH，通常用于更新资源的部分内容。
//   
// - `"/repos/:owner/:repo"`：表示URL路径，其中`:owner`和`:repo`是占位符参数，分别代表仓库所有者和仓库名称。这条路径通常用于访问或修改GitHub等平台上的特定用户的所有仓库中的某个仓库。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/comments/:id"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/comments/:id"} 翻译为：
// 
// 使用HTTP方法PATCH，访问"/repos/:owner/:repo/comments/:id"路径
// 
// 其中：
// - `http.MethodPatch`表示使用HTTP的PATCH方法，用于更新资源的部分内容。
// - `"/repos/:owner/:repo/comments/:id"`是一个URL路径模板，其中`:owner`、`:repo`和`:id`是参数占位符，分别代表仓库所有者名称、仓库名称和评论ID。在实际请求时，这三个占位符会被具体的值替换。
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodPut, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodDelete, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodGet, "/repos/:owner/:repo/:archive_format/:ref"},
<原文结束>

# <翻译开始>
// {http.MethodGet, "/repos/:owner/:repo/contents/*path"}：GET 方法，访问路径为 /repos/仓库拥有者/仓库名/contents/任意路径
// {http.MethodPut, "/repos/:owner/:repo/contents/*path"}：PUT 方法，访问路径为 /repos/仓库拥有者/仓库名/contents/任意路径
// {http.MethodDelete, "/repos/:owner/:repo/contents/*path"}：DELETE 方法，访问路径为 /repos/仓库拥有者/仓库名/contents/任意路径
// {http.MethodGet, "/repos/:owner/:repo/:archive_format/:ref"}：GET 方法，访问路径为 /repos/仓库拥有者/仓库名/指定归档格式/分支或提交引用
// 
// 这些注释描述了在Go语言中HTTP请求处理的路由规则，其中`:owner`、`:repo`、`:path` 和 `:archive_format` 是动态参数，分别表示仓库拥有者名称、仓库名称、任意路径和归档格式，`:ref` 表示特定的分支或提交引用。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/keys/:id"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/keys/:id"} 的中文注释可以这样表述：
// 
// 使用HTTP方法Patch，请求路径为"/repos/:owner/:repo/keys/:id"，
// 
// 这里的":owner"、":repo"和":id"是动态参数，分别代表仓库所有者用户名、仓库名和密钥ID。这个API路由主要用于更新指定仓库的特定密钥信息。
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/hooks/:id"},
<原文结束>

# <翻译开始>
// {http.MethodPatch, "/repos/:owner/:repo/hooks/:id"} 
// 
// 这段Go注释的翻译为：
// 
// 使用HTTP方法PATCH，访问路径为"/repos/:owner/:repo/hooks/:id"
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/repos/:owner/:repo/releases/:id"},
<原文结束>

# <翻译开始>
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
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/user"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/user"} 的翻译为：
// 
// （HTTP方法）PATCH，"/user"（URL路径）
# <翻译结束>


<原文开始>
//{http.MethodPatch, "/user/keys/:id"},
<原文结束>

# <翻译开始>
//{http.MethodPatch, "/user/keys/:id"} 的中文注释翻译为：
// 
// 使用HTTP方法Patch，请求路径为"/user/keys/:id"
# <翻译结束>


<原文开始>
// Sending a copy of the Context to two separate routines
<原文结束>

# <翻译开始>
// 向两个独立的goroutine发送Context的副本
# <翻译结束>


<原文开始>
// Each goroutine has its own bytes.Buffer.
<原文结束>

# <翻译开始>
// 每个goroutine都有自己独立的bytes.Buffer。
# <翻译结束>

