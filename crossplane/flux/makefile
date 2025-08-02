
bootstrap:
	kind delete cluster --name crossplane || true
	kind create cluster --name crossplane
	@echo "Please enter your GitHub token:"
	@read -s github_token && \
	export GITHUB_TOKEN=$$github_token && \
	flux bootstrap github \
		--owner=dag-andersen \
		--repository=kubecon-us-2025-code \
		--branch=main \
		--path=./crossplane/flux/clusters/kind \
		--personal
	kubectl apply -f apps/crossplane.yaml
	kubectl apply -f apps/compositions.yaml
	kubectl apply -f apps/website.yaml

# run-manual:
# 	kind delete cluster --name crossplane || true
# 	kind create cluster --name crossplane
# 	# install flux manually
# 	flux install
# 	# apply secrets and resources manually
# 	kubectl apply -f secrets/secret.yaml
# 	kubectl apply -f auth/gitrepository.yaml
# 	kubectl apply -f apps/crossplane.yaml
# 	kubectl apply -f apps/compositions.yaml
# 	kubectl apply -f apps/website.yaml

# gen-secret:
# 	mkdir -p secrets
# 	@echo "Please enter your GitHub token:"
# 	@read github_token && cat secret_template.yaml | sed "s/GITHUB_TOKEN/$${github_token}/g" > secrets/secret.yaml

local-render:
	crossplane render app.yaml composition.yaml pkg.yaml
