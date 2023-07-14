package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err := h.Service.GetUser(uID)
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User is not found", "error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"name": user.Name,
			"surname": user.Surname,
			"id":      user.Id})
	}
}

func (h *handler) Subscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		profileId, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
		if err != nil {
			log.Println("Error with get key from context")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		subID := c.Keys["userId"].(int)

		if profileId == subID {
			log.Println("Error with profileId == subID")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err = h.Service.Subscribe(subID, profileId); err != nil {
			log.Println("Error with create new subscribe", err)
			c.AbortWithStatus(http.StatusNotImplemented)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "You're subscribed"})
	}
}

func (h *handler) Unsubscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		profileId, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
		if err != nil {
			log.Println("Error with get key from context")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		subID := c.Keys["userId"].(int)

		if profileId == subID {
			log.Println("Error with profileId == subID")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err = h.Service.Unsubscribe(subID, profileId); err != nil {
			log.Println("Error with  unsubscribe", err)
			c.AbortWithStatus(http.StatusNotImplemented)
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "You're unsubscribed"})
	}
}
