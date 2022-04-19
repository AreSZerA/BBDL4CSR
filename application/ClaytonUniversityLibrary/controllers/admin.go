package controllers

import (
	"ClaytonUniversityLibrary/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
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
	users, err := models.FindUsers()
	if err != nil {
		c.Abort("500")
		return
	}
	c.Data["isLogin"] = true
	c.Layout = "layout.html"
	c.TplName = "admin.html"
	c.Data["isReviewer"] = user.(models.User).IsReviewer
	c.Data["isAdmin"] = user.(models.User).IsAdmin
	c.Data["users"] = users
}

func (c *AdminController) Post() {
	user := c.GetSession("user")
	if user == nil {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		return
	}
	if !user.(models.User).IsAdmin {
		c.Ctx.Output.SetStatus(http.StatusForbidden)
		return
	}
	obj := struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotAcceptable)
		return
	}
	if obj.Password != user.(models.User).Password {
		c.Ctx.WriteString(NewResp(false, ""))
		return
	}
	err = models.UpdateUserIsReviewer(obj.Email)
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	c.Ctx.WriteString(NewResp(true, ""))
}
