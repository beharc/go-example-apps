.PHONY: build-all run-all test-all

build-all: build-balance build-transfer

run-all: run-balance run-transfer

test-all: test-balance test-transfer

tidy-all: tidy-balance tidy-transfer

build-balance:
	cd services/balance && go build -o ../../bin/balance ./cmd

run-balance:
	cd services/balance && go run ./cmd

test-balance:
	cd services/balance && go test ./...

build-transfer:
	cd services/transfer && go build -o ../../bin/transfer ./cmd

run-transfer:
	cd services/transfer && go run ./cmd

test-transfer:
	cd services/transfer && go test ./...

tidy-transfer:
	cd services/transfer && go mod tidy

tidy-balance:
	cd services/balance && go mod tidy

docker-build-balance:
	docker build --build-arg SERVICE_NAME=balance -t balance-service .

docker-build-transfer:
	docker build --build-arg SERVICE_NAME=transfer -t transfer-service .

docker-build-all: docker-build-user docker-build-balance docker-build-transfer
