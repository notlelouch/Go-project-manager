build:
	@go build -o bin/api
test:
	@go test -v ./...
run: build
	@./bin/api 