vercel-prd:
	@echo "Deploying to Vercel production..."
	vercel --prod

run:
	@echo "Running go server..."
	go run main.go

test:
	@echo "Running go tests..."
	go test ./... -v

pre-commit:
	@echo "Running pre-commit hooks..."
	golangci-lint run --fix
	@echo "Pre-commit hooks completed."

tidy:
	@echo "Running go mod tidy and vendor..."
	go mod tidy
	go clean -modcache
	go mod vendor
	@echo "All are good to go"