// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for updating paper information.

package update

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// PaperById in package `update` synchronizes the status and publishing time value automatically.
// It requires one necessary argument: paper ID.
func PaperById(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]

	// retrieve paper by paper ID
	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, paperId)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the paper does not exist
	if len(paperBytes) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	var paper lib.Paper
	// deserialize the string to paper object
	_ = json.Unmarshal(paperBytes, &paper)

	var statuses []string
	var final int64
	// traverse the Paper.Reviewers to find the peer review information
	for _, reviewer := range paper.Reviewers {
		var peerReview lib.PeerReview
		// retrieve peer review information by paper ID and reviewer email
		result, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, reviewer)
		if err != nil {
			return shim.Error(err.Error())
		}
		// if the result is empty
		if len(result) == 0 {
			// construct peer review object
			peerReview = lib.PeerReview{Paper: paperId, Reviewer: reviewer, Status: lib.StatusReviewing}
			// put the peer review object to the ledger
			_, _ = utils.PutLedger(stub, peerReview)
		} else {
			// otherwise, deserialize the string to peer review object
			_ = json.Unmarshal(result, &peerReview)
			// record the peer review time
			if final < peerReview.Time {
				final = peerReview.Time
			}
		}
		// record status
		statuses = append(statuses, peerReview.Status)
	}

	// get the final status
	paper.Status = utils.GetStatus(statuses[0], statuses[1], statuses[2])
	// if the current status is not reviewing, update the publishing time
	if paper.Status != lib.StatusReviewing {
		paper.PublishTime = final
	}
	// put the paper object into the ledger to update
	paperBytes, err = utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
