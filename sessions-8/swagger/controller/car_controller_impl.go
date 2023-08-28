package controller

import (
	"go-swagger/entity"
	"go-swagger/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarControllerImpl struct {
}

// AddCar implements CarController.
// AddCar godoc
// @Summary Add car
// @Description Add new car
// @Tags car
// @Accept json
// @Produce json
// @Param entity.Car body entity.Car true "body request for add new car"
// @Success 201 {object} entity.Car
// @Router /car [post]
func (controller *CarControllerImpl) AddCar(c *gin.Context) {
	car := entity.Car{}

	if err := c.ShouldBindJSON(&car); err != nil {
		return
	}

	car.Id = uint(len(model.Cars) + 1)
	model.Cars = append(model.Cars, car)

	c.JSON(http.StatusCreated, car)
}

// FindCarById implements CarController.
// FindCarById godoc
// @Summary Find car
// @Description Find car by id
// @Tags car
// @Accept json
// @Produce json
// @Param id path int true "id for the car"
// @Success 200 {object} entity.Car
// @Router /car/{id} [get]
func (controller *CarControllerImpl) FindCarById(c *gin.Context) {
	paramId := c.Param("id")
	id, _ := strconv.Atoi(paramId)

	car := entity.Car{}

	for _, v := range model.Cars {
		if v.Id == uint(id) {
			car = entity.Car{
				Id: v.Id,
				Name: v.Name,
				Price: v.Price,
			}
		}
	}

	c.JSON(http.StatusOK, car)
}

// GetAllCar implements CarController.
// GetAllCar godoc
// @Summary Get car
// @Description Get all car
// @Tags car
// @Accept json
// @Produce json
// @Success 200 {object} model.Response
// @Router /car [get]
func (controller *CarControllerImpl) GetAllCar(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{Data: model.Cars})
}

func NewCarController() CarController {
	return &CarControllerImpl{}
}
