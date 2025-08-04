# UI
# kubectl -n flux-system port-forward svc/capacitor 9000:9000
# https://fluxcd.io/blog/2024/02/introducing-capacitor/

# kubectl tree -n default core.oam.dev/my-awesome-vela-app
# vs.
# flux tree kustomization kubevela-applications

bootstrap:
	kind delete cluster --name kubevela || true
	kind create cluster --name kubevela
	@echo "Please enter your GitHub token:"
	@read -s github_token && \
	export GITHUB_TOKEN=$$github_token && \
	flux bootstrap github \
		--owner=dag-andersen \
		--repository=kubecon-us-2025-code \
		--branch=main \
		--path=./kubevela/flux/clusters/kind \
		--personal