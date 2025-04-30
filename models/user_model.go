package models

import (
	"gvb_server/models/ctype"
)

type UserModel struct {
	MODEL
	NickName      string           `grom : "size:42" json:"nick_name"`                                                      //昵称
	UserName      string           `grom : "size:36" json:"user_name"`                                                      //用户名
	Password      string           `grom : "size:64" json:"password"`                                                       //密码
	Avatar        string           `grom : "size:256" json:"avator"`                                                        //头像
	Email         string           `json: "email"`                                                                          //邮箱
	Tel           string           `grom : "size:18" json:"telephone"`                                                      //电话号码
	addr          string           `grom : "size:64" json:"addr"`                                                           //地址
	Token         string           `grom : "size:64" json:"token"`                                                          //其他平台唯一的id
	IP            string           `grom : "size:20" json:"ip"`                                                             //ip地址
	Role          ctype.Role       `grom : "size:4;default:1;" json:"role"`                                                 //权限
	SignStatus    ctype.SignStatus `grom : "type = smallint(6)" json:"sign_status"`                                         //注册来源
	ArticleModels []ArticleModel   `gorm:"foreignKey:AuthId" json:"-"`                                                      //发布文章列表
	CollectModels []ArticleModel   `gorm:"many2many:auth2_collect;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"-"` //收藏文章列表
}
