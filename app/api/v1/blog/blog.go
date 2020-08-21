package blog

import (
	"gf-app/app/api/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) GetBlog(r *ghttp.Request) {
	//blogs, err := sys_article.Model.FindAll()
	//if err != nil {
	//	panic(err)
	//}
	//g.Dump(blogs)
	response.Success(r, 200,1)
}
