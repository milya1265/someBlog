package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"someBlog/configs"
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
	Config  configs.Config
}

func NewHandler(s *auth.Service, cfg *configs.Config) Handler {
	return &handler{Service: *s,
		Config: *cfg}
}

func HashPassword(u *user2.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
