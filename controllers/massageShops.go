package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
)

func GetMassageShops(c *gin.Context) {
	var massageShops []models.MassageShopSchema
	initializers.DB.Find(&massageShops)

	if massageShops != nil {
		c.JSON(400,  gin.H{
			"success": false,
		})
	}

	c.JSON(200, gin.H{
		"success" : true,
		"data" : massageShops,
	})
}