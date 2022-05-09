// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the PapersController for route "/papers".

package controllers

import (
	"ClaytonUniversityLibrary/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"os"
)

func init() {
	// create a directory to place uploaded papers
	err := os.MkdirAll("static/papers/", os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create dir:", err.Error())
		fmt.Println("Please create manually")
	}
}

type PapersController struct {
	beego.Controller
}

func (c *PapersController) Get() {
	// parse parameter
	title := c.GetString("t")
	// find published papers according to the parameter
	papers, err := models.FindPublishedPapersByTitle(title)
	if err != nil {
		c.Abort("500")
		return
	}
	c.Layout = "layout.html"
	c.TplName = "papers.html"
	user := c.GetSession("user")
	c.Data["isLogin"] = user != nil
	if user != nil {
		c.Data["isReviewer"] = user.(models.User).IsReviewer
		c.Data["isAdmin"] = user.(models.User).IsAdmin
		c.Data["username"] = user.(models.User).Name
	}
	c.Data["papers"] = papers
}

func (c *PapersController) Post() {
	// check login status
	sess := c.GetSession("user")
	if sess == nil {
		c.Ctx.Output.Status = http.StatusUnauthorized
		return
	}
	// parse parameters
	c.Ctx.Output.Header("Content-Type", "application/json")
	user := sess.(models.User)
	title := c.GetString("title")
	abstract := c.GetString("abstract")
	authors := c.GetString("authors")
	keywords := c.GetString("keywords")
	// read file from
	f, _, err := c.GetFile("file")
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		return
	}
	defer f.Close()
	// put paper to the ledger
	paper, err := models.UploadPaper(user.Email, title, abstract, authors, keywords)
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	// if the updated successfully, save the file from memory
	err = c.SaveToFile("file", "static/papers/"+paper.ID+".pdf")
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		return
	}
	// response {"result":true}
	c.Ctx.WriteString(NewResp(true, ""))
}
