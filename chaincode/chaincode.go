package main

import (
	"chaincode/routes"
	"chaincode/routes/create"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"os"
)

type DigitalLibrary struct {
}

func (library *DigitalLibrary) Init(stub shim.ChaincodeStubInterface) peer.Response {
	_ = create.Admin(stub, []string{"admin@dl4csr.org", "admin", fmt.Sprintf("%x", md5.Sum([]byte("12345678")))})
	return shim.Success(nil)
}

func (library *DigitalLibrary) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	function, ok := routes.FuncMap[funcName]
	if ok {
		return function(stub, args)
	}
	funcNamesBytes, _ := json.Marshal(routes.FuncNames)
	return shim.Error(fmt.Sprintf("Function not implemeted: %s. Available functions: %s.", funcName, string(funcNamesBytes)))
}

func main() {
	err := shim.Start(new(DigitalLibrary))
	if err != nil {
		fmt.Printf("Error: failed to start chaincode: %s\n", err.Error())
		os.Exit(-1)
	}
}
