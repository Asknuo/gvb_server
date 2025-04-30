package models

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"` //发送人id
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"sen_user_avator"`
	RevUserID        uint      `gorm:"primaryKey" json:"rev_user_id"`
	RevUserModel     UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickName  string    `gorm:"size:42" json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avator"`
	IsRead           bool      `gorm:"default:false" json:"is_read"` //接收方是否查看
	Content          string    `json:"content"`                      //消息内容
}
