package infrastructure

import (
	"demo/src/employees/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type EmployeeRouter struct {
	listController *controllers.ListEmployeeController
	createController *controllers.CreateEmployeeController
}

func NewEmployeeRouter(listController *controllers.ListEmployeeController, createController *controllers.CreateEmployeeController) *EmployeeRouter {
	return &EmployeeRouter{listController: listController, createController: createController}
}



func (er *EmployeeRouter) SetupRoutes(router *gin.Engine) {
	employeesGroup := router.Group("/employees")
	{
		employeesGroup.POST("", er.createController.Execute)
		employeesGroup.GET("", er.listController.Execute)
		
	}
}


