package impl

import (
	"context"
	"microshop/domain/entity"
)


type BannerRepo interface {
	List(ctx context.Context, page, size int)([]entity.Banner, error)
}
