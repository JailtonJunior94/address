tests:
	go test -v --coverprofile test/coverage.out ./... 
	go tool cover -html=test/coverage.out

swagger:
	swag fmt
	swag init -g cmd/main.go

build-address-api:
	@echo "Compiling Address API..."
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o ./bin/address ./cmd/main.go