package router

import (
	"github.com/gin-gonic/gin"
	"video-api/handler"
	"video-api/middleware"
)

func SetupRouter(
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./uploads")
	apigroup := r.Group("/")
	userGroup := apigroup.Group("/user")
	{
		userGroup.POST("register/", userHandler.Register)
		userGroup.POST("login/", userHandler.Login)
		userGroup.GET("/", middleware.AuthMiddleware(), userHandler.GetUserInfo)
		userGroup.POST("avatar/upload/", middleware.AuthMiddleware(), userHandler.UploadAvatar)
	}
	return r
}
