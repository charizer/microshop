package infrastructure

import (
	"microshop/infrastructure/cache"
	"microshop/infrastructure/config"
	"microshop/infrastructure/db"
	"microshop/infrastructure/logger"
)

var (
	log = logger.GetLogger()
	cfg = config.GetConfig()
)

func StartUp(){
	cache.StartUp()
	db.StartUp()
}
