package entity

const (
	ORDER_ALL       = -1
	ORDER_WAIT_PAY  = 0
	ORDER_WAIT_RECV = 1
	ORDER_CANCEL    = 2
	ORDER_SUCC      = 3
	ORDER_FINISH    = 4
	ORDER_DELETE    = 5
)

type SubmitOrderReq struct {
	AddressId int `json:"addressId"`
	GoodsId   int `json:"goodsId"`
	Number    int `json:"number"`
	ProductId int `json:"productId"`
}

type Order struct {
	ActualPrice    float64 `json:"actual_price"`
	Address        string  `json:"address"`
	CallbackStatus string  `json:"callback_status"`
	ConfirmTime    int     `json:"confirm_time"`
	Consignee      string  `json:"consignee"`
	FreightPrice   float64 `json:"freight_price"`
	Id             int     `json:"id"`
	Mobile         string  `json:"mobile"`
	OrderPrice     float64 `json:"order_price"`
	OrderSn        string  `json:"order_sn"`
	OrderStatus    int     `json:"order_status"`
	PayId          int     `json:"pay_id"`
	PayName        string  `json:"pay_name"`
	PayStatus      int     `json:"pay_status"`
	PayTime        int     `json:"pay_time"`
	UserId         int     `json:"user_id"`
	ProvinceName   string  `json:"province_name"`
	CityName       string  `json:"city_name"`
	DistrictName   string  `json:"district_name"`
	CreateTime     int64   `json:"create_time"`
	UpdateTime     int64   `json:"update_time"`
}

type OrderGoods struct {
	GoodsId     int     `json:"goods_id"`
	GoodsName   string  `json:"goods_name"`
	Id          int     `json:"id"`
	ListPicUrl  string  `json:"list_pic_url"`
	Number      int     `json:"number"`
	OrderId     int     `json:"order_id"`
	ProductId   int     `json:"product_id"`
	RetailPrice float64 `json:"retail_price"`
	GoodsBrief  string  `json:"goods_brief"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

type OrderInfo struct {
	Order
	GoodsList       []OrderGoods `json:"goodList"`
	GoodsCount      int          `json:"goodsCount"`
	OrderStatusText string       `json:"order_status_text"`
	OrderDate       string       `json:"order_date"`
}

type UpdateOrderStatusReq struct {
	OrderId int `json:"orderId"`
	Status  int `json:"status"`
}

