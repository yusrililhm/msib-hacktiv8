package controller

import "github.com/gin-gonic/gin"

type CarController interface {
	AddCar(c *gin.Context)
	GetAllCar(c *gin.Context)
	FindCarById(c *gin.Context)
}
