package routes

type field uint8

const (
	fieldUserName = iota
	fieldUserPasswd
	fieldUserIsReviewer
	fieldUserIsAdmin
	fieldUserReviewing
)
