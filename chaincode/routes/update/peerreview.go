package update

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

func PeerReviewByPaperAndReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	now := time.Now().UnixNano()
	err := utils.CheckArgs(args, 4)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId, reviewer, status, comment := args[0], args[1], args[2], args[3]
	result, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, reviewer)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrPeerReviewNotFound.Error())
	}
	var peerReview lib.PeerReview
	_ = json.Unmarshal(result, &peerReview)
	if status != "accepted" && status != "rejected" {
		return shim.Success(result)
	}
	peerReview.Status = status
	peerReview.Comment = comment
	peerReview.Time = now
	peerReviewBytes, err := utils.PutLedger(stub, peerReview)
	if err != nil {
		return shim.Error(err.Error())
	}

	userBytes, err := utils.GetByKeys(stub, lib.ObjectTypeUser, reviewer)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(userBytes) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(userBytes, &user)
	user.Reviewing--
	_, err = utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}

	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, paperId)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(paperBytes) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	var paper lib.Paper
	_ = json.Unmarshal(paperBytes, &paper)
	var statuses = []string{status}
	var final int64
	for _, r := range paper.Reviewers {
		if r != reviewer {
			var pr lib.PeerReview
			res, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, r)
			if err != nil {
				return shim.Error(err.Error())
			}
			if len(res) == 0 {
				shim.Error(lib.ErrDataHasLost.Error())
			} else {
				_ = json.Unmarshal(res, &pr)
				if final < pr.Time {
					final = pr.Time
				}
			}
			statuses = append(statuses, pr.Status)
		}
	}
	paper.Status = utils.GetStatus(statuses[0], statuses[1], statuses[2])
	if paper.Status != lib.StatusReviewing {
		paper.PublishTime = final
	}
	paperBytes, err = utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(peerReviewBytes)
}
