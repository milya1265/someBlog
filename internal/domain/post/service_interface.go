package post

type Service interface {
	Create(newPost *Post) error
	CreateFeed(userId, numTenPost int) ([]Post, error)
	GetByID(idPost int) (*Post, error)
	GetUserPosts(userId int) ([]Post, error)
	Edit(idPost, idUser int, newBody string) error
	Delete(idPost, idUser int) error
}

type service struct {
	repository Repository
}

func NewService(storage *Repository) Service {
	return &service{repository: *storage}
}
