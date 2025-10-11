cluster-name := kubecon

gen-secret:
	mkdir -p secrets
	@echo "Please enter your GitHub token:"
	@read github_token && cat secret_template.yaml | sed "s/GITHUB_TOKEN/$${github_token}/g" > secrets/secret.yaml

create-cluster:
	kind delete cluster --name ${cluster-name} || true
	kind create cluster --name ${cluster-name}

install-argocd:
	kubectl create namespace argocd || true
	helm repo add argo https://argoproj.github.io/argo-helm
	helm repo update
	helm upgrade --install argocd argo/argo-cd -n argocd -f helm-values.yaml
	kubectl apply -f secrets/secret.yaml

install-helm:
	kubectl apply -f helm/apps

install-kustomize:
	kubectl apply -f kustomize/apps

install-kro:
	kubectl apply -f kro/apps

install-crossplane:
	kubectl apply -f crossplane/apps

install-kubevela:
	kubectl apply -f kubevela/apps

open-ui:
	argocd login --core
	kubectl config set-context --current --namespace=argocd
	kubectl wait deployment argocd-server \
		--for=condition=Available=True \
		--namespace=argocd \
		--timeout=120s
	open http://localhost:8082
	argocd admin dashboard --port 8082

run: create-cluster install-argocd open-ui
