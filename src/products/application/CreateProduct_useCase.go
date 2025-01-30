package application

import (
	"demo/src/products/domain"
	"demo/src/products/domain/entities"
)

type CreateProduct struct {
	repository domain.ProductRepository
}

func NewCreateProduct(repository domain.ProductRepository) *CreateProduct {
	return &CreateProduct{repository: repository}
}

func (cp *CreateProduct) Execute(product entities.Product) error {
	return cp.repository.Create(product)
}
