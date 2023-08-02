package comment

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"someBlog/internal/domain/comment"
	"strconv"
	"time"
)

var timeNow = time.Now

func (h *handler) Get() gin.HandlerFunc {
	log.Println("INFO --> starting get comment handler")

	return func(c *gin.Context) {
		idCom, err := strconv.Atoi(c.Param("idCom"))
		if err != nil {
			log.Println("ERROR -->", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Com, err := h.Service.Get(idCom)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comment": Com})

	}
}

func (h *handler) Create() gin.HandlerFunc {
	log.Println("INFO --> starting create comment handler ")

	return func(c *gin.Context) {
		newCom := &comment.Comment{}

		newCom.Time = timeNow()

		var err error
		newCom.IdPost, err = strconv.Atoi(c.Param("idPost"))

		if err != nil {
			log.Println("ERROR -->", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "not int in url param"})
			return
		}

		newCom.Author = c.Keys["userId"].(int)

		if err = c.BindJSON(newCom); err != nil {
			log.Println("ERROR --> bind JSON comment ", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
			return
		}

		idCom, err := h.Service.Create(newCom)
		if err != nil {
			log.Println("ERROR --> insert new comment", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"comment ID": idCom, "message": "comment is created"})
	}
}

func (h *handler) Edit() gin.HandlerFunc {
	log.Println("INFO --> starting edit comment handler")

	return func(c *gin.Context) {
		var newComment comment.Comment

		idCom, err := strconv.Atoi(c.Param("idCom"))
		if err != nil {
			log.Println("ERROR --> convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = c.BindJSON(&newComment); err != nil {
			log.Println("ERROR --> bind comment to JSON:", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		UserId := c.Keys["userId"].(int)
		if err = h.Service.Edit(idCom, UserId, newComment.Body); err != nil {
			log.Println("ERROR --> edit comment", err)
			c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "comment is changed"})
	}
}

func (h *handler) Delete() gin.HandlerFunc {
	log.Println("INFO --> starting delete comment handler")

	return func(c *gin.Context) {
		idCom, err := strconv.Atoi(c.Param("idCom"))

		if err != nil {
			log.Println("ERROR --> convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		UserId := c.Keys["userId"].(int)

		if err = h.Service.Delete(idCom, UserId); err != nil {
			log.Println("ERROR --> delete comment", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "comment has been deleted"})

	}
}

func (h *handler) GetPostComment() gin.HandlerFunc {
	log.Println("INFO --> starting get post comments handler")

	return func(c *gin.Context) {
		idPost, err := strconv.Atoi(c.Param("idPost"))
		if err != nil {
			log.Println("ERROR --> convert to int", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Comments, err := h.Service.GetPostComment(idPost)
		if err != nil {
			log.Println("ERROR --> returning comments", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"comments": Comments})

	}
}
