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

func createUserBase(stub shim.ChaincodeStubInterface, isReviewer bool, args []string) peer.Response {
	if len(args) < 3 {
		if isReviewer {
			return shim.Error("function CreateReviewer requires 3 arguments")
		}
		return shim.Error("function CreateUser requires 3 arguments")
	}
	if args[0] == "" || args[1] == "" || args[2] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	name := args[1]
	passwd := args[2]
	user := lib.User{Email: email, Name: name, Passwd: passwd, IsReviewer: isReviewer}
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

func CreateUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return createUserBase(stub, false, args)
}

func CreateReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return createUserBase(stub, true, args)
}

func updateUserBase(stub shim.ChaincodeStubInterface, email string, field field, newVal string) peer.Response {
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
	case fieldUserReviewing:
		if newVal == "+" {
			user.Reviewing++
		} else if newVal == "-" {
			user.Reviewing--
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

func UpdateUserName(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function UpdateUserName requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	newName := args[1]
	return updateUserBase(stub, email, fieldUserName, newName)
}

func UpdateUserPasswd(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function UpdateUserPasswd requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	newPasswd := args[1]
	return updateUserBase(stub, email, fieldUserPasswd, newPasswd)
}

func UpdateUserIsReviewer(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function UpdateUserIsReviewer requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	isReviewer := args[1]
	return updateUserBase(stub, email, fieldUserIsReviewer, isReviewer)
}

func UpdateUserIsAdmin(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function UpdateUserIsAdmin requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	isAdmin := args[1]
	return updateUserBase(stub, email, fieldUserIsAdmin, isAdmin)
}

func updateUserReviewingAdd(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		shim.Error("function updateUserReviewingAdd requires 2 arguments")
	}
	if args[0] == "" {
		shim.Error("argument should be nonempty")
	}
	email := args[0]
	return updateUserBase(stub, email, fieldUserReviewing, "+")
}

func updateUserReviewingSub(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		shim.Error("function updateUserReviewingAdd requires 2 arguments")
	}
	if args[0] == "" {
		shim.Error("argument should be nonempty")
	}
	email := args[0]
	return updateUserBase(stub, email, fieldUserReviewing, "-")
}

func RetrieveUserByKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
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
	if len(args) < 1 {
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

func RetrieveAllUsers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	var users []lib.User
	results, err := util.GetAll(stub, lib.ObjectTypeUser)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve users: %s", err.Error()))
	}
	for _, userBytes := range results {
		var user lib.User
		err = json.Unmarshal(userBytes, &user)
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to unserialise user: %s", err.Error()))
		}
		users = append(users, user)
	}
	usersBytes, err := json.Marshal(users)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise users: %s", err.Error()))
	}
	return shim.Success(usersBytes)
}

func RetrieveAllReviewers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	var users []lib.User
	results, err := util.GetByQuery(stub, `{"selector":{"is_reviewer":true}}`)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve users: %s", err.Error()))
	}
	for _, userBytes := range results {
		var user lib.User
		err = json.Unmarshal(userBytes, &user)
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to unserialise user: %s", err.Error()))
		}
		users = append(users, user)
	}
	usersBytes, err := json.Marshal(users)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise users: %s", err.Error()))
	}
	return shim.Success(usersBytes)
}

func retrieveAllSortedReviewers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	var users []lib.User
	results, err := util.GetByQuery(stub, `{"selector":{"is_reviewer":true},"sort":[{"reviewing":"asc"}]}`)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve users: %s", err.Error()))
	}
	for _, userBytes := range results {
		var user lib.User
		err = json.Unmarshal(userBytes, &user)
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to unserialise user: %s", err.Error()))
		}
		users = append(users, user)
	}
	usersBytes, err := json.Marshal(users)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise users: %s", err.Error()))
	}
	return shim.Success(usersBytes)
}

func RetrieveCountAllUsers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	results, err := util.GetAll(stub, lib.ObjectTypeUser)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrive ledger: %s", err.Error()))
	}
	return shim.Success([]byte(strconv.Itoa(len(results))))
}

func RetrieveCountAllReviewers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"is_reviewer":true}}`
	results, err := util.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrive ledger: %s", err.Error()))
	}
	return shim.Success([]byte(strconv.Itoa(len(results))))
}
