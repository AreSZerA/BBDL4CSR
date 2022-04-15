package main

import (
	"chaincode/routes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"os"
)

var funcMap = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	"ping": func(shim.ChaincodeStubInterface, []string) peer.Response { return shim.Success([]byte("pong")) },

	"CreateUser":                routes.CreateUser,
	"CreateReviewer":            routes.CreateReviewer,
	"CreateAdmin":               routes.CreateAdmin,
	"UpdateUserName":            routes.UpdateUserName,
	"UpdateUserPasswd":          routes.UpdateUserPasswd,
	"UpdateUserIsReviewer":      routes.UpdateUserIsReviewer,
	"UpdateUserIsAdmin":         routes.UpdateUserIsAdmin,
	"RetrieveUserByEmail":       routes.RetrieveUserByEmail,
	"RetrieveAllUsers":          routes.RetrieveAllUsers,
	"RetrieveAllReviewers":      routes.RetrieveAllReviewers,
	"RetrieveCountAllUsers":     routes.RetrieveCountAllUsers,
	"RetrieveCountAllReviewers": routes.RetrieveCountAllReviewers,

	"CreatePaper":             routes.CreatePaper,
	"RetrieveAllPapers":       routes.RetrieveAllPapers,
	"RetrieveAcceptedPapers":  routes.RetrieveAcceptedPapers,
	"RetrieveRejectedPapers":  routes.RetrieveRejectedPapers,
	"RetrieveReviewingPapers": routes.RetrieveReviewingPapers,
	"RetrievePapersByEmail":   routes.RetrievePapersByEmail,
	"RetrievePapersByTitle":   routes.RetrievePapersByTitle,
	"RetrievePaperById":       routes.RetrievePaperById,
	"UpdatePaperStatus":       routes.UpdatePaperStatus,

	"UpdatePeerReview":             routes.UpdatePeerReview,
	"RetrievePeerReviewsByPaperId": routes.RetrievePeerReviewsByPaperId,
	"RetrievePeerReviewByIds":      routes.RetrievePeerReviewByIds,
}

var funcNames = []string{
	"ping",

	"CreateUser",
	"CreateReviewer",
	"CreateAdmin",
	"UpdateUserName",
	"UpdateUserPasswd",
	"UpdateUserIsReviewer",
	"UpdateUserIsAdmin",
	"RetrieveUserByEmail",
	"RetrieveAllUsers",
	"RetrieveAllReviewers",
	"RetrieveCountAllUsers",
	"RetrieveCountAllReviewers",

	"CreatePaper",
	"RetrieveAllPapers",
	"RetrieveAcceptedPapers",
	"RetrieveRejectedPapers",
	"RetrieveReviewingPapers",
	"RetrievePapersByEmail",
	"RetrievePapersByTitle",
	"RetrievePaperById",
	"UpdatePaperStatus",

	"UpdatePeerReview",
	"RetrievePeerReviewsByPaperId",
	"RetrievePeerReviewByIds",
}

type DigitalLibrary struct {
}

func (library *DigitalLibrary) Init(stub shim.ChaincodeStubInterface) peer.Response {
	resp := routes.CreateAdmin(stub, []string{"admin@dl4csr.org", "admin", fmt.Sprintf("%x", md5.Sum([]byte("12345678")))})
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
