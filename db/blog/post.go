package blog

import (
	"database/sql"
	errors2 "errors"
	"log"
	db2 "someBlog/db/auth"
	"someBlog/pkg"
)

func InsertPost(newPost *pkg.Post, db *sql.DB) error {
	query := "INSERT INTO posts (author, time, name, body) VALUES ($1, $2, $3, $4);"

	_, err := db.Exec(query, newPost.Author, newPost.Time, newPost.Name, newPost.Body)
	if err != nil {
		return err
	}

	return nil
}

func FetchPostByID(idPost int, db *sql.DB) (*pkg.Post, error) {
	query := "SELECT * FROM posts WHERE id = $1"

	res := db.QueryRow(query, idPost)

	postDB := &pkg.Post{}

	err := res.Scan(&postDB.Id, &postDB.Author, &postDB.Time, &postDB.Name, &postDB.Body)
	if err != nil {
		return nil, err
	}

	return postDB, nil
}

func FetchUserPosts(username string, db *sql.DB) (*[]pkg.Post, error) {
	query := "SELECT * FROM posts WHERE author = $1;"

	user := db2.SearchUserByName(username, db)
	if user == nil {
		return nil, errors2.New("User is not found")
	}

	res, err := db.Query(query, user.Id)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	posts := &[]pkg.Post{}

	for res.Next() {
		log.Println("lolkek")
		var post pkg.Post
		if err = res.Scan(&post.Id, &post.Author, &post.Time, &post.Name, &post.Body); err != nil {
			return nil, err
		}
		*posts = append(*posts, post)
		log.Println(post)
		log.Println(posts)
	}
	if err = res.Err(); err != nil {
		return nil, err
	}
	return posts, nil

}
