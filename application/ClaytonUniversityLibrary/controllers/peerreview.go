// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the PeerReviewController for route "/papers/peer_review".

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
	// check login status
	user := c.GetSession("user")
	if user == nil {
		c.Abort("401")
		return
	}
	// check identity
	if !user.(models.User).IsReviewer {
		c.Abort("403")
		return
	}
	// find three peer review information in three types of status
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
	// pass values and render
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
	// check login status
	user := c.GetSession("user")
	if user == nil {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		return
	}
	// check identity
	if !user.(models.User).IsReviewer {
		c.Ctx.Output.SetStatus(http.StatusForbidden)
		return
	}
	obj := struct {
		Id      string `json:"id"`
		Status  string `json:"status"`
		Comment string `json:"comment"`
	}{}
	// parse parameters from request body
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &obj)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotAcceptable)
		return
	}
	email := user.(models.User).Email
	// update peer review information
	err = models.UpdatePeerReview(obj.Id, email, obj.Status, obj.Comment)
	if err != nil {
		// response {"result":false,"error":"#{err}"}
		c.Ctx.WriteString(NewResp(false, err.Error()))
		return
	}
	// response {"result":true}
	c.Ctx.WriteString(NewResp(true, ""))
}
