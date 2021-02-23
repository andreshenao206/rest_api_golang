package models

import (
	"github.com/jinzhu/gorm"
)

type Userjson struct {
	gorm.Model
	User_id   int  `json:"user_id" `
	Json_user JSON `json:"json_user" gorm:"type:varchar(255);"`
}
