# Variables
APP_NAME := knovault
THEMES_DIR := ./internal/themes


.PHONY: templ-generate
templ-generate:
	TEMPL_EXPERIMENT=rawgo templ generate

.PHONY: build-themes
build-themes:
	@echo "Building theme plugins..."
	@for theme in $(THEMES_DIR)/*; do \
		if [ -d "$$theme" ]; then \
			theme_name=$$(basename $$theme); \
			echo "Building $$theme_name..."; \
			go build -buildmode=plugin -o $$theme/$$theme_name.so $$theme/$$theme_name.go; \
		fi \
	done

.PHONY: dev
dev: build-themes templ-generate
	go build -o ./bin/$(APP_NAME) ./cmd/main.go && air

.PHONY: build
build: build-themes templ-generate
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

.PHONY: docker-build-base
docker-build-base:
	@if [ -z "$$(docker images -q knovault_base:latest 2> /dev/null)" ]; then \
		echo "Building knovault_base image..."; \
		docker build -t knovault_base:latest -f Dockerfile_base .; \
	else \
		echo "knovault_base image already exists. Skipping build."; \
	fi

.PHONY: docker-build-dev
docker-build-dev: docker-build-base
	docker build -t knovault:dev -f Dockerfile.dev .

.PHONY: docker-build-prod
docker-build-prod: docker-build-base
	docker build -t knovault:prod -f Dockerfile.prod .

.PHONY: docker-run-dev
docker-run-dev:
	docker run -d --name knovault-dev -p 1323:1323 -v $(PWD):/app knovault:dev

.PHONY: docker-run-prod
docker-run-prod:
	docker run -d --name knovault-prod -p 1323:1323 -v $(PWD)/data:/app/data knovault:prod

.PHONY: docker-stop
docker-stop:
	-docker stop knovault-dev
	-docker rm knovault-dev
	-docker stop knovault-prod
	-docker rm knovault-prod