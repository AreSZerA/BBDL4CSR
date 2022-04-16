package controllers

import (
	"ClaytonUniversityLibrary/models"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type UserLoginController struct {
	beego.Controller
}

func (c *UserLoginController) Prepare() {
	c.Ctx.Output.Header("content-type", "application/json")
}

func (c *UserLoginController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")
	user := c.GetSession("user")
	if user == nil {
		user, err := models.FindUserByEmail(email)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		if user.Password != password {
			c.Ctx.WriteString(NewResp(false, ""))
			return
		}
		err = c.SetSession("user", user)
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
	user := c.GetSession("user")
	if user == nil {
		c.Ctx.WriteString(NewResp(false, "you have not login"))
		return
	}
	err := c.DestroySession()
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
	email := c.GetString("email")
	username := c.GetString("username")
	password := c.GetString("password")
	err := models.RegisterUser(email, username, password)
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	c.Ctx.WriteString(NewResp(true, ""))
	return
}

type UserUpdateController struct {
	beego.Controller
}

func (c *UserUpdateController) Prepare() {
	c.Ctx.Output.Header("content-type", "application/json")
}

func (c *UserUpdateController) Post() {
	sess := c.GetSession("user")
	if sess == nil {
		c.Ctx.Output.Status = http.StatusUnauthorized
		return
	}
	field := c.GetString("field")
	user := sess.(models.User)
	switch field {
	case "username":
		username := c.GetString("username")
		password := c.GetString("password")
		if password != sess.(models.User).Password {
			c.Ctx.WriteString(NewResp(false, "wrong password"))
			return
		}
		user.Name = username
		err := models.UpdateUsername(user.Email, username)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		_ = c.SetSession("user", user)
		c.Ctx.WriteString(NewResp(true, ""))
		return
	case "password":
		oldPassword := c.GetString("oldPassword")
		newPassword := c.GetString("newPassword")
		if oldPassword != sess.(models.User).Password {
			c.Ctx.WriteString(NewResp(false, "wrong password"))
			return
		}
		user.Password = newPassword
		err := models.UpdatePassword(user.Email, newPassword)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		_ = c.SetSession("user", user)
		c.Ctx.WriteString(NewResp(true, ""))
		return
	default:
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
}
