package main

import (
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
	case "test":
		return shim.Success([]byte("chaincode invocation test passed"))
	default:
		return shim.Error(fmt.Sprintf("Cannot handle args: %s", args))
	}
}

func main() {
	err := shim.Start(new(DigitalLibrary))
	if err != nil {
		fmt.Printf("Error: failed to start chaincode: %s\n", err.Error())
		os.Exit(-1)
	}
}
