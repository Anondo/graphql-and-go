package resolvers

import (
	"context"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/Anondo/graphql-and-go/graph/models"
)

func (r *queryResolver) Orders(ctx context.Context, userID int) ([]models.Order, error) {
	oo := []models.Order{}

	oRepo := repos.NewOrderRepo(conn.Default())

	orders, err := oRepo.GetOrders(userID)
	if err != nil {
		return oo, err
	}

	for _, order := range orders {
		oo = append(oo, models.Order{
			ID:         order.ID,
			CustomerID: userID,
			Address:    order.Address,
			DateAdded:  order.DateAdded.String(),
		})
	}

	return oo, nil

}

func (r *queryResolver) Order(ctx context.Context, id int) (*models.Order, error) {

	orderRepo := repos.NewOrderRepo(conn.Default())

	order, err := orderRepo.GetOrder(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	o := &models.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Address:    order.Address,
		DateAdded:  order.DateAdded.String(),
	}

	return o, nil
}

func (r *orderResolver) Products(ctx context.Context, o *models.Order) ([]models.OrderProduct, error) {

	ops := []models.OrderProduct{}

	orderProducts, err := repos.NewOrderProductRepo(conn.Default()).GetOrderProducts(o.ID)
	if err != nil {
		return ops, err
	}

	for _, op := range orderProducts {
		ops = append(ops, models.OrderProduct{
			ID:        op.ID,
			ProductID: op.ProductID,
			Quantity:  op.Quantity,
		})
	}

	return ops, nil
}

func (r *orderResolver) Customer(ctx context.Context, o *models.Order) (*models.User, error) {

	user, err := repos.NewUserRepo(conn.Default()).GetUser(o.CustomerID)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	customer := &models.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PhoneNo:   user.PhoneNo,
	}

	return customer, nil
}

func (r *orderProductResolver) Product(ctx context.Context, op *models.OrderProduct) (*models.Product, error) {

	product, err := repos.NewProductRepo(conn.Default()).GetProduct(op.ProductID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	p := &models.Product{
		ID:   product.ID,
		Name: product.Name,
		Type: product.Type,
	}

	return p, nil
}
