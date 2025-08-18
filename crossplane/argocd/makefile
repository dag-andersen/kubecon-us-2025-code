gen-secret:
	mkdir -p secrets
	@echo "Please enter your GitHub token:"
	@read github_token && cat secret_template.yaml | sed "s/GITHUB_TOKEN/$${github_token}/g" > secrets/secret.yaml

create-cluster:
	kind delete cluster --name crossplane || true
	kind create cluster --name crossplane

install-argocd:
	kubectl create namespace argocd
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	kubectl apply -f secrets/secret.yaml

install-crossplane:
	kubectl apply -f apps/crossplane.yaml
	kubectl apply -f apps/compositions.yaml
	kubectl apply -f apps/website.yaml

run: create-cluster install-argocd install-crossplane

local-render:	
	crossplane render app.yaml composition.yaml pkg.yaml