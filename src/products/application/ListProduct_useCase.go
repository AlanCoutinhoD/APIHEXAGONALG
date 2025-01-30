package application

import "demo/src/products/domain/entities"
import "demo/src/products/domain"


type ListProduct struct {
	repository domain.ProductRepository
}

func NewListProduct(repository domain.ProductRepository) *ListProduct {
	return &ListProduct{repository: repository}
}

func (lp *ListProduct) Execute() []entities.Product {
	return lp.repository.GetAll()
}