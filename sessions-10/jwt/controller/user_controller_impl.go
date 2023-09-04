package controller

import (
	"errors"
	"go-jwt/helper"
	"go-jwt/model"
	"go-jwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// UserRegister implements UserController.
// UserRegister godoc
// @Summary	User register
// @Description	Endpoint for User register
// @Tags	order
// @Accept	json
// @Produce	json
// @Param	model.UserRegister body model.UserRegister true "body request for user register"
// @Success	201     {object} model.UserRegisterResponse
// @Failure	400     {object} helper.Response
// @Router	/user/register [post]
func (controller *UserControllerImpl) UserRegister(c *gin.Context) {
	user := &model.UserRegister{}

	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, helper.Response{Message: helper.BadRequest})
		return
	}

	res, err := controller.UserService.UserRegisterService(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.Response{Message: helper.BadRequest})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UserLogin implements UserController.
func (controller *UserControllerImpl) UserLogin(c *gin.Context) {
	user := &model.UserLogin{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("bad request"))
		return
	}

	data, err := controller.UserService.UserLoginService(user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.Response{Message: helper.Unathorized})
		return
	}

	token := helper.GenerateToken(data.Id, data.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
