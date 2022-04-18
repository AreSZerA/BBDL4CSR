package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Prepare() {
	user := c.GetSession("user")
	c.Data["isLogin"] = user != nil
	if user != nil {
		c.Data["isReviewer"] = user.(models.User).IsReviewer
		c.Data["isAdmin"] = user.(models.User).IsAdmin
		c.Data["username"] = user.(models.User).Name
	}
}

func (c *ErrorController) Error401() {
	c.Layout = "layout.html"
	c.TplName = "error.html"
	c.Data["statusCode"] = 401
	c.Data["info"] = "Unauthorized"
}

func (c *ErrorController) Error403() {
	c.Layout = "layout.html"
	c.TplName = "error.html"
	c.Data["statusCode"] = 403
	c.Data["info"] = "Forbidden"
}

func (c *ErrorController) Error404() {
	c.Layout = "layout.html"
	c.TplName = "error.html"
	c.Data["statusCode"] = 404
	c.Data["info"] = "Not Found"
}

func (c *ErrorController) Error500() {
	c.Layout = "layout.html"
	c.TplName = "error.html"
	c.Data["statusCode"] = 500
	c.Data["info"] = "Internal Server Error"
}

func (c *ErrorController) Error503() {
	c.Layout = "layout.html"
	c.TplName = "error.html"
	c.Data["statusCode"] = 503
	c.Data["info"] = "Service Unavailable"
}
