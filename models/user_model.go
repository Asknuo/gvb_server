package models

import (
	"gvb_server/models/ctype"
)

type UserModel struct {
	MODEL
	NickName      string           `gorm:"size:42" json:"nick_name"`
	UserName      string           `gorm:"size:36" json:"user_name"`
	Password      string           `gorm:"size:64" json:"password"`
	Avatar        string           `gorm:"size:256" json:"avator"`
	Email         string           `json:"email"`
	Tel           string           `gorm:"size:18" json:"telephone"`
	Addr          string           `gorm:"size:64" json:"addr"`
	Token         string           `gorm:"size:64" json:"token"`
	IP            string           `gorm:"size:20" json:"ip"`
	Role          ctype.Role       `gorm:"size:4;default:1;" json:"role"`
	SignStatus    ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
	ArticleModels []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`                                                            // 修改为 UserID（大写 D）
	CollectModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` // 收藏文章列表
}
