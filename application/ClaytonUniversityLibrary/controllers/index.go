// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the IndexController for route "/".

package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Layout = "layout.html"
	c.TplName = "index.html"
	user := c.GetSession("user")
	c.Data["isLogin"] = user != nil
	if user != nil {
		c.Data["isReviewer"] = user.(models.User).IsReviewer
		c.Data["isAdmin"] = user.(models.User).IsAdmin
		c.Data["username"] = user.(models.User).Name
	}
}
