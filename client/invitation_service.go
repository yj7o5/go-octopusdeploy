package client

import "github.com/dghubble/sling"

type invitationService struct {
	service
}

func newInvitationService(sling *sling.Sling, uriTemplate string) *invitationService {
	return &invitationService{
		service: newService(serviceInvitationService, sling, uriTemplate, nil),
	}
}
