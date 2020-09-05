package repos

import (
	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/models"
	"github.com/jinzhu/gorm"
)

// OrderProductRepo ...
type OrderProductRepo struct {
	db *conn.DB
}

// NewOrderProductRepo ...
func NewOrderProductRepo(db *conn.DB) OrderProduct {
	return &OrderProductRepo{
		db: db,
	}
}

// GetOrderProducts ...
func (op *OrderProductRepo) GetOrderProducts(orderIDs ...int) ([]models.OrderProduct, error) {
	orderProducts := []models.OrderProduct{}

	if err := op.db.Where("order_id IN (?)", orderIDs).Find(&orderProducts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return orderProducts, nil
		}
		return orderProducts, err
	}

	return orderProducts, nil
}
