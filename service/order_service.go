package service

import (
	"fmt"
	
	"h8-assignment-2/dto"
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/item_repository"
	"h8-assignment-2/repository/order_repository"

	"net/http"
)

type OrderService interface {
	Add(newOrderRequest *dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error)
	Fetch() (*dto.GetOrdersResponse, errs.Error)
	Edit(orderId int, updateOrderRequest *dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error)
	Remove(orderId int) errs.Error
}

type OrderServiceImpl struct {
	orderRepository order_repository.OrderRepository
	itemRepository  item_repository.ItemRepository
}

func NewOrderServiceImpl(orderRepository order_repository.OrderRepository, itemRepository item_repository.ItemRepository) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
		itemRepository:  itemRepository,
	}
}

// Add implements OrderService.
func (orderServiceImpl *OrderServiceImpl) Add(newOrderRequest *dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error) {

	orderPayload := &entity.Order{
		OrderedAt:    newOrderRequest.OrderedAt,
		CustomerName: newOrderRequest.CustomerName,
	}

	itemPayload := []*entity.Item{}

	for _, eachItemPayload := range newOrderRequest.Items {
		item := &entity.Item{
			ItemCode:    eachItemPayload.ItemCode,
			Quantity:    eachItemPayload.Quantity,
			Description: eachItemPayload.Description,
		}

		itemPayload = append(itemPayload, item)
	}

	err := orderServiceImpl.orderRepository.Create(orderPayload, itemPayload)

	if err != nil {
		return nil, err
	}

	return &dto.NewOrderResponse{
		StatusCode: http.StatusCreated,
		Message:    "new order successfully created",
		Data:       nil,
	}, nil
}

// Edit implements OrderService.
func (orderServiceImpl *OrderServiceImpl) Edit(orderId int, updateOrderRequest *dto.NewOrderRequest) (*dto.NewOrderResponse, errs.Error) {

	_, err := orderServiceImpl.orderRepository.ReadOrderById(orderId)

	if err != nil {
		return nil, err
	}

	itemCodes := []any{}

	for _, eachItem := range updateOrderRequest.Items {
		itemCodes = append(itemCodes, eachItem.ItemCode)
	}

	items, err := orderServiceImpl.itemRepository.GetItemsByCode(itemCodes)

	if err != nil {
		return nil, err
	}

	for _, eachItemFromRequest := range updateOrderRequest.Items {

		isFound := false

		for _, eachItem := range items {

			if orderId != eachItem.OrderId {
				return nil, errs.NewBadRequestError(fmt.Sprintf("item with item code %s doesn't belong to the order with id %d", eachItem.ItemCode, orderId))
			}

			if eachItem.ItemCode == eachItemFromRequest.ItemCode {
				isFound = true
				break
			}
		}

		if !isFound {
			return nil, errs.NewNotFoundError(fmt.Sprintf("item with item code %s is not found", eachItemFromRequest.ItemCode))
		}
	}

	orderPayload := &entity.Order{
		OrderId:      orderId,
		OrderedAt:    updateOrderRequest.OrderedAt,
		CustomerName: updateOrderRequest.CustomerName,
	}

	itemPayload := []*entity.Item{}

	for _, eachItemFromRequest := range updateOrderRequest.Items {

		item := &entity.Item{
			ItemCode:    eachItemFromRequest.ItemCode,
			Description: eachItemFromRequest.Description,
			Quantity:    eachItemFromRequest.Quantity,
		}

		itemPayload = append(itemPayload, item)
	}

	err = orderServiceImpl.orderRepository.Update(orderId, orderPayload, itemPayload)

	if err != nil {
		return nil, err
	}

	return &dto.NewOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "order successfully updated",
		Data:       nil,
	}, nil
}

// Fetch implements OrderService.
func (orderServiceImpl *OrderServiceImpl) Fetch() (*dto.GetOrdersResponse, errs.Error) {

	orders, err := orderServiceImpl.orderRepository.Read()

	if err != nil {
		return nil, err
	}

	data := []dto.OrderWithItems{}

	// melakukan iterasi slice of pointer struct OrderItemMapped
	for _, eachOrder := range orders {
		fetchOrder := dto.OrderWithItems{
			OrderId:      eachOrder.Order.OrderId,
			CustomerName: eachOrder.Order.CustomerName,
			OrderedAt:    eachOrder.Order.OrderedAt,
			Items:        []dto.GetItemResponse{},
			CreatedAt:    eachOrder.Order.CreatedAt,
			UpdatedAt:    eachOrder.Order.UpdatedAt,
		}

		for _, eachItem := range eachOrder.Item {
			item := dto.GetItemResponse{
				ItemId:      eachItem.ItemId,
				ItemCode:    eachItem.ItemCode,
				OrderId:     eachItem.OrderId,
				Quantity:    eachItem.Quantity,
				Description: eachItem.Description,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
			}

			fetchOrder.Items = append(fetchOrder.Items, item)
		}

		data = append(data, fetchOrder)
	}

	return &dto.GetOrdersResponse{
		StatusCode: http.StatusOK,
		Message:    "orders successfully fetched",
		Data:       data,
	}, nil
}

// Remove implements OrderService.
func (orderServiceImpl *OrderServiceImpl) Remove(orderId int) errs.Error {

	order, err := orderServiceImpl.orderRepository.ReadOrderById(orderId)

	if err != nil {
		return err
	}

	if order.OrderId != orderId {
		return err
	}

	err = orderServiceImpl.orderRepository.Delete(orderId)

	if err != nil {
		return err
	}

	return nil
}
