package purchase

import (
	"github.com/jinzhu/gorm"
)

type repositoryInterface interface {
	createPurchase(purchaseData Purchase) (Purchase, error)
	updateNumberReceivedByID(id string, updatePurchaseRepositoryData Purchase) (Purchase, error)
	findPurchasesDataByProductID(productID string) ([]Purchase, error)
	findPurchaseDataByID(id string) (Purchase, error)
}

type purchaseRepository struct {
	DB *gorm.DB
}

func initRepository(db *gorm.DB) repositoryInterface {
	return &purchaseRepository{
		DB: db,
	}
}

func (repo *purchaseRepository) createPurchase(purchaseData Purchase) (Purchase, error) {
	result := repo.DB.Create(&purchaseData)
	return purchaseData, result.Error
}

func (repo *purchaseRepository) findPurchasesDataByProductID(productID string) ([]Purchase, error) {
	var purchases []Purchase
	result := repo.DB.Where("product_id = ?", productID).Find(&purchases)
	return purchases, result.Error
}

func (repo *purchaseRepository) updateNumberReceivedByID(id string, updatePurchaseRepositoryData Purchase) (Purchase, error) {
	purchase := Purchase{}

	searchExistingPurchaseResult := repo.DB.First(&purchase, id)

	if searchExistingPurchaseResult.Error != nil {
		return updatePurchaseRepositoryData, searchExistingPurchaseResult.Error
	}

	purchase.NumberReceived = updatePurchaseRepositoryData.NumberReceived

	updateNumberReceivedResult := repo.DB.Save(&purchase)

	if updateNumberReceivedResult.Error != nil {
		return updatePurchaseRepositoryData, updateNumberReceivedResult.Error
	}

	return purchase, nil
}

func (repo *purchaseRepository) findPurchaseDataByID(id string) (Purchase, error) {
	var purchase Purchase
	result := repo.DB.First(&purchase, id)
	return purchase, result.Error
}
