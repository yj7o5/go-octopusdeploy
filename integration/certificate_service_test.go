package integration

import (
	"reflect"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testCert1Data = `MIIDiDCCAnACCQDXHofnqz05ITANBgkqhkiG9w0BAQsFADCBhTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE9rbGFob21hMQ8wDQYDVQQHDAZOb3JtYW4xEzARBgNVBAoMCk1vb25zd2l0Y2gxGTAXBgNVBAMMEGRlbW8ub2N0b3B1cy5jb20xIjAgBgkqhkiG9w0BCQEWE2plZmZAbW9vbnN3aXRjaC5jb20wHhcNMTkwNjE0MjExMzI1WhcNMjAwNjEzMjExMzI1WjCBhTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE9rbGFob21hMQ8wDQYDVQQHDAZOb3JtYW4xEzARBgNVBAoMCk1vb25zd2l0Y2gxGTAXBgNVBAMMEGRlbW8ub2N0b3B1cy5jb20xIjAgBgkqhkiG9w0BCQEWE2plZmZAbW9vbnN3aXRjaC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDSTiD0OHyFDMH9O+d/h3AiqcuvpvUgRkKjf+whZ6mVlQnGkvPddRTUY48xCEaQ4QD1MAVJcGaJ2PU4NxwhrQgHqWW8TQkAZESL4wfzSwIKO2NX/I2tWqyv7a0uA/WdtlWQye+2oPV5rCnS0kM75X+gjEwOTpFh/ryS6KhMPFDb0zeNGREdg6564FdxWSvN4ppUZMqhvMpfzM7rsDWqEzYsMaQ4CNJDFdWkG89D4j5qk4b4Qb4m+l7QINdmYIXf4qO/0LE1WcfIkCpAS65tjc/hefIHmYtj/E/ijoNJbWKZDK3WLZg3zq99Ipqv/9DFvSiMQFBhZT0jO2B5d5zBUuIHAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAKsa4gaW7GhByu8aq56h99DaIl1LauI5WMVH8Q9Qpapho2VLRIpfwGeI5eENFoXwuKrnJp1ncsCqGnMQnugQHS+SrruS3Yyl0Uog4Zak9GbbK6qn+olx7GNJbsckmD371lqQOaKITLqYzK6kTc7/v8Cv0BwHFCBda1OCrmeVBSaarucPxZhGxzLAielzHHdlkZFQT/oO2VR3thhURIqtni7jVQ2MoeZF1ccvmAfVbzr/QnlNe/jrcmyPYymuF2JyrezzIjlKuiDhalKqwqkCHpOOgzV4y6BFuS+0w3DS8pa07nUudZ6E0kZzvhjjiyAx/sBdX6ZDdUjP9TDJMM4f5YA=`
var testCert2Data = `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`

const testCert2Thumbprint = `0454EAD1FC6F3F60F22AE82230165C14C4FD7FA7`
const testCert2Password = `HCWVMo7u`

func AssertEqualCertificateResources(t *testing.T, expected *octopusdeploy.CertificateResource, actual *octopusdeploy.CertificateResource) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, reflect.DeepEqual(expected.GetLinks(), actual.GetLinks()))

	// CertificateResource
	assert.Equal(t, expected.Archived, actual.Archived)
	assert.Equal(t, expected.CertificateData, actual.CertificateData)
	assert.Equal(t, expected.CertificateDataFormat, actual.CertificateDataFormat)
	assert.True(t, reflect.DeepEqual(expected.EnvironmentIDs, actual.EnvironmentIDs))
	assert.Equal(t, expected.HasPrivateKey, actual.HasPrivateKey)
	assert.Equal(t, expected.IsExpired, actual.IsExpired)
	assert.Equal(t, expected.IssuerCommonName, actual.IssuerCommonName)
	assert.Equal(t, expected.IssuerDistinguishedName, actual.IssuerDistinguishedName)
	assert.Equal(t, expected.IssuerOrganization, actual.IssuerOrganization)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.NotAfter, actual.NotAfter)
	assert.Equal(t, expected.NotBefore, actual.NotBefore)
	assert.Equal(t, expected.Notes, actual.Notes)
	assert.Equal(t, expected.Password, actual.Password)
	assert.Equal(t, expected.ReplacedBy, actual.ReplacedBy)
	assert.Equal(t, expected.SerialNumber, actual.SerialNumber)
	assert.Equal(t, expected.SignatureAlgorithmName, actual.SignatureAlgorithmName)
	assert.True(t, reflect.DeepEqual(expected.SubjectAlternativeNames, actual.SubjectAlternativeNames))
	assert.Equal(t, expected.SubjectDistinguishedName, actual.SubjectDistinguishedName)
	assert.Equal(t, expected.SubjectCommonName, actual.SubjectCommonName)
	assert.Equal(t, expected.SubjectOrganization, actual.SubjectOrganization)
	assert.Equal(t, expected.SelfSigned, actual.SelfSigned)
	assert.Equal(t, expected.TenantedDeploymentMode, actual.TenantedDeploymentMode)
	assert.True(t, reflect.DeepEqual(expected.TenantIDs, actual.TenantIDs))
	assert.True(t, reflect.DeepEqual(expected.TenantTags, actual.TenantTags))
	assert.Equal(t, expected.Thumbprint, actual.Thumbprint)
	assert.Equal(t, expected.Version, actual.Version)
}

