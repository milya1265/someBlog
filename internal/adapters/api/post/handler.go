package post

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/internal/domain/post"
	"strconv"
	"time"
)

func (h *handler) CreateNewPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPost post.Post

		if err := c.BindJSON(&newPost); err != nil {
			log.Println("Error with bind JSON post: ", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		newPost.Author = c.Keys["userId"].(int)
		newPost.Time = time.Now()

		if err := h.Service.Create(&newPost); err != nil {
			log.Println("Error with insert post to db:", err)
			c.AbortWithStatus(http.StatusNotImplemented)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post has been created"})

	}

}

func (h *handler) GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		idPost, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
		if err != nil {
			log.Println("Error with parse URI", err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		post, err := h.Service.GetByID(idPost)
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

func (h *handler) GetUserPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Error with convert postID to int ", err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		posts, err := h.Service.GetUserPosts(userId)
		if err != nil {
			log.Println("Error with get posts from db", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": userId, "posts": posts})
	}
}

func (h *handler) CreateFeed() gin.HandlerFunc {
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

		posts, err := h.Service.CreateFeed(userId, numTenPost)
		if err != nil {
			log.Println("Error with returning posts from db:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": posts})

	}
}

//func EditPost(db *sql.DB) gin.HandlerFunc {
//
//}
