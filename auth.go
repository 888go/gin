// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package gin

import (
	"crypto/subtle"
	"encoding/base64"
	"net/http"
	"strconv"
	
	"github.com/888go/gin/internal/bytesconv"
)

// AuthUserKey是基本验证中用户凭据的cookie名称
const AuthUserKey = "user"

// Accounts为授权登录的用户/通行证列表定义一个键/值
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

// BasicAuthForRealm返回一个基本HTTP授权中间件
// 它接受一个map[string]字符串作为参数，其中键是用户名，值是密码，以及Realm的名称
// 如果领域为空，则“授权需要”;将默认使用
// (见http://tools.ietf.org/html/rfc2617 - 1.2节)

// ff:
// realm:
// accounts:
func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := processAccounts(accounts)
	return func(c *Context) {
// 在允许的凭据片中搜索用户
		user, found := pairs.searchCredential(c.requestHeader("Authorization"))
		if !found {
// 凭证不匹配，我们返回401并中止处理程序链
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

// 找到用户凭据，在此上下文中将用户id设置为密钥AuthUserKey，稍后可以使用c.MustGet(gin.AuthUserKey)读取用户id
		c.Set(AuthUserKey, user)
	}
}

// BasicAuth返回一个基本HTTP授权中间件
// 它接受一个map[string]字符串作为参数，其中键是用户名，值是密码

// ff:
// accounts:
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
