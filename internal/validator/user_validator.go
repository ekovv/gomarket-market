package validator

import (
	"go-rest-api/internal/model"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	//return validation.ValidateStruct(&user,
	//	validation.Field(
	//		&user.Username,
	//		validation.Required.Error("email is required"),
	//		validation.RuneLength(1, 30).Error("limited max 30 char"),
	//		is.Email.Error("is not valid email format"),
	//	),
	//	validation.Field(
	//		&user.Password,
	//		validation.Required.Error("password is required"),
	//		validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
	//	),
	//)
	return nil
}
