package entity

type Banner struct {
	Id        int    `json:"id"`
	Content   string `json:"content"`
	Enabled   int    `json:"enabled"`
	ImageUrl  string `json:"image_url"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}
