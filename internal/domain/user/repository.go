package user

import (
	"database/sql"
	"log"
)

func (r *repository) SearchUserByID(idUser int) (*User, error) {
	log.Println("INFO   search user by ID in database")

	query := "SELECT * FROM users WHERE id = $1;"

	row := r.DataBase.QueryRow(query, idUser)

	var u User

	if err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Email, &u.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil, err
	}

	return &u, nil
}

func (r *repository) NewSubscribe(idSub int, idProfile int) error {
	log.Println("INFO   new subscribe from ", idSub, "to ", idProfile)

	query := "SELECT * FROM subscription WHERE (subscriber = $1 AND profile = $2)"

	row := r.DataBase.QueryRow(query, idSub, idProfile)

	var val1 int
	var val2 int

	err := row.Scan(&val1, &val2)
	if err != sql.ErrNoRows {
		return err
	}

	query = "INSERT INTO subscription (subscriber, profile) VALUES ($1, $2);"
	_, err = r.DataBase.Exec(query, idSub, idProfile)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteSubscribe(idSub int, idProfile int) error {
	log.Println("INFO   delete  subscribe from database")

	query := "DELETE FROM subscription WHERE (subscriber = $1 AND profile = $2)"

	_, err := r.DataBase.Exec(query, idSub, idProfile)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ChangeUser(idUser int, newName, newSurname string) error {
	log.Println("INFO   change user in database")

	query := "UPDATE users SET name = $1, surname = $2 WHERE id = $3"

	_, err := r.DataBase.Exec(query, newName, newSurname, idUser)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(idUser int) error {
	log.Println("INFO   delete user from database")

	query := "DELETE FROM users WHERE id = $1"

	_, err := r.DataBase.Exec(query, idUser)
	if err != nil {
		return err
	}

	return nil
}
