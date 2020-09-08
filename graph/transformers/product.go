package transformers

import (
	dbmodels "github.com/Anondo/graphql-and-go/database/models"
	gmodels "github.com/Anondo/graphql-and-go/graph/models"
)

// TransformProductsToGraph ...
func TransformProductsToGraph(dbProducts []dbmodels.Product) []gmodels.Product {
	gprdcts := []gmodels.Product{}

	for _, product := range dbProducts {
		gprdcts = append(gprdcts, gmodels.Product{
			ID:   product.ID,
			Name: product.Name,
			Type: product.Type,
		})
	}

	return gprdcts
}

// TransformProductToGraph ...
func TransformProductToGraph(dbProduct *dbmodels.Product) *gmodels.Product {
	return &gmodels.Product{
		ID:   dbProduct.ID,
		Name: dbProduct.Name,
		Type: dbProduct.Type,
	}
}
