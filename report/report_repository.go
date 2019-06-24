package report

import (
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	getProductValuesReport() ([]productStock, error)
	getSalesReport() ([]order, error)
}

type reportRepository struct {
	DB *gorm.DB
}

func initRepository(db *gorm.DB) RepositoryInterface {
	return &reportRepository{
		DB: db,
	}
}

func (repo *reportRepository) getProductValuesReport() ([]productStock, error) {
	var productStocks []productStock
	result := repo.DB.Joins("join products on products.id = product_stocks.product_id").Preload("Product").Find(&productStocks)
	return productStocks, result.Error
}

func (repo *reportRepository) getSalesReport() ([]order, error) {
	var orders []order
	result := repo.DB.
		Joins("join products on products.id = orders.product_id").
		Preload("Product").
		Joins("join product_stocks on product_stocks.id = orders.product_stock_id").
		Preload("ProductStock").
		Find(&orders)
	return orders, result.Error
}
