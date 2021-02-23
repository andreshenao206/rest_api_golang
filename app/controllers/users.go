package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"rest_golang/app/helpers"
	"rest_golang/app/httpmodels"
	"rest_golang/app/models"
	"rest_golang/db"
	"sort"

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
func orderDuplicate(intSlice []int) []int {
	keys := make(map[int]bool)
	Uni := []int{}
	Dup := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			Uni = append(Uni, entry)
		} else {
			Dup = append(Uni, entry)
		}
	}
	z := append([]int{}, append(Uni, Dup...)...)
	return z
}

func CreateTest(c *gin.Context) {
	var userjson httpmodels.Userjson
	c.Bind(&userjson)
	jsonOrig := userjson.Json_user
	sort.Ints(userjson.Json_user)
	emailUser := helpers.EmailbyToken(c)
	sorted := orderDuplicate(userjson.Json_user)
	c.JSON(http.StatusOK, gin.H{
		"Code":      200,
		"unsorted":  jsonOrig,
		"sorted":    sorted,
		"emailUser": emailUser,
	})
}
func GenerateValidUser(email string) (*models.User, error) {
	user, err := helpers.FindUserByEmailAll(email)
	if err == nil {
		return user, errors.New("User exist")
	}
	user.Email = email
	return user, nil
}
