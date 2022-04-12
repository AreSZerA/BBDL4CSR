package main

import (
	"chaincode/route"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"os"
)

var funcMap = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	"test": func(stub shim.ChaincodeStubInterface, args []string) peer.Response {
		return shim.Success([]byte(fmt.Sprintf("Chaincode invocation test passed, args: %s", args)))
	},
	"CreateUser":           route.CreateUser,
	"CreateReviewer":       route.CreateReviewer,
	"UpdateUserName":       route.UpdateUserName,
	"UpdateUserPasswd":     route.UpdateUserPasswd,
	"UpdateUserIsReviewer": route.UpdateUserIsReviewer,
	"UpdateUserIsAdmin":    route.UpdateUserIsAdmin,
	"RetrieveUserByKey":    route.RetrieveUserByKey,
	"RetrieveUserByEmail":  route.RetrieveUserByEmail,
	"RetrieveAllUsers":     route.RetrieveAllUsers,
	"RetrieveAllReviewers": route.RetrieveAllReviewers,
}

type DigitalLibrary struct {
}

func (library *DigitalLibrary) Init(_ shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (library *DigitalLibrary) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	functionName, args := stub.GetFunctionAndParameters()
	function, ok := funcMap[functionName]
	if ok {
		return function(stub, args)
	}
	return shim.Error(fmt.Sprintf("Function not implemeted: %s", functionName))
}

func main() {
	err := shim.Start(new(DigitalLibrary))
	if err != nil {
		fmt.Printf("Error: failed to start chaincode: %s\n", err.Error())
		os.Exit(-1)
	}
}
