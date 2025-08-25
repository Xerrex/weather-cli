package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"xerrex/display"
	"xerrex/station"
)

func main() {
	fmt.Println("Welcome to weather CLI")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	latitude := -1.2970091674162187
	longitude := 36.821917335677504

	WEATHER_API_KEY := os.Getenv("WEATHER_API_KEY")

	var BASE_URL string = "https://api.openweathermap.org/data/2.5/weather?"
	FULL_URL := fmt.Sprintf("%slat=%f&lon=%f&appid=%s&units=metric", BASE_URL, latitude, longitude, WEATHER_API_KEY)

	raw_response, weather, err := station.Fetchweather(latitude, longitude, FULL_URL)

	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Raw response")
	fmt.Println(raw_response)
	fmt.Println(strings.Repeat("=", 50) + "\n")

	display.ShowWeatherData(weather)

}
