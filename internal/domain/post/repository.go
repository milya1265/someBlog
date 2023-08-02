package post

import (
	"errors"
	"log"
)

func (r *repository) Insert(newPost *Post) error {
	log.Println("INFO --> insert post in database")

	query := "INSERT INTO posts (author, time, body) VALUES ($1, $2, $3);"

	_, err := r.DataBase.Exec(query, newPost.Author, newPost.Time, newPost.Body)
	if err != nil {
		return err
	}

	return nil
}
func (r *repository) SearchPostByID(idPost int) (*Post, error) {
	log.Println("INFO --> searching posts by post id")

	query := "SELECT * FROM posts WHERE id = $1"

	res := r.DataBase.QueryRow(query, idPost)

	postDB := &Post{}

	err := res.Scan(&postDB.Id, &postDB.Author, &postDB.Time, &postDB.Body)
	if err != nil {
		return nil, err
	}

	return postDB, nil
}
func (r *repository) ReturnUserPosts(userId int) ([]Post, error) {
	log.Println("INFO --> return user posts")
	query := "SELECT * FROM posts WHERE author = $1;"

	posts := []Post{}

	res, err := r.DataBase.Query(query, userId)
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
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *repository) ReturnTenPosts(idSub int, numPost int) ([]Post, error) {
	log.Println("INFO --> take 10 posts from database")

	query := "SELECT p.id, p.author, p.time, p.body FROM posts p JOIN subscription r ON p.author = r.profile WHERE r.subscriber = $1 LIMIT $2 OFFSET $3"

	posts := []Post{}

	res, err := r.DataBase.Query(query, idSub, 10, numPost) // Хреново, что тут хардкодом 10 постов
	if err != nil {
		return posts, err
	}

	defer res.Close()

	for res.Next() {
		var post Post
		if err = res.Scan(&post.Id, &post.Author, &post.Time, &post.Body); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = res.Err(); err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *repository) ChangeBody(idPost, idUser int, newBody string) error {
	log.Println("INFO --> change post body")

	query := "UPDATE posts SET body = $1 WHERE id = $2 AND author = $3;"

	res, err := r.DataBase.Exec(query, newBody, idPost, idUser)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return errors.New("you can't change this post")
	}

	return nil
}

func (r *repository) Delete(idPost, idUser int) error {
	log.Println("INFO --> delete post from database")

	query := "DELETE FROM posts WHERE id = $1 AND  author = $2;"

	res, err := r.DataBase.Exec(query, idPost, idUser)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return errors.New("you can't delete this post")
	}

	return nil
}
