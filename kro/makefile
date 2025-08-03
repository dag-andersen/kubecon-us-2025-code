run:
	kind delete cluster --name kro || true
	kind create cluster --name kro
	helm install kro oci://ghcr.io/kro-run/kro/kro \
		--namespace kro \
		--create-namespace \
		--version=0.4.0
	helm -n kro list
	kubectl apply -f resource-graph-definition.yaml
	kubectl get rgd -o wide