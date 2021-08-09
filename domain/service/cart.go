package service

import (
	"context"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type CartService struct {
	CartRepo impl.CartRepo
}

func NewCartService() CartService {
	return CartService{
		CartRepo: repository.NewCartRepo(),
	}
}

func (o CartService) ListCart(ctx context.Context, page, size, userId int) ([]entity.Cart, error) {
	return o.CartRepo.List(ctx, page, size, userId)
}

func (o CartService) AllCart(ctx context.Context, filter map[string]interface{}) ([]entity.Cart, error) {
	return o.CartRepo.All(ctx, filter)
}

func (o CartService) GetCart(ctx context.Context, filter map[string]interface{}) (*entity.Cart, error){
	return o.CartRepo.One(ctx, filter)
}

func (o CartService) AddCart(ctx context.Context, cart entity.Cart) (int, error) {
	return o.CartRepo.Insert(ctx, cart)
}

func (o CartService) UpdateCart(ctx context.Context, where, update map[string]interface{}) error {
	return o.CartRepo.Update(ctx, where, update)
}

func (o CartService) DeleteCart(ctx context.Context, where map[string]interface{}) error {
	return o.CartRepo.Delete(ctx, where)
}