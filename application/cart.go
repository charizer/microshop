package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"microshop/utils"
	"net/http"
	"time"
)

func CartList(c *gin.Context) {
	ctx := c.Request.Context()
	page, size := common.GetPageParam(c)
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	errMsg := ""
	list, err := service.NewCartService().ListCart(ctx, page, size, userId)
	if err != nil {
		errMsg = fmt.Sprintf("list cart err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	total := entity.CartTotal{}
	for _, cart := range list {
		total.GoodsCount += cart.Number
		total.GoodsAmount += float64(cart.Number) * cart.RetailPrice
		if cart.Checked == 1 {
			total.CheckedGoodsCount += cart.Number
			total.CheckedGoodsAmount += float64(cart.Number) * cart.RetailPrice
		}
	}
	common.HandleSucc(c, http.StatusOK, "", entity.CartListResponse{
		CartList:  list,
		CartTotal: total,
	})
}

func CartAdd(c *gin.Context) {
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	var req entity.CartAddReq
	c.BindJSON(&req)
	goods, err := service.NewGoodsService().GetGoods(ctx, req.GoodsId)
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d detail err:%s", req.GoodsId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if goods == nil {
		errMsg = fmt.Sprintf("get goods:%d not exsit", req.GoodsId)
		common.HandleError(c, http.StatusBadRequest, common.C_GOODS_NOT_EXSIT_ERR, errMsg)
		return
	}
	cart, err := service.NewCartService().GetCart(ctx, map[string]interface{}{
		"user_id":    userId,
		"goods_id":   goods.Id,
	})
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d cart err:%s", req.GoodsId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if cart == nil {
		insertCart(c, *goods, req.Number, userId)
	} else {
		updateCart(c, *goods, *cart, req.Number, userId)
	}
}

func insertCart(c *gin.Context, goods entity.Goods, number, userId int) {
	ctx := c.Request.Context()
	errMsg := ""
	now := time.Now().UnixNano() / 1e6
	cart := entity.Cart{
		GoodsId:     goods.Id,
		GoodsName:   goods.Name,
		ListPicUrl:  goods.ListPicUrl,
		Number:      number,
		UserId:      userId,
		RetailPrice: goods.RetailPrice,
		Checked:     1,
		GoodsBrief:  goods.GoodsBrief,
		CreateTime:  now,
		UpdateTime:  now,
	}
	_, err := service.NewCartService().AddCart(ctx, cart)
	if err != nil {
		errMsg = fmt.Sprintf("add cart goods:%d err:%s", goods.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{})
	}
}

func updateCart(c *gin.Context, goods entity.Goods, cart entity.Cart, number, userId int) {
	ctx := c.Request.Context()
	errMsg := ""
	now := time.Now().UnixNano() / 1e6
	err := service.NewCartService().UpdateCart(ctx, map[string]interface{}{
		"id": cart.Id,
	}, map[string]interface{}{
		"number":      cart.Number + number,
		"update_time": now,
	})
	if err != nil {
		errMsg = fmt.Sprintf("update cart:%d goods:%d err:%s", cart.Id, goods.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{})
	}
}

func CartCount(c *gin.Context) {
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	goodsId := utils.GetIntParam(c, "goodsId")
	if goodsId == -1 {
		errMsg = fmt.Sprintf("get cart count param err")
		common.HandleError(c, http.StatusBadRequest, common.C_REQ_PARAM_ERR, errMsg)
		return
	}
	cart, err := service.NewCartService().GetCart(ctx, map[string]interface{}{
		"goods_id":   goodsId,
		"user_id":    userId,
	})
	if err != nil {
		errMsg = fmt.Sprintf("get cart count goods:%d err:%s", goodsId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if cart == nil {
		common.HandleSucc(c, http.StatusOK, "", gin.H{"count": 0})
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{"count": cart.Number})
	}
}

func CartUpdate(c *gin.Context) {
	now := time.Now().UnixNano() / 1e6
	var req entity.CartUpdateReq
	c.BindJSON(&req)
	ctx := c.Request.Context()
	errMsg := ""
	cart, err := service.NewCartService().GetCart(ctx, map[string]interface{}{
		"id": req.Id,
	})
	if err != nil {
		errMsg = fmt.Sprintf("update cart before get cart id:%d err:%s",req.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if cart == nil {
		errMsg = fmt.Sprintf("update cart before get cart id:%d not exist", cart.Id)
		common.HandleError(c, http.StatusBadRequest, common.C_CART_NOT_EXSIT_ERR, errMsg)
		return
	}
	err = service.NewCartService().UpdateCart(ctx, map[string]interface{}{
		"id": req.Id,
	}, map[string]interface{}{
		"number":      req.Number,
		"update_time": now,
	})
	if err != nil {
		errMsg = fmt.Sprintf("update cart:%d goods:%d err:%s", cart.Id, req.GoodsId,  err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{})
	}
}

func CartDelete(c *gin.Context) {
	var req entity.CartDeleteReq
	c.BindJSON(&req)
	ctx := c.Request.Context()
	errMsg := ""
	err := service.NewCartService().DeleteCart(ctx, map[string]interface{}{
		"id": req.Id,
	})
	if err != nil {
		errMsg = fmt.Sprintf("delete cart:%d err:%s", req.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{})
	}
}

func CartChecked(c *gin.Context) {
	var req entity.CartCheckedReq
	c.BindJSON(&req)
	ctx := c.Request.Context()
	errMsg := ""
	err := service.NewCartService().UpdateCart(ctx, map[string]interface{}{
		"id in": req.Carts,
	},map[string]interface{}{
		"checked": req.IsChecked,
	})
	if err != nil {
		errMsg = fmt.Sprintf("checked err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{})
	}
}

func UserCartCount(c *gin.Context) {
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	count, err := service.NewCartService().CountCart(ctx, map[string]interface{}{
		"userId": userId,})
	if err != nil {
		errMsg = fmt.Sprintf("count cart err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	}else{
		common.HandleSucc(c, http.StatusOK, "", gin.H{"count": count})
	}
}

func CartCheckout(c *gin.Context) {
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	var address entity.Address
	addrs, err := service.NewAddressService().List(ctx, 0, 10, userId)
	if err != nil {
		errMsg = fmt.Sprintf("cart checkout get address list err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if len(addrs) > 0 {
		address = addrs[0]
	}
	list, err := service.NewCartService().ListCart(ctx, 0, 20, userId)
	if err != nil {
		errMsg = fmt.Sprintf("cart checkout list car err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	total := entity.CartTotal{}
	checkedCarts := []entity.Cart{}
	for _, cart := range list {
		total.GoodsCount += cart.Number
		total.GoodsAmount += float64(cart.Number) * cart.RetailPrice
		if cart.Checked == 1 {
			total.CheckedGoodsCount += cart.Number
			total.CheckedGoodsAmount += float64(cart.Number) * cart.RetailPrice
			checkedCarts = append(checkedCarts, cart)
		}
	}
	common.HandleSucc(c, http.StatusOK, "", entity.CheckOutCartResponse{
		Address: address,
		CheckedGoodsList: checkedCarts,
		OrderTotalPrice: total.CheckedGoodsAmount,
	})
}
