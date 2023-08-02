package post

func (s *service) Create(newPost *Post) error {
	return s.repository.Insert(newPost)
}

func (s *service) GetByID(idPost int) (*Post, error) {
	return s.repository.SearchPostByID(idPost)
}

func (s *service) GetUserPosts(userId int) ([]Post, error) {
	return s.repository.ReturnUserPosts(userId)
}

func (s *service) Edit(idPost, idUser int, newBody string) error {
	return s.repository.ChangeBody(idPost, idUser, newBody)
}

func (s *service) CreateFeed(idSub, numTenPost int) ([]Post, error) {
	return s.repository.ReturnTenPosts(idSub, numTenPost)
}

func (s *service) Delete(idPost, idUser int) error {
	return s.repository.Delete(idPost, idUser)
}
