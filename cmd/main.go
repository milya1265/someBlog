package main

import (
	_ "github.com/lib/pq"
	"log"
	"someBlog/db"
	"someBlog/pkg/handlers"
	"someBlog/pkg/server"
)

// docker run --name BlogDB -p 5432:5432 -e POSTGRES_USER=dmilya -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=BlogDB -d postgres:13.3

// Invoke-WebRequest -Uri "http://localhost:8080/newUser" -Method POST -Body '{"name":"Egor","age":12}' -Headers @{"Content-Type"="application/json"}

func main() {
	db := db.DB{}
	if err := db.Open("postgres://dmilya:qwerty@localhost:5432/BlogDB?sslmode=disable"); err != nil {
		log.Fatalln("Error with open database", err)
	}

	var serv server.Server
	serv.Srv = handlers.InitRoutes(&db)
	serv.Start("localhost:8080") // исправить хардкод порта

}
