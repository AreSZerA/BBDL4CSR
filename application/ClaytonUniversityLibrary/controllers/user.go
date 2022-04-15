package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
)

type UserLoginController struct {
	beego.Controller
}

func (c *UserLoginController) Prepare() {
	c.Ctx.Output.Header("content-type", "application/json")
}

func (c *UserLoginController) Post() {
	email := c.GetString("login-email")
	passwd := c.GetString("login-password")
	sess := c.GetSession("user")
	if sess == nil {
		ok, err := models.ValidateUser(email, passwd)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		if !ok {
			c.Ctx.WriteString(NewResp(false, ""))
			return
		}
		err = c.SetSession("user", email)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		c.Ctx.WriteString(NewResp(true, ""))
		return
	}
	c.Ctx.WriteString(NewResp(true, "you have already logged in"))
	return
}

type UserLogoutController struct {
	beego.Controller
}

func (c *UserLogoutController) Prepare() {
	c.Ctx.Output.Header("content-type", "application/json")
}

func (c *UserLogoutController) Post() {
	sess := c.GetSession("user")
	if sess == nil {
		c.Ctx.WriteString(NewResp(false, "you have not login"))
		return
	}
	err := c.DelSession("user")
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	c.Ctx.WriteString(NewResp(true, ""))
	return
}

type UserRegisterController struct {
	beego.Controller
}

func (c *UserRegisterController) Prepare() {
	c.Ctx.Output.Header("content-type", "application/json")
}

func (c *UserRegisterController) Post() {
	email := c.GetString("register-email")
	username := c.GetString("register-username")
	passwd := c.GetString("register-password")
	err := models.RegisterUser(email, username, passwd)
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	c.Ctx.WriteString(NewResp(true, ""))
	return
}
