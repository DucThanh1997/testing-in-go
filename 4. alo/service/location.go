package service

import (
	"backend/providers/locations_provider"
	"backend/errors"
	"backend/location"
)

type locationServiceInterface interface {
	GetCountry(countryID string) (*location.Country, *errors.ApiError)
}

type locationService struct{}

var (
	LocationService locationServiceInterface
)

func init() {
	LocationService = &locationService{}
}

func (s *locationService) GetCountry(countryID string) (*location.Country, *errors.ApiError) {
	return locations_provider.GetCountry(countryID)
}