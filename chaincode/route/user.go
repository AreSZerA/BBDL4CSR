package route

import (
	"chaincode/lib"
	"chaincode/util"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

func CreateUser(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	if len(args) != 3 {
		return shim.Error("function CreateUser requires 3 arguments")
	}
	email := args[0]
	name := args[1]
	passwd := args[2]
	user := lib.User{Email: email, Name: name, Passwd: passwd}
	resp := RetrieveUserByEmail(stub, email)
	if resp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to check data uniqueness: %s", resp.Message))
	}
	if resp.Payload != nil {
		return shim.Error(fmt.Sprintf("email %s has been registered", email))
	}
	payload, err := util.PutLedger(stub, user)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to create user: %s", err.Error()))
	}
	return shim.Success(payload)
}

func updateUserField(stub shim.ChaincodeStubInterface, email string, field field, newVal interface{}) peer.Response {
	resp := RetrieveUserByEmail(stub, email)
	if resp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to check email %s: %s", email, resp.Message))
	}
	if resp.Payload == nil {
		return shim.Error(fmt.Sprintf("email %s has not been registered", email))
	}
	var user lib.User
	err := json.Unmarshal(resp.Payload, &user)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to unserialise user: %s", err.Error()))
	}
	switch field {
	case fieldUserName:
		switch newVal.(type) {
		case string:
			user.Name = fmt.Sprintf("%s", newVal)
		default:
			return shim.Error("invalid value type: type must be string")
		}
	case fieldUserPasswd:
		switch newVal.(type) {
		case string:
			user.Passwd = fmt.Sprintf("%s", newVal)
		default:
			return shim.Error("invalid value type: type must be string")
		}
	case fieldUserIsReviewer:
		switch newVal.(type) {
		case bool:
			user.IsReviewer, _ = strconv.ParseBool(fmt.Sprintf("%s", newVal))
		default:
			return shim.Error("invalid value type: type must be bool")
		}
	case fieldUserIsAdmin:
		switch newVal.(type) {
		case bool:
			user.IsAdmin, _ = strconv.ParseBool(fmt.Sprintf("%s", newVal))
		default:
			return shim.Error("invalid value type: type must be bool")
		}
	default:
		return shim.Error("invalid field type")
	}
	payload, err := util.PutLedger(stub, user)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to update ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}

func UpdateUserName(stub shim.ChaincodeStubInterface, email string, newName string) peer.Response {
	return updateUserField(stub, email, fieldUserName, newName)
}

func UpdateUserPasswd(stub shim.ChaincodeStubInterface, email string, newPasswd string) peer.Response {
	return updateUserField(stub, email, fieldUserPasswd, newPasswd)
}

func UpdateUserIsReviewer(stub shim.ChaincodeStubInterface, email string, isReviewer bool) peer.Response {
	return updateUserField(stub, email, fieldUserIsReviewer, isReviewer)
}

func UpdateUserIsAdmin(stub shim.ChaincodeStubInterface, email string, isAdmin bool) peer.Response {
	return updateUserField(stub, email, fieldUserIsAdmin, isAdmin)
}

func RetrieveUserByKey(stub shim.ChaincodeStubInterface, key string) peer.Response {
	payload, err := util.GetByKey(stub, key)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve user by %s: %s", key, err.Error()))
	}
	return shim.Success(payload)
}

func RetrieveUserByEmail(stub shim.ChaincodeStubInterface, email string) peer.Response {
	payload, err := util.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve user by %s: %s", email, err.Error()))
	}
	return shim.Success(payload)
}
