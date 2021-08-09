package impl

import (
	"context"
	"microshop/domain/entity"
)

type OrderRepo interface {
	List(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.Order, error)
	One(ctx context.Context, filter map[string]interface{}) (*entity.Order, error)
	Insert(ctx context.Context, cart entity.Order) (int, error)
	Update(ctx context.Context, filter, update map[string]interface{}) error
	Delete(ctx context.Context, filter map[string]interface{}) error
	InsertOrderGoods(ctx context.Context, orderGoods entity.OrderGoods) (int, error)
	ListOrderGoods(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.OrderGoods, error)
}
