package controllers

import beego "github.com/beego/beego/v2/server/web"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Layout = "frame.gohtml"
	c.TplName = "index.gohtml"
}
