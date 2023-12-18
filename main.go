package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Timezone  int64  `json:"timezone"`
		Date_time int64  `json:"dt"`
		Details   struct {
			Country string `json:"country"`
			Sunrise int64  `json:"sunrise"`
			Sunset  int64  `json:"sunset"`
		} `json:"sys"`
	}
	// Forecast []struct {
	// 	Brief       string `json:"main"`
	// 	Description string `json:"description"`
	// 	Icon        string `json:"icon"`
	// } `json:"weather"`

	// Main struct {
	// 	Temp_High float32 `json:"temp_max"`
	// 	Temp_Low  float32 `json:"temp_min"`
	// 	Humidity  int     `json:"humidity"`
	// } `json:"main"`
}

func main() {
	fmt.Println("Weather CLI")

	var latitude, longitude float32 = -1.292066, 36.821945
	var API_KEY string = "1bb59edb0bef71820e6a62339ec24627"
	var BASE_URL string = "https://api.openweathermap.org/data/2.5/weather?"
	// var BASE_URL string = "https://api.openweathermap.org/data/3.0/onecall?lat=-1.292066&lon=36.821945"
	FULL_URL := fmt.Sprintf("%slat=%f&lon=%f&appid=%s&units=metric", BASE_URL, latitude, longitude, API_KEY)

	resp, err := http.Get(FULL_URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Call to Weather API is not Available")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(weather)

}
