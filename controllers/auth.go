package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassWord(password string) (string, error){
	resultPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(resultPassWord), err
}


func Register (c *gin.Context){
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid input data",
		})
	}

	hash, err := hashPassWord(input.Password)

	if(err != nil){
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	newUser := models.User{
		Name: input.Name,
		Tel: input.Tel,
		Email: input.Email,
		Role: input.Role,
		Password: hash,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    input,
	})

}