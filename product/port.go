package product

import (
	"ecommerce/domain"
	productHandler "ecommerce/rest/handlers/product"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(productID int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(productID int) error
	Update(product domain.Product) (*domain.Product, error)
}
