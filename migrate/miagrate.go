package main

import (
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}