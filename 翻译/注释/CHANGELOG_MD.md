
<原文开始>
BUG FIXES 

* fix Request.Context() checks [#3512](https://github.com/gin-gonic/gin/pull/3512)

#
<原文结束>

# <翻译开始>
# BUG 修复

* 修复 Request.Context() 检查问题 [#3512](https://github.com/gin-gonic/gin/pull/3512)

# <翻译结束>


<原文开始>
SECURITY

* fix lack of escaping of filename in Content-Disposition [#3556](https://github.com/gin-gonic/gin/pull/3556) 

#
<原文结束>

# <翻译开始>
# SECURITY

* 修复Content-Disposition中文件名缺乏转义的问题 [#3556](https://github.com/gin-gonic/gin/pull/3556)

#

# <翻译结束>


<原文开始>
ENHANCEMENTS

* refactor: use bytes.ReplaceAll directly [#3455](https://github.com/gin-gonic/gin/pull/3455)
* convert strings and slices using the officially recommended way [#3344](https://github.com/gin-gonic/gin/pull/3344)
* improve render code coverage [#3525](https://github.com/gin-gonic/gin/pull/3525)

#
<原文结束>

# <翻译开始>
# 增强功能

* 重构：直接使用 bytes.ReplaceAll [#3455](https://github.com/gin-gonic/gin/pull/3455)
* 使用官方推荐的方式转换字符串和切片 [#3344](https://github.com/gin-gonic/gin/pull/3344)
* 提高渲染代码覆盖率 [#3525](https://github.com/gin-gonic/gin/pull/3525)

# <翻译结束>


<原文开始>
DOCS

* docs: changed documentation link for trusted proxies [#3575](https://github.com/gin-gonic/gin/pull/3575)
* chore: improve linting, testing, and GitHub Actions setup [#3583](https://github.com/gin-gonic/gin/pull/3583)


<原文结束>

# <翻译开始>
# DOCS

* 文档：更改了受信任代理的文档链接 [#3575](https://github.com/gin-gonic/gin/pull/3575)
* 构建改进：优化代码格式化检查、测试及 GitHub Actions 配置设置 [#3583](https://github.com/gin-gonic/gin/pull/3583)

# <翻译结束>


<原文开始>
BREAK CHANGES

* Stop useless panicking in context and render [#2150](https://github.com/gin-gonic/gin/pull/2150)

#
<原文结束>

# <翻译开始>
# BREAK CHANGES

* 停止在 context 和 render 中的无用恐慌 [#2150](https://github.com/gin-gonic/gin/pull/2150)

#

# <翻译结束>


<原文开始>
BUG FIXES

* fix(router): tree bug where loop index is not decremented. [#3460](https://github.com/gin-gonic/gin/pull/3460)
* fix(context): panic on NegotiateFormat - index out of range [#3397](https://github.com/gin-gonic/gin/pull/3397)
* Add escape logic for header [#3500](https://github.com/gin-gonic/gin/pull/3500) and [#3503](https://github.com/gin-gonic/gin/pull/3503)

#
<原文结束>

# <翻译开始>
# BUG 修复

* 修复(router)：循环索引未递减导致的树型结构错误。[#3460](https://github.com/gin-gonic/gin/pull/3460)
* 修复(context)：协商格式时引发的 panic（索引超出范围）问题。[#3397](https://github.com/gin-gonic/gin/pull/3397)
* 为头部添加转义逻辑。[#3500](https://github.com/gin-gonic/gin/pull/3500) 和 [#3503](https://github.com/gin-gonic/gin/pull/3503)

##

# <翻译结束>


<原文开始>
SECURITY

* Fix the GO-2022-0969 and GO-2022-0288 vulnerabilities [#3333](https://github.com/gin-gonic/gin/pull/3333)
* fix(security): vulnerability GO-2023-1571 [#3505](https://github.com/gin-gonic/gin/pull/3505)

#
<原文结束>

# <翻译开始>
# SECURITY

* 修复 GO-2022-0969 和 GO-2022-0288 漏洞 [#3333](https://github.com/gin-gonic/gin/pull/3333)
* 修复（安全）：GO-2023-1571 漏洞 [#3505](https://github.com/gin-gonic/gin/pull/3505)

#

# <翻译结束>


<原文开始>
ENHANCEMENTS

* feat: add sonic json support [#3184](https://github.com/gin-gonic/gin/pull/3184)
* chore(file): Creates a directory named path [#3316](https://github.com/gin-gonic/gin/pull/3316)
* fix: modify interface check way [#3327](https://github.com/gin-gonic/gin/pull/3327)
* remove deprecated of package io/ioutil [#3395](https://github.com/gin-gonic/gin/pull/3395)
* refactor: avoid calling strings.ToLower twice [#3343](https://github.com/gin-gonic/gin/pull/3433)
* console logger HTTP status code bug fixed [#3453](https://github.com/gin-gonic/gin/pull/3453)
* chore(yaml): upgrade dependency to v3 version [#3456](https://github.com/gin-gonic/gin/pull/3456)
* chore(router): match method added to routergroup for multiple HTTP methods supporting [#3464](https://github.com/gin-gonic/gin/pull/3464)
* chore(http): add support for go1.20 http.rwUnwrapper to gin.responseWriter [#3489](https://github.com/gin-gonic/gin/pull/3489)

#
<原文结束>

# <翻译开始>
# 增强功能

* 新特性：添加对Sonic JSON的支持 [#3184](https://github.com/gin-gonic/gin/pull/3184)
* 优化（文件）：创建名为path的目录 [#3316](https://github.com/gin-gonic/gin/pull/3316)
* 修复：修改接口检查方式 [#3327](https://github.com/gin-gonic/gin/pull/3327)
* 移除包io/ioutil的废弃内容 [#3395](https://github.com/gin-gonic/gin/pull/3395)
* 重构：避免两次调用strings.ToLower [#3343](https://github.com/gin-gonic/gin/pull/3433)
* 控制台日志记录器已修复HTTP状态码错误 [#3453](https://github.com/gin-gonic/gin/pull/3453)
* 优化（yaml）：升级依赖至v3版本 [#3456](https://github.com/gin-gonic/gin/pull/3456)
* 优化（路由器）：为routergroup添加match方法以支持多种HTTP方法 [#3464](https://github.com/gin-gonic/gin/pull/3464)
* 优化（HTTP）：为gin.responseWriter添加对go1.20 http.rwUnwrapper的支持 [#3489](https://github.com/gin-gonic/gin/pull/3489)

# <翻译结束>


<原文开始>
DOCS

* docs: update markdown format [#3260](https://github.com/gin-gonic/gin/pull/3260)
* docs(readme): Add the TOML rendering example [#3400](https://github.com/gin-gonic/gin/pull/3400)
* docs(readme): move more example to docs/doc.md [#3449](https://github.com/gin-gonic/gin/pull/3449)
* docs: update markdown format [#3446](https://github.com/gin-gonic/gin/pull/3446)


<原文结束>

# <翻译开始>
# DOCS

* 文档：更新Markdown格式 [#3260](https://github.com/gin-gonic/gin/pull/3260)
* 文档(README)：添加TOML渲染示例 [#3400](https://github.com/gin-gonic/gin/pull/3400)
* 文档(README)：将更多示例移动到docs/doc.md文件中 [#3449](https://github.com/gin-gonic/gin/pull/3449)
* 文档：更新Markdown格式 [#3446](https://github.com/gin-gonic/gin/pull/3446)

# <翻译结束>


<原文开始>
BUG FIXES

* fix(route): redirectSlash bug ([#3227]((https://github.com/gin-gonic/gin/pull/3227)))
* fix(engine): missing route params for CreateTestContext ([#2778]((https://github.com/gin-gonic/gin/pull/2778))) ([#2803]((https://github.com/gin-gonic/gin/pull/2803)))

#
<原文结束>

# <翻译开始>
# **错误修复**

* 修复(route)：redirectSlash 问题 ([#3227](https://github.com/gin-gonic/gin/pull/3227))
* 修复(engine)：CreateTestContext 中缺失的路由参数 ([#2778](https://github.com/gin-gonic/gin/pull/2778)) ([#2803](https://github.com/gin-gonic/gin/pull/2803))

---

以下是翻译后的中文内容：

**错误修复**

* 修复(route)：解决redirectSlash错误 ([#3227](https://github.com/gin-gonic/gin/pull/3227))
* 修复(engine)：为CreateTestContext添加缺失的路由参数 ([#2778](https://github.com/gin-gonic/gin/pull/2778))，参考([#2803](https://github.com/gin-gonic/gin/pull/2803))

# <翻译结束>


<原文开始>
SECURITY

* Fix the GO-2022-1144 vulnerability ([#3432]((https://github.com/gin-gonic/gin/pull/3432)))


<原文结束>

# <翻译开始>
# SECURITY

* 修复 GO-2022-1144 漏洞 ([#3432](https://github.com/gin-gonic/gin/pull/3432))

# <翻译结束>


<原文开始>
ENHANCEMENTS

* feat(context): add ContextWithFallback feature flag [#3172](https://github.com/gin-gonic/gin/pull/3172)


<原文结束>

# <翻译开始>
# 增强功能

* 新特性(context): 添加 ContextWithFallback 功能标志 [#3172](https://github.com/gin-gonic/gin/pull/3172)

# <翻译结束>


<原文开始>
BREAK CHANGES

* TrustedProxies: Add default IPv6 support and refactor [#2967](https://github.com/gin-gonic/gin/pull/2967). Please replace `RemoteIP() (net.IP, bool)` with `RemoteIP() net.IP`
* gin.Context with fallback value from gin.Context.Request.Context() [#2751](https://github.com/gin-gonic/gin/pull/2751)

#
<原文结束>

# <翻译开始>
# BREAK CHANGES

* TrustedProxies：添加默认的IPv6支持并重构 [#2967](https://github.com/gin-gonic/gin/pull/2967)。请将 `RemoteIP() (net.IP, bool)` 替换为 `RemoteIP() net.IP`
* gin.Context 现在具有从 gin.Context.Request.Context() 获取的回退值 [#2751](https://github.com/gin-gonic/gin/pull/2751)

#

# <翻译结束>


<原文开始>
BUG FIXES

* Fixed SetOutput() panics on go 1.17 [#2861](https://github.com/gin-gonic/gin/pull/2861)
* Fix: wrong when wildcard follows named param [#2983](https://github.com/gin-gonic/gin/pull/2983)
* Fix: missing sameSite when do context.reset() [#3123](https://github.com/gin-gonic/gin/pull/3123)

#
<原文结束>

# <翻译开始>
# BUG 修复

* 解决了在 go 1.17 上 SetOutput() 函数导致的 panic 问题 [#2861](https://github.com/gin-gonic/gin/pull/2861)
* 修复：通配符紧跟命名参数时出现错误的情况 [#2983](https://github.com/gin-gonic/gin/pull/2983)
* 修复：在执行 context.reset() 时丢失 sameSite 属性的问题 [#3123](https://github.com/gin-gonic/gin/pull/3123)

# <翻译结束>


<原文开始>
ENHANCEMENTS

* Use Header() instead of deprecated HeaderMap [#2694](https://github.com/gin-gonic/gin/pull/2694)
* RouterGroup.Handle regular match optimization of http method [#2685](https://github.com/gin-gonic/gin/pull/2685)
* Add support go-json, another drop-in json replacement [#2680](https://github.com/gin-gonic/gin/pull/2680)
* Use errors.New to replace fmt.Errorf will much better [#2707](https://github.com/gin-gonic/gin/pull/2707)
* Use Duration.Truncate for truncating precision [#2711](https://github.com/gin-gonic/gin/pull/2711)
* Get client IP when using Cloudflare [#2723](https://github.com/gin-gonic/gin/pull/2723)
* Optimize code adjust [#2700](https://github.com/gin-gonic/gin/pull/2700/files)
* Optimize code and reduce code cyclomatic complexity [#2737](https://github.com/gin-gonic/gin/pull/2737)
* Improve sliceValidateError.Error performance [#2765](https://github.com/gin-gonic/gin/pull/2765)
* Support custom struct tag [#2720](https://github.com/gin-gonic/gin/pull/2720)
* Improve router group tests [#2787](https://github.com/gin-gonic/gin/pull/2787)
* Fallback Context.Deadline() Context.Done() Context.Err() to Context.Request.Context() [#2769](https://github.com/gin-gonic/gin/pull/2769)
* Some codes optimize [#2830](https://github.com/gin-gonic/gin/pull/2830) [#2834](https://github.com/gin-gonic/gin/pull/2834) [#2838](https://github.com/gin-gonic/gin/pull/2838) [#2837](https://github.com/gin-gonic/gin/pull/2837) [#2788](https://github.com/gin-gonic/gin/pull/2788) [#2848](https://github.com/gin-gonic/gin/pull/2848) [#2851](https://github.com/gin-gonic/gin/pull/2851) [#2701](https://github.com/gin-gonic/gin/pull/2701)
* TrustedProxies: Add default IPv6 support and refactor [#2967](https://github.com/gin-gonic/gin/pull/2967)
* Test(route): expose performRequest func [#3012](https://github.com/gin-gonic/gin/pull/3012)
* Support h2c with prior knowledge [#1398](https://github.com/gin-gonic/gin/pull/1398)
* Feat attachment filename support utf8 [#3071](https://github.com/gin-gonic/gin/pull/3071)
* Feat: add StaticFileFS [#2749](https://github.com/gin-gonic/gin/pull/2749)
* Feat(context): return GIN Context from Value method [#2825](https://github.com/gin-gonic/gin/pull/2825)
* Feat: automatically SetMode to TestMode when run go test [#3139](https://github.com/gin-gonic/gin/pull/3139)
* Add TOML bining for gin [#3081](https://github.com/gin-gonic/gin/pull/3081)
* IPv6 add default trusted proxies [#3033](https://github.com/gin-gonic/gin/pull/3033)

#
<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空，请提供需要翻译的.MD格式的具体内容。

# <翻译结束>


<原文开始>
DOCS

* Add note about nomsgpack tag to the readme [#2703](https://github.com/gin-gonic/gin/pull/2703)


<原文结束>

# <翻译开始>
# DOCS

* 在自述文件中添加关于nomsgpack标签的注释 [#2703](https://github.com/gin-gonic/gin/pull/2703)

# <翻译结束>


<原文开始>
BUG FIXES

* Fixed X-Forwarded-For unsafe handling of CVE-2020-28483 [#2844](https://github.com/gin-gonic/gin/pull/2844), closed issue [#2862](https://github.com/gin-gonic/gin/issues/2862).
* Tree: updated the code logic for `latestNode` [#2897](https://github.com/gin-gonic/gin/pull/2897), closed issue [#2894](https://github.com/gin-gonic/gin/issues/2894) [#2878](https://github.com/gin-gonic/gin/issues/2878).
* Tree: fixed the misplacement of adding slashes [#2847](https://github.com/gin-gonic/gin/pull/2847), closed issue [#2843](https://github.com/gin-gonic/gin/issues/2843).
* Tree: fixed tsr with mixed static and wildcard paths [#2924](https://github.com/gin-gonic/gin/pull/2924), closed issue [#2918](https://github.com/gin-gonic/gin/issues/2918).

#
<原文结束>

# <翻译开始>
# 错误修复

* 修复了 X-Forwarded-For 对 CVE-2020-28483 不安全的处理方式 [#2844](https://github.com/gin-gonic/gin/pull/2844)，已关闭问题 [#2862](https://github.com/gin-gonic/gin/issues/2862)。
* 树结构：更新了 `latestNode` 的代码逻辑 [#2897](https://github.com/gin-gonic/gin/pull/2897)，已关闭问题 [#2894](https://github.com/gin-gonic/gin/issues/2894) 和 [#2878](https://github.com/gin-gonic/gin/issues/2878)。
* 树结构：修复了添加斜杠时的位置错误 [#2847](https://github.com/gin-gonic/gin/pull/2847)，已关闭问题 [#2843](https://github.com/gin-gonic/gin/issues/2843)。
* 树结构：修复了混合静态路径和通配符路径时的 tsr 问题 [#2924](https://github.com/gin-gonic/gin/pull/2924)，已关闭问题 [#2918](https://github.com/gin-gonic/gin/issues/2918)。

# <翻译结束>


<原文开始>
ENHANCEMENTS

* TrustedProxies: make it backward-compatible [#2887](https://github.com/gin-gonic/gin/pull/2887), closed issue [#2819](https://github.com/gin-gonic/gin/issues/2819).
* TrustedPlatform: provide custom options for another CDN services [#2906](https://github.com/gin-gonic/gin/pull/2906).

#
<原文结束>

# <翻译开始>
# 增强功能

* TrustedProxies：使其向后兼容 [#2887](https://github.com/gin-gonic/gin/pull/2887)，已关闭问题 [#2819](https://github.com/gin-gonic/gin/issues/2819)。
* TrustedPlatform：为其他 CDN 服务提供自定义选项 [#2906](https://github.com/gin-gonic/gin/pull/2906)。

# <翻译结束>


<原文开始>
DOCS

* NoMethod: added usage annotation ([#2832](https://github.com/gin-gonic/gin/pull/2832#issuecomment-929954463)).


<原文结束>

# <翻译开始>
# DOCS

* NoMethod：添加了使用注释（[#2832](https://github.com/gin-gonic/gin/pull/2832#issuecomment-929954463)）。

# <翻译结束>


<原文开始>
BUG FIXES

* bump new release to fix v1.7.5 release error by using v1.7.4 codes.


<原文结束>

# <翻译开始>
# 错误修复

* 利用v1.7.4的代码，更新新版本以修复v1.7.5版本发布时的错误。

# <翻译结束>


<原文开始>
BUG FIXES

* bump new release to fix checksum mismatch


<原文结束>

# <翻译开始>
# BUG 修复

* 更新新版本以修复校验和不匹配的问题

# <翻译结束>


<原文开始>
BUG FIXES

* fix level 1 router match [#2767](https://github.com/gin-gonic/gin/issues/2767), [#2796](https://github.com/gin-gonic/gin/issues/2796)


<原文结束>

# <翻译开始>
# **错误修复**

* 修复一级路由器匹配问题 [#2767](https://github.com/gin-gonic/gin/issues/2767), [#2796](https://github.com/gin-gonic/gin/issues/2796)

# <翻译结束>


<原文开始>
BUG FIXES

* Fix conflict between param and exact path [#2706](https://github.com/gin-gonic/gin/issues/2706). Close issue [#2682](https://github.com/gin-gonic/gin/issues/2682) [#2696](https://github.com/gin-gonic/gin/issues/2696).


<原文结束>

# <翻译开始>
# 错误修复

* 修复参数与精确路径之间的冲突问题 [#2706](https://github.com/gin-gonic/gin/issues/2706)。关闭问题 [#2682](https://github.com/gin-gonic/gin/issues/2682) 和 [#2696](https://github.com/gin-gonic/gin/issues/2696)。

# <翻译结束>


<原文开始>
BUG FIXES

* fix: data race with trustedCIDRs from [#2674](https://github.com/gin-gonic/gin/issues/2674)([#2675](https://github.com/gin-gonic/gin/pull/2675))


<原文结束>

# <翻译开始>
# BUG 修复

* 修复：与 [#2674](https://github.com/gin-gonic/gin/issues/2674) 中的 trustedCIDRs 存在的数据竞争问题 ([#2675](https://github.com/gin-gonic/gin/pull/2675))

# <翻译结束>


<原文开始>
BUG FIXES

* fix compile error from [#2572](https://github.com/gin-gonic/gin/pull/2572) ([#2600](https://github.com/gin-gonic/gin/pull/2600))
* fix: print headers without Authorization header on broken pipe ([#2528](https://github.com/gin-gonic/gin/pull/2528))
* fix(tree): reassign fullpath when register new node ([#2366](https://github.com/gin-gonic/gin/pull/2366))

#
<原文结束>

# <翻译开始>
# 错误修复

* 修复由 [#2572](https://github.com/gin-gonic/gin/pull/2572) 引起的编译错误 ([#2600](https://github.com/gin-gonic/gin/pull/2600))
* 修复：在管道损坏时打印不包含 Authorization 头的头信息 ([#2528](https://github.com/gin-gonic/gin/pull/2528))
* 修复(tree)：注册新节点时重新分配完整路径 ([#2366](https://github.com/gin-gonic/gin/pull/2366))

# <翻译结束>


<原文开始>
ENHANCEMENTS

* Support params and exact routes without creating conflicts ([#2663](https://github.com/gin-gonic/gin/pull/2663))
* chore: improve render string performance ([#2365](https://github.com/gin-gonic/gin/pull/2365))
* Sync route tree to httprouter latest code ([#2368](https://github.com/gin-gonic/gin/pull/2368))
* chore: rename getQueryCache/getFormCache to initQueryCache/initFormCa ([#2375](https://github.com/gin-gonic/gin/pull/2375))
* chore(performance): improve countParams ([#2378](https://github.com/gin-gonic/gin/pull/2378))
* Remove some functions that have the same effect as the bytes package ([#2387](https://github.com/gin-gonic/gin/pull/2387))
* update:SetMode function ([#2321](https://github.com/gin-gonic/gin/pull/2321))
* remove an unused type SecureJSONPrefix ([#2391](https://github.com/gin-gonic/gin/pull/2391))
* Add a redirect sample for POST method ([#2389](https://github.com/gin-gonic/gin/pull/2389))
* Add CustomRecovery builtin middleware ([#2322](https://github.com/gin-gonic/gin/pull/2322))
* binding: avoid 2038 problem on 32-bit architectures ([#2450](https://github.com/gin-gonic/gin/pull/2450))
* Prevent panic in Context.GetQuery() when there is no Request ([#2412](https://github.com/gin-gonic/gin/pull/2412))
* Add GetUint and GetUint64 method on gin.context ([#2487](https://github.com/gin-gonic/gin/pull/2487))
* update content-disposition header to MIME-style ([#2512](https://github.com/gin-gonic/gin/pull/2512))
* reduce allocs and improve the render `WriteString` ([#2508](https://github.com/gin-gonic/gin/pull/2508))
* implement ".Unwrap() error" on Error type ([#2525](https://github.com/gin-gonic/gin/pull/2525)) ([#2526](https://github.com/gin-gonic/gin/pull/2526))
* Allow bind with a map[string]string ([#2484](https://github.com/gin-gonic/gin/pull/2484))
* chore: update tree ([#2371](https://github.com/gin-gonic/gin/pull/2371))
* Support binding for slice/array obj [Rewrite] ([#2302](https://github.com/gin-gonic/gin/pull/2302))
* basic auth: fix timing oracle ([#2609](https://github.com/gin-gonic/gin/pull/2609))
* Add mixed param and non-param paths (port of httprouter[#329](https://github.com/gin-gonic/gin/pull/329)) ([#2663](https://github.com/gin-gonic/gin/pull/2663))
* feat(engine): add trustedproxies and remoteIP ([#2632](https://github.com/gin-gonic/gin/pull/2632))


<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空，请提供需要翻译的.MD格式的具体内容。

# <翻译结束>


<原文开始>
ENHANCEMENTS

  * Improve performance: Change `*sync.RWMutex` to `sync.RWMutex` in context. [#2351](https://github.com/gin-gonic/gin/pull/2351)


<原文结束>

# <翻译开始>
# 增强功能

  * 提升性能：在上下文中将 `*sync.RWMutex` 更改为 `sync.RWMutex`。 [#2351](https://github.com/gin-gonic/gin/pull/2351)

# <翻译结束>


<原文开始>
BUG FIXES

  * fix missing initial sync.RWMutex [#2305](https://github.com/gin-gonic/gin/pull/2305)

#
<原文结束>

# <翻译开始>
# BUG 修复

  * 修复初始化时缺少的 sync.RWMutex 问题 [#2305](https://github.com/gin-gonic/gin/pull/2305)

# <翻译结束>


<原文开始>
ENHANCEMENTS

  * Add set samesite in cookie. [#2306](https://github.com/gin-gonic/gin/pull/2306)


<原文结束>

# <翻译开始>
# 增强功能

  * 添加在cookie中设置samesite属性。 [#2306](https://github.com/gin-gonic/gin/pull/2306)

# <翻译结束>


<原文开始>
BUG FIXES

  * Revert "fix accept incoming network connections" [#2294](https://github.com/gin-gonic/gin/pull/2294)


<原文结束>

# <翻译开始>
# 错误修复

  * 恢复“修复接受传入网络连接” [#2294](https://github.com/gin-gonic/gin/pull/2294)

# <翻译结束>


<原文开始>
BREAKING

  * chore(performance): Improve performance for adding RemoveExtraSlash flag [#2159](https://github.com/gin-gonic/gin/pull/2159)
  * drop support govendor [#2148](https://github.com/gin-gonic/gin/pull/2148)
  * Added support for SameSite cookie flag [#1615](https://github.com/gin-gonic/gin/pull/1615)

#
<原文结束>

# <翻译开始>
# 重大更新

  * 构建任务(性能提升)：为添加 RemoveExtraSlash 标志改善性能 [#2159](https://github.com/gin-gonic/gin/pull/2159)
  * 停止支持 govendor [#2148](https://github.com/gin-gonic/gin/pull/2148)
  * 新增对 SameSite cookie 标志的支持 [#1615](https://github.com/gin-gonic/gin/pull/1615)

#

# <翻译结束>


<原文开始>
FEATURES

  * add yaml negotiation [#2220](https://github.com/gin-gonic/gin/pull/2220)
  * FileFromFS [#2112](https://github.com/gin-gonic/gin/pull/2112)

#
<原文结束>

# <翻译开始>
# 特性

  * 添加yaml协商功能 [#2220](https://github.com/gin-gonic/gin/pull/2220)
  * FileFromFS 功能 [#2112](https://github.com/gin-gonic/gin/pull/2112)

#

# <翻译结束>


<原文开始>
BUG FIXES

  * Unix Socket Handling [#2280](https://github.com/gin-gonic/gin/pull/2280)
  * Use json marshall in context json to fix breaking new line issue. Fixes #2209 [#2228](https://github.com/gin-gonic/gin/pull/2228)
  * fix accept incoming network connections [#2216](https://github.com/gin-gonic/gin/pull/2216)
  * Fixed a bug in the calculation of the maximum number of parameters [#2166](https://github.com/gin-gonic/gin/pull/2166)
  * [FIX] allow empty headers on DataFromReader [#2121](https://github.com/gin-gonic/gin/pull/2121)
  * Add mutex for protect Context.Keys map [#1391](https://github.com/gin-gonic/gin/pull/1391)

#
<原文结束>

# <翻译开始>
# BUG 修复

  * Unix 套接字处理 [#2280](https://github.com/gin-gonic/gin/pull/2280)
  * 使用 json 序列化在 context.json 中修复换行符问题。修复 #2209 [#2228](https://github.com/gin-gonic/gin/pull/2228)
  * 修复接收网络连接问题 [#2216](https://github.com/gin-gonic/gin/pull/2216)
  * 修复计算最大参数数量时的错误 [#2166](https://github.com/gin-gonic/gin/pull/2166)
  * [FIX] 允许 DataFromReader 上的空头信息 [#2121](https://github.com/gin-gonic/gin/pull/2121)
  * 添加互斥锁以保护 Context.Keys 映射 [#1391](https://github.com/gin-gonic/gin/pull/1391)

#

# <翻译结束>


<原文开始>
ENHANCEMENTS

  * Add mitigation for log injection [#2277](https://github.com/gin-gonic/gin/pull/2277)
  * tree: range over nodes values [#2229](https://github.com/gin-gonic/gin/pull/2229)
  * tree: remove duplicate assignment [#2222](https://github.com/gin-gonic/gin/pull/2222)
  * chore: upgrade go-isatty and json-iterator/go [#2215](https://github.com/gin-gonic/gin/pull/2215)
  * path: sync code with httprouter [#2212](https://github.com/gin-gonic/gin/pull/2212)
  * Use zero-copy approach to convert types between string and byte slice [#2206](https://github.com/gin-gonic/gin/pull/2206)
  * Reuse bytes when cleaning the URL paths [#2179](https://github.com/gin-gonic/gin/pull/2179)
  * tree: remove one else statement [#2177](https://github.com/gin-gonic/gin/pull/2177)
  * tree: sync httprouter update (#2173) (#2172) [#2171](https://github.com/gin-gonic/gin/pull/2171)
  * tree: sync part httprouter codes and reduce if/else [#2163](https://github.com/gin-gonic/gin/pull/2163)
  * use http method constant [#2155](https://github.com/gin-gonic/gin/pull/2155)
  * upgrade go-validator to v10 [#2149](https://github.com/gin-gonic/gin/pull/2149)
  * Refactor redirect request in gin.go [#1970](https://github.com/gin-gonic/gin/pull/1970)
  * Add build tag nomsgpack [#1852](https://github.com/gin-gonic/gin/pull/1852)

#
<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空，请提供需要翻译的.MD格式的具体内容。

# <翻译结束>


<原文开始>
DOCS

  * docs(path): improve comments [#2223](https://github.com/gin-gonic/gin/pull/2223)
  * Renew README to fit the modification of SetCookie method [#2217](https://github.com/gin-gonic/gin/pull/2217)
  * Fix spelling [#2202](https://github.com/gin-gonic/gin/pull/2202)
  * Remove broken link from README. [#2198](https://github.com/gin-gonic/gin/pull/2198)
  * Update docs on Context.Done(), Context.Deadline() and Context.Err() [#2196](https://github.com/gin-gonic/gin/pull/2196)
  * Update validator to v10 [#2190](https://github.com/gin-gonic/gin/pull/2190)
  * upgrade go-validator to v10 for README [#2189](https://github.com/gin-gonic/gin/pull/2189)
  * Update to currently output [#2188](https://github.com/gin-gonic/gin/pull/2188)
  * Fix "Custom Validators" example [#2186](https://github.com/gin-gonic/gin/pull/2186)
  * Add project to README [#2165](https://github.com/gin-gonic/gin/pull/2165)
  * docs(benchmarks): for gin v1.5 [#2153](https://github.com/gin-gonic/gin/pull/2153)
  * Changed wording for clarity in README.md [#2122](https://github.com/gin-gonic/gin/pull/2122)

#
<原文结束>

# <翻译开始>
# DOCS

  * docs(path): 提高注释质量 [#2223](https://github.com/gin-gonic/gin/pull/2223)
  * 更新README以适应SetCookie方法的修改 [#2217](https://github.com/gin-gonic/gin/pull/2217)
  * 修正拼写错误 [#2202](https://github.com/gin-gonic/gin/pull/2202)
  * 从README中移除失效链接 [#2198](https://github.com/gin-gonic/gin/pull/2198)
  * 更新关于Context.Done()、Context.Deadline()和Context.Err()的文档说明 [#2196](https://github.com/gin-gonic/gin/pull/2196)
  * 更新validator至v10版本 [#2190](https://github.com/gin-gonic/gin/pull/2190)
  * 将go-validator升级至v10版本（针对README） [#2189](https://github.com/gin-gonic/gin/pull/2189)
  * 更新至当前输出内容 [#2188](https://github.com/gin-gonic/gin/pull/2188)
  * 修复“Custom Validators”示例问题 [#2186](https://github.com/gin-gonic/gin/pull/2186)
  * 在README中添加项目信息 [#2165](https://github.com/gin-gonic/gin/pull/2165)
  * docs(benchmarks): 针对gin v1.5版本的基准测试 [#2153](https://github.com/gin-gonic/gin/pull/2153)
  * 修改README.md中的措辞以提高清晰度 [#2122](https://github.com/gin-gonic/gin/pull/2122)

#

# <翻译结束>


<原文开始>
MISC

  * ci support go1.14 [#2262](https://github.com/gin-gonic/gin/pull/2262)
  * chore: upgrade depend version [#2231](https://github.com/gin-gonic/gin/pull/2231)
  * Drop support go1.10 [#2147](https://github.com/gin-gonic/gin/pull/2147)
  * fix comment in `mode.go` [#2129](https://github.com/gin-gonic/gin/pull/2129)


<原文结束>

# <翻译开始>
# MISC

  * 支持 go1.14 [#2262](https://github.com/gin-gonic/gin/pull/2262)
  * 构建任务：升级依赖版本 [#2231](https://github.com/gin-gonic/gin/pull/2231)
  * 停止支持 go1.10 [#2147](https://github.com/gin-gonic/gin/pull/2147)
  * 修复 `mode.go` 中的注释 [#2129](https://github.com/gin-gonic/gin/pull/2129)

# <翻译结束>


<原文开始>
Gin v1.5.0

- [FIX] Use DefaultWriter and DefaultErrorWriter for debug messages [#1891](https://github.com/gin-gonic/gin/pull/1891)
- [NEW] Now you can parse the inline lowercase start structure [#1893](https://github.com/gin-gonic/gin/pull/1893)
- [FIX] Some code improvements [#1909](https://github.com/gin-gonic/gin/pull/1909)
- [FIX] Use encode replace json marshal increase json encoder speed [#1546](https://github.com/gin-gonic/gin/pull/1546)
- [NEW] Hold matched route full path in the Context [#1826](https://github.com/gin-gonic/gin/pull/1826)
- [FIX] Fix context.Params race condition on Copy() [#1841](https://github.com/gin-gonic/gin/pull/1841)
- [NEW] Add context param query cache [#1450](https://github.com/gin-gonic/gin/pull/1450)
- [FIX] Improve GetQueryMap performance [#1918](https://github.com/gin-gonic/gin/pull/1918)
- [FIX] Improve get post data [#1920](https://github.com/gin-gonic/gin/pull/1920)
- [FIX] Use context instead of x/net/context [#1922](https://github.com/gin-gonic/gin/pull/1922)
- [FIX] Attempt to fix PostForm cache bug [#1931](https://github.com/gin-gonic/gin/pull/1931)
- [NEW] Add support of multipart multi files [#1949](https://github.com/gin-gonic/gin/pull/1949)
- [NEW] Support bind http header param [#1957](https://github.com/gin-gonic/gin/pull/1957)
- [FIX] Drop support for go1.8 and go1.9 [#1933](https://github.com/gin-gonic/gin/pull/1933)
- [FIX] Bugfix for the FullPath feature [#1919](https://github.com/gin-gonic/gin/pull/1919)
- [FIX] Gin1.5 bytes.Buffer to strings.Builder [#1939](https://github.com/gin-gonic/gin/pull/1939)
- [FIX] Upgrade github.com/ugorji/go/codec [#1969](https://github.com/gin-gonic/gin/pull/1969)
- [NEW] Support bind unix time [#1980](https://github.com/gin-gonic/gin/pull/1980)
- [FIX] Simplify code [#2004](https://github.com/gin-gonic/gin/pull/2004)
- [NEW] Support negative Content-Length in DataFromReader [#1981](https://github.com/gin-gonic/gin/pull/1981)
- [FIX] Identify terminal on a RISC-V architecture for auto-colored logs [#2019](https://github.com/gin-gonic/gin/pull/2019)
- [BREAKING] `Context.JSONP()` now expects a semicolon (`;`) at the end [#2007](https://github.com/gin-gonic/gin/pull/2007)
- [BREAKING] Upgrade default `binding.Validator` to v9 (see [its changelog](https://github.com/go-playground/validator/releases/tag/v9.0.0)) [#1015](https://github.com/gin-gonic/gin/pull/1015)
- [NEW] Add `DisallowUnknownFields()` in `Context.BindJSON()` [#2028](https://github.com/gin-gonic/gin/pull/2028)
- [NEW] Use specific `net.Listener` with `Engine.RunListener()` [#2023](https://github.com/gin-gonic/gin/pull/2023)
- [FIX] Fix some typo [#2079](https://github.com/gin-gonic/gin/pull/2079) [#2080](https://github.com/gin-gonic/gin/pull/2080)
- [FIX] Relocate binding body tests [#2086](https://github.com/gin-gonic/gin/pull/2086)
- [FIX] Use Writer in Context.Status [#1606](https://github.com/gin-gonic/gin/pull/1606)
- [FIX] `Engine.RunUnix()` now returns the error if it can't change the file mode [#2093](https://github.com/gin-gonic/gin/pull/2093)
- [FIX] `RouterGroup.StaticFS()` leaked files. Now it closes them. [#2118](https://github.com/gin-gonic/gin/pull/2118)
- [FIX] `Context.Request.FormFile` leaked file. Now it closes it. [#2114](https://github.com/gin-gonic/gin/pull/2114)
- [FIX] Ignore walking on `form:"-"` mapping [#1943](https://github.com/gin-gonic/gin/pull/1943)

#
<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空，请提供需要翻译的.MD格式的具体内容。

# <翻译结束>


<原文开始>
Gin v1.4.0

- [NEW] Support for [Go Modules](https://github.com/golang/go/wiki/Modules)  [#1569](https://github.com/gin-gonic/gin/pull/1569)
- [NEW] Refactor of form mapping multipart request [#1829](https://github.com/gin-gonic/gin/pull/1829)
- [FIX] Truncate Latency precision in long running request [#1830](https://github.com/gin-gonic/gin/pull/1830)
- [FIX] IsTerm flag should not be affected by DisableConsoleColor method. [#1802](https://github.com/gin-gonic/gin/pull/1802)
- [NEW] Supporting file binding [#1264](https://github.com/gin-gonic/gin/pull/1264)
- [NEW] Add support for mapping arrays [#1797](https://github.com/gin-gonic/gin/pull/1797)
- [FIX] Readme updates [#1793](https://github.com/gin-gonic/gin/pull/1793) [#1788](https://github.com/gin-gonic/gin/pull/1788) [1789](https://github.com/gin-gonic/gin/pull/1789)
- [FIX] StaticFS: Fixed Logging two log lines on 404.  [#1805](https://github.com/gin-gonic/gin/pull/1805), [#1804](https://github.com/gin-gonic/gin/pull/1804)
- [NEW] Make context.Keys available as LogFormatterParams [#1779](https://github.com/gin-gonic/gin/pull/1779)
- [NEW] Use internal/json for Marshal/Unmarshal [#1791](https://github.com/gin-gonic/gin/pull/1791)
- [NEW] Support mapping time.Duration [#1794](https://github.com/gin-gonic/gin/pull/1794)
- [NEW] Refactor form mappings [#1749](https://github.com/gin-gonic/gin/pull/1749)
- [NEW] Added flag to context.Stream indicates if client disconnected in middle of stream [#1252](https://github.com/gin-gonic/gin/pull/1252)
- [FIX] Moved [examples](https://github.com/gin-gonic/examples) to stand alone Repo [#1775](https://github.com/gin-gonic/gin/pull/1775)
- [NEW] Extend context.File to allow for the content-disposition attachments via a new method context.Attachment [#1260](https://github.com/gin-gonic/gin/pull/1260)
- [FIX] Support HTTP content negotiation wildcards [#1112](https://github.com/gin-gonic/gin/pull/1112)
- [NEW] Add prefix from X-Forwarded-Prefix in redirectTrailingSlash [#1238](https://github.com/gin-gonic/gin/pull/1238)
- [FIX] context.Copy() race condition [#1020](https://github.com/gin-gonic/gin/pull/1020)
- [NEW] Add context.HandlerNames() [#1729](https://github.com/gin-gonic/gin/pull/1729)
- [FIX] Change color methods to public in the defaultLogger. [#1771](https://github.com/gin-gonic/gin/pull/1771)
- [FIX] Update writeHeaders method to use http.Header.Set [#1722](https://github.com/gin-gonic/gin/pull/1722)
- [NEW] Add response size to LogFormatterParams [#1752](https://github.com/gin-gonic/gin/pull/1752)
- [NEW] Allow ignoring field on form mapping [#1733](https://github.com/gin-gonic/gin/pull/1733)
- [NEW] Add a function to force color in console output. [#1724](https://github.com/gin-gonic/gin/pull/1724)
- [FIX] Context.Next() - recheck len of handlers on every iteration. [#1745](https://github.com/gin-gonic/gin/pull/1745)
- [FIX] Fix all errcheck warnings [#1739](https://github.com/gin-gonic/gin/pull/1739) [#1653](https://github.com/gin-gonic/gin/pull/1653)
- [NEW] context: inherits context cancellation and deadline from http.Request context for Go>=1.7 [#1690](https://github.com/gin-gonic/gin/pull/1690)
- [NEW] Binding for URL Params [#1694](https://github.com/gin-gonic/gin/pull/1694)
- [NEW] Add LoggerWithFormatter method [#1677](https://github.com/gin-gonic/gin/pull/1677)
- [FIX] CI testing updates [#1671](https://github.com/gin-gonic/gin/pull/1671) [#1670](https://github.com/gin-gonic/gin/pull/1670) [#1682](https://github.com/gin-gonic/gin/pull/1682) [#1669](https://github.com/gin-gonic/gin/pull/1669)
- [FIX] StaticFS(): Send 404 when path does not exist [#1663](https://github.com/gin-gonic/gin/pull/1663)
- [FIX] Handle nil body for JSON binding [#1638](https://github.com/gin-gonic/gin/pull/1638)
- [FIX] Support bind uri param [#1612](https://github.com/gin-gonic/gin/pull/1612)
- [FIX] recovery: fix issue with syscall import on google app engine [#1640](https://github.com/gin-gonic/gin/pull/1640)
- [FIX] Make sure the debug log contains line breaks [#1650](https://github.com/gin-gonic/gin/pull/1650)
- [FIX] Panic stack trace being printed during recovery of broken pipe [#1089](https://github.com/gin-gonic/gin/pull/1089) [#1259](https://github.com/gin-gonic/gin/pull/1259)
- [NEW] RunFd method to run http.Server through a file descriptor [#1609](https://github.com/gin-gonic/gin/pull/1609)
- [NEW] Yaml binding support [#1618](https://github.com/gin-gonic/gin/pull/1618)
- [FIX] Pass MaxMultipartMemory when FormFile is called [#1600](https://github.com/gin-gonic/gin/pull/1600)
- [FIX] LoadHTML* tests [#1559](https://github.com/gin-gonic/gin/pull/1559)
- [FIX] Removed use of sync.pool from HandleContext [#1565](https://github.com/gin-gonic/gin/pull/1565)
- [FIX] Format output log to os.Stderr [#1571](https://github.com/gin-gonic/gin/pull/1571)
- [FIX] Make logger use a yellow background and a darkgray text for legibility [#1570](https://github.com/gin-gonic/gin/pull/1570)
- [FIX] Remove sensitive request information from panic log. [#1370](https://github.com/gin-gonic/gin/pull/1370)
- [FIX] log.Println() does not print timestamp [#829](https://github.com/gin-gonic/gin/pull/829) [#1560](https://github.com/gin-gonic/gin/pull/1560)
- [NEW] Add PureJSON renderer [#694](https://github.com/gin-gonic/gin/pull/694)
- [FIX] Add missing copyright and update if/else [#1497](https://github.com/gin-gonic/gin/pull/1497)
- [FIX] Update msgpack usage [#1498](https://github.com/gin-gonic/gin/pull/1498)
- [FIX] Use protobuf on render [#1496](https://github.com/gin-gonic/gin/pull/1496)
- [FIX] Add support for Protobuf format response [#1479](https://github.com/gin-gonic/gin/pull/1479)
- [NEW] Set default time format in form binding [#1487](https://github.com/gin-gonic/gin/pull/1487)
- [FIX] Add BindXML and ShouldBindXML [#1485](https://github.com/gin-gonic/gin/pull/1485)
- [NEW] Upgrade dependency libraries [#1491](https://github.com/gin-gonic/gin/pull/1491)



<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空，请提供需要翻译的.MD格式的具体内容。

# <翻译结束>


<原文开始>
Gin v1.3.0

- [NEW] Add [`func (*Context) QueryMap`](https://godoc.org/github.com/gin-gonic/gin#Context.QueryMap), [`func (*Context) GetQueryMap`](https://godoc.org/github.com/gin-gonic/gin#Context.GetQueryMap), [`func (*Context) PostFormMap`](https://godoc.org/github.com/gin-gonic/gin#Context.PostFormMap) and [`func (*Context) GetPostFormMap`](https://godoc.org/github.com/gin-gonic/gin#Context.GetPostFormMap) to support `type map[string]string` as query string or form parameters, see [#1383](https://github.com/gin-gonic/gin/pull/1383)
- [NEW] Add [`func (*Context) AsciiJSON`](https://godoc.org/github.com/gin-gonic/gin#Context.AsciiJSON), see [#1358](https://github.com/gin-gonic/gin/pull/1358)
- [NEW] Add `Pusher()` in [`type ResponseWriter`](https://godoc.org/github.com/gin-gonic/gin#ResponseWriter) for supporting http2 push, see [#1273](https://github.com/gin-gonic/gin/pull/1273)
- [NEW] Add [`func (*Context) DataFromReader`](https://godoc.org/github.com/gin-gonic/gin#Context.DataFromReader) for serving dynamic data, see [#1304](https://github.com/gin-gonic/gin/pull/1304)
- [NEW] Add [`func (*Context) ShouldBindBodyWith`](https://godoc.org/github.com/gin-gonic/gin#Context.ShouldBindBodyWith) allowing to call binding multiple times, see [#1341](https://github.com/gin-gonic/gin/pull/1341)
- [NEW] Support pointers in form binding, see [#1336](https://github.com/gin-gonic/gin/pull/1336)
- [NEW] Add [`func (*Context) JSONP`](https://godoc.org/github.com/gin-gonic/gin#Context.JSONP), see [#1333](https://github.com/gin-gonic/gin/pull/1333)
- [NEW] Support default value in form binding, see [#1138](https://github.com/gin-gonic/gin/pull/1138)
- [NEW] Expose validator engine in [`type StructValidator`](https://godoc.org/github.com/gin-gonic/gin/binding#StructValidator), see [#1277](https://github.com/gin-gonic/gin/pull/1277)
- [NEW] Add [`func (*Context) ShouldBind`](https://godoc.org/github.com/gin-gonic/gin#Context.ShouldBind), [`func (*Context) ShouldBindQuery`](https://godoc.org/github.com/gin-gonic/gin#Context.ShouldBindQuery) and [`func (*Context) ShouldBindJSON`](https://godoc.org/github.com/gin-gonic/gin#Context.ShouldBindJSON), see [#1047](https://github.com/gin-gonic/gin/pull/1047)
- [NEW] Add support for `time.Time` location in form binding, see [#1117](https://github.com/gin-gonic/gin/pull/1117)
- [NEW] Add [`func (*Context) BindQuery`](https://godoc.org/github.com/gin-gonic/gin#Context.BindQuery), see [#1029](https://github.com/gin-gonic/gin/pull/1029)
- [NEW] Make [jsonite](https://github.com/json-iterator/go) optional with build tags, see [#1026](https://github.com/gin-gonic/gin/pull/1026)
- [NEW] Show query string in logger, see [#999](https://github.com/gin-gonic/gin/pull/999)
- [NEW] Add [`func (*Context) SecureJSON`](https://godoc.org/github.com/gin-gonic/gin#Context.SecureJSON), see [#987](https://github.com/gin-gonic/gin/pull/987) and [#993](https://github.com/gin-gonic/gin/pull/993)
- [DEPRECATE] `func (*Context) GetCookie` for [`func (*Context) Cookie`](https://godoc.org/github.com/gin-gonic/gin#Context.Cookie)
- [FIX] Don't display color tags if [`func DisableConsoleColor`](https://godoc.org/github.com/gin-gonic/gin#DisableConsoleColor) called, see [#1072](https://github.com/gin-gonic/gin/pull/1072)
- [FIX] Gin Mode `""` when calling [`func Mode`](https://godoc.org/github.com/gin-gonic/gin#Mode) now returns `const DebugMode`, see [#1250](https://github.com/gin-gonic/gin/pull/1250)
- [FIX] `Flush()` now doesn't overwrite `responseWriter` status code, see [#1460](https://github.com/gin-gonic/gin/pull/1460)


<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空，请提供需要翻译的.MD格式的具体内容。

# <翻译结束>


<原文开始>
Gin 1.2.0

- [NEW] Switch from godeps to govendor
- [NEW] Add support for Let's Encrypt via gin-gonic/autotls
- [NEW] Improve README examples and add extra at examples folder
- [NEW] Improved support with App Engine
- [NEW] Add custom template delimiters, see #860
- [NEW] Add Template Func Maps, see #962
- [NEW] Add \*context.Handler(), see #928
- [NEW] Add \*context.GetRawData()
- [NEW] Add \*context.GetHeader() (request)
- [NEW] Add \*context.AbortWithStatusJSON() (JSON content type)
- [NEW] Add \*context.Keys type cast helpers
- [NEW] Add \*context.ShouldBindWith()
- [NEW] Add \*context.MustBindWith()
- [NEW] Add \*engine.SetFuncMap()
- [DEPRECATE] On next release: \*context.BindWith(), see #855
- [FIX] Refactor render
- [FIX] Reworked tests
- [FIX] logger now supports cygwin
- [FIX] Use X-Forwarded-For before X-Real-Ip
- [FIX] time.Time binding (#904)


<原文结束>

# <翻译开始>
# Gin 1.2.0 更新内容：

- [新特性] 从godeps切换到govendor
- [新特性] 添加对Let's Encrypt的支持，通过gin-gonic/autotls
- [新特性] 改进README示例，并在examples文件夹中添加额外示例
- [新特性] 提高与App Engine的兼容性
- [新特性] 添加自定义模板分隔符功能，参见#860
- [新特性] 添加模板Func Maps功能，参见#962
- [新特性] 添加*\context.Handler()方法，参见#928
- [新特性] 添加*\context.GetRawData()方法
- [新特性] 添加获取请求头信息的*\context.GetHeader()方法
- [新特性] 添加*\context.AbortWithStatusJSON()方法（设置JSON内容类型）
- [新特性] 添加*\context.Keys类型的转换辅助方法
- [新特性] 添加*\context.ShouldBindWith()方法
- [新特性] 添加*\context.MustBindWith()方法
- [新特性] 添加*\engine.SetFuncMap()方法
- [弃用] 在下个版本中将弃用*\context.BindWith()方法，参见#855
- [修复] 重构渲染模块
- [修复] 重新编写测试
- [修复] 日志模块现支持cygwin环境
- [修复] 先使用X-Forwarded-For，再使用X-Real-Ip
- [修复] 修复时间类型(time.Time)绑定问题(#904)

# <翻译结束>


<原文开始>
Gin 1.1.4

- [NEW] Support google appengine for IsTerminal func


<原文结束>

# <翻译开始>
# Gin 1.1.4

- [新特性] 在 IsTerminal 函数中支持 Google App Engine

# <翻译结束>


<原文开始>
Gin 1.1.3

- [FIX] Reverted Logger: skip ANSI color commands


<原文结束>

# <翻译开始>
# Gin 1.1.3

- [修复] 恢复日志器：跳过 ANSI 颜色命令

# <翻译结束>


<原文开始>
Gin 1.1

- [NEW] Implement QueryArray and PostArray methods
- [NEW] Refactor GetQuery and GetPostForm
- [NEW] Add contribution guide
- [FIX] Corrected typos in README
- [FIX] Removed additional Iota
- [FIX] Changed imports to gopkg instead of github in README (#733)
- [FIX] Logger: skip ANSI color commands if output is not a tty


<原文结束>

# <翻译开始>
# Gin 1.1

- [新功能] 实现 QueryArray 和 PostArray 方法
- [重构] 优化 GetQuery 和 GetPostForm
- [新增] 添加贡献指南
- [修复] 修正 README 中的拼写错误
- [修复] 移除多余的 Iota
- [修复] 将 README 中的导入路径从 github 改为 gopkg (#733)
- [修复] Logger：如果输出不是 tty，则跳过 ANSI 颜色命令

# <翻译结束>


<原文开始>
Gin 1.0rc2 (...)

- [PERFORMANCE] Fast path for writing Content-Type.
- [PERFORMANCE] Much faster 404 routing
- [PERFORMANCE] Allocation optimizations
- [PERFORMANCE] Faster root tree lookup
- [PERFORMANCE] Zero overhead, String() and JSON() rendering.
- [PERFORMANCE] Faster ClientIP parsing
- [PERFORMANCE] Much faster SSE implementation
- [NEW] Benchmarks suite
- [NEW] Bind validation can be disabled and replaced with custom validators.
- [NEW] More flexible HTML render
- [NEW] Multipart and PostForm bindings
- [NEW] Adds method to return all the registered routes
- [NEW] Context.HandlerName() returns the main handler's name
- [NEW] Adds Error.IsType() helper
- [FIX] Binding multipart form
- [FIX] Integration tests
- [FIX] Crash when binding non struct object in Context.
- [FIX] RunTLS() implementation
- [FIX] Logger() unit tests
- [FIX] Adds SetHTMLTemplate() warning
- [FIX] Context.IsAborted()
- [FIX] More unit tests
- [FIX] JSON, XML, HTML renders accept custom content-types
- [FIX] gin.AbortIndex is unexported
- [FIX] Better approach to avoid directory listing in StaticFS()
- [FIX] Context.ClientIP() always returns the IP with trimmed spaces.
- [FIX] Better warning when running in debug mode.
- [FIX] Google App Engine integration. debugPrint does not use os.Stdout
- [FIX] Fixes integer overflow in error type
- [FIX] Error implements the json.Marshaller interface
- [FIX] MIT license in every file



<原文结束>

# <翻译开始>
# Gin 1.0rc2 (...)

- [性能] 快速设置Content-Type的路径
- [性能] 大幅提升404路由速度
- [性能] 内存分配优化
- [性能] 加快根树查找速度
- [性能] 零开销，String()和JSON()渲染方法
- [性能] 提高ClientIP解析速度
- [性能] 更快速的SSE（Server-Sent Events）实现
- [新增] 性能基准测试套件
- [新增] 可禁用绑定验证并替换为自定义验证器
- [新增] 更灵活的HTML渲染功能
- [新增] 支持Multipart和PostForm绑定
- [新增] 新增返回所有已注册路由的方法
- [新增] Context.HandlerName() 返回主处理程序名称
- [新增] 添加Error.IsType()辅助方法
- [修复] 绑定multipart表单问题
- [修复] 集成测试问题
- [修复] 当在Context中绑定非结构体对象时导致崩溃的问题
- [修复] RunTLS() 方法实现问题
- [修复] Logger() 单元测试问题
- [修复] 添加SetHTMLTemplate()警告提示
- [修复] Context.IsAborted() 方法
- [修复] 增加更多单元测试
- [修复] JSON、XML、HTML渲染接受自定义内容类型
- [修复] gin.AbortIndex 现为未导出状态
- [修复] StaticFS() 中改进避免目录列表显示的方法
- [修复] Context.ClientIP() 方法始终返回去除空格后的IP地址
- [修复] 调试模式下的更友好警告信息
- [修复] Google App Engine集成，debugPrint不再使用os.Stdout
- [修复] 修复错误类型中的整数溢出问题
- [修复] Error 类型实现json.Marshaller接口
- [修复] 每个文件采用MIT许可证

# <翻译结束>


<原文开始>
Gin 1.0rc1 (May 22, 2015)

- [PERFORMANCE] Zero allocation router
- [PERFORMANCE] Faster JSON, XML and text rendering
- [PERFORMANCE] Custom hand optimized HttpRouter for Gin
- [PERFORMANCE] Misc code optimizations. Inlining, tail call optimizations
- [NEW] Built-in support for golang.org/x/net/context
- [NEW] Any(path, handler). Create a route that matches any path
- [NEW] Refactored rendering pipeline (faster and static typed)
- [NEW] Refactored errors API
- [NEW] IndentedJSON() prints pretty JSON
- [NEW] Added gin.DefaultWriter
- [NEW] UNIX socket support
- [NEW] RouterGroup.BasePath is exposed
- [NEW] JSON validation using go-validate-yourself (very powerful options)
- [NEW] Completed suite of unit tests
- [NEW] HTTP streaming with c.Stream()
- [NEW] StaticFile() creates a router for serving just one file.
- [NEW] StaticFS() has an option to disable directory listing.
- [NEW] StaticFS() for serving static files through virtual filesystems
- [NEW] Server-Sent Events native support
- [NEW] WrapF() and WrapH() helpers for wrapping http.HandlerFunc and http.Handler
- [NEW] Added LoggerWithWriter() middleware
- [NEW] Added RecoveryWithWriter() middleware
- [NEW] Added DefaultPostFormValue()
- [NEW] Added DefaultFormValue()
- [NEW] Added DefaultParamValue()
- [FIX] BasicAuth() when using custom realm
- [FIX] Bug when serving static files in nested routing group
- [FIX] Redirect using built-in http.Redirect()
- [FIX] Logger when printing the requested path
- [FIX] Documentation typos
- [FIX] Context.Engine renamed to Context.engine
- [FIX] Better debugging messages
- [FIX] ErrorLogger
- [FIX] Debug HTTP render
- [FIX] Refactored binding and render modules
- [FIX] Refactored Context initialization
- [FIX] Refactored BasicAuth()
- [FIX] NoMethod/NoRoute handlers
- [FIX] Hijacking http
- [FIX] Better support for Google App Engine (using log instead of fmt)



<原文结束>

# <翻译开始>
# 很抱歉，您提供的内容为空。请提供具体的MD格式内容以便我为您翻译成中文。

# <翻译结束>


<原文开始>
Gin 0.6 (Mar 9, 2015)

- [NEW] Support multipart/form-data
- [NEW] NoMethod handler
- [NEW] Validate sub structures
- [NEW] Support for HTTP Realm Auth
- [FIX] Unsigned integers in binding
- [FIX] Improve color logger



<原文结束>

# <翻译开始>
# Gin 0.6 (2015年3月9日)

- [新特性] 支持 multipart/form-data
- [新特性] 新增 NoMethod 处理器
- [新特性] 验证子结构体
- [新特性] 支持 HTTP 实体认证（HTTP Realm Auth）
- [修复] 绑定中的无符号整数问题
- [修复] 提升颜色日志记录器的表现

# <翻译结束>


<原文开始>
Gin 0.5 (Feb 7, 2015)

- [NEW] Content Negotiation
- [FIX] Solved security bug that allow a client to spoof ip
- [FIX] Fix unexported/ignored fields in binding



<原文结束>

# <翻译开始>
# Gin 0.5（2015年2月7日）

- [新功能] 内容协商
- [修复] 解决了允许客户端伪造IP的安全漏洞
- [修复] 修复绑定中未导出/被忽略的字段问题

# <翻译结束>


<原文开始>
Gin 0.4 (Aug 21, 2014)

- [NEW] Development mode
- [NEW] Unit tests
- [NEW] Add Content.Redirect()
- [FIX] Deferring WriteHeader()
- [FIX] Improved documentation for model binding



<原文结束>

# <翻译开始>
# Gin 0.4（2014年8月21日）

- [新增] 开发模式
- [新增] 单元测试
- [新增] 添加 Content.Redirect() 方法
- [修复] 延迟 WriteHeader() 调用
- [改进] 对模型绑定的文档进行了优化

# <翻译结束>


<原文开始>
Gin 0.3 (Jul 18, 2014)

- [PERFORMANCE] Normal log and error log are printed in the same call.
- [PERFORMANCE] Improve performance of NoRouter()
- [PERFORMANCE] Improve context's memory locality, reduce CPU cache faults.
- [NEW] Flexible rendering API
- [NEW] Add Context.File()
- [NEW] Add shortcut RunTLS() for http.ListenAndServeTLS
- [FIX] Rename NotFound404() to NoRoute()
- [FIX] Errors in context are purged
- [FIX] Adds HEAD method in Static file serving
- [FIX] Refactors Static() file serving
- [FIX] Using keyed initialization to fix app-engine integration
- [FIX] Can't unmarshal JSON array, #63
- [FIX] Renaming Context.Req to Context.Request
- [FIX] Check application/x-www-form-urlencoded when parsing form



<原文结束>

# <翻译开始>
# Gin 0.3 (2014年7月18日)

- [性能优化] 正常日志和错误日志在同一个调用中打印。
- [性能优化] 提高 NoRouter() 的性能表现。
- [性能优化] 改善上下文的内存局部性，减少 CPU 缓存故障。
- [新特性] 引入灵活的渲染 API。
- [新特性] 添加 Context.File() 方法。
- [新特性] 添加快捷方法 RunTLS() 用于 http.ListenAndServeTLS。
- [修复] 将 NotFound404() 重命名为 NoRoute()。
- [修复] 清除上下文中的错误信息。
- [修复] 在静态文件服务中增加 HEAD 请求方法支持。
- [修复] 重构 Static() 静态文件服务功能。
- [修复] 使用键控初始化以解决与 app-engine 的集成问题。
- [修复] 无法反序列化 JSON 数组，问题 #63。
- [修复] 将 Context.Req 重命名为 Context.Request。
- [修复] 在解析表单时检查 application/x-www-form-urlencoded 类型。

# <翻译结束>


<原文开始>
Gin 0.2b (Jul 08, 2014)
- [PERFORMANCE] Using sync.Pool to allocatio/gc overhead
- [NEW] Travis CI integration
- [NEW] Completely new logger
- [NEW] New API for serving static files. gin.Static()
- [NEW] gin.H() can be serialized into XML
- [NEW] Typed errors. Errors can be typed. Internet/external/custom.
- [NEW] Support for Godeps
- [NEW] Travis/Godocs badges in README
- [NEW] New Bind() and BindWith() methods for parsing request body.
- [NEW] Add Content.Copy()
- [NEW] Add context.LastError()
- [NEW] Add shortcut for OPTIONS HTTP method
- [FIX] Tons of README fixes
- [FIX] Header is written before body
- [FIX] BasicAuth() and changes API a little bit
- [FIX] Recovery() middleware only prints panics
- [FIX] Context.Get() does not panic anymore. Use MustGet() instead.
- [FIX] Multiple http.WriteHeader() in NotFound handlers
- [FIX] Engine.Run() panics if http server can't be set up
- [FIX] Crash when route path doesn't start with '/'
- [FIX] Do not update header when status code is negative
- [FIX] Setting response headers before calling WriteHeader in context.String()
- [FIX] Add MIT license
- [FIX] Changes behaviour of ErrorLogger() and Logger()

<原文结束>

# <翻译开始>
# Gin 0.2b (2014年7月8日)
- [性能优化] 使用sync.Pool降低内存分配和垃圾回收开销
- [新增] 集成Travis CI
- [全新] 完全重新设计的日志系统
- [新增] 新的静态文件服务API：gin.Static()
- [新增] gin.H() 现在可以被序列化为XML
- [新增] 支持类型错误。错误现在可以具有类型，如互联网错误、外部错误或自定义错误。
- [新增] 支持Godeps
- [新增] 在README中添加Travis和Godocs徽章
- [新增] 新增用于解析请求体的Bind()和BindWith()方法
- [新增] 添加Content.Copy()方法
- [新增] 添加context.LastError()方法
- [新增] 为OPTIONS HTTP方法添加快捷方式
- [修复] 大量README文档修正
- [修复] 先写入Header再写入body
- [修复] BasicAuth() 方法及其API小幅调整
- [修复] Recovery() 中间件仅打印panic信息
- [修复] Context.Get() 不再引发panic，改用MustGet()代替
- [修复] NotFound处理器中的多次http.WriteHeader()调用问题
- [修复] 若无法设置HTTP服务器，Engine.Run()将引发panic
- [修复] 路由路径不以'/'开头时导致程序崩溃的问题
- [修复] 当状态码为负数时不更新header
- [修复] 在context.String()中先调用WriteHeader再设置响应头
- [修复] 添加MIT许可证
- [修复] 更改ErrorLogger()和Logger()的行为

# <翻译结束>

