gen-secret:
	mkdir -p secrets
	@echo "Please enter your GitHub token:"
	@read github_token && cat secret_template.yaml | sed "s/GITHUB_TOKEN/$${github_token}/g" > secrets/secret.yaml

create-cluster:
	kind delete cluster --name helm || true
	kind create cluster --name helm

install-argocd:
	kubectl create namespace argocd
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	kubectl apply -f secrets/secret.yaml

install-helm:
	kubectl apply -f apps/website.yaml

open-ui:
	argocd login --core
	kubectl config set-context --current --namespace=argocd
	open http://localhost:8082
	argocd admin dashboard --port 8082

run: create-cluster install-argocd install-helm open-ui

local-render:	
	helm template my-release ./my-chart