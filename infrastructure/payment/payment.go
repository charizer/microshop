package payment

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"microshop/common"
	"microshop/infrastructure/config"
	"microshop/infrastructure/logger"
	"net/url"
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
	return &AliPayment{aliClient: aliClient,}
}

func (o *AliPayment) Pay() (string, error){
	var p = alipay.TradeAppPay{}
	p.NotifyURL = cfg.AliPayNotifyUrl
	p.Subject = "标题"
	p.OutTradeNo = common.GenerateOrderNumber()
	p.TotalAmount = "0.01"
	values := url.Values{} //回调参数，这里只能这样写，要进行urlEncode才能传给支付宝
	//需要回传的参数
	values.Add("aaa", "aaa")
	values.Add("bbb", "bbb")
	p.PassbackParams = values.Encode()
	var url, err = o.aliClient.TradeAppPay(p)
	if err != nil {
		logger.GetLogger().Errorf("pay err:%s", err.Error())
	}
	logger.GetLogger().Infof("trade pay return url:", url)
	return url, err
}

func (o *AliPayment) ParseNotify(c *gin.Context) error {
	c.Request.ParseForm()
	ok, err := o.aliClient.VerifySign(c.Request.Form)
	if err != nil {
		log.Infoln("异步通知验证签名发生错误", err)
		return err
	}
	if ok == false {
		log.Infoln("异步通知验证签名未通过")
		return errors.New("异步通知验证签名未通过")
	}
	log.Infoln("异步通知验证签名通过")
	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := o.aliClient.TradeQuery(p)
	if err != nil {
		log.Infof("异步通知验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		return err
	}
	if rsp.IsSuccess() == false {
		log.Infof("异步通知验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		return errors.New("异步通知信息发生错误")
	}

	log.Infof("订单 %s 支付成功 详细信息: %+v", outTradeNo, c.Request.Form)
	return nil
}
