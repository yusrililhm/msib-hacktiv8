package service

import (
	"hacktiv8-assignment-3/dto"
	"hacktiv8-assignment-3/entity"
)

type WeatherService interface {
	Update(weatherPayload *entity.Weather) *dto.GetWeatherResponse
}

type weatherServiceImpl struct {
}

func NewWeatherService() WeatherService {
	return &weatherServiceImpl{}
}

// Update implements WeatherService.
func (ws *weatherServiceImpl) Update(weatherPayload *entity.Weather) *dto.GetWeatherResponse {
	waterResponse := dto.WaterResponse{Water: weatherPayload.Water}
	windResponse := dto.WindResponse{Wind: weatherPayload.Wind}

	waterResponse.WaterStatus()
	windResponse.WindStatus()

	return &dto.GetWeatherResponse{
		WaterResponses: waterResponse,
		WindResponses: windResponse,
	}
}
