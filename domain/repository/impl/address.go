package impl

import (
	"context"
	"microshop/domain/entity"
)

type AddressRepo interface {
	List(ctx context.Context, page, size int, userId int) ([]entity.Address, error)
	One(ctx context.Context, id, userId int) (*entity.Address, error)
	Insert(ctx context.Context, address entity.Address) (int, error)
	UpdateOne(ctx context.Context, id int, address entity.Address) error
	Update(ctx context.Context, filter, update map[string]interface{}) error
	Delete(ctx context.Context, filter map[string]interface{}) error
}
