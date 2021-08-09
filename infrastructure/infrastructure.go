package infrastructure

import (
	"microshop/infrastructure/cache"
	"microshop/infrastructure/db"
	"microshop/infrastructure/logger"
)

var (
	log = logger.GetLogger()
)

func StartUp(){
	cache.StartUp()
	db.StartUp()
}
