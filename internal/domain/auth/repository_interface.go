package auth

import (
	"database/sql"
	user2 "someBlog/internal/domain/user"
)

type Repository interface {
	CheckPasswordAndReturnUser(user *user2.User) *user2.User
	InsertUser(user *user2.User) (int, error)
	SearchUserByID(idUser int) (*user2.User, error)
}

type repository struct {
	DataBase *sql.DB
}

func NewStorage(DB *sql.DB) Repository {
	return &repository{DataBase: DB}
}
