package models

type MenuImageModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuID"`
	BannerID    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:bannerID"`
	Sort        int         `gorm:"size:10"json:"sort"`
}
