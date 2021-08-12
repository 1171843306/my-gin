package routers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


// 结构体验证 Person
type Person struct {
	// required 必须, gt 大于
	Name string    `form:"name" binding:"required"`
	Age  int       `form:"age" binding:"required,gt=10"`
	Day  time.Time `form:"day" time_format:"2006-01-02" time_utc:"1"`
}

/*
	自定义参数验证 需要 import "github.com/go-playground/validator/v10"
	对绑定到结构体上的参数，有时候无法直接使用现有方法，这时候自定义验证就来勒
   	比如我们要对 name 字段做校验，不等于apple
*/
type Goods struct {
	// 使用validate自定义的校验方法函数注册时候的名称
	Name string `form:"name" validate:"customFunc"`
	Type string `form:"type" binding:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// 参数验证常用 (使用gin框架的数据验证，可以减少代码手写判断。)
func VerificationRouter(e *gin.Engine) {

	// 使用路由组
	r := e.Group("/verification")

	// 结构体验证
	{
		// http://127.0.0.1:8080/verification/structure?age=27&name=%E9%9B%A8%E4%B8%AD%E7%AC%91&day=1995-03-17
		r.GET("/structure", structure)
	}

	// 自定义验证 http://127.0.0.1:8080/verification/custom?age=27&name=apple&type=fg
	{
		validate = validator.New()
		validate.RegisterValidation("customFunc", customFunc)

		r.GET("/custom", func(c *gin.Context) {
			var goods Goods

			if e := c.ShouldBind(&goods); e != nil {
				c.String(http.StatusOK, "Goods bind err:%v", e.Error())
			}

			err := validate.Struct(goods)
			if err != nil {
				fmt.Printf("Err(s):\n%+v\n", err)
				c.String(http.StatusOK, "Goods err:%v", err)
			} else {
				c.String(http.StatusOK, "%v", goods)
			}

		})

		// 更多验证文档看 https://pkg.go.dev/github.com/go-playground/validator/v10
	}

}

// 结构体验证
func structure(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		return
	}
	c.String(200, fmt.Sprintf("%#v", person))
}

// 1、自定义的校验方法
func customFunc(fl validator.FieldLevel) bool {

	return fl.Field().String() != "apple"
}
