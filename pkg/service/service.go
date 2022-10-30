package service

import (
	"github.com/nik17s/auth/pkg/repository"
)

type Service struct {
	repository repository.Implementation
}

type AuthImplementation interface {
	CreateUser(login string, password string, email string, phone string) (int, error)
	GenerateJWT(login string, password string) (string, error)
	ParseJWT(tokenString string) (int, error)
}

type Implementation interface {
	AuthImplementation
}

func NewService(repository repository.Implementation) Implementation {
	return &Service{repository: repository}
}
