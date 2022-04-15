package lib

import "strconv"

type Paper struct {
	ID         string    `json:"paper_id"`       // MD5(Uploader + UploadTime)
	Uploader   string    `json:"paper_uploader"` // User email
	UploadTime int64     `json:"paper_upload_time"`
	Title      string    `json:"paper_title"`
	Abstract   string    `json:"paper_abstract"`
	Authors    []string  `json:"paper_authors"`
	Keywords   []string  `json:"paper_keywords,omitempty"`
	Reviewers  [3]string `json:"paper_reviewers"` // Reviewers' emails
	Status     string    `json:"paper_status"`
	ReviewTime int64     `json:"paper_review_time"`
}

func (p Paper) Type() string {
	return ObjectTypePaper
}

func (p Paper) Keys() []string {
	return []string{p.Uploader + strconv.FormatInt(p.UploadTime, 10)}
}
