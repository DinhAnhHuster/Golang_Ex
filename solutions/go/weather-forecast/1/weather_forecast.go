// Package weather provides tools to forecast the current weather conditions for cities in Goblinocus.
package weather

var (
	// CurrentCondition stores the current weather condition (e.g. "sunny", "rainy") for the last forecast location.
	CurrentCondition string
	// CurrentLocation stores the city name for which the current condition was last recorded.
	CurrentLocation string
)

// Forecast returns a formatted forecast string for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

