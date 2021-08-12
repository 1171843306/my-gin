package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("something-very-secret"))

// 会话控制常用
func ConversationRouter(e *gin.Engine) {

	// cookie
	{
		// 使用路由组
		r := e.Group("/conversation")

		// 服务端要给客户端cookie
		r.GET("cookie", cookie)
	}

	// session (此方法可直接替换main,为了归一特意放在这里)
	{
		// curl http://localhost:8080/save
		http.HandleFunc("/save", SaveSession)

		// curl http://localhost:8080/get
		http.HandleFunc("/get", GetSession)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("HTTP server failed,err:", err)
			return
		}
	}

}

// cookie 服务端发送cookie给客户端，客户端请求时携带cookie
func cookie(c *gin.Context) {
	// 获取客户端是否携带cookie
	cookie, err := c.Cookie("my_cookie")
	if err != nil {
		cookie = "没有设置"
		// 给客户端设置cookie
		// maxAge int, 单位为秒
		// path,cookie所在目录
		// domain string,域名
		// secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		// 谷歌浏览器，F12=>Application=>Cookies可看到
		c.SetCookie("my_cookie", "yzx-fjl.cn", 60, "/",
			"localhost", false, true)
	}
	fmt.Printf("cookie的值是： %s\n", cookie)
}

// 保存session
func SaveSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.

	//　获取一个session对象，my-session-name是session的名字
	session, err := store.Get(r, "my-session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("设置session")

	// 在session中存储值
	session.Values["name"] = "我是session"
	session.Values["sex"] = "girl"

	//// 将session的最大存储时间设置为小于零的数即为删除
	//session.Options.MaxAge = -1

	// 保存更改
	session.Save(r, w)
}

// 获取session
func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "my-session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := session.Values["name"]
	fmt.Println(name)
}
