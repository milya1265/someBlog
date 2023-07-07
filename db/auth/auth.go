package auth

import (
	"context"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"someBlog/pkg"
)

func HashPassword(u *pkg.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

func PasswordEquality(hashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func InsertUser(user *pkg.User, db *sql.DB) (int, error) {
	query := "INSERT INTO users (name, password) VALUES ($1, $2) RETURNING id;"

	row := db.QueryRow(query, user.Name, user.Password)

	var idNewUser int

	if err := row.Scan(&idNewUser); err != nil {
		log.Println("Error with scan row:", err)
		return -1, err
	}

	return idNewUser, nil
}

func FetchUser(user *pkg.User, db *sql.DB) *pkg.User {
	query := "SELECT * FROM users WHERE name = $1;"

	log.Println(user.Name)

	row := db.QueryRow(query, user.Name)

	var uDB pkg.User

	if err := row.Scan(&uDB.Id, &uDB.Name, &uDB.Password); err != nil {
		log.Println(uDB.Name, uDB.Password)
		log.Println("Error with scan row:", err)
		return nil
	}

	if uDB.Name == user.Name && PasswordEquality(uDB.Password, user.Password) {
		return &uDB
	} else {
		return nil
	}
}

func SearchUserByID(idUser int, db *sql.DB) *pkg.User {
	query := "SELECT * FROM users WHERE id = $1;"

	row := db.QueryRowContext(context.Background(), query, idUser)

	var u pkg.User

	if err := row.Scan(&u.Id, &u.Name, &u.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil
	}

	return &u
}

func SearchUserByName(nameUser string, db *sql.DB) *pkg.User {
	query := "SELECT * FROM users WHERE name = $1;"

	row := db.QueryRowContext(context.Background(), query, nameUser)

	var u pkg.User

	if err := row.Scan(&u.Id, &u.Name, &u.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil
	}

	return &u
}
