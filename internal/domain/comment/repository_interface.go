package comment

import (
	"database/sql"
)

type Repository interface {
	Get(idCom int) (*Comment, error)
	GetPostComment(idPost int) ([]Comment, error)
	InsertNewComment(com *Comment) error
	ChangeBody(idCom int, authorId int, newBody string) error
	Delete(idCom, authorID int) error
}

type repository struct {
	DataBase sql.DB
}

func NewRepository(DB *sql.DB) Repository {
	return &repository{DataBase: *DB}
}
