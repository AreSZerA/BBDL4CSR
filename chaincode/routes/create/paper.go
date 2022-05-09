// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for creating paper.

package create

import (
	"chaincode/lib"
	"chaincode/utils"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

// Paper in package `create` puts a paper object to the ledger.
// It requires 5 necessary arguments: uploader, title, abstract, authors in JSON, and keywords in JSON.
func Paper(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// record current time
	now := time.Now().UnixNano()
	// check arguments
	err := utils.CheckArgs(args, 5)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader, title, abstract, authorsString, keywordsString := args[0], args[1], args[2], args[3], args[4]
	// deserialize authors to slice
	var authors []string
	err = json.Unmarshal([]byte(authorsString), &authors)
	if err != nil {
		return shim.Error(err.Error())
	}
	// deserialize keywords to slice
	var keywords []string
	err = json.Unmarshal([]byte(keywordsString), &keywords)
	if err != nil {
		return shim.Error(err.Error())
	}

	// retrieve user by email to make sure that the user exists
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, uploader)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}

	// get the reviewer list sorted by number of reviewing papers in ascending order
	// SELECT * FROM user WHERE user_email != #{uploader} AND user_is_reviewer ORDER BY (user_reviewing ASC)
	query := `{"selector":{"$and":[{"$ne":{"user_email":"` + uploader + `"}},{"user_is_reviewer":true}]},"sort":[{"user_reviewing":"asc"}]}`
	// execute the query to get results
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// if the number of reviewers is less than 3, response error
	if len(results) < 3 {
		return shim.Error(lib.ErrReviewerNotEnough.Error())
	}
	// calculate paper ID by MD5
	id := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", uploader, now))))
	var reviewers [3]string
	// select the top three reviewers to peer review this paper
	for i := 0; i < 3; i++ {
		var user lib.User
		_ = json.Unmarshal(results[i], &user)
		user.Reviewing++
		reviewers[i] = user.Email
		peerReview := lib.PeerReview{Paper: id, Reviewer: user.Email, CreateTime: now, Status: lib.StatusReviewing}
		// update peer review and user information, errors will be ignored
		_, _ = utils.PutLedger(stub, peerReview)
		_, _ = utils.PutLedger(stub, user)
	}

	// construct paper object
	paper := lib.Paper{
		ID:         id,
		Uploader:   uploader,
		UploadTime: now,
		Title:      title,
		Abstract:   abstract,
		Authors:    authors,
		Keywords:   keywords,
		Reviewers:  reviewers,
		Status:     lib.StatusReviewing,
	}
	// put the paper object into the ledger
	paperBytes, err := utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
