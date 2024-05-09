package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
)

func GetReservations (c *gin.Context){
	var reservation []models.Reservation
	result := initializers.DB.Find(&reservation)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": reservation,
	})
}

func GetReservation (c *gin.Context){
	id := c.Param("id")
	var reservation []models.Reservation
	result := initializers.DB.Find(&reservation, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": reservation,
	})
}

func CreateReservation (c *gin.Context) {
	id := c.Param("id")
	userClaims, exists  := c.Get("user")

	if !exists {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Not authorize to access this route",
		})
		return
	}

	claims, ok := userClaims.(jwt.MapClaims)
	userID  := claims["ID"].(float64)
	if !ok {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Not authorize to access this route",
		})
		return
	}

	var input models.Reservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	massageShopFloatID, err := (strconv.Atoi(id))
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}	
	massageShopID := float64(massageShopFloatID)

	var massageShop []models.MassageShops 
	result := initializers.DB.Find(&massageShop, massageShopID)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	newReservation := models.Reservation{
		PickupDate: input.PickupDate,
		User: userID,
		Name: input.Name,
		Email: input.Email,
		PhoneNumber: input.PhoneNumber,
		MassageShop: massageShopID,
	}

	result = initializers.DB.Create(&newReservation)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": newReservation,
	})
}