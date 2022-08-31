package controllers

import (
	"backend/errors"
	"fmt"
	// "backend/location"
	"encoding/json"

	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"backend/location"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryID string) (*location.Country, *errors.ApiError)
)

type locationServiceMock struct{}

func (*locationServiceMock) GetCountry(countryID string) (*location.Country, *errors.ApiError) {
	return getCountryFunc
}

func TestGetCountryNotFound(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request,_ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "id", Value: "ARS"},
	}

	GetCountry(c)
	var apiErr errors.ApiError
	body, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &apiErr ); err != nil {
		assert.NotNil(t, err)
	}

	fmt.Println("body: ", body)

	// assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}