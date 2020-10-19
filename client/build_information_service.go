package client

import (
	"github.com/dghubble/sling"
)

type buildInformationService struct {
	bulkPath string

	service
}

func newBuildInformationService(sling *sling.Sling, uriTemplate string, bulkPath string) *buildInformationService {
	buildInformationService := &buildInformationService{
		bulkPath: bulkPath,
	}
	buildInformationService.service = newService(serviceBuildInformationService, sling, uriTemplate, nil)

	return buildInformationService
}
