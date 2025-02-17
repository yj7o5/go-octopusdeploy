package octopusdeploy

import (
	"encoding/json"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPropertyValueBehaviour(t *testing.T) {
	pv := SensitiveValue{}
	require.NotNil(t, pv)
	require.False(t, pv.HasValue)
	require.Nil(t, pv.NewValue)

	pvp := &SensitiveValue{}
	require.NotNil(t, pvp)
	require.False(t, pvp.HasValue)
	require.Nil(t, pvp.NewValue)
}

func TestNewPropertyValueBehaviour(t *testing.T) {
	pvp := NewPropertyValue("", false)
	require.NotNil(t, pvp)
	require.False(t, pvp.IsSensitive)
	require.Nil(t, pvp.SensitiveValue)
	require.Len(t, pvp.Value, 0)

	pvp = NewPropertyValue("", true)
	require.NotNil(t, pvp)
	require.True(t, pvp.IsSensitive)
	require.NotNil(t, pvp.SensitiveValue)
	require.False(t, pvp.SensitiveValue.HasValue)
	require.Nil(t, pvp.SensitiveValue.NewValue)
	require.Len(t, pvp.Value, 0)

	pvp = NewPropertyValue("test", false)
	require.NotNil(t, pvp)
	require.False(t, pvp.IsSensitive)
	require.Nil(t, pvp.SensitiveValue)
	require.Equal(t, "test", pvp.Value)

	pvp = NewPropertyValue("sensitive value", true)
	require.NotNil(t, pvp)
	require.True(t, pvp.IsSensitive)
	require.NotNil(t, pvp.SensitiveValue)
	require.True(t, pvp.SensitiveValue.HasValue)
	require.Equal(t, "sensitive value", *pvp.SensitiveValue.NewValue)
	require.Len(t, pvp.Value, 0)
}

func TestNewPropertyValueMarshalJSON(t *testing.T) {
	propertyValue := NewPropertyValue("", true)
	propertyValueAsJSON, err := json.Marshal(propertyValue)
	require.NoError(t, err)
	require.NotNil(t, propertyValueAsJSON)
	jsonassert.New(t).Assertf(string(propertyValueAsJSON), emptySensitivePropertyValueAsJSON)

	propertyValue = NewPropertyValue("non-sensitive value", false)
	propertyValueAsJSON, err = json.Marshal(propertyValue)
	require.NoError(t, err)
	require.NotNil(t, propertyValueAsJSON)
	jsonassert.New(t).Assertf(string(propertyValueAsJSON), testNonSensitivePropertyValueAsJSON)

	propertyValue = NewPropertyValue("test", true)
	propertyValueAsJSON, err = json.Marshal(propertyValue)
	require.NoError(t, err)
	require.NotNil(t, propertyValueAsJSON)
	jsonassert.New(t).Assertf(string(propertyValueAsJSON), testSensitivePropertyValueAsJSON)
}

func TestNewPropertyValueUnmarshalJSON(t *testing.T) {
	var propertyValue PropertyValue
	err := json.Unmarshal([]byte(testNonSensitivePropertyValueAsJSON), &propertyValue)
	require.NoError(t, err)
	require.NotNil(t, propertyValue)
	require.Equal(t, "non-sensitive value", propertyValue.Value)
	require.False(t, propertyValue.IsSensitive)
	require.Nil(t, propertyValue.SensitiveValue)

	var emptySensitiveValue PropertyValue
	err = json.Unmarshal([]byte(emptySensitivePropertyValueAsJSON), &emptySensitiveValue)
	require.NoError(t, err)
	require.NotNil(t, emptySensitiveValue)
	require.False(t, emptySensitiveValue.SensitiveValue.HasValue)
	require.Nil(t, emptySensitiveValue.SensitiveValue.NewValue)
	require.Empty(t, emptySensitiveValue.Value)

	var testPropertyValue PropertyValue
	err = json.Unmarshal([]byte(testSensitivePropertyValueAsJSON), &testPropertyValue)
	require.NoError(t, err)
	require.NotNil(t, testPropertyValue)
	require.NotNil(t, testPropertyValue.SensitiveValue.NewValue)
	require.True(t, testPropertyValue.SensitiveValue.HasValue)
	require.Equal(t, "test", *testPropertyValue.SensitiveValue.NewValue)
	require.Empty(t, testPropertyValue.Value)
}

const emptySensitivePropertyValueAsJSON string = `{
	"HasValue": false,
	"NewValue": null,
	"Hint": null
}`

const testSensitivePropertyValueAsJSON string = `{
	"HasValue": true,
	"NewValue": "test",
	"Hint": null
}`

const testNonSensitivePropertyValueAsJSON string = `"non-sensitive value"`
