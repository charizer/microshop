package entity

type Cart struct {
	Checked     int     `json:"checked"`
	GoodsId     int     `json:"goods_id"`
	GoodsName   string  `json:"goods_name"`
	Id          int     `json:"id"`
	ListPicUrl  string  `json:"list_pic_url"`
	Number      int     `json:"number"`
	ProductId   int     `json:"product_id"`
	RetailPrice float64 `json:"retail_price"`
	UserId      int     `json:"user_id"`
	GoodsBrief  string  `json:"goods_brief"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

type CartAddReq struct {
	GoodsId   int `json:"goodsId"`
	ProductId int `json:"productId"`
	Number    int `json:"number"`
}

type CartTotal struct {
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float64 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float64 `json:"checkedGoodsAmount"`
}

type CartListResponse struct {
	CartList  []Cart    `json:"cartList"`
	CartTotal CartTotal `json:"cartTotal"`
}

type CartUpdateReq struct {
	GoodsId   int `json:"goodsId"`
	ProductId int `json:"productId"`
	Number    int `json:"number"`
	Id        int `json:"id"`
}

type CartDeleteReq struct {
	Id int `json:"id"`
}

type CartCheckedReq struct {
	IsChecked int   `json:"isChecked"`
	Carts     []int `json:"carts"`
}

type CheckOutCartResponse struct {
	Address          Address      `json:"checkedAddress"`
	CheckedGoodsList []Cart       `json:"checkedGoodsList"`
	OrderTotalPrice  float64      `json:"orderTotalPrice"`
}
