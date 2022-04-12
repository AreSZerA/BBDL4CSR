package lib

type PeerReview struct {
	ID       string `json:"id"`    // CompositeKey(Paper + Reviewer)
	Paper    string `json:"paper"` // Paper composite key
	Reviewer string `json:"reviewer"`
	Status   string `json:"status"`
	Comment  string `json:"comment,omitempty"`
	Time     int64  `json:"time,omitempty"`
}

func (p PeerReview) ObjectType() string {
	return ObjectTypePeerReview
}
