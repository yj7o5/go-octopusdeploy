package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureSubscriptionAccount represents an Azure subscription account.
type AzureSubscriptionAccount struct {
	AccountType           string          `json:"AccountType" validate:"required,eq=AzureSubscription"`
	AzureEnvironment      string          `json:"AzureEnvironment,omitempty" validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	CertificateBytes      *SensitiveValue `json:"CertificateBytes,omitempty"`
	CertificateThumbprint string          `json:"CertificateThumbprint,omitempty"`
	ManagementEndpoint    string          `json:"ServiceManagementEndpointBaseUri,omitempty" validate:"omitempty,uri"`
	StorageEndpointSuffix string          `json:"ServiceManagementEndpointSuffix,omitempty" validate:"omitempty,hostname"`
	SubscriptionID        *uuid.UUID      `json:"SubscriptionNumber" validate:"required"`

	AccountResource
}

// NewAzureSubscriptionAccount creates and initializes an Azure subscription
// account with a name.
func NewAzureSubscriptionAccount(name string, subscriptionID uuid.UUID) *AzureSubscriptionAccount {
	return &AzureSubscriptionAccount{
		AccountType:     "AzureSubscription",
		SubscriptionID:  &subscriptionID,
		AccountResource: *newAccountResource(name),
	}
}

// GetAccountType returns the account type for this account.
func (a *AzureSubscriptionAccount) GetAccountType() string {
	return a.AccountType
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AzureSubscriptionAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IAccount = &AzureSubscriptionAccount{}
