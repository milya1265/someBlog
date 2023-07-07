package blog

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/db/blog"
	"someBlog/pkg"
	"strconv"
	"time"
)

func CreateNewComment(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		newCom := &pkg.Comment{}

		idPost := c.Param("idPost")

		newCom.Time = time.Now()

		var err error
		newCom.PostID, err = strconv.Atoi(idPost)

		if err != nil {
			log.Println("Error with convert to int ", err)
			c.Abort() // Вставить ошибку
			return
		}
		newCom.Author = c.Keys["user"].(pkg.User).Name

		if err = c.BindJSON(newCom); err != nil {
			log.Println("Error with bind JSON comment ", err)
			c.Abort() // вставить ошибку
			return
		}
		//Проверить , что забиндилось только тело комментария
		log.Println(*newCom)

		if err = blog.InsertNewComment(newCom, database); err != nil {
			log.Println("Error with insert new comment", err)
			c.Abort() // вставить ошибку
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "comment is created"})
	}
}
