package post

import "database/sql"

type Repository interface {
	Insert(newPost *Post) error
	SearchPostByID(idPost int) (*Post, error)
	ReturnUserPosts(userId int) ([]Post, error)
	ReturnTenPosts(idSub, numPost int) ([]Post, error)
	ChangeBody(idPost, idUser int, newBody string) error
	Delete(idPost, idUser int) error
}

type repository struct {
	DataBase sql.DB
}

func NewRepository(DB *sql.DB) Repository {
	return &repository{DataBase: *DB}
}
