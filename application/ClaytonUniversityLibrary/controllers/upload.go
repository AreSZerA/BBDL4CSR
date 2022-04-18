package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UploadPaperController struct {
	beego.Controller
}

func (c *UploadPaperController) Get() {
	user := c.GetSession("user")
	if user == nil {
		c.Abort("401")
		return
	}
	c.Data["isLogin"] = true
	c.Layout = "layout.html"
	c.TplName = "upload.html"
	c.Data["isReviewer"] = user.(models.User).IsReviewer
	c.Data["username"] = user.(models.User).Name
}
