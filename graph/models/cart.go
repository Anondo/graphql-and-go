package models

// CartLineItem ...
type CartLineItem struct {
	ID       int     `json:"id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// Cart ...
type Cart struct {
	CustomerID    int            `json:"customer_id"`
	CartLineItems []CartLineItem `json:"cart_line_items"`
}

// NewCartItem ...
type NewCartItem struct {
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

// CheckoutInfo ...
type CheckoutInfo struct {
	CustomerID int    `json:"customer_id"`
	Address    string `json:"address"`
}
