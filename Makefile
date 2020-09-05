

boot_containers:
	@echo "booting up dependency containers..."
	docker-compose up -d
	@echo "containers up & running..."

destroy_containers:
	@echo "destroying dependency containers..."
	docker-compose down
	@echo "containers destroyed..."

build_app:
	@echo "building the app..."
	go build .
	@echo "App build..."

migration_up:
	@echo "running migrations..."
	./graphql-and-go migration up
	@echo "migrations ran successfully..."

migration_down:
	@echo "running migrations..."
	./graphql-and-go migration down
	@echo "migrations ran successfully..."

gen:
	@echo "generating gqlgen codes..."
	go generate ./...
	@echo "gqlgen codes generated..."

init: boot_containers gen build_app migration_up 
	@echo "App initialized"

run:
	./graphql-and-go serve


down: migration_down destroy_containers
	@echo "App is down"
