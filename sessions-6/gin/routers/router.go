package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/car", controllers.CreateCar)
	r.PUT("/car/:id", controllers.UpdateCar)
	r.GET("/car/:id", controllers.GetCar)
	r.GET("/car", controllers.GetAllCar)
	r.DELETE("/car/:id", controllers.DeleteCar)

	return r
}
