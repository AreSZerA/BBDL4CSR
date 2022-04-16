package models

import (
	"ClaytonUniversityLibrary/blockchain"
	"encoding/json"
	"fmt"
)

type User struct {
	Email      string `json:"user_email"`
	Name       string `json:"user_name"`
	Password   string `json:"user_password"`
	IsReviewer bool   `json:"user_is_reviewer"`
	IsAdmin    bool   `json:"user_is_admin"`
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

func UpdateUsername(email string, username string) error {
	_, err := blockchain.Execute(blockchain.FuncUpdateUserName, []byte(email), []byte(username))
	if err != nil {
		return err
	}
	return nil
}

func UpdatePassword(email string, password string) error {
	_, err := blockchain.Execute(blockchain.FuncUpdateUserPassword, []byte(email), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
