// Copyright 2022 AreSZerA. All rights reserved.
// This file defines utility function related to the ledger.

package utils

import (
	"chaincode/lib"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// PutLedger puts an object implements lib.BlockchainObjectInterface to ledger.
// Identified by the composite key, if the object already exist, the fields of it will be updated.
func PutLedger(stub shim.ChaincodeStubInterface, object lib.BlockchainObjectInterface) ([]byte, error) {
	// create composite key
	key, err := stub.CreateCompositeKey(object.Type(), object.Attributes())
	if err != nil {
		return nil, err
	}
	// serialise object
	objectBytes, _ := json.Marshal(object)
	// write ledger
	err = stub.PutState(key, objectBytes)
	if err != nil {
		return nil, err
	}
	return objectBytes, nil
}

// DelLedger deletes object implements lib.BlockchainObjectInterface to ledger.
func DelLedger(stub shim.ChaincodeStubInterface, object lib.BlockchainObjectInterface) error {
	// create composite key
	key, err := stub.CreateCompositeKey(object.Type(), object.Attributes())
	if err != nil {
		return err
	}
	// write ledger
	err = stub.DelState(key)
	if err != nil {
		return err
	}
	return nil
}

// GetAll retrieves object according to the object type and attributes.
// This function is based on shim.ChaincodeStubInterface#GetStateByPartialCompositeKey,
// mainly used to get all the records of a type of objects.
func GetAll(stub shim.ChaincodeStubInterface, objectType string, attributes ...string) ([][]byte, error) {
	var results [][]byte
	// retrieve by partial composite key
	itr, err := stub.GetStateByPartialCompositeKey(objectType, attributes)
	if err != nil {
		return nil, err
	}
	defer itr.Close()
	// traverse the results to append to the slice
	for itr.HasNext() {
		result, err := itr.Next()
		if err != nil {
			return nil, err
		}
		results = append(results, result.GetValue())
	}
	return results, nil
}

// GetByKeys retrieves objects by object type name and complete composite.
// It is used for searching the exact object.
func GetByKeys(stub shim.ChaincodeStubInterface, objectType string, attributes ...string) ([]byte, error) {
	key, err := stub.CreateCompositeKey(objectType, attributes)
	if err != nil {
		return nil, err
	}
	result, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByQuery uses CouchDB to retrieve objects by query.
func GetByQuery(stub shim.ChaincodeStubInterface, query string) ([][]byte, error) {
	var results [][]byte
	// retrieve by query
	itr, err := stub.GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer itr.Close()
	// traverse the results to append to the slice
	for itr.HasNext() {
		result, err := itr.Next()
		if err != nil {
			return nil, err
		}
		results = append(results, result.GetValue())
	}
	return results, nil
}
