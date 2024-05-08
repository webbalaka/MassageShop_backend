package models

import (
	"gorm.io/gorm"
)

type MassageShops struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:address`
	Tel     string `json:tel`
}
