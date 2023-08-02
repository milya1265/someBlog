package post

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/domain/post"
)

type Handler interface {
	Create() gin.HandlerFunc
	Get() gin.HandlerFunc
	Edit() gin.HandlerFunc
	GetUserPosts() gin.HandlerFunc
	CreateFeed() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type handler struct {
	Service post.Service
}

func NewHandler(s *post.Service) Handler {
	return &handler{Service: *s}
}
