package service

import (
	"context"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type OrderService struct {
	OrderRepo impl.OrderRepo
}

func NewOrderService() OrderService {
	return OrderService{
		OrderRepo: repository.NewOrderRepo(),
	}
}

func (o OrderService) ListOrder(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.Order, error) {
	return o.OrderRepo.List(ctx, page, size, filter)
}

func (o OrderService) GetOrder(ctx context.Context, filter map[string]interface{}) (*entity.Order, error){
	return o.OrderRepo.One(ctx, filter)
}

func (o OrderService) AddOrder(ctx context.Context, order entity.Order) (int, error) {
	return o.OrderRepo.Insert(ctx, order)
}

func (o OrderService) UpdateOrder(ctx context.Context, where, update map[string]interface{}) error {
	return o.OrderRepo.Update(ctx, where, update)
}

func (o OrderService) DeleteOrder(ctx context.Context, where map[string]interface{}) error {
	return o.OrderRepo.Delete(ctx, where)
}

func (o OrderService) AddOrderGoods(ctx context.Context, orderGoods entity.OrderGoods) (int, error) {
	return o.OrderRepo.InsertOrderGoods(ctx, orderGoods)
}

func (o OrderService) ListOrderGoods(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.OrderGoods, error) {
	return o.OrderRepo.ListOrderGoods(ctx, page, size, filter)
}