package services

import (
	"cmdb/services/k8s"

	appsV1 "k8s.io/api/apps/v1"

	coreV1 "k8s.io/api/core/v1"
)

type deploymentService struct{}

var DeploymentService = new(deploymentService)

func (s *deploymentService) Query() []appsV1.Deployment {
	client := k8s.NewClient("conf/k8s/config")
	deployments, _ := client.Deployments()
	return deployments
}

func (s *deploymentService) Create(namespace, name, image string, labels map[string]string, ports []coreV1.ContainerPort, replicas int) {
	client := k8s.NewClient("conf/k8s/config")
	client.CreateDeployment(namespace, name, image, labels, ports, replicas)
}

func (s *deploymentService) Get(namespace, name string) *appsV1.Deployment {
	client := k8s.NewClient("conf/k8s/config")
	return client.GetDeployment(namespace, name)
}

func (s *deploymentService) Modify(namespace, name, image string, ports []coreV1.ContainerPort, replicas int) {
	client := k8s.NewClient("conf/k8s/config")
	client.ModifyDeployment(namespace, name, image, ports, replicas)
}

func (s *deploymentService) Delete(name, namespace string) {
	client := k8s.NewClient("conf/k8s/config")
	client.DeleteDeployment(name, namespace)
}
