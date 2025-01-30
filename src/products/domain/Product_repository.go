package domain

import "demo/src/products/domain/entities"

type IProduct interface {
	Save()
	GetAll() []entities.Product
}

type ProductRepository interface {
	GetAll() []entities.Product
	GetByID(id string) entities.Product    // Método para obtener un producto por ID
	Create(product entities.Product) error // Método para crear un nuevo producto

}
