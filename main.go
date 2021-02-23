package main

import (
	"fmt"
	"log"
	"os"
	"rest_golang/app/models"
	"rest_golang/app/routes"
	db "rest_golang/db"
	"rest_golang/envflag"
	"rest_golang/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	envflag.Init()
	router := gin.Default()
	router.Use(utils.CORSMiddleware())
	db.InitDB()
	db.DB.AutoMigrate(&models.User{})
	routes.User(router)

	port := os.Getenv("PORT")
	fmt.Printf("%+v\n", port)
	erro := router.Run(":" + port)
	log.Panic(erro)
}
