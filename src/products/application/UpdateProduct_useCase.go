package application

import (
	"demo/src/products/domain/repositories"
	"demo/src/products/domain/entities"
	"log"
	"strconv"
)

type UpdateProduct struct {
	repository domain.IProductRepository
}

func NewUpdateProduct(repository domain.IProductRepository) *UpdateProduct {
	return &UpdateProduct{repository: repository}
}

func (up *UpdateProduct) Execute(product entities.Product) error {
	return up.repository.Update(product)
}

func (up *UpdateProduct) ExecuteByID(id string, name string, price float64) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error al convertir ID: %v", err)
		return err
	}

	product, err := up.repository.GetByID(intID)
	if err != nil {
		log.Printf("Error al obtener producto: %v", err)
		return err
	}

	product.Name = name
	product.Price = price
	return up.repository.Update(*product)
}
