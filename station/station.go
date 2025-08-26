package station

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Date     int64  `json:"dt"`
	Timezone int    `json:"timezone"`
	City     string `json:"name"`
	Coord    struct {
		Longitude float32 `json:"lon"`
		Latitude  float32 `json:"lat"`
	} `json:"coord"`

	System struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`

	Description []struct {
		Short string `json:"main"`
		Long  string `json:"description"`
	} `json:"weather"`

	Main struct {
		TempAvg     float32 `json:"temp"`
		TempHigh    float32 `json:"temp_max"`
		TempLow     float32 `json:"temp_min"`
		ATMPressure int32   `json:"pressure"`
		Humidity    int32   `json:"humidity"`
		FeelsLike   float32 `json:"feels_like"`
	} `json:"main"`

	Wind struct {
		Speed   float32 `json:"speed"`
		Degrees float32 `json:"deg"`
	} `json:"wind"`

	Visibility int `json:"visibility"`
}

func FetchWeather(apiUrl string) (string, Weather, error) {
	// Get the weather at the provided longitude(lon)
	// and latitude(lat)

	resp, err := http.Get(apiUrl)

	if err != nil {
		return "", Weather{}, errors.New("Error occurred when calling the API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorMessage := fmt.Sprintf("Weather API call error with status code '%v'", resp.StatusCode)
		return "", Weather{}, errors.New(errorMessage)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		// panic(err)
		return "", Weather{}, err
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		return "", Weather{}, err
	}

	return string(body), weather, nil
}
