package fakes

// Report represents a faked weather report
type Report struct {
	TemperatureFahrenheit float32 `json:"temp_f"`
	TemperatureString     string  `json:"temperature_string"`
}

// GetWeatherConditions returns a dummy weather conditions report
func GetWeatherConditions(city string, state string) Report {
	report := Report{
		TemperatureFahrenheit: 59.3,
		TemperatureString:     "59.3 F (15.2 C)",
	}

	return report
}
