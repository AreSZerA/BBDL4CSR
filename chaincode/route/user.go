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

func CreateUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("function CreateUser requires 3 arguments")
	}
	if args[0] == "" || args[1] == "" || args[2] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	name := args[1]
	passwd := args[2]
	user := lib.User{Email: email, Name: name, Passwd: passwd}
	resp := RetrieveUserByEmail(stub, []string{email})
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

func updateUserField(stub shim.ChaincodeStubInterface, email string, field field, newVal string) peer.Response {
	resp := RetrieveUserByEmail(stub, []string{email})
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
		user.Name = newVal
	case fieldUserPasswd:
		user.Passwd = newVal
	case fieldUserIsReviewer:
		user.IsReviewer, _ = strconv.ParseBool(fmt.Sprintf("%s", newVal))
	case fieldUserIsAdmin:
		user.IsAdmin, _ = strconv.ParseBool(fmt.Sprintf("%s", newVal))
	default:
		return shim.Error("invalid field type")
	}
	payload, err := util.PutLedger(stub, user)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to update ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}

func UpdateUserName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("function UpdateUserName requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	newName := args[1]
	return updateUserField(stub, email, fieldUserName, newName)
}

func UpdateUserPasswd(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("function UpdateUserPasswd requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	newPasswd := args[1]
	return updateUserField(stub, email, fieldUserPasswd, newPasswd)
}

func UpdateUserIsReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("function UpdateUserIsReviewer requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	isReviewer := args[1]
	return updateUserField(stub, email, fieldUserIsReviewer, isReviewer)
}

func UpdateUserIsAdmin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("function UpdateUserIsAdmin requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	isAdmin := args[1]
	return updateUserField(stub, email, fieldUserIsAdmin, isAdmin)
}

func RetrieveUserByKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("function RetrieveUserByKey requires 1 arguments")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	key := args[0]
	payload, err := util.GetByKey(stub, key)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve user by %s: %s", key, err.Error()))
	}
	return shim.Success(payload)
}

func RetrieveUserByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("function RetrieveUserByEmail requires 1 arguments")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	email := args[0]
	payload, err := util.GetByKeys(stub, lib.ObjectTypeUser, email)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve user by %s: %s", email, err.Error()))
	}
	return shim.Success(payload)
}
