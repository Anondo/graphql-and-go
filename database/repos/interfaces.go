package repos

import "github.com/Anondo/graphql-and-go/database/models"

// User ...
type User interface {
	GetUsers() ([]models.User, error)
	GetUser(int) (*models.User, error)
}

// Product ...
type Product interface {
	GetProducts() ([]models.Product, error)
	GetProduct(int) (*models.Product, error)
	GetProductsIn(...int) ([]models.Product, error)
}

// Cart ...
type Cart interface {
	AddToCart(int, int, int) (int, error)
	GetCarts(int) ([]models.Cart, error)
	Checkout(int, string) error
}

// Order ...
type Order interface {
	GetOrders(int) ([]models.Order, error)
	GetOrder(int) (*models.Order, error)
}

// OrderProduct ...
type OrderProduct interface {
	GetOrderProducts(...int) ([]models.OrderProduct, error)
}
