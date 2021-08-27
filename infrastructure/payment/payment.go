package payment

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"microshop/domain/entity"
	"microshop/infrastructure/config"
	"microshop/infrastructure/logger"
	"net/url"
	"strconv"
	"sync"
)

var (
	log            = logger.GetLogger()
	cfg 		   = config.GetConfig()
	aliPayment     *AliPayment
	onceAliPayment sync.Once
)

type AliPayment struct {
	aliClient *alipay.Client
}

func init(){
	GetInstance()
}

func GetInstance() *AliPayment {
	onceAliPayment.Do(func() {
		aliPayment = newAliPayment(cfg.AliPayAppId, cfg.AliPayPrivateKey, cfg.AliPayPublicKey)
	})
	return aliPayment
}

func newAliPayment(appId, privateKey, publicKey string) *AliPayment{
	aliClient, err := alipay.New(appId, privateKey, true)
	if err != nil {
		logger.GetLogger().Errorf("init alipay err:%s", err.Error())
		panic(err)
	}
	err = aliClient.LoadAliPayPublicKey(publicKey)
	if err != nil {
		logger.GetLogger().Errorf("load alipay public key err:%s", err.Error())
		panic(err)
	}
	return &AliPayment{aliClient: aliClient}
}

func (o *AliPayment) Pay(ctx context.Context, req entity.SubmitPayReq) (string, error){
	var p = alipay.TradeAppPay{}
	p.NotifyURL = cfg.AliPayNotifyUrl
	p.Subject = "app 支付"
	p.OutTradeNo = req.OrderSn
	p.TotalAmount = strconv.FormatFloat(req.Amount, 'f', 2, 64)
	log.Infof("pay detail orderId:%d orderSn:%s amount:%f", req.OrderId, req.OrderSn, req.Amount)
	values := url.Values{} //回调参数，这里只能这样写，要进行urlEncode才能传给支付宝
	//需要回传的参数
	values.Add("orderId", strconv.Itoa(req.OrderId))
	p.PassbackParams = values.Encode()
	var url, err = o.aliClient.TradeAppPay(p)
	if err != nil {
		logger.GetLogger().Errorf("pay err:%s", err.Error())
	}
	logger.GetLogger().Infof("trade pay return url:", url)
	return url, err
}

func (o *AliPayment) ParseNotify(c *gin.Context) (entity.NotifyPayResp,error) {
	c.Request.ParseForm()
	errMsg := ""
	var outTradeNo = c.Request.Form.Get("out_trade_no")
	log.Infof("订单 %s 支付 详细信息: %+v", outTradeNo, c.Request.Form)
	ok, err := o.aliClient.VerifySign(c.Request.Form)
	if err != nil {
		errMsg = fmt.Sprintf("sn:%s 异步通知验证签名发生错误:%s", outTradeNo, err.Error())
		log.Infoln(errMsg)
		return entity.NotifyPayResp{
			OrderSn: outTradeNo,
			Status: entity.PAY_STATUS_FAIL,
			Desc: errMsg,
		}, err
	}
	if ok == false {
		errMsg = fmt.Sprintf("sn:%s 异步通知验证签名未通过", outTradeNo)
		log.Infoln(errMsg)
		return entity.NotifyPayResp{
			OrderSn: outTradeNo,
			Status: entity.PAY_STATUS_FAIL,
			Desc: errMsg,
		}, errors.New("异步通知验证签名未通过")
	}
	log.Infoln("异步通知验证签名通过")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := o.aliClient.TradeQuery(p)
	if err != nil {
		errMsg = fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		log.Infoln(errMsg)
		return entity.NotifyPayResp{
			OrderSn: outTradeNo,
			Status: entity.PAY_STATUS_FAIL,
			Desc: errMsg,
		}, err
	}
	if rsp.IsSuccess() == false {
		errMsg = fmt.Sprintf("异步通知验证订单不通过 %s 信息发生错误: %s-%s", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		log.Infoln(errMsg)
		return entity.NotifyPayResp{
			OrderSn: outTradeNo,
			Status: entity.PAY_STATUS_FAIL,
			Desc: errMsg,
		}, errors.New("异步通知信息发生错误")
	}
	msg := fmt.Sprintf("订单 %s 支付成功 详细信息: %+v", outTradeNo, c.Request.Form)
	return entity.NotifyPayResp{
		OrderSn: outTradeNo,
		Status: entity.PAY_STATUS_SUCC,
		Desc: msg,
	},nil
}


