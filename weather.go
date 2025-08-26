package main

import (
	"fmt"
	"log"
	"os"
	"xerrex/weather/city_reader"
	"xerrex/weather/cli_display"
	"xerrex/weather/station"

	"github.com/joho/godotenv"
)

const WEATHER_BASE_URL = "https://api.openweathermap.org/data/2.5/weather"

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
	case "city":
		var cityName string = ""
		var cmdOpt string = ""

		if len(cmdLineArgs) > 2 {
			cityName = cmdLineArgs[2]
		}

		if len(cmdLineArgs) > 3 {
			cmdOpt = cmdLineArgs[3]
		}

		handleCityCmd(cityName, cmdOpt)
		return
	default:
		var cmdOpt string = ""
		if len(cmdLineArgs) > 2 {
			cmdOpt = cmdLineArgs[2]
		}

		handleCityNameCmd(command, cmdOpt, CITIES_JSON_FILE)

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
	fmt.Println("city		Show the weather in the city provided as an option")

	fmt.Println("\nOptions:")
	fmt.Println("-f	Show the weather in the city ie. Nairobi & json response")
	fmt.Println("<city>	Name of city to show the weather ie. Nairobi")

	fmt.Println("\nExamples:")
	fmt.Println("weather --help")
	fmt.Println("weather -h")
	fmt.Println("weather --cities")
	fmt.Println("weather -l")
	fmt.Println("weather Nairobi")
	fmt.Println("weather Nairobi -f")
	fmt.Println("weather city Nairobi -f")

}

func handleCitiesCmd(citiesJsonFile string) {
	cities, err := city_reader.ReadCitiesJson(citiesJsonFile)
	if err != nil {
		panic(err)
	}

	cli_display.ShowCities(cities)
}

func handleCityNameCmd(cityName string, cmdOpt string, citiesJsonFile string) {

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

	lat := city.Latitude
	lon := city.Longitude
	FULL_URL := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s&units=metric", WEATHER_BASE_URL, lat, lon, API_KEY)

	cityWeatherHandler(FULL_URL, cmdOpt)
}

func handleCityCmd(cityName string, cmdOpt string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("WEATHER_API_KEY")
	FULL_URL := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", WEATHER_BASE_URL, cityName, API_KEY)

	cityWeatherHandler(FULL_URL, cmdOpt)
}

func cityWeatherHandler(weatherURL string, cmdOpt string) {

	raw_response, weather, err := station.FetchWeather(weatherURL)

	if err != nil {
		panic(err)
	}

	if cmdOpt == "-f" {
		cli_display.ShowWeatherRawResponse(raw_response)
	}

	cli_display.ShowWeatherData(weather)
}
