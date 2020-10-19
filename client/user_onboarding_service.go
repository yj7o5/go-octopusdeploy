package client

import "github.com/dghubble/sling"

type userOnboardingService struct {
	service
}

func newUserOnboardingService(sling *sling.Sling, uriTemplate string) *userOnboardingService {
	return &userOnboardingService{
		service: newService(serviceUserOnboardingService, sling, uriTemplate, nil),
	}
}
