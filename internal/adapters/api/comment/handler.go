package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/internal/domain/comment"
	"strconv"
	"time"
)

func (h *handler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		idCom, err := strconv.Atoi(c.Param("idCom"))
		if err != nil {
			log.Println("Error with convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Com, err := h.Service.Get(idCom)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "you can't get this comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comment": Com})

	}
}

func (h *handler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		newCom := &comment.Comment{}

		idPost := c.Param("idPost")

		newCom.Time = time.Now()

		var err error
		newCom.IdPost, err = strconv.Atoi(idPost)

		if err != nil {
			log.Println("Error with convert to int ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newCom.Author = c.Keys["userId"].(int)

		if err = c.BindJSON(newCom); err != nil {
			log.Println("Error with bind JSON comment ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = h.Service.Create(newCom); err != nil {
			log.Println("Error with insert new comment", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "comment is created"})
	}
}

func (h *handler) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newComment comment.Comment

		idCom, err := strconv.Atoi(c.Param("idCom"))
		if err != nil {
			log.Println("Error with convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = c.BindJSON(&newComment); err != nil {
			log.Println("Error with bind comment to JSON:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		UserId := c.Keys["userId"].(int)
		if err = h.Service.Edit(idCom, UserId, newComment.Body); err != nil {
			log.Println("Error with edit comment", err)
			c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "comment is changed"})
	}
}

func (h *handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idCom, err := strconv.Atoi(c.Param("idCom"))

		if err != nil {
			log.Println("Error with convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		UserId := c.Keys["userId"].(int)

		if err = h.Service.Delete(idCom, UserId); err != nil {
			log.Println("Error with delete comment", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}
}

func (h *handler) GetPostComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		idPost, err := strconv.Atoi(c.Param("idPost"))
		if err != nil {
			log.Println("Error with convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Comments, err := h.Service.GetPostComment(idPost)
		if err != nil {
			log.Println("Error with returning comments", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comments": Comments})

	}
}
