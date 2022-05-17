// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions the peer review DAO and functions to update and query.

package models

import (
	"ClaytonUniversityLibrary/blockchain"
	"encoding/json"
	"time"
)

type jsonPeerReview struct {
	Paper      string `json:"peer_review_paper"`
	Reviewer   string `json:"peer_review_reviewer"`
	CreateTime int64  `json:"peer_review_create_time"`
	Status     string `json:"peer_review_status"`
	Comment    string `json:"peer_review_comment,omitempty"`
	Time       int64  `json:"peer_review_time,omitempty"`
}

func (pr jsonPeerReview) convert(i int) PeerReview {
	createTime := time.UnixMicro(pr.CreateTime / 1000).Format("2006-01-02 15:04:05")
	reviewTime := ""
	if pr.Time != 0 {
		reviewTime = time.UnixMicro(pr.Time / 1000).Format("2006-01-02 15:04:05")
	}
	return PeerReview{
		Count:      i,
		Paper:      pr.Paper,
		Reviewer:   pr.Reviewer,
		CreateTime: createTime,
		Status:     pr.Status,
		Comment:    pr.Comment,
		Time:       reviewTime,
	}
}

type PeerReview struct {
	Count      int
	Paper      string
	Reviewer   string
	CreateTime string
	Status     string
	Comment    string
	Time       string
}

func UpdatePeerReview(paperId string, email string, status string, comment string) error {
	_, err := blockchain.Execute(blockchain.FuncUpdatePeerReviewByPaperAndReviewer, []byte(paperId), []byte(email), []byte(status), []byte(comment))
	if err != nil {
		return err
	}
	return nil
}

func FindReviewingPeerReviewsByReviewer(email string) ([]PeerReview, error) {
	resp, err := blockchain.Query(blockchain.FuncRetrieveReviewingPeerReviewsByReviewerSortByCreateTime, []byte(email))
	if err != nil {
		return nil, err
	}
	var jpr []jsonPeerReview
	err = json.Unmarshal(resp.Payload, &jpr)
	if err != nil {
		return nil, err
	}
	var peerReviews []PeerReview
	for i, pr := range jpr {
		peerReviews = append(peerReviews, pr.convert(i+1))
	}
	return peerReviews, nil
}

func FindAcceptedPeerReviewsByReviewer(email string) ([]PeerReview, error) {
	resp, err := blockchain.Query(blockchain.FuncRetrieveAcceptedPeerReviewsByReviewerSortByCreateTime, []byte(email))
	if err != nil {
		return nil, err
	}
	var jpr []jsonPeerReview
	err = json.Unmarshal(resp.Payload, &jpr)
	if err != nil {
		return nil, err
	}
	var peerReviews []PeerReview
	for i, pr := range jpr {
		peerReviews = append(peerReviews, pr.convert(i+1))
	}
	return peerReviews, nil
}

func FindRejectedPeerReviewsByReviewer(email string) ([]PeerReview, error) {
	resp, err := blockchain.Query(blockchain.FuncRetrieveRejectedPeerReviewsByReviewerSortByCreateTime, []byte(email))
	if err != nil {
		return nil, err
	}
	var jpr []jsonPeerReview
	err = json.Unmarshal(resp.Payload, &jpr)
	if err != nil {
		return nil, err
	}
	var peerReviews []PeerReview
	for i, pr := range jpr {
		peerReviews = append(peerReviews, pr.convert(i+1))
	}
	return peerReviews, nil
}
