package impl

import "context"

const (
	CART_C_NAME = "cart"
)

type CartRepo interface {
	List(ctx context.Context) error
}
