package main

import (
	"github.com/gin-gonic/gin"
	"microshop/infrastructure"
	"microshop/infrastructure/config"
	"microshop/infrastructure/logger"
	"microshop/router"
)

var (
	log = logger.GetLogger()
	cfg = config.GetConfig()
)

func main() {
	infrastructure.StartUp()
	gin.SetMode(cfg.Mode)
	g := gin.Default()
	router.Load(g)
	err := g.Run(":" + cfg.HttpPort)
	if err != nil {
		log.Errorln("start err:", err.Error())
	}
}