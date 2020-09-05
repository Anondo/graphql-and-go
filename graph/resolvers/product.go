package resolvers

import (
	"context"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/Anondo/graphql-and-go/graph/models"
)

func (r *queryResolver) Products(ctx context.Context) ([]models.Product, error) {
	productRepo := repos.NewProductRepo(conn.Default())
	pmm := []models.Product{}

	products, err := productRepo.GetProducts()
	if err != nil {
		return pmm, err
	}

	for _, product := range products {
		pmm = append(pmm, models.Product{
			ID:   product.ID,
			Name: product.Name,
			Type: product.Type,
		})
	}

	return pmm, nil
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

	p := &models.Product{
		ID:   id,
		Name: product.Name,
		Type: product.Type,
	}

	return p, nil
}
