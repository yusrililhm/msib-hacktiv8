package middleware

import (
	"go-jwt/database"
	"go-jwt/entity"
	"go-jwt/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyTkn, err := helper.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   helper.Unathorized,
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", verifyTkn)
		ctx.Next()
	}
}

func ProductAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.NewConn()

		if err != nil {
			return
		}

		productId, _ := strconv.Atoi(c.Param("productId"))
		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		product := &entity.Product{}

		if err := db.Select("user_id").First(product, uint(productId)).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "data not found",
			})
			return
		}

		if product.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": helper.Unathorized,
			})
			return
		}

		c.Next()
	}
}
