package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体，此处用作接收参数的验证
type Login struct {
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// 数据解析常用
func DataRouter(e *gin.Engine) {

	// 使用路由组
	r := e.Group("/data")

	// Json 数据解析和绑定
	r.POST("/loginJson", loginJson)

	// 表单数据解析和绑定
	r.POST("/loginForm", loginForm)

	// URI数据解析和绑定
	r.GET("/:user/:password", user)

}

// URI数据解析和绑定
func user(c *gin.Context) {
	// 声明接收的变量
	var login Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// 表单数据解析和绑定
func loginForm(c *gin.Context) {
	// 声明接收的变量
	var form Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

//Json 数据解析和绑定
//curl 127.0.0.1:8080/data/loginJson -H 'content-type:application/json' -d "{\"user\":\"root\",\"password\":\"admin\"}" -X POST
func loginJson(c *gin.Context) {
	// 声明接收的变量
	var json Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if json.User != "root" || json.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
