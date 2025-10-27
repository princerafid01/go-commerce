package product

import "ecommerce/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.productRepo.Create(p)
}
func (svc *service) Get(productID int) (*domain.Product, error) {
	return svc.productRepo.Get(productID)
}
func (svc *service) List() ([]*domain.Product, error) {
	return svc.productRepo.List()
}
func (svc *service) Delete(productID int) error {
	return svc.productRepo.Delete(productID)
}
func (svc *service) Update(product domain.Product) (*domain.Product, error) {
	return svc.productRepo.Update(product)
}
