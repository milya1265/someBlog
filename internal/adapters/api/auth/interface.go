package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"someBlog/internal/domain/auth"
	user2 "someBlog/internal/domain/user"
)

type Handler interface {
	SignUp() gin.HandlerFunc
	SignIn() gin.HandlerFunc
	Authorize() gin.HandlerFunc
}

type handler struct {
	Service auth.Service
}

func NewHandler(s *auth.Service) Handler {
	return &handler{Service: *s}
}

func HashPassword(u *user2.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
