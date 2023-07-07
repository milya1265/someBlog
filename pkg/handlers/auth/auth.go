package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	db2 "someBlog/db/auth"
	"someBlog/pkg"
	"time"
)

var JWTKey = []byte("lolkekcheburek")

func SignUp(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser pkg.User

		if err := c.BindJSON(&newUser); err != nil {
			log.Println("Error input with bind JSON:", err)
			c.JSON(http.StatusBadRequest, err)
			c.Abort()
			return
		}

		if err := db2.HashPassword(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		idNewUser, err := db2.InsertUser(&newUser, database)
		if err != nil {
			log.Println("Error with insert to database", err)
			c.JSON(http.StatusNotImplemented, err)
			c.Abort()
			return
		}
		newUser.Id = idNewUser

		c.JSON(http.StatusCreated, gin.H{"message": "Insert successfully", "id": newUser.Id, "name": newUser.Name})
	}
}

func SignIn(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := &pkg.User{}

		if err := c.BindJSON(u); err != nil {
			log.Println("Error with bind JSON:", err)
			c.Abort()
			return
		}

		u = db2.CheckPasswordAndReturnUser(u, database)
		if u == nil {
			log.Println("User not found in database")
			c.JSON(http.StatusNotFound, gin.H{"message": "Invalid username or password"})
			c.Abort()
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": u.Id,
			"exp": time.Now().Add(time.Minute * 60).Unix(),
		})

		tokenString, err := token.SignedString(JWTKey)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"user": u.Id,
		})

	}
}

func Authorize(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			log.Println("Error with read cookie:", err)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTKey, nil
		})
		if err != nil {
			log.Println("Error with parse token: ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				log.Println("Error: time is out")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			user := &pkg.User{}
			user = db2.SearchUserByID(int(claims["sub"].(float64)), database)
			if user == nil {
				log.Println("User not found")
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("userId", user.Id)
			c.Next()
		} else {
			log.Println("Validation error")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
