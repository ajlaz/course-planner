package models

import (
	"errors"
)

var (
	ErrIncorrectPassword = errors.New("Incorrect Password")
	ErrUserNotFound      = errors.New("User Not Found")
	ErrUserAlreadyExists = errors.New("User Already Exists")
	ErrInvalidDate       = errors.New("Invalid Date")
)
