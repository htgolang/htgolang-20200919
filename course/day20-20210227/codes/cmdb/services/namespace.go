package services

import (
	"cmdb/services/k8s"

	coreV1 "k8s.io/api/core/v1"
)

type namespaceService struct{}

var NamespaceService = new(namespaceService)

func (s *namespaceService) Query() []coreV1.Namespace {
	client := k8s.NewClient("conf/k8s/config")
	namespaces, _ := client.Namespaces()
	return namespaces
}
