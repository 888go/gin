// 版权所有 2014 Manu Martinez-Almeida。保留所有权利。
// 使用本源代码受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package gin

import (
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strconv"
	
	"github.com/888go/gin/internal/bytesconv"
)

// AuthUserKey 是基本认证中用于存储用户凭证的cookie名称。
const AuthUserKey = "user"

// Accounts 定义了一个用于存储授权登录用户/密码键值对的列表。
type Accounts map[string]string

type authPair struct {
	value string
	user  string
}

type authPairs []authPair

func (a authPairs) searchCredential(authValue string) (string, bool) {
	if authValue == "" {
		return "", false
	}
	for _, pair := range a {
		if subtle.ConstantTimeCompare(bytesconv.StringToBytes(pair.value), bytesconv.StringToBytes(authValue)) == 1 {
			return pair.user, true
		}
	}
	return "", false
}

// BasicAuthForRealm 返回一个基础HTTP身份验证中间件。它接受两个参数：一个map[string]string，其中键是用户名，值是密码；以及一个realm（领域）名称。
// 如果realm为空，则默认使用"Authorization Required"。
// （参见http://tools.ietf.org/html/rfc2617#section-1.2）

// ff:中间件函数_简单认证2
// realm:
// accounts:账号密码Map
func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := processAccounts(accounts)
	return func(c *Context) {
		// 在允许的凭据切片中搜索用户
		user, found := pairs.searchCredential(c.requestHeader("Authorization"))
		if !found {
			// 凭证不匹配，我们返回 401 状态码并中断处理程序链。
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 已找到用户凭据，将用户的ID设置为当前上下文中的AuthUserKey键，稍后可以通过
		// c.MustGet(gin.AuthUserKey)读取用户的ID。
		c.Set(AuthUserKey, user)
	}
}

// BasicAuth 返回一个基础HTTP授权中间件。它接受一个map[string]string作为参数，
// 其中键是用户名，值是密码。

// ff:中间件函数_简单认证
// accounts:账号密码Map
func BasicAuth(accounts Accounts) HandlerFunc {
	return BasicAuthForRealm(accounts, "")
}

func processAccounts(accounts Accounts) authPairs {
	length := len(accounts)
	assert1(length > 0, "Empty list of authorized credentials")
	pairs := make(authPairs, 0, length)
	for user, password := range accounts {
		assert1(user != "", "User can not be empty")
		value := authorizationHeader(user, password)
		pairs = append(pairs, authPair{
			value: value,
			user:  user,
		})
	}
	return pairs
}

func authorizationHeader(user, password string) string {
	base := user + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString(bytesconv.StringToBytes(base))
}
