package routes

import (
	controllers "project_restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemsRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/order_items", controllers.GetOrderItems())
	incomingRoutes.GET("/order_items/:order_id", controllers.GetOrderItem())
	incomingRoutes.GET("/order_items-order/:order_id", controllers.GetOrderItemsByOrderId())
	incomingRoutes.POST("/order_items", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order_items/:order_id", controllers.UpdateOrderItem())
}