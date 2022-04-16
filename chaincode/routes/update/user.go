package update

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func UserByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	query := `{"selector":{"$and":[{"peer_review_reviewer":"` + email + `"},{"peer_review_status":"reviewing"}]}}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	if user.Reviewing != uint16(len(results)) {
		user.Reviewing = uint16(len(results))
		userBytes, err := utils.PutLedger(stub, user)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(userBytes)
	}
	return shim.Success(result)
}

func UserName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 2)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, name := args[0], args[1]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	user.Name = name
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func UserPassword(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 2)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, password := args[0], args[1]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	user.Password = password
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func UserIsReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	user.IsReviewer = true
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func UserIsNotReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	user.IsReviewer = false
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func UserIsAdmin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	user.IsAdmin = true
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func UserIsNotAdmin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	email := args[0]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	user.IsAdmin = false
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func UserReviewing(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 2)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, reviewingString := args[0], args[1]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) == 0 {
		return shim.Error(lib.ErrUserNotFound.Error())
	}
	var user lib.User
	_ = json.Unmarshal(result, &user)
	reviewing, err := strconv.Atoi(reviewingString)
	if err == nil && reviewing >= 0 {
		user.Reviewing = uint16(reviewing)
	}
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}
