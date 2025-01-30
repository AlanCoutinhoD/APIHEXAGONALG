package infrastructure

import (
	"demo/src/products/application"
	"demo/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// Crear producto
	ps := NewMySQL()
	createProduct := application.NewCreateProduct(ps)
	createProductController := controllers.NewCreateProductController(createProduct)

	// Listar todos los productos
	productRepo := ps
	getAllProduct := application.NewListProduct(productRepo)
	listProductController := controllers.NewListProductController(*getAllProduct)

	// Listar producto por id
	listProductForId := application.NewListProductForId(productRepo)
	listProductForIdController := controllers.NewListProductForIdController(*listProductForId)

	// Configura las rutas
	productRouter := NewProductRouter(listProductController, createProductController, listProductForIdController)
	productRouter.SetupRoutes(router)
}
