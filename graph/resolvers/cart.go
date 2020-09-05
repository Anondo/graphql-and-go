package resolvers

import (
	"context"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/Anondo/graphql-and-go/graph/models"
)

func (r *queryResolver) Cart(ctx context.Context, userID int) (*models.Cart, error) {
	cartRepo := repos.NewCartRepo(conn.Default())

	carts, err := cartRepo.GetCarts(userID)
	if err != nil {
		return nil, err
	}

	productRepo := repos.NewProductRepo(conn.Default())
	pids := []int{}

	for _, cart := range carts {
		pids = append(pids, cart.ProductID)
	}

	products, err := productRepo.GetProductsIn(pids...)
	if err != nil {
		return nil, err
	}

	proMap := map[int]models.Product{}

	for _, product := range products {
		proMap[product.ID] = models.Product{
			ID:   product.ID,
			Name: product.Name,
			Type: product.Type,
		}
	}

	c := &models.Cart{
		CustomerID:    userID,
		CartLineItems: []models.CartLineItem{},
	}

	for _, cart := range carts {
		c.CartLineItems = append(c.CartLineItems, models.CartLineItem{
			ID:       cart.ID,
			Product:  proMap[cart.ProductID],
			Quantity: cart.Quantity,
		})
	}

	return c, nil

}

func (r *mutationResolver) AddToCart(ctx context.Context, cartItem models.NewCartItem) (*models.CartLineItem, error) {

	productRepo := repos.NewProductRepo(conn.Default())
	cartRepo := repos.NewCartRepo(conn.Default())

	product, err := productRepo.GetProduct(cartItem.ProductID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	cid, err := cartRepo.AddToCart(cartItem.CustomerID, cartItem.ProductID, cartItem.Quantity)
	if err != nil {
		return nil, err
	}

	cl := models.CartLineItem{
		ID:       cid,
		Quantity: cartItem.Quantity,
		Product: models.Product{
			ID:   product.ID,
			Name: product.Name,
			Type: product.Type,
		},
	}

	return &cl, nil

}
