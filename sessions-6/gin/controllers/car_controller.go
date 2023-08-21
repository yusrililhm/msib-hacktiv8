package controllers

import (
	"belajar-gin/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCar(c *gin.Context) {
	// body request
	car := model.Car{}

	// validate request
	if err := c.ShouldBindJSON(&car); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// generate id
	car.Id = fmt.Sprintf("c%d", len(model.CarData)+1)
	model.CarData = append(model.CarData, car)

	// response
	c.JSON(
		http.StatusCreated,
		gin.H{
			"car": car,
		},
	)
}

func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	condition := false
	car := model.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, v := range model.CarData {
		if v.Id == id {
			condition = true
			model.CarData[i] = car
			model.CarData[i].Id = id
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errror":        "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully update", id),
	})
}

func GetCar(c *gin.Context) {
	id := c.Param("id")
	condition := false
	car := model.Car{}

	for i, v := range model.CarData {
		if v.Id == id {
			condition = true
			car = model.CarData[i]
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errror":        "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	car := model.Car{}
	condition := false
	carIndex := 0

	for i, v := range model.CarData {
		if v.Id == id {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errror":        "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", id),
		})
		return
	}

	copy(model.CarData[carIndex:], model.CarData[carIndex+1:])
	model.CarData[len(model.CarData)-1] = car
	model.CarData = model.CarData[:len(model.CarData)-1]

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("delete car id %v success", id),
	})
}

func GetAllCar(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"car": model.CarData,
	})
}
