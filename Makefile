.PHONY: setup
setup: ## Downloads all docker dependencies
	docker pull postgres
	docker pull alpine

.PHONY: up
up: setup docker-backend-build ## Setup services and start them
	docker-compose up -d

.PHONY: down
down: ## Move services to down
	docker-compose down -v

.PHONY: build
build: ## Build the binary of backend
	go build -o ./BUILD/backend ./cmd/backend

docker-backend-build: build
	docker build -t angular-twitter -f cmd/backend/Dockerfile .

default: up ## It will do everything instead of you, need only for production

.PHONY: help
help: ## Display this help screen
	@echo "Commands:"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'