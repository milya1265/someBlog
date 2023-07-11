package repository

import (
	"database/sql"
	"log"
	"someBlog/pkg"
)

func SearchUserByID(idUser int, db *sql.DB) (*pkg.User, error) {
	query := "SELECT * FROM users WHERE id = $1;"

	row := db.QueryRow(query, idUser)

	var u pkg.User

	if err := row.Scan(&u.Id, &u.Name, &u.Surname, &u.Email, &u.Password); err != nil {
		log.Println("Error with scan row:", err)
		return nil, err
	}

	return &u, nil
}

//	INSERT INTO the_table (id, column_1, column_2)
//	VALUES (1, 'A', 'X'), (2, 'B', 'Y'), (3, 'C', 'Z')
//	ON CONFLICT (id) DO UPDATE
//		SET column_1 = excluded.column_1,
//			column_2 = excluded.column_2;

//func NewSubscribe(idSub int, idProfile int, db *sql.DB) error {
//	query := "INSERT INTO subscription (subscriber, profile) VALUES ($1, $2) WHERE NOT EXIST(SELECT hiu FROM subscription WHERE subscriber = $1 AND profile = $2)	;"
//	_, err := db.Exec(query, idSub, idProfile)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func NewSubscribe(idSub int, idProfile int, db *sql.DB) error {
	log.Println("New subscribe from ", idSub, "to ", idProfile)

	query := "SELECT * FROM subscription WHERE (subscriber = $1 AND profile = $2)"

	row := db.QueryRow(query, idSub, idProfile)

	var val1 int
	var val2 int

	err := row.Scan(&val1, &val2)
	if err != sql.ErrNoRows {
		return err
	}

	query = "INSERT INTO subscription (subscriber, profile) VALUES ($1, $2);"
	_, err = db.Exec(query, idSub, idProfile)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubscribe(idSub int, idProfile int, db *sql.DB) error {
	query := "DELETE FROM subscription WHERE (subscriber = $1 AND profile = $2)"

	_, err := db.Exec(query, idSub, idProfile)
	if err != nil {
		return err
	}

	return nil
}
