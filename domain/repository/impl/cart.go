package impl

import (
	"context"
	"microshop/domain/entity"
)

type CartRepo interface {
	List(ctx context.Context, page, size, userId int) ([]entity.Cart, error)
	All(ctx context.Context, filter map[string]interface{}) ([]entity.Cart, error)
	One(ctx context.Context, filter map[string]interface{}) (*entity.Cart, error)
	Insert(ctx context.Context, cart entity.Cart) (int, error)
	Update(ctx context.Context, filter, update map[string]interface{}) error
	Delete(ctx context.Context, filter map[string]interface{}) error
	Count(ctx context.Context, filter map[string]interface{}) (int64, error)
}
