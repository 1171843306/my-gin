package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 中间件常用
func MiddlewareRouter(e *gin.Engine) {

	//// 注册全局中间件 (测试其它时注释)
	//e.Use(global())
	//
	//// 注册中间件next() (测试其它时注释)
	//e.Use(next())

	// 使用路由组
	r := e.Group("/middleware")

	// 注册全局中间件
	{
		// curl 127.0.0.1:8080/middleware/global
		r.GET("/global", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("user_id")
			fmt.Println("用户ID:", req)
			// 页面接收
			c.JSON(200, gin.H{"user_id": req})
		})
	}

	//局部中间键使用(个别路由可用)
	{
		r.GET("/local", next(), func(c *gin.Context) {
			// 取值
			req, _ := c.Get("user_id")
			fmt.Println("user_id:", req)
			// 页面接收
			c.JSON(200, gin.H{"user_id": req})
		})
	}

}

// 全局中间件
func global() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("middleware begin")
		// 设置变量到Context的key中，可以通过Get()取，如现实项目通常会设置用户id
		c.Set("user_id", 9527)
		status := c.Writer.Status()
		fmt.Println("middleware end", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// Next()方法  curl 127.0.0.1:8080/middleware/global
func next() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("middleware-next begin")
		// 设置变量到Context的key中，可以通过Get()取，如现实项目通常会设置用户id
		c.Set("user_id", 9527)
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("middleware-next end", status)
		t2 := time.Since(t)
		fmt.Println("执行时间:", t2)
	}
}
