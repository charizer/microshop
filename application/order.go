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

func OrderSubmit(c *gin.Context) {
	var req entity.SubmitOrderReq
	c.BindJSON(&req)
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	addr, err := service.NewAddressService().GetAddress(ctx, req.AddressId, userId)
	if err != nil {
		errMsg = fmt.Sprintf("order submit get address:%d err:%s", req.AddressId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if addr == nil {
		errMsg = fmt.Sprintf("order submit get address:%d not exsit", req.AddressId)
		common.HandleError(c, http.StatusBadRequest, common.C_ADDR_NOT_EXSIT_ERR, errMsg)
		return
	}
	if req.GoodsId != 0 {
		submitQuick(c, *addr, req.GoodsId, req.Number, userId)
	} else {
		submitNormal(c, *addr, userId)
	}
}

func submitQuick(c *gin.Context, address entity.Address, goodsId, number, userId int) {
	ctx := c.Request.Context()
	errMsg := ""
	goods, err := service.NewGoodsService().GetGoods(ctx, goodsId)
	if err != nil {
		errMsg = fmt.Sprintf("get goods:%d detail err:%s", goodsId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if goods == nil {
		errMsg = fmt.Sprintf("get goods:%d not exsit", goodsId)
		common.HandleError(c, http.StatusBadRequest, common.C_GOODS_NOT_EXSIT_ERR, errMsg)
		return
	}
	var freightPrice float64 = 0
	var goodstotalprice float64 = 0
	goodstotalprice += float64(number) * goods.RetailPrice
	ordertotalprice := goodstotalprice + freightPrice
	actualprice := ordertotalprice - 0
	now := time.Now().UnixNano() / 1e6
	order := entity.Order{
		OrderSn:        common.GenerateOrderNumber(),
		UserId:         userId,
		Consignee:      address.Name,
		Mobile:         address.Mobile,
		Address:        address.Address,
		FreightPrice:   0,
		OrderPrice:     ordertotalprice,
		ActualPrice:    actualprice,
		ProvinceName:   address.ProvinceName,
		CityName:       address.CityName,
		DistrictName:   address.DistrictName,
		CallbackStatus: "true",
		CreateTime:     now,
		UpdateTime:     now,
	}
	orderId, err := service.NewOrderService().AddOrder(ctx, order)
	if err != nil {
		errMsg = fmt.Sprintf("submit order add order err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	orderGoods := entity.OrderGoods{
		OrderId:     orderId,
		GoodsId:     goods.Id,
		GoodsName:   goods.Name,
		ListPicUrl:  goods.ListPicUrl,
		RetailPrice: goods.RetailPrice,
		Number:      number,
		GoodsBrief:  goods.GoodsBrief,
		CreateTime:  now,
		UpdateTime:  now,
	}
	_, err = service.NewOrderService().AddOrderGoods(ctx, orderGoods)
	if err != nil {
		errMsg = fmt.Sprintf("submit quick order add order goods err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	order.Id = orderId
	common.HandleSucc(c, http.StatusOK, "", order)
}

func submitNormal(c *gin.Context, address entity.Address, userId int) {
	ctx := c.Request.Context()
	errMsg := ""
	carts, err := service.NewCartService().AllCart(ctx, map[string]interface{}{
		"user_id": userId,
		"checked": 1,
	})
	if err != nil {
		errMsg = fmt.Sprintf("submit normal get cart err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if len(carts) <= 0 {
		errMsg = fmt.Sprintf("submit normal cart empty", err.Error())
		common.HandleError(c, http.StatusBadRequest, common.C_CART_EMPTY_ERR, errMsg)
		return
	}
	var freightPrice float64 = 0
	var goodstotalprice float64 = 0
	for _, cart := range carts {
		goodstotalprice += float64(cart.Number) * cart.RetailPrice
	}
	ordertotalprice := goodstotalprice + freightPrice
	actualprice := ordertotalprice - 0
	now := time.Now().UnixNano() / 1e6
	order := entity.Order{
		OrderSn:        common.GenerateOrderNumber(),
		UserId:         userId,
		Consignee:      address.Name,
		Mobile:         address.Mobile,
		Address:        address.Address,
		FreightPrice:   0,
		OrderPrice:     ordertotalprice,
		ActualPrice:    actualprice,
		ProvinceName:   address.ProvinceName,
		CityName:       address.Address,
		DistrictName:   address.DistrictName,
		CallbackStatus: "true",
		CreateTime:     now,
		UpdateTime:     now,
	}
	orderId, err := service.NewOrderService().AddOrder(ctx, order)
	if err != nil {
		errMsg = fmt.Sprintf("submit normal order add order err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	for _, item := range carts {
		orderGoods := entity.OrderGoods{
			OrderId:    orderId,
			GoodsId:    item.GoodsId,
			GoodsName:  item.GoodsName,
			ListPicUrl: item.ListPicUrl,
			RetailPrice: item.RetailPrice,
			Number:      item.Number,
			GoodsBrief:  item.GoodsBrief,
			CreateTime:  now,
			UpdateTime:  now,
		}
		_, err = service.NewOrderService().AddOrderGoods(ctx, orderGoods)
		if err != nil {
			errMsg = fmt.Sprintf("submit normal order add order goods err:%s", err.Error())
			log.Errorln(errMsg)
		}
	}
	err = service.NewCartService().DeleteCart(ctx, map[string]interface{}{
		"user_id": userId,
		"checked": 1,
	})
	if err != nil {
		errMsg = fmt.Sprintf("submit normal del cart goods err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	order.Id = orderId
	common.HandleSucc(c, http.StatusOK, "", order)
}

func OrderList(c *gin.Context) {
	page, size := common.GetPageParam(c)
	status := utils.GetIntParam(c, "status")
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	response := []entity.OrderInfo{}
	list := []entity.Order{}
	var err error
	switch status {
	case entity.ORDER_ALL:
		list, err = service.NewOrderService().ListOrder(ctx, page, size, map[string]interface{}{
			"user_id": userId,
			"order_status <>": entity.ORDER_DELETE,
		})
	case entity.ORDER_FINISH:
		list, err = service.NewOrderService().ListOrder(ctx, page, size, map[string]interface{}{
			"user_id": userId,
			"order_status in": []int{entity.ORDER_CANCEL, entity.ORDER_SUCC},
		})
	default:
		list, err = service.NewOrderService().ListOrder(ctx, page, size, map[string]interface{}{
			"user_id": userId,
			"order_status": status,
		})
	}
	if err != nil {
		errMsg = fmt.Sprintf("order list err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if len(list) <= 0 {
		common.HandleSucc(c, http.StatusOK, "", response)
		return
	}
	orderIds := make([]int,0,len(list))
	for _, order := range list {
		orderIds = append(orderIds, order.Id)
	}
	orderGoodsList, err := service.NewOrderService().ListOrderGoods(ctx, 0, 100, map[string]interface{}{
		"order_id in": orderIds,
	})
	if err != nil {
		errMsg = fmt.Sprintf("OrderList list order goods err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	hash := getOrderGoods(orderGoodsList)
	for _, order := range list {
		orderDate := utils.FormatTimestamp(order.CreateTime, "2006-01-02 15:04:05")
		orderStatusText := getOrderStatusText(order)
		rsp := entity.OrderInfo{order, hash[order.Id], len(hash[order.Id]), orderStatusText , orderDate}
		response = append(response, rsp)
	}
	common.HandleSucc(c, http.StatusOK, "", response)
}

func getOrderStatusText(order entity.Order) string {
	statustext := "待付款"
	switch order.OrderStatus {
	case entity.ORDER_WAIT_PAY:
		statustext = "待付款"
	case entity.ORDER_WAIT_RECV:
		statustext = "待收货"
	case entity.ORDER_CANCEL:
		statustext = "已取消"
	case entity.ORDER_SUCC:
		statustext = "交易成功"
	}
	return statustext
}

func getOrderGoods(goods []entity.OrderGoods) map[int][]entity.OrderGoods{
	hash := make(map[int][]entity.OrderGoods)
	for _, good := range goods {
		if _, ok := hash[good.OrderId]; ok {
			hash[good.OrderId] = append(hash[good.OrderId], good)
		}else{
			hash[good.OrderId] = []entity.OrderGoods{}
			hash[good.OrderId] = append(hash[good.OrderId], good)
		}
	}
	return hash
}

func OrderDetail(c *gin.Context) {
	orderId := utils.GetIntParam(c, "orderId")
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	ctx := c.Request.Context()
	errMsg := ""
	order, err := service.NewOrderService().GetOrder(ctx, map[string]interface{}{
		"id": orderId,
		"user_id": userId,
	})
	if err != nil {
		errMsg = fmt.Sprintf("OrderDetail get order:%d err:%s", orderId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if order == nil {
		errMsg = fmt.Sprintf("OrderDetail get order:%d not exsit", orderId, err.Error())
		common.HandleError(c, http.StatusBadRequest, common.C_ORDER_NOT_EXSIT_ERR, errMsg)
		return
	}
	orderGoodsList, err := service.NewOrderService().ListOrderGoods(ctx, 0, 100, map[string]interface{}{
		"order_id": order.Id,
	})
	if err != nil {
		errMsg = fmt.Sprintf("OrderDetail list order goods err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	orderStatusText := getOrderStatusText(*order)
	orderDate := utils.FormatTimestamp(order.CreateTime, "2006-01-02 15:04:05")
	response := entity.OrderInfo{*order, orderGoodsList, len(orderGoodsList), orderStatusText , orderDate}
	common.HandleSucc(c, http.StatusOK, "", response)
}

func OrderStatusUpdate(c *gin.Context) {
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	ctx := c.Request.Context()
	errMsg := ""
	var req entity.UpdateOrderStatusReq
	c.BindJSON(&req)
	order, err := service.NewOrderService().GetOrder(ctx, map[string]interface{}{
		"id": req.OrderId,
		"user_id": userId,
	})
	if err != nil {
		errMsg = fmt.Sprintf("OrderStatusUpdate get order:%d err:%s", req.OrderId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if order == nil {
		errMsg = fmt.Sprintf("OrderStatusUpdate get order:%d not exsit", req.OrderId)
		common.HandleError(c, http.StatusBadRequest, common.C_ORDER_NOT_EXSIT_ERR, errMsg)
		return
	}
	err = service.NewOrderService().UpdateOrder(ctx, map[string]interface{}{
		"id": req.OrderId,
		"user_id": userId,
	},map[string]interface{}{
		"order_status": req.Status,
		"update_time": time.Now().UnixNano()/1e6,
	})
	if err != nil {
		errMsg = fmt.Sprintf("OrderStatusUpdate update order:%d err:%s", req.OrderId, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	common.HandleSucc(c, http.StatusOK, "", gin.H{})
}
