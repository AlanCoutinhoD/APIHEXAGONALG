package application

import (
	"demo/src/products/domain/entities"
	"demo/src/products/domain"
)

type ListProductForId struct {
	repository domain.ProductRepository
}


func NewListProductForId(repository domain.ProductRepository) *ListProductForId {
	return &ListProductForId{repository: repository}
}

func (lp *ListProductForId) Execute(id string) entities.Product {
	return lp.repository.GetByID(id)
}
