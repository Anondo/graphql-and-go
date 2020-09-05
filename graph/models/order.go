package models

// OrderProduct ...
type OrderProduct struct {
	ID       int     `json:"id"`
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// Order ...
type Order struct {
	ID        int            `json:"id"`
	Address   string         `json:"address"`
	DateAdded string         `json:"date_added"`
	Customer  User           `json:"customer"`
	Products  []OrderProduct `json:"products"`
}
