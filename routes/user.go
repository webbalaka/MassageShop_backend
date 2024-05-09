package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/controllers"
	"github.com/webbalaka/MassageShop_backend/middleware"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/api/v1/auth/register", controllers.Register)
	r.POST("/api/v1/auth/login", controllers.Login)
	r.GET("/api/v1/auth/me",  middleware.Protect(), controllers.GetMe)
	r.GET("/api/v1/auth/logout", controllers.Logout)
}