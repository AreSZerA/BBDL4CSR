package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	user := c.GetSession("user")
	if user == nil {
		c.Abort("401")
		return
	}
	if !user.(models.User).IsAdmin {
		c.Abort("403")
		return
	}
	c.Data["isLogin"] = true
	c.Layout = "layout.html"
	c.TplName = "admin.html"
	c.Data["isReviewer"] = user.(models.User).IsReviewer
	c.Data["isAdmin"] = user.(models.User).IsAdmin
}
