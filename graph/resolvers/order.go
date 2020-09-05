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
	opRepo := repos.NewOrderProductRepo(conn.Default())
	pRepo := repos.NewProductRepo(conn.Default())
	userRepo := repos.NewUserRepo(conn.Default())

	user, err := userRepo.GetUser(userID)
	if err != nil {
		return oo, err
	}

	if user == nil {
		return oo, nil
	}

	orders, err := oRepo.GetOrders(userID)
	if err != nil {
		return oo, err
	}

	oids := []int{}
	for _, order := range orders {
		oids = append(oids, order.ID)
	}

	ops, err := opRepo.GetOrderProducts(oids...)
	if err != nil {
		return oo, err
	}

	pids := []int{}
	for _, op := range ops {
		pids = append(pids, op.ProductID)
	}

	products, err := pRepo.GetProductsIn(pids...)
	if err != nil {
		return oo, err
	}

	pMap := map[int]models.Product{}
	opMap := map[int][]models.OrderProduct{}

	for _, product := range products {
		pMap[product.ID] = models.Product{
			ID:   product.ID,
			Name: product.Name,
			Type: product.Type,
		}
	}
	for _, op := range ops {
		if _, exists := opMap[op.OrderID]; !exists {
			opMap[op.OrderID] = []models.OrderProduct{}
		}

		opMap[op.OrderID] = append(opMap[op.OrderID], models.OrderProduct{
			ID:       op.ID,
			Product:  pMap[op.ProductID],
			Quantity: op.Quantity,
		})
	}

	for _, order := range orders {
		oo = append(oo, models.Order{
			ID: order.ID,
			Customer: models.User{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				PhoneNo:   user.PhoneNo,
			},
			Address:   order.Address,
			DateAdded: order.DateAdded.String(),
			Products:  opMap[order.ID],
		})
	}

	return oo, nil

}

func (r *queryResolver) Order(ctx context.Context, id int) (*models.Order, error) {

	orderRepo := repos.NewOrderRepo(conn.Default())
	userRepo := repos.NewUserRepo(conn.Default())
	orderProductRepo := repos.NewOrderProductRepo(conn.Default())
	productRepo := repos.NewProductRepo(conn.Default())

	order, err := orderRepo.GetOrder(id)
	if err != nil {
		return nil, err
	}

	if order == nil {
		return nil, nil
	}

	orderProducts, err := orderProductRepo.GetOrderProducts(order.ID)
	if err != nil {
		return nil, err
	}

	user, err := userRepo.GetUser(order.CustomerID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	pids := []int{}
	for _, op := range orderProducts {
		pids = append(pids, op.ProductID)
	}

	products, err := productRepo.GetProductsIn(pids...)
	if err != nil {
		return nil, err
	}

	pMap := map[int]models.Product{}
	ops := []models.OrderProduct{}

	for _, product := range products {
		pMap[product.ID] = models.Product{
			ID:   product.ID,
			Name: product.Name,
			Type: product.Type,
		}
	}
	for _, op := range orderProducts {
		ops = append(ops, models.OrderProduct{
			ID:       op.ID,
			Product:  pMap[op.ProductID],
			Quantity: op.Quantity,
		})
	}

	o := &models.Order{
		ID: order.ID,
		Customer: models.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			PhoneNo:   user.PhoneNo,
		},
		Address:   order.Address,
		DateAdded: order.DateAdded.String(),
		Products:  ops,
	}

	return o, nil
}
