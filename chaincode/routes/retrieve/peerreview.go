package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func queryPeerReviews(stub shim.ChaincodeStubInterface, query string) peer.Response {
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var peerReviewers []lib.PeerReview
	for _, result := range results {
		var peerReview lib.PeerReview
		_ = json.Unmarshal(result, &peerReview)
		peerReviewers = append(peerReviewers, peerReview)
	}
	peerReviewersBytes, _ := json.Marshal(peerReviewers)
	return shim.Success(peerReviewersBytes)
}

func PeerReviewsByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	return queryPeerReviews(stub, query)
}

func PeerReviewsByReviewerSortByCreateTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	reviewer := args[0]
	query := fmt.Sprintf(`{"selector":{"peer_review_reviewer":"%s"},"sort":[{"peer_review_create_time":"desc"}]}`, reviewer)
	return queryPeerReviews(stub, query)
}
