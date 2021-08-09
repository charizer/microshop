package entity

type HomeResponse struct {
	Banners  []Banner      `json:"banner"`
	Newgoods []SampleGoods `json:"newGoodsList"`
	Hotgoods []SampleGoods `json:"hotGoodsList"`
}
