package locations_provider

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountryTimeoutFromAPI(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response when trying to get country AR", err.Message)
}

func TestCountryNotFound(t *testing.T) {
	country, err := GetCountry("ARS")
	// assert.Nil(t, err)
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestCountryInvalidErrorInterface(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when getting country for AR", err.Message)
}

func TestCountryInvalidJsonResponse(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AR", err.Message)
}

func TestCountryNoError(t *testing.T) {
	country, err := GetCountry("AR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.ID)
}