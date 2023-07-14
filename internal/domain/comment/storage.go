package comment

import "database/sql"

type Storage interface {
	InsertNewComment(com *Comment) error
	Delete(idCom int) error
}

type storage struct {
	DataBase sql.DB
}

func NewStorage(DB *sql.DB) Storage {
	return &storage{DataBase: *DB}
}

func (s *storage) InsertNewComment(com *Comment) error {
	query := "INSERT INTO post_comments (id_post, author, time, body) VALUES ($1, $2, $3, $4)"

	_, err := s.DataBase.Exec(query, com.IdPost, com.Author, com.Time, com.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) Delete(idCom int) error {
	return nil
}