func CreateTestCertificateResource(t *testing.T, client *octopusdeploy.Client) *octopusdeploy.CertificateResource {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	certificateData := octopusdeploy.NewSensitiveValue(testCert2Data)
	require.NotNil(t, certificateData)

	password := octopusdeploy.NewSensitiveValue(testCert2Password)
	require.NotNil(t, password)

	certificate := octopusdeploy.NewCertificateResource(name, certificateData, password)
	require.NotNil(t, certificate)
	require.NoError(t, certificate.Validate())

	createdCertificate, err := client.Certificates.Add(certificate)
	require.NoError(t, err)
	require.NotNil(t, createdCertificate)
	require.NotEmpty(t, createdCertificate.GetID())

	// verify the add operation was successful
	certificateToCompare, err := client.Certificates.GetByID(createdCertificate.GetID())
	require.NoError(t, err)
	require.NotNil(t, certificateToCompare)
	AssertEqualCertificateResources(t, createdCertificate, certificateToCompare)

	return createdCertificate
}

func DeleteTestCertificateResource(t *testing.T, client *octopusdeploy.Client, certificate *octopusdeploy.CertificateResource) {
	require.NotNil(t, certificate)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Certificates.DeleteByID(certificate.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedCertificate, err := client.Certificates.GetByID(certificate.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedCertificate)
}

func TestCertificateServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	certificate := CreateTestCertificateResource(t, client)
	require.NotNil(t, certificate)
	defer DeleteTestCertificateResource(t, client, certificate)

	certificateToCompare, err := client.Certificates.GetByID(certificate.GetID())
	require.NotNil(t, certificateToCompare)
	require.NoError(t, err)
	AssertEqualCertificateResources(t, certificate, certificateToCompare)
}

func TestCertificateServiceAddReplaceDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	certificate := CreateTestCertificateResource(t, client)
	require.NotNil(t, certificate)
	defer DeleteTestCertificateResource(t, client, certificate)

	replacementCertificate := octopusdeploy.NewReplacementCertificate(testCert2Data, testCert2Password)
	require.NotNil(t, replacementCertificate)

	updatedCertificate, err := client.Certificates.Replace(certificate.GetID(), replacementCertificate)
	require.NoError(t, err)
	require.NotNil(t, updatedCertificate)
	require.NotEmpty(t, updatedCertificate.GetID())
	require.Equal(t, testCert2Thumbprint, updatedCertificate.Thumbprint)
}

func TestCertificateServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	certificates, err := client.Certificates.GetAll()
	require.NoError(t, err)
	require.NotNil(t, certificates)

	for _, certificate := range certificates {
		defer DeleteTestCertificateResource(t, client, certificate)
	}
}
