package middlewares

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/repos"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware ...
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		userID, _ := strconv.Atoi(c.Request().Header.Get("User-id"))

		repo := repos.NewUserRepo(conn.Default())

		user, err := repo.GetUser(userID)
		if err != nil {
			log.Println("Failed to fetch user info for authentication: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "something went wrong",
			})
		}

		if user == nil {
			log.Printf("User of user-id:%d is not authorized\n", userID)
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "User not authorized",
			})
		}

		return next(c)
	}
}
