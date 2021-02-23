package routes

import (
	"rest_golang/app/controllers"
	auth "rest_golang/app/jwt"

	_ "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	route := r.Group("/api/user")
	route.POST("/", controllers.CreateUserPublic)
	route.POST("/login", auth.JwtMiddleware().LoginHandler)
	route.Use(auth.JwtMiddleware().MiddlewareFunc())
	{
		route.POST("/json", controllers.CreateTest)

	}
}
