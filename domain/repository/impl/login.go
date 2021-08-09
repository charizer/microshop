package impl

import (
	"context"
	"microshop/domain/entity"
)

type LoginRepo interface {
	One(ctx context.Context, mobile string) (*entity.UserInfo, error)
	Insert(ctx context.Context, userInfo entity.UserInfo) error
	Update(mobile string, update map[string]interface{}) error
}
