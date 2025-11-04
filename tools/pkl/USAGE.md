# Pkl Usage Examples

## Quick Start

```bash
# Generate manifests for production
make generate-prod

# Generate manifests for staging
make generate-staging

# Generate all environments
make generate-all
```

## Direct Pkl Commands

### Evaluate and view output
```bash
pkl eval prod.pkl
```

### Generate YAML
```bash
pkl eval -f yaml prod.pkl
```

### Generate JSON instead
```bash
pkl eval -f json prod.pkl
```

### Save to file
```bash
pkl eval -f yaml prod.pkl > manifests.yaml
```

## Validate Before Generating

```bash
# Check if configuration is valid
make validate

# Or directly with Pkl
pkl eval prod.pkl > /dev/null && echo "✓ Valid"
```

## Apply to Kubernetes

### Apply production
```bash
pkl eval -f yaml prod.pkl | kubectl apply -f -
```

### Apply staging
```bash
pkl eval -f yaml staging.pkl | kubectl apply -f -
```

### Dry run first
```bash
pkl eval -f yaml prod.pkl | kubectl apply -f - --dry-run=client
```

### Apply to specific context
```bash
pkl eval -f yaml prod.pkl | kubectl apply -f - --context=production-cluster
```

## View Differences

### Compare with current cluster state
```bash
# Generate and diff with production cluster
pkl eval -f yaml prod.pkl | kubectl diff -f -
```

### Compare environments
```bash
make diff
```

### Compare specific resources
```bash
diff <(pkl eval -f yaml prod.pkl | yq e 'select(.kind == "Deployment")' -) \
     <(pkl eval -f yaml staging.pkl | yq e 'select(.kind == "Deployment")' -)
```

## Customize Output

### Pretty print
```bash
pkl eval -f yaml prod.pkl | bat -l yaml
# or
pkl eval -f yaml prod.pkl | pygmentize -l yaml
```

### Extract specific resource
```bash
# Get only the Deployment
pkl eval -f yaml prod.pkl | yq e 'select(.kind == "Deployment")' -

# Get only the Service
pkl eval -f yaml prod.pkl | yq e 'select(.kind == "Service")' -
```

## Advanced: Override Values

You can override values at evaluation time:

```bash
# Override replicas
pkl eval -f yaml prod.pkl -p 'replicas=10'

# Override image tag
pkl eval -f yaml prod.pkl -p 'tag="1.22.0"'

# Override multiple values
pkl eval -f yaml prod.pkl -p 'replicas=10' -p 'tag="1.22.0"'
```

## Integration with CI/CD

### GitHub Actions
```yaml
- name: Generate Kubernetes manifests
  run: |
    pkl eval -f yaml prod.pkl > manifests.yaml
    kubectl apply -f manifests.yaml
```

### GitLab CI
```yaml
deploy:
  script:
    - pkl eval -f yaml prod.pkl | kubectl apply -f -
```

### Argo CD
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app
spec:
  source:
    plugin:
      name: pkl
      env:
        - name: PKL_FILE
          value: prod.pkl
```

## Debugging

### Check what properties are available
```bash
pkl eval -f yaml template.pkl
```

### Validate with verbose output
```bash
pkl eval prod.pkl -v
```

### Check Pkl version
```bash
pkl --version
```

## Creating New Environments

1. Copy an existing environment file:
```bash
cp staging.pkl qa.pkl
```

2. Modify the values:
```pkl
amends "template.pkl"

namespace = "qa"
replicas = 2
image = "nginx"
tag = "1.21.6"
host = "myapp-qa.example.com"
```

3. Generate and apply:
```bash
pkl eval -f yaml qa.pkl | kubectl apply -f -
```

## Testing Changes

### Test locally with Kind/Minikube
```bash
# Start local cluster
kind create cluster --name test

# Generate and apply
pkl eval -f yaml staging.pkl | kubectl apply -f -

# Check resources
kubectl get deployments,services,ingress -n staging

# Cleanup
kind delete cluster --name test
```

### Port forward to test service
```bash
kubectl port-forward -n production svc/my-app-service 8080:80
curl localhost:8080
```

## Tips and Tricks

### 1. Use environment variables
```pkl
// In your .pkl file
namespace = read("env:NAMESPACE")
```

Then:
```bash
NAMESPACE=production pkl eval -f yaml prod.pkl
```

### 2. Include additional resources
Edit `template.pkl` to add ConfigMaps, Secrets, etc.:
```pkl
configMap = new {
  apiVersion = "v1"
  kind = "ConfigMap"
  metadata {
    name = "\(appName)-config"
    namespace = module.namespace
  }
  data {
    ["app.env"] = "production"
  }
}

output {
  renderer = new YamlRenderer { isStream = true }
  value = List(deployment, service, ingress, configMap)
}
```

### 3. Conditional resources
```pkl
// Only create ingress in production
output {
  renderer = new YamlRenderer { isStream = true }
  value = if (namespace == "production")
    then List(deployment, service, ingress)
    else List(deployment, service)
}
```

## Common Issues

### Issue: "Module not found"
Make sure you're in the correct directory:
```bash
cd /path/to/kubecon-us-2025-code/tools/pkl
```

### Issue: "Circular reference"
Make sure to use `module.property` when referencing module-level properties inside objects.

### Issue: "Type constraint violated"
Check that your values meet the constraints defined in the template:
```pkl
replicas: Int(this >= 1)  // Must be at least 1
```

## Resources

- [Pkl Official Documentation](https://pkl-lang.org/)
- [Pkl Language Tutorial](https://pkl-lang.org/main/current/language-tutorial/02_filling_out_a_template.html)
- [Pkl Standard Library](https://pkl-lang.org/package-docs/pkl/0.29.1/index.html)
- [Pkl GitHub Repository](https://github.com/apple/pkl)

