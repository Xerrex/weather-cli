package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	} `json:"main"`

	Wind struct {
		Speed   float32 `json:"speed"`
		Degrees float32 `json:"deg"`
	} `json:"wind"`
}

func main() {
	fmt.Println("Weather CLI: Nairobi(-1.292066, 36.821945)")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var latitude, longitude float32 = -1.292066, 36.821945
	WEATHER_API_KEY := os.Getenv("WEATHER_API_KEY")
	var BASE_URL string = "https://api.openweathermap.org/data/2.5/weather?"
	FULL_URL := fmt.Sprintf("%slat=%f&lon=%f&appid=%s&units=metric", BASE_URL, latitude, longitude, WEATHER_API_KEY)

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

	// fmt.Println(weather)

	city := weather.City
	country := weather.System.Country
	current := weather.Description[0].Long
	tempAvg := weather.Main.TempAvg
	date := time.Unix(weather.Date, 0).Local()
	fmt.Printf("%s, %s(%.0f°C): %s\n", city, country, tempAvg, current)
	fmt.Printf("Today %s \n", date)

	tempHigh := weather.Main.TempHigh
	tempLow := weather.Main.TempLow

	fmt.Printf("Temperature range %.0f - %.0f\n", tempLow, tempHigh)

}
