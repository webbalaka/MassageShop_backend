package models

import (
	"time"

	"gorm.io/gorm"
)

type MassageShopSchema struct {
	gorm.Model
	ID uint
	name string 
	address string
	open time.Time
	close time.Time
	picture *string

}