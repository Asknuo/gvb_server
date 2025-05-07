package models

import (
	"gvb_server/models/ctype"

	"gorm.io/gorm"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片的哈希值，判断是否重复
	Name      string          `gorm:"size:38" json:"name"`         //图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"iamge_type"` //图片类型
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {

	return
}
