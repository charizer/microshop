package service

import (
	"context"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type BannerService struct {
	BannerRepo impl.BannerRepo
}

func NewBannerService() BannerService {
	s := BannerService{
		BannerRepo: repository.NewBannerRepo(),
	}
	return s
}

func (o BannerService) List(ctx context.Context, page, size int) ([]entity.Banner, error){
	return o.BannerRepo.List(ctx, page, size)
}
