package handlers

import (
	"github.com/gin-gonic/gin"
	"someBlog/db"
	auth2 "someBlog/pkg/handlers/auth"
	blog2 "someBlog/pkg/handlers/blog"
)

func InitRoutes(db *db.DB) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", auth2.SignIn(db.DataBase))
		auth.POST("/sign-up", auth2.SignUp(db.DataBase))
	}
	blog := router.Group("/blog")
	{
		blog.POST("/post", auth2.Authorize(db.DataBase), blog2.CreateNewPost(db.DataBase))
		blog.GET("/post", auth2.Authorize(db.DataBase), blog2.GetPost(db.DataBase))
		blog.POST("/:idPost/comment", auth2.Authorize(db.DataBase), blog2.CreateNewComment(db.DataBase))
		blog.GET("/:id/posts", auth2.Authorize(db.DataBase), blog2.GetUserPosts(db.DataBase))
	}
	return router
}
