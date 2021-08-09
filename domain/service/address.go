package service

import (
	"context"
	"microshop/domain/entity"
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type AddressService struct {
	AddressRepo impl.AddressRepo
}

func NewAddressService() AddressService {
	return AddressService{
		AddressRepo: repository.NewAddressRepo(),
	}
}

func (o AddressService) List(ctx context.Context, page, size int, userId int) ([]entity.Address, error){
	return o.AddressRepo.List(ctx, page, size, userId)
}

func (o AddressService) GetAddress(ctx context.Context, id, userId int) (*entity.Address, error){
	return o.AddressRepo.One(ctx, id, userId)
}

func (o AddressService) AddAddress(ctx context.Context, address entity.Address) (int, error) {
	return o.AddressRepo.Insert(ctx, address)
}

func (o AddressService) UpdateAddress(ctx context.Context, id int, address entity.Address) error {
	return o.AddressRepo.UpdateOne(ctx, id, address)
}

func (o AddressService) UpdateAll(ctx context.Context, where, update map[string]interface{}) error {
	return o.AddressRepo.Update(ctx, where, update)
}

func (o AddressService) DeleteAddress(ctx context.Context, where map[string]interface{}) error {
	return o.AddressRepo.Delete(ctx, where)
}
