package handler

import (
	"h8-assignment-2/dto"
	"h8-assignment-2/service"
	"h8-assignment-2/pkg/errs"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: orderService,
	}
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create new orders
// @Tags Orders
// @Accept json
// @Produce json
// @Param dto.NewOrderRequest body dto.NewOrderRequest true "body request for create new order"
// @Success 201 {object} dto.NewOrderResponse
// @Router /orders [post]
func (orderHandler *orderHandler) CreateOrder(ctx *gin.Context) {
	// body request
	newOrderRequest := dto.NewOrderRequest{}

	// validate body request
	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	// call method Add
	response, err := orderHandler.OrderService.Add(&newOrderRequest)

	// validate error and response error
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// response success
	ctx.JSON(response.StatusCode, response)
}

// FetchOrder godoc
// @Summary Fetch order
// @Description Get all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetOrdersResponse
// @Router /orders [get]
func (orderHandler *orderHandler) FetchOrder(ctx *gin.Context) {
	// call method Fetch
	order, err := orderHandler.OrderService.Fetch()

	// validate and response error
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// response success
	ctx.JSON(order.StatusCode, order)
}

// UpdateOrder godoc
// @Summary Edit order
// @Description Update order by order id
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderId path int true "order id"
// @Param dto.NewOrderRequest body dto.NewOrderRequest true "update order body request"
// @Success 200 {object} dto.NewOrderResponse
// @Router /orders/{orderId} [put]
func (orderHandler *orderHandler) UpdateOrder(ctx *gin.Context) {
	// cast param string to int
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))

	// body request for update orders
	updateOrderRequest := dto.NewOrderRequest{}

	// validate body request
	if err := ctx.ShouldBindJSON(&updateOrderRequest); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	// call method Edit
	response, err := orderHandler.OrderService.Edit(orderId, &updateOrderRequest)

	// validate and respoonse error
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// response success
	ctx.JSON(response.StatusCode, response)
}

// CreateOrder godoc
// @Summary Remove order
// @Description Delete order by order id
// @Tags Orders
// @Accept json
// @Produce json
// @Param orderId path int true "order id"
// @Success 200
// @Router /orders/{orderId} [delete]
func (orderHandler *orderHandler) DeleteOrder(ctx *gin.Context) {
	// cast param string to int
	orderId, _ := strconv.Atoi(ctx.Param("orderId"))

	// call method Remove
	err := orderHandler.OrderService.Remove(orderId)

	// validate and response error
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// response success
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete order success",
	})
}
