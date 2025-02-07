package routes

import (
	"github.com/gin-gonic/gin"
	"golang-web-server/config"
	"golang-web-server/controllers"
	"golang-web-server/repositories"
)

func SetupProductRoutes(r *gin.Engine) {
	productRepo := repositories.NewProductRepository(config.DB)
	productsController := controllers.NewProductController(productRepo)

	product := r.Group("/api/product")
	{
		product.POST("/create", productsController.CreateProduct)
		product.GET("/", productsController.GetProducts)
		product.GET("/:id", productsController.GetProduct)
		product.PUT("/:id", productsController.UpdateProduct)
		product.DELETE("/:id", productsController.DeleteProduct)

		product.PUT("/inventory", productsController.UpdateInventory)
		product.GET("/inventory/:id", productsController.GetInventoryByProductID)
	}

	order := r.Group("/api/order")
	{
		order.POST("/", productsController.CreateNewOrder)
		order.GET("/:id", productsController.GetOrder)
	}
}
