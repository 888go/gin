package main

import (
	"net/http"
	
	"github.com/888go/gin"
	"github.com/888go/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// User contains user information.
type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For example, this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't not just do this outside of validator. Doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
func UserStructLevelValidation(sl validator.StructLevel) {
	// user := structLevel.CurrentStruct.Interface().(User)
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}

	// plus can to more, even with different tag than "fnameorlname"
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
