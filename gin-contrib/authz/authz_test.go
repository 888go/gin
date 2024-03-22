// 版权所有 ? 2014 Manu Martinez-Almeida。保留所有权利。
// 本源代码的使用受 MIT 风格许可证约束，
// 该许可证可在 LICENSE 文件中找到。

package authz

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/casbin/casbin/v2"
	"github.com/888go/gin"
)

func testAuthzRequest(t *testing.T, router *gin类.Engine, user string, path string, method string, code int) {
	r, _ := http.NewRequestWithContext(context.Background(), method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, w.Code, code)
	}
}

func TestBasic(t *testing.T) {
	router := gin类.X创建()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.X中间件(NewAuthorizer(e))
	router.X绑定Any("/*anypath", func(c *gin类.Context) {
		c.X设置状态码(200)
	})

	testAuthzRequest(t, router, "alice", "/dataset1/resource1", "GET", 200)
	testAuthzRequest(t, router, "alice", "/dataset1/resource1", "POST", 200)
	testAuthzRequest(t, router, "alice", "/dataset1/resource2", "GET", 200)
	testAuthzRequest(t, router, "alice", "/dataset1/resource2", "POST", 403)
}

func TestPathWildcard(t *testing.T) {
	router := gin类.X创建()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.X中间件(NewAuthorizer(e))
	router.X绑定Any("/*anypath", func(c *gin类.Context) {
		c.X设置状态码(200)
	})

	testAuthzRequest(t, router, "bob", "/dataset2/resource1", "GET", 200)
	testAuthzRequest(t, router, "bob", "/dataset2/resource1", "POST", 200)
	testAuthzRequest(t, router, "bob", "/dataset2/resource1", "DELETE", 200)
	testAuthzRequest(t, router, "bob", "/dataset2/resource2", "GET", 200)
	testAuthzRequest(t, router, "bob", "/dataset2/resource2", "POST", 403)
	testAuthzRequest(t, router, "bob", "/dataset2/resource2", "DELETE", 403)

	testAuthzRequest(t, router, "bob", "/dataset2/folder1/item1", "GET", 403)
	testAuthzRequest(t, router, "bob", "/dataset2/folder1/item1", "POST", 200)
	testAuthzRequest(t, router, "bob", "/dataset2/folder1/item1", "DELETE", 403)
	testAuthzRequest(t, router, "bob", "/dataset2/folder1/item2", "GET", 403)
	testAuthzRequest(t, router, "bob", "/dataset2/folder1/item2", "POST", 200)
	testAuthzRequest(t, router, "bob", "/dataset2/folder1/item2", "DELETE", 403)
}

func TestRBAC(t *testing.T) {
	router := gin类.X创建()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.X中间件(NewAuthorizer(e))
	router.X绑定Any("/*anypath", func(c *gin类.Context) {
		c.X设置状态码(200)
	})

	// Cathy 可以通过所有方法访问 /dataset1/* 下的所有资源，因为它具有 dataset1_admin 角色。
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "GET", 200)
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "POST", 200)
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "DELETE", 200)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "GET", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "POST", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "DELETE", 403)

	// 删除用户cathy的所有角色，因此cathy现在无法访问任何资源。
	_, err := e.DeleteRolesForUser("cathy")
	if err != nil {
		t.Errorf("got error %v", err)
	}

	testAuthzRequest(t, router, "cathy", "/dataset1/item", "GET", 403)
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "POST", 403)
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "DELETE", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "GET", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "POST", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "DELETE", 403)
}
