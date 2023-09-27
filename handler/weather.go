package handler

import (
	"encoding/json"
	"hacktiv8-assignment-3/entity"
	"hacktiv8-assignment-3/service"
	"log"
	"math/rand"
	"time"
)

type handler struct {
	service service.WeatherService
}

func newWeatherHandler(ws service.WeatherService) handler {
	return handler{
		service: ws,
	}
}

func (h *handler) weatherHandler() []byte {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().Unix()))

	weatherPayload := &entity.Weather{
		Water: randomizer.Intn(20),
		Wind:  randomizer.Intn(20),
	}

	response := h.service.Update(weatherPayload)

	b, err := json.Marshal(response)

	if err != nil {
		log.Panicln(err.Error())
		return nil
	}

	return b
}
