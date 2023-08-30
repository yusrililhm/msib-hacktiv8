package repository

import (
	"fmt"
	"hacktiv8-assignment-3/entity"
	"log"
	"time"

	"gorm.io/gorm"
)

type WeatherRepositoryImpl struct {
	*gorm.DB
}

func NewWeatherRepository(db *gorm.DB) WeatherRepository {
	return &WeatherRepositoryImpl{
		DB: db,
	}
}

// UpdateData implements WeatherRepository.
func (repository *WeatherRepositoryImpl) UpdateData(weather entity.Weather) {
	if err := repository.Create(&weather).Error; err != nil {
		return
	}

	msg := fmt.Sprintf("\nWater\t: %d m;\t\tStatus Water\t: %s\nWind\t: %d m/s;\tStatus Wind\t: %s", weather.Water, weather.StatusWater(), weather.Wind, weather.StatusWind())
	log.Println(msg)

	time.Sleep(15 * time.Second)
}
