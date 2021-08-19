package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-gin/toolkit/config"
	"net/http"
	"strings"
)

// 基础
func BaseRouter(e *gin.Engine) {

	// 使用路由组
	r := e.Group("/base")

	// 获取API参数 http://127.0.0.1:8080/base/view/123/like
	r.GET("/view/:id/*action", api)

	// 获取url参数 http://127.0.0.1:8080/base/view?name=world
	r.GET("/view", url)

	// 获取表单参数 打开html/base/form.html
	r.POST("/form", form)

	// 上传单个文件 打开html/base/upload.html
	r.POST("/upload", upload)

	// 上传多个文件
	r.POST("/uploads", uploads)

	// 获取url参数 http://127.0.0.1:8080/base/view?name=world
	r.GET("/config", getConfig)

}

// 获取API参数
func api(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	// 截取/
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, id+" is "+action)
}

// 获取URL参数
func url(c *gin.Context) {
	name := c.DefaultQuery("name", "yzx-fjl.cn")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

// 获取表单参数
func form(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}

// 上传单个文件
func upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
		return
	}

	// 文件大小不能超过5M
	if file.Size > 5000000 {
		c.String(500, "上传图片超过5M")
		return
	}

	// 上传路径
	path := "./upload/" + file.Filename
	c.SaveUploadedFile(file, path)
	c.String(http.StatusOK, fmt.Sprintf("Filename:%s,Size:%d字节", file.Filename, file.Size))
}

// 上传多个文件
func uploads(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		return
	}

	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		path := "./upload/" + file.Filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}

// 获取配置
// 复制 configs/config.toml.default 为config.toml
func getConfig(c *gin.Context) {
	conf := config.Configure()
	conf.Set("redis.port", 6381)
	c.JSON(200, gin.H{
		"listen": conf.GetString("http.listen"),
		"app_name": conf.Get("app_name"),
		"ip":       conf.Get("mysql.ip"),
		"config":   conf.AllSettings(),
	})
}
