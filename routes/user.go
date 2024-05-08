package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/controllers"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/api/v1/auth/register", controllers.Register)

}