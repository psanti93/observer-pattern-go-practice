package main

type WeatherData struct {
	Observer
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func (w *WeatherData) RegisterObserver() {
}

func (w *WeatherData) RemoveObserver() {

}

func (w *WeatherData) NotifyObservers() {

	w.Observer.Update()

}

func (w *WeatherData) SetMeasurements(temperature float64, humidty float64, pressure float64) {
	w.Temperature = temperature
	w.Humidity = humidty
	w.Pressure = pressure

	w.NotifyObservers()
}
