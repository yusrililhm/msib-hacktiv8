package handler

import (
	"log"
	"time"
	
	"hacktiv8-assignment-3/service"
)

func StartApplication() {

	ws := service.NewWeatherService()
	h := newWeatherHandler(ws)

	for {
		time.Sleep(2 * time.Second)
		log.Println(string(h.weatherHandler()))
	}
}
