package main

import (
	_ "github.com/lib/pq"
	"log"
	auth2 "someBlog/internal/adapters/api/auth"
	comment2 "someBlog/internal/adapters/api/comment"
	post2 "someBlog/internal/adapters/api/post"
	user2 "someBlog/internal/adapters/api/user"
	router2 "someBlog/internal/adapters/router"
	"someBlog/internal/adapters/server"
	"someBlog/internal/domain/auth"
	"someBlog/internal/domain/comment"
	"someBlog/internal/domain/post"
	"someBlog/internal/domain/user"
	"someBlog/pkg/repository"
)

// docker run --name BlogDB -p 5432:5432 -e POSTGRES_USER=dmilya -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=BlogDB -d postgres:13.3

// Invoke-WebRequest -Uri "http://localhost:8080/newUser" -Method POST -Body '{"name":"Egor","age":12}' -Headers @{"Content-Type"="application/json"}

func main() {
	repos := repository.Repository{}
	DB, err := repos.Open("postgres://dmilya:qwerty@localhost:5432/BlogDB?sslmode=disable")
	if err != nil {
		log.Fatalln("Error with open database", err)
	}

	authStorage := auth.NewStorage(DB)
	authService := auth.NewService(&authStorage)
	authHandlers := auth2.NewHandler(&authService)

	userStorage := user.NewStorage(DB)
	userService := user.NewService(&userStorage)
	userHandlers := user2.NewHandler(&userService)

	postStorage := post.NewStorage(DB)
	postService := post.NewService(&postStorage)
	postHandlers := post2.NewHandler(&postService)

	commentStorage := comment.NewStorage(DB)
	commentService := comment.NewService(&commentStorage)
	commentHandlers := comment2.NewHandler(&commentService)

	router := router2.InitRoutes(authHandlers, userHandlers, postHandlers, commentHandlers)

	serv := server.Server{}
	log.Fatalln(serv.Run("localhost:8080", router))

}
