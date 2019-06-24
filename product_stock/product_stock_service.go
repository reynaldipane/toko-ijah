package productstock

/*ServiceInterface wraps the productstock service ability
so it can be easier to initialize the productHandler with different kind of service
*/
type ServiceInterface interface {
	CreateProductStock(ps ProductStock) (ProductStock, error)
	UpdateProductStock(productID string, updateProductStockData ProductStock) (ProductStock, error)
	FindProductStockByProductID(productID string) (ProductStock, error)
	FindAllProductStock() ([]ProductStock, error)
}

/*Service is product stock service object that has access to the
product stock repo */
type Service struct {
	Repo RepositoryInterface
}

/*CreateProductStock wraps Repo method to create the stock
and additional business logic will added here as well
*/
func (service *Service) CreateProductStock(ps ProductStock) (ProductStock, error) {
	return service.Repo.createProductStock(ps)
}

/*UpdateProductStock wraps Repo method to update the stock
and additional business logic will added here as well
*/
func (service *Service) UpdateProductStock(productID string, updateProductStockData ProductStock) (ProductStock, error) {
	return service.Repo.updateProductStock(productID, updateProductStockData)
}

/*FindProductStockByProductID wraps Repo method to find the product's stock
by product id, and additional business logic will added here as well
*/
func (service *Service) FindProductStockByProductID(productID string) (ProductStock, error) {
	return service.Repo.findProductStockByProductID(productID)
}

/*FindAllProductStock wraps Repo method to find all the product's stock
, and additional business logic will added here as well
*/
func (service *Service) FindAllProductStock() ([]ProductStock, error) {
	return service.Repo.findAllProductStock()
}
