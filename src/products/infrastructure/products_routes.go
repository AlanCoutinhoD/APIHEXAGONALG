package infrastructure

import (
	"demo/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
	listController      *controllers.ListProductController
	createController    *controllers.CreateProductController
	listForIdController *controllers.ListProductForIdController
	deleteController    *controllers.DeleteProductController
	updateController    *controllers.UpdateProductController
}

func NewProductRouter(listController *controllers.ListProductController, createController *controllers.CreateProductController, listForIdController *controllers.ListProductForIdController, deleteController *controllers.DeleteProductController, updateController *controllers.UpdateProductController) *ProductRouter {
	return &ProductRouter{
		listController:      listController,
		createController:    createController,
		listForIdController: listForIdController,
		deleteController:    deleteController,
		updateController:    updateController,
	}
}

func (pr *ProductRouter) SetupRoutes(router *gin.Engine) {
	productsGroup := router.Group("/products")
	{
		productsGroup.GET("", pr.listController.Execute)
		productsGroup.GET("/:id", pr.listForIdController.Execute)
		productsGroup.POST("", pr.createController.Execute)
		productsGroup.DELETE("/:id", pr.deleteController.Execute)
		productsGroup.PUT("/:id", pr.updateController.Execute)
	}
}
