// Copyright 2022 AreSZerA. All rights reserved.
// This file defines controllers related to users for routes "/users/login", "/users/logout", "/users/register", and "/users/update".

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
	// parse parameters from request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
	// check login status
	user := c.GetSession("user")
	// if not login
	if user == nil {
		// retrieve user by email
		user, err := models.FindUserByEmail(obj.Email)
		if err != nil {
			// response {"result":false,"error":"#{err}"}
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		// check password
		if user.Password != obj.Password {
			// response {"result":false,"error":"#{err}"}
			c.Ctx.WriteString(NewResp(false, ""))
			return
		}
		// store user instance in session
		err = c.SetSession("user", user)
		if err != nil {
			// response {"result":false,"error":"#{err}"}
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		// response {"result":true}
		c.Ctx.WriteString(NewResp(true, ""))
		return
	}
	// otherwise, response {"result":true,"error":"you have already logged in"}
	c.Ctx.WriteString(NewResp(true, "you have already logged in"))
}

type UserLogoutController struct {
	beego.Controller
}

func (c *UserLogoutController) Prepare() {
	c.Ctx.Output.Header("content-type", "application/json")
}

func (c *UserLogoutController) Post() {
	// check login status
	user := c.GetSession("user")
	if user == nil {
		// response {"result":false,"error":"you have not login"}
		c.Ctx.WriteString(NewResp(false, "you have not login"))
		return
	}
	// destroy session
	err := c.DestroySession()
	if err != nil {
		// {"result":false,"error":"#{err}"}
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	// {"result"true}
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
	// parse parameters from request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
	// register user
	err = models.RegisterUser(obj.Email, obj.Username, obj.Password)
	if err != nil {
		// response {"result":false,"error":"#{err}"}
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	// response {"result":true}
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
	// parse parameters from request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
	user := sess.(models.User)
	// do different actions according to the obj.Field
	switch obj.Field {
	case "username":
		// check password
		if obj.Password != sess.(models.User).Password {
			// response {"result":false,"error":"wrong password"}
			c.Ctx.WriteString(NewResp(false, "wrong password"))
			return
		}
		// reassign and update
		user.Name = obj.Value
		err := models.UpdateUsername(user.Email, user.Name)
		if err != nil {
			// response {"result":false,"error":"#{err}"}
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		// update session
		_ = c.SetSession("user", user)
		// response {"result":true}
		c.Ctx.WriteString(NewResp(true, ""))
		return
	case "password":
		if obj.Password != sess.(models.User).Password {
			// response {"result":false,"error":"wrong password"}
			c.Ctx.WriteString(NewResp(false, "wrong password"))
			return
		}
		// reassign and update
		user.Password = obj.Value
		err := models.UpdatePassword(user.Email, user.Password)
		if err != nil {
			// response {"result":false,"error":"#{err}"}
			c.Ctx.WriteString(NewResp(false, err.Error()))
			return
		}
		// update session
		_ = c.SetSession("user", user)
		// response {"result":true}
		c.Ctx.WriteString(NewResp(true, ""))
		return
	default:
		c.Ctx.Output.Status = http.StatusNotAcceptable
		return
	}
}
