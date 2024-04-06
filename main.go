package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webbalaka/MassageShop_backend/initializers"
	"github.com/webbalaka/MassageShop_backend/routes"
)

func init() {
	initializers.LoadEnvVariables();
	initializers.ConnectToDB();
}

func main() {
	r := gin.Default()
	routes.MassageShopRouter(r)
	r.Run()

}
