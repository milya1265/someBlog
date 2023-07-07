package pkg

import (
	"time"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Comment struct {
	Id     int       `json:"id"`
	Author string    `json:"author"`
	PostID int       `json:"post_id"`
	Time   time.Time `json:"time"`
	Body   string    `json:"body"`
}

type Article struct {
	Author int       `json:"author"`
	Id     int       `json:"id"`
	Time   time.Time `json:"time"`
	Name   string    `json:"name"`
	Body   string    `json:"body"`
}

type Post struct {
	Id     int       `json:"id"`
	Author int       `json:"author"`
	Time   time.Time `json:"time"`
	Name   string    `json:"name"`
	Body   string    `json:"body"`
}
