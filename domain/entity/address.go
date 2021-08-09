package entity

type Address struct {
	Address      string `json:"address"`
	Id           int    `json:"id"`
	IsDefault    *bool  `json:"is_default"`
	Mobile       string `json:"mobile"`
	Name         string `json:"name"`
	UserId       int    `json:"user_id"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}
