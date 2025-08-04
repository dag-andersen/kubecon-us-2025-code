# KubeVela Setup Options
# 
# Choose between ArgoCD or Flux for GitOps management:

argocd:
	@echo "Setting up KubeVela with ArgoCD..."
	cd argocd && make gen-secret && make run

flux:
	@echo "Setting up KubeVela with Flux..."
	cd flux && make bootstrap

help:
	@echo "Available commands:"
	@echo "  make argocd  - Deploy KubeVela using ArgoCD"
	@echo "  make flux    - Deploy KubeVela using Flux"
	@echo "  make help    - Show this help message"