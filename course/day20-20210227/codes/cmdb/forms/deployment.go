package forms

import (
	"fmt"
	"strconv"
	"strings"

	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
)

type AddDeploymentForm struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
	Image     string `form:"image"`
	Label     string `form:"label"`
	Expose    string `form:"expose"`
	Replicas  int    `form:"replicas"`
}

func (f *AddDeploymentForm) Labels() map[string]string {
	/*
		label:
		key1:value1
		key2:value2
	*/
	labels := make(map[string]string)
	for _, line := range strings.Split(f.Label, "\n") {
		nodes := strings.SplitN(strings.TrimSpace(line), ":", 2)
		if len(nodes) == 2 {
			labels[nodes[0]] = nodes[1]
		}
	}
	return labels
}

func (f *AddDeploymentForm) Exposes() []coreV1.ContainerPort {
	/*
		expose
		name:port:protocol
	*/
	ports := []coreV1.ContainerPort{}
	for _, line := range strings.Split(f.Expose, "\n") {
		nodes := strings.SplitN(strings.TrimSpace(line), ":", 3)
		if len(nodes) == 3 {
			if port, err := strconv.Atoi(nodes[1]); err == nil {
				ports = append(ports, coreV1.ContainerPort{
					Name:          nodes[0],
					ContainerPort: int32(port),
					Protocol:      coreV1.Protocol(strings.ToUpper(nodes[2])),
				})
			}

		}
	}
	return ports
}

type ModifyDeploymentForm struct {
	Namespace string `form:"namespace"`
	Name      string `form:"name"`
	Image     string `form:"image"`
	Expose    string `form:"expose"`
	Replicas  int    `form:"replicas"`
}

func (f *ModifyDeploymentForm) FromModel(deployment *appsV1.Deployment) {
	if deployment == nil {
		return
	}
	exposes := []string{}
	for _, port := range deployment.Spec.Template.Spec.Containers[0].Ports {
		exposes = append(exposes, fmt.Sprintf("%s:%d:%s", port.Name, port.ContainerPort, port.Protocol))
	}

	f.Namespace = deployment.Namespace
	f.Name = deployment.Name
	f.Image = deployment.Spec.Template.Spec.Containers[0].Image
	f.Expose = strings.Join(exposes, "\n")
	f.Replicas = int(*deployment.Spec.Replicas)
}

func (f *ModifyDeploymentForm) Exposes() []coreV1.ContainerPort {
	/*
		expose
		name:port:protocol
	*/
	ports := []coreV1.ContainerPort{}
	for _, line := range strings.Split(f.Expose, "\n") {
		nodes := strings.SplitN(strings.TrimSpace(line), ":", 3)
		if len(nodes) == 3 {
			if port, err := strconv.Atoi(nodes[1]); err == nil {
				ports = append(ports, coreV1.ContainerPort{
					Name:          nodes[0],
					ContainerPort: int32(port),
					Protocol:      coreV1.Protocol(strings.ToUpper(nodes[2])),
				})
			}

		}
	}
	return ports
}
