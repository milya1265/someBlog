package postgreSQL

import (
	"database/sql"
	"someBlog/configs"
)

type Repository struct {
	DataBase *sql.DB
}

func (repo *Repository) Open(cfg configs.Config) (*sql.DB, error) {
	var err error = nil

	repo.DataBase, err = sql.Open(cfg.Storage.DbDriver, cfg.Storage.DbDriver+"://"+cfg.Storage.Username+":"+
		""+cfg.Storage.Password+"@"+cfg.Storage.Host+":"+cfg.Storage.Port+"/"+cfg.Storage.Database+""+
		"?sslmode="+cfg.Storage.SSLMode)
	//repo.DataBase, err = sql.Open("postgres", "postgres://dmilya:qwerty@localhost:5432/BlogDB?sslmode=disable")

	if err != nil {
		return nil, err
	}

	if err = repo.DataBase.Ping(); err != nil {
		return nil, err
	}

	return repo.DataBase, nil
}
