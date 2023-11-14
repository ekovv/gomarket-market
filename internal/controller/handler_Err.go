package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Handler(c echo.Context, err error) error {
	var UnmarshalTypeError *json.UnmarshalTypeError
	if err != nil {
		if errors.Is(err, ErrInvalidData) {
			err := fmt.Sprintf("enabled data %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		if errors.As(err, &UnmarshalTypeError) {
			err := fmt.Sprintf("bad json %s", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		if errors.Is(err, ErrInvalidLogin) {
			err := fmt.Sprintf("bad login and password %s", err)
			return c.JSON(http.StatusUnauthorized, err)
		}
	}
	c.Response().Status = http.StatusOK
	return nil
}
