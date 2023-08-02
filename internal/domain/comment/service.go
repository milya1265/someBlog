package comment

func (s *service) Get(idСom int) (*Comment, error) {
	return s.repository.Get(idСom)
}

func (s *service) Create(com *Comment) (int, error) {
	return s.repository.InsertNewComment(com)
}

func (s *service) Edit(idCom, userId int, newBody string) error {
	return s.repository.ChangeBody(idCom, userId, newBody)
}

func (s *service) Delete(idCom, authorId int) error {
	return s.repository.Delete(idCom, authorId)
}

func (s *service) GetPostComment(idPost int) ([]Comment, error) {
	return s.repository.GetPostComment(idPost)
}
