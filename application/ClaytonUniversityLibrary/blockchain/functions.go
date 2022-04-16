package blockchain

const (
	FuncPing = "ping"

	FuncCreateUser     = "CreateUser"
	FuncCreateReviewer = "CreateReviewer"
	FuncCreateAdmin    = "CreateAdmin"
	FuncCreatePaper    = "CreatePaper"

	FuncUpdateUserByEmail                  = "UpdateUserByEmail"
	FuncUpdateUserName                     = "UpdateUserName"
	FuncUpdateUserPassword                 = "UpdateUserPassword"
	FuncUpdateUserIsReviewer               = "UpdateUserIsReviewer"
	FuncUpdateUserIsNotReviewer            = "UpdateUserIsNotReviewer"
	FuncUpdateUserIsAdmin                  = "UpdateUserIsAdmin"
	FuncUpdateUserIsNotAdmin               = "UpdateUserIsNotAdmin"
	FuncUpdateUserReviewing                = "UpdateUserReviewing"
	FuncUpdatePaperById                    = "UpdatePaperById"
	FuncUpdatePeerReviewByPaperAndReviewer = "UpdatePeerReviewByPaperAndReviewer"

	FuncRetrieveUsers                   = "RetrieveUsers"
	FuncRetrieveUsersByName             = "RetrieveUsersByName"
	FuncRetrieveUsersIsReviewer         = "RetrieveUsersIsReviewer"
	FuncRetrieveUsersIsAdmin            = "RetrieveUsersIsAdmin"
	FuncRetrieveUsersByPaperId          = "RetrieveUsersByPaperId"
	FuncRetrieveUserByEmail             = "RetrieveUserByEmail"
	FuncRetrievePapers                  = "RetrievePapers"
	FuncRetrieveAcceptedPapersByTitle   = "RetrieveAcceptedPapersByTitle"
	FuncRetrieveAcceptedPapersByAuthor  = "RetrieveAcceptedPapersByAuthor"
	FuncRetrieveAcceptedPapersByKeyword = "RetrieveAcceptedPapersByKeyword"
	FuncRetrieveAcceptedPapers          = "RetrieveAcceptedPapers"
	FuncRetrieveRejectedPapers          = "RetrieveRejectedPapers"
	FuncRetrieveRejectedPapersByAuthor  = "RetrieveRejectedPapersByAuthor"
	FuncRetrieveReviewingPapers         = "RetrieveReviewingPapers"
	FuncRetrieveReviewingPapersByAuthor = "RetrieveReviewingPapersByAuthor"
	FuncRetrievePaperById               = "RetrievePaperById"
	FuncRetrievePeerReviewsByReviewer   = "RetrievePeerReviewsByReviewer"
)
