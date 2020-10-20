package octopusdeploy

import "github.com/dghubble/sling"

type userRoleService struct {
	canDeleteService
}

func newUserRoleService(sling *sling.Sling, uriTemplate string) *userRoleService {
	userRoleService := &userRoleService{}
	userRoleService.service = newService(serviceUserRoleService, sling, uriTemplate, nil)

	return userRoleService
}
