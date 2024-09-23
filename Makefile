.PHONY: help build server galaxy test coverage docker-build

default: help

help:
	@echo "--------------------------------------------------------------------"
	@echo "| Command        | Description                                     |"
	@echo "--------------------------------------------------------------------"
	@echo "| docker-build   | builds the project docker release image		  |"
	@echo "--------------------------------------------------------------------"
	@echo "| build          | builds the project and outputs to ./release     |"
	@echo "--------------------------------------------------------------------"
	@echo "| galaxy         | Runs the galaxy generator code                  |" 
	@echo "--------------------------------------------------------------------"
	@echo "| api            | runs the project web server                     |" 
	@echo "--------------------------------------------------------------------"

docker-build:
	docker build -t galaxy-generator:${VERSION} -f Dockerfile .

build:
	cd src && go build -o ./release/ ./cmd/api/main.go

galaxy:
	cd src && go run ./cmd/generator/main.go

api :
	cd src && go run ./cmd/api/main.go
