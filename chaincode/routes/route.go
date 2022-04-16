package routes

import (
	"chaincode/routes/create"
	"chaincode/routes/retrieve"
	"chaincode/routes/update"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var FuncNames []string

func init() {
	for key := range FuncMap {
		FuncNames = append(FuncNames, key)
	}
}

var FuncMap = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	"ping": func(shim.ChaincodeStubInterface, []string) peer.Response { return shim.Success([]byte("pong")) },
	// Create
	"CreateUser":     create.User,
	"CreateReviewer": create.Reviewer,
	"CreateAdmin":    create.Admin,
	"CreatePaper":    create.Paper,
	// Update
	"UpdateUserByEmail":                  update.UserByEmail,
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
	"RetrieveAcceptedPapers":          retrieve.AcceptedPapers,
	"RetrieveAcceptedPapersByTitle":   retrieve.AcceptedPapersByTitle,
	"RetrieveAcceptedPapersByAuthor":  retrieve.AcceptedPapersByAuthor,
	"RetrieveAcceptedPapersByKeyword": retrieve.AcceptedPapersByKeyword,
	"RetrieveRejectedPapers":          retrieve.RejectedPapers,
	"RetrieveRejectedPapersByAuthor":  retrieve.RejectedPapersByAuthor,
	"RetrieveReviewingPapers":         retrieve.ReviewingPapers,
	"RetrieveReviewingPapersByAuthor": retrieve.ReviewingPapersByAuthor,
	"RetrievePaperById":               retrieve.PaperById,
	"RetrievePeerReviewsByReviewer":   retrieve.PeerReviewsByReviewer,
	// Delete
}
