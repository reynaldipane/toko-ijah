package purchasehistory

import (
	"github.com/jinzhu/gorm"
)

/*RepositoryInterface wraps the purchasehistory repository ability
so it can be easier to initialize the purchaseHistoryService with different kind of repository
*/
type RepositoryInterface interface {
	createPurchaseHistory(purchaseHistoryData PurchaseHistory) (PurchaseHistory, error)
}

type purchaseHistoryRepository struct {
	DB *gorm.DB
}

/*InitRepository will returns the repository object that required
to initialize purchaseHistoryService
*/
func InitRepository(db *gorm.DB) RepositoryInterface {
	return &purchaseHistoryRepository{
		DB: db,
	}
}

func (repo purchaseHistoryRepository) createPurchaseHistory(purchaseHistoryData PurchaseHistory) (PurchaseHistory, error) {
	result := repo.DB.Create(&purchaseHistoryData)
	return purchaseHistoryData, result.Error
}
