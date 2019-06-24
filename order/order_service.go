package order

import (
	"errors"
	"strconv"

	productstock "github.com/reynaldipane/toko-ijah/product_stock"
)

/*ServiceInterface wraps the order service ability
so it can be easier to initialize the orderHandler with different kind of service
*/
type ServiceInterface interface {
	createOrder(order Order) (Order, error)
}

/*Service is order service object that has access to the
order repo */
type Service struct {
	Repo                RepositoryInterface
	productStockService productstock.ServiceInterface
}

func (service *Service) createOrder(order Order) (Order, error) {
	productID := strconv.Itoa(int(order.ProductID))
	productStockData, err := service.productStockService.FindProductStockByProductID(productID)
	if err != nil {
		return order, err
	}

	isInvalidNumberOfOrder := productStockData.Stock-order.NumberSold < 0
	if isInvalidNumberOfOrder {
		return order, errors.New("Insufficient stock")
	}

	productStockData.Stock = productStockData.Stock - order.NumberSold

	_, err = service.productStockService.UpdateProductStock(productID, productStockData)
	if err != nil {
		return order, err
	}

	order.TotalPrice = order.NumberSold * order.SellPrice
	order.ProductStockID = productStockData.ID
	return service.Repo.createOrder(order)
}
