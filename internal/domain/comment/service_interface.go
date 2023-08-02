package comment

type Service interface {
	Get(idCom int) (*Comment, error)
	GetPostComment(idPost int) ([]Comment, error)
	Create(com *Comment) error
	Edit(idCom int, authorId int, newBody string) error
	Delete(idCom, authorId int) error
}

type service struct {
	repository Repository
}

func NewService(storage *Repository) Service {
	return &service{repository: *storage}
}
