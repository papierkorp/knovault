# Variables
APP_NAME := knovault
THEMES_DIR := ./internal/themes
PLUGINS_DIR := ./internal/plugins

.PHONY: clean-plugins
clean-plugins:
	@echo "Cleaning up plugin artifacts..."
	@find $(PLUGINS_DIR)/core -name "internal" -type d -exec rm -rf {} +
	@find $(PLUGINS_DIR)/core -name "*.so" -type f -delete

.PHONY: clean-themes
clean-themes:
	@echo "Cleaning up theme artifacts..."
	@find $(THEMES_DIR)/core -name "internal" -type d -exec rm -rf {} +
	@find $(THEMES_DIR)/core -name "*.so" -type f -delete

.PHONY: templ-generate
templ-generate:
	TEMPL_EXPERIMENT=rawgo templ generate

.PHONY: build-themes
build-themes: clean-themes
	@echo "Building theme plugins..."
	@for theme in $(THEMES_DIR)/core/*; do \
		if [ -d "$$theme" ]; then \
			theme_name=$$(basename $$theme); \
			echo "Building $$theme_name..."; \
			cd $$theme && go build -buildmode=plugin -o $$theme_name.so main.go && cd -; \
		fi \
	done

.PHONY: build-plugins
build-plugins: clean-plugins
	@echo "Building plugins..."
	@for plugin in $(PLUGINS_DIR)/core/*; do \
		if [ -d "$$plugin" ]; then \
			plugin_name=$$(basename $$plugin); \
			echo "Building $$plugin_name..."; \
			cd $$plugin && go build -buildmode=plugin -o $$plugin_name.so main.go && cd -; \
		fi \
	done

.PHONY: ensure-dirs
ensure-dirs:
	@mkdir -p $(THEMES_DIR)/core
	@mkdir -p $(THEMES_DIR)/common
	@mkdir -p $(PLUGINS_DIR)/core
	@mkdir -p $(PLUGINS_DIR)/common

.PHONY: dev
dev: ensure-dirs templ-generate build-themes build-plugins
	go build -o ./bin/$(APP_NAME) ./cmd/main.go && air

.PHONY: build
build: ensure-dirs templ-generate build-themes build-plugins
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