package main

import (
	"chaincode/route"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"os"
)

type DigitalLibrary struct {
}

func (library *DigitalLibrary) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (library *DigitalLibrary) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fcn, args := stub.GetFunctionAndParameters()
	switch fcn {
	case "CreateUser":
		return route.CreateUser(stub, args)
	case "UpdateUserName":
		return route.UpdateUserName(stub, args)
	case "UpdateUserPasswd":
		return route.UpdateUserPasswd(stub, args)
	case "UpdateUserIsReviewer":
		return route.UpdateUserIsReviewer(stub, args)
	case "UpdateUserIsAdmin":
		return route.UpdateUserIsAdmin(stub, args)
	case "RetrieveUserByKey":
		return route.RetrieveUserByKey(stub, args)
	case "RetrieveUserByEmail":
		return route.RetrieveUserByEmail(stub, args)
	case "RetrieveAllUsers":
		return route.RetrieveAllUsers(stub)
	case "RetrieveAllReviewers":
		return route.RetrieveAllReviewers(stub)
	case "test":
		return shim.Success([]byte(fmt.Sprintf("Chaincode invocation test passed, args: %s", args)))
	default:
		return shim.Error(fmt.Sprintf("Function not implemeted: %s", fcn))
	}
}

func main() {
	err := shim.Start(new(DigitalLibrary))
	if err != nil {
		fmt.Printf("Error: failed to start chaincode: %s\n", err.Error())
		os.Exit(-1)
	}
}
