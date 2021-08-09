package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"microshop/utils"
	"net/http"
)

func GoodsList(c *gin.Context){
	ctx := c.Request.Context()
	cateGoryStr := c.DefaultQuery("categoryId", "")
	cateGory := utils.String2Int(cateGoryStr)
	page, size := common.GetPageParam(c)
	filter := make(map[string]interface{})
	if cateGory != -1 {
		filter["category_id"] = cateGory
	}
	errMsg := ""
	list, err := service.NewGoodsService().List(ctx, page, size, filter)
	if err != nil {
		errMsg = fmt.Sprintf("list catalog err:%s",err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	common.HandleSucc(c, http.StatusOK, "", entity.GoodsListResponse{GoodsList: list})
}

func GoodsDetail(c *gin.Context){
	ctx := c.Request.Context()
	errMsg := ""
	idStr := c.DefaultQuery("id", "")
	id := utils.String2Int(idStr)
	if id == -1 {
		errMsg = fmt.Sprintf("good detail require goods id:%s err", idStr)
		common.HandleError(c, http.StatusBadRequest, common.C_REQ_PARAM_ERR, errMsg)
		return
	}
	goods, err := service.NewGoodsService().GetGoods(ctx, id)
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d detail err:%s", id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if goods == nil {
		errMsg = fmt.Sprintf("get goods:%d not exsit", id)
		common.HandleError(c, http.StatusBadRequest, common.C_GOODS_NOT_EXSIT_ERR, errMsg)
		return
	}
	gallerys, err := service.NewGoodsService().GetGoodsGallerys(ctx, id)
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d gallerys err:%s", id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	attrs, err := service.NewGoodsService().GetGoodsAttrs(ctx, id)
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d attrs err:%s", id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	products, err := service.NewGoodsService().GetGoodsProducts(ctx, id)
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d products err:%s", id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	galleries, imageText := []entity.GoodsGallery{}, []entity.GoodsGallery{}
	for _, item := range gallerys {
		if item.ImgType == 1 {
			galleries = append(galleries, item)
		}else{
			imageText = append(imageText, item)
		}
	}
	common.HandleSucc(c, http.StatusOK, "", entity.GoodsDetailResponse{
		Goods: *goods,
		Galleries: galleries,
		Attribute: attrs,
		ProductList: products,
		ImageText: imageText,
	})
}
