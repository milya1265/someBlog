package auth

//
//import (
//	user2 "someBlog/internal/domain/user"
//	"testing"
//)
//
////func (r *repository) InsertUser(user *user2.User) (int, error) {
////	query := "INSERT INTO users (name, surname, email, password) VALUES ($1, $2, $3, $4) RETURNING id;"
////
////	row := r.DataBase.QueryRow(query, user.Name, user.Surname, user.Email, user.Password)
////
////	var idNewUser int
////
////	if err := row.Scan(&idNewUser); err != nil {
////		log.Println("Error with scan row:", err)
////		return -1, err
////	}
////
////	return idNewUser, nil
////}
//
//func (r *repository) TestInsertUser(t *testing.T) {
//
//	tests := []struct {
//		name string
//		user user2.User
//		id   int
//		err  error
//	}{
//		{
//			name: "OK",
//			user: user2.User{
//				Name:     "test",
//				Surname:  "test",
//				Email:    "test@lol.ru",
//				Password: "test",
//			},
//			id:  1,
//			err: nil,
//		},
//	}
//
//	query := "INSERT INTO users (name, surname, email, password) VALUES ($1, $2, $3, $4) RETURNING id;"
//
//}
