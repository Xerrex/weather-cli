package city_reader

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// City represents the latitude and longitude of a location.
type City struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ReadCitiesJson(file_path string) ([]City, error) {
	fileContent, err := os.ReadFile(file_path)

	if err != nil {
		return nil, err
	}

	var cities []City
	err = json.Unmarshal(fileContent, &cities)

	if err != nil {
		return nil, err
	}

	return cities, nil
}

func DisplayCities(cities []City) {

	for index, city := range cities {
		fmt.Printf("%d.  %s -- latitude: %.8f, longitude: %.8f\n", index+1, city.Name, city.Latitude, city.Longitude)
	}
}

func GetCityDetails(name string, cities []City) (City, error) {
	var cityName string = strings.ToLower(name)

	for _, city := range cities {
		// var city_name string = strings.ToLower(city.Name)
		if cityName == strings.ToLower(city.Name) {
			return city, nil
		}
	}
	return City{}, fmt.Errorf("City named '%s' was not found in default list", name)
}
