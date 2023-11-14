package domains

import "go-rest-api/internal/model"

//go:generate go run github.com/vektra/mockery/v3 --name=Service
type Service interface {
	SignUp(user model.User) (string, error)
	Login(user model.User) error
}
