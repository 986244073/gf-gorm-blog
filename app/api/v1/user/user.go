package user

import (
	"gf-app/app/api/response"
	"gf-app/app/model"
	user "gf-app/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) Add(r *ghttp.Request) {
	var data model.User
	err := r.Parse(&data)
	g.Dump(data)
	if err != nil {
		panic("解析错误")
	}
	err = user.CheckUser(data.Username)
	if err != nil {
		panic("用户已存在")
	}
	response.Success(r, 200, "添加成功")
}
