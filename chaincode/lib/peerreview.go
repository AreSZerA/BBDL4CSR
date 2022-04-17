package lib

type PeerReview struct {
	Paper      string `json:"peer_review_paper"`    // Paper ID
	Reviewer   string `json:"peer_review_reviewer"` // User email
	CreateTime int64  `json:"peer_review_create_time"`
	Status     string `json:"peer_review_status"`
	Comment    string `json:"peer_review_comment,omitempty"`
	Time       int64  `json:"peer_review_time,omitempty"`
}

func (p PeerReview) Type() string {
	return ObjectTypePeerReview
}

func (p PeerReview) Keys() []string {
	return []string{p.Paper, p.Reviewer}
}
