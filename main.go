package main

import (
	"set-up-Golang/config"
	"set-up-Golang/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	config.ConnectDB()
	r := route.Group("/user")
	{
		r.POST("/register", controller.Register)
		r.POST("login", controller.Login)
		r.POST("logout", controller.Logout)
	}

	route.Run(":8080")
}
