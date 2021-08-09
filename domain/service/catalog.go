package service

import (
	"context"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type CatalogService struct {
	CatalogRepo impl.CatalogRepo
}

func NewCatalogService() CatalogService {
	s := CatalogService{
		CatalogRepo: repository.NewCatalogRepo(),
	}
	return s
}

func (o CatalogService) List(ctx context.Context, page, size int) ([]entity.Catalog, error){
	return o.CatalogRepo.List(ctx, page, size)
}

func (o CatalogService) One(ctx context.Context, id int) (*entity.Catalog, error){
	return o.CatalogRepo.One(ctx, id)
}
