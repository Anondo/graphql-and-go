package transformers

import (
	dbmodels "github.com/Anondo/graphql-and-go/database/models"
	gmodels "github.com/Anondo/graphql-and-go/graph/models"
)

// TransformOrdersToGraph ...
func TransformOrdersToGraph(dbOrders []dbmodels.Order) []gmodels.Order {
	gos := []gmodels.Order{}

	for _, order := range dbOrders {
		gos = append(gos, gmodels.Order{
			ID:         order.ID,
			CustomerID: order.CustomerID,
			Address:    order.Address,
			DateAdded:  order.DateAdded.String(),
		})
	}

	return gos
}

// TransformOrderToGraph ...
func TransformOrderToGraph(dbOrder *dbmodels.Order) *gmodels.Order {
	return &gmodels.Order{
		ID:         dbOrder.ID,
		CustomerID: dbOrder.CustomerID,
		Address:    dbOrder.Address,
		DateAdded:  dbOrder.DateAdded.String(),
	}
}

// TransformOrderProductsToGraph ...
func TransformOrderProductsToGraph(dbOPs []dbmodels.OrderProduct) []gmodels.OrderProduct {
	gOPs := []gmodels.OrderProduct{}

	for _, op := range dbOPs {
		gOPs = append(gOPs, gmodels.OrderProduct{
			ID:        op.ID,
			ProductID: op.ProductID,
			Quantity:  op.Quantity,
		})
	}

	return gOPs
}
