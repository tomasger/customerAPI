install:
	go build -o customer_api cmd/main.go

test:
    go test ./...