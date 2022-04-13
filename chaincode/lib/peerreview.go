package lib

type PeerReview struct {
	Paper    string `json:"paper"`    // Paper ID
	Reviewer string `json:"reviewer"` // User email
	Status   string `json:"status"`
	Comment  string `json:"comment,omitempty"`
	Time     int64  `json:"time,omitempty"`
}

func (p PeerReview) Type() string {
	return ObjectTypePeerReview
}

func (p PeerReview) Keys() []string {
	return []string{p.Paper, p.Reviewer}
}
