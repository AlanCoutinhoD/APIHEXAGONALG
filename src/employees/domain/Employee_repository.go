package domain

import "demo/src/employees/domain/entities"


type IEmployeeRepository interface {
	GetAll() []entities.Employee
	Create(employee entities.Employee)
}
