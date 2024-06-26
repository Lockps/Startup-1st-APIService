build:
	@go build -buildvcs=false -o ./bin/Service ./cmd/api

run: build
	@./bin/Service

test:
	@go test -v ./cmd/test