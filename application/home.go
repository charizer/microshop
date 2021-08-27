package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"net/http"
)

// 首页
func Home(c *gin.Context) {
	ctx := c.Request.Context()
	errMsg := ""
	banners, err := service.NewGoodsService().BannerGoods(ctx, 0, 3)
	if err != nil {
		errMsg = fmt.Sprintf("list banner err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	newGoods, err := service.NewGoodsService().NewGoods(ctx, 0, 4)
	if err != nil {
		errMsg = fmt.Sprintf("list new goods err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	hotGoods, err := service.NewGoodsService().HotGoods(ctx, 0, 3)
	if err != nil {
		errMsg = fmt.Sprintf("list hot goods err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	common.HandleSucc(c, http.StatusOK, "", entity.HomeResponse{
		Banners: banners,
		Newgoods: newGoods,
		Hotgoods: hotGoods,
	})
}
