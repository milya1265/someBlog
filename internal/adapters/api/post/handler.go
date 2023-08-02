package post

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/internal/domain/post"
	"strconv"
	"time"
)

func (h *handler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		idPost, err := strconv.Atoi(c.Param("idPost"))
		if err != nil {
			log.Println("Error with parse URI", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post, err := h.Service.GetByID(idPost)
		if err != nil {
			log.Println("Error with select post:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if post != nil {
			c.JSON(http.StatusOK, gin.H{"post": post})
		} else {
			log.Println("Post ID is not found")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	}
}

func (h *handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPost post.Post

		if err := c.BindJSON(&newPost); err != nil {
			log.Println("Error with bind JSON post: ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newPost.Author = c.Keys["userId"].(int)
		newPost.Time = time.Now()

		if err := h.Service.Create(&newPost); err != nil {
			log.Println("Error with insert post to db:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post has been created"})

	}

}

func (h *handler) GetUserPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("idUser"))
		if err != nil {
			log.Println("Error with convert postID to int ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		posts, err := h.Service.GetUserPosts(userId)
		if err != nil {
			log.Println("Error with get posts from db", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": userId, "posts": posts})
	}
}

func (h *handler) CreateFeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Keys["userId"].(int)
		var numPosts = 1

		if c.Param("page") != "" {
			var err error
			numPosts, err = strconv.Atoi(c.Param("page"))
			if err != nil {
				log.Println(c.Request.URL.Query().Get("num"))
				log.Println("Error with get num of feed", err)
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		numTenPost := (numPosts - 1) * 10

		posts, err := h.Service.CreateFeed(userId, numTenPost)
		if err != nil {
			log.Println("Error with returning posts from db:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"post": posts})

	}
}

func (h *handler) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Post post.Post

		if err := c.BindJSON(&Post); err != nil {
			log.Println("Error with bind JSON post: ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idPost, err := strconv.Atoi(c.Param("idPost"))
		if err != nil {
			log.Println("Error with convert to int: ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		idUser := c.Keys["userId"].(int)

		if err = h.Service.Edit(idPost, idUser, Post.Body); err != nil {
			log.Println("Error with edit post:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post has been edited"})

	}
}

func (h *handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idPost, err := strconv.Atoi(c.Param("idPost"))
		if err != nil {
			log.Println("Error with convert to int: ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idUser := c.Keys["userId"].(int)

		if err = h.Service.Delete(idPost, idUser); err != nil {
			log.Println("Error with insert delete:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post has been deleted"})

	}
}
