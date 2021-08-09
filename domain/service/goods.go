package service

import (
	"context"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type GoodsService struct {
	GoodsRepo impl.GoodsRepo
}

func NewGoodsService() GoodsService {
	s := GoodsService{
		GoodsRepo: repository.NewGoodsRepo(),
	}
	return s
}

func (o GoodsService) List(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.SampleGoods, error){
	return o.GoodsRepo.List(ctx, page, size, filter)
}

func (o GoodsService) HotGoods(ctx context.Context, page, size int) ([]entity.SampleGoods, error){
	filter := make(map[string]interface{})
	filter["is_hot"] = 1
	return o.GoodsRepo.List(ctx, page, size, filter)
}

func (o GoodsService) NewGoods(ctx context.Context, page, size int) ([]entity.SampleGoods, error){
	filter := make(map[string]interface{})
	filter["is_new"] = 1
	return o.GoodsRepo.List(ctx, page, size, filter)
}

func (o GoodsService) GetGoods(ctx context.Context, id int) (*entity.Goods, error){
	return o.GoodsRepo.One(ctx, id)
}

func (o GoodsService) GetGoodsGallerys(ctx context.Context, id int)([]entity.GoodsGallery, error){
	return o.GoodsRepo.GetGallerys(ctx, 0, 10, id)
}

func (o GoodsService) GetGoodsAttrs(ctx context.Context, id int)([]entity.GoodsAttr, error){
	return o.GoodsRepo.GetAttrs(ctx, 0, 10, id)
}

func (o GoodsService) GetGoodsProducts(ctx context.Context, id int)([]entity.GoodsProduct, error){
	return o.GoodsRepo.GetProducts(ctx, 0, 10, id)
}

func (o GoodsService) GetGoodsProduct(ctx context.Context, filter map[string]interface{}) (*entity.GoodsProduct, error){
	return o.GoodsRepo.GetProduct(ctx, filter)
}
