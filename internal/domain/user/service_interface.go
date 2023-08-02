package user

type Service interface {
	Get(idUser int) (*User, error)
	Subscribe(idSub int, idProfile int) error
	Unsubscribe(idSub int, idProfile int) error
	EditUser(idUser int, newName, newSurname string) error
	Delete(idUser int) error
}

type service struct {
	repository Repository
}

func NewService(storage *Repository) Service {
	return &service{repository: *storage}
}
