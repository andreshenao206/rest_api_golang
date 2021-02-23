package helpers

import (
	"errors"
	"fmt"
	"log"
	"rest_golang/db"

	//"fmt"

	"rest_golang/app/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func LoginUser(identifier string) (*models.User, error) {
	user, error := FindUserByEmail(identifier)
	if error != nil {
		return nil, errors.New("error login")
	} else {
		return user, nil
	}
}
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
func FindUserByEmailAll(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Unscoped().Where("email = ?", email).First(&user).Error
	return &user, err
}
func FindUserById(id string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func ListUser(email string) ([]models.User, error) {
	var user []models.User
	var err error
	if email != "" {
		err = db.DB.Where("email = ?", email).Find(&user).Error
	} else {
		err = db.DB.Find(&user).Error
	}
	return user, err
}
func EmailbyToken(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	log.Println(claims)
	emailLogin, okE := claims["id"]
	log.Println(emailLogin)
	if !okE {
		return ""
	}
	return fmt.Sprintf("%v", emailLogin)

}
