package router

import (
	"go-swagger/controller"
	_ "go-swagger/docs"

	"github.com/gin-gonic/gin"
	gimSwagger "github.com/swaggo/gin-swagger"
	swaggerFile "github.com/swaggo/files"
)

// @title Car API
// @version 1.0
// @description An API for Car

// @contact.name Yusril Ilham
// @contact.email yusrililham62@gmail.com
// @contact.url https://github.com/yusrililhm

// @host localhost:9000
// @BasePath /

func router(controller controller.CarController) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", gimSwagger.WrapHandler(swaggerFile.Handler))

	car := r.Group("car")

	{
		car.POST("/", controller.AddCar)
		car.GET("/", controller.GetAllCar)
		car.GET("/:id", controller.FindCarById)
	}

	return r
}

func StartServer()  {
	carController := controller.NewCarController()
	r := router(carController)
	r.Run(":9000")
}