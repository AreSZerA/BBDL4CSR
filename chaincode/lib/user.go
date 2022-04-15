package lib

type User struct {
	Email      string `json:"user_email"`
	Name       string `json:"user_name"`
	Passwd     string `json:"user_password"`
	IsAdmin    bool   `json:"user_is_admin"`
	IsReviewer bool   `json:"user_is_reviewer"`
	Reviewing  uint16 `json:"user_reviewing"`
}

func (u User) Type() string {
	return ObjectTypeUser
}

func (u User) Keys() []string {
	return []string{u.Email}
}
