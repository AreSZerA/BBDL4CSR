package models

import (
	"ClaytonUniversityLibrary/blockchain"
	"encoding/json"
	"fmt"
)

type User struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Passwd     string `json:"passwd"`
	IsAdmin    bool   `json:"is_admin"`
	IsReviewer bool   `json:"is_reviewer"`
	Reviewing  uint16 `json:"reviewing"`
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
