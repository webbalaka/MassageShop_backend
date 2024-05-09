package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
)

func GetMassageShops(c *gin.Context) {
	var massageShops []models.MassageShops
	result := initializers.DB.Find(&massageShops)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    massageShops,
	})
}

func GetMassageShop(c *gin.Context) {
	id := c.Param("id")
	var massageShops []models.MassageShops
	result := initializers.DB.First(&massageShops, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    massageShops,
	})
}

func CreateMassageShop(c *gin.Context) {
	var input models.MassageShops
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid input data",
		})
	}
	newMassageShop := models.MassageShops{
		Name:    input.Name,
		Address: input.Address,
		Tel:     input.Tel,
	}
	result := initializers.DB.Create(&newMassageShop)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    newMassageShop,
	})
}

func UpdateMassageShop(c *gin.Context) {
	id := c.Param("id")
	var massageShops []models.MassageShops
	result := initializers.DB.Find(&massageShops, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error,
		})
		return
	}

	var input models.MassageShops
	if error := c.ShouldBindJSON(&input); error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error,
		})
		return
	}

	result = initializers.DB.Model(&massageShops).Updates(input)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    massageShops,
	})
}

func DeleteMassageShop(c *gin.Context) {
	id := c.Param("id")
	var massageShops []models.MassageShops
	result := initializers.DB.Find(&massageShops, id)
	if result.Error != nil {
		if result.Error != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": result.Error,
			})
			return
		}
	}

	result = initializers.DB.Delete(&massageShops)
	if result.Error != nil {
		if result.Error != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": result.Error,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "delete successful",
	})
}
