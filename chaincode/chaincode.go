package main

import (
	"chaincode/routes/create"
	"chaincode/routes/retrieve"
	"chaincode/routes/update"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"os"
)

var funcNames []string

func init() {
	for key := range funcMap {
		funcNames = append(funcNames, key)
	}
}

var funcMap = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	"ping": func(shim.ChaincodeStubInterface, []string) peer.Response { return shim.Success([]byte("pong")) },
	// Create
	"CreateUser":     create.User,
	"CreateReviewer": create.Reviewer,
	"CreateAdmin":    create.Admin,
	"CreatePaper":    create.Paper,
	// Update
	"UpdateUserByName":                   update.UserByEmail,
	"UpdateUserName":                     update.UserName,
	"UpdateUserPassword":                 update.UserPassword,
	"UpdateUserIsReviewer":               update.UserIsReviewer,
	"UpdateUserIsNotReviewer":            update.UserIsNotReviewer,
	"UpdateUserIsAdmin":                  update.UserIsAdmin,
	"UpdateUserIsNotAdmin":               update.UserIsNotAdmin,
	"UpdateUserReviewing":                update.UserReviewing,
	"UpdatePaperById":                    update.PaperById,
	"UpdatePeerReviewByPaperAndReviewer": update.PeerReviewByPaperAndReviewer,
	// Retrieve
	"RetrieveUsers":                   retrieve.Users,
	"RetrieveUsersByName":             retrieve.UsersByName,
	"RetrieveUsersIsReviewer":         retrieve.UsersIsReviewer,
	"RetrieveUsersIsAdmin":            retrieve.UsersIsAdmin,
	"RetrieveUsersByPaperId":          retrieve.UsersByPaperId,
	"RetrieveUserByEmail":             retrieve.UserByEmail,
	"RetrievePapers":                  retrieve.Papers,
	"RetrieveAcceptedPapersByTitle":   retrieve.AcceptedPapersByTitle,
	"RetrieveAcceptedPapersByAuthor":  retrieve.AcceptedPapersByAuthor,
	"RetrieveAcceptedPapersByKeyword": retrieve.AcceptedPapersByKeyword,
	"RetrieveAcceptedPapers":          retrieve.AcceptedPapers,
	"RejectedPapers":                  retrieve.RejectedPapers,
	"RejectedPapersByAuthor":          retrieve.RejectedPapersByAuthor,
	"ReviewingPapers":                 retrieve.ReviewingPapers,
	"ReviewingPapersByAuthor":         retrieve.ReviewingPapersByAuthor,
	"RetrievePaperById":               retrieve.PaperById,
	"RetrievePeerReviewsByReviewer":   retrieve.PeerReviewsByReviewer,
	// Delete
}

type DigitalLibrary struct {
}

func (library *DigitalLibrary) Init(stub shim.ChaincodeStubInterface) peer.Response {
	resp := create.Admin(stub, []string{"admin@dl4csr.org", "admin", fmt.Sprintf("%x", md5.Sum([]byte("12345678")))})
	if resp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to create admin: %s", resp.Message))
	}
	return shim.Success(nil)
}

func (library *DigitalLibrary) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	function, ok := funcMap[funcName]
	if ok {
		return function(stub, args)
	}
	funcNamesBytes, _ := json.Marshal(funcNames)
	return shim.Error(fmt.Sprintf("Function not implemeted: %s. Available functions: %s.", funcName, string(funcNamesBytes)))
}

func main() {
	err := shim.Start(new(DigitalLibrary))
	if err != nil {
		fmt.Printf("Error: failed to start chaincode: %s\n", err.Error())
		os.Exit(-1)
	}
}
