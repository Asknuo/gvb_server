package models

type AdvertModel struct {
	MODEL
	Title  string `gorm:size:32" json:"title"` //显示的标题
	Herf   string `json:"herf"`                //跳转到的链接
	Images string `json:"images"`              //图片
	IsShow bool   `json:"is_show"`             //是否展示
}
