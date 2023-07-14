package user

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/domain/user"
)

type Handler interface {
	GetUser() gin.HandlerFunc
	Subscribe() gin.HandlerFunc
	Unsubscribe() gin.HandlerFunc
}

type handler struct {
	Service user.Service
}

func NewHandler(s *user.Service) Handler {
	return &handler{Service: *s}
}
