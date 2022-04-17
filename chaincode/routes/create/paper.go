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

func Paper(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	now := time.Now().UnixNano()
	err := utils.CheckArgs(args, 5)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader, title, abstract, authorsString, keywordsString := args[0], args[1], args[2], args[3], args[4]
	var authors []string
	err = json.Unmarshal([]byte(authorsString), &authors)
	if err != nil {
		return shim.Error(err.Error())
	}
	var keywords []string
	err = json.Unmarshal([]byte(keywordsString), &keywords)
	if err != nil {
		return shim.Error(err.Error())
	}
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, uploader)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	query := `{"selector":{"$and":[{"$ne":{"user_email":"` + uploader + `"}},{"user_is_reviewer":true}]},"sort":[{"user_reviewing":"asc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(results) < 3 {
		return shim.Error(lib.ErrReviewerNotEnough.Error())
	}
	id := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", uploader, now))))
	var reviewers [3]string
	for i := 0; i < 3; i++ {
		var user lib.User
		_ = json.Unmarshal(results[i], &user)
		user.Reviewing++
		reviewers[i] = user.Email
		peerReview := lib.PeerReview{Paper: id, Reviewer: user.Email, CreateTime: now, Status: lib.StatusReviewing}
		_, _ = utils.PutLedger(stub, peerReview)
		_, _ = utils.PutLedger(stub, user)
	}
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
	paperBytes, err := utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
