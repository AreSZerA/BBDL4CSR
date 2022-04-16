package create

import (
	"chaincode/lib"
	"chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func User(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 3)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, username, password := args[0], args[1], args[2]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) != 0 {
		return shim.Error(lib.ErrUserExists.Error())
	}
	user := lib.User{Email: email, Name: username, Password: password}
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func Reviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 3)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, username, password := args[0], args[1], args[2]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) != 0 {
		return shim.Error(lib.ErrUserExists.Error())
	}
	user := lib.User{Email: email, Name: username, Password: password, IsReviewer: true}
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}

func Admin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 3)
	if err != nil {
		return shim.Error(err.Error())
	}
	email, username, password := args[0], args[1], args[2]
	result, err := utils.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(err.Error())
	}
	if len(result) != 0 {
		return shim.Error(lib.ErrUserExists.Error())
	}
	user := lib.User{Email: email, Name: username, Password: password, IsAdmin: true}
	userBytes, err := utils.PutLedger(stub, user)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(userBytes)
}
