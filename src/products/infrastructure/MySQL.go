package infrastructure

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
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

func (mysql *MySQL) Save() {
	log.Println("Método Save() está deprecado. Por favor usar Create()")
}

func (mysql *MySQL) GetAll() []entities.Product {
	query := "SELECT * FROM product"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var idproduct int
		var name string
		var price float32
		if err := rows.Scan(&idproduct, &name, &price); err != nil {
			fmt.Println("error al escanear la fila: %w", err)
		}
		products = append(products, entities.Product{ID: idproduct, Name: name, Price: float64(price)})
	}

	if err := rows.Err(); err != nil {
		fmt.Println("error iterando sobre las filas: %w", err)
	}

	return products
}

func (m *MySQL) GetByID(id string) entities.Product {
	query := "SELECT * FROM product WHERE id = ?"
	rows := m.conn.FetchRows(query, id)
	defer rows.Close()

	var product entities.Product
	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Price)
	}

	return product
}

func (m *MySQL) Create(product entities.Product) error {
	query := "INSERT INTO product (name, price) VALUES (?, ?)"

	result, err := m.conn.ExecutePreparedQuery(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("error al crear el producto: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Producto creado correctamente")
	}

	return nil
}
