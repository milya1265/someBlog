package router

import (
	"github.com/gin-gonic/gin"
	"someBlog/internal/adapters/api/auth"
	"someBlog/internal/adapters/api/comment"
	"someBlog/internal/adapters/api/post"
	"someBlog/internal/adapters/api/user"
)

func InitRoutes(authHandler auth.Handler, userHandler user.Handler, postHandler post.Handler,
	comHandler comment.Handler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", authHandler.SignIn())
		auth.POST("/sign-up", authHandler.SignUp())
	}
	blog := router.Group("/blog")
	{
		post := blog.Group("/post")
		{
			post.POST("", authHandler.Authorize(), postHandler.Create())
			post.GET("/:idPost", authHandler.Authorize(), postHandler.Get())
			post.PUT("/:idPost", authHandler.Authorize(), postHandler.Edit())
			post.DELETE("/:idPost", authHandler.Authorize(), postHandler.Delete())

			post.POST("/:idPost/comment", authHandler.Authorize(), comHandler.Create())
			post.GET("/:idPost/comment", authHandler.Authorize(), comHandler.GetPostComment())

			post.GET("/comment/:idCom", authHandler.Authorize(), comHandler.Get())
			post.PUT("/comment/:idCom", authHandler.Authorize(), comHandler.Edit())
			post.DELETE("/comment/:idCom", authHandler.Authorize(), comHandler.Delete())

		}

		blog.GET("/feed", authHandler.Authorize(), postHandler.CreateFeed())
		blog.GET("/feed/:page", authHandler.Authorize(), postHandler.CreateFeed())
	}
	user := router.Group("/user")

	{
		user.GET("/:idUser", authHandler.Authorize(), userHandler.Get())
		user.PATCH("", authHandler.Authorize(), userHandler.EditProfile())
		user.DELETE("", authHandler.Authorize(), userHandler.Delete())

		user.GET("/:idUser/posts", authHandler.Authorize(), postHandler.GetUserPosts())

		user.POST("/sub/:userId", authHandler.Authorize(), userHandler.Subscribe())
		user.DELETE("/unsub/:userId", authHandler.Authorize(), userHandler.Unsubscribe())
	}
	return router
}
