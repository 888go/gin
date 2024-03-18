// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package authz

import (
	"net/http"
	
	"github.com/casbin/casbin/v2"
	"github.com/888go/gin"
)

// NewAuthorizer返回授权器，使用Casbin强制器作为输入

// ff:
// e:

// ff:
// e:

// ff:
// e:

// ff:
// e:

// ff:
// e:

// ff:
// e:

// ff:
// e:
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c.Request) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer存储casbin处理程序
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName从请求中获取用户名
// 目前只支持HTTP基本认证

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission检查请求中的用户/方法/路径组合
// 返回true(授予权限)或false(禁止权限)

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:

// ff:
// r:
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	user := a.GetUserName(r)
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		panic(err)
	}

	return allowed
}

// RequirePermission返回403 Forbidden给客户端

// ff:
// c:

// ff:
// c:

// ff:
// c:

// ff:
// c:

// ff:
// c:

// ff:
// c:

// ff:
// c:
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}
