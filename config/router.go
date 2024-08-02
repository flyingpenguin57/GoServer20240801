package config

import (
	"github.com/gin-gonic/gin"
	"goserver/controller"
	"goserver/middleware"
)

func CreateRouter(router *gin.Engine) {

	// 创建一个基础路径为 /api 的路由组
	apiGroup := router.Group("/api")
	{
		// 在 /api 路由组中创建一个 /users 路由组
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", controller.Register)
			userGroup.POST("/login", controller.Login)
			userGroup.POST("/userInfo", middleware.JWTAuthMiddleware(), controller.QueryUserInfo)
			userGroup.POST("/updateUserInfo",  middleware.JWTAuthMiddleware(), controller.UpdateUserInfo)
			userGroup.POST("/deleteUser",  middleware.JWTAuthMiddleware(), controller.DeleteUser)
		}
	}
}