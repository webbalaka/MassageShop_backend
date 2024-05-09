package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	PickupDate  string  `json:pickupDate`
	User        float64 `json:user`
	Name        string  `json:name`
	Email       string  `json:email`
	PhoneNumber string  `json:phoneNumber`
	MassageShop float64 `json:massageShop`
}
