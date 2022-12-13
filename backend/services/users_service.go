package services

import (
	"user_auth_golang/backend/domain/users"
	"user_auth_golang/backend/utils/errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	//encrypt password

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to encrypt the password")
	}
	user.Password = string(pwSlice[:])
	user.Save()
}
