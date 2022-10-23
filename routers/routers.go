package routers

import (
	"rest-api-project/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	ordersRoute := router.Group("/orders") 
	{
		ordersRoute.POST("", controllers.CreateOrders)
		ordersRoute.GET("", controllers.GetOrders)
		ordersRoute.PATCH("/:orderId", controllers.UpdateOrders)
		ordersRoute.DELETE("/:orderId")
	}
	//Write route here
	router.GET("/")
	

	return router
}