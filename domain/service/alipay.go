package service

import (
	"github.com/gin-gonic/gin"
	"microshop/infrastructure/payment"
)

type AlipayService struct {

}

func (o AlipayService) Pay() (string, error){
	return payment.GetInstance().Pay()
}

func (o AlipayService) ParseNotify(c *gin.Context) error {
	return payment.GetInstance().ParseNotify(c)
}
