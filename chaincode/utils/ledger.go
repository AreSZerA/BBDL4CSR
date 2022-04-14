package utils

import (
	"chaincode/lib"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func PutLedger(stub shim.ChaincodeStubInterface, object lib.BlockchainObject) ([]byte, error) {
	key, err := stub.CreateCompositeKey(object.Type(), object.Keys())
	if err != nil {
		return nil, fmt.Errorf("failed to create composite ke %s", err.Error())
	}
	objectBytes, err := json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("failed to serialise object to JSON: %s", err.Error())
	}
	err = stub.PutState(key, objectBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to update ledger: %s", err.Error())
	}
	return objectBytes, nil
}

func GetAll(stub shim.ChaincodeStubInterface, objectType string, keys ...string) ([][]byte, error) {
	var results [][]byte
	itr, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, fmt.Errorf("failed to get from ledger: %s", err.Error())
	}
	defer itr.Close()
	for itr.HasNext() {
		result, err := itr.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to traverse results: %s", err.Error())
		}
		results = append(results, result.GetValue())
	}
	return results, nil
}

func GetByKeys(stub shim.ChaincodeStubInterface, objectType string, keys ...string) ([]byte, error) {
	key, err := stub.CreateCompositeKey(objectType, keys)
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %s", err.Error())
	}
	result, err := stub.GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get from ledger: %s", err.Error())
	}
	return result, nil
}

func GetByQuery(stub shim.ChaincodeStubInterface, query string) ([][]byte, error) {
	var results [][]byte
	itr, err := stub.GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %s", err.Error())
	}
	defer itr.Close()
	for itr.HasNext() {
		result, err := itr.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to traverse results: %s", err.Error())
		}
		results = append(results, result.GetValue())
	}
	return results, nil
}
