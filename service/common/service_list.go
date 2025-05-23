package common

import (
	"gvb_server/global"
	"gvb_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (List []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	query := DB.Where(model)
	count = query.Select("id").Find(&List).RowsAffected
	//这里query会受到上面查询的影响
	query = DB.Where(model)
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = global.DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&List).Error
	return List, count, err
}
