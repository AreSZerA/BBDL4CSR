package blockchain

const (
	FuncPing = "ping"

	FuncCreateUser                = "CreateUser"
	FuncCreateReviewer            = "CreateReviewer"
	FuncUpdateUserName            = "UpdateUserName"
	FuncUpdateUserPasswd          = "UpdateUserPasswd"
	FuncUpdateUserIsReviewer      = "UpdateUserIsReviewer"
	FuncUpdateUserIsAdmin         = "UpdateUserIsAdmin"
	FuncRetrieveUserByEmail       = "RetrieveUserByEmail"
	FuncRetrieveAllUsers          = "RetrieveAllUsers"
	FuncRetrieveAllReviewers      = "RetrieveAllReviewers"
	FuncRetrieveCountAllUsers     = "RetrieveCountAllUsers"
	FuncRetrieveCountAllReviewers = "RetrieveCountAllReviewers"

	FuncCreatePaper             = "CreatePaper"
	FuncRetrieveAllPapers       = "RetrieveAllPapers"
	FuncRetrieveAcceptedPapers  = "RetrieveAcceptedPapers"
	FuncRetrieveRejectedPapers  = "RetrieveRejectedPapers"
	FuncRetrieveReviewingPapers = "RetrieveReviewingPapers"
	FuncRetrievePapersByEmail   = "RetrievePapersByEmail"
	FuncRetrievePapersByTitle   = "RetrievePapersByTitle"
	FuncRetrievePaperById       = "RetrievePaperById"
	FuncUpdatePaperStatus       = "UpdatePaperStatus"

	FuncUpdatePeerReview             = "UpdatePeerReview"
	FuncRetrievePeerReviewsByPaperId = "RetrievePeerReviewsByPaperId"
	FuncRetrievePeerReviewByIds      = "RetrievePeerReviewByIds"
)
