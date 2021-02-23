package helpers

import (
	"errors"
	"rest_golang/db"

	//"fmt"

	"rest_golang/app/models"
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
