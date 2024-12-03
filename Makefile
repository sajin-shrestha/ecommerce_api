# command to build the GO application
build:
	@go build -o bin/ecommerce_api main.go
	@echo "Build Complete."

# command to run the GO application
run:
	@go run main.go
