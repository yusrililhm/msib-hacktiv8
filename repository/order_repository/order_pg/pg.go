package order_pg

import (
	"database/sql"
	"errors"

	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/order_repository"
)

type OrderRepositoryImpl struct {
	db *sql.DB
}

const (
	createOrderQuery = `
		INSERT INTO
			orders
			(ordered_at, customer_name)
		VALUES
			($1, $2)
		RETURNING order_id
	`

	createItemQuery = `
		INSERT INTO
			items
			(item_code, description, quantity, order_id)
		VALUES
			($1, $2, $3, $4)
	`

	fetchOrderQuery = `
		SELECT
			o.order_id,
			o.customer_name,
			o.ordered_at,
			o.created_at,
			o.updated_at,
			i.item_id,
			i.item_code,
			i.order_id,
			i.description,
			i.quantity,
			i.created_at,
			i.updated_at
		FROM
			orders AS o
		LEFT JOIN
			items AS i
		ON
			o.order_id = i.order_id
		ORDER BY
			o.order_id
		ASC
	`

	fetchOrderByIdQuery = `
		SELECT 
			order_id,
			customer_name,
			ordered_at
		FROM
			orders
		WHERE
			order_id = $1
	`

	updateOrderQuery = `
		UPDATE
			orders
		SET
			customer_name = $2,
			ordered_at = $3,
			updated_at = now()
		WHERE
			order_id = $1
	`

	updateItemQuery = `
		UPDATE
			items
		SET
			description = $2,
			quantity = $3
		WHERE
			item_id = $1
	`

	deleteOrderQuery = `
		DELETE FROM
			orders
		WHERE
			order_id = $1
	`
)

func NewOrderRepositoryImpl(db *sql.DB) order_repository.OrderRepository {
	return &OrderRepositoryImpl{
		db: db,
	}
}

// Create implements order_repository.OrderRepository.
func (orderRepositoryImpl *OrderRepositoryImpl) Create(orderPayload *entity.Order, itemsPayload []*entity.Item) errs.Error {
	// start transaction
	tx, err := orderRepositoryImpl.db.Begin()

	// rollback if found error
	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	// return id
	var orderId int

	// exec query
	result := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)

	// rollback if found error
	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	// get returning id
	err = result.Scan(&orderId)

	// rollback if found error
	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	for _, v := range itemsPayload {
		// exec query
		_, err = tx.Exec(createItemQuery, v.ItemCode, v.Description, v.Quantity, orderId)

		// rollback if found error
		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError("something went wrong")
		}
	}

	// commit
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	// return error as nil if transaction success
	return nil
}

// Delete implements order_repository.OrderRepository.
func (orderRepositoryImpl *OrderRepositoryImpl) Delete(orderId int) errs.Error {
	tx, err := orderRepositoryImpl.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteOrderQuery, orderId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

// Read implements order_repository.OrderRepository.
func (orderRepositoryImpl *OrderRepositoryImpl) Read() ([]*order_repository.OrderItemMapped, errs.Error) {
	row, err := orderRepositoryImpl.db.Query(fetchOrderQuery)
	var orderItems []*order_repository.OrderItem

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for row.Next() {
		orderItem := order_repository.OrderItem{}
		err = row.Scan(
			&orderItem.Order.OrderId,
			&orderItem.Order.CustomerName,
			&orderItem.Order.OrderedAt,
			&orderItem.Order.CreatedAt,
			&orderItem.Order.UpdatedAt,
			&orderItem.Item.ItemId,
			&orderItem.Item.ItemCode,
			&orderItem.Item.OrderId,
			&orderItem.Item.Description,
			&orderItem.Item.Quantity,
			&orderItem.Item.CreatedAt,
			&orderItem.Item.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("order not found")
			}
			return nil, errs.NewInternalServerError("something went wrong")
		}

		orderItems = append(orderItems, &orderItem)
	}

	result := order_repository.OrderItemMapped{}
	return result.HandleMappingOrderWithItem(orderItems), nil
}

// ReadOrderById implements order_repository.OrderRepository.
func (orderRepositoryImpl *OrderRepositoryImpl) ReadOrderById(orderId int) (*entity.Order, errs.Error) {
	row := orderRepositoryImpl.db.QueryRow(fetchOrderByIdQuery, orderId)
	order := entity.Order{}

	err := row.Scan(
		&order.OrderId,
		&order.CustomerName,
		&order.OrderedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("order not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &order, nil
}

// Update implements order_repository.OrderRepository.
func (orderRepositoryImpl *OrderRepositoryImpl) Update(orderId int, orderPayload *entity.Order, itemPayload []*entity.Item) errs.Error {
	tx, err := orderRepositoryImpl.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(updateOrderQuery, orderId, orderPayload.CustomerName, orderPayload.OrderedAt)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	for _, eachItem := range itemPayload {
		_, err = tx.Exec(updateItemQuery, eachItem.ItemId, eachItem.Description, eachItem.Quantity)

		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError("something went wrong")
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
