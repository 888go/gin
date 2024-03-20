
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
// AuthUserKey is the cookie name for user credential in basic auth.
<原文结束>

# <翻译开始>
// AuthUserKey 是基本认证中用于存储用户凭证的cookie名称。
# <翻译结束>


<原文开始>
// Accounts defines a key/value for user/pass list of authorized logins.
<原文结束>

# <翻译开始>
// Accounts 定义了一个用于存储授权登录用户/密码键值对的列表。
# <翻译结束>


<原文开始>
// BasicAuthForRealm returns a Basic HTTP Authorization middleware. It takes as arguments a map[string]string where
// the key is the user name and the value is the password, as well as the name of the Realm.
// If the realm is empty, "Authorization Required" will be used by default.
// (see http://tools.ietf.org/html/rfc2617#section-1.2)
<原文结束>

# <翻译开始>
// BasicAuthForRealm 返回一个基础HTTP身份验证中间件。它接受两个参数：一个map[string]string，其中键是用户名，值是密码；以及一个realm（领域）名称。
// 如果realm为空，则默认使用"Authorization Required"。
// （参见http://tools.ietf.org/html/rfc2617#section-1.2）
//
// 学习备注:这种认证方式现代几乎已经没用,访问的时候浏览器会弹出一个简单的认证框,类似访问ftp,会提示输入账号密码.
# <翻译结束>


<原文开始>
// Search user in the slice of allowed credentials
<原文结束>

# <翻译开始>
// 在允许的凭据切片中搜索用户
# <翻译结束>


<原文开始>
// Credentials doesn't match, we return 401 and abort handlers chain.
<原文结束>

# <翻译开始>
// 凭证不匹配，我们返回 401 状态码并中断处理程序链。
# <翻译结束>


<原文开始>
		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
<原文结束>

# <翻译开始>
// 已找到用户凭据，将用户的ID设置为当前上下文中的AuthUserKey键，稍后可以通过
// c.MustGet(gin.AuthUserKey)读取用户的ID。
# <翻译结束>


<原文开始>
// BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
// the key is the user name and the value is the password.
<原文结束>

# <翻译开始>
// BasicAuth 返回一个基础HTTP授权中间件。它接受一个map[string]string作为参数，
// 其中键是用户名，值是密码。
//
// 学习备注:这种认证方式现代几乎已经没用,访问的时候浏览器会弹出一个简单的认证框,类似访问ftp,会提示输入账号密码.
# <翻译结束>

