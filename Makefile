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

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

proxy:
	templ generate --watch --proxy=http://192.168.206.1:9001

.PHONY: build run test mod gen

