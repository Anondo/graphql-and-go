package models

// Cart ...
type Cart struct {
	ID         int `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	CustomerID int `json:"customer_id" gorm:"column:customer_id;not null"`
	ProductID  int `json:"product_id" gorm:"column:product_id;not null"`
	Quantity   int `json:"quantity" gorm:"column:quantity;not null;default:1"`
}

// TableName ...
func (c *Cart) TableName() string {
	return "cart"
}
