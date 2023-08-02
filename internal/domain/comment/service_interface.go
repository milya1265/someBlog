package comment

//go:generate mockgen -source=service_interface.go -destination=mocks/mock.go

type Service interface {
	Get(idCom int) (*Comment, error)
	GetPostComment(idPost int) ([]Comment, error)
	Create(com *Comment) (int, error)
	Edit(idCom int, authorId int, newBody string) error
	Delete(idCom, authorId int) error
}

type service struct {
	repository Repository
}

func NewService(storage *Repository) Service {
	return &service{repository: *storage}
}
