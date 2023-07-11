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

func CreateNewPost(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPost pkg.Post

		if err := c.BindJSON(&newPost); err != nil {
			log.Println("Error with bind JSON post: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		newPost.Author = c.Keys["userId"].(int)
		newPost.Time = time.Now()

		if err := repository.InsertPost(&newPost, database); err != nil {
			log.Println("Error with insert post to db:", err)
			c.AbortWithStatus(http.StatusNotImplemented)
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

		post, err := repository.SearchPostByID(idPost, database)
		if err != nil {
			log.Println("Error with select post:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if post != nil {
			c.JSON(http.StatusOK, gin.H{"post": post})
		} else {
			log.Println("Post ID is not found")
			c.AbortWithStatus(http.StatusInternalServerError)
		}

	}
}

func GetUserPosts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Error with convert postID to int ", err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		posts, err := repository.ReturnUserPosts(userId, db)
		if err != nil {
			log.Println("Error with get posts from db", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": userId, "posts": posts})
	}
}

//func EditPost(db *sql.DB) gin.HandlerFunc {
//
//}
