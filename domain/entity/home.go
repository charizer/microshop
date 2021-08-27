package entity

type HomeResponse struct {
	Banners  []SampleGoods `json:"banner"`
	Newgoods []SampleGoods `json:"newGoodsList"`
	Hotgoods []SampleGoods `json:"hotGoodsList"`
}
