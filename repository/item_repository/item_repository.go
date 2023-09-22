package item_repository

import (
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
)

type ItemRepository interface {
	GetItemsByCode(itemsCode []any) ([]*entity.Item, errs.Error)
}
