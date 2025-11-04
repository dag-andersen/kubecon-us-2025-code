# Pkl Kubernetes Manifests

This directory contains Pkl configurations for generating Kubernetes manifests (Deployment, Service, and Ingress) for different environments.

## What is Pkl?

[Pkl](https://pkl-lang.org/) is a configuration-as-code language from Apple that provides:
- Type-safe configuration
- Template inheritance through "amending"
- Built-in validation
- Multiple output formats (YAML, JSON, etc.)

## Structure

```
pkl/
├── template.pkl    # Base template with all resource definitions
├── prod.pkl        # Production environment (amends template.pkl)
├── staging.pkl     # Staging environment (amends template.pkl)
├── Makefile        # Build targets for generating manifests
└── dist/           # Generated YAML manifests (gitignored)
    ├── production/
    └── staging/
```

## How It Works

The pattern follows Pkl's "amending" concept (similar to filling out a form):

1. **template.pkl** - Defines the base structure for all Kubernetes resources:
   - Deployment with configurable replicas, image, and namespace
   - Service pointing to the Deployment
   - Ingress routing to the Service

2. **Environment files** (prod.pkl, staging.pkl) - "Amend" the template by providing environment-specific values:
   - Production: 5 replicas, production namespace
   - Staging: 2 replicas, staging namespace

## Installation

Install Pkl on macOS:

```bash
brew install pkl
```

Or use the Makefile:

```bash
make install
```

For other platforms, see: https://pkl-lang.org/main/current/pkl-cli/index.html#installation

## Usage

### Generate all environments:
```bash
make generate-all
```

### Generate specific environment:
```bash
make generate-prod     # Generates dist/production/manifests.yaml
make generate-staging  # Generates dist/staging/manifests.yaml
```

### Validate configurations:
```bash
make validate
```

### Compare environments:
```bash
make diff
```

### Clean generated files:
```bash
make clean
```

## Manual Usage (without Makefile)

Generate YAML for production:
```bash
pkl eval -f yaml prod.pkl > manifests.yaml
```

Generate YAML for staging:
```bash
pkl eval -f yaml staging.pkl > manifests.yaml
```

## Customization

To add a new environment:

1. Create a new file (e.g., `dev.pkl`)
2. Amend the base template:
   ```pkl
   amends "template.pkl"
   
   namespace = "development"
   replicas = 1
   image = "nginx"
   tag = "latest"
   host = "myapp-dev.localhost"
   ```

3. Generate manifests:
   ```bash
   pkl eval -f yaml dev.pkl > dist/development/manifests.yaml
   ```

## Comparison with cdk8s

| Feature | Pkl | cdk8s |
|---------|-----|-------|
| Language | Pkl (declarative) | Go/TypeScript/Python (imperative) |
| Type Safety | ✅ Built-in | ✅ Via TypeScript/Go types |
| Validation | ✅ Constraints in config | Code-based validation |
| Template Reuse | "Amending" | Object-oriented inheritance |
| Learning Curve | Low | Medium (requires programming) |
| IDE Support | ✅ VS Code, IntelliJ | ✅ Full IDE support |

## Benefits of Pkl

1. **Type Safety**: Catch errors before deployment
   - Example: `replicas: Int(this >= 1)` ensures at least 1 replica

2. **Template Inheritance**: Environment configs only specify what's different
   - No code duplication
   - Clear separation of base and environment-specific config

3. **Validation**: Built-in constraints
   ```pkl
   replicas: Int(this >= 1)  // Must be at least 1
   namespace: String          // Required field
   ```

4. **Multi-format Output**: Generate YAML, JSON, or other formats from the same source

5. **No Build Step**: Pure configuration - no compilation needed

## Resources

- [Pkl Documentation](https://pkl-lang.org/main/current/language-tutorial/02_filling_out_a_template.html)
- [Pkl GitHub](https://github.com/apple/pkl)
- [Pkl Standard Library](https://pkl-lang.org/package-docs/pkl/0.29.1/index.html)

