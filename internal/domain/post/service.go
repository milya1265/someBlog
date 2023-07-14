package post

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	Create(newPost *Post) error
	CreateFeed(userId, numTenPost int) ([]Post, error)
	GetByID(idPost int) (*Post, error)
	GetUserPosts(userId int) ([]Post, error)
	Delete(ctx gin.Context) error
}

type service struct {
	storage Storage
}

func NewService(storage *Storage) Service {
	return &service{storage: *storage}
}

func (s *service) Create(newPost *Post) error {
	return s.storage.Insert(newPost)
}

func (s *service) GetByID(idPost int) (*Post, error) {
	return s.storage.SearchPostByID(idPost)
}

func (s *service) GetUserPosts(userId int) ([]Post, error) {
	return s.storage.ReturnUserPosts(userId)
}

func (s *service) Delete(ctx gin.Context) error {
	return nil
}

func (s *service) CreateFeed(idSub, numTenPost int) ([]Post, error) {
	return s.storage.ReturnTenPosts(idSub, numTenPost)
}
