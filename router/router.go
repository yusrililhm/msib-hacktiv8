package router

import (
	"hacktiv8-assignment-3/config"
	"hacktiv8-assignment-3/controller"
	"hacktiv8-assignment-3/entity"
	"hacktiv8-assignment-3/repository"
	"hacktiv8-assignment-3/service"
	"log"

	"github.com/gin-gonic/gin"
)

var PORT = config.Init("PORT")

func StartServer()  {
	db, err := config.NewConn()

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	db.AutoMigrate(&entity.Weather{})

	d, _ := db.DB()
	defer d.Close()

	weatherRepository := repository.NewWeatherRepository(db)
	weatherService := service.NewWeatherService(weatherRepository)
	weatherController := controller.NewWeatherController(weatherService)
	
	go weatherController.WeatherUpdate()

	router := gin.Default()
	router.Run(":" + PORT)
}
