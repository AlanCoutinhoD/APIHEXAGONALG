package controllers

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	useCase *application.CreateProduct
}

func NewCreateProductController(useCase *application.CreateProduct) *CreateProductController {
	return &CreateProductController{useCase: useCase}
}

func (cp_c *CreateProductController) Execute(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := cp_c.useCase.Execute(product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Producto creado correctamente"})
}
