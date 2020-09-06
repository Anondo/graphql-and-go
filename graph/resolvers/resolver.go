//go:generate go run github.com/99designs/gqlgen --verbose
package resolvers

import "github.com/Anondo/graphql-and-go/graph/generated"

// Resolver ...
type Resolver struct {
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query ...
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// Order ...
func (r *Resolver) Order() generated.OrderResolver {
	return &orderResolver{r}
}

// OrderProduct ...
func (r *Resolver) OrderProduct() generated.OrderProductResolver {
	return &orderProductResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type orderProductResolver struct{ *Resolver }
