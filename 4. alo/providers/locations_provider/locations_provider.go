package locations_provider

import (
	"backend/location"
	"backend/errors"
	"encoding/json"
	"fmt"
	// "net"
	"net/http"
	"io/ioutil"
)

func GetCountry(countryId string) (*location.Country, *errors.ApiError) {
	response, _ := http.Get(fmt.Sprintf("https://api.mercadolibre.com/countries/%s", countryId))
	if response == nil{
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
		}
	}
	if response.StatusCode == 404 {
		return nil, &errors.ApiError{
				Status: http.StatusNotFound,
				Message: fmt.Sprintf("Country not found"),
			}
	}
	if response.StatusCode > 299 {
		var apiErr errors.ApiError
		body, err := ioutil.ReadAll(response.Body)
		fmt.Println("err: ", err)
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status: http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid restclient response when trying to get country %s", countryId),
			}
		}
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid when getting country %s", countryId),
		}
	}

	var result location.Country
	body, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, &errors.ApiError{
			Status: http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}
	return &result, nil
}