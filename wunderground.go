package wunderground

// Report represents a wunderground response body
type Report struct {
	TemperatureFahrenheit float32 `json:"temp_f"`
	TemperatureString     string  `json:"temperature_string"`
}

// GetWeatherConditions returns a report of the current weather conditions
//  for a city and state.
func GetWeatherConditions(city, state) {
	return nil
}
