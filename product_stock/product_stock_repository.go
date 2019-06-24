package productstock

import (
	"github.com/jinzhu/gorm"
)

/*RepositoryInterface wraps the productstock repository ability
so it can be easier to initialize the productService with different kind of repository
*/
type RepositoryInterface interface {
	createProductStock(ps ProductStock) (ProductStock, error)
	updateProductStock(productID string, updateProductStockData ProductStock) (ProductStock, error)
	findProductStockByProductID(productID string) (ProductStock, error)
	findAllProductStock() ([]ProductStock, error)
}

type productStockRepository struct {
	DB *gorm.DB
}

/*InitRepository will returns the repository object that required
to initialize productStockService
*/
func InitRepository(db *gorm.DB) RepositoryInterface {
	return &productStockRepository{
		DB: db,
	}
}

func (repo *productStockRepository) createProductStock(productStockData ProductStock) (ProductStock, error) {
	result := repo.DB.Create(&productStockData)
	return productStockData, result.Error
}

func (repo *productStockRepository) updateProductStock(productID string, updateProductStockData ProductStock) (ProductStock, error) {
	productStock := ProductStock{}

	searchExistingProductStockResult := repo.DB.Where("product_id = ?", productID).First(&productStock)

	if searchExistingProductStockResult.Error != nil {
		return updateProductStockData, searchExistingProductStockResult.Error
	}

	productStock.Stock = updateProductStockData.Stock
	productStock.TotalOrdered = updateProductStockData.TotalOrdered
	productStock.AverageBuyPrice = updateProductStockData.AverageBuyPrice

	updateProductStockResult := repo.DB.Save(&productStock)

	if updateProductStockResult.Error != nil {
		return productStock, updateProductStockResult.Error
	}

	return productStock, nil
}

func (repo *productStockRepository) findProductStockByProductID(productID string) (ProductStock, error) {
	var productStock ProductStock
	result := repo.DB.Where("product_id = ?", productID).Find(&productStock)
	return productStock, result.Error
}

func (repo *productStockRepository) findAllProductStock() ([]ProductStock, error) {
	var productStocks []ProductStock
	result := repo.DB.Find(&productStocks)
	return productStocks, result.Error
}
