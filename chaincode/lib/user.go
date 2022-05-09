// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the structure of user.

package lib

// User implements BlockchainObjectInterface for storing in the ledger.
type User struct {
	// Unique email address of the user.
	Email string `json:"user_email"`
	// Username.
	Name string `json:"user_name"`
	// User password.
	Password string `json:"user_password"`
	// IsReviewer is used for distinguishing ordinary users with reviewers.
	IsReviewer bool `json:"user_is_reviewer"`
	// IsAdmin is used for distinguishing ordinary users with administrators.
	IsAdmin bool `json:"user_is_admin"`
	// The number of ongoing peer review tasks of the user.
	Reviewing uint16 `json:"user_reviewing"`
}

// Type returns ObjectTypeUser.
func (u User) Type() string {
	return ObjectTypeUser
}

// Attributes returns a slice constructed with Email.
func (u User) Attributes() []string {
	return []string{u.Email}
}
