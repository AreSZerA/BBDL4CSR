package utils

import (
	"chaincode/lib"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func PutLedger(stub shim.ChaincodeStubInterface, object lib.BlockchainObject) ([]byte, error) {
	key, err := stub.CreateCompositeKey(object.Type(), object.Keys())
	if err != nil {
		return nil, err
	}
	objectBytes, _ := json.Marshal(object)
	err = stub.PutState(key, objectBytes)
	if err != nil {
		return nil, err
	}
	return objectBytes, nil
}

func DelLedger(stub shim.ChaincodeStubInterface, object lib.BlockchainObject) error {
	key, err := stub.CreateCompositeKey(object.Type(), object.Keys())
	if err != nil {
		return err
	}
	err = stub.DelState(key)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(stub shim.ChaincodeStubInterface, objectType string, keys ...string) ([][]byte, error) {
	var results [][]byte
	itr, err := stub.GetStateByPartialCompositeKey(objectType, keys)
	if err != nil {
		return nil, err
	}
	defer itr.Close()
	for itr.HasNext() {
		result, err := itr.Next()
		if err != nil {
			return nil, err
		}
		results = append(results, result.GetValue())
	}
	return results, nil
}

func GetByKeys(stub shim.ChaincodeStubInterface, objectType string, keys ...string) ([]byte, error) {
	key, err := stub.CreateCompositeKey(objectType, keys)
	if err != nil {
		return nil, err
	}
	result, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetByQuery(stub shim.ChaincodeStubInterface, query string) ([][]byte, error) {
	var results [][]byte
	itr, err := stub.GetQueryResult(query)
	if err != nil {
		return nil, err
	}
	defer itr.Close()
	for itr.HasNext() {
		result, err := itr.Next()
		if err != nil {
			return nil, err
		}
		results = append(results, result.GetValue())
	}
	return results, nil
}
