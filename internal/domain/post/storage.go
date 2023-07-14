package post

import (
	"database/sql"
	"log"
)

type Storage interface {
	Insert(newPost *Post) error
	SearchPostByID(idPost int) (*Post, error)
	ReturnUserPosts(userId int) ([]Post, error)
	ReturnTenPosts(idSub, numPost int) ([]Post, error)
}

type storage struct {
	DataBase sql.DB
}

func NewStorage(DB *sql.DB) Storage {
	return &storage{DataBase: *DB}
}

func (s *storage) Insert(newPost *Post) error {
	query := "INSERT INTO posts (author, time, body) VALUES ($1, $2, $3);"

	_, err := s.DataBase.Exec(query, newPost.Author, newPost.Time, newPost.Body)
	if err != nil {
		return err
	}

	return nil
}
func (s *storage) SearchPostByID(idPost int) (*Post, error) {
	query := "SELECT * FROM posts WHERE id = $1"

	res := s.DataBase.QueryRow(query, idPost)

	postDB := &Post{}

	err := res.Scan(&postDB.Id, &postDB.Author, &postDB.Time, &postDB.Body)
	if err != nil {
		return nil, err
	}

	return postDB, nil
}
func (s *storage) ReturnUserPosts(userId int) ([]Post, error) {
	log.Println("INFO   return user posts")
	log.Println(userId)
	query := "SELECT * FROM posts WHERE author = $1;"

	posts := []Post{}

	res, err := s.DataBase.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var post Post
		if err = res.Scan(&post.Id, &post.Author, &post.Time, &post.Body); err != nil {
			return nil, err
		}
		posts = append(posts, post)
		log.Println(post)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *storage) ReturnTenPosts(idSub int, numPost int) ([]Post, error) {

	query := "SELECT p.id, p.author, p.time, p.body FROM posts p JOIN subscription s ON p.author = s.profile WHERE s.subscriber = $1 LIMIT $2 OFFSET $3"

	posts := []Post{}

	res, err := s.DataBase.Query(query, idSub, 10, numPost) // Хреново, что тут хардкодом 10 постов
	if err != nil {
		return posts, err
	}

	defer res.Close()

	for res.Next() {
		log.Println("kek")
		var post Post
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
