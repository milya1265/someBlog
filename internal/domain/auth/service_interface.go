package auth

import (
	user "someBlog/internal/domain/user"
)

type Service interface {
	SignUp(user *user.User) (int, error)
	SignIn(user *user.User) *user.User
	GetUser(idUser int) (*user.User, error)
}

type service struct {
	repository Repository
}

func NewService(storage *Repository) Service {
	return &service{repository: *storage}
}
