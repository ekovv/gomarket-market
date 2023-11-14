package repository

import (
	"fmt"
	"go-rest-api/internal/domains"
	"go-rest-api/internal/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domains.IRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByUsername(user model.User, username string) (model.User, error) {
	data := user
	err := ur.db.Where("username=?", username).Find(&data).Error
	if err != nil {
		return model.User{}, fmt.Errorf("could not find user %w", err)
	}
	return data, nil
}
