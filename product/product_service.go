package product

import (
	productstock "github.com/reynaldipane/test-ijah/product_stock"
)

type serviceInterface interface {
	createProduct(p Product) (Product, error)
}

/*Service is the product's service object
that have access to the product repository*/
type Service struct {
	repo                repositoryInterface
	productStockService productstock.ServiceInterface
}

func (service *Service) createProduct(product Product) (Product, error) {
	createdProduct, err := service.repo.createProduct(product)

	if err != nil {
		return createdProduct, err
	}

	_, err = service.productStockService.CreateProductStock(productstock.ProductStock{
		ProductID:       createdProduct.ID,
		Stock:           0,
		TotalOrdered:    0,
		AverageBuyPrice: 0.0,
	})

	return createdProduct, err
}
