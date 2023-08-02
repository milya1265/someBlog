package user

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
