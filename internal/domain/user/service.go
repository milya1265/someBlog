package user

type Service interface {
	GetUser(idUser int) (*User, error)
	Subscribe(idSub int, idProfile int) error
	Unsubscribe(idSub int, idProfile int) error
}

type service struct {
	storage Storage
}

func NewService(storage *Storage) Service {
	return &service{storage: *storage}
}

func (s *service) GetUser(idUser int) (*User, error) {
	return s.storage.SearchUserByID(idUser)
}

func (s *service) Subscribe(idSub, idProfile int) error {
	return s.storage.NewSubscribe(idSub, idProfile)
}

func (s *service) Unsubscribe(idSub int, idProfile int) error {
	return s.storage.DeleteSubscribe(idSub, idProfile)
}
