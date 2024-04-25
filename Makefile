build:
	@go build -o bin/app .

run: build
	@./bin/app

test:
	@go test -v ./...

mod:
	@go mod tidy

gen:
	@templ generate

.PHONY: build run test mod gen

