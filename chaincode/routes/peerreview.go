package routes

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

func createPeerReview(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function createPeerReview requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	paperId := args[0]
	reviewer := args[1]
	peerReview := lib.PeerReview{
		Paper:    paperId,
		Reviewer: reviewer,
		Status:   lib.StatusReviewing,
	}
	payload, err := utils.PutLedger(stub, peerReview)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}

func UpdatePeerReview(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	now := time.Now().UnixNano()
	if len(args) < 4 {
		return shim.Error("function UpdatePeerReview requires 4 arguments")
	}
	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" {
		return shim.Error("arguments should be nonempty")
	}
	paperId := args[0]
	reviewer := args[1]
	status := args[2]
	comment := args[3]
	resp := RetrievePeerReviewByIds(stub, []string{paperId, reviewer})
	if resp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to retrieve peer review information: %s", resp.Message))
	}
	if resp.Payload == nil {
		return shim.Error("peer review information not found")
	}
	var peerReview lib.PeerReview
	err := json.Unmarshal(resp.Payload, &peerReview)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to unserialise peer review information: %s", err.Error()))
	}
	if peerReview.Status != lib.StatusReviewing {
		return shim.Error("this paper has finished peer review")
	}
	if status != lib.StatusAccepted && status != lib.StatusRejected {
		return shim.Error(fmt.Sprintf("invalid status: %s", status))
	}
	peerReview.Status = status
	peerReview.Comment = comment
	peerReview.Time = now
	payload, err := utils.PutLedger(stub, peerReview)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to update ledger: %s", err.Error()))
	}
	_ = UpdatePaperStatus(stub, []string{paperId})
	_ = updateUserReviewingSub(stub, []string{reviewer})
	return shim.Success(payload)
}

func RetrievePeerReviewsByPaperId(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("function RetrievePeerReviewsByPaperId requires 1 argument")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	paperId := args[0]
	query := `{"selector":{"peer_review_paper":"` + paperId + `"}}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to query: %s", err.Error()))
	}
	var peerReviews []lib.PeerReview
	for _, peerReviewBytes := range results {
		var peerReview lib.PeerReview
		err := json.Unmarshal(peerReviewBytes, &peerReview)
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to unserialise peer review information: %s", err.Error()))
		}
		peerReviews = append(peerReviews, peerReview)
	}
	if len(peerReviews) != 3 {
		return shim.Error(fmt.Sprintf("invalid length of peer review: %d", len(peerReviews)))
	}
	peerReviewsBytes, err := json.Marshal(peerReviews)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise peer review information: %s", err.Error()))
	}
	return shim.Success(peerReviewsBytes)
}

func RetrievePeerReviewByIds(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function RetrievePeerReviewByIds requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	paperId := args[0]
	reviewer := args[1]
	peerReviewBytes, err := utils.GetByKeys(stub, lib.ObjectTypePeerReview, paperId, reviewer)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve peer review: %s", err.Error()))
	}
	return shim.Success(peerReviewBytes)
}
