cluster-name := kubecon


# UI
# kubectl -n flux-system port-forward svc/capacitor 9000:9000
# https://fluxcd.io/blog/2024/02/introducing-capacitor/

# kubectl tree -n default applications.kro.run/my-application-instance
# vs.
# flux tree kustomization kro-applications

create-cluster:
	kind delete cluster --name ${cluster-name} || true
	kind create cluster --name ${cluster-name}

install-flux:
	@echo "Please enter your GitHub token:"
	@read -s github_token && \
	export GITHUB_TOKEN=$$github_token && \
	flux bootstrap github \
		--owner=dag-andersen \
		--repository=kubecon-us-2025-code \
		--branch=main \
		--path=./flux/clusters/kind \
		--personal

install-kro:
	kubectl apply -f kro/apps

run: create-cluster install-flux
