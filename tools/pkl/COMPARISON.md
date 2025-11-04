# Pkl vs cdk8s - Side by Side Comparison

## Overview

This document compares the Pkl and cdk8s approaches for generating Kubernetes manifests.

## File Structure

### cdk8s
```
cdk8s/
├── main.go           # Main application logic (129 lines)
├── config.go         # Environment configurations (79 lines)
├── go.mod            # Go dependencies
├── go.sum            # Go dependency checksums
├── Makefile          # Build commands
└── dist/             # Generated YAML
    ├── production/
    └── staging/
```

### Pkl
```
pkl/
├── template.pkl      # Base template (35 lines)
├── prod.pkl          # Production config (7 lines)
├── staging.pkl       # Staging config (7 lines)
├── Makefile          # Build commands
└── dist/             # Generated YAML
    ├── production/
    └── staging/
```

**Result**: Pkl requires ~80% less code for the same functionality.

## Code Comparison

### Defining a Deployment

#### cdk8s (Go)
```go
// In main.go - imperative, programmatic
k8s.NewKubeDeployment(chart, jsii.String("deployment"), &k8s.KubeDeploymentProps{
    Metadata: &k8s.ObjectMeta{},
    Spec: &k8s.DeploymentSpec{
        Replicas: jsii.Number(float64(env.Replicas)),
        Selector: &k8s.LabelSelector{
            MatchLabels: &podSelector,
        },
        Template: &k8s.PodTemplateSpec{
            Metadata: &k8s.ObjectMeta{
                Labels: &podSelector,
            },
            Spec: &k8s.PodSpec{
                Containers: &[]*k8s.Container{{
                    Name:  jsii.String("app-container"),
                    Image: jsii.String(fmt.Sprintf("%s:%s", env.Image, env.Tag)),
                    Ports: &[]*k8s.ContainerPort{{
                        ContainerPort: jsii.Number(80),
                        Name:          jsii.String("http"),
                    }},
                }},
            },
        },
    },
})
```

#### Pkl (Declarative)
```pkl
// In template.pkl - declarative, configuration-focused
deployment = new {
  apiVersion = "apps/v1"
  kind = "Deployment"
  metadata {
    name = appName
    namespace = module.namespace
  }
  spec {
    replicas = module.replicas
    selector {
      matchLabels = module.labels
    }
    template {
      metadata {
        labels = module.labels
      }
      spec {
        containers {
          new {
            name = "app-container"
            image = "\(module.image):\(module.tag)"
            ports {
              new {
                name = "http"
                containerPort = 80
              }
            }
          }
        }
      }
    }
  }
}
```

**Key Differences**:
- Pkl: Direct YAML-like structure, easier to read and map to Kubernetes docs
- cdk8s: Go structs with pointers, requires understanding of the library API

### Environment Configuration

#### cdk8s (Go)
```go
// In config.go - requires a full struct and switch statement
func GetEnvironmentConfig(env string) Environment {
    switch env {
    case "production", "prod":
        return Environment{
            Name:      "production",
            Namespace: "production",
            Replicas:  5,
            Image:     "nginx",
            Tag:       "1.21.6",
            Host:      "myapp-prod.example.com",
        }
    case "staging":
        return Environment{
            Name:      "staging",
            Namespace: "staging",
            Replicas:  2,
            Image:     "nginx",
            Tag:       "1.21.6",
            Host:      "myapp-staging.example.com",
        }
    }
}
```

#### Pkl
```pkl
// In prod.pkl - amends (inherits from) template
amends "template.pkl"

namespace = "production"
replicas = 5
image = "nginx"
tag = "1.21.6"
host = "myapp-prod.example.com"
```

```pkl
// In staging.pkl - only specifies what's different
amends "template.pkl"

namespace = "staging"
replicas = 2
image = "nginx"
tag = "1.21.6"
host = "myapp-staging.example.com"
```

**Key Differences**:
- Pkl: Uses "amending" - only specify what's different from the base
- cdk8s: Requires full environment struct definition with all fields

## Type Safety and Validation

### cdk8s
```go
// Type safety through Go's type system
// Validation through code
if env.Replicas < 1 {
    return errors.New("replicas must be at least 1")
}
```

### Pkl
```pkl
// Built-in type constraints
replicas: Int(this >= 1)  // Validated automatically
namespace: String         // Required field
```

**Winner**: Tie - both provide strong type safety, but Pkl's is more declarative

## Workflow

### cdk8s
```bash
# Need to compile Go code
make generate ENV=production
# Output: dist/production/*.yaml

# Behind the scenes:
# 1. go run with environment flag
# 2. CDK8s synthesizes to YAML
```

### Pkl
```bash
# Direct evaluation, no compilation
make generate-prod
# Output: dist/production/manifests.yaml

# Behind the scenes:
# 1. pkl eval prod.pkl
# 2. Direct YAML output
```

## Adding a New Environment

### cdk8s
1. Edit `config.go`, add new case to switch statement
2. Define all environment properties
3. Run `make generate ENV=dev`

### Pkl
1. Create `dev.pkl`:
```pkl
amends "template.pkl"

namespace = "development"
replicas = 1
image = "nginx"
tag = "latest"
host = "myapp-dev.localhost"
```
2. Run `pkl eval -f yaml dev.pkl > dist/dev/manifests.yaml`

## Learning Curve

| Aspect | cdk8s | Pkl |
|--------|-------|-----|
| Prerequisites | Go knowledge, cdk8s API | Pkl syntax (YAML-like) |
| Setup Time | Medium (Go modules, imports) | Low (install Pkl CLI) |
| IDE Support | Full (Go tooling) | Good (VS Code, IntelliJ) |
| Debugging | Go debugger | Pkl errors with stack traces |
| Documentation | Extensive | Growing |

## Performance

| Operation | cdk8s | Pkl |
|-----------|-------|-----|
| Initial Build | ~5s (includes Go compilation) | <1s (direct evaluation) |
| Incremental | ~3s | <1s |
| Cold Start | ~5s | <1s |

## When to Use Each

### Use cdk8s when:
- ✅ You need complex programmatic logic
- ✅ You want to leverage existing Go/TS/Python ecosystem
- ✅ Your team is already proficient in these languages
- ✅ You need to integrate with existing Go applications
- ✅ You require complex conditional logic or loops

### Use Pkl when:
- ✅ You prefer declarative configuration over code
- ✅ You want minimal boilerplate
- ✅ You need fast iteration cycles
- ✅ You want YAML-like syntax with type safety
- ✅ Your use case is primarily configuration templating
- ✅ You want to avoid a compilation step

## Actual Output Comparison

Both generate identical Kubernetes manifests:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
  namespace: production
spec:
  replicas: 5
  # ... rest of deployment
---
apiVersion: v1
kind: Service
# ... service definition
---
apiVersion: networking.k8s.io/v1
kind: Ingress
# ... ingress definition
```

## Conclusion

**cdk8s** excels at:
- Complex logic and conditionals
- Integration with existing codebases
- Programmatic resource generation

**Pkl** excels at:
- Simple, readable configuration
- Fast iteration
- Minimal boilerplate
- Direct YAML-like syntax

Both are excellent tools - the choice depends on your team's preferences and use case complexity.

