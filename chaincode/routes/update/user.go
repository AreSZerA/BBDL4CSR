// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for updating user information.

package update

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// UserByEmail in package `update` synchronizes the user_reviewing value automatically.
// It requires one necessary argument: email.
func UserByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	// SELECT * FROM peer_review WHERE reviewer = #{email} AND status = reviewing
	query := `{"selector":{"$and":[{"peer_review_reviewer":"` + email + `"},{"peer_review_status":"reviewing"}]}}`
	// query to find the peer review tasks
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the reviewing numbers are different
	if user.Reviewing != uint16(len(results)) {
		// update the User.Reviewing field
		user.Reviewing = uint16(len(results))
		// update the ledger
		userBytes, err := utils.PutLedger(stub, user)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(userBytes)
	}
	return shim.Success(result)
}

// UserName in package `update` updates the username.
// It requires 2 necessary arguments: email and new username.
func UserName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 2)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, name := args[0], args[1]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	// update the User.Name field
	user.Name = name
	// update the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// UserPassword in package `update` updates the user password.
// It requires 2 necessary arguments: email and new password.
func UserPassword(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 2)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, password := args[0], args[1]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	user.Password = password
	// update the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// UserIsReviewer in package `update` updates the user to reviewer.
// It requires 1 necessary argument: email.
func UserIsReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	// update the User.IsReviewer field to true
	user.IsReviewer = true
	// update the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// UserIsNotReviewer in package `update` updates the reviewer to ordinary user.
// It requires 1 necessary argument: email.
func UserIsNotReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	// update the User.IsReviewer field to false
	user.IsReviewer = false
	// update the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// UserIsAdmin in package `update` updates the user to administrator.
// It requires 1 necessary argument: email.
func UserIsAdmin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	// update the User.IsAdmin field to true
	user.IsAdmin = true
	// update the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// UserIsNotAdmin UserIsReviewer in package `update` updates the administrator to ordinary user.
// It requires 1 necessary argument: email.
func UserIsNotAdmin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is empty, the email has not been registered
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	// deserialize the string to user object
	_ = json.Unmarshal(result, &user)

	// update the User.IsAdmin field to false
	user.IsAdmin = false
	// update the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}
