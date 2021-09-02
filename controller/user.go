package controller

import (
	"testmongo/models"
	"testmongo/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	result, err := repository.GetAllUsers()
	if err != nil {
		messege := models.Get_data_error()
		c.JSON(401, messege)
	} else {
		c.JSON(200, result)
	}
}

func GetUserByID(c *gin.Context) {
	var user models.User
	user.ID = c.Param("id")
	result, err := repository.GetUser(user.ID)
	if err != nil {
		messege := models.Get_data_error()
		c.JSON(401, messege)
	} else {
		c.JSON(200, result)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	_, err := repository.GetUser(user.ID)
	if err != nil {

		if err := repository.CreateUser(user); err != nil {
			messege := models.Create_user_fail()
			c.JSON(401, messege)
		} else {
			messege := models.Create_user_success()
			c.JSON(200, messege)
		}
	} else {
		message := models.Check_ID_fail()
		c.JSON(401, message)
	}

}

func DeleteUserByID(c *gin.Context) {
	var user models.User
	user.ID = c.Param("id")
	if err := repository.DeleteUser(user.ID); err != nil {
		messege := models.Delete_user_fail()
		c.JSON(401, messege)
	} else {
		messege := models.Delete_user_success()
		c.JSON(200, messege)
	}
}

func PutUserByID(c *gin.Context) {
	var user models.User
	user.ID = c.Param("id")
	c.BindJSON(&user)

	_, err := repository.GetUser(user.ID)
	if err != nil {
		message := models.Check_ID_fail()
		c.JSON(401, message)
	} else {
		if err := repository.PutUser(user); err != nil {
			messege := models.Update_user_fail()
			c.JSON(401, messege)
		} else {
			messege := models.Update_user_success()
			c.JSON(200, messege)
		}
	}

}
