package main

import (
	"net/http"
	"time"
	
	"github.com/888go/gin"
	"github.com/888go/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	route := gin类.X创建默认对象()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.X绑定GET("/bookable", getBookable)
	route.X监听(":8085")
}

func getBookable(c *gin类.Context) {
	var b Booking
	if err := c.X取参数到指针并按类型(&b, binding.Query); err == nil {
		c.X输出JSON(http.StatusOK, gin类.H{"message": "Booking dates are valid!"})
	} else {
		c.X输出JSON(http.StatusBadRequest, gin类.H{"error": err.Error()})
	}
}
