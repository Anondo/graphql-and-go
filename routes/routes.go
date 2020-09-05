package routes

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Anondo/graphql-and-go/graph/generated"
	"github.com/Anondo/graphql-and-go/graph/resolvers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func init() {
	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
}

// Router list all the routes used by this app
func Router() *echo.Echo {

	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL Playground", "/query")))

	v1 := e.Group("/api/v1")

	graphHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	v1.POST("/query", echo.WrapHandler(graphHandler))

	return e
}
