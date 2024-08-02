package main

import (
	"goserver/config"
	//"goserver/dao"
	"github.com/gin-gonic/gin"
	"goserver/middleware"
	"goserver/rpc/user_rpc"
)

func main() {
	router := gin.Default()

	router.Use(middleware.RecoveryMiddleware())

	//create router
	config.CreateRouter(router)

	//connect db
	//dao.ConnectDB();

	//run rpc server
	go func() {
		user_rpc.StartRpcServer()
	}()

	//run server
	router.Run("localhost:8080")
}