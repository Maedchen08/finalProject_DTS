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
	mockgen -source=./services/transaction_service.go -destination=./mocks/service/transaction_mock.go -package=mocks
	mockgen -source=./services/agent_service.go -destination=./mocks/service/agent_mock.go -package=mocks

coverage-out:
	go tool cover -html=coverage.out

coverage:
	go test ./... -coverprofile coverage.out -count=1 -v



