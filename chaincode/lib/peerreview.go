// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the structure of peer review action.

package lib

// PeerReview implements BlockchainObjectInterface for storing in the ledger.
type PeerReview struct {
	// Unique ID for identifying papers, related to Paper.ID.
	Paper string `json:"peer_review_paper"`
	// Email of the reviewer, related to User.Email.
	Reviewer string `json:"peer_review_reviewer"`
	// Timestamp of create time, initialized when the object is instantiated.
	CreateTime int64 `json:"peer_review_create_time"`
	// Status of peer review, could be StatusReviewing, StatusAccepted, or StatusRejected.
	Status string `json:"peer_review_status"`
	// Comment to the paper, given by the reviewer.
	Comment string `json:"peer_review_comment,omitempty"`
	// Timestamp of peer review, updated by peer review action.
	Time int64 `json:"peer_review_time,omitempty"`
}

// Type returns ObjectTypePeerReview.
func (p PeerReview) Type() string {
	return ObjectTypePeerReview
}

// Attributes returns a slice constructed with PeerReview.Paper and Reviewer.
func (p PeerReview) Attributes() []string {
	return []string{p.Paper, p.Reviewer}
}
