package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Tel      string `json:tel`
	Email    string `json:email`
	Role     string `json:role`
	Password string `json:password`
}
