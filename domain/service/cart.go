package service

import (
	"microshop/domain/repository"
	"microshop/domain/repository/impl"
)

type CartService struct {
	CartRepo impl.CartRepo
}

func NewCartService() *CartService {
	s := &CartService{
		CartRepo: repository.NewCartRepo(),
	}
	return s
}
