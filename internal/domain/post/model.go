package post

import "time"

type Post struct {
	Id     int       `json:"id"`
	Author int       `json:"author"`
	Time   time.Time `json:"time"`
	Body   string    `json:"body" binding:"required"`
}
