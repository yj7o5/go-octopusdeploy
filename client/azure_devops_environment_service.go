package client

import (
	"github.com/dghubble/sling"
)

type azureEnvironmentService struct {
	service
}

func newAzureEnvironmentService(sling *sling.Sling, uriTemplate string) *azureEnvironmentService {
	return &azureEnvironmentService{
		service: newService(serviceAzureEnvironmentService, sling, uriTemplate, nil),
	}
}
