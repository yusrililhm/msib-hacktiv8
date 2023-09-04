package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
}
