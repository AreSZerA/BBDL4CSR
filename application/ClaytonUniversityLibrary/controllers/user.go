package controllers

import (
	"ClaytonUniversityLibrary/models"
	"encoding/json"
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
	obj := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
	user := c.GetSession("user")
	if user == nil {
		user, err := models.FindUserByEmail(obj.Email)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		if user.Password != obj.Password {
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
	obj := struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
	err = models.RegisterUser(obj.Email, obj.Username, obj.Password)
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
	obj := struct {
		Field    string `json:"field"`
		Value    string `json:"value"`
		Password string `json:"password"`
	}{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
	user := sess.(models.User)
	switch obj.Field {
	case "username":
		if obj.Password != sess.(models.User).Password {
			c.Ctx.WriteString(NewResp(false, "wrong password"))
			return
		}
		user.Name = obj.Value
		err := models.UpdateUsername(user.Email, user.Name)
		if err != nil {
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		_ = c.SetSession("user", user)
		c.Ctx.WriteString(NewResp(true, ""))
		return
	case "password":
		if obj.Password != sess.(models.User).Password {
			c.Ctx.WriteString(NewResp(false, "wrong password"))
			return
		}
		user.Password = obj.Value
		err := models.UpdatePassword(user.Email, user.Password)
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
