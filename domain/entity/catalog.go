package entity

type Catalog struct {
	Id           int    `json:"id"`
	ImgUrl       string `json:"img_url"`
	IsShow       int    `json:"is_show"`
	Name         string `json:"name"`
	ParentId     int    `json:"parent_id"`
	SortOrder    int    `json:"sort_order"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type CataListResponse struct {
	CategoryList    []Catalog `json:"categoryList"`
}
