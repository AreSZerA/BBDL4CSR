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
	_ = UserByEmail(stub, []string{reviewer})
	_ = PaperById(stub, []string{paperId})
	return shim.Success(peerReviewBytes)
}
