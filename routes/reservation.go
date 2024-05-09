package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/controllers"
	"github.com/webbalaka/MassageShop_backend/middleware"
)

func ReservationRouter(r *gin.Engine){
	r.GET("/api/v1/reservations", controllers.GetReservations)
	r.GET("/api/v1/MassageShop/:id/reservation", controllers.GetReservation)
	r.POST("/api/v1/MassageShop/:id/reservation", middleware.Protect(), controllers.CreateReservation)
}