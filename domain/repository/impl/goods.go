package impl

import (
	"context"
	"microshop/domain/entity"
)

type GoodsRepo interface {
	List(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.SampleGoods, error)
	One(ctx context.Context, id int) (*entity.Goods, error)
	GetGallerys(ctx context.Context,page,size,id int)([]entity.GoodsGallery, error)
	GetAttrs(ctx context.Context,page,size,id int)([]entity.GoodsAttr, error)
	GetProducts(ctx context.Context,page,size,id int)([]entity.GoodsProduct, error)
	GetProduct(ctx context.Context,filter map[string]interface{})(*entity.GoodsProduct, error)
}
