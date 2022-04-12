package lib

type Paper struct {
	ID          string    `json:"id"`       // CompositeKey(Uploader + UploadTime)
	Uploader    string    `json:"uploader"` // User composite key
	UploadTime  int64     `json:"upload_time"`
	Title       string    `json:"title"`
	Abstract    string    `json:"abstract"`
	Authors     []string  `json:"authors"`
	Keywords    []string  `json:"keywords,omitempty"`
	PeerReviews [3]string `json:"peer_reviews"` // PeerReview composite keys
	Status      string    `json:"status"`
	ReviewTime  int64     `json:"review_time"`
}

func (p Paper) ObjectType() string {
	return ObjectTypePaper
}
