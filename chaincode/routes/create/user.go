// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the functions for creating ordinary users, reviewers, and administrators.

package create

import (
	"chaincode/lib"
	"chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// User in package `create` puts a user object to the ledger.
// It requires 3 necessary arguments: email, username, password.
func User(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 3)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, username, password := args[0], args[1], args[2]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is not empty, this email has been registered, so response error
	if len(result) != 0 {
		return shim.Error(lib.ErrUserExists.Error())
	}

	// construct user object
	user := lib.User{Email: email, Name: username, Password: password}
	// put the user object into the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// Reviewer in package `create` puts a user object to the ledger, in which the IsReview field is true.
// It requires 3 necessary arguments: email, username, password.
func Reviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 3)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, username, password := args[0], args[1], args[2]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is not empty, this email has been registered, so response error
	if len(result) != 0 {
		return shim.Error(lib.ErrUserExists.Error())
	}

	// construct user object
	user := lib.User{Email: email, Name: username, Password: password, IsReviewer: true}
	// put the user object into the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

// Admin in package `create` puts a user object to the ledger, in which the IsAdmin field is true.
// It requires 3 necessary arguments: email, username, password.
func Admin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 3)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, username, password := args[0], args[1], args[2]

	// retrieve user by email
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	// when the result is not empty, this email has been registered, so response error
	if len(result) != 0 {
		return shim.Error(lib.ErrUserExists.Error())
	}

	// construct user object
	user := lib.User{Email: email, Name: username, Password: password, IsAdmin: true}
	// put the user object into the ledger
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}
