package lib

type PeerReview struct {
	Paper    string `json:"paper"`    // Paper composite key
	Reviewer string `json:"reviewer"` // User composite key
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
