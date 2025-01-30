package controllers

import (
	"demo/src/products/application"

	"github.com/gin-gonic/gin"
)

type ListProductForIdController struct {
	useCase application.ListProductForId
}

func NewListProductForIdController(useCase application.ListProductForId) *ListProductForIdController {
	return &ListProductForIdController{useCase: useCase}
}

func (lp_c *ListProductForIdController) Execute(c *gin.Context) {
	products := lp_c.useCase.Execute(c.Param("id"))
	c.JSON(200, products)
}
