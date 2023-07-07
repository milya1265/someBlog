package blog

import (
	"database/sql"
	"someBlog/pkg"
)

func InsertNewComment(com *pkg.Comment, db *sql.DB) error {
	query := "INSERT INTO post_comments (id_post, author, time, body) VALUES ($1, $2, $3, $4)"

	_, err := db.Exec(query, com.PostID, com.Author, com.Time, com.Body)
	if err != nil {
		return err
	}

	return nil
}
