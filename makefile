vercel-prd:
	@echo "Deploying to Vercel production..."
	vercel --regions sin1 --prod

#Please run 'make tidy' before you first run the project
tidy:
	@echo "Running go mod tidy and vendor..."
	go mod tidy
	go clean -modcache
	go mod vendor
	@echo "All are good to go"

run:
	@echo "Running go server..."
	go run main.go

test:
	@echo "Running go tests..."
	go test ./... -v

swag:
	@echo "Generating swagger documentation..."
	swag init --parseDependency -g api/swagger.go
	@echo "Swagger documentation generated."

pre-commit:
	@echo "Running pre-commit hooks..."
	golangci-lint run --fix
	@echo "Pre-commit hooks completed."