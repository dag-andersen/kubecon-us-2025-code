package main

// Environment represents the configuration for different environments
type Environment struct {
	Name      string
	Namespace string
	Replicas  int
	Image     string
	Tag       string
	Host      string
	ConfigMap ConfigMapConfig
}

// ConfigMapConfig defines configuration data for the environment
type ConfigMapConfig struct {
	Data map[string]string
}

// GetEnvironmentConfig returns configuration for the specified environment
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
			ConfigMap: ConfigMapConfig{
				Data: map[string]string{
					"environment":     "production",
					"log_level":       "warn",
					"database_url":    "postgres://prod-db:5432/myapp",
					"redis_url":       "redis://prod-redis:6379",
					"max_connections": "100",
				},
			},
		}
	case "staging":
		return Environment{
			Name:      "staging",
			Namespace: "staging",
			Replicas:  2,
			Image:     "nginx",
			Tag:       "1.21.6",
			Host:      "myapp-staging.example.com",
			ConfigMap: ConfigMapConfig{
				Data: map[string]string{
					"environment":     "staging",
					"log_level":       "info",
					"database_url":    "postgres://staging-db:5432/myapp",
					"redis_url":       "redis://staging-redis:6379",
					"max_connections": "50",
				},
			},
		}
	default:
		// Default to development/local environment
		return Environment{
			Name:      "development",
			Namespace: "default",
			Replicas:  1,
			Image:     "nginx",
			Tag:       "1.19.10",
			Host:      "myapp-dev.localhost",
			ConfigMap: ConfigMapConfig{
				Data: map[string]string{
					"environment":     "development",
					"log_level":       "debug",
					"database_url":    "postgres://localhost:5432/myapp_dev",
					"redis_url":       "redis://localhost:6379",
					"max_connections": "10",
				},
			},
		}
	}
}
