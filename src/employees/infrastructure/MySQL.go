package infrastructure

import (
	"demo/src/core"
	"demo/src/employees/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (m *MySQL) Create(employee entities.Employee) {
	query := "INSERT INTO employees (name) VALUES (?)"
	fmt.Println(employee)
	_, err := m.conn.ExecutePreparedQuery(query, employee.Name)
	if err != nil {
		log.Fatalf("Error al crear el empleado: %v", err)
	}
}

func (m *MySQL) GetAll() []entities.Employee {
	query := "SELECT * FROM employees"
	rows := m.conn.FetchRows(query)
	defer rows.Close()

	var employees []entities.Employee
	for rows.Next() {
		var employee entities.Employee
		rows.Scan(&employee.ID, &employee.Name)
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error al iterar sobre las filas: %w", err)
	}

	return employees
}
