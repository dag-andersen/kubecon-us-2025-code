run:
	kind delete cluster --name crossplane || true
	kind create cluster --name crossplane
	helm repo add crossplane-preview https://charts.crossplane.io/preview
	sleep 5
	helm install crossplane \
		--namespace crossplane-system \
		--create-namespace crossplane-preview/crossplane \
		--version v2.0.0-preview.1

local-render:
	crossplane render app.yaml composition.yaml pkg.yaml