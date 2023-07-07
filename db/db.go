package db

import "database/sql"

type DB struct {
	DataBase *sql.DB
}

func (db *DB) Open(pathDB string) error {
	var err error = nil
	db.DataBase, err = sql.Open("postgres", pathDB)

	if err != nil {
		return err
	}

	if err = db.DataBase.Ping(); err != nil {
		return err
	}

	return nil
}
