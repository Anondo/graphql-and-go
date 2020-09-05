package resolvers

import (
	"context"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/Anondo/graphql-and-go/graph/models"
)

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	userRepo := repos.NewUserRepo(conn.Default())
	umm := []models.User{}

	users, err := userRepo.GetUsers()
	if err != nil {
		return umm, err
	}

	for _, user := range users {
		umm = append(umm, models.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			PhoneNo:   user.PhoneNo,
		})
	}

	return umm, nil
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

	u := &models.User{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		PhoneNo:   user.PhoneNo,
	}

	return u, nil
}
