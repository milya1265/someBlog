package repository

import (
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
	query := "INSERT INTO users (name, surname, email, password) VALUES ($1, $2, $3, $4) RETURNING id;"

	row := db.QueryRow(query, user.Name, user.Surname, user.Email, user.Password)

	var idNewUser int

	if err := row.Scan(&idNewUser); err != nil {
		log.Println("Error with scan row:", err)
		return -1, err
	}

	return idNewUser, nil
}

func CheckPasswordAndReturnUser(user *pkg.User, db *sql.DB) *pkg.User {
	query := "SELECT * FROM users WHERE email = $1;"

	row := db.QueryRow(query, user.Email)

	var uDB pkg.User

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
