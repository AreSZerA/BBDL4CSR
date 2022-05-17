// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for updating peer review information.

package update

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

// PeerReviewByPaperAndReviewer in package `update` updates the peer review information.
// It requires 4 necessary argument: paper ID, reviewer email, acceptance, and comment.
func PeerReviewByPaperAndReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// record current timestamp
	now := time.Now().UnixNano()
	// check arguments
	err := utils.CheckArgs(args, 4)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId, reviewer, status, comment := args[0], args[1], args[2], args[3]

	//  retrieve peer review object by paper ID and reviewer email
	result, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, reviewer)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the peer review information does not exist
	if len(result) == 0 {
		return shim.Error(lib.ErrPeerReviewNotFound.Error())
	}
	var peerReview lib.PeerReview
	// deserialize the string to peer review object
	_ = json.Unmarshal(result, &peerReview)

	// if the status has been reviewed, response error
	if peerReview.Status == lib.StatusAccepted || status == lib.StatusRejected {
		return shim.Error(lib.ErrPeerReviewDone.Error())
	}
	// otherwise, update the fields
	peerReview.Status = status
	peerReview.Comment = comment
	peerReview.Time = now
	// put the peer review object into the ledger to update
	peerReviewBytes, err := utils.PutLedger(stub, peerReview)
	if err != nil {
		return shim.Error(err.Error())
	}

	// update reviewer information
	// retrieve review by email
	userBytes, err := utils.GetByKeys(stub, lib.ObjectTypeUser, reviewer)
	if err != nil {
		return shim.Error(err.Error())
	}
	// if the result is empty, user does not exist
	if len(userBytes) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(userBytes, &user)
	// update the User.Reviewing field
	user.Reviewing--
	// put the user object into the ledger to update
	_, err = utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}

	// update paper information
	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, paperId)
	if err != nil {
		return shim.Error(err.Error())
	}
	// if the result is empty, the paper does not exist
	if len(paperBytes) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	var paper lib.Paper
	// deserialize the string to paper object
	_ = json.Unmarshal(paperBytes, &paper)
	// record statuses
	var statuses = []string{status}
	// traverse the Paper.Reviewers to find the peer review information
	for _, r := range paper.Reviewers {
		if r != reviewer {
			var pr lib.PeerReview
			// retrieve peer review information by paper ID and reviewer email
			res, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, r)
			if err != nil {
				return shim.Error(err.Error())
			}
			// if the result is empty, response error message
			if len(res) == 0 {
				shim.Error(lib.ErrDataHasLost.Error())
			} else {
				// otherwise, deserialize the string to peer review object
				_ = json.Unmarshal(res, &pr)
			}
			// record status
			statuses = append(statuses, pr.Status)
		}
	}

	// get the final status
	paper.Status = utils.GetStatus(statuses[0], statuses[1], statuses[2])
	// if the current status is not reviewing, update the publishing time
	if paper.Status != lib.StatusReviewing {
		paper.PublishTime = now
	}
	// put the paper object into the ledger to update
	_, err = utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(peerReviewBytes)
}
