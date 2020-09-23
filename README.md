# Graphql & Go

Trying to fit graphql with a project structure i often for Go projects wihtout making too much mess!

## Dependencies

1. **Docker Compose** - for `mysql` database required in the project(& other dependencies if needed)

1. [**gqlgen**](https://github.com/99designs/gqlgen) - The `Schema` & `Generate` based Go package for graphql which i find awesome

1. [**echo**](https://github.com/labstack/echo) - The router

1. [**viper**](https://github.com/spf13/viper) & [**cobra**](https://github.com/spf13/cobra) - For config management & command line interface

1. [**gorm**](https://github.com/jinzhu/gorm) - The ORM

## Running

Go to your terminal and hit

```console
  make init && make run
```

This should start the graphql server along with the graphql playground to play around with the API
