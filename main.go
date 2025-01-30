package main

import (
	employeesInfrastructure "demo/src/employees/infrastructure"
	productsInfrastructure "demo/src/products/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Inicializar infraestructura de productos
	productsInfrastructure.Init(router)

	// Inicializar infraestructura de empleados
	employeesInfrastructure.Init(router)

	router.Run(":8080")
}
