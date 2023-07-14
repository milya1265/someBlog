package user

import (
	"database/sql"
	"log"
)

type Storage interface {
	SearchUserByID(idUser int) (*User, error)
	NewSubscribe(idSub int, idProfile int) error
	DeleteSubscribe(idSub int, idProfile int) error
}

type storage struct {
	DataBase sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{DataBase: *db}
}

func (s *storage) SearchUserByID(idUser int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1;"

	row := s.DataBase.QueryRow(query, idUser)

	var u User

	if err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Email, &u.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil, err
	}

	return &u, nil
}

func (s *storage) NewSubscribe(idSub int, idProfile int) error {
	log.Println("New subscribe from ", idSub, "to ", idProfile)

	query := "SELECT * FROM subscription WHERE (subscriber = $1 AND profile = $2)"

	row := s.DataBase.QueryRow(query, idSub, idProfile)

	var val1 int
	var val2 int

	err := row.Scan(&val1, &val2)
	if err != sql.ErrNoRows {
		return err
	}

	query = "INSERT INTO subscription (subscriber, profile) VALUES ($1, $2);"
	_, err = s.DataBase.Exec(query, idSub, idProfile)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) DeleteSubscribe(idSub int, idProfile int) error {
	query := "DELETE FROM subscription WHERE (subscriber = $1 AND profile = $2)"

	_, err := s.DataBase.Exec(query, idSub, idProfile)
	if err != nil {
		return err
	}

	return nil
}
