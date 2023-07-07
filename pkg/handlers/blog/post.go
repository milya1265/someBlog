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

func CreateNewPost(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPost pkg.Post

		if err := c.BindJSON(&newPost); err != nil {
			log.Println("Error with bind JSON post: ", err)
			c.Abort() // вставить ошибку
			return
		}
		newPost.Author = c.Keys["user"].(pkg.User).Id
		newPost.Time = time.Now()

		if err := blog.InsertPost(&newPost, database); err != nil {
			log.Println("Error with insert post to db:", err)
			c.Abort() // вставить ошибку
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post has been created"})

	}

}

func GetPost(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idPost, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
		if err != nil {
			log.Println("Error with parse URI", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		post, err := blog.FetchPostByID(idPost, database)
		if err != nil {
			log.Println("Error with select post:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if post != nil {
			c.JSON(http.StatusIMUsed, gin.H{"post": post})
		} else {
			log.Println("Post ID is not found")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

	}
}

func GetUserPosts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		posts, err := blog.FetchUserPosts(username, db)
		if err != nil {
			log.Println("Error with get posts from db", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": username, "posts": *posts})
	}
}
