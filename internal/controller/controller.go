package controller

import (
	"github.com/labstack/echo/v4"
	"go-rest-api/internal/domains"
	"go-rest-api/internal/model"
	"net/http"
	"time"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
}

type userController struct {
	userService domains.Service
}

func NewUserController(userService domains.Service) IUserController {
	return &userController{userService}
}

func (s *userController) SignUp(c echo.Context) error {
	data := model.User{}
	if err := c.Bind(&data); err != nil {
		return Handler(c, err)
	}

	uuid, err := s.userService.SignUp(data)
	if err != nil {
		return Handler(c, err)
	}

	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = uuid
	cookie.Expires = time.Now().Add(56 * time.Hour)
	c.SetCookie(&cookie)
	c.Response().Status = http.StatusOK
	return nil
}

func (s *userController) Login(c echo.Context) error {
	data := model.User{}
	if err := c.Bind(&data); err != nil {
		return Handler(c, err)
	}

	err := s.userService.Login(data)
	if err != nil {
		return Handler(c, err)
	}

	c.Response().Status = http.StatusOK
	return nil
}
