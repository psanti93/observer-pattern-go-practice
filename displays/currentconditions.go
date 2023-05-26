package currentconditions

import "fmt"

type CurrentConditions struct {
	Temperature float64
	Humidity    float64
	WeatherData *WeatherData
}

func (c *CurrentConditions) Display() {
	fmt.Printf("Current Conditions: Temperature: %v F degrees and %v  humdity", c.Temperature, c.Humidity)
}

func (c *CurrentConditions) Update() {
	c.Temperature = c.WeatherData.Temperature
	c.Humidity = c.WeatherData.Humidity
	c.Display()
}
