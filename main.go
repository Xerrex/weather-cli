package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)


type CityDetails struct{
	name string
	latitude float64
	longitude float64
}

var CitiesDetails = map[string]CityDetails{
	"mombasa": CityDetails{"Mombasa", -4.044420488531392, 39.67001728536259},
	"nairobi": CityDetails{"Nairobi", -1.2970091674162187, 36.821917335677504},
	"naivasha": CityDetails{"Naivasha", -0.7185111907604552, 36.440378974982345},
	"nakuru": CityDetails{"Nakuru", -0.3017192761345334, 36.062820762552924},
	"eldoret": CityDetails{"Eldoret", 0.5444355122581617, 35.250484945199055},
	"kisumu": CityDetails{"Kisumu", -0.09311678011290307, 34.752140939624525},
	"thika": CityDetails{"Thika", -0.9850818030818558, 37.04363218548822},
	"nyahururu": CityDetails{"Nyahururu", 0.07355733472573683, 36.35617270988145},
	"nanyuki": CityDetails{"Nanyuki", 0.05293356117980731, 37.02300840122002},
	"marsabit": CityDetails{"Marsabit", 2.4170771425557778, 37.99232626182558},
	"garissa": CityDetails{"Garissa", -0.4145353024440266, 39.621605219013645},
}



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


func printCityList(){
	fmt.Println("\nList of cities:")

	for _, cityDetail := range CitiesDetails {
		fmt.Printf("%s -- latitude: %.4f, longitude: %.4f\n", cityDetail.name, cityDetail.latitude, cityDetail.longitude)
	}
}



func getCityCoordinates(cityName string) (CityDetails, error){
	//get city coordinates from the name
	cityDetails, found := CitiesDetails[cityName]
	if !found {
		return CityDetails{}, fmt.Errorf("City not found: %s", cityName)
	}
	return cityDetails, nil
}


func printHelp(){
	fmt.Println("Usage: weather_snack [options] - works after build")
	fmt.Println("or")
	fmt.Println("go run main.go [options]")

	fmt.Println("Options:")
	fmt.Println(" --help, -h 		Show this help message")
	fmt.Println(" --city-list, -l 		Show the city list")
  fmt.Println(" cityName 		Show the weather in the cityName, ie. Nairobi")
}


func getCityWeather(cityDetails CityDetails){
	// Get the city weather

	// var latitude, longitude float32 = -1.292066, 36.821945
	latitude, longitude := cityDetails.latitude, cityDetails.longitude

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

	city := weather.City
	country := weather.System.Country
	current := weather.Description[0].Long
	tempAvg := weather.Main.TempAvg
	date := time.Unix(weather.Date, 0).Local()
	fmt.Printf("%s, %s(%.0f°C): %s\n", city, country, tempAvg, current)
	fmt.Printf("Today %s \n", date)

	tempHigh := weather.Main.TempHigh
	tempLow := weather.Main.TempLow

	fmt.Printf("Temperature range %.0f - %.0f\n\n", tempLow, tempHigh)
}

func main() {
	fmt.Println("Weather CLI")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help", "-h":
			printHelp()
			return
		case "--city-list", "-l":
			printCityList()
			return
		default:
			// If none of the recognized options, treat it as a city name
			cityName := strings.ToLower(os.Args[1])
			cityDetails, err := getCityCoordinates(cityName)

			if err != nil {
				panic(err)
			}
			getCityWeather(cityDetails)
			return
		}
	}

	// If no arguments print help message Usage message.
	printHelp()
}
