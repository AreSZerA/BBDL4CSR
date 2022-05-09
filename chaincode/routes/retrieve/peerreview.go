// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for retrieving peer review objects.

package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// queryPapers is used in this package to simplify the process of query.
func queryPeerReviews(stub shim.ChaincodeStubInterface, query string) ([]lib.PeerReview, error) {
	// retrieve results by query
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return nil, err
	}
	var peerReviewers []lib.PeerReview
	// traverse the results, deserialize and append to the slice
	for _, result := range results {
		var peerReview lib.PeerReview
		// deserialize the string to user object
		_ = json.Unmarshal(result, &peerReview)
		peerReviewers = append(peerReviewers, peerReview)
	}
	return peerReviewers, nil
}

// PeerReviewsByQuery in package `retrieve` retrieves peer review objects by query.
// It requires one necessary argument: query.
func PeerReviewsByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	// query to get results
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	peerReviewsBytes, err := utils.MarshalByArgs(peerReviews, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(peerReviewsBytes)
}

// PeerReviewsByReviewerSortByCreateTime in package `retrieve` retrieves peer review objects by reviewer email and sort by create time in descending order.
// It requires one necessary argument: reviewer email.
func PeerReviewsByReviewerSortByCreateTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	reviewer := args[0]
	// SELECT * FROM peer_review WHERE reviewer = #{reviewer} ORDER BY (create_time DESC)
	query := fmt.Sprintf(`{"selector":{"peer_review_reviewer":"%s"},"sort":[{"peer_review_create_time":"desc"}]}`, reviewer)
	// query to get results
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	peerReviewsBytes, err := utils.MarshalByArgs(peerReviews, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(peerReviewsBytes)
}

// AcceptedPeerReviewsByReviewerSortByCreateTime in package `retrieve` retrieves accepted peer review objects by reviewer email and sort by create time in descending order.
// It requires one necessary argument: reviewer email.
func AcceptedPeerReviewsByReviewerSortByCreateTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	reviewer := args[0]
	// SELECT * FROM peer_review WHERE reviewer = #{reviewer} AND status = accepted ORDER BY (create_time DESC)
	query := fmt.Sprintf(`{"selector":{"peer_review_reviewer":"%s","peer_review_status":"accepted"},"sort":[{"peer_review_create_time":"desc"}]}`, reviewer)
	// query to get results
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	peerReviewsBytes, err := utils.MarshalByArgs(peerReviews, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(peerReviewsBytes)
}

// RejectedPeerReviewsByReviewerSortByCreateTime in package `retrieve` retrieves rejected peer review objects by reviewer email and sort by create time in descending order.
// It requires one necessary argument: reviewer email.
func RejectedPeerReviewsByReviewerSortByCreateTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	reviewer := args[0]
	// SELECT * FROM peer_review WHERE reviewer = #{reviewer} AND status = rejected ORDER BY (create_time DESC)
	query := fmt.Sprintf(`{"selector":{"peer_review_reviewer":"%s","peer_review_status":"rejected"},"sort":[{"peer_review_create_time":"desc"}]}`, reviewer)
	// query to get results
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	peerReviewsBytes, err := utils.MarshalByArgs(peerReviews, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(peerReviewsBytes)
}

// ReviewingPeerReviewsByReviewerSortByCreateTime in package `retrieve` retrieves ongoing peer review objects by reviewer email and sort by create time in descending order.
// It requires one necessary argument: reviewer email.
func ReviewingPeerReviewsByReviewerSortByCreateTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	reviewer := args[0]
	// SELECT * FROM peer_review WHERE reviewer = #{reviewer} AND status = reviewing ORDER BY (create_time DESC)
	query := fmt.Sprintf(`{"selector":{"peer_review_reviewer":"%s","peer_review_status":"reviewing"},"sort":[{"peer_review_create_time":"desc"}]}`, reviewer)
	// query to get results
	peerReviews, err := queryPeerReviews(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	peerReviewsBytes, err := utils.MarshalByArgs(peerReviews, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(peerReviewsBytes)
}
