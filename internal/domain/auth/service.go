package auth

import (
	"someBlog/internal/domain/user"
)

func (s *service) SignUp(user *user.User) (int, error) {
	return s.repository.InsertUser(user)
}

func (s *service) SignIn(user *user.User) *user.User {
	return s.repository.CheckPasswordAndReturnUser(user)
}

func (s *service) GetUser(idUser int) (*user.User, error) {
	return s.repository.SearchUserByID(idUser)
}
