package models

type PeerReview struct {
	Paper    string `json:"peer_review_paper"`
	Reviewer string `json:"peer_review_reviewer"`
	Status   string `json:"peer_review_status"`
	Comment  string `json:"peer_review_comment,omitempty"`
	Time     int64  `json:"peer_review_time,omitempty"`
}
