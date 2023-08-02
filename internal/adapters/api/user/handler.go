package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		uID, err := strconv.Atoi(c.Param("idUser"))
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		user, err := h.Service.Get(uID)
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
		profileId, err := strconv.Atoi(c.Param("userId"))
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
		profileId, err := strconv.Atoi(c.Param("userId"))
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

func (h *handler) EditProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Keys["userId"].(int)

		dtoUser := &struct {
			name    string `json:"name"`
			surname string `json:"surname"`
		}{}

		var response []byte
		_, err := c.Request.Body.Read(response)

		log.Println(response)

		if err != nil {
			log.Println("Error with read JSON", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.BindJSON(dtoUser); err != nil {
			log.Println("Error with bind JSON", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println(dtoUser)

		if err := h.Service.EditUser(userId, dtoUser.name, dtoUser.surname); err != nil {
			log.Println("Error with edit profile", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile is changed"})
	}
}

func (h *handler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId := c.Keys["userId"].(int)

		if err := h.Service.Delete(userId); err != nil {
			log.Println("Error with deleting user", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Profile is deleted"})
	}
}
