package repository

import "database/sql"

type Repository struct {
	DataBase *sql.DB
}

func (repo *Repository) Open(pathDB string) error {
	var err error = nil
	repo.DataBase, err = sql.Open("postgres", pathDB)

	if err != nil {
		return err
	}

	if err = repo.DataBase.Ping(); err != nil {
		return err
	}

	return nil
}
