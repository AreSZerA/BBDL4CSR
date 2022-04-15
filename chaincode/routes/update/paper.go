package update

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func PaperById(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]
	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, paperId)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(paperBytes) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	var paper lib.Paper
	_ = json.Unmarshal(paperBytes, &paper)
	var statuses []string
	var final int64
	for _, reviewer := range paper.Reviewers {
		var peerReview lib.PeerReview
		result, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, reviewer)
		if err != nil {
			return shim.Error(err.Error())
		}
		if len(result) == 0 {
			peerReview = lib.PeerReview{Paper: paperId, Reviewer: reviewer, Status: lib.StatusReviewing}
			go func() { _, _ = utils.PutLedger(stub, peerReview) }()
		} else {
			_ = json.Unmarshal(result, &peerReview)
			if final > peerReview.Time {
				final = peerReview.Time
			}
		}
		statuses = append(statuses, peerReview.Status)
	}
	paper.Status = utils.GetStatus(statuses[0], statuses[1], statuses[2])
	if paper.Status != lib.StatusReviewing {
		paper.ReviewTime = final
	}
	paperBytes, err = utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
