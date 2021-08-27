package entity

const (
	// 所有
	PAY_STATUS_ALL = -1
	// 待支付
	PAY_STATUS_WAIT = 0
	// 支付失败
	PAY_STATUS_FAIL = 1
	// 支付成功
	PAY_STATUS_SUCC = 2
)

type SubmitPayReq struct {
	OrderId int     `json:"id"`
	Amount  float64 `json:"amount"`
	OrderSn string  `json:"sn"`
}

type NotifyPayResp struct {
	OrderSn string `json:"sn"`
	Status  int    `json:"status"`
	Desc    string `json:"desc"`
}
