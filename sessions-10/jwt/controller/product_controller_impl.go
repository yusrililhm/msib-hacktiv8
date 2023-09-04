package controller

import (
	"go-jwt/model"
	"go-jwt/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ProductControllerImpl struct {
	service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

// CreateProduct implements ProductController.
func (controller *ProductControllerImpl) CreateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)

	product := &model.ProductModel{}
	userId := uint(userData["id"].(float64))

	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	data, err := controller.ProductService.CreateProductService(userId, product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, data)
}

// UpdateProduct implements ProductController.
func (controller *ProductControllerImpl) UpdateProduct(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	product := &model.ProductModel{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))

	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	data, err := controller.ProductService.UpdateProductService(uint(productId), userId, product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
