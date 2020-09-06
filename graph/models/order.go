package models

// OrderProduct ...
type OrderProduct struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// Order ...
type Order struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	DateAdded  string `json:"date_added"`
	CustomerID int    `json:"customer_id"`
}
