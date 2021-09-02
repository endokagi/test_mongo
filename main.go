package main

import (
	"os"
	"testmongo/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("PORT", "4000")

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
	route.Run(":" + os.Getenv("PORT"))

}
