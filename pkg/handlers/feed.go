package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/pkg/repository"
	"strconv"
)

func CreateFeed(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Keys["userId"].(int)
		numPosts := 1

		if c.Request.URL.Query().Get("num") != "" {
			var err error
			numPosts, err = strconv.Atoi(c.Request.URL.Query().Get("num"))
			if err != nil {
				log.Println(c.Request.URL.Query().Get("num"))
				log.Println("Error with get num of feed", err)
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}

		if numPosts == 1 {
		}
		numTenPost := (numPosts - 1) * 10

		posts, err := repository.ReturnTenPosts(userId, numTenPost, db)
		if err != nil {
			log.Println("Error with returning posts from db:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": posts})

	}
}
