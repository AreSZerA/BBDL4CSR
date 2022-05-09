// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the structure of paper.

package lib

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// Paper implements BlockchainObjectInterface for storing in the ledger.
type Paper struct {
	// Unique ID for identifying papers, calculated by MD5(Uploader + UploadTime).
	ID string `json:"paper_id"`
	// Email address of the uploader, related to User.Email.
	Uploader string `json:"paper_uploader"`
	// Timestamp of the upload time, initialized when the object is instantiated.
	UploadTime int64 `json:"paper_upload_time"`
	// Title of the paper.
	Title string `json:"paper_title"`
	// Abstract of the paper.
	Abstract string `json:"paper_abstract"`
	// Authors of the paper.
	Authors []string `json:"paper_authors"`
	// Keywords of the paper.
	Keywords []string `json:"paper_keywords,omitempty"`
	// Emails of the three reviewers, related to User.Email.
	Reviewers [3]string `json:"paper_reviewers"`
	// Status of the paper, could be StatusReviewing, StatusAccepted, or StatusRejected.
	Status string `json:"paper_status"`
	// Timestamp of the publishing time of the paper, will be updated by peer review.
	PublishTime int64 `json:"paper_publish_time"`
}

// Type returns ObjectTypePaper.
func (p Paper) Type() string {
	return ObjectTypePaper
}

// Attributes returns a slice constructed with MD5(Uploader + UploadTime), which is also the paper ID.
func (p Paper) Attributes() []string {
	return []string{fmt.Sprintf("%x", md5.Sum([]byte(p.Uploader+strconv.FormatInt(p.UploadTime, 10))))}
}
