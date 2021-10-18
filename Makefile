lint:
	go fmt ./...

test:
	go test -v ./...

build:
	go build -v .

run:
	go run main.go

generate-mock:
	go generate -v ./... 

generate:
	mockgen -source=./repositories/transaction_repository.go -destination=./mocks/repository/transaction_mock.go -package=mocks
	mockgen -source=./services/transaction_service.go -destination=./mocks/service/transaction_mock.go -package=mocks


