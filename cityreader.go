package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
)

// City represents the latitude and longitude of a location.
type CityDetails struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}


func readCitiesDetails(citiesJsonFilePath string) ([]CityDetails, error) {
	// ReadCitiesDetails

	fileContent, err := ioutil.ReadFile(citiesJsonFilePath)
	if err != nil {
		return nil, err
	}

	var cities []CityDetails
	err = json.Unmarshal(fileContent, &cities)

	if err != nil {
		return nil, err
	}
	return cities, nil
}


func showCitiesList(cities []CityDetails){
	
	for _, city := range cities {
		fmt.Printf("%s -- latitude: %.4f, longitude: %.4f\n", city.Name, city.Latitude, city.Longitude)
	}

}


func getCityDetails(cityName sting, cities []cityDetails)(CityDetails, error){

	var cityNameLower string = strings.ToLower(cityName)

	for _, city := range cities {
		var city_name string = strings.ToLower(city.Name)
		if cityNameLower == city_name {
			return city, nil
	}
	return CityDetails{}, fmt.Errorf("City not found: %s", cityName)
}


func main() {

	cities, err := readCitiesDetails("./cities.json")

	
	if err != nil {
		panic(err)
	}

	// showCitiesList(cities)

	city, error2 := getCityDetails("Nairobi", cities)

	if error2 != nil {
		panic(error2)
	}

	fmt.Printf("%s -- latitude: %.4f, longitude: %.4f\n", city.Name, city.Latitude, city.Longitude)
}


