package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	user2 "someBlog/internal/domain/user"
)

func (r *repository) InsertUser(user *user2.User) (int, error) {
	log.Println("INFO --> insert user in database")

	query := "INSERT INTO users (name, surname, email, password) VALUES ($1, $2, $3, $4) RETURNING id;"

	row := r.DataBase.QueryRow(query, user.Name, user.Surname, user.Email, user.Password)

	var idNewUser int

	if err := row.Scan(&idNewUser); err != nil {
		log.Println("ERROR --> scan row:", err)
		return -1, err
	}

	return idNewUser, nil
}

func (r *repository) CheckPasswordAndReturnUser(user *user2.User) *user2.User {
	log.Println("INFO --> check password and return user from database")

	query := "SELECT * FROM users WHERE email = $1;"

	row := r.DataBase.QueryRow(query, user.Email)

	var uDB user2.User

	if err := row.Scan(&uDB.Id, &uDB.Name, &uDB.Surname, &uDB.Email, &uDB.Password); err != nil {
		log.Println("ERROR --> scan row:", err)
		return nil
	}

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

func (r *repository) SearchUserByID(idUser int) (*user2.User, error) {
	log.Println("INFO --> search user by ID in database")

	query := "SELECT * FROM users WHERE id = $1;"

	row := r.DataBase.QueryRow(query, idUser)

	var u user2.User

	if err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Email, &u.Password); err != nil {
		log.Println("ERROR --> scan row:", err)
		return nil, err
	}

	return &u, nil
}
