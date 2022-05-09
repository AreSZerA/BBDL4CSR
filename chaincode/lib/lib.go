// Copyright 2022 AreSZerA. All rights reserved.
// This file defines Status, BlockchainObjectInterface interface, and some constants.

package lib

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

// Status is an alias of uint8.
type Status uint8

// String converts a numeric Status to string.
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

// StatusCode converts a string to numeric Status.
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

// BlockchainObjectInterface must be implemented by the objects who will be stored in ledger.
// utils.PutLedger and utils.DelLedger use its method to obtain object type and attributes to create composite key,
// then update the ledger by it.
type BlockchainObjectInterface interface {
	// Type returns a string as the type name of the object.
	Type() string
	// Attributes returns a string list as the attributes to create composite key.
	Attributes() []string
}
