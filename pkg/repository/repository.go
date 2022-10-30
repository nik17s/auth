package repository

import entity "github.com/nik17s/auth"

type AuthImplementation interface {
	CreateUser(login, password, email, phone string) (int, error)
	GetUser(login, password string) (*entity.User, error)
}

type Implementation interface {
	AuthImplementation
}
