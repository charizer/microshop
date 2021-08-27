package entity

type Goods struct {
	CategoryId    int     `json:"category_id"`
	GoodsBrief    string  `json:"goods_brief"`
	GoodsDesc     string  `json:"goods_desc"`
	GoodsSn       string  `json:"goods_sn"`
	Id            int     `json:"id"`
	IsDelete      bool    `json:"is_delete"`
	IsHot         int     `json:"is_hot"`
	IsLimited     int     `json:"is_limited"`
	IsNew         int     `json:"is_new"`
	IsOnSale      int     `json:"is_on_sale"`
	ListPicUrl    string  `json:"list_pic_url"`
	Name          string  `json:"name"`
	PrimaryPicUrl string  `json:"primary_pic_url"`
	RetailPrice   float64 `json:"retail_price"`
	SortOrder     int     `json:"sort_order"`
	CreateTime    int64   `json:"create_time"`
	UpdateTime    int64   `json:"update_time"`
}

type SampleGoods struct {
	Id            int     `json:"id"`
	CategoryId    int     `json:"category_id"`
	ListPicUrl    string  `json:"list_pic_url"`
	Name          string  `json:"name"`
	RetailPrice   float64 `json:"retail_price"`
	GoodsBrief    string  `json:"goods_brief"`
	PrimaryPicUrl string  `json:"primary_pic_url"`
}

type GoodsGallery struct {
	GoodsId    int    `json:"goods_id"`
	Id         int    `json:"id"`
	ImgDesc    string `json:"img_desc"`
	ImgUrl     string `json:"img_url"`
	ImgType    int    `json:"img_type"`
	SortOrder  int    `json:"sort_order"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type GoodsAttr struct {
	GoodsId    int    `json:"goods_id"`
	Id         int    `json:"id"`
	Value      string `json:"value"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type GoodsProduct struct {
	GoodsId     int     `json:"goods_id"`
	GoodsNumber int     `json:"goods_number"`
	Id          int     `json:"id"`
	RetailPrice float64 `json:"retail_price"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

type GoodsListResponse struct {
	GoodsList []SampleGoods `json:"goodsList"`
}

type GoodsDetailResponse struct {
	Goods       Goods          `json:"info"`
	Galleries   []GoodsGallery `json:"gallery"`
	Attribute   []GoodsAttr    `json:"attribute"`
	ProductList []GoodsProduct `json:"productList"`
	ImageText   []GoodsGallery `json:"imageText"`
}
