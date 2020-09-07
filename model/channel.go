package model

import (
	"github.com/go-playground/validator/v10"
)

type Channels struct {
	Items []Channel `json:"Items"`
	PagedResults
}

type Channel struct {
	Description string        `json:"Description"`
	IsDefault   bool          `json:"IsDefault"`
	LifecycleID string        `json:"LifecycleId"`
	Name        string        `json:"Name"`
	ProjectID   string        `json:"ProjectId"`
	Rules       []ChannelRule `json:"Rules,omitempty"`
	TenantTags  []string      `json:"TenantedDeploymentMode,omitempty"`
	Resource
}

func (c *Channel) GetID() string {
	return c.ID
}

func (c *Channel) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return ValidateMultipleProperties([]error{
		ValidateRequiredPropertyValue("Name", c.Name),
		ValidateRequiredPropertyValue("ProjectID", c.ProjectID),
	})
}

func NewChannel(name, description, projectID string) *Channel {
	return &Channel{
		Name:        name,
		ProjectID:   projectID,
		Description: description,
	}
}

var _ ResourceInterface = &Channel{}
