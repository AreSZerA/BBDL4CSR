package lib

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

type Paper struct {
	ID          string    `json:"paper_id"`       // MD5(Uploader + UploadTime)
	Uploader    string    `json:"paper_uploader"` // User email
	UploadTime  int64     `json:"paper_upload_time"`
	Title       string    `json:"paper_title"`
	Abstract    string    `json:"paper_abstract"`
	Authors     []string  `json:"paper_authors"`
	Keywords    []string  `json:"paper_keywords,omitempty"`
	Reviewers   [3]string `json:"paper_reviewers"` // Reviewers' emails
	Status      string    `json:"paper_status"`
	PublishTime int64     `json:"paper_publish_time"`
}

func (p Paper) Type() string {
	return ObjectTypePaper
}

func (p Paper) Keys() []string {
	return []string{fmt.Sprintf("%x", md5.Sum([]byte(p.Uploader+strconv.FormatInt(p.UploadTime, 10))))}
}
