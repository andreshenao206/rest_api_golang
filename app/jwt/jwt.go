package jwt

import (
	"log"
	"rest_golang/app/helpers"
	"rest_golang/app/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	_ "github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

type login struct {
	Email string `form:"email" json:"email" binding:"required"`
}

func JwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {

			if v, ok := data.(*models.User); ok {

				return jwt.MapClaims{
					identityKey: v.Email,
					"username":  v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			log.Printf("claims = %+v\n", claims)
			return &models.User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			return helpers.LoginUser(loginVals.Email)
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.UserName != "" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
