package main

import (
	"flag"
	"fmt"

	"example.com/cdk8s/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewChart(scope constructs.Construct, id string, env Environment) cdk8s.Chart {

	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
		Namespace: jsii.String(env.Namespace),
	})

	// Generate consistent app label
	appLabel := "my-app"

	podSelector := map[string]*string{
		"app": jsii.String(appLabel),
	}

	// Create Deployment
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

	// Create Service pointing to the Deployment pods
	serviceName := fmt.Sprintf("%s-service", appLabel)
	k8s.NewKubeService(chart, jsii.String("service"), &k8s.KubeServiceProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(serviceName),
		},
		Spec: &k8s.ServiceSpec{
			Type:     jsii.String("ClusterIP"),
			Selector: &podSelector,
			Ports: &[]*k8s.ServicePort{{
				Name:       jsii.String("http"),
				Port:       jsii.Number(80),
				TargetPort: k8s.IntOrString_FromString(jsii.String("http")),
				Protocol:   jsii.String("TCP"),
			}},
		},
	})

	// Create Ingress pointing to the Service
	pathTypePrefix := "Prefix"
	k8s.NewKubeIngress(chart, jsii.String("ingress"), &k8s.KubeIngressProps{
		Metadata: &k8s.ObjectMeta{
			Name: jsii.String(fmt.Sprintf("%s-ingress", appLabel)),
			Annotations: &map[string]*string{
				"kubernetes.io/ingress.class": jsii.String("nginx"),
			},
		},
		Spec: &k8s.IngressSpec{
			Rules: &[]*k8s.IngressRule{{
				Host: jsii.String(env.Host),
				Http: &k8s.HttpIngressRuleValue{
					Paths: &[]*k8s.HttpIngressPath{{
						Path:     jsii.String("/"),
						PathType: jsii.String(pathTypePrefix),
						Backend: &k8s.IngressBackend{
							Service: &k8s.IngressServiceBackend{
								Name: jsii.String(serviceName),
								Port: &k8s.ServiceBackendPort{
									Number: jsii.Number(80),
								},
							},
						},
					}},
				},
			}},
		},
	})

	return chart
}

func main() {
	// Parse command line flags
	envFlag := flag.String("env", "development", "Environment to deploy (development, staging, production)")
	flag.Parse()

	// Get environment configuration
	env := GetEnvironmentConfig(*envFlag)

	fmt.Printf("Generating manifests for %s environment...\n", env.Name)
	fmt.Printf("Namespace: %s\n", env.Namespace)
	fmt.Printf("Replicas: %d\n", env.Replicas)
	fmt.Printf("Image: %s:%s\n", env.Image, env.Tag)
	fmt.Printf("Host: %s\n", env.Host)

	// Create CDK8s app with custom output directory for environment
	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir: jsii.String(fmt.Sprintf("dist/%s", env.Name)),
	})

	// Create chart with environment-specific configuration
	NewChart(app, fmt.Sprintf("app-%s", env.Name), env)

	app.Synth()

	fmt.Printf("Manifests generated successfully in dist/%s/\n", env.Name)
}
