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

	// Create user
	"CreateUser":     create.User,
	"CreateReviewer": create.Reviewer,
	"CreateAdmin":    create.Admin,
	// Create paper
	"CreatePaper": create.Paper,

	// Update user
	"UpdateUserByEmail":       update.UserByEmail,
	"UpdateUserName":          update.UserName,
	"UpdateUserPassword":      update.UserPassword,
	"UpdateUserIsReviewer":    update.UserIsReviewer,
	"UpdateUserIsNotReviewer": update.UserIsNotReviewer,
	"UpdateUserIsAdmin":       update.UserIsAdmin,
	"UpdateUserIsNotAdmin":    update.UserIsNotAdmin,
	"UpdateUserReviewing":     update.UserReviewing,
	// Update paper
	"UpdatePaperById": update.PaperById,
	// Update peer review
	"UpdatePeerReviewByPaperAndReviewer": update.PeerReviewByPaperAndReviewer,

	// Retrieve user
	"RetrieveUsersByQuery":                  retrieve.UsersByQuery,
	"RetrieveUsers":                         retrieve.Users,
	"RetrieveUsersSortByEmail":              retrieve.UsersSortByEmail,
	"RetrieveUsersSortByName":               retrieve.UsersSortByName,
	"RetrieveUsersByNameSortByEmail":        retrieve.UsersByNameSortByEmail,
	"RetrieveReviewersSortByEmail":          retrieve.ReviewersSortByEmail,
	"RetrieveReviewersSortByName":           retrieve.ReviewersSortByName,
	"RetrieveReviewersByPaperIdSortByEmail": retrieve.ReviewersByPaperIdSortByEmail,
	"RetrieveReviewersByPaperIdSortByName":  retrieve.ReviewersByPaperIdSortByName,
	"RetrieveAdminsSortByEmail":             retrieve.AdminsSortByEmail,
	"RetrieveAdminsSortByName":              retrieve.AdminsSortByName,
	"RetrieveUserByEmail":                   retrieve.UserByEmail,
	// Retrieve paper
	"RetrievePapersByQuery":                             retrieve.PapersByQuery,
	"RetrievePapers":                                    retrieve.Papers,
	"RetrievePapersSortByTitle":                         retrieve.PapersSortByTitle,
	"RetrievePapersSortByUploadTime":                    retrieve.PapersSortByUploadTime,
	"RetrieveAcceptedPapersSortByTitle":                 retrieve.AcceptedPapersSortByTitle,
	"RetrieveAcceptedPapersSortByUploadTime":            retrieve.AcceptedPapersSortByUploadTime,
	"RetrieveAcceptedPapersSortByPublishTime":           retrieve.AcceptedPapersSortByPublishTime,
	"RetrieveRejectedPapersSortByTitle":                 retrieve.RejectedPapersSortByTitle,
	"RetrieveRejectedPapersSortByUploadTime":            retrieve.RejectedPapersSortByUploadTime,
	"RetrieveReviewingPapersSortByTitle":                retrieve.ReviewingPapersSortByTitle,
	"RetrieveReviewingPapersSortByUploadTime":           retrieve.ReviewingPapersSortByUploadTime,
	"RetrieveAcceptedPapersByTitleSortByTitle":          retrieve.AcceptedPapersByTitleSortByTitle,
	"RetrieveAcceptedPapersByTitleSortByPublishTime":    retrieve.AcceptedPapersByTitleSortByPublishTime,
	"RetrieveAcceptedPapersByAuthorSortByTitle":         retrieve.AcceptedPapersByAuthorSortByTitle,
	"RetrieveAcceptedPapersByAuthorSortByPublishTime":   retrieve.AcceptedPapersByAuthorSortByPublishTime,
	"RetrieveAcceptedPapersByKeywordSortByTitle":        retrieve.AcceptedPapersByKeywordSortByTitle,
	"RetrieveAcceptedPapersByKeywordSortByPublishTime":  retrieve.AcceptedPapersByKeywordSortByPublishTime,
	"RetrieveAcceptedPapersByUploaderSortByTitle":       retrieve.AcceptedPapersByUploaderSortByTitle,
	"RetrieveAcceptedPapersByUploaderSortByUploadTime":  retrieve.AcceptedPapersByUploaderSortByUploadTime,
	"RetrieveAcceptedPapersByUploaderSortByPublishTime": retrieve.AcceptedPapersByUploaderSortByPublishTime,
	"RetrieveRejectedPapersByUploaderSortByTitle":       retrieve.RejectedPapersByUploaderSortByTitle,
	"RetrieveRejectedPapersByUploaderSortByUploadTime":  retrieve.RejectedPapersByUploaderSortByUploadTime,
	"RetrieveRejectedPapersByUploaderSortByPublishTime": retrieve.RejectedPapersByUploaderSortByPublishTime,
	"RetrieveReviewingPapersByUploaderSortByTitle":      retrieve.ReviewingPapersByUploaderSortByTitle,
	"RetrieveReviewingPapersByUploaderSortByUploadTime": retrieve.ReviewingPapersByUploaderSortByUploadTime,
	"RetrievePaperById":                                 retrieve.PaperById,
	// Retrieve peer review
	"RetrievePeerReviewsByQuery":                    retrieve.PeerReviewsByQuery,
	"RetrievePeerReviewsByReviewerSortByCreateTime": retrieve.PeerReviewsByReviewerSortByCreateTime,

	// Delete
}
