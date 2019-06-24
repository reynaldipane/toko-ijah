package purchase

import (
	"errors"
	"strconv"

	productstock "github.com/reynaldipane/test-ijah/product_stock"
	purchasehistory "github.com/reynaldipane/test-ijah/purchase_history"
)

type serviceInterface interface {
	createPurchase(purchaseData Purchase) (Purchase, error)
	updateNumberReceivedByID(id string, updatePurchaseData Purchase) (Purchase, error)
}

/*Service is the purchase's service object
that have access to the product repository*/
type Service struct {
	repo                   repositoryInterface
	productStockService    productstock.ServiceInterface
	purchaseHistoryService purchasehistory.ServiceInterface
}

func (service *Service) createPurchase(purchaseData Purchase) (Purchase, error) {
	productID := strconv.Itoa(int(purchaseData.ProductID))

	productStockData, err := service.productStockService.FindProductStockByProductID(productID)
	if err != nil {
		return purchaseData, err
	}

	isInvalidNumberReceivedPurchase := purchaseData.NumberOrdered < purchaseData.NumberReceived
	if isInvalidNumberReceivedPurchase {
		return purchaseData, errors.New("Incorrect number received, check again")
	}

	purchaseData.TotalPrice = purchaseData.NumberOrdered * purchaseData.BuyPrice

	createdPurchase, err := service.repo.createPurchase(purchaseData)
	if err != nil {
		return createdPurchase, err
	}

	_, err = service.purchaseHistoryService.CreatePurchaseHistory(purchasehistory.PurchaseHistory{
		PurchaseID:     createdPurchase.ID,
		NumberReceived: createdPurchase.NumberReceived,
	})

	if err != nil {
		return purchaseData, err
	}

	err = service.processUpdateProductStock(productID, productStockData, createdPurchase)

	if err != nil {
		return purchaseData, err
	}

	return createdPurchase, nil
}

func (service *Service) updateNumberReceivedByID(id string, updatePurchaseData Purchase) (Purchase, error) {
	oldPurchaseData, err := service.repo.findPurchaseDataByID(id)
	if err != nil {
		return updatePurchaseData, err
	}

	isInvalidNumberReceived := (oldPurchaseData.NumberReceived+updatePurchaseData.NumberReceived > oldPurchaseData.NumberOrdered) ||
		(oldPurchaseData.NumberReceived <= updatePurchaseData.NumberReceived)

	if isInvalidNumberReceived {
		return updatePurchaseData, errors.New("Incorrect update number received, check again")
	}

	productID := strconv.Itoa(int(oldPurchaseData.ProductID))

	productStockData, err := service.productStockService.FindProductStockByProductID(productID)
	if err != nil {
		return updatePurchaseData, err
	}

	updatePurchaseData.NumberReceived = oldPurchaseData.NumberReceived + updatePurchaseData.NumberReceived
	updatedPurchaseData, err := service.repo.updateNumberReceivedByID(id, updatePurchaseData)
	if err != nil {
		return updatePurchaseData, err
	}

	err = service.processUpdateProductStock(productID, productStockData, updatedPurchaseData)
	if err != nil {
		return updatePurchaseData, err
	}

	_, err = service.purchaseHistoryService.CreatePurchaseHistory(purchasehistory.PurchaseHistory{
		PurchaseID:     updatedPurchaseData.ID,
		NumberReceived: updatePurchaseData.NumberReceived - oldPurchaseData.NumberReceived,
	})

	if err != nil {
		return updatePurchaseData, err
	}

	return updatedPurchaseData, nil
}

func calculateProductStockDetails(purchases []Purchase) (float64, int, int) {
	sumTotalPrice := 0
	sumNumberReceived := 0
	sumNumberOrdered := 0

	for _, v := range purchases {
		sumTotalPrice = sumTotalPrice + v.TotalPrice
		sumNumberReceived = sumNumberReceived + v.NumberReceived
		sumNumberOrdered = sumNumberOrdered + v.NumberOrdered
	}

	averageBuyPrice := float64(sumTotalPrice / sumNumberReceived)

	return averageBuyPrice, sumNumberReceived, sumNumberOrdered
}

func (service *Service) processUpdateProductStock(
	productID string,
	productStockData productstock.ProductStock,
	purchaseData Purchase,
) error {
	purchasesData, err := service.repo.findPurchasesDataByProductID(productID)

	if err != nil {
		return err
	}

	productStockData.AverageBuyPrice, productStockData.Stock, productStockData.TotalOrdered = calculateProductStockDetails(purchasesData)

	_, err = service.productStockService.UpdateProductStock(productID, productStockData)

	return err
}
