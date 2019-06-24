package product

import (
	"github.com/jinzhu/gorm"
)

type repositoryInterface interface {
	createProduct(p Product) (Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func initRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		DB: db,
	}
}

func (productRepo *productRepository) createProduct(product Product) (Product, error) {
	result := productRepo.DB.Create(&product)
	return product, result.Error
}
