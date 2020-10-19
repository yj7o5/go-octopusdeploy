package client

import "github.com/dghubble/sling"

type tenantVariableService struct {
	service
}

func newTenantVariableService(sling *sling.Sling, uriTemplate string) *tenantVariableService {
	return &tenantVariableService{
		service: newService(serviceTenantVariableService, sling, uriTemplate, nil),
	}
}
