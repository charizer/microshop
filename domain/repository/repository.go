package repository

import (
	"microshop/domain/repository/impl"
	"microshop/domain/repository/impl/mysql"
	"microshop/infrastructure/config"
)

var (
	cfg = config.GetConfig()
)

func NewCartRepo() impl.CartRepo {
	return mysql.CartRepo{}
}

func NewOrderRepo() impl.OrderRepo {
	return mysql.OrderRepo{}
}

func NewBannerRepo() impl.BannerRepo {
	return mysql.BannerRepo{}
}

func NewCatalogRepo() impl.CatalogRepo {
	return mysql.CatalogRepo{}
}

func NewGoodsRepo() impl.GoodsRepo {
	return mysql.GoodsRepo{}
}

func NewLoginRepo() impl.LoginRepo {
	return mysql.LoginRepo{}
}

func NewAddressRepo() impl.AddressRepo {
	return mysql.AddressRepo{}
}


