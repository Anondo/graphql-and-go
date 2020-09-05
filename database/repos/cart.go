package repos

import (
	"time"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/models"
	"github.com/jinzhu/gorm"
)

// CartRepo ...
type CartRepo struct {
	db *conn.DB
}

// NewCartRepo ...
func NewCartRepo(db *conn.DB) Cart {
	return &CartRepo{
		db: db,
	}
}

// AddToCart ...
func (c *CartRepo) AddToCart(userID, productID, quantity int) (int, error) {
	cart := models.Cart{
		CustomerID: userID,
		ProductID:  productID,
		Quantity:   quantity,
	}

	if err := c.db.Create(&cart).Error; err != nil {
		return 0, err
	}

	return cart.ID, nil
}

// GetCarts ...
func (c *CartRepo) GetCarts(userID int) ([]models.Cart, error) {
	carts := []models.Cart{}

	if err := c.db.Where("customer_id = ?", userID).Find(&carts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return carts, nil
		}
		return carts, err
	}

	return carts, nil

}

// Checkout ...
func (c *CartRepo) Checkout(userID int, address string) error {
	order := models.Order{
		CustomerID: userID,
		Address:    address,
		DateAdded:  time.Now().UTC(),
	}

	if err := c.db.Create(&order).Error; err != nil {
		return err
	}

	carts, err := c.GetCarts(userID)

	if err != nil {
		return err
	}

	for _, cart := range carts {
		orderProduct := models.OrderProduct{
			OrderID:   order.ID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
		}

		if err := c.db.Create(&orderProduct).Error; err != nil {
			return err
		}
	}

	return nil

}
