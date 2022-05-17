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
	keyword := c.GetString("k")
	var mode string
	var countPapers int
	var papers []models.Paper
	var err error
	if title == "" && keyword == "" {
		mode = "global"
		countPapers, err = models.CountPublishedPapers()
		if err != nil {
			c.Abort("500")
			return
		}
		papers, err = models.FindPublishedPapers()
		if err != nil {
			c.Abort("500")
			return
		}
	} else if title != "" {
		mode = "title"
		countPapers, err = models.CountPublishedPapersByTitle(title)
		if err != nil {
			c.Abort("500")
			return
		}
		papers, err = models.FindPublishedPapersByTitle(title)
		if err != nil {
			c.Abort("500")
			return
		}
	} else {
		mode = "keyword"
		countPapers, err = models.CountPublishedPapersByKeyword(keyword)
		if err != nil {
			c.Abort("500")
			return
		}
		papers, err = models.FindPublishedPapersByKeyword(keyword)
		if err != nil {
			c.Abort("500")
			return
		}
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
	c.Data["mode"] = mode
	c.Data["title"] = title
	c.Data["keyword"] = keyword
	c.Data["countPapers"] = countPapers
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
