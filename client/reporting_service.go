package client

import "github.com/dghubble/sling"

type reportingService struct {
	deploymentsCountedByWeekPath string

	service
}

func newReportingService(sling *sling.Sling, uriTemplate string, deploymentsCountedByWeekPath string) *reportingService {
	return &reportingService{
		deploymentsCountedByWeekPath: deploymentsCountedByWeekPath,

		service: newService(serviceReportingService, sling, uriTemplate, nil),
	}
}
