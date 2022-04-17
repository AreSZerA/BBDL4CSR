package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func queryPeerReviews(stub shim.ChaincodeStubInterface, query string) ([]lib.PeerReview, error) {
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return nil, err
	}
	var peerReviewers []lib.PeerReview
	for _, result := range results {
		var peerReview lib.PeerReview
		_ = json.Unmarshal(result, &peerReview)
		peerReviewers = append(peerReviewers, peerReview)
	}
	return peerReviewers, nil
}

func PeerReviewsByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var peerReviewsBytes []byte
	if len(args) >= 3 {
		peerReviewsBytes, _ = utils.MarshalWithOffsetAndLimit(peerReviews, args[1], args[2])
	} else {
		peerReviewsBytes, _ = json.Marshal(peerReviews)
	}
	return shim.Success(peerReviewsBytes)
}

func PeerReviewsByReviewerSortByCreateTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	reviewer := args[0]
	query := fmt.Sprintf(`{"selector":{"peer_review_reviewer":"%s"},"sort":[{"peer_review_create_time":"desc"}]}`, reviewer)
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var peerReviewsBytes []byte
	if len(args) >= 3 {
		peerReviewsBytes, _ = utils.MarshalWithOffsetAndLimit(peerReviews, args[1], args[2])
	} else {
		peerReviewsBytes, _ = json.Marshal(peerReviews)
	}
	return shim.Success(peerReviewsBytes)
}
