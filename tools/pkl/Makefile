.PHONY: help install generate-all generate-prod generate-staging clean

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

install: ## Install Pkl (macOS)
	@echo "Installing Pkl..."
	@if command -v pkl > /dev/null; then \
		echo "Pkl is already installed: $$(pkl --version)"; \
	else \
		echo "Installing Pkl via Homebrew..."; \
		brew install pkl; \
	fi

generate-all: generate-prod generate-staging ## Generate manifests for all environments

generate-prod: ## Generate production manifests
	@echo "Generating production manifests..."
	@mkdir -p dist/production
	@pkl eval -f yaml prod.pkl > dist/production/manifests.yaml
	@echo "✓ Production manifests generated at dist/production/manifests.yaml"

generate-staging: ## Generate staging manifests
	@echo "Generating staging manifests..."
	@mkdir -p dist/staging
	@pkl eval -f yaml staging.pkl > dist/staging/manifests.yaml
	@echo "✓ Staging manifests generated at dist/staging/manifests.yaml"

clean: ## Remove generated files
	@echo "Cleaning up..."
	@rm -rf dist/
	@echo "✓ Cleaned up generated files"

validate: ## Validate Pkl files
	@echo "Validating Pkl files..."
	@pkl eval prod.pkl > /dev/null && echo "✓ prod.pkl is valid"
	@pkl eval staging.pkl > /dev/null && echo "✓ staging.pkl is valid"

diff: ## Show differences between prod and staging
	@echo "Differences between production and staging:"
	@mkdir -p dist/production dist/staging
	@pkl eval -f yaml prod.pkl > dist/production/manifests.yaml
	@pkl eval -f yaml staging.pkl > dist/staging/manifests.yaml
	@diff -u dist/staging/manifests.yaml dist/production/manifests.yaml || true

