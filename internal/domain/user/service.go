package user

func (s *service) Get(idUser int) (*User, error) {
	return s.repository.SearchUserByID(idUser)
}

func (s *service) Subscribe(idSub, idProfile int) error {
	return s.repository.NewSubscribe(idSub, idProfile)
}

func (s *service) Unsubscribe(idSub int, idProfile int) error {
	return s.repository.DeleteSubscribe(idSub, idProfile)
}

func (s *service) EditUser(idUser int, newName, newSurname string) error {
	return s.repository.ChangeUser(idUser, newName, newSurname)
}
func (s *service) Delete(idUser int) error {
	return s.repository.Delete(idUser)
}
