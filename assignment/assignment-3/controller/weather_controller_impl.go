package controller

import (
	"hacktiv8-assignment-3/service"
)

type WeatherControllerImpl struct {
	service.WeatherService
}

func NewWeatherController(weatherService service.WeatherService) WeatherController {
	return &WeatherControllerImpl{
		WeatherService: weatherService,
	}
}

// WeatherUpdate implements WeatherController.
func (controller *WeatherControllerImpl) WeatherUpdate() {
	controller.WeatherService.WeatherUpdateService()
}
