.PHONY: build-all run-all test-all

build-all: build-user build-balance build-transfer

run-all: run-user run-balance run-transfer

test-all: test-user test-balance test-transfer

build-user:
	cd services/user && go build -o ../../bin/user ./cmd

run-user:
	cd services/user && go run ./cmd

test-user:
	cd services/user && go test ./...

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

docker-build-user:
	docker build --build-arg SERVICE_NAME=user -t user-service .

docker-build-balance:
	docker build --build-arg SERVICE_NAME=balance -t balance-service .

docker-build-transfer:
	docker build --build-arg SERVICE_NAME=transfer -t transfer-service .

docker-build-all: docker-build-user docker-build-balance docker-build-transfer
