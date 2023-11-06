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
}

type userController struct {
	userService domains.UserUseCase
}

func NewUserController(userService domains.UserUseCase) IUserController {
	return &userController{userService}
}

func (s *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	uuid, err := s.userService.SignUp(user)
	if err != nil {
		// сделай логику обработки ошибки как у артема
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := http.Cookie{}
	cookie.Name = "token"
	cookie.Value = uuid
	cookie.Expires = time.Now().Add(56 * time.Hour)
	c.SetCookie(&cookie)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Response().Status = http.StatusOK
	return nil
}

//func (uc *userController) LogIn(c echo.Context) error {
//	user := model.User{}
//	if err := c.Bind(&user); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	tokenString, err := uc.uu.Login(user)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//	cookie := new(http.Cookie)
//	cookie.Name = "token"
//	cookie.Value = tokenString
//	cookie.Expires = time.Now().Add(24 * time.Hour)
//	cookie.Path = "/"
//	cookie.Domain = os.Getenv("API_DOMAIN")
//	cookie.Secure = true
//	cookie.HttpOnly = true
//	cookie.SameSite = http.SameSiteNoneMode
//	c.SetCookie(cookie)
//	return c.NoContent(http.StatusOK)
//}
//
//func (uc *userController) LogOut(c echo.Context) error {
//	cookie := new(http.Cookie)
//	cookie.Name = "token"
//	cookie.Value = ""
//	cookie.Expires = time.Now()
//	cookie.Path = "/"
//	cookie.Domain = os.Getenv("API_DOMAIN")
//	cookie.Secure = true
//	cookie.HttpOnly = true
//	cookie.SameSite = http.SameSiteNoneMode
//	c.SetCookie(cookie)
//	return c.NoContent(http.StatusOK)
//}
