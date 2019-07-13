package model

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	mobile string `json:"mobile"`
}