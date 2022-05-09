// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for retrieving users.

package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// queryUsers is used in this package to simplify the process of query.
func queryUsers(stub shim.ChaincodeStubInterface, query string) ([]lib.User, error) {
	// retrieve results by query
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return nil, err
	}
	var users []lib.User
	// traverse the results, deserialize and append to the slice
	for _, result := range results {
		var user lib.User
		// deserialize the string to user object
		_ = json.Unmarshal(result, &user)
		users = append(users, user)
	}
	return users, nil
}

// UsersByQuery in package `retrieve` retrieves users by query.
// It requires one necessary argument: query.
func UsersByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// Users in package `retrieve` retrieves all the users.
// No necessary argument is required.
func Users(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// get all the users
	results, err := utils.GetAll(stub, lib.ObjectTypeUser)
	if err != nil {
		return shim.Error(err.Error())
	}
	var users []lib.User
	// traverse the results, deserialize and append to the slice
	for _, result := range results {
		var user lib.User
		// deserialize the string to user object
		_ = json.Unmarshal(result, &user)
		users = append(users, user)
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// UsersSortByEmail in package `retrieve` retrieves all the users and sort by email in ascending order.
// No necessary argument is required.
func UsersSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM user ORDER BY (email ASC)
	query := `{"selector":{},"sort":[{"user_email":"asc"}]}`
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// UsersSortByName in package `retrieve` retrieves all the users and sort by name in ascending order.
// No necessary argument is required.
func UsersSortByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM user ORDER BY (name ASC)
	query := `{"selector":{},"sort":[{"user_name":"asc"}]}`
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// UsersByNameSortByEmail in package `retrieve` retrieves users by username and sort by email in ascending order.
// It requires one necessary argument: username.
func UsersByNameSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	name := args[0]
	// SELECT user WHERE name = #{name} ORDER BY (email ASC)
	query := fmt.Sprintf(`{"selector":{"user_name":"%s"},"sort":[{"user_email":"asc"}]}`, name)
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// ReviewersSortByEmail in package `retrieve` retrieves reviewers sort by email in ascending order.
// No necessary argument is required.
func ReviewersSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM user WHERE is_reviewer ORDER BY (email ASC)
	query := `{"selector":{"user_is_reviewer":true},"sort":[{"user_email":"asc"}]}`
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// ReviewersSortByName in package `retrieve` retrieves reviewers sort by name in ascending order.
// No necessary argument is required.
func ReviewersSortByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM user WHERE is_reviewer ORDER BY (name ASC)
	query := `{"selector":{"user_is_reviewer":true},"sort":[{"user_name":"asc"}]}`
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// ReviewersByPaperIdSortByEmail in package `retrieve` retrieves reviewers by paper ID and sort by email in ascending order.
// It requires one necessary argument: paper ID.
func ReviewersByPaperIdSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]
	// SELECT * FROM peer_review WHERE paper = #{paperId}
	query := `{"selector":{"peer_review_paper":"` + paperId + `"}}`
	// retrieve peer reviews by paper ID
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// if the result is empty, the paper does not exist
	if len(results) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	// SELECT * FROM user WHERE email = #{email1} OR email = #{email2} OR email = #{email2} ORDER BY (email ASC)
	query = `{"selector":{"user_email":{"$or":["`
	// retrieve the peer reviews to find emails
	for i, result := range results {
		var peerReview lib.PeerReview
		// deserialize the string to peer review object
		_ = json.Unmarshal(result, &peerReview)
		query += peerReview.Reviewer
		if i < len(results)-1 {
			query += `","`
		} else {
			query += `"]}},"sort":[{"user_email":"asc"}]}`
		}
	}
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// ReviewersByPaperIdSortByName in package `retrieve` retrieves reviewers by paper ID and sort by name in ascending order.
// It requires one necessary argument: paper ID.
func ReviewersByPaperIdSortByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]
	// SELECT * FROM peer_review WHERE paper = #{paperId}
	query := `{"selector":{"peer_review_paper":"` + paperId + `"}}`
	// retrieve peer reviews by paper ID
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// if the result is empty, the paper does not exist
	if len(results) == 0 {
		return shim.Error(lib.ErrPaperNotFound.Error())
	}
	// SELECT * FROM user WHERE email = #{email1} OR email = #{email2} OR email = #{email2} ORDER BY (name ASC)
	query = `{"selector":{"user_email":{"$or":["`
	// retrieve the peer reviews to find emails
	for i, result := range results {
		var peerReview lib.PeerReview
		// deserialize the string to peer review object
		_ = json.Unmarshal(result, &peerReview)
		query += peerReview.Reviewer
		if i < len(results)-1 {
			query += `","`
		} else {
			query += `"]}},"sort":[{"user_name":"asc"}]}`
		}
	}
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// AdminsSortByEmail in package `retrieve` retrieves administrators sort by email in ascending order.
// No necessary argument is required.
func AdminsSortByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM user WHERE is_admin ORDER BY (email ASC)
	query := `{"selector":{"user_is_admin":true},"sort":[{"user_email":"asc"}]}`
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// AdminsSortByName in package `retrieve` retrieves administrators sort by name in ascending order.
// No necessary argument is required.
func AdminsSortByName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM user WHERE is_admin ORDER BY (name ASC)
	query := `{"selector":{"user_is_admin":true},"sort":[{"user_name":"asc"}]}`
	// query to get results
	users, err := queryUsers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	usersBytes, err := utils.MarshalByArgs(users, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(usersBytes)
}

// UserByEmail in package `retrieve` retrieves the exact one user by email.
// It requires one necessary argument: email.
func UserByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	// retrieve paper by composite key
	userBytes, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}
