package controllers

import (
	"demo/src/products/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListProductForIdController struct {
	useCase *application.ListProductForId
}

func NewListProductForIdController(useCase *application.ListProductForId) *ListProductForIdController {
	return &ListProductForIdController{useCase: useCase}
}

func (c *ListProductForIdController) Execute(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
