package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

// 渲染常用
func RenderRouter(e *gin.Engine) {
	// 使用路由组
	r := e.Group("/render")

	// 各种数据格式的响应
	{
		// 1.json   curl 127.0.0.1:8080/render/someJSON
		r.GET("/someJSON", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "someJSON", "status": 200})
		})
		// 2. 结构体响应
		r.GET("/someStruct", someStruct)
		// 3.XML
		r.GET("/someXML", func(c *gin.Context) {
			c.XML(200, gin.H{"message": "abc"})
		})
		// 4.YAML响应
		r.GET("/someYAML", func(c *gin.Context) {
			c.YAML(200, gin.H{"name": "zhangsan"})
		})
		// 5.protobuf
		r.GET("/someProtoBuf", someProtoBuf)
	}

	// HTML模板渲染
	{
		// 加载模板文件
		e.LoadHTMLGlob("html/**/*")

		// 引入静态文件
		// r.Static("/assets", "./assets")

		// http://127.0.0.1:8080/render/index
		r.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "render/index.html", gin.H{"title": "这个锅我不背", "address": "http://yzx-fjl.cn"})
		})

		// html文件拆分例子
		r.GET("/view", func(c *gin.Context) {
			c.HTML(http.StatusOK, "render/view.html", gin.H{"title": "这个锅我不背", "address": "http://yzx-fjl.cn"})
		})
	}

	// 重定向
	{
		r.GET("/redirect", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "http://yzx-fjl.cn")
		})
	}

	// 同步异步
	{
		// 1.异步
		r.GET("/async", longAsync)
		// 2.同步
		r.GET("/sync", longSync)
	}

}

// protobuf
func someProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	// 定义数据
	label := "label"
	// 传protobuf格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}

// 结构体响应
func someStruct(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}
	msg.Name = "root"
	msg.Message = "message"
	msg.Number = 123
	c.JSON(200, msg)
}

// 异步  curl http://localhost:8080/render/async
func longAsync(c *gin.Context) {
	// 需要搞一个副本
	copyContext := c.Copy()
	// 异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}

// 同步 curl http://localhost:8080/render/sync
func longSync(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}