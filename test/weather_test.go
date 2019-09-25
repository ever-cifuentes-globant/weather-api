package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ever-cifuentes-globant/weather-api/mediators"
	"github.com/ever-cifuentes-globant/weather-api/models"
)

func TestIsDataUpdated(t *testing.T) {
	var weatherShouldPass models.Weather
	weatherShouldPass.RequestedTime = "2019-09-17T16:47:07-03:00"
	actualRespone := mediators.IsDataUpdated(weatherShouldPass)
	expectedResponse := true
	assert.Equal(t, actualRespone, expectedResponse)
}

func TestIsNotDataUpdated(t *testing.T) {
	var weatherShouldNotPass models.Weather
	weatherShouldNotPass.RequestedTime = "2019-09-17T16:40:07-03:00"
	actualRespone := mediators.IsDataUpdated(weatherShouldNotPass)
	expectedResponse := false
	assert.Equal(t, actualRespone, expectedResponse)
}
