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
	product, err := service.NewGoodsService().GetGoodsProduct(ctx, map[string]interface{}{
		"goods_id": req.GoodsId,
		"id":       req.ProductId,
	})
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d product:%d err:%s", req.GoodsId, req.ProductId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if product == nil || product.GoodsNumber < req.Number {
		errMsg = fmt.Sprintf("get goods:%d product:%d not exsit", req.GoodsId, req.ProductId)
		common.HandleError(c, http.StatusBadRequest, common.C_PRODUCT_NOT_EXSIT_ERR, errMsg)
		return
	}
	cart, err := service.NewCartService().GetCart(ctx, map[string]interface{}{
		"user_id":    userId,
		"goods_id":   goods.Id,
		"product_id": product.Id,
	})
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d product:%d cart err:%s", req.GoodsId, req.ProductId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if cart == nil {
		insertCart(c, *goods, *product, req.Number, userId)
	} else {
		updateCart(c, *goods, *product, *cart, req.Number, userId)
	}
}

func insertCart(c *gin.Context, goods entity.Goods, product entity.GoodsProduct, number, userId int) {
	ctx := c.Request.Context()
	errMsg := ""
	now := time.Now().UnixNano() / 1e6
	cart := entity.Cart{
		GoodsId:     goods.Id,
		ProductId:   product.Id,
		GoodsName:   goods.Name,
		ListPicUrl:  goods.ListPicUrl,
		Number:      number,
		UserId:      userId,
		RetailPrice: product.RetailPrice,
		Checked:     1,
		GoodsBrief:  goods.GoodsBrief,
		CreateTime:  now,
		UpdateTime:  now,
	}
	_, err := service.NewCartService().AddCart(ctx, cart)
	if err != nil {
		errMsg = fmt.Sprintf("add cart goods:%d product:%d err:%s", goods.Id, product.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	} else {
		common.HandleSucc(c, http.StatusOK, "", gin.H{})
	}
}

func updateCart(c *gin.Context, goods entity.Goods, product entity.GoodsProduct, cart entity.Cart, number, userId int) {
	ctx := c.Request.Context()
	errMsg := ""
	now := time.Now().UnixNano() / 1e6
	if product.GoodsNumber < (number + cart.Number) {
		errMsg = fmt.Sprintf("update cart:%d goods:%d product:%d number less", cart.Id, goods.Id, product.Id)
		common.HandleError(c, http.StatusBadRequest, common.C_PRODUCT_NOT_EXSIT_ERR, errMsg)
		return
	}
	err := service.NewCartService().UpdateCart(ctx, map[string]interface{}{
		"id": cart.Id,
	}, map[string]interface{}{
		"number":      cart.Number + number,
		"update_time": now,
	})
	if err != nil {
		errMsg = fmt.Sprintf("update cart:%d goods:%d product:%d err:%s", cart.Id, goods.Id, product.Id, err.Error())
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
	productId := utils.GetIntParam(c, "productId")
	if goodsId == -1 || productId == -1 {
		errMsg = fmt.Sprintf("get cart count param err")
		common.HandleError(c, http.StatusBadRequest, common.C_REQ_PARAM_ERR, errMsg)
		return
	}
	cart, err := service.NewCartService().GetCart(ctx, map[string]interface{}{
		"goods_id":   goodsId,
		"product_id": productId,
		"user_id":    userId,
	})
	if err != nil {
		errMsg = fmt.Sprintf("get cart count goods:%d product:%d err:%s", goodsId, productId, err.Error())
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
	product, err := service.NewGoodsService().GetGoodsProduct(ctx, map[string]interface{}{
		"id":       req.ProductId,
	})
	if product == nil || product.GoodsNumber < req.Number {
		errMsg = fmt.Sprintf("update cart before get goods:%d product:%d not exsit", req.GoodsId, req.ProductId)
		common.HandleError(c, http.StatusBadRequest, common.C_PRODUCT_NOT_EXSIT_ERR, errMsg)
		return
	}
	err = service.NewCartService().UpdateCart(ctx, map[string]interface{}{
		"id": req.Id,
	}, map[string]interface{}{
		"number":      req.Number,
		"update_time": now,
	})
	if err != nil {
		errMsg = fmt.Sprintf("update cart:%d goods:%d product:%d err:%s", cart.Id, req.GoodsId, product.Id, err.Error())
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
