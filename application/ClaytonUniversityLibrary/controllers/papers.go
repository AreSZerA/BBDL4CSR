package controllers

import (
	"ClaytonUniversityLibrary/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"os"
)

func init() {
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
	title := c.GetString("t")
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
		c.Data["username"] = user.(models.User).Name
	}
	c.Data["papers"] = papers
}

func (c *PapersController) Post() {
	sess := c.GetSession("user")
	if sess == nil {
		c.Ctx.Output.Status = http.StatusUnauthorized
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/json")
	user := sess.(models.User)
	title := c.GetString("title")
	abstract := c.GetString("abstract")
	authors := c.GetString("authors")
	keywords := c.GetString("keywords")
	paper, err := models.UploadPaper(user.Email, title, abstract, authors, keywords)
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	f, _, err := c.GetFile("file")
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		return
	}
	defer f.Close()
	err = c.SaveToFile("file", "static/papers/"+paper.ID+".pdf")
	if err != nil {
		c.Ctx.Output.Status = http.StatusInternalServerError
		return
	}
	c.Ctx.WriteString(NewResp(true, ""))
}
