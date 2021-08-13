package common

import "github.com/gin-gonic/gin"

func ErrorResponse(errcode, error, detail string) gin.H {
	return gin.H{
		"status":  errcode,
		"msg":     error,
		"data":    detail,
		"service": "microshop",
	}
}

const (
	C_REQ_PARAM_ERR = iota + 1
	C_MOBILE_ERR
	C_SEND_MUCH_ERR
	C_VERIFY_CODE_ERR
	C_ADDR_NOT_EXSIT_ERR
	C_GOODS_NOT_EXSIT_ERR
	C_PRODUCT_NOT_EXSIT_ERR
	C_CART_NOT_EXSIT_ERR
	C_CART_EMPTY_ERR
	C_ORDER_NOT_EXSIT_ERR
	S_MYSQL_ERR
	S_SEND_SMS_ERR
	S_ALIPAY_REQ_ERR
)

var errorCode = map[int]string{
	C_REQ_PARAM_ERR:         "C_REQ_PARAM_ERR",
	C_MOBILE_ERR:            "C_MOBILE_ERR",
	C_SEND_MUCH_ERR:         "C_SEND_MUCH_ERR",
	C_VERIFY_CODE_ERR:       "C_VERIFY_CODE_ERR",
	C_ADDR_NOT_EXSIT_ERR:    "C_ADDR_NOT_EXSIT_ERR",
	C_GOODS_NOT_EXSIT_ERR:   "C_GOODS_NOT_EXSIT_ERR",
	C_PRODUCT_NOT_EXSIT_ERR: "C_PRODUCT_NOT_EXSIT_ERR",
	C_CART_NOT_EXSIT_ERR:    "C_CART_NOT_EXSIT_ERR",
	C_CART_EMPTY_ERR:        "C_CART_EMPTY_ERR",
	C_ORDER_NOT_EXSIT_ERR:   "C_ORDER_NOT_EXSIT_ERR",
	S_MYSQL_ERR:             "S_MYSQL_ERR",
	S_SEND_SMS_ERR:          "S_SEND_SMS_ERR",
	S_ALIPAY_REQ_ERR:		"S_ALIPAY_REQ_ERR",
}

var errorText = map[int]string{
	C_REQ_PARAM_ERR:         "请求参数错误",
	C_MOBILE_ERR:            "手机号码格式错误",
	C_SEND_MUCH_ERR:         "发送太频繁",
	C_VERIFY_CODE_ERR:       "验证码不对",
	C_ADDR_NOT_EXSIT_ERR:    "修改的地址不存在",
	C_GOODS_NOT_EXSIT_ERR:   "商品不存在",
	C_PRODUCT_NOT_EXSIT_ERR: "商品库存不足",
	C_CART_NOT_EXSIT_ERR:    "购物车条目不存在",
	C_CART_EMPTY_ERR:        "购物车没有商品",
	C_ORDER_NOT_EXSIT_ERR:   "订单不存在",
	S_MYSQL_ERR:             "mysql错误",
	S_SEND_SMS_ERR:          "短信发送失败",
	S_ALIPAY_REQ_ERR: "支付请求失败",
}

func ErrorCode(code int) string {
	return errorCode[code]
}

func ErrorText(code int) string {
	return errorText[code]
}
