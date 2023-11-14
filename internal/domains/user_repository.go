package domains

import "go-rest-api/internal/model"

//go:generate go run github.com/vektra/mockery/v3 --name=IRepository
type IRepository interface {
	CreateUser(user model.User) error
	GetUserByUsername(user model.User, username string) (model.User, error)
}
