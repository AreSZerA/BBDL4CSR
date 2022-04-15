package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Layout = "frame.gohtml"
	c.TplName = "index.gohtml"
	user := c.GetSession("user")
	c.Data["isLogin"] = user != nil
	if user != nil {
		c.Data["isReviewer"] = user.(models.User).IsReviewer
		c.Data["username"] = user.(models.User).Name
	}
}
