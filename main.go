package main

func main() {
	weatherData := &WeatherData{
		Observers:   []Observer{},
		Temperature: 80.0,
		Humidity:    65.0,
		Pressure:    30.0,
	}
	currentConditions := &CurrentConditions{
		WeatherData: weatherData,
	}

	weatherData.RegisterObserver(currentConditions)

}
