build:
	@go build -o bin/cas

run: build
	@ ./bin/cas

test:
	@go test ./... -v