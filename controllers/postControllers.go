package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/models"
)

func PostsCrate(c *gin.Context) {

	post := models.Post{Title: "Jinzhu", Body: "Post body"}

	result := initializers.DB.Create(&post) 

	if result.Error != nil {
		c.Status(400);
		return;
	}

	c.JSON(200, gin.H{
		"post" : post,
	})
}

func PostsIndex(c *gin.Context){
	var post []models.Post
	initializers.DB.Find(&post)

	c.JSON(200, gin.H{
		"message": "success",
		"data": post,
	})
}