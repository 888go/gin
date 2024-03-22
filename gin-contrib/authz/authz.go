// 版权所有 ? 2014 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package authz

import (
	"net/http"
	
	"github.com/casbin/casbin/v2"
	"github.com/888go/gin"
)

// NewAuthorizer 返回一个鉴权器，使用 Casbin 执行器作为输入参数
func NewAuthorizer(e *casbin.Enforcer) gin类.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin类.Context) {
		if !a.CheckPermission(c.X请求) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer 基础授权器存储了 casbin 处理器
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName 从请求中获取用户名。
// 当前仅支持HTTP基本认证。
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission 检查请求中的用户/方法/路径组合。
// 返回 true（权限授予）或 false（权限禁止）
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

// RequirePermission 返回 403 Forbidden 给客户端
func (a *BasicAuthorizer) RequirePermission(c *gin类.Context) {
	c.X停止并带状态码(http.StatusForbidden)
}
