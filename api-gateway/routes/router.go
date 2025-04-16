package routes

import (
	"github.com/gin-gonic/gin"
	"seotrang.com/rgpc-microservice/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Route lấy tất cả users
		api.GET("/users", controllers.GetUsers)
		// Route lấy user theo id
		api.GET("/users/:id", controllers.GetUser)
		// Route lấy tất cả products
		api.GET("/products", controllers.GetAllProducts)
		// Route lấy sản phẩm theo id
		api.GET("/products/:id", controllers.GetProduct)

		api.GET("/orders", controllers.GetAllOrders)
		api.GET("/orders/:id", controllers.GetOrder)

	}
}
