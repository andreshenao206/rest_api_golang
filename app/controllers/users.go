package controllers

import (
	"errors"
	"fmt"
	"rest_golang/app/helpers"
	"rest_golang/app/httpmodels"
	"rest_golang/app/models"
	"rest_golang/db"

	"github.com/gin-gonic/gin"
)

func CreateUserPublic(c *gin.Context) {
	var userRequest models.User
	err := c.ShouldBindJSON(&userRequest)
	errorResponse := &httpmodels.ErrorRespose{}
	if err != nil {
		fmt.Println(err)
		errorResponse.Msg = err.Error()
		c.AbortWithStatusJSON(400, errorResponse)
		return
	}

	user, error := GenerateValidUser(userRequest.Email)
	if error != nil {
		errorResponse.Msg = error.Error()
		c.AbortWithStatusJSON(400, errorResponse)
		return
	}
	user.UserName = userRequest.UserName
	db.DB.Create(&user)
	c.JSON(200, user)
}

func GenerateValidUser(email string) (*models.User, error) {
	user, err := helpers.FindUserByEmailAll(email)
	if err == nil {
		return user, errors.New("User exist")
	}
	user.Email = email
	return user, nil
}
