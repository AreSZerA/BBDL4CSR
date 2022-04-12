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

type BlockchainObject interface {
	ObjectType() string
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
