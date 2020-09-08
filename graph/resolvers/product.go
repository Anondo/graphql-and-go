package resolvers

import (
	"context"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/Anondo/graphql-and-go/graph/models"
	"github.com/Anondo/graphql-and-go/graph/transformers"
)

func (r *queryResolver) Products(ctx context.Context) ([]models.Product, error) {
	productRepo := repos.NewProductRepo(conn.Default())
	pmm := []models.Product{}

	products, err := productRepo.GetProducts()
	if err != nil {
		return pmm, err
	}

	return transformers.TransformProductsToGraph(products), nil
}

func (r *queryResolver) Product(ctx context.Context, id int) (*models.Product, error) {
	productRepo := repos.NewProductRepo(conn.Default())

	product, err := productRepo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	return transformers.TransformProductToGraph(product), nil
}
