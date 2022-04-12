package lib

type User struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Passwd     string `json:"passwd"`
	IsAdmin    bool   `json:"is_admin"`
	IsReviewer bool   `json:"is_reviewer"`
}

func (u User) Type() string {
	return ObjectTypeUser
}

func (u User) Keys() []string {
	return []string{u.Email}
}
