#! /bin/sh

echo "building gng(graphql-and-go)..."
export GO111MODULE=on
CGO_ENABLED=0 GOFLAGS=-mod=vendor go build -o gng
echo "gng built successfully!"

echo "running migrations..."
./gng migration up
echo "Migrations ran successfully!"
./gng serve
