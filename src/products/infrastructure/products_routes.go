package infrastructure

import (
	"demo/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	listController   *controllers.ListProductController
	createController *controllers.CreateProductController
	listForIdController *controllers.ListProductForIdController
}

func NewProductRouter(listController *controllers.ListProductController, createController *controllers.CreateProductController, listForIdController *controllers.ListProductForIdController) *ProductRouter {
	return &ProductRouter{
		listController:   listController,
		createController: createController,
		listForIdController: listForIdController,
	}
}

// SetupRoutes configura todas las rutas relacionadas con productos
func (pr *ProductRouter) SetupRoutes(router *gin.Engine) {
	productsGroup := router.Group("/products")
	{
		productsGroup.GET("", pr.listController.Execute)
		productsGroup.GET("/:id", pr.listForIdController.Execute)
		productsGroup.POST("", pr.createController.Execute)
	}
}
