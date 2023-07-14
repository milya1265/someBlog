package post

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/domain/post"
)

type Handler interface {
	CreateNewPost() gin.HandlerFunc
	GetPost() gin.HandlerFunc
	GetUserPosts() gin.HandlerFunc
	CreateFeed() gin.HandlerFunc
}

type handler struct {
	Service post.Service
}

func NewHandler(s *post.Service) Handler {
	return &handler{Service: *s}
}
