package impl

import (
	"context"
	"microshop/domain/entity"
)

type CatalogRepo interface {
	List(ctx context.Context, page, size int)([]entity.Catalog, error)
	One(ctx context.Context, id int) (*entity.Catalog, error)
}
