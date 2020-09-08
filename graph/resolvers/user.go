package resolvers

import (
	"context"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/Anondo/graphql-and-go/graph/models"
	"github.com/Anondo/graphql-and-go/graph/transformers"
)

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	userRepo := repos.NewUserRepo(conn.Default())
	umm := []models.User{}

	users, err := userRepo.GetUsers()
	if err != nil {
		return umm, err
	}

	return transformers.TransformUsersToGraph(users), nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	userRepo := repos.NewUserRepo(conn.Default())

	user, err := userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return transformers.TransformUserToGraph(user), nil
}
