package models

type MenuBannerModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BannerID"` // 修改为 BannerID（大写 ID）
	Sort        int         `gorm:"size:10" json:"sort"` // 修正 json 标签格式
}
