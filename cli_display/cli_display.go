package cli_display

import (
	"fmt"
	"strings"
	"time"
	"xerrex/weather/city_reader"
	"xerrex/weather/station"
)

/* NOTE: when to user pointer or value
1. Use pointers when structs are large or you need to modify the original
2. Use values when structs are small and you want to work with a copy */

func ShowWeatherData(weatherData station.Weather) {
	fmt.Printf("%s  Weather in %s (%s):\n", "🌤️", weatherData.City, weatherData.System.Country)

	latDirection, lonDirection := coordsDirection(weatherData.Coord.Latitude, weatherData.Coord.Longitude)

	fmt.Printf("%s Coordinates: %s %s, %s %s\n", "📍",
		fmt.Sprintf("%.8f", weatherData.Coord.Latitude), latDirection,
		fmt.Sprintf("%.8f", weatherData.Coord.Longitude), lonDirection)

	fmt.Printf("%s  Temperature: %s°C\n", "🌡️", fmt.Sprintf("%.1f", weatherData.Main.TempAvg))
	fmt.Printf("%s Feels like: %s°C\n", "🔥", fmt.Sprintf("%.1f", weatherData.Main.FeelsLike))
	fmt.Printf("%s Min/Max: %s°C / %s°C\n", "📊", fmt.Sprintf("%.1f", weatherData.Main.TempLow), fmt.Sprintf("%.1f", weatherData.Main.TempHigh))
	fmt.Printf("%s Humidity: %s%%\n", "💧", fmt.Sprintf("%d", weatherData.Main.Humidity))
	fmt.Printf("%s Pressure: %s hPa\n", "📏", fmt.Sprintf("%d", weatherData.Main.ATMPressure))

	if len(weatherData.Description) > 0 {
		fmt.Printf("%s  Description: %s(%s)\n", "☁️", weatherData.Description[0].Short, weatherData.Description[0].Long)
	}
	fmt.Printf("%s Wind Speed: %s m/s\n", "💨", fmt.Sprintf("%.1f", weatherData.Wind.Speed))
	fmt.Printf("%s Wind Degrees: %s°\n", "🧭", fmt.Sprintf("%.1f", weatherData.Wind.Degrees))
	fmt.Printf("%s Visibility: %s metres\n", "👀", fmt.Sprintf("%d", weatherData.Visibility))

	fmt.Printf("%s TimeZone: %s\n", "🕐", timezoneFormatter(weatherData.Timezone))
	fmt.Printf("%s Sunrise: %s \n", "🌅", convertTimestampsTo24h(weatherData.System.Sunrise))
	fmt.Printf("%s Sunset: %s\n", "🌇", convertTimestampsTo24h(weatherData.System.Sunset))

}

func coordsDirection(latitude float32, longitude float32) (string, string) {
	var latDirection string = "N"
	var lonDirection string = "E"

	if latitude < 0 {
		latDirection = "S"
	}
	if longitude < 0 {
		lonDirection = "W"
	}

	return latDirection, lonDirection
}

func timezoneFormatter(seconds int) string {
	// Consider: Using the time function
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60

	sign := "+"
	if hours < 0 {
		sign = "-"
		hours = -hours // Make hours positive for formatting
	}

	return fmt.Sprintf("UTC%s%02d:%02d", sign, hours, minutes)
}

func convertTimestampsTo24h(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("15:04")
}

func ShowWeatherRawResponse(raw_response string) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("🟦 Raw response { } ")
	fmt.Println(raw_response)
	fmt.Println(strings.Repeat("=", 50) + "\n")
}

func ShowCities(cities []city_reader.City) {

	fmt.Printf("%-4s %-15s %-16s %-16s \n", "#", "Name", "Latitude", "Longitude")
	fmt.Println("------------------------------------------------------------")

	for index, city := range cities {
		fmt.Printf("%-4d %-15s %-16.8f %-16.8f\n", index+1, city.Name, city.Latitude, city.Longitude)
		fmt.Println("------------------------------------------------------------")
	}
}
