package order_repository

import (
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
)

type OrderRepository interface {
	Create(orderPayload *entity.Order, itemsPayload []*entity.Item) errs.Error
	Read() ([]*OrderItemMapped, errs.Error)
	ReadOrderById(orderId int) (*entity.Order, errs.Error)
	Update(orderId int, orderPayload *entity.Order, itemPayload []*entity.Item) errs.Error
	Delete(orderId int) errs.Error
}
