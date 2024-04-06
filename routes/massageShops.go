package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/controllers"
)

func MassageShopRouter(r *gin.Engine) {
	r.GET("/api/massageShops", controllers.GetMassageShops)
}