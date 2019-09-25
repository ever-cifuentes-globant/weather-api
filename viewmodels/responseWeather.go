package viewmodels

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

type Sys struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type ErrorResponse struct {
	StatusCode int64  `json:"statuscode"`
	Message    string `json:"message"`
}

type WeatherJSON struct {
	City       string    `json:"name"`
	APISys     Sys       `json:"sys"`
	APIMain    Main      `json:"main"`
	APIWind    Wind      `json:"wind"`
	APICoords  Coord     `json:"coord"`
	APIWeather []Weather `json:"weather"`
}

type WeatherPersist struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func NewErrorResponse(err error) ErrorResponse {
	errorResponse := ErrorResponse{404, err.Error()}
	return errorResponse
}

func NewWeatherResponse(response io.ReadCloser) (*WeatherJSON, error) {
	data, err := ioutil.ReadAll(response)

	if err != nil {
		return nil, err
	}

	var JSONParsed WeatherJSON
	errUnmashal := json.Unmarshal(data, &JSONParsed)
	if errUnmashal != nil {
		return nil, errUnmashal
	}

	return &JSONParsed, nil
}
