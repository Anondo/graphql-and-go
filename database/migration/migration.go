package migration

import "github.com/Anondo/graphql-and-go/database/models"

var (
	// Models ...
	Models []interface{}
)

func init() {
	Models = append(Models, &models.User{})
	Models = append(Models, &models.Product{})
	Models = append(Models, &models.Cart{})
	Models = append(Models, &models.Order{})
	Models = append(Models, &models.OrderProduct{})
}
