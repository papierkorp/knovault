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