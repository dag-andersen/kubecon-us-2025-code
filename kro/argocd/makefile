gen-secret:
	mkdir -p secrets
	@echo "Please enter your GitHub token:"
	@read github_token && cat secret_template.yaml | sed "s/GITHUB_TOKEN/$${github_token}/g" > secrets/secret.yaml

run:
	kind delete cluster --name kro || true
	kind create cluster --name kro
	# install argocd
	kubectl create namespace argocd
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
	kubectl apply -f secrets/secret.yaml
	kubectl apply -f apps/kro.yaml
	kubectl apply -f apps/kro-definitions.yaml
	kubectl apply -f apps/applications.yaml
