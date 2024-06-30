build:
	@go build -o bin/ecom-go cmd/main.go

test:
	go test -cover  ./...

run: build
	@./bin/ecom-go