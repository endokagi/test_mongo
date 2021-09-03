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

	route.Use(controller.CORSMiddleware())
	// users
	route.GET("/users", controller.GetAllUser)
	// user
	route.POST("/user", controller.CreateUser)

	// user/:name
	route.GET("/user/:id", controller.GetUserByID)
	route.DELETE("/user/:id", controller.DeleteUserByID)
	route.PUT("/user/:id", controller.PutUserByID)

	route.Run(":" + os.Getenv("PORT"))

}
