package main

import (
	"os"
	middleware "project_restaurant/middleware"
	routes "project_restaurant/routes"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)
}
