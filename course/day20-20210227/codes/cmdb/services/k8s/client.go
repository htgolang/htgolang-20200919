package k8s

import (
	"context"
	"fmt"
	"strings"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"cmdb/utils"
)

type client struct {
	client     *kubernetes.Clientset
	kubeconfig string
}

func NewClient(kubeconfig string) *client {
	return &client{
		kubeconfig: kubeconfig,
	}
}

func (c *client) clientset() (*kubernetes.Clientset, error) {
	if c.client == nil {
		// k8s 配置
		var err error
		config, err := clientcmd.BuildConfigFromFlags("", c.kubeconfig)
		if err != nil {
			return nil, err
		}
		c.client, err = kubernetes.NewForConfig(config)
	}
	return c.client, nil
}

func (c *client) Namespaces() ([]coreV1.Namespace, error) {
	clientset, err := c.clientset()
	if err != nil {
		return nil, err
	}
	namespaceClient := clientset.CoreV1().Namespaces()
	namespaceResult, err := namespaceClient.List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaceResult.Items, nil
}

func (c *client) namespaces() ([]string, error) {
	list, err := c.Namespaces()
	if err != nil {
		return nil, err
	}
	rt := []string{}
	for _, namespace := range list {
		rt = append(rt, namespace.Name)
	}
	return rt, nil
}

func (c *client) Deployments() ([]appsV1.Deployment, error) {
	clientset, err := c.clientset()
	namespaces, err := c.namespaces()
	if err != nil {
		return nil, err
	}

	deployments := make([]appsV1.Deployment, 0, 100)
	for _, namespace := range namespaces {
		deploymentClient := clientset.AppsV1().Deployments(namespace)
		depoymentResult, err := deploymentClient.List(context.TODO(), metaV1.ListOptions{})
		if err == nil {
			deployments = append(deployments, depoymentResult.Items...)
		}
	}
	return deployments, nil
}

func (c *client) DeleteDeployment(name, namespace string) {
	clientset, err := c.clientset()
	if err != nil {
		return
	}
	deploymentClient := clientset.AppsV1().Deployments(namespace)
	deploymentClient.Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func (c *client) CreateDeployment(namespace, name, image string, labels map[string]string, ports []coreV1.ContainerPort, replicas int) {
	clientset, err := c.clientset()
	if err != nil {
		return
	}

	names := strings.SplitN(image, ":", 2)
	imageName := names[0]
	deploymentClient := clientset.AppsV1().Deployments(namespace)
	deployment := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: utils.Int32Ptr(int32(replicas)),
			Selector: &metaV1.LabelSelector{
				MatchLabels: labels,
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: labels,
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  imageName,
							Image: image,
							Ports: ports,
						},
					},
				},
			},
		},
	}
	deploymentClient.Create(context.TODO(), deployment, metaV1.CreateOptions{})
}

func (c *client) GetDeployment(namespace, name string) *appsV1.Deployment {
	clientset, err := c.clientset()
	if err != nil {
		return nil
	}
	deploymentClient := clientset.AppsV1().Deployments(namespace)
	deployment, _ := deploymentClient.Get(context.TODO(), name, metaV1.GetOptions{})
	return deployment
}

func (c *client) ModifyDeployment(namespace, name, image string, ports []coreV1.ContainerPort, replicas int) {
	clientset, err := c.clientset()
	if err != nil {
		return
	}

	fmt.Println(namespace, name, image, ports, replicas)

	names := strings.SplitN(image, ":", 2)
	imageName := names[0]
	deploymentClient := clientset.AppsV1().Deployments(namespace)

	deployment, err := deploymentClient.Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return
	}
	deployment.Spec.Replicas = utils.Int32Ptr(int32(replicas))
	deployment.Spec.Template.Spec.Containers[0].Name = imageName
	deployment.Spec.Template.Spec.Containers[0].Image = image
	deployment.Spec.Template.Spec.Containers[0].Ports = ports

	_, err = deploymentClient.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	fmt.Println(err)
}
