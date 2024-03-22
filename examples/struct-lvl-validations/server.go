package main

import (
	"net/http"
	
	"github.com/888go/gin"
	"github.com/888go/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// User 包含用户信息。
type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

// UserStructLevelValidation 包含一些自定义的结构级别验证，这些验证在字段级别上并不总是适用。例如，此函数验证 FirstName 或 LastName 至少有一个存在；虽然也可以通过自定义字段验证来实现，但那样就需要在两个字段上都添加该验证逻辑，导致代码重复和额外开销，而这种方式只需验证一次。
// 
// 注意：你可能会问为什么不直接在 validator 之外进行这种验证。采用这种方式将验证过程直接融入到 validator 中，可以与验证标签结合使用，并且仍然保持统一的错误输出格式。
func UserStructLevelValidation(sl validator.StructLevel) {
	// 获取当前结构体的接口表示，并将其转换为 User 类型，赋值给 user 变量
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}

	// plus 可以做更多事情，即使标签不同于 "fnameorlname"
}

func main() {
	route := gin类.X创建默认对象()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}

	route.X绑定POST("/user", validateUser)
	route.X监听(":8085")
}

func validateUser(c *gin类.Context) {
	var u User
	if err := c.X取JSON参数到指针(&u); err == nil {
		c.X输出JSON(http.StatusOK, gin类.H{"message": "User validation successful."})
	} else {
		c.X输出JSON(http.StatusBadRequest, gin类.H{
			"message": "User validation failed!",
			"error":   err.Error(),
		})
	}
}
