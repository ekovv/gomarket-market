package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-rest-api/internal/domains/mocks"
	"go-rest-api/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type serviceMock func(c *mocks.Service)

func Test_UserController_SignUp(t *testing.T) {
	type jso struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	tests := []struct {
		name        string
		body        jso
		serviceMock serviceMock
		wantCode    int
	}{
		{
			name: "OK1",
			body: jso{
				Login:    "dima",
				Password: "test1",
			},
			serviceMock: func(c *mocks.Service) {
				c.Mock.On("SignUp", model.User{Username: "dima", Password: "test1", Session: ""}).Return("ahsjufil12-fk", nil).Times(1)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "BAD",
			body: jso{
				Login:    "1",
				Password: "1",
			},
			serviceMock: func(c *mocks.Service) {
				user := model.User{Username: "1", Password: "1", Session: ""}
				c.Mock.On("SignUp", user).Return("ahsjufil12-fk", ErrInvalidData).Times(1)
			},
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			service := mocks.NewService(t)
			h := NewUserController(service)
			tt.serviceMock(service)
			b, err := json.Marshal(tt.body)
			if err != nil {
				return
			}
			path := "/v1/user/register"
			e.POST(path, h.SignUp)

			w := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, path, strings.NewReader(string(b)))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			// создаём новый Recorder

			e.ServeHTTP(w, request)

			if w.Code != tt.wantCode {
				t.Errorf("got %d, want %d", w.Code, tt.wantCode)
			}
		})
	}
}

func Test_userController_Login(t *testing.T) {
	type jso struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	tests := []struct {
		name        string
		body        jso
		serviceMock serviceMock
		wantCode    int
	}{
		{
			name: "OK1",
			body: jso{
				Login:    "dima",
				Password: "test1",
			},
			serviceMock: func(c *mocks.Service) {
				c.Mock.On("Login", model.User{Username: "dima", Password: "test1", Session: ""}).Return(nil).Times(1)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "BAD",
			body: jso{
				Login:    "1",
				Password: "1",
			},
			serviceMock: func(c *mocks.Service) {
				user := model.User{Username: "1", Password: "1", Session: ""}
				c.Mock.On("Login", user).Return(ErrInvalidLogin).Times(1)
			},
			wantCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			service := mocks.NewService(t)
			h := NewUserController(service)
			tt.serviceMock(service)
			b, err := json.Marshal(tt.body)
			if err != nil {
				return
			}
			path := "/v1/user/login"
			e.POST(path, h.Login)

			w := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, path, strings.NewReader(string(b)))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			// создаём новый Recorder

			e.ServeHTTP(w, request)

			if w.Code != tt.wantCode {
				t.Errorf("got %d, want %d", w.Code, tt.wantCode)
			}
		})
	}
}
