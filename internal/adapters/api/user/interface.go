package user

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/domain/user"
)

type Handler interface {
	Get() gin.HandlerFunc
	Subscribe() gin.HandlerFunc
	Unsubscribe() gin.HandlerFunc
	EditProfile() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type handler struct {
	Service user.Service
}

func NewHandler(s *user.Service) Handler {
	return &handler{Service: *s}
}
