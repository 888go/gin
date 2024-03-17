// Manu Martinez-Almeida版权所有
// 版权所有
// 此源代码的使用受MIT风格许可的约束，该许可可以在license文件中找到

package authz

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/casbin/casbin/v2"
	"github.com/888go/gin"
)

func testAuthzRequest(t *testing.T, router *gin.Engine, user string, path string, method string, code int) {
	r, _ := http.NewRequestWithContext(context.Background(), method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, w.Code, code)
	}
}


// ff:
// t:

// ff:
// t:
func TestBasic(t *testing.T) {
	router := gin.New()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.Use(NewAuthorizer(e))
	router.Any("/*anypath", func(c *gin.Context) {
		c.Status(200)
	})

	testAuthzRequest(t, router, "alice", "/dataset1/resource1", "GET", 200)
	testAuthzRequest(t, router, "alice", "/dataset1/resource1", "POST", 200)
	testAuthzRequest(t, router, "alice", "/dataset1/resource2", "GET", 200)
	testAuthzRequest(t, router, "alice", "/dataset1/resource2", "POST", 403)
}


// ff:
// t:

// ff:
// t:
func TestPathWildcard(t *testing.T) {
	router := gin.New()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.Use(NewAuthorizer(e))
	router.Any("/*anypath", func(c *gin.Context) {
		c.Status(200)
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


// ff:
// t:

// ff:
// t:
func TestRBAC(t *testing.T) {
	router := gin.New()
	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	router.Use(NewAuthorizer(e))
	router.Any("/*anypath", func(c *gin.Context) {
		c.Status(200)
	})

// Cathy可以通过所有方法访问所有/dataset1/*资源，因为它具有dataset1_admin角色
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "GET", 200)
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "POST", 200)
	testAuthzRequest(t, router, "cathy", "/dataset1/item", "DELETE", 200)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "GET", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "POST", 403)
	testAuthzRequest(t, router, "cathy", "/dataset2/item", "DELETE", 403)

// 删除用户cathy上的所有角色，因此cathy现在不能访问任何资源
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
