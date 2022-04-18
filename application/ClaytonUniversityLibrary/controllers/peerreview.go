package controllers

import (
	"ClaytonUniversityLibrary/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type PeerReviewController struct {
	beego.Controller
}

func (c *PeerReviewController) Get() {
	user := c.GetSession("user")
	if user == nil {
		c.Abort("401")
		return
	}
	if !user.(models.User).IsReviewer {
		c.Abort("403")
		return
	}
	reviewingPeerReviews, err := models.FindReviewingPeerReviewsByReviewer(user.(models.User).Email)
	if err != nil {
		c.Abort("500")
		return
	}
	acceptedPeerReviews, err := models.FindAcceptedPeerReviewsByReviewer(user.(models.User).Email)
	if err != nil {
		c.Abort("500")
		return
	}
	rejectedPeerReviews, err := models.FindRejectedPeerReviewsByReviewer(user.(models.User).Email)
	if err != nil {
		c.Abort("500")
		return
	}
	c.Data["isLogin"] = true
	c.Layout = "layout.html"
	c.TplName = "peerreview.html"
	c.Data["isReviewer"] = user.(models.User).IsReviewer
	c.Data["isAdmin"] = user.(models.User).IsAdmin
	c.Data["reviewingPeerReviews"] = reviewingPeerReviews
	c.Data["acceptedPeerReviews"] = acceptedPeerReviews
	c.Data["rejectedPeerReviews"] = rejectedPeerReviews
}

func (c *PeerReviewController) Post() {
	user := c.GetSession("user")
	if user == nil {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		return
	}
	if !user.(models.User).IsReviewer {
		c.Ctx.Output.SetStatus(http.StatusForbidden)
		return
	}
	obj := struct {
		Id      string `json:"id"`
		Status  string `json:"status"`
		Comment string `json:"comment"`
	}{}
	email := user.(models.User).Email
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotAcceptable)
		return
	}
	err = models.UpdatePeerReview(obj.Id, email, obj.Status, obj.Comment)
	if err != nil {
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	c.Ctx.WriteString(NewResp(true, ""))
}
