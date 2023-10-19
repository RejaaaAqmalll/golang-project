package main

import (
	"log"
	"os"
	"set-up-Golang/config"
	"set-up-Golang/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error can,t get .env")
	}
	route := gin.Default()
	config.ConnectDB()
	r := route.Group("/user")
	{
		r.POST("/register", controller.Register)
		r.POST("login", controller.Login)
		r.POST("logout", controller.Logout)
	}

	// port := os.Getenv("PORT_DATA")
	route.Run(":" + os.Getenv("PORT_DATA"))
}
