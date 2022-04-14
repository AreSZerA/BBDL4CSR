package lib

type Status uint8

func (s Status) String() string {
	switch s {
	case StatusCodeRejected:
		return StatusRejected
	case StatusCodeAccepted:
		return StatusAccepted
	default:
		return StatusReviewing
	}
}

func StatusCode(s string) Status {
	switch s {
	case StatusRejected:
		return StatusCodeRejected
	case StatusAccepted:
		return StatusCodeAccepted
	default:
		return StatusCodeReviewing
	}
}

type BlockchainObject interface {
	Type() string
	Keys() []string
}

const (
	StatusCodeReviewing Status = 1
	StatusCodeRejected  Status = 2
	StatusCodeAccepted  Status = 4

	StatusReviewing = "reviewing"
	StatusRejected  = "rejected"
	StatusAccepted  = "accepted"

	ObjectTypeUser       = "ObjectUser"
	ObjectTypePaper      = "ObjectPaper"
	ObjectTypePeerReview = "ObjectPeerReview"
)
