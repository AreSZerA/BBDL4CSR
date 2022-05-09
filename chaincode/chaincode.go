// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the structure implements Chaincode interface for starting the blockchain.

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

// DigitalLibrary implements shim.Chaincode.
type DigitalLibrary struct {
}

// Init will be invoked for chaincode container to initialize.
func (library *DigitalLibrary) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// create an account with administrator privilege
	// the Init method will be invoked when upgrading chaincode, so error is ignored here
	_ = create.Admin(stub, []string{"admin@dl4csr.org", "admin", fmt.Sprintf("%x", md5.Sum([]byte("12345678")))})
	return shim.Success(nil)
}

// Invoke will be called by clients to query or update the ledger with function name and parameters.
func (library *DigitalLibrary) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// get the function name and arguments from stub
	funcName, args := stub.GetFunctionAndParameters()
	// retrieve function from map
	function, ok := routes.FuncMap[funcName]
	// if the function is found in the map, run the function to response
	if ok {
		return function(stub, args)
	}
	// otherwise, deserialize the list of function names as the response message
	funcNamesBytes, _ := json.Marshal(routes.FuncNames)
	return shim.Error(fmt.Sprintf("Function not implemeted: %s. Available functions: %s.", funcName, string(funcNamesBytes)))
}

// The entrance function for starting chaincode.
func main() {
	// run the smart contract
	err := shim.Start(new(DigitalLibrary))
	if err != nil {
		fmt.Printf("Error: failed to start chaincode: %s\n", err.Error())
		os.Exit(-1)
	}
}
