package transformers

import (
	dbmodels "github.com/Anondo/graphql-and-go/database/models"
	gmodels "github.com/Anondo/graphql-and-go/graph/models"
)

// TransformUsersToGraph ...
func TransformUsersToGraph(dbUsers []dbmodels.User) []gmodels.User {

	gUsers := []gmodels.User{}

	for _, user := range dbUsers {
		gUsers = append(gUsers, gmodels.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			PhoneNo:   user.PhoneNo,
		})
	}

	return gUsers
}

// TransformUserToGraph ...
func TransformUserToGraph(dbUser *dbmodels.User) *gmodels.User {
	return &gmodels.User{
		ID:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		PhoneNo:   dbUser.PhoneNo,
	}
}
