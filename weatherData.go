package main

type WeatherData struct {
	Observers   []Observer
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func (w *WeatherData) RegisterObserver(observer Observer) {
	w.Observers = append(w.Observers, observer)
}

func (w *WeatherData) RemoveObserver(observer Observer) {

}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.Observers {
		observer.Update()
	}

}

func (w *WeatherData) SetMeasurements(temperature float64, humidty float64, pressure float64) {
	w.Temperature = temperature
	w.Humidity = humidty
	w.Pressure = pressure

	w.NotifyObservers()
}
