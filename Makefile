.PHONY: help build server galaxy test coverage docker-build

default: help

help:
	@echo "--------------------------------------------------------------------"
	@echo "| Command        | Description                                     |"
	@echo "--------------------------------------------------------------------"
	@echo "| docker-build   | builds the project docker image		          |"
	@echo "--------------------------------------------------------------------"
	@echo "| build          | builds the project and outputs to /release      |"
	@echo "--------------------------------------------------------------------"
	@echo "| galaxy         | Generates a new galaxy        		          |" 
	@echo "--------------------------------------------------------------------"
	@echo "| server         | runs the project web server                     |" 
	@echo "--------------------------------------------------------------------"
	@echo "| test           | runs the project test files                     |" 
	@echo "--------------------------------------------------------------------"

docker-build:
	docker build -t galaxy-generator:${VERSION} -f Dockerfile .

build-generator:
	cd src && go build -o ./release/ ./cmd/generator/main.go

build-api:
	cd src && GIN_MODE=release go build -o ./release/ ./cmd/api/main.go

galaxy:
	cd src && go run ./cmd/generator/main.go

server:
	cd src && go run ./cmd/api/main.go

test:
	cd src && go test ./...
