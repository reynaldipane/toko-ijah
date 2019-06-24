package order

import (
	"github.com/jinzhu/gorm"
)

/*RepositoryInterface wraps the order repository ability
so it can be easier to initialize the orderService with different kind of repository
*/
type RepositoryInterface interface {
	createOrder(order Order) (Order, error)
}

type orderRepository struct {
	DB *gorm.DB
}

/*InitRepository will returns the repository object that required
to initialize orderService
*/
func InitRepository(db *gorm.DB) RepositoryInterface {
	return &orderRepository{
		DB: db,
	}
}

func (repo *orderRepository) createOrder(order Order) (Order, error) {
	result := repo.DB.Create(&order)
	return order, result.Error
}
