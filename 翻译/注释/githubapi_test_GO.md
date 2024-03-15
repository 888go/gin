
<原文开始>
// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到
# <翻译结束>


<原文开始>
// http://developer.github.com/v3/
<原文结束>

# <翻译开始>
// http://developer.github.com/v3/
# <翻译结束>


<原文开始>
	// OAuth Authorizations
<原文结束>

# <翻译开始>
// OAuth授权
# <翻译结束>


<原文开始>
	//{http.MethodPut, "/authorizations/clients/:client_id"},
	//{http.MethodPatch, "/authorizations/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPut， "/authorizations/clients/:client_id"}， {http;“MethodPatch; /授权/:id"},
# <翻译结束>


<原文开始>
	// Activity
<原文结束>

# <翻译开始>
// 活动
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/notifications/threads/:id"},
<原文结束>

# <翻译开始>
// {http
// “MethodPatch; / /线程/通知:id"},
# <翻译结束>


<原文开始>
	// Gists
<原文结束>

# <翻译开始>
// 依据
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/gists/public"},
	//{http.MethodGet, "/gists/starred"},
<原文结束>

# <翻译开始>
// {http
// MethodGet， "/gist /public"}， {http;MethodGet,“/丰子恺/ starred"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/gists/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“丰子恺/:id"},
# <翻译结束>


<原文开始>
	// Git Data
<原文结束>

# <翻译开始>
// Git数据
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/git/refs/*ref"},
<原文结束>

# <翻译开始>
// {http
// MethodGet " / /:业主休息,回购:git / refs ref"} / *,
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/git/refs/*ref"},
	//{http.MethodDelete, "/repos/:owner/:repo/git/refs/*ref"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch， "/repos/:owner/:repo/git/refs/* refquot;}， {http;MethodDelete,“/回购:所有者/:回购/ git / refs / * ref"},
# <翻译结束>


<原文开始>
	// Issues
<原文结束>

# <翻译开始>
// 问题
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/issues/:number"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/问题/:number"},
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/issues/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/comments/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodGet， "/repos/:owner/:repo/issues/comments"}， {http;MethodGet,“/回购:所有者/:回购/问题/评论/:id"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/issues/comments/:id"},
	//{http.MethodDelete, "/repos/:owner/:repo/issues/comments/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch， "/repos/:owner/:repo/issues/comments/:id"}， {http;MethodDelete,“/回购:所有者/:回购/问题/评论/:id"},
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/issues/events"},
	//{http.MethodGet, "/repos/:owner/:repo/issues/events/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodGet， "/repos/:owner/:repo/issues/events"}， {http;MethodGet,“/回购:所有者/:回购/问题/事件/:id"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/labels/:name"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/标签/:name"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/milestones/:number"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/里程碑:number"},
# <翻译结束>


<原文开始>
	// Miscellaneous
<原文结束>

# <翻译开始>
// 杂项
# <翻译结束>


<原文开始>
	// Organizations
<原文结束>

# <翻译开始>
// 组织
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/orgs/:org"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch &quot / orgs: org"},,
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/teams/:id"},
<原文结束>

# <翻译开始>
// {http
// “MethodPatch; /团队/:id"},
# <翻译结束>


<原文开始>
	// Pull Requests
<原文结束>

# <翻译开始>
// 把请求
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/pulls/:number"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/拉/:number"},
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/pulls/comments"},
	//{http.MethodGet, "/repos/:owner/:repo/pulls/comments/:number"},
<原文结束>

# <翻译开始>
// {http
// MethodGet， "/repos/:owner/:repo/拉/评论"}， {http;MethodGet,“/回购:所有者/:回购/拉/评论/:number"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/pulls/comments/:number"},
	//{http.MethodDelete, "/repos/:owner/:repo/pulls/comments/:number"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch， "/repos/:owner/:repo/pull /comments/:number"}， {http;MethodDelete,“/回购:所有者/:回购/拉/评论/:number"},
# <翻译结束>


<原文开始>
	// Repositories
<原文结束>

# <翻译开始>
// 存储库
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:repo"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/comments/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/评论/:id"},
# <翻译结束>


<原文开始>
	//{http.MethodGet, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodPut, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodDelete, "/repos/:owner/:repo/contents/*path"},
	//{http.MethodGet, "/repos/:owner/:repo/:archive_format/:ref"},
<原文结束>

# <翻译开始>
// {http
// MethodGet， "/repos/:owner/:repo/contents/*path"}， {http;MethodPut， "/repos/:owner/:repo/contents/*path"}， {http;MethodDelete， "/repos/:owner/:repo/contents/*path"}， {http;MethodGet,“/回购:所有者/:回购/:archive_format /: ref"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/keys/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/键/:id"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/hooks/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/钩子:id"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/repos/:owner/:repo/releases/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/回购:所有者/:回购/版本/:id"},
# <翻译结束>


<原文开始>
	// Search
<原文结束>

# <翻译开始>
// 搜索
# <翻译结束>


<原文开始>
	// Users
<原文结束>

# <翻译开始>
// 用户
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/user"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch,“/ user"},
# <翻译结束>


<原文开始>
	//{http.MethodPatch, "/user/keys/:id"},
<原文结束>

# <翻译开始>
// {http
// MethodPatch“/用户/键/:id"},
# <翻译结束>


<原文开始>
		// Sending a copy of the Context to two separate routines
<原文结束>

# <翻译开始>
// 将Context的副本发送给两个独立的例程
# <翻译结束>


<原文开始>
		// TEST
<原文结束>

# <翻译开始>
// 测试
# <翻译结束>


<原文开始>
		// Each goroutine has its own bytes.Buffer.
<原文结束>

# <翻译开始>
// 每个程序都有自己的bytes.Buffer
# <翻译结束>

