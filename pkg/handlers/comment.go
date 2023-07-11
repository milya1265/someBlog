package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/pkg"
	"someBlog/pkg/repository"
	"strconv"
	"time"
)

func CreateNewComment(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		newCom := &pkg.Comment{}

		idPost := c.Param("idPost")

		newCom.Time = time.Now()

		var err error
		newCom.IdPost, err = strconv.Atoi(idPost)

		if err != nil {
			log.Println("Error with convert to int ", err)
			c.Abort() // Вставить ошибку
			return
		}

		newCom.Author = c.Keys["userId"].(int)

		if err = c.BindJSON(newCom); err != nil {
			log.Println("Error with bind JSON comment ", err)
			c.Abort() // вставить ошибку
			return
		}
		//Проверить , что забиндилось только тело комментария
		log.Println(*newCom)

		if err = repository.InsertNewComment(newCom, database); err != nil {
			log.Println("Error with insert new comment", err)
			c.Abort() // вставить ошибку
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "comment is created"})
	}
}