package user

import "database/sql"

type Repository interface {
	SearchUserByID(idUser int) (*User, error)
	NewSubscribe(idSub int, idProfile int) error
	DeleteSubscribe(idSub int, idProfile int) error
	ChangeUser(idUser int, newName, newSurname string) error
	Delete(idUser int) error
}

type repository struct {
	DataBase sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{DataBase: *db}
}
