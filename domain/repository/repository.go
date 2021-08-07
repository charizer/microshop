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

