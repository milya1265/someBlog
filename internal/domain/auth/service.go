package auth

import (
	"someBlog/internal/domain/user"
)

type Service interface {
	SignUp(user *user.User) (int, error)
	SignIn(user *user.User) *user.User
	GetUser(idUser int) (*user.User, error)
}

type service struct {
	storage Storage
}

func NewService(storage *Storage) Service {
	return &service{storage: *storage}
}

func (s *service) SignUp(user *user.User) (int, error) {
	return s.storage.InsertUser(user)
}

func (s *service) SignIn(user *user.User) *user.User {
	return s.storage.CheckPasswordAndReturnUser(user)
}

func (s *service) GetUser(idUser int) (*user.User, error) {
	return s.storage.SearchUserByID(idUser)
}
