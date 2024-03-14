
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
// AuthUserKey is the cookie name for user credential in basic auth.
<原文结束>

# <翻译开始>
// AuthUserKey是基本验证中用户凭据的cookie名称
# <翻译结束>


<原文开始>
// Accounts defines a key/value for user/pass list of authorized logins.
<原文结束>

# <翻译开始>
// Accounts为授权登录的用户/通行证列表定义一个键/值
# <翻译结束>


<原文开始>
// BasicAuthForRealm returns a Basic HTTP Authorization middleware. It takes as arguments a map[string]string where
// the key is the user name and the value is the password, as well as the name of the Realm.
// If the realm is empty, "Authorization Required" will be used by default.
// (see http://tools.ietf.org/html/rfc2617#section-1.2)
<原文结束>

# <翻译开始>
// BasicAuthForRealm返回一个基本HTTP授权中间件
// 它接受一个map[string]字符串作为参数，其中键是用户名，值是密码，以及Realm的名称
// 如果领域为空，则“授权需要”;将默认使用
// (见http://tools.ietf.org/html/rfc2617 - 1.2节)
# <翻译结束>


<原文开始>
		// Search user in the slice of allowed credentials
<原文结束>

# <翻译开始>
// 在允许的凭据片中搜索用户
# <翻译结束>


<原文开始>
			// Credentials doesn't match, we return 401 and abort handlers chain.
<原文结束>

# <翻译开始>
// 凭证不匹配，我们返回401并中止处理程序链
# <翻译结束>


<原文开始>
		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
<原文结束>

# <翻译开始>
// 找到用户凭据，在此上下文中将用户id设置为密钥AuthUserKey，稍后可以使用c.MustGet(gin.AuthUserKey)读取用户id
# <翻译结束>


<原文开始>
// BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
// the key is the user name and the value is the password.
<原文结束>

# <翻译开始>
// BasicAuth返回一个基本HTTP授权中间件
// 它接受一个map[string]字符串作为参数，其中键是用户名，值是密码
# <翻译结束>

