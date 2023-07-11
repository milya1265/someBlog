package repository

import (
	"database/sql"
	"log"
	"someBlog/pkg"
)

func ReturnTenPosts(idSub int, numPost int, db *sql.DB) ([]pkg.Post, error) {

	query := "SELECT p.id, p.author, p.time, p.body FROM posts p JOIN subscription s ON p.author = s.profile WHERE s.subscriber = $1 LIMIT $2 OFFSET $3"

	posts := []pkg.Post{}

	res, err := db.Query(query, idSub, 10, numPost) // Хреново, что тут хардкодом 10 постов
	if err != nil {
		return posts, err
	}

	defer res.Close()

	for res.Next() {
		log.Println("kek")
		var post pkg.Post
		if err = res.Scan(&post.Id, &post.Author, &post.Time, &post.Body); err != nil {
			return nil, err
		}
		log.Println(post)
		posts = append(posts, post)
	}

	if err = res.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}
