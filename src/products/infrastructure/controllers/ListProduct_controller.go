package controllers

import (
	"github.com/gin-gonic/gin"
	"demo/src/products/application"
)

type ListProductController struct {
	useCase application.ListProduct
}

func NewListProductController(useCase application.ListProduct) *ListProductController {
	return &ListProductController{useCase: useCase}
}

func (lp_c *ListProductController) Execute(c *gin.Context) {
	products := lp_c.useCase.Execute()
	c.JSON(200, products)
}