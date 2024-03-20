
<原文开始>
// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
<原文结束>

# <翻译开始>
// 版权所有 ? 2014 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。
# <翻译结束>


<原文开始>
// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
<原文结束>

# <翻译开始>
// NewAuthorizer 返回一个鉴权器，使用 Casbin 执行器作为输入参数
# <翻译结束>


<原文开始>
// BasicAuthorizer stores the casbin handler
<原文结束>

# <翻译开始>
// BasicAuthorizer 基础授权器存储了 casbin 处理器
# <翻译结束>


<原文开始>
// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
<原文结束>

# <翻译开始>
// GetUserName 从请求中获取用户名。
// 当前仅支持HTTP基本认证。
# <翻译结束>


<原文开始>
// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
<原文结束>

# <翻译开始>
// CheckPermission 检查请求中的用户/方法/路径组合。
// 返回 true（权限授予）或 false（权限禁止）
# <翻译结束>


<原文开始>
// RequirePermission returns the 403 Forbidden to the client
<原文结束>

# <翻译开始>
// RequirePermission 返回 403 Forbidden 给客户端
# <翻译结束>

