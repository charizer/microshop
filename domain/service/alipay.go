package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"microshop/domain/entity"
	"microshop/infrastructure/payment"
)

type AlipayService struct {

}

func (o AlipayService) Pay(ctx context.Context, req entity.SubmitPayReq) (string, error){
	return payment.GetInstance().Pay(ctx, req)
}

func (o AlipayService) ParseNotify(c *gin.Context) (entity.NotifyPayResp, error) {
	return payment.GetInstance().ParseNotify(c)
}
