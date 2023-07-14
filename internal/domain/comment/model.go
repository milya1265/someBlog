package comment

import "time"

type Comment struct {
	Id     int       `json:"id"`
	Author int       `json:"author"`
	IdPost int       `json:"idPost"`
	Time   time.Time `json:"time"`
	Body   string    `json:"body"`
}
