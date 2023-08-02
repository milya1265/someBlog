package comment

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/domain/comment"
)

type Handler interface {
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
	GetPostComment() gin.HandlerFunc
	Edit() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type handler struct {
	Service comment.Service
}

func NewHandler(s *comment.Service) Handler {
	return &handler{Service: *s}
}
