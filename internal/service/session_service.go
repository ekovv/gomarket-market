package service

import "github.com/google/uuid"

type SessionService struct {
}

func NewSessionService() SessionService {
	return SessionService{}
}

func (s *SessionService) Generate() (string, error) {
	session := uuid.New().String()
	return session, nil
}
