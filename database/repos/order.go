package repos

import (
	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/models"
	"github.com/jinzhu/gorm"
)

// OrderRepo ...
type OrderRepo struct {
	db *conn.DB
}

// NewOrderRepo ...
func NewOrderRepo(db *conn.DB) Order {
	return &OrderRepo{
		db: db,
	}
}

// GetOrders ...
func (o *OrderRepo) GetOrders(userID int) ([]models.Order, error) {
	orders := []models.Order{}

	if err := o.db.Where("customer_id = ?", userID).Find(&orders).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return orders, nil
		}
		return orders, err
	}

	return orders, nil
}

// GetOrder ...
func (o *OrderRepo) GetOrder(orderID int) (*models.Order, error) {
	order := models.Order{}

	if err := o.db.Where("id = ?", orderID).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}
