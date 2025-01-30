package infrastructure

import (
	"demo/src/employees/application"
	"demo/src/employees/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	employeeRepo := NewMySQL()
	
	getAllEmployee := application.NewListEmployee(employeeRepo)
	listEmployeeController := controllers.NewListEmployeeController(*getAllEmployee)

	createEmployee := application.NewCreateEmployee(employeeRepo)
	createEmployeeController := controllers.NewCreateEmployeeController(*createEmployee)

	employeeRouter := NewEmployeeRouter(listEmployeeController, createEmployeeController)
	employeeRouter.SetupRoutes(router)



	
}
