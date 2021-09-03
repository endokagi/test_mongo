package main

import (
	"os"
	"testmongo/controller"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

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
	route.Run(":" + os.Getenv("PORT"))

}
