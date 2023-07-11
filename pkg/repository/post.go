package repository

import (
	"database/sql"
	"someBlog/pkg"
)

func InsertPost(newPost *pkg.Post, db *sql.DB) error {
	query := "INSERT INTO posts (author, time, name, body) VALUES ($1, $2, $3, $4);"

	_, err := db.Exec(query, newPost.Author, newPost.Time, newPost.Body)
	if err != nil {
		return err
	}

	return nil
}

func SearchPostByID(idPost int, db *sql.DB) (*pkg.Post, error) {
	query := "SELECT * FROM posts WHERE id = $1"

	res := db.QueryRow(query, idPost)

	postDB := &pkg.Post{}

	err := res.Scan(&postDB.Id, &postDB.Author, &postDB.Time, &postDB.Body)
	if err != nil {
		return nil, err
	}

	return postDB, nil
}

func ReturnUserPosts(userId int, db *sql.DB) ([]pkg.Post, error) {
	query := "SELECT * FROM posts WHERE author = $1;"

	posts := []pkg.Post{}

	res, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var post pkg.Post
		if err = res.Scan(&post.Id, &post.Author, &post.Time, &post.Body); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return posts, nil

}
