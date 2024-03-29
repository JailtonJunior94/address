tests:
	go test -v --coverprofile test/coverage.out ./... 
	go tool cover -html=test/coverage.out

swagger:
	swag fmt
	swag init -g cmd/main.go

build-address-api:
	@echo "Compiling Address API..."
	@CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/address ./cmd/main.go

.PHONY: mockery
mock:
	@mockery --dir=internal/interfaces --name=CorreiosService --filename=correios_service_mock.go --output=internal/services/mocks --outpkg=serviceMocks
	@mockery --dir=internal/interfaces --name=HttpClient --filename=http_client_mock.go --output=internal/services/mocks --outpkg=serviceMocks
	@mockery --dir=internal/interfaces --name=ViaCepService --filename=via_cep_service_mock.go --output=internal/services/mocks --outpkg=serviceMocks
	@mockery --dir=pkg/logger --name=Logger --filename=logger_mock.go --output=pkg/logger/mocks --outpkg=loggerMocks