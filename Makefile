# Variables
APP_NAME := pewito
THEMES_DIR := ./internal/themes

.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./static/css/main.css -o ./static/css/output.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i ./static/css/main.css -o ./static/css/output.css --minify

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
dev: build-themes templ-generate tailwind-build
	go build -o ./bin/$(APP_NAME) ./cmd/main.go && air

.PHONY: build
build: build-themes tailwind-build templ-generate
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

.PHONY: docker-build-base
docker-build-base:
	@if [ -z "$$(docker images -q pewito_base:latest 2> /dev/null)" ]; then \
		echo "Building pewito_base image..."; \
		docker build -t pewito_base:latest -f Dockerfile_base .; \
	else \
		echo "pewito_base image already exists. Skipping build."; \
	fi

.PHONY: docker-build-dev
docker-build-dev: docker-build-base
	docker build -t pewito:dev -f Dockerfile.dev .

.PHONY: docker-build-prod
docker-build-prod: docker-build-base
	docker build -t pewito:prod -f Dockerfile.prod .

.PHONY: docker-run-dev
docker-run-dev:
	docker run -d --name pewito-dev -p 1323:1323 -v $(PWD):/app pewito:dev

.PHONY: docker-run-prod
docker-run-prod:
	docker run -d --name pewito-prod -p 1323:1323 -v $(PWD)/data:/app/data pewito:prod

.PHONY: docker-stop
docker-stop:
	-docker stop pewito-dev
	-docker rm pewito-dev
	-docker stop pewito-prod
	-docker rm pewito-prod