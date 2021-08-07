package mysql

import (
	"context"
	"microshop/domain/repository/impl"
)

var (
	_ impl.CartRepo = new(CartRepo)
)

type CartRepo struct {
}

func (o CartRepo) List(ctx context.Context) error {
	return nil
}
