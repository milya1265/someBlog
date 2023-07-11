package handlers

import (
	"github.com/gin-gonic/gin"
	"someBlog/pkg/repository"
)

func InitRoutes(repo *repository.Repository) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", SignIn(repo.DataBase))
		auth.POST("/sign-up", SignUp(repo.DataBase))
	}
	blog := router.Group("/blog")
	{
		blog.POST("/post", Authorize(repo.DataBase), CreateNewPost(repo.DataBase))
		blog.GET("/post", Authorize(repo.DataBase), GetPost(repo.DataBase))
		blog.POST("/post/:idPost/comment", Authorize(repo.DataBase), CreateNewComment(repo.DataBase))
		blog.GET("/feed", Authorize(repo.DataBase), CreateFeed(repo.DataBase))
	}
	user := router.Group("/user")
	{
		user.GET("/:id", Authorize(repo.DataBase), GetUser(repo.DataBase))
		user.GET("/:id/posts", Authorize(repo.DataBase), GetUserPosts(repo.DataBase))
		user.POST("/sub", Authorize(repo.DataBase), Subscribe(repo.DataBase))
		user.DELETE("/unsub", Authorize(repo.DataBase), Unsubscribe(repo.DataBase))
	}
	return router
}
