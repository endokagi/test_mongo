package main

import (
	"testmongo/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()

	// users
	route.GET("/users", controller.GetAllUser)
	// user
	route.POST("/user", controller.CreateUser)

	// user/:name
	route.GET("/user/:id", controller.GetUserByID)
	route.DELETE("/user/:id", controller.DeleteUserByID)
	route.PUT("/user/:id", controller.PutUserByID)

	route.Use(controller.CORSMiddleware())
	route.Run(":4000")

}
