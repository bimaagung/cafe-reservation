build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

gen-swagger:
	swag init -g cmd/main.go

test:
	go test -v ./...