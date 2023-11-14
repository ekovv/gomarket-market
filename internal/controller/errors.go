package controller

import "errors"

var (
	ErrInvalidData  = errors.New("invalid data")
	ErrInvalidLogin = errors.New("invalid data for login")
)
