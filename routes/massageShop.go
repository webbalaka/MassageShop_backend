package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/controllers"
)

func MassageShopRouter(r *gin.Engine) {
	r.GET("/api/v1/MassageShops", controllers.GetMassageShops)
	r.GET("/api/v1/MassageShop/:id", controllers.GetMassageShop)
	r.POST("/api/v1/MassageShops", controllers.CreateMassageShop)
	r.PUT("/api/v1/MassageShop/:id", controllers.UpdateMassageShop)
	r.DELETE("/api/v1/MassageShop/:id", controllers.DeleteMassageShop)
}