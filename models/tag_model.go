package models

type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`                  //标签名称
	Articles []ArticleModel `gorm:"many2many:article_tag_models" json:"-"` //关联读标签的文章列表
}
