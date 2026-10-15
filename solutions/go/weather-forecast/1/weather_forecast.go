// Package weather provides tools to weather.
package weather


var (
	// CurrentCondition init.
	CurrentCondition string
	// CurrentLocation init.
	CurrentLocation  string
)

// Forecast return the formart string about somewhere's weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
