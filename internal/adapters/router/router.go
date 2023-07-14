package router

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/adapters/api/auth"
	"someBlog/internal/adapters/api/comment"
	"someBlog/internal/adapters/api/post"
	"someBlog/internal/adapters/api/user"
)

func InitRoutes(authHandler auth.Handler, userHandler user.Handler, postHandler post.Handler, comHandler comment.Handler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", authHandler.SignIn())
		auth.POST("/sign-up", authHandler.SignUp())
	}
	blog := router.Group("/blog")
	{
		blog.POST("/post", authHandler.Authorize(), postHandler.CreateNewPost())
		blog.GET("/post", authHandler.Authorize(), postHandler.GetPost())
		blog.POST("/post/:idPost/comment", authHandler.Authorize(), comHandler.CreateNewComment())
		blog.GET("/feed", authHandler.Authorize(), postHandler.CreateFeed())
	}
	user := router.Group("/user")
	{
		user.GET("/:id", authHandler.Authorize(), userHandler.GetUser())
		user.GET("/:id/posts", authHandler.Authorize(), postHandler.GetUserPosts())
		user.POST("/sub", authHandler.Authorize(), userHandler.Subscribe())
		user.DELETE("/unsub", authHandler.Authorize(), userHandler.Unsubscribe())
	}
	return router
}
