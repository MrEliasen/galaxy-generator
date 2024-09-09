include .env
.PHONY: help build dev-setup server galaxy test coverage

default: help

help:
	@echo "--------------------------------------------------------------------"
	@echo "| Command        | Description                                     |"
	@echo "--------------------------------------------------------------------"
	@echo "| build          | builds the project and outputs to /release      |"
	@echo "--------------------------------------------------------------------"
	@echo "| dev-setup      | Install required global go packages             |" 
	@echo "--------------------------------------------------------------------"
	@echo "| server         | runs the project web server                     |" 
	@echo "--------------------------------------------------------------------"
	@echo "| galaxy         | Generates a new galaxy        		          |" 
	@echo "--------------------------------------------------------------------"
	@echo "| test           | runs the project test files                     |" 
	@echo "--------------------------------------------------------------------"
	@echo "| cover          | outputs test cover to ./coverage.html           |" 
	@echo "--------------------------------------------------------------------"

build:
	CGO_ENABLED=1 go build -o ./release/ ./cmd/scraper/main.go

dev-setup:
	go install github.com/air-verse/air@latest	
	curl -sSf https://atlasgo.sh | sh
	go mod tidy

galaxy:
	go run ./cmd/generator/main.go

server:
	CGO_ENABLED=1 air 

test:
	go test ./...

coverage:
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out -o ./coverage.html
	open ./coverage.html

