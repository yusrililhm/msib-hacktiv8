package order_repository

import "h8-assignment-2/entity"

type OrderItem struct {
	Order entity.Order
	Item  entity.Item
}

type OrderItemMapped struct {
	Order entity.Order
	Item  []entity.Item `json:"items"`
}

func (oim *OrderItemMapped) HandleMappingOrderWithItem(orderItem []*OrderItem) []*OrderItemMapped {
	orderItemsMapped := []*OrderItemMapped{}

	for _, eachOrderItem := range orderItem {

		isOrderExist := false

		for i := range orderItemsMapped {
			if eachOrderItem.Order.OrderId == orderItem[i].Order.OrderId {
				isOrderExist = true
				orderItemsMapped[i].Item = append(orderItemsMapped[i].Item, eachOrderItem.Item)
				break
			}
		}

		if !isOrderExist {
			orderItemMapped := &OrderItemMapped{
				Order: eachOrderItem.Order,
			}

			orderItemMapped.Item = append(orderItemMapped.Item, eachOrderItem.Item)
			orderItemsMapped = append(orderItemsMapped, orderItemMapped)
		}
	}

	return orderItemsMapped
}
