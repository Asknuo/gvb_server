package models

import (
	"gvb_server/models/ctype"
)

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`           // 添加空格
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`        // 添加空格
	Slogan       string        `gorm:"size:64" json:"slogan"`               // 添加空格
	Abstract     ctype.Array   `gorm:"type:string;size:64" json:"abstract"` // 添加 gorm 标签
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `gorm:"size:10" json:"sort"` // 添加空格
}
