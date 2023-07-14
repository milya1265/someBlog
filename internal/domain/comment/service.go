package comment

type Service interface {
	Create(com *Comment) error
	Delete(idCom int) error
}

type service struct {
	storage Storage
}

func NewService(storage *Storage) Service {
	return &service{storage: *storage}
}

func (s *service) Create(com *Comment) error {
	return s.storage.InsertNewComment(com)
}
func (s *service) Delete(idCom int) error {
	return s.storage.Delete(idCom)
}
