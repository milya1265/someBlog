package auth

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	user2 "someBlog/internal/domain/user"
)

type Storage interface {
	CheckPasswordAndReturnUser(user *user2.User) *user2.User
	InsertUser(user *user2.User) (int, error)
	SearchUserByID(idUser int) (*user2.User, error)
}

type storage struct {
	DataBase *sql.DB
}

func NewStorage(DB *sql.DB) Storage {
	return &storage{DataBase: DB}
}

func (s *storage) InsertUser(user *user2.User) (int, error) {
	query := "INSERT INTO users (name, surname, email, password) VALUES ($1, $2, $3, $4) RETURNING id;"

	row := s.DataBase.QueryRow(query, user.Name, user.Surname, user.Email, user.Password)

	var idNewUser int

	if err := row.Scan(&idNewUser); err != nil {
		log.Println("Error with scan row:", err)
		return -1, err
	}

	return idNewUser, nil
}

func (s *storage) CheckPasswordAndReturnUser(user *user2.User) *user2.User {
	query := "SELECT * FROM users WHERE email = $1;"

	row := s.DataBase.QueryRow(query, user.Email)

	var uDB user2.User

	if err := row.Scan(&uDB.Id, &uDB.Name, &uDB.Surname, &uDB.Email, &uDB.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil
	}
	log.Println(uDB.Email, uDB.Password)

	if uDB.Email == user.Email && PasswordEquality(uDB.Password, user.Password) {
		return &uDB
	} else {
		return nil
	}
}

func PasswordEquality(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (s *storage) SearchUserByID(idUser int) (*user2.User, error) {
	query := "SELECT * FROM users WHERE id = $1;"

	row := s.DataBase.QueryRow(query, idUser)

	var u user2.User

	if err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Email, &u.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil, err
	}

	return &u, nil
}
