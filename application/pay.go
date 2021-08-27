package application

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"net/http"
	"time"
)

func PayNotify(c *gin.Context) {
	payResult, err := service.AlipayService{}.ParseNotify(c)
	go updateOrder(context.Background(), payResult)
	if err != nil {
		c.JSON(http.StatusOK, "fail")
	}else{
		c.JSON(http.StatusOK, "success")
	}
}

func updateOrder(ctx context.Context, payResult entity.NotifyPayResp){
	var err error
	if payResult.Status == entity.PAY_STATUS_SUCC {
		err = service.NewOrderService().UpdateOrder(ctx, map[string]interface{}{
			"order_sn": payResult.OrderSn,
		},map[string]interface{}{
			"pay_status": payResult.Status,
			"order_status": entity.ORDER_WAIT_RECV,
			"update_time": time.Now().UnixNano()/1e6,
		})
	}else{
		err = service.NewOrderService().UpdateOrder(ctx, map[string]interface{}{
			"order_sn": payResult.OrderSn,
		},map[string]interface{}{
			"pay_status": payResult.Status,
			"update_time": time.Now().UnixNano()/1e6,
		})
	}
	if err != nil {
		log.Errorf("update order:%s after pay err:%s", payResult.OrderSn, err.Error())
	}else{
		log.Infof("update order:%s after pay succ", payResult.OrderSn)
	}
}

func Pay(c *gin.Context) {
	var req entity.SubmitPayReq
	c.BindJSON(&req)
	ctx := c.Request.Context()
	errMsg := ""
	data, err := service.AlipayService{}.Pay(ctx, req)
	if err != nil {
		errMsg = fmt.Sprintf("do pay err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError,common.S_ALIPAY_REQ_ERR, errMsg)
	}else{
		common.HandleSucc(c, http.StatusOK, "", data)
	}
}
