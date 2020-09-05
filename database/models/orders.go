package models

import "time"

type (
	// Order ...
	Order struct {
		ID         int       `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
		CustomerID int       `json:"customer_id" gorm:"column:customer_id;not null"`
		Address    string    `json:"address" gorm:"column:address;type:text"`
		DateAdded  time.Time `json:"date_added" gorm:"column:date_added"`
	}
	// OrderProduct ...
	OrderProduct struct {
		ID        int `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
		OrderID   int `json:"order_id" gorm:"column:order_id;not null"`
		ProductID int `json:"product_id" gorm:"column:product_id;not null"`
		Quantity  int `json:"quantity" gorm:"column:quantity;not null; default:1"`
	}
)

// TableName ...
func (o *Order) TableName() string {
	return "order"
}

// TableName ...
func (op *OrderProduct) TableName() string {
	return "order_product"
}
