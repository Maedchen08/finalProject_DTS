lint:
	go fmt ./...

test:
	go test -v ./...

build:
	go build -v .

run:
	go run main.go

generate:
	mockgen -source=./repositories/transaction_repository.go -destination=./mocks/repository/transaction_mock.go -package=mocks
	mockgen -source=./repositories/agent_repository.go -destination=./mocks/repository/agent_mock.go -package=mocks
	mockgen -source=./repositories/customer_repository.go -destination=./mocks/repository/customer_mock.go -package=mocks
	mockgen -source=./repositories/authrepository.go -destination=./mocks/repository/auth_mock.go -package=mocks
	mockgen -source=./repositories/service_transaction_repository.go -destination=./mocks/repository/service_transaction_mock.go -package=mocks
	mockgen -source=./repositories/type_service_transaction_repository.go -destination=./mocks/repository/type_service_transaction_mock.go -package=mocks
	mockgen -source=./services/transaction_service.go -destination=./mocks/service/transaction_mock.go -package=mocks
	mockgen -source=./services/agent_service.go -destination=./mocks/service/agent_mock.go -package=mocks
	mockgen -source=./services/authservice.go -destination=./mocks/service/auth_mock.go -package=mocks
	mockgen -source=./services/customer_service.go -destination=./mocks/service/customer_mock.go -package=mocks
	mockgen -source=./services/serviceTransaction_service.go -destination=./mocks/service/serviceTransaction_mock.go -package=mocks
	mockgen -source=./services/type_service_transaction.go -destination=./mocks/service/type_service_transaction_mock.go -package=mocks

coverage-out:
	go tool cover -html=coverage.out

coverage:
	go test ./... -coverprofile coverage.out -count=1 -v



