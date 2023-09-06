package repository

import "hacktiv8-assignment-3/entity"

type WeatherRepository interface {
	UpdateData(weather entity.Weather)
}