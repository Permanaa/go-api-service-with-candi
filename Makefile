.PHONY : build run

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

proto:
	$(foreach proto_file, $(shell find api/proto -name '*.proto'),\
	protoc --proto_path=api/proto --go_out=plugins=grpc:api/proto \
	--go_opt=paths=source_relative $(proto_file);)

migration:
	@go run cmd/migration/migration.go $(call args,up)

build:
	@go build -o bin

run: build
	@./bin

docker:
	docker build -t go-api-service-with-candi:latest .

run-container:
	docker run --rm --name=go-api-service-with-candi --network="host" -d go-api-service-with-candi

clear-docker:
	docker rm -f go-api-service-with-candi
	docker rmi -f go-api-service-with-candi

# mocks all interfaces for unit test
mocks:
	@mockery --all --keeptree --dir=internal --output=pkg/mocks --case underscore
	@mockery --all --keeptree --dir=pkg --output=pkg/mocks --case underscore

# unit test & calculate code coverage
test:
	@echo "\x1b[32;1m>>> running unit test and calculate coverage\x1b[0m"
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count \
		-coverpkg=$$(go list ./... | grep -v mocks | tr '\n' ',')
	@go tool cover -func=coverage.txt
