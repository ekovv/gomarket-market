package service

import (
	"errors"
	"go-rest-api/internal/controller"
	"go-rest-api/internal/domains/mocks"
	"go-rest-api/internal/model"
	"testing"
)

type sessionMock[A any] func(c *mocks.SessionService, args A)
type repositoryMock[A any] func(c *mocks.IRepository, args A)

func Test_userService_SignUp(t *testing.T) {
	invalidErr := errors.New("invalid")
	tests := []struct {
		name           string
		args           model.User
		sessionMock    sessionMock[model.User]
		repositoryMock repositoryMock[model.User]
		wantErr        error
	}{
		{
			name: "OK1",
			args: model.User{
				Username: "dima",
				Password: "test1",
				Session:  "ahsjufil12-fk",
			},
			sessionMock: func(c *mocks.SessionService, user model.User) {
				c.Mock.On("Generate").Return(user.Session, nil).Times(1)
			},
			repositoryMock: func(c *mocks.IRepository, user model.User) {
				c.Mock.On("CreateUser", user).Return(nil)
			},
			wantErr: nil,
		},
		{
			name: "BAD1",
			args: model.User{
				Username: "dima",
				Password: "test1",
				Session:  "ahsjufil12-fk",
			},
			sessionMock: func(c *mocks.SessionService, user model.User) {
				c.Mock.On("Generate").Return(user.Session, nil).Times(1)
			},
			repositoryMock: func(c *mocks.IRepository, user model.User) {
				c.Mock.On("CreateUser", user).Return(invalidErr).Times(1)
			},
			wantErr: invalidErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := mocks.NewIRepository(t)
			session := mocks.NewSessionService(t)
			tt.repositoryMock(storage, tt.args)
			tt.sessionMock(session, tt.args)
			service := Service{
				database:       storage,
				sessionService: session,
			}
			cook, err := service.SignUp(tt.args)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
			if tt.wantErr == nil && cook != tt.args.Session {
				t.Errorf("got %s, want %s", cook, tt.args.Session)
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	tests := []struct {
		name           string
		args           model.User
		repositoryMock repositoryMock[model.User]
		wantErr        error
	}{
		{
			name: "OK1",
			args: model.User{
				Username: "dima",
				Password: "test1",
			},
			repositoryMock: func(c *mocks.IRepository, user model.User) {
				data := model.User{}
				c.Mock.On("GetUserByUsername", data, user.Username).Return(model.User{Username: "dima", Password: "test1", Session: "ahsjdfurol-12"}, nil)
			},
			wantErr: nil,
		},
		{
			name: "BAD1",
			args: model.User{
				Username: "1",
				Password: "1",
			},
			repositoryMock: func(c *mocks.IRepository, user model.User) {
				data := model.User{}
				c.Mock.On("GetUserByUsername", data, user.Username).Return(model.User{}, controller.ErrInvalidLogin).Times(1)
			},
			wantErr: controller.ErrInvalidLogin,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := mocks.NewIRepository(t)
			tt.repositoryMock(storage, tt.args)
			service := Service{
				database: storage,
			}
			err := service.Login(tt.args)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}
		})
	}
}
