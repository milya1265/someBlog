package repository

import "database/sql"

type Repository struct {
	DataBase *sql.DB
}

func (repo *Repository) Open(pathDB string) (*sql.DB, error) {
	var err error = nil
	repo.DataBase, err = sql.Open("postgres", pathDB)

	if err != nil {
		return nil, err
	}

	if err = repo.DataBase.Ping(); err != nil {
		return nil, err
	}

	return repo.DataBase, nil
}
