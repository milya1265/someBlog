package comment

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/domain/comment"
)

type Handler interface {
	CreateNewComment() gin.HandlerFunc
}

type handler struct {
	Service comment.Service
}

func NewHandler(s *comment.Service) Handler {
	return &handler{Service: *s}
}
