package routes

import (
	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/gin-gonic/gin"
)

func InitProductRoutes(router *gin.Engine, productController controller.ProductController) {
	router.GET("/products", productController.GetProducts)
	router.POST("/product", productController.CreateProduct)
	router.GET("/product/:productId", productController.GetProductById)
	router.DELETE("/product/:productId", productController.DeleteProduct)
	router.PUT("/product/:productId", productController.UpdateProduct)
}