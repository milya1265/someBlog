package comment

import (
	"errors"
	"log"
)

func (r *repository) Get(idCom int) (*Comment, error) {
	log.Println("INFO   search comment in database")

	query := "SELECT * FROM post_comments WHERE id = $1;"

	row := r.DataBase.QueryRow(query, idCom)

	var Com Comment
	if err := row.Scan(&Com.Id, &Com.IdPost, &Com.Author, &Com.Time, &Com.Body); err != nil {
		return &Com, err
	}

	return &Com, nil
}

func (r *repository) InsertNewComment(com *Comment) error {
	log.Println("INFO   insert comment in database")

	query := "INSERT INTO post_comments (id_post, author, time, body) VALUES ($1, $2, $3, $4)"

	if _, err := r.DataBase.Exec(query, com.IdPost, com.Author, com.Time, com.Body); err != nil {
		return err
	}

	return nil
}

func (r *repository) ChangeBody(idCom int, authorId int, newBody string) error {
	log.Println("INFO   change comment in database")

	query := "UPDATE post_comments SET body = $1 WHERE id = $2 AND author = $3;"

	res, err := r.DataBase.Exec(query, newBody, idCom, authorId)
	if err != nil {
		return err
	}

	if cnt, err := res.RowsAffected(); cnt == 0 {
		if err != nil {
			return err
		}
		err := errors.New("you can't change foreign comment")
		return err
	}

	return nil
}

func (r *repository) Delete(idCom, authirId int) error {
	log.Println("INFO   delete comment in database")

	query := "DELETE FROM post_comments	WHERE id = $1 AND author = $2;"

	res, err := r.DataBase.Exec(query, idCom, authirId)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		err := errors.New("you can't delete this comment")
		return err
	}

	return nil
}

func (r *repository) GetPostComment(idPost int) ([]Comment, error) {
	log.Println("INFO   take post comments from database")

	query := "SELECT * FROM post_comments WHERE id_post = $1 ORDER BY time;"

	res, err := r.DataBase.Query(query, idPost)
	if err != nil {
		return nil, err
	}

	var comments []Comment

	defer res.Close()

	for res.Next() {
		var com Comment
		if err = res.Scan(&com.Id, &com.IdPost, &com.Author, &com.Time, &com.Body); err != nil {
			return nil, err
		}
		comments = append(comments, com)
	}

	return comments, nil
}
