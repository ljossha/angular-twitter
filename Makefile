.PHONY: setup
setup:
	docker pull postgres
	docker pull alpine

.PHONY: up
up: setup docker-backend-build
	docker-compose up -d

.PHONY: down
down:
	docker-compose down -v

.PHONY: build
build:
	go build -o ./BUILD/backend ./cmd/backend

docker-backend-build: build
	docker build -t angular-twitter -f cmd/backend/Dockerfile .

.PHONY: help
help:
	@echo "How to use Makefile"
	@echo ""
	@echo "make					- It will do everything instead of you, need only for production"
	@echo "make setup           - Downloads all docker dependencies"
	@echo "make build			- Build the binary of backend"
	@echo ""

default: up