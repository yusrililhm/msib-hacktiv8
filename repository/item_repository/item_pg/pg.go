package item_pg

import (
	"database/sql"
	"fmt"

	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/item_repository"
)

type itemRepositoryImpl struct {
	db *sql.DB
}

func NewItemRepositoryImpl(db *sql.DB) item_repository.ItemRepository {
	return &itemRepositoryImpl{
		db: db,
	}
}

func (itemPg itemRepositoryImpl) generateByItemsCodesQuery(itemAmount int) string {
	baseQuery := `
		SELECT
			item_id,
			item_code,
			order_id,
			quantity,
			description
		FROM
			items
		WHERE
			item_code
		IN
	`

	statement := "("

	for i := 1; i <= itemAmount; i++ {
		if i == itemAmount {
			statement += fmt.Sprintf("$%d)", i)
			break
		}

		statement += fmt.Sprintf("$%d, ", i)
	}
	
	return fmt.Sprintf("%s %s", baseQuery, statement)
}

// GetItemsByCode implements item_repository.ItemRepository.
func (itemPg *itemRepositoryImpl) GetItemsByCode(itemsCode []any) ([]*entity.Item, errs.Error) {
	query := itemPg.generateByItemsCodesQuery(len(itemsCode))
	row, err := itemPg.db.Query(query, itemsCode...)

	if err != nil {
		return nil, errs.NewInternalServerError("something wrong")
	}

	items := []*entity.Item{}

	for row.Next() {
		item := entity.Item{}
		err = row.Scan(
			&item.ItemId,
			&item.ItemCode,
			&item.OrderId,
			&item.Quantity,
			&item.Description,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something wrong")
		}

		items = append(items, &item)
	}

	return items, nil
}
