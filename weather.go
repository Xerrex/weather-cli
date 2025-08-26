package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"xerrex/weather/city_reader"
	"xerrex/weather/cli_display"
	"xerrex/weather/station"

	// "strings"
	// "xerrex/weather/cli_display"
	// "xerrex/weather/station"
	"github.com/joho/godotenv"
)

func main() {

	cmdLineArgs := os.Args
	if len(cmdLineArgs) < 2 {
		handleHelpCmd()
		return
	}

	const CITIES_JSON_FILE string = "cities.json"

	command := cmdLineArgs[1]

	switch command {
	case "--help", "-h":
		handleHelpCmd()
		return
	case "--cities", "-c":
		handleCitiesCmd(CITIES_JSON_FILE)
		return
	default:
		var cmdOpt string = ""
		if len(cmdLineArgs) > 2 {
			cmdOpt = cmdLineArgs[2]
		}

		handleCityCmd(command, cmdOpt, CITIES_JSON_FILE)

	}
}

func handleHelpCmd() {
	fmt.Println("Welcome to weather CLI")
	fmt.Println("\nUsage:")
	fmt.Println("weather [command] [options]")

	fmt.Println("\nCommands:")
	fmt.Println("--help/-h	Show this help message")
	fmt.Println("--cities/-l	Show this help message")
	fmt.Println("<city>		Show the weather in the city ie. Nairobi")

	fmt.Println("\nOptions:")
	fmt.Println("-f	Show the weather in the city ie. Nairobi & json response")

	fmt.Println("\nExamples:")
	fmt.Println("weather --help")
	fmt.Println("weather -h")
	fmt.Println("weather --cities")
	fmt.Println("weather -l")
	fmt.Println("weather Nairobi")
	fmt.Println("weather Nairobi -f")

}

func handleCitiesCmd(citiesJsonFile string) {
	cities, err := city_reader.ReadCitiesJson(citiesJsonFile)
	if err != nil {
		panic(err)
	}
	city_reader.DisplayCities(cities)
}

func handleCityCmd(cityName string, cmdOpt string, citiesJsonFile string) {

	cities, err := city_reader.ReadCitiesJson(citiesJsonFile)
	if err != nil {
		panic(err)
	}
	city, err := city_reader.GetCityDetails(cityName, cities)

	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("WEATHER_API_KEY")
	BASE_URL := os.Getenv("WEATHER_API_BASE_URL")
	lat := city.Latitude
	lon := city.Longitude
	FULL_URL := fmt.Sprintf("%slat=%f&lon=%f&appid=%s&units=metric", BASE_URL, lat, lon, API_KEY)

	raw_response, weather, err := station.FetchWeather(FULL_URL)

	if err != nil {
		panic(err)
	}

	if cmdOpt == "-f" {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Raw response")
		fmt.Println(raw_response)
		fmt.Println(strings.Repeat("=", 50) + "\n")
	}

	cli_display.ShowWeatherData(weather)
}
