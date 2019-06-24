package purchasehistory

/*ServiceInterface wraps the purchasehistory service ability
so it can be easier to initialize the purchaseHandler with different kind of service
*/
type ServiceInterface interface {
	CreatePurchaseHistory(purchaseHistoryData PurchaseHistory) (PurchaseHistory, error)
}

/*Service is purchase history service object that has access to the
product stock repo */
type Service struct {
	Repo RepositoryInterface
}

/*CreatePurchaseHistory wraps Repo method to create the purchase history
and additional business logic will added here as well
*/
func (service *Service) CreatePurchaseHistory(purchaseHistoryData PurchaseHistory) (PurchaseHistory, error) {
	return service.Repo.createPurchaseHistory(purchaseHistoryData)
}
