package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/require"
)

func TestHelmDeployment(t *testing.T) {
	t.Parallel()

	helmOptions := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", "default"),
		ValuesFiles:    []string{"../helm/mock-app/values-dev.yaml"},
	}

	releaseName := "mock-app"
	helm.Install(t, helmOptions, "../helm/mock-app", releaseName)

	kubeOptions := k8s.NewKubectlOptions("", "", "default")

	// Check if Deployment is available
	deployment := k8s.GetDeployment(t, kubeOptions, "mock-app-deployment")
	require.True(t, k8s.IsDeploymentAvailable(deployment), "Deployment should be available")

	// Check if Service is available
	service := k8s.GetService(t, kubeOptions, "mock-app-service")
	require.NotNil(t, service, "Service should be available")

	// Check if ConfigMap is available
	configMap := k8s.GetConfigMap(t, kubeOptions, "myhelmapp-configmap-v1")
	require.NotNil(t, configMap, "ConfigMap should be available")

	// Clean up Helm release after verification
	helm.Delete(t, helmOptions, releaseName, true)
}
