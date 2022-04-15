package models

import (
	"ClaytonUniversityLibrary/blockchain"
	"encoding/json"
	"fmt"
)

type User struct {
	Email      string `json:"user_email"`
	Name       string `json:"user_name"`
	Passwd     string `json:"user_password"`
	IsAdmin    bool   `json:"user_is_admin"`
	IsReviewer bool   `json:"user_is_reviewer"`
	Reviewing  uint16 `json:"user_reviewing"`
}

func FindUserByEmail(email string) (User, error) {
	resp, err := blockchain.Query(blockchain.FuncRetrieveUserByEmail, []byte(email))
	if err != nil {
		return User{}, err
	}
	var user User
	if resp.Payload == nil {
		return User{}, fmt.Errorf("user not exist")
	}
	err = json.Unmarshal(resp.Payload, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func RegisterUser(email string, username string, passwd string) error {
	_, err := blockchain.Execute(blockchain.FuncCreateUser, []byte(email), []byte(username), []byte(passwd))
	if err != nil {
		return err
	}
	return nil
}
