package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func queryUsers(stub shim.ChaincodeStubInterface, query string) peer.Response {
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var users []lib.User
	for _, result := range results {
		var user lib.User
		_ = json.Unmarshal(result, &user)
		users = append(users, user)
	}
	usersBytes, _ := json.Marshal(users)
	return shim.Success(usersBytes)
}

func UsersByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	return queryUsers(stub, query)
}

func Users(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	results, err := utils.GetAll(stub, lib.ObjectTypeUser)
	if err != nil {
		return shim.Error(err.Error())
	}
	var users []lib.User
	for _, result := range results {
		var user lib.User
		_ = json.Unmarshal(result, &user)
		users = append(users, user)
	}
	usersBytes, _ := json.Marshal(users)
	return shim.Success(usersBytes)
}

func UsersSortByEmail(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"sort":[{"user_email":"asc"}]"}`
	return queryUsers(stub, query)
}

func UsersSortByName(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"sort":[{"user_name":"asc"}]"}`
	return queryUsers(stub, query)
}

func UsersByNameSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	name := args[0]
	query := fmt.Sprintf(`{"selector":{"user_name":"%s"},"sort":[{"user_email":"asc"}]}`, name)
	return queryUsers(stub, query)
}

func ReviewersSortByEmail(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"user_is_reviewer":true},"sort":[{"user_email":"asc"}]}`
	return queryUsers(stub, query)
}

func ReviewersSortByName(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"user_is_reviewer":true},"sort":[{"user_name":"asc"}]}`
	return queryUsers(stub, query)
}

func ReviewersByPaperIdSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]
	query := `{"selector":{"peer_review_paper":"` + paperId + `"}}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(results) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	query = `{"selector":{"user_email":{"$or":["`
	for i, result := range results {
		var peerReview lib.PeerReview
		_ = json.Unmarshal(result, &peerReview)
		query += peerReview.Reviewer
		if i < len(results)-1 {
			query += `","`
		} else {
			query += `"]}},"sort":[{"user_email":"asc"}]}`
		}
	}
	return queryUsers(stub, query)
}

func ReviewersByPaperIdSortByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]
	query := `{"selector":{"peer_review_paper":"` + paperId + `"}}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(results) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	query = `{"selector":{"user_email":{"$or":["`
	for i, result := range results {
		var peerReview lib.PeerReview
		_ = json.Unmarshal(result, &peerReview)
		query += peerReview.Reviewer
		if i < len(results)-1 {
			query += `","`
		} else {
			query += `"]}},"sort":[{"user_name":"asc"}]}`
		}
	}
	return queryUsers(stub, query)
}

func AdminsSortByEmail(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"user_is_admin":true},"sort":[{"user_email":"asc"}]}`
	return queryUsers(stub, query)
}

func AdminsSortByName(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"user_is_admin":true},"sort":[{"user_name":"asc"}]}`
	return queryUsers(stub, query)
}

func UserByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	userBytes, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(userBytes) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	return shim.Success(userBytes)
}
