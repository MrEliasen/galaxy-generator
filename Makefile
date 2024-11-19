.PHONY: help build server galaxy test coverage docker-build

default: help

VERSION ?= latest
REGISTRY = registry.ooga-booga.dev

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
	docker login registry.ooga-booga.dev
	docker build -t galaxy-generator:${VERSION} -f Dockerfile .
	docker tag galaxy-generator:${VERSION} ${REGISTRY}/galaxy-generator:${VERSION}
	docker push ${REGISTRY}/galaxy-generator:${VERSION}

build:
	cd src && go build -o ./release/ ./cmd/api/main.go

galaxy:
	cd src && go run ./cmd/generator/main.go

api :
	cd src && go run ./cmd/api/main.go
