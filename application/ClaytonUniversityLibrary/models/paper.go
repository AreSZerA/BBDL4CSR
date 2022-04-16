package models

type Paper struct {
	ID         string    `json:"paper_id"`
	Uploader   string    `json:"paper_uploader"`
	UploadTime int64     `json:"paper_upload_time"`
	Title      string    `json:"paper_title"`
	Abstract   string    `json:"paper_abstract"`
	Authors    []string  `json:"paper_authors"`
	Keywords   []string  `json:"paper_keywords,omitempty"`
	Reviewers  [3]string `json:"paper_reviewers"`
	Status     string    `json:"paper_status"`
	ReviewTime int64     `json:"paper_review_time"`
}
