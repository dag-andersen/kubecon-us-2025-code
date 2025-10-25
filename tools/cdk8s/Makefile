# CDK8s Multi-Environment Build System

.PHONY: help dev staging prod clean all

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

dev: ## Generate manifests for development environment
	@echo "🔨 Building development manifests..."
	go run . -env=development
	@echo "✅ Development manifests ready in dist/development/"

staging: ## Generate manifests for staging environment
	@echo "🔨 Building staging manifests..."
	go run . -env=staging
	@echo "✅ Staging manifests ready in dist/staging/"

prod: ## Generate manifests for production environment
	@echo "🔨 Building production manifests..."
	go run . -env=production
	@echo "✅ Production manifests ready in dist/production/"

all: dev staging prod ## Generate manifests for all environments

clean: ## Clean all generated manifests
	@echo "🧹 Cleaning generated manifests..."
	rm -rf dist/
	@echo "✅ All manifests cleaned"

# Quick deployment helpers (requires kubectl context to be set)
deploy-dev: dev ## Deploy development manifests to current kubectl context
	kubectl apply -f dist/development/

deploy-staging: staging ## Deploy staging manifests to current kubectl context
	kubectl apply -f dist/staging/

deploy-prod: prod ## Deploy production manifests to current kubectl context
	kubectl apply -f dist/production/

# Validation helpers
validate-dev: dev ## Validate development manifests
	kubectl --dry-run=client apply -f dist/development/

validate-staging: staging ## Validate staging manifests
	kubectl --dry-run=client apply -f dist/staging/

validate-prod: prod ## Validate production manifests
	kubectl --dry-run=client apply -f dist/production/

validate-all: all ## Validate all environment manifests
	kubectl --dry-run=client apply -f dist/development/
	kubectl --dry-run=client apply -f dist/staging/
	kubectl --dry-run=client apply -f dist/production/

# Show differences between environments
diff-staging-prod: staging prod ## Show differences between staging and production manifests
	@echo "Differences between staging and production:"
	diff -u dist/staging/ dist/production/ || true

diff-dev-staging: dev staging ## Show differences between development and staging manifests
	@echo "Differences between development and staging:"
	diff -u dist/development/ dist/staging/ || true
