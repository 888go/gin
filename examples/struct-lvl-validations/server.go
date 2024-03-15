package main

import (
	"net/http"
	
	"github.com/888go/gin"
	"github.com/888go/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// User包含用户信息
type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

// UserStructLevelValidation包含自定义结构级验证，这些验证在字段验证级别上并不总是有意义的
// 例如，这个函数验证FirstName或LastName是否存在;本可以使用自定义字段验证来完成此操作，但随后必须将其添加到复制逻辑+开销的两个字段中，这样只验证一次
// 注意:你可能会问为什么不在验证器之外做这个
// 这样做可以直接与验证器挂钩，并且可以与验证标记结合使用，并且仍然具有常见的错误输出格式
func UserStructLevelValidation(sl validator.StructLevel) {
// user:= structLevel.CurrentStruct.Interface().(user)
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}

// Plus可以添加更多，即使标签与“fnameorlname”不同
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}

	route.POST("/user", validateUser)
	route.Run(":8085")
}

func validateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed!",
			"error":   err.Error(),
		})
	}
}
