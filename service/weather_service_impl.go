package service

import (
	"hacktiv8-assignment-3/entity"
	"hacktiv8-assignment-3/repository"
	"math/rand"
	"time"
)

type WeatherServiceImpl struct {
	repository.WeatherRepository
}

func NewWeatherService(WeatherRepository repository.WeatherRepository) WeatherService {
	return &WeatherServiceImpl{
		WeatherRepository: WeatherRepository,
	}
}

// WeatherUpdate implements WeatherService.
func (service *WeatherServiceImpl) WeatherUpdateService() {
	for {
		
		randomize := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	
		weather := entity.Weather{
			Water: uint(randomize.Intn(100) + 1),
			Wind: uint(randomize.Intn(100) + 1),
		}
	
		service.WeatherRepository.UpdateData(weather)
	}
}
