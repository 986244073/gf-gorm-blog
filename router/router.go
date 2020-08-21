package router

import (
	"gf-app/app/api/hello"
	"gf-app/app/api/response"
	"gf-app/app/api/v1/blog"
	"gf-app/app/api/v1/user"
	"gf-app/app/service/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		//返回固定的友好信息
		r.Response.ClearBuffer()
		response.Failed(r, 500, "服务器居然开小差了，请稍后再试吧！")
	}
}
func init() {
	s := g.Server()
	// 某些浏览器直接请求favicon.ico文件，特别是产生404时
	s.Use(MiddlewareErrorHandler)


	s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
	// 分组路由
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//ctlChat := new(chat.Controller)
		ctlBlog := new(blog.Controller)
		ctlUser := new(user.Controller)
		group.ALL("/user", ctlUser)
		group.ALL("/blog", ctlBlog)
		group.Middleware(middleware.CORS)
		group.ALL("/", hello.Hello)
		//group.Group("/", func(group *ghttp.RouterGroup) {
		//	group.Middleware(middleware.Auth)
		//	group.ALL("/user/profile", ctlUser, "Profile")
		//})
	})
}
